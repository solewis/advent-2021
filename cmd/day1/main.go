package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	depths := parse.Ints("cmd/day1/input.txt", "\n")
	fmt.Printf("Part 1: %d\n", countIncreases(depths))
	fmt.Printf("Part 2: %d\n", countWindowIncreases(depths))
}


func countIncreases(depths []int) int {
	count := 0
	for i := 0; i < len(depths) - 1; i++ {
		if depths[i+1] > depths[i] {
			count++
		}
	}
	return count
}

func countWindowIncreases(depths []int) int {
	var windows []int
	for i := 0; i < len(depths) - 2; i++ {
		windows = append(windows, depths[i] + depths[i+1] + depths[i+2])
	}
	return countIncreases(windows)
}
