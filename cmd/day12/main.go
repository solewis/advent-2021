package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	lines := parse.Lines("cmd/day12/input.txt")
	paths := parsePaths(lines)
	distinctPaths := countPathsToEnd("start", nil, paths)
	fmt.Printf("Part 1: %d\n", distinctPaths)
	distinctPaths2 := countPathsToEnd2("start", nil, false, paths)
	fmt.Printf("Part 2: %d", distinctPaths2)
}

func parsePaths(lines []string) map[string][]string {
	paths := map[string][]string{}
	lineRegex := regexp.MustCompile(`^(\w+)-(\w+)$`)
	for _, line := range lines {
		matches := lineRegex.FindStringSubmatch(line)
		key := matches[1]
		val := matches[2]
		if vals, ok := paths[key]; ok {
			paths[key] = append(vals, val)
		} else {
			paths[key] = []string{val}
		}
		if keys, ok := paths[val]; ok {
			paths[val] = append(keys, key)
		} else {
			paths[val] = []string{key}
		}
	}
	return paths
}

func countPathsToEnd(next string, visited map[string]bool, paths map[string][]string) int {
	if next == "end" {
		return 1
	}
	if unicode.IsLower(rune(next[0])) && visited[next] {
		return 0
	}
	count := 0
	newVisited := map[string]bool{}
	for k, v := range visited {
		newVisited[k] = v
	}
	newVisited[next] = true
	for _, n := range paths[next] {
		count += countPathsToEnd(n, newVisited, paths)
	}
	return count
}

func countPathsToEnd2(next string, visited map[string]bool, smallCaveTwice bool, paths map[string][]string) int {
	if next == "end" {
		return 1
	}
	if next == "start" && visited[next] {
		return 0
	}
	if smallCaveTwice && unicode.IsLower(rune(next[0])) && visited[next] {
		return 0
	}
	if unicode.IsLower(rune(next[0])) && visited[next] {
		smallCaveTwice = true
	}
	count := 0
	newVisited := map[string]bool{}
	for k, v := range visited {
		newVisited[k] = v
	}
	newVisited[next] = true
	for _, n := range paths[next] {
		count += countPathsToEnd2(n, newVisited, smallCaveTwice, paths)
	}
	return count
}
