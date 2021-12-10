package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestOverlaps(t *testing.T) {
	fileRows := parse.Lines("input.test.txt")
	lines := parseLines(fileRows)

	t.Run("without diagonals", func(t *testing.T) {
		overlapCount := overlaps(lines, false)
		if overlapCount != 5 {
			t.Errorf("Expected 5 points of overlap without diagonals but was %d", overlapCount)
		}
	})

	t.Run("with diagonals", func(t *testing.T) {
		overlapCount := overlaps(lines, true)
		if overlapCount != 12 {
			t.Errorf("Expected 12 points of overlap with diagonals but was %d", overlapCount)
		}
	})
}
