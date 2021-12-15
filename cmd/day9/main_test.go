package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestLowPointCount(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	grid := parseGrid(lines)
	r := riskLevelSum(grid)
	if r != 15 {
		t.Errorf("Expected risk level sum of 15 but was %d", r)
	}
}

func TestLargestBasinsProduct(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	grid := parseGrid(lines)
	p := largestBasinsProduct(grid)
	if p != 1134 {
		t.Errorf("Expected a product of 1134 but was %d", p)
	}
}
