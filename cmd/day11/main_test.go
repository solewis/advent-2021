package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestCountFlashes(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	t.Run("10 steps", func(t *testing.T) {
		g := parseGrid(lines)
		c := countFlashes(g, 10)
		if c != 204 {
			t.Errorf("expected 204 flashes but was %d", c)
		}
	})
	t.Run("100 steps", func(t *testing.T) {
		g := parseGrid(lines)
		c := countFlashes(g, 100)
		if c != 1656 {
			t.Errorf("expected 1656 flashes but was %d", c)
		}
	})
}

func TestFirstSynchronizedFlash(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	g := parseGrid(lines)
	step := firstSynchronizedFlash(g)
	if step != 195 {
		t.Errorf("Expected step 195 but was %d", step)
	}
}
