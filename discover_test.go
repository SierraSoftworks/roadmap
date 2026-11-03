package roadmap_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/SierraSoftworks/roadmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createDiscoveryTree(t *testing.T) string {
	dir := t.TempDir()

	for _, f := range []string{
		"roadmap.yml",
		"other.yml",
		"docs/roadmap.yml",
		"docs/team/roadmap.yml",
		"docs/team/roadmap.yaml",
		".hidden/roadmap.yml",
	} {
		p := filepath.Join(dir, filepath.FromSlash(f))
		require.NoError(t, os.MkdirAll(filepath.Dir(p), 0o755))
		require.NoError(t, os.WriteFile(p, []byte("title: Test"), 0o644))
	}

	return dir
}

func TestDiscover(t *testing.T) {
	dir := createDiscoveryTree(t)
	root := filepath.ToSlash(dir)

	cases := []struct {
		pattern  string
		expected []string
	}{
		{"**/roadmap.yml", []string{"docs/roadmap.yml", "docs/team/roadmap.yml", "roadmap.yml"}},
		{"*.yml", []string{"other.yml", "roadmap.yml"}},
		{"docs/**/roadmap.y*ml", []string{"docs/roadmap.yml", "docs/team/roadmap.yaml", "docs/team/roadmap.yml"}},
		{"docs/*/roadmap.yml", []string{"docs/team/roadmap.yml"}},
		{"**/missing.yml", []string{}},
		{".hidden/*.yml", []string{".hidden/roadmap.yml"}},
	}

	for _, c := range cases {
		t.Run(c.pattern, func(t *testing.T) {
			matches, err := roadmap.Discover(root + "/" + c.pattern)
			require.NoError(t, err)

			expected := make([]string, len(c.expected))
			for i, e := range c.expected {
				expected[i] = filepath.Join(dir, filepath.FromSlash(e))
			}

			assert.Equal(t, expected, matches)
		})
	}
}

func TestDiscoverRelative(t *testing.T) {
	dir := createDiscoveryTree(t)
	t.Chdir(dir)

	matches, err := roadmap.Discover("**/roadmap.yml")
	require.NoError(t, err)
	assert.Equal(t, []string{
		filepath.FromSlash("docs/roadmap.yml"),
		filepath.FromSlash("docs/team/roadmap.yml"),
		"roadmap.yml",
	}, matches)
}

func TestDiscoverLiteral(t *testing.T) {
	matches, err := roadmap.Discover("path/to/missing.yml")
	require.NoError(t, err)
	assert.Equal(t, []string{filepath.FromSlash("path/to/missing.yml")}, matches, "literal paths should be returned as-is")
}

func TestDiscoverInvalidPattern(t *testing.T) {
	_, err := roadmap.Discover("docs/[/*.yml")
	assert.Error(t, err)
}
