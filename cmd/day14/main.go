package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"math"
	"strings"
)

func main() {
	template, rules := parseManual(parse.String("cmd/day14/input.txt"))
	fmt.Printf("Part 1: %d\n", solve(template, rules, 10))
	fmt.Printf("Part 2: %d\n", solve2(template, rules, 40))
}

func solve(template string, rules map[string]string, steps int) int {
	result := template
	for i := 0; i < steps; i++ {
		result = runPairInsertion(result, rules)
	}
	counts := map[rune]int{}
	for _, c := range result {
		counts[c]++
	}
	min, max := math.MaxInt64, 0
	for _, ct := range counts {
		if ct > max {
			max = ct
		}
		if ct < min {
			min = ct
		}
	}
	return max - min
}

func solve2(template string, rules map[string]string, steps int) int {
	counts := map[string]int{}
	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		pairs[pair]++
		counts[template[i:i+1]]++
		//counts[template[i+1:i+2]]++
	}
	counts[template[len(template)-1:]]++
	for i := 0; i < steps; i++ {
		newPairs := map[string]int{}
		for pair, ct := range pairs {
			if ct > 0 {
				c := rules[pair]
				p1 := pair[0:1] + c
				p2 := c + pair[1:2]
				counts[c] += ct
				newPairs[pair] -= ct
				newPairs[p1]+=ct
				newPairs[p2]+=ct
			}
		}
		for pair, ct := range newPairs {
			pairs[pair] += ct
		}
	}
	min, max := math.MaxInt64, 0
	for _, ct := range counts {
		if ct > max {
			max = ct
		}
		if ct < min {
			min = ct
		}
	}
	return max - min
}

func optimized(pair string, rules map[string]string, numSteps int, maxSteps int, counts map[string]int) {
	if numSteps == maxSteps {
		return
	}
	if c, ok := rules[pair]; ok {
		counts[c]++
		optimized(pair[0:1]+c, rules, numSteps+1, maxSteps, counts)
		optimized(c+pair[1:2], rules, numSteps+1, maxSteps, counts)
	}
	//return counts
}

func runPairInsertion(polymer string, rules map[string]string) string {
	var p strings.Builder
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		if rule, ok := rules[pair]; ok {
			p.WriteRune(rune(polymer[i]))
			p.WriteString(rule)
		} else {
			p.WriteRune(rune(polymer[i]))
		}
	}
	p.WriteRune(rune(polymer[len(polymer)-1]))
	return p.String()
}

func parseManual(manual string) (string, map[string]string) {
	parts := strings.Split(manual, "\n\n")
	template := parts[0]
	rules := map[string]string{}
	for _, line := range strings.Split(parts[1], "\n") {
		sections := strings.Split(line, " -> ")
		rules[sections[0]] = sections[1]
	}
	return template, rules
}
