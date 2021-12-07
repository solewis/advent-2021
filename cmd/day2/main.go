package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day2/input.txt")
	h, d := calculateCourse(lines)
	fmt.Printf("Part 1: %d\n", h * d)
	h, d = calculateCourseRuleset2(lines)
	fmt.Printf("Part 2: %d\n", h * d)
}

func calculateCourse(lines []string) (horizontal int, depth int) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		change := parse.Int(parts[1])
		switch parts[0] {
		case "forward":
			horizontal += change
		case "up":
			depth -= change
		case "down":
			depth += change
		}
	}
	return
}

func calculateCourseRuleset2(lines []string) (horizontal int, depth int) {
	aim := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		change := parse.Int(parts[1])
		switch parts[0] {
		case "forward":
			horizontal += change
			depth += aim * change
		case "up":
			aim -= change
		case "down":
			aim += change
		}
	}
	return
}