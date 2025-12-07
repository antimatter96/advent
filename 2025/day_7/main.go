package main

import (
	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) common.Graph[rune] {
	board := make(common.Graph[rune], len(inp))
	for i, inp_row := range inp {
		board[i] = []rune(inp_row)
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

func Part1(board common.Graph[rune]) int {
	var sum int

	board.Print()
	for j := range board[0] {
		if board[0][j] == 'S' {
			board[1][j] = '|'
		}
	}
	board.Print()

	for i := range board {
		if i == len(board)-1 {
			break
		}
		if i == 0 {
			continue
		}

		for j := range board[i] {
			if board[i][j] == '|' {
				if board[i+1][j] == '.' {
					board[i+1][j] = '|'
					continue
				} else if board[i+1][j] == '^' {
					if j-1 >= 0 {
						board[i+1][j-1] = '|'
					}
					if j+1 < len(board[i]) {
						board[i+1][j+1] = '|'
					}
				}
			}
		}

		board.Print()
	}

	for i := range board {
		if i == 0 {
			continue
		}
		for j := range board[i] {
			if board[i][j] == '^' && board[i-1][j] == '|' {
				sum++
			}
		}

	}

	return sum
}

func Part2(board common.Graph[rune]) int {
	var sum int

	return sum

}
