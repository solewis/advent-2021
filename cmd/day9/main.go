package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"sort"
)

type grid [][]int

func (g grid) adjacentPoints(row, col int) []point {
	points := []point{
		{row - 1, col},
		{row + 1, col},
		{row, col - 1},
		{row, col + 1},
	}
	var validPoints []point
	for _, p := range points {
		if p.row >= 0 && p.row < len(g) && p.col >= 0 && p.col < len(g[0]) {
			validPoints = append(validPoints, p)
		}
	}

	return validPoints
}

type point struct {
	row, col int
}

func main() {
	lines := parse.Lines("cmd/day9/input.txt")
	grid := parseGrid(lines)
	c := riskLevelSum(grid)
	fmt.Printf("Part 1: %d\n", c)
	p := largestBasinsProduct(grid)
	fmt.Printf("Part 2: %d\n", p)
}

func lowPoints(grid grid) []point {
	var lowPoints []point
	for rowIndex, row := range grid {
	rowLoop:
		for colIndex, val := range row {
			for _, p := range grid.adjacentPoints(rowIndex, colIndex) {
				if grid[p.row][p.col] <= val {
					continue rowLoop
				}
			}
			lowPoints = append(lowPoints, point{rowIndex, colIndex})
		}
	}
	return lowPoints
}

func riskLevelSum(grid grid) int {
	riskLevel := 0
	for _, lp := range lowPoints(grid) {
		riskLevel += 1 + grid[lp.row][lp.col]
	}
	return riskLevel
}

func largestBasinsProduct(grid grid) int {
	lowPoints := lowPoints(grid)
	var basinSizes []int
	for _, lp := range lowPoints {
		basinSizes = append(basinSizes, basinSize(lp, grid))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	product := 1
	for i := 0; i < 3; i++ {
		product *= basinSizes[i]
	}
	return product
}

func basinSize(lowPoint point, grid grid) int {
	visited := map[point]bool{}
	queue := []point{lowPoint}
	size := 0
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		if visited[next] {
			continue
		}
		visited[next] = true
		size++
		for _, adj := range grid.adjacentPoints(next.row, next.col) {
			if grid[adj.row][adj.col] < 9 {
				queue = append(queue, adj)
			}
		}
	}
	return size
}

func parseGrid(lines []string) grid {
	var grid grid
	for _, line := range lines {
		var row []int
		for _, val := range line {
			row = append(row, parse.Int(string(val)))
		}
		grid = append(grid, row)
	}
	return grid
}
