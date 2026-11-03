// Package toolconfig provides the shared plumbing which allows the road map rendering tools
// to discover their input files using glob patterns and to read their options from the
// `config` block of each road map file.
package toolconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/SierraSoftworks/roadmap"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v3"
)

// ProfileFlag returns the `--profile` flag used to select which entry in a road map's
// `config` block a tool should read its options from. By default, tools use their own name.
func ProfileFlag(defaultProfile string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "profile",
		Aliases: []string{"p"},
		Usage:   "The `NAME` of the entry in the roadmap's config block to read options from.",
		Value:   defaultProfile,
	}
}

// Options resolves the value of a tool's options for a specific road map file. Flags which
// are explicitly set on the command line take precedence over values in the road map's config
// profile, which in turn take precedence over the flag's default value.
type Options struct {
	c       *cli.Context
	profile roadmap.ToolConfig
	dir     string
}

// NewOptions creates a new set of options for a road map loaded from the provided file.
func NewOptions(c *cli.Context, r *roadmap.Roadmap, file string) *Options {
	return &Options{
		c:       c,
		profile: r.Profile(c.String("profile")),
		dir:     filepath.Dir(file),
	}
}

// String returns the value of the named option as a string. Structured config values (such
// as a `frontmatter` mapping) are returned in their YAML representation.
func (o *Options) String(name string) string {
	if o.c.IsSet(name) {
		return o.c.String(name)
	}

	v, ok := o.profile[name]
	if !ok || v == nil {
		return o.c.String(name)
	}

	if s, ok := v.(string); ok {
		return s
	}

	out, err := yaml.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}

	return strings.TrimSpace(string(out))
}

// Bool returns the value of the named option as a boolean.
func (o *Options) Bool(name string) bool {
	if o.c.IsSet(name) {
		return o.c.Bool(name)
	}

	v, ok := o.profile[name]
	if !ok || v == nil {
		return o.c.Bool(name)
	}

	if b, ok := v.(bool); ok {
		return b
	}

	b, err := strconv.ParseBool(o.String(name))
	if err != nil {
		return o.c.Bool(name)
	}

	return b
}

// Path returns the value of the named option as a file path. Paths provided in the road map's
// config block are resolved relative to the directory containing the road map file, while
// paths provided on the command line are used as-is.
func (o *Options) Path(name string) string {
	p := o.String(name)
	if p == "" || o.c.IsSet(name) || filepath.IsAbs(p) {
		return p
	}

	if _, ok := o.profile[name]; !ok {
		return p
	}

	return filepath.Join(o.dir, filepath.FromSlash(p))
}

// Render will discover every road map file matching the `--input` pattern, parse it and
// pass it to the provided render function before writing the result to the resolved
// `output` path (or stdout, if no output has been configured).
func Render(c *cli.Context, render func(r *roadmap.Roadmap, opts *Options) (string, error)) error {
	files, err := roadmap.Discover(c.String("input"))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("no roadmap files matched the input pattern %q", c.String("input"))
	}

	if len(files) > 1 && c.IsSet("output") {
		return fmt.Errorf("the --output flag cannot be used when the input pattern %q matches multiple files, configure the output in each roadmap's config block instead", c.String("input"))
	}

	for _, file := range files {
		f, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		r, err := roadmap.Parse(f)
		if err != nil {
			return errors.Wrapf(err, "%s", file)
		}

		opts := NewOptions(c, r, file)
		out, err := render(r, opts)
		if err != nil {
			return errors.Wrapf(err, "%s", file)
		}

		if output := opts.Path("output"); output != "" {
			if err := os.WriteFile(output, []byte(out), os.ModePerm); err != nil {
				return err
			}
		} else if _, err := fmt.Println(out); err != nil {
			return err
		}
	}

	return nil
}
