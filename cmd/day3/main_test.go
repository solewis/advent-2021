package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestCalculateGammaEpsilon(t *testing.T) {
	report := parse.Lines("input.test.txt")
	g, e := calculateGammaEpsilon(report)
	if g != 22 {
		t.Errorf("Expected gamma of 22 but was %d", g)
	}
	if e != 9 {
		t.Errorf("Expected epsilon of 9 but was %d", e)
	}
}

func TestCalcRating(t *testing.T) {
	report := parse.Lines("input.test.txt")
	o := calcRating(report, oxygenRatingTest)
	if o != 23 {
		t.Errorf("Expected oxygen rating of 23 but was %d", o)
	}
	c := calcRating(report, co2ScrubberTest)
	if c != 10 {
		t.Errorf("Expected CO2 scrubber rating of 10 but was %d", c)
	}
}
