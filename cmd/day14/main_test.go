package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestRunPairInsertion(t *testing.T) {
	template, rules := parseManual(parse.String("input.test.txt"))
	t.Run("has correct value after steps", func(t *testing.T) {
		result := runPairInsertion(template, rules)
		if result != "NCNBCHB" {
			t.Errorf("Expected NCNBCHB after 1 step but was %s", result)
		}
		result = runPairInsertion(result, rules)
		if result != "NBCCNBBBCBHCB" {
			t.Errorf("Expected NBCCNBBBCBHCB after 2 steps but was %s", result)
		}
	})
	t.Run("runs many steps", func(t *testing.T) {
		result := template
		for i := 0; i < 10; i++ {
			result = runPairInsertion(result, rules)
		}
		if len(result) != 3073 {
			t.Errorf("Expected a polymer of length 3073 after 10 steps but was %d", len(result))
		}
	})
}

func TestSolve(t *testing.T) {
	template, rules := parseManual(parse.String("input.test.txt"))
	r := solve2(template, rules, 10)
	if r != 1588 {
		t.Errorf("Expected 1588 but was %d", r)
	}
	r = solve2(template, rules, 40)
	if r != 2188189693529 {
		t.Errorf("Expected 2188189693529 but was %d", r)
	}
}
