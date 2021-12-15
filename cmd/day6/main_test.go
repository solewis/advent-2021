package main

import "testing"

func TestCountAfterDays(t *testing.T) {
	t.Run("3", func(t *testing.T) {
		ct := countAfterDays(3, 3)
		if ct != 1 {
			t.Errorf("expected 1 but was %d", ct)
		}
	})
	t.Run("4", func(t *testing.T) {
		ct := countAfterDays(3, 4)
		if ct != 2 {
			t.Errorf("expected 2 but was %d", ct)
		}
	})
	t.Run("10", func(t *testing.T) {
		ct := countAfterDays(3, 10)
		if ct != 2 {
			t.Errorf("expected 2 but was %d", ct)
		}
	})
	t.Run("11", func(t *testing.T) {
		ct := countAfterDays(3, 11)
		if ct != 3 {
			t.Errorf("expected 3 but was %d", ct)
		}
	})
	t.Run("13", func(t *testing.T) {
		ct := countAfterDays(3, 13)
		if ct != 4 {
			t.Errorf("expected 4 but was %d", ct)
		}
	})
	t.Run("18", func(t *testing.T) {
		ct := countAfterDays(3, 18)
		if ct != 5 {
			t.Errorf("expected 5 but was %d", ct)
		}
	})
}

func TestPart1(t *testing.T) {
	initialStates := []int{3, 4, 3, 1, 2}
	t.Run("After 18 days", func(t *testing.T) {
		ct := countAllAfterDays(initialStates, 18)
		if ct != 26 {
			t.Errorf("Expected 26 lanternfish after 18 days but was %d", ct)
		}
	})
	t.Run("After 80 days", func(t *testing.T) {
		ct := countAllAfterDays(initialStates, 80)
		if ct != 5934 {
			t.Errorf("Expected 5934 lanternfish after 80 days but was %d", ct)
		}
	})
	t.Run("After 256 days", func(t *testing.T) {
		ct := countAllAfterDays(initialStates, 256)
		if ct != 26984457539 {
			t.Errorf("Expected 26984457539 lanternfish after 256 days but was %d", ct)
		}
	})
}
