package roadmap

import (
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Discover will expand the provided glob pattern into the list of road map files which
// it matches. In addition to the standard `*`, `?` and `[...]` wildcards supported by
// path.Match, a `**` path segment will match zero or more directories, allowing patterns
// like `**/roadmap.yml` to find road maps anywhere within a repository.
//
// Patterns which do not contain any wildcards are returned as-is (without checking that
// the file exists) so that callers report the usual "file not found" error for them.
// Hidden directories (those starting with a `.`) are not searched by wildcards.
func Discover(pattern string) ([]string, error) {
	pattern = filepath.ToSlash(pattern)
	if !hasMeta(pattern) {
		return []string{filepath.FromSlash(pattern)}, nil
	}

	segments := strings.Split(pattern, "/")
	for _, s := range segments {
		if _, err := path.Match(s, ""); err != nil {
			return nil, errors.Wrapf(err, "roadmap: invalid input pattern %q", pattern)
		}
	}

	// Walk from the longest leading portion of the pattern which contains no wildcards,
	// so that `docs/**/roadmap.yml` doesn't need to search the whole working directory.
	static := 0
	for static < len(segments)-1 && !hasMeta(segments[static]) {
		static++
	}

	base := strings.Join(segments[:static], "/")
	if base == "" {
		base = "."
		if strings.HasPrefix(pattern, "/") {
			base = "/"
		}
	}
	remaining := segments[static:]

	matches := []string{}
	err := filepath.WalkDir(filepath.FromSlash(base), func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(filepath.FromSlash(base), p)
		if err != nil {
			return err
		}

		if d.IsDir() {
			if rel != "." && strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		if matchSegments(remaining, strings.Split(filepath.ToSlash(rel), "/")) {
			matches = append(matches, p)
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "roadmap: unable to search for files matching %q", pattern)
	}

	return matches, nil
}

func hasMeta(pattern string) bool {
	return strings.ContainsAny(pattern, "*?[")
}

// matchSegments determines whether the provided path segments match the pattern segments,
// where a `**` pattern segment matches zero or more path segments.
func matchSegments(pattern, name []string) bool {
	if len(pattern) == 0 {
		return len(name) == 0
	}

	if pattern[0] == "**" {
		return matchSegments(pattern[1:], name) || (len(name) > 0 && matchSegments(pattern, name[1:]))
	}

	if len(name) == 0 {
		return false
	}

	// The pattern has already been validated, so we can safely ignore errors here.
	if ok, _ := path.Match(pattern[0], name[0]); !ok {
		return false
	}

	return matchSegments(pattern[1:], name[1:])
}
