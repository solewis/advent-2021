package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestCalculateCourse(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	h, d := calculateCourse(lines)
	if h != 15 {
		t.Errorf("Expected horizontal position of 15 but was %d", h)
	}
	if d != 10 {
		t.Errorf("Expected depth of 10 but was %d", d)
	}
}

func TestCalculateCourseRuleset2(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	h, d := calculateCourseRuleset2(lines)
	if h != 15 {
		t.Errorf("Expected horizontal position of 15 but was %d", h)
	}
	if d != 60 {
		t.Errorf("Expected depth of 60 but was %d", d)
	}
}
