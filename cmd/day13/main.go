package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	manual := parse.String("cmd/day13/input.txt")
	dots, folds := parseManual(manual)
	newDots := fold(dots, folds[0])
	fmt.Printf("Part 1: %d\n", len(newDots))
	printFinalMap(dots, folds)
}

func printFinalMap(dots []point, folds []point) {
	for _, f := range folds {
		dots = fold(dots, f)
	}
	dotMap := map[point]bool{}
	var width, height int
	for _, d := range dots {
		if d.x > width {
			width = d.x
		}
		if d.y > height {
			height = d.y
		}
		dotMap[d] = true
	}
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			if dotMap[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func fold(dots []point, fold point) []point {
	dotMap := map[point]bool{}
	for _, dot := range dots {
		if fold.x > 0 && dot.x > fold.x {
			// vertical fold
			dotMap[point{x: fold.x - (dot.x - fold.x), y: dot.y}] = true
		} else if fold.y > 0 && dot.y > fold.y {
			// horizontal fold
			dotMap[point{y: fold.y - (dot.y - fold.y), x: dot.x}] = true
		} else {
			dotMap[dot] = true
		}
	}
	var newDots []point
	for d := range dotMap {
		newDots = append(newDots, d)
	}
	return newDots
}

func parseManual(manual string) ([]point, []point) {
	sections := strings.Split(manual, "\n\n")
	var dots []point
	dotRegex := regexp.MustCompile(`^(\d+),(\d+)$`)
	for _, dot := range strings.Split(sections[0], "\n") {
		matches := dotRegex.FindStringSubmatch(dot)
		dots = append(dots, point{x: parse.Int(matches[1]), y: parse.Int(matches[2])})
	}

	var folds []point
	foldRegex := regexp.MustCompile(`^fold along ([xy])=(\d+)$`)
	for _, fold := range strings.Split(sections[1], "\n") {
		matches := foldRegex.FindStringSubmatch(fold)
		var p point
		if matches[1] == "x" {
			p = point{x: parse.Int(matches[2])}
		} else {
			p = point{y: parse.Int(matches[2])}
		}
		folds = append(folds, p)
	}
	return dots, folds
}
