package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

type grid [][]int

func (g grid) inBounds(p point) bool {
	return p.x >= 0 && p.x < len(g[0]) && p.y >= 0 && p.y < len(g)
}

type point struct {
	x, y int
}

func (p point) adj() []point {
	return []point{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

func main() {
	lines := parse.Lines("cmd/day11/input.txt")
	g := parseGrid(lines)
	count := countFlashes(g, 100)
	fmt.Printf("Part 1: %d\n", count)

	g = parseGrid(lines)
	step := firstSynchronizedFlash(g)
	fmt.Printf("Part 2: %d\n", step)
}

func firstSynchronizedFlash(grid [][]int) int {
	for step := 1; ; step++ {
		// increment by one
		for y, row := range grid {
			for x := range row {
				grid[y][x]++
			}
		}

		// handle flashes
		flashed := map[point]bool{}
		for y, row := range grid {
			for x := range row {
				flash(point{x, y}, grid, flashed)
			}
		}

		// reset flashed points
		for y, row := range grid {
			for x := range row {
				if grid[y][x] > 9 {
					grid[y][x] = 0
				}
			}
		}
		if len(flashed) == len(grid) * len(grid[0]) {
			return step
		}
	}
}

func countFlashes(grid [][]int, steps int) int {
	flashCount := 0
	for step := 0; step < steps; step++ {
		// increment by one
		for y, row := range grid {
			for x := range row {
				grid[y][x]++
			}
		}

		// handle flashes
		flashed := map[point]bool{}
		for y, row := range grid {
			for x := range row {
				flash(point{x, y}, grid, flashed)
			}
		}

		// reset flashed points
		for y, row := range grid {
			for x := range row {
				if grid[y][x] > 9 {
					grid[y][x] = 0
				}
			}
		}
		flashCount += len(flashed)
	}
	return flashCount
}

// recursively handles flashing at a given point
func flash(p point, g grid, flashMap map[point]bool) {
	if flashMap[p] || g[p.y][p.x] <= 9 {
		return
	}
	flashMap[p] = true
	for _, adj := range p.adj() {
		if g.inBounds(adj) {
			g[adj.y][adj.x]++
			flash(adj, g, flashMap)
		}
	}
}

func parseGrid(lines []string) grid {
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
	}
	return grid
}
