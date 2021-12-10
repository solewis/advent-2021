package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestWinnerBingo(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	draws, boards := parseBingo(lines)
	score := findWinnerScore(draws, boards)
	if score != 4512 {
		t.Errorf("Expected score of 4512 but was %d", score)
	}
}

func TestLoserBingo(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	draws, boards := parseBingo(lines)
	score := findLoserScore(draws, boards)
	if score != 1924 {
		t.Errorf("Expected score of 1924 but was %d", score)
	}
}
