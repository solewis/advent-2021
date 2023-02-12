package main

import (
	"advent-2020/internal/parse"
	"reflect"
	"strings"
	"testing"
)

func TestParsePaths(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	paths := parsePaths(lines)
	if !reflect.DeepEqual(paths["start"], []string{"A", "b"}) {
		t.Errorf("expected start to connect to A and b, but was connected to (%s)", strings.Join(paths["start"], ", "))
	}
	if !reflect.DeepEqual(paths["A"], []string{"start", "c", "b", "end"}) {
		t.Errorf("expected A to connect to c, b, start, end, but was connected to (%s)", strings.Join(paths["A"], ", "))
	}
	if !reflect.DeepEqual(paths["b"], []string{"start", "A", "d", "end"}) {
		t.Errorf("expected b to connect to start, A and d, end, but was connected to (%s)", strings.Join(paths["b"], ", "))
	}
	if !reflect.DeepEqual(paths["d"], []string{"b"}) {
		t.Errorf("expected d to connect to b, but was connected to (%s)", strings.Join(paths["d"], ", "))
	}
	if !reflect.DeepEqual(paths["c"], []string{"A"}) {
		t.Errorf("expected c to connect to A, but was connected to (%s)", strings.Join(paths["c"], ", "))
	}
	if !reflect.DeepEqual(paths["end"], []string{"A", "b"}) {
		t.Errorf("expected end to connect to A and b, but was connected to (%s)", strings.Join(paths["end"], ", "))
	}
}

func TestCountPathsToEnd(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		lines := parse.Lines("input.test.txt")
		paths := parsePaths(lines)
		distinctPaths := countPathsToEnd("start", map[string]bool{}, paths)
		if distinctPaths != 10 {
			t.Errorf("Expected 10 paths through the cave, but was %d", distinctPaths)
		}
	})
	t.Run("example 2", func(t *testing.T) {
		lines := parse.Lines("input.test2.txt")
		paths := parsePaths(lines)
		distinctPaths := countPathsToEnd("start", map[string]bool{}, paths)
		if distinctPaths != 19 {
			t.Errorf("Expected 19 paths through the cave, but was %d", distinctPaths)
		}
	})
	t.Run("example 3", func(t *testing.T) {
		lines := parse.Lines("input.test3.txt")
		paths := parsePaths(lines)
		distinctPaths := countPathsToEnd("start", map[string]bool{}, paths)
		if distinctPaths != 226 {
			t.Errorf("Expected 226 paths through the cave, but was %d", distinctPaths)
		}
	})
}

func TestCountPathsToEnd2(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		lines := parse.Lines("input.test.txt")
		paths := parsePaths(lines)
		distinctPaths := countPathsToEnd2("start", map[string]bool{}, false, paths)
		if distinctPaths != 36 {
			t.Errorf("Expected 36 paths through the cave, but was %d", distinctPaths)
		}
	})
	t.Run("example 2", func(t *testing.T) {
		lines := parse.Lines("input.test2.txt")
		paths := parsePaths(lines)
		distinctPaths := countPathsToEnd2("start", map[string]bool{}, false, paths)
		if distinctPaths != 103 {
			t.Errorf("Expected 103 paths through the cave, but was %d", distinctPaths)
		}
	})
	t.Run("example 3", func(t *testing.T) {
		lines := parse.Lines("input.test3.txt")
		paths := parsePaths(lines)
		distinctPaths := countPathsToEnd2("start", map[string]bool{}, false, paths)
		if distinctPaths != 3509 {
			t.Errorf("Expected 3509 paths through the cave, but was %d", distinctPaths)
		}
	})
}
