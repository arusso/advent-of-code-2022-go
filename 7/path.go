package main

import "strings"

// NewPath creates a new Path from a given path string.
func NewPath(p string) Path {
	c := cleanPath(p)
	np := Path{length: len(c), segments: c}
	return np
}

// Path represents a filesystem path
type Path struct {
	length   int
	segments []string
}

func (p Path) String() string {
	builder := strings.Builder{}
	builder.WriteString("/" + strings.Join(p.segments, "/"))
	return builder.String()
}

// Shift off a segment from the base of the path. If path is empty, returns an
// empty path and empty segment.
func (p Path) Shift() (Path, string) {
	if p.length == 0 {
		return Path{}, ""
	}

	newlen := len(p.segments) - 1
	newsegments := make([]string, newlen)
	copy(newsegments, p.segments[1:])
	return Path{length: newlen, segments: newsegments}, p.segments[0]
}

// PopOff the last path segment returning the resulting path and popped segment
// If no segments left, an empty path is returned with an empty string segment.
func (p Path) PopOff() (Path, string) {
	if p.length == 0 {
		return Path{length: 0, segments: []string{}}, ""
	}
	newlen := len(p.segments) - 1
	newsegments := make([]string, newlen)
	copy(newsegments, p.segments[:newlen])
	return Path{length: newlen, segments: newsegments}, p.segments[newlen]
}

// Pop the last path segment off return the resulting path.
func (p Path) Pop() Path {
	np, _ := p.PopOff()
	return np
}

// Push a path string onto our path. Empty segments will be ignored.
func (p Path) Push(str string) Path {
	newc := cleanPath(str)
	return Path{
		length:   p.length + len(newc),
		segments: append(p.segments, newc...),
	}
}

// cleanPath takes a path string and returns the segments of that path,
// removing empty segments.
func cleanPath(path string) []string {
	newsegments := []string{}
	for _, val := range strings.Split(path, "/") {
		if val == "" {
			continue
		}
		newsegments = append(newsegments, val)
	}

	return newsegments
}

// Base returns the base path
func (p Path) Base() Path {
	return p.Pop()
}

func (p Path) Segments() []string {
	return p.segments
}
