package main

import (
	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) [][]rune {
	board := make([][]rune, 0, len(inp))
	for _, inp_row := range inp {
		board = append(board, []rune(inp_row))
	}

	return board
}

func parsePart2(inp []string) [][]rune {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(board [][]rune) int {
	var sum int

	for i, row := range board {
		for j, ele := range row {
			if ele == 'X' {
				sum += searchXMAS(i, j, board)
			}
		}
	}

	return sum
}

func Part2(board [][]rune) int {
	var sum int

	for i, row := range board {
		for j, ele := range row {
			if ele == 'A' {
				sum += searchMAS(i, j, board)
			}
		}
	}

	return sum
}

func searchRec(target []rune, board [][]rune, move *common.Directions) bool {
	if len(target) == 0 {
		return true
	}
	if move.I >= len(board) || move.I < 0 {
		return false
	}
	if move.J >= len(board[0]) || move.J < 0 {
		return false
	}
	if board[move.I][move.J] == target[0] {
		return searchRec(target[1:], board, &common.Directions{I: move.NextI(move.I), J: move.NextJ(move.J), NextI: move.NextI, NextJ: move.NextJ})
	}
	return false
}

func searchXMAS(i, j int, board [][]rune) int {
	total := 0

	movements := common.GenerateDirections(i, j)

	for _, movement := range movements {
		if searchRec([]rune("MAS"), board, movement) {
			total += 1
		}
	}

	return total
}

func searchMAS(i, j int, board [][]rune) int {
	movements := common.GenerateDirections(i, j)

	backSlash := searchRec([]rune("M"), board, movements["UP-LEFT"]) && searchRec([]rune("S"), board, movements["DOWN-RIGHT"])
	if !backSlash {
		backSlash = searchRec([]rune("S"), board, movements["UP-LEFT"]) && searchRec([]rune("M"), board, movements["DOWN-RIGHT"])
	}
	if !backSlash {
		return 0
	}

	forwardSlash := searchRec([]rune("M"), board, movements["UP-RIGHT"]) && searchRec([]rune("S"), board, movements["DOWN-LEFT"])
	if !forwardSlash {
		forwardSlash = searchRec([]rune("S"), board, movements["UP-RIGHT"]) && searchRec([]rune("M"), board, movements["DOWN-LEFT"])
	}

	if !forwardSlash {
		return 0
	}

	return 1
}
