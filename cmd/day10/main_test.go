package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestSyntaxScore(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	s := syntaxErrorScore(lines)
	if s != 26397 {
		t.Errorf("Expected a score of 26397 but was %d", s)
	}
}

func TestAutocompleteScore(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	s := autocompleteScore(lines)
	if s != 288957 {
		t.Errorf("Expected 288957 but was %d", s)
	}
}
