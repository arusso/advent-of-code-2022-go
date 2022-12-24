package main

import (
	"fmt"
	"testing"
)

func TestPathString(t *testing.T) {
	tests := map[string]string{
		"/foo/bar":  "/foo/bar",
		"/":         "/",
		"/foo/bar/": "/foo/bar",
	}

	for input, output := range tests {
		t.Run(fmt.Sprintf("input='%s'", input), func(t *testing.T) {
			p := NewPath(input)
			if p.String() != output {
				t.Errorf("Expected path %s, got %s", output, p.String())
			}
		})
	}
}

func TestPathPushNPop(t *testing.T) {
	path := NewPath("/")
	path = path.
		Push("foo").
		Push("bar/baz").
		Pop()

	if path.String() != "/foo/bar" {
		t.Errorf("Expected '/foo/bar/', got '%s'", path.String())
	}

	path = path.Push("z")
	if path.String() != "/foo/bar/z" {
		t.Errorf("Expected '/foo/bar/z', got '%s'", path.String())
	}
}
