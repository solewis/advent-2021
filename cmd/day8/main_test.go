package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	entries := parseEntries(lines)
	p1 := part1(entries)
	if p1 != 26 {
		t.Errorf("expected 26 occurrences but was %d", p1)
	}
}

func TestOutputSum(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	entries := parseEntries(lines)
	sum := outputSum(entries)
	if sum != 61229 {
		t.Errorf("Expected a sum of 61229 but was %d", sum)
	}
}
