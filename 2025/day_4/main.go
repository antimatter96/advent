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
			if ele == '@' {
				adjacent := countAdjacentRolls(i, j, board)
				if adjacent < 4 {
					sum += 1
				}
				common.Log.Debug().Int("_i", i).Int("_j", j).Int("adjacent", adjacent).Send()
			}
		}
	}

	return sum
}

func Part2(board [][]rune) int {
	var sum int

	for {
		thisRound := 0
		pointsToRemove := make([]*common.Point, 0)

		for i, row := range board {
			for j, ele := range row {
				if ele == '@' {
					adjacent := countAdjacentRolls(i, j, board)
					if adjacent < 4 {
						thisRound += 1
						pointsToRemove = append(pointsToRemove, &common.Point{X: i, Y: j})
					}
				}
			}
		}

		if thisRound == 0 {
			break
		}

		for _, point := range pointsToRemove {
			board[point.X][point.Y] = '.'

			common.Log.Debug().Int("_i", point.X).Int("_j", point.Y).Msg("Removed")
		}

		sum += thisRound
	}

	return sum
}

func isTissueRoll(target []rune, board [][]rune, move *common.Directions) bool {
	if len(target) == 0 {
		return false
	}
	if move.I >= len(board) || move.I < 0 {
		return false
	}
	if move.J >= len(board[0]) || move.J < 0 {
		return false
	}
	return board[move.I][move.J] == '@'
}

func countAdjacentRolls(i, j int, board [][]rune) int {
	total := 0

	movements := common.GenerateDirections(i, j)

	for _, movement := range movements {
		if isTissueRoll([]rune("MAS"), board, movement) {
			total += 1
		}
	}

	return total
}
