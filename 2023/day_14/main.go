package main

import (
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) common.Graph[string] {

	for i := 0; i < len(inp); i++ {
		inp[i] = strings.TrimSpace(inp[i])
	}

	board := make(common.Graph[string], 0)

	for _, rowString := range inp {
		row := strings.Split(rowString, "")

		board = append(board, row)
	}

	return board
}

func parsePart2(inp []string) common.Graph[string] {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(board common.Graph[string]) int {
	sum := 0

	board.Print()

	shiftNorth(board)

	board.Print()

	sum = sumWeights(board)

	return sum
}

func Part2(board common.Graph[string]) int {
	const limit = 1000000000
	sum := 0

	board.Print()

	seen := make(map[string]int)
	reverse := make(map[int]string)

	str := badString(board)

	seen[str] = 0
	reverse[0] = str

	start := 0
	diff := 0

	for i := 0; i < limit; i++ {

		shiftNorthNew(board)
		shiftWestNew(board)
		shiftSouthNew(board)
		shiftEastNew(board)

		str = badString(board)

		if when, present := seen[str]; present {
			common.Log.Debug().Int("seen", i).Int("when", when).Send()

			start = when
			diff = i - when

			break
		}
		seen[str] = i
		reverse[i] = str

	}

	common.Log.Debug().Int("start", start).Int("diff", diff).Send()

	cycles := (limit - start) / diff
	done := cycles * diff
	left := limit - start - done
	common.Log.Debug().Int("cycles", cycles).Int("done", done).Int("left", left).Send()

	targetBoard := recreate(reverse[start+left-1], len(board), len(board[0]))

	targetBoard.Print()

	sum = sumWeights(targetBoard)

	return sum
}

func badString(board common.Graph[string]) string {
	res := strings.Builder{}

	for _, row := range board {
		for _, cell := range row {
			res.WriteString(cell)
		}
	}

	return res.String()
}

func recreate(badString string, m, n int) common.Graph[string] {
	board := make(common.Graph[string], 0)

	for i := 0; i < n; i++ {

		row := make([]string, 0)
		for j := 0; j < m; j++ {
			row = append(row, string(badString[(i*n)+j]))
		}

		board = append(board, row)

	}

	return board
}

func sumWeights(board common.Graph[string]) int {
	sum := 0

	for i := 0; i < len(board); i++ {
		count := 0
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "O" {
				count++
			}
		}

		sum += (len(board) - i) * count
	}

	return sum
}

func shiftNorthNew(board common.Graph[string]) {

	shift(board,
		0, len(board[0]),
		0, limitMax(len(board)-1), inc,
		innerDiff(same), outerDiff(same),
		innerDiff(inc), outerDiff(same),
	)

}

func shiftSouthNew(board common.Graph[string]) {

	shift(board,
		0, len(board[0]),
		len(board)-1, limitMin(0), dec,
		innerDiff(same), outerDiff(same),
		innerDiff(dec), outerDiff(same),
	)

}

func shiftWestNew(board common.Graph[string]) {

	shift(board,
		0, len(board),
		0, limitMax(len(board[0])-1), inc,
		outerDiff(same), innerDiff(same),
		outerDiff(same), innerDiff(inc),
	)

}

func shiftEastNew(board common.Graph[string]) {

	shift(board,
		0, len(board),
		len(board[0])-1, limitMin(0), dec,
		outerDiff(same), innerDiff(same),
		outerDiff(same), innerDiff(dec),
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

func dec(i int) int {
	return i - 1
}

func inc(i int) int {
	return i + 1
}

func same(i int) int {
	return i
}

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
