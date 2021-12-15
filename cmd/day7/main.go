package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"math"
)

func main() {
	crabPositions := parse.Ints("cmd/day7/input.txt", ",")
	c := fuelCostAtBestPosition(crabPositions, fuelCostToAlign)
	fmt.Printf("Part 1: %d\n", c)
	c = fuelCostAtBestPosition(crabPositions, fuelCostToAlignPart2)
	fmt.Printf("Part 2: %d\n", c)
}

func fuelCostAtBestPosition(crabPositions []int, costCalc func([]int, int) int) int {
	min, max := math.MaxInt64, 0
	for _, pos := range crabPositions {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}
	lowestCost := math.MaxInt64
	for i := min; i < max; i++ {
		cost := costCalc(crabPositions, i)
		if cost < lowestCost {
			lowestCost = cost
		}
	}
	return lowestCost
}
func fuelCostToAlign(crabPositions []int, target int) int {
	cost := 0
	for _, pos := range crabPositions {
		cost += abs(pos - target)
	}
	return cost
}

func fuelCostToAlignPart2(crabPositions []int, target int) int {
	cost := 0
	for _, pos := range crabPositions {
		steps := abs(pos - target)
		cost += (steps * (steps + 1)) / 2
	}
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
