package main

import (
	"github.com/antimatter96/advent/2023/common"
)

func shiftNorthNew(board common.Graph[string]) {

	shift(board,
		0, len(board[0]),
		0, limitMax(len(board)-1), inc,
		innerDiffSame, outerDiffSame,
		innerDiffInc, outerDiffSame,
	)

}

func shiftSouthNew(board common.Graph[string]) {

	shift(board,
		0, len(board[0]),
		len(board)-1, limitMinZero, dec,
		innerDiffSame, outerDiffSame,
		innerDiffDec, outerDiffSame,
	)

}

func shiftWestNew(board common.Graph[string]) {

	shift(board,
		0, len(board),
		0, limitMax(len(board[0])-1), inc,
		outerDiffSame, innerDiffSame,
		outerDiffSame, innerDiffInc,
	)

}

func shiftEastNew(board common.Graph[string]) {

	shift(board,
		0, len(board),
		len(board[0])-1, limitMinZero, dec,
		outerDiffSame, innerDiffSame,
		outerDiffSame, innerDiffDec,
	)

}

func shift(board common.Graph[string], outerStart, outerLimit int, innerStart int, innerLimit func(int) bool, innerFunc func(int) int, checkXFunc, checkYFunc, againstXFunc, againstYFunc func(int, int) int) {

	for outer := outerStart; outer < outerLimit; outer++ {

		changed := true
		for changed {
			changed = false

			for inner := innerStart; innerLimit(inner); inner = innerFunc(inner) {

				checkX := checkXFunc(inner, outer)
				checkY := checkYFunc(inner, outer)
				againstX := againstXFunc(inner, outer)
				againstY := againstYFunc(inner, outer)

				if board[checkX][checkY] == "." && board[againstX][againstY] == "O" {
					board[checkX][checkY] = "O"
					board[againstX][againstY] = "."
					changed = true
				}

			}

		}
	}
}

func limitMax(limit int) func(int) bool {
	return func(i int) bool {

		return i < limit
	}
}

func limitMin(limit int) func(int) bool {
	return func(i int) bool {

		return i > limit
	}
}

var limitMinZero = limitMin(0)

func dec(i int) int {
	return i - 1
}

func inc(i int) int {
	return i + 1
}

func same(i int) int {
	return i
}

var outerDiffSame = outerDiff(same)
var innerDiffSame = innerDiff(same)
var innerDiffInc = innerDiff(inc)
var innerDiffDec = innerDiff(dec)

func innerDiff(changeFunc func(int) int) func(int, int) int {
	return func(inner, outer int) int {
		return changeFunc(inner)
	}
}

func outerDiff(changeFunc func(int) int) func(int, int) int {
	return func(inner, outer int) int {
		return changeFunc(outer)
	}
}

func shiftNorth(board common.Graph[string]) {

	for j := 0; j < len(board[0]); j++ {

		changed := true
		for changed {
			changed = false

			for i := 0; i < len(board)-1; i++ {
				if board[i][j] == "." && board[i+1][j] == "O" {
					board[i][j] = "O"
					board[i+1][j] = "."
					changed = true
				}
			}

		}
	}
}

func shiftSouth(board common.Graph[string]) {

	for j := 0; j < len(board[0]); j++ {

		changed := true
		for changed {
			changed = false

			for i := len(board) - 1; i > 0; i-- {
				if board[i][j] == "." && board[i-1][j] == "O" {
					board[i][j] = "O"
					board[i-1][j] = "."
					changed = true
				}
			}

		}
	}
}

func shiftWest(board common.Graph[string]) {

	for i := 0; i < len(board); i++ {

		changed := true
		for changed {
			changed = false

			for j := 0; j < len(board[0])-1; j++ {
				if board[i][j] == "." && board[i][j+1] == "O" {
					board[i][j] = "O"
					board[i][j+1] = "."
					changed = true
				}
			}

		}
	}
}

func shiftEast(board common.Graph[string]) {

	for i := 0; i < len(board); i++ {

		changed := true
		for changed {
			changed = false

			for j := len(board[0]) - 1; j > 0; j-- {
				if board[i][j] == "." && board[i][j-1] == "O" {
					board[i][j] = "O"
					board[i][j-1] = "."
					changed = true
				}
			}

		}
	}
}
