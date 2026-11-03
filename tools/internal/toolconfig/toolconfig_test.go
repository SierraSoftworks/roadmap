package toolconfig_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/SierraSoftworks/roadmap"
	"github.com/SierraSoftworks/roadmap/tools/internal/toolconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

// testApp builds a minimal tool which renders the resolved options for each road map,
// allowing us to verify how flags and config profiles are combined.
func testApp() *cli.App {
	return &cli.App{
		Name: "test",
		Flags: []cli.Flag{
			toolconfig.ProfileFlag("test"),
			&cli.StringFlag{Name: "input", Value: "roadmap.yml"},
			&cli.StringFlag{Name: "output"},
			&cli.StringFlag{Name: "template", Value: "default.tmpl"},
			&cli.BoolFlag{Name: "simple"},
			&cli.BoolFlag{Name: "collapsed"},
		},
		Action: func(c *cli.Context) error {
			return toolconfig.Render(c, func(r *roadmap.Roadmap, opts *toolconfig.Options) (string, error) {
				return fmt.Sprintf("%s simple=%v collapsed=%v template=%s", r.Title, opts.Bool("simple"), opts.Bool("collapsed"), opts.String("template")), nil
			})
		},
	}
}

func writeFile(t *testing.T, path, content string) {
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
}

func readFile(t *testing.T, path string) string {
	content, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(content)
}

func TestRenderUsesProfile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "a", "roadmap.yml"), `
title: A
config:
  test:
    output: out/a.md
    simple: true
  other:
    output: other.md
    collapsed: true
`)
	writeFile(t, filepath.Join(dir, "b", "roadmap.yml"), `
title: B
config:
  test:
    output: b.md
    template: custom.tmpl
`)
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "a", "out"), 0o755))

	require.NoError(t, testApp().Run([]string{"test", "--input", filepath.ToSlash(dir) + "/**/roadmap.yml"}))

	assert.Equal(t, "A simple=true collapsed=false template=default.tmpl", readFile(t, filepath.Join(dir, "a", "out", "a.md")), "outputs should be relative to the road map file")
	assert.Equal(t, "B simple=false collapsed=false template=custom.tmpl", readFile(t, filepath.Join(dir, "b", "b.md")))

	require.NoError(t, testApp().Run([]string{"test", "--profile", "other", "--input", filepath.Join(dir, "a", "roadmap.yml")}))
	assert.Equal(t, "A simple=false collapsed=true template=default.tmpl", readFile(t, filepath.Join(dir, "a", "other.md")), "the selected profile should be used")
}

func TestRenderFlagsOverrideProfile(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "roadmap.yml"), `
title: A
config:
  test:
    output: config.md
    simple: true
`)

	output := filepath.Join(dir, "flag.md")
	require.NoError(t, testApp().Run([]string{"test", "--input", filepath.Join(dir, "roadmap.yml"), "--output", output, "--simple=false", "--collapsed"}))

	assert.Equal(t, "A simple=false collapsed=true template=default.tmpl", readFile(t, output))
	assert.NoFileExists(t, filepath.Join(dir, "config.md"), "the configured output should be overridden")
}

func TestRenderOutputWithMultipleInputs(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "a", "roadmap.yml"), "title: A")
	writeFile(t, filepath.Join(dir, "b", "roadmap.yml"), "title: B")

	err := testApp().Run([]string{"test", "--input", filepath.ToSlash(dir) + "/*/roadmap.yml", "--output", filepath.Join(dir, "out.md")})
	assert.Error(t, err, "a single output cannot be used for multiple inputs")
}

func TestRenderNoMatches(t *testing.T) {
	dir := t.TempDir()

	err := testApp().Run([]string{"test", "--input", filepath.ToSlash(dir) + "/**/roadmap.yml"})
	assert.ErrorContains(t, err, "no roadmap files matched")
}
