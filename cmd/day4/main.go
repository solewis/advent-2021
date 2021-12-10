package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"strings"
)

type board [][]int

func (board board) isWinner() bool {
row:
	for _, row := range board {
		for _, space := range row {
			if space != -1 {
				continue row
			}
		}
		return true
	}
column:
	for i := 0; i < 5; i++ {
		for _, row := range board {
			if row[i] != -1 {
				continue column
			}
		}
		return true
	}
	return false
}

func (board board) mark(draw int) {
	for j, row := range board {
		for k, space := range row {
			if space == draw {
				board[j][k] = -1
			}
		}
	}
}

func (board board) score(draw int) int {
	unmarked := 0
	for _, row := range board {
		for _, space := range row {
			if space != -1 {
				unmarked += space
			}
		}
	}
	return unmarked * draw
}

func main() {
	lines := parse.Lines("cmd/day4/input.txt")
	draws, boards := parseBingo(lines)
	score := findWinnerScore(draws, boards)
	fmt.Printf("Part 1: %d\n", score)
	draws, boards = parseBingo(lines)
	score = findLoserScore(draws, boards)
	fmt.Printf("Part 2: %d\n", score)
}

func findWinnerScore(draws []int, boards []board) int {
	for _, draw := range draws {
		for _, board := range boards {
			board.mark(draw)
			if board.isWinner() {
				return board.score(draw)
			}
		}
	}
	panic("could not solve")
}

func findLoserScore(draws []int, boards []board) int {
	finishedBoards := map[int]bool{}
	for _, draw := range draws {
		for boardIndex, board := range boards {
			if finishedBoards[boardIndex] {
				continue
			}
			board.mark(draw)
			if board.isWinner() {
				finishedBoards[boardIndex] = true
				if len(finishedBoards) == len(boards) {
					return board.score(draw)
				}
			}
		}
	}
	panic("could not solve")
}

func parseBingo(lines []string) ([]int, []board) {
	drawsStr := strings.Split(lines[0], ",")
	var draws []int
	for _, draw := range drawsStr {
		draws = append(draws, parse.Int(draw))
	}

	var boards []board
	var board board
	for _, line := range lines[2:] {
		if line == "" {
			boards = append(boards, board)
			board = [][]int{}
			continue
		}
		var row []int
		for _, space := range strings.Fields(line) {
			row = append(row, parse.Int(space))
		}
		board = append(board, row)
	}
	boards = append(boards, board)

	return draws, boards
}
