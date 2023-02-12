package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestFold(t *testing.T) {
	manual := parse.String("input.test.txt")
	dots, folds := parseManual(manual)
	newDots := fold(dots, folds[0])
	if len(newDots) != 17 {
		t.Errorf("Expected 17 dots visible, but was %d", len(newDots))
	}
}
