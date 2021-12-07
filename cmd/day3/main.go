package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"strconv"
)

func main() {
	report := parse.Lines("cmd/day3/input.txt")
	g, e := calculateGammaEpsilon(report)
	fmt.Printf("Part 1: %d\n", g*e)
	o := calcRating(report, oxygenRatingTest)
	c := calcRating(report, co2ScrubberTest)
	fmt.Printf("Part 2: %d\n", o*c)
}

func calculateGammaEpsilon(report []string) (gamma, epsilon int64) {
	for i := 0; i < len(report[0]); i++ {
		zeroCount, oneCount := 0, 0
		for _, num := range report {
			if num[i] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		epsilon <<= 1
		gamma <<= 1
		if zeroCount > oneCount {
			epsilon |= 1
		} else {
			gamma |= 1
		}
	}
	return
}

func calcRating(report []string, test func(int, int, uint8) bool) int64 {
	remaining := report
	pointer := 0
	for len(remaining) > 1 {
		var newRemaining []string
		zeroCount, oneCount := 0, 0
		for _, num := range remaining {
			if num[pointer] == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		for _, num := range remaining {
			if test(zeroCount, oneCount, num[pointer]) {
				newRemaining = append(newRemaining, num)
			}
		}
		pointer++
		remaining = newRemaining
	}
	rating, _ := strconv.ParseInt(remaining[0], 2, 64)
	return rating
}

func oxygenRatingTest(zeroCount, oneCount int, bit uint8) bool {
	if zeroCount > oneCount && bit == '0' {
		return true
	}
	if oneCount >= zeroCount && bit == '1' {
		return true
	}
	return false
}

func co2ScrubberTest(zeroCount, oneCount int, bit uint8) bool {
	if zeroCount <= oneCount && bit == '0' {
		return true
	}
	if oneCount < zeroCount && bit == '1' {
		return true
	}
	return false
}
