package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"sort"
)

func main() {
	lines := parse.Lines("cmd/day10/input.txt")
	score := syntaxErrorScore(lines)
	fmt.Printf("Part 1: %d\n", score)
	score = autocompleteScore(lines)
	fmt.Printf("Part 2: %d\n", score)
}

func autocompleteScore(lines []string) int {
	var scores []int
lineLoop:
	for _, line := range lines {
		r := reduce(line)
		for _, c := range r {
			if c == ')' || c == ']' || c == '}' || c == '>' {
				continue lineLoop //corrupted
			}
		}
		score := 0
		for i := len(r) - 1; i >= 0; i-- {
			switch r[i] {
			case '(':
				score = score*5 + 1
			case '[':
				score = score*5 + 2
			case '{':
				score = score*5 + 3
			case '<':
				score = score*5 + 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func syntaxErrorScore(lines []string) int {
	syntaxScore := 0
lineLoop:
	for _, line := range lines {
		r := reduce(line)
		for _, c := range r {
			switch c {
			case ')':
				syntaxScore += 3
				continue lineLoop
			case ']':
				syntaxScore += 57
				continue lineLoop
			case '}':
				syntaxScore += 1197
				continue lineLoop
			case '>':
				syntaxScore += 25137
				continue lineLoop
			}
		}
	}
	return syntaxScore
}

func reduce(s string) string {
	newStr := ""
	for i := 0; i < len(s); {
		a := s[i]
		if i == len(s)-1 {
			newStr += string(a)
			break
		}
		b := s[i+1]
		if (a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}') || (a == '<' && b == '>') {
			i += 2
			continue
		}
		newStr += string(a)
		i++
	}
	if len(newStr) == len(s) {
		return newStr
	}
	return reduce(newStr)
}
