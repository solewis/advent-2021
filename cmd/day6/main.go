package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	states := parse.Ints("cmd/day6/input.txt", ",")
	ct := countAllAfterDays(states, 80)
	fmt.Printf("Part 1: %d\n", ct)
	ct = countAllAfterDays(states, 256)
	fmt.Printf("Part 2: %d\n", ct)
}

func countAllAfterDays(states []int, days int) int {
	count := 0
	for _, state := range states {
		count += countAfterDays(state, days)
	}
	return count
}

var cache = make(map[int]int)

func countAfterDays(state, days int) int {
	if days <= state {
		return 1
	}
	nextSpawn := days - state
	if _, ok := cache[nextSpawn]; !ok {
		cache[nextSpawn] = countAfterDays(8, nextSpawn-1) + countAfterDays(6, nextSpawn-1)
	}
	return cache[nextSpawn]
}
