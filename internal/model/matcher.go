package model

import (
	"fmt"
	"path/filepath"
)

type PathMatcher interface {
	Matches(f string, isDir bool) (bool, error)
}

// A Matcher that matches nothing.
type emptyMatcher struct{}

func (m emptyMatcher) Matches(f string, isDir bool) (bool, error) {
	return false, nil
}

var EmptyMatcher PathMatcher = emptyMatcher{}

// A matcher that matches one file.
type fileMatcher struct {
	path string
}

func (m fileMatcher) Matches(f string, isDir bool) (bool, error) {
	return f == m.path, nil
}

func NewSimpleFileMatcher(f string) (fileMatcher, error) {
	// Get the absolute path of the file, because PathMatchers expect to always
	// work with absolute paths.
	f, err := filepath.Abs(f)
	if err != nil {
		return fileMatcher{}, fmt.Errorf("NewSimpleFileMatcher: %v", err)
	}
	return fileMatcher{path: f}, nil
}

type PatternMatcher interface {
	PathMatcher

	// Express this PathMatcher as a sequence of filepath.Match
	// patterns. These patterns are widely useful in Docker-land because
	// they're suitable in .dockerignore or Dockerfile ADD statements
	// https://docs.docker.com/engine/reference/builder/#add
	AsMatchPatterns() []string
}

type CompositePathMatcher struct {
	Matchers []PathMatcher
}

func NewCompositeMatcher(matchers []PathMatcher) PathMatcher {
	cMatcher := CompositePathMatcher{Matchers: matchers}
	pMatchers := make([]PatternMatcher, len(matchers))
	for i, m := range matchers {
		pm, ok := m.(CompositePatternMatcher)
		if !ok {
			return cMatcher
		}
		pMatchers[i] = pm
	}
	return CompositePatternMatcher{
		CompositePathMatcher: cMatcher,
		Matchers:             pMatchers,
	}
}

func (c CompositePathMatcher) Matches(f string, isDir bool) (bool, error) {
	for _, t := range c.Matchers {
		ret, err := t.Matches(f, isDir)
		if err != nil {
			return false, err
		}
		if ret {
			return true, nil
		}
	}
	return false, nil
}

type CompositePatternMatcher struct {
	CompositePathMatcher
	Matchers []PatternMatcher
}

func (c CompositePatternMatcher) AsMatchPatterns() []string {
	result := []string{}
	for _, m := range c.Matchers {
		result = append(result, m.AsMatchPatterns()...)
	}
	return result
}

var _ PathMatcher = CompositePathMatcher{}
var _ PatternMatcher = CompositePatternMatcher{}
