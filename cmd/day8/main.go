package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"sort"
	"strings"
)

type entry struct {
	patterns, output []string
}

func main() {
	lines := parse.Lines("cmd/day8/input.txt")
	entries := parseEntries(lines)
	p1 := part1(entries)
	fmt.Printf("Part 1: %d\n", p1)
	p2 := outputSum(entries)
	fmt.Printf("Part 2: %d\n", p2)
}

func part1(entries []entry) int {
	count := 0
	for _, entry := range entries {
		for _, v := range entry.output {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	return count
}

func outputSum(entries []entry) int {
	totalOutput := 0
	for _, entry := range entries { // solve 1, 7, 4, 9 first
		key := map[string]string{}
		bKey := map[int]string{}
		for _, pattern := range entry.patterns {
			l := len(pattern)
			s := sortString(pattern)
			switch l {
			case 2: // number one
				key[s] = "1"
				bKey[1] = s
			case 3: // number 7
				key[s] = "7"
				bKey[7] = s
			case 4: // number 4
				key[s] = "4"
				bKey[4] = s
			case 7: // number 8
				key[s] = "8"
				bKey[8] = s
			}
		}
		for _, pattern := range entry.patterns { // solve 0, 6, 9 next
			l := len(pattern)
			s := sortString(pattern)
			switch l {
			case 6: // number 0, 6 or 9
				// if missing letter is in 4 and 8 -> 0
				// if missing letter is in all 4 -> 6
				// if missing letter is in none -> 9
				missingLetter := missingLetters(s)[0]
				if !strings.Contains(bKey[4], missingLetter) {
					// if missing section is not in 4, must be 9
					key[s] = "9"
					bKey[9] = s
				} else if strings.Contains(bKey[7], missingLetter) {
					// if missing section is in 7 (and 4 from first check), must be 6
					key[s] = "6"
					bKey[6] = s
				} else {
					// if missing section is in 4 but not 7, must be 0
					key[s] = "0"
					bKey[0] = s
				}
			}
		}
		for _, pattern := range entry.patterns { // solve 2, 3, 5 last
			l := len(pattern)
			s := sortString(pattern)
			switch l {
			case 5: // number 2, 3 or 5
				missing := missingLetters(s)
				if !strings.Contains(bKey[1], missing[0]) && !strings.Contains(bKey[1], missing[1]) {
					// if missing letters not in 1, must be 3
					key[s] = "3"
					bKey[3] = s
				} else if strings.Contains(missingLetters(bKey[6])[0], missing[0]) || strings.Contains(missingLetters(bKey[6])[0], missing[1]) {
					// if shares a missing letter with 6, must be 5
					key[s] = "5"
					bKey[5] = s
				} else {
					key[s] = "2"
					bKey[2] = s
				}
			}
		}
		output := ""
		for _, o := range entry.output {
			s := sortString(o)
			output += key[s]
		}
		totalOutput += parse.Int(output)
	}
	return totalOutput
}

func missingLetters(pattern string) []string {
	var missing []string
	for _, l := range "abcdefg" {
		if !strings.Contains(pattern, string(l)) {
			missing = append(missing, string(l))
		}
	}
	return missing
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func parseEntries(lines []string) []entry {
	var entries []entry
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		entries = append(entries, entry{
			patterns: strings.Split(parts[0], " "),
			output:   strings.Split(parts[1], " "),
		})
	}
	return entries
}
