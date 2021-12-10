package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
)

type point struct {
	x, y int
}

type line struct {
	start, end point
}

func (line line) isHorizontalVertical() bool {
	return line.start.x == line.end.x || line.start.y == line.end.y
}

func (line line) coveredPoints() []point {
	var points []point
	slopeX, slopeY := slope(line.start.x, line.end.x), slope(line.start.y, line.end.y)
	x, y := line.start.x, line.start.y
	for {
		points = append(points, point{x: x, y: y})
		if x == line.end.x && y == line.end.y {
			break
		}
		x += slopeX
		y += slopeY
	}
	return points
}

func main() {
	fileRows := parse.Lines("cmd/day5/input.txt")
	lines := parseLines(fileRows)
	overlapCount := overlaps(lines, false)
	fmt.Printf("Part 1: %d\n", overlapCount)
	overlapCount = overlaps(lines, true)
	fmt.Printf("Part 2: %d\n", overlapCount)
}

func overlaps(lines []line, includeDiagonal bool) int {
	oceanFloor := map[point]int{}
	for _, line := range lines {
		if includeDiagonal || line.isHorizontalVertical() {
			for _, p := range line.coveredPoints() {
				oceanFloor[p]++
			}
		}
	}

	overlapCount := 0
	for _, lines := range oceanFloor {
		if lines > 1 {
			overlapCount++
		}
	}
	return overlapCount
}

func parseLines(fileRows []string) []line {
	lineRegex := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	var lines []line
	for _, row := range fileRows {
		matches := lineRegex.FindStringSubmatch(row)
		lines = append(lines, line{
			start: point{x: parse.Int(matches[1]), y: parse.Int(matches[2])},
			end:   point{x: parse.Int(matches[3]), y: parse.Int(matches[4])},
		})
	}
	return lines
}

func slope(start, end int) int {
	if start == end {
		return 0
	}
	if start < end {
		return 1
	}
	return -1
}
