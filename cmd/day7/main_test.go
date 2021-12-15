package main

import "testing"

func TestLowestFuelCost(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		crabPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
		lowestCost := fuelCostAtBestPosition(crabPositions, fuelCostToAlign)
		if lowestCost != 37 {
			t.Errorf("Expected a cost of 37 but was %d", lowestCost)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		crabPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
		lowestCost := fuelCostAtBestPosition(crabPositions, fuelCostToAlignPart2)
		if lowestCost != 168 {
			t.Errorf("Expected a cost of 168 but was %d", lowestCost)
		}
	})
}
