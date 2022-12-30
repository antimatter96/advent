package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp []string) [][]int {
	forest := make([][]int, 0)

	for _, row := range inp {
		split := strings.Split(row, "")

		intRow := make([]int, 0)

		for _, tree := range split {
			treeInt, _ := strconv.Atoi(tree)

			intRow = append(intRow, treeInt)
		}

		forest = append(forest, intRow)

	}

	return forest
}

func parsePart2(inp []string) [][]int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func genFromArrays(forest [][]int) ([][]int, [][]int, [][]int, [][]int) {
	fromT := make([][]int, 0)
	fromB := make([][]int, 0)
	fromL := make([][]int, 0)
	fromR := make([][]int, 0)

	for _, l := range forest {
		fromT = append(fromT, make([]int, len(l)))
		fromB = append(fromB, make([]int, len(l)))
		fromL = append(fromL, make([]int, len(l)))
		fromR = append(fromR, make([]int, len(l)))
	}

	for i := 0; i < len(forest); i++ {
		for j := 1; j < len(forest[0]); j++ {
			fromL[i][j] = forest[i][j-1]
			fromT[j][i] = forest[j-1][i]
		}
		for j := len(forest) - 2; j > -1; j-- {
			fromR[i][j] = forest[i][j+1]
			fromB[j][i] = forest[j+1][i]
		}
	}

	for i := 0; i < len(forest); i++ {
		for j := 1; j < len(forest[0]); j++ {
			if fromL[i][j] <= fromL[i][j-1] {
				fromL[i][j] = fromL[i][j-1]
			}

			if fromT[j][i] <= fromT[j-1][i] {
				fromT[j][i] = fromT[j-1][i]
			}
		}

		for j := len(forest) - 2; j > -1; j-- {
			if fromR[i][j] <= fromR[i][j+1] {
				fromR[i][j] = fromR[i][j+1]
			}
			if fromB[j][i] <= fromB[j+1][i] {
				fromB[j][i] = fromB[j+1][i]
			}
		}
	}

	return fromT, fromL, fromB, fromR
}

func Part1(forest [][]int) int {
	fromT, fromL, fromB, fromR := genFromArrays(forest)

	visibleInside := 0

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[0])-1; j++ {
			if forest[i][j] > fromL[i][j] || forest[i][j] > fromR[i][j] || forest[i][j] > fromT[i][j] || forest[i][j] > fromB[i][j] {
				visibleInside++
			}
		}
	}

	return visibleInside + (4*len(forest) - 4)
}

func Part2(forest [][]int) int {
	fromT, fromL, fromB, fromR := genFromArrays(forest)

	mMax := 0

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[0])-1; j++ {
			if forest[i][j] > fromL[i][j] || forest[i][j] > fromR[i][j] || forest[i][j] > fromT[i][j] || forest[i][j] > fromB[i][j] {
				l, r, t, b := 1, 1, 1, 1

				for ; (j + r) < len(forest); r++ {
					if forest[i][j] > forest[i][j+r] {
						continue
					} else {
						break
					}
				}
				if j+r == len(forest) {
					r--
				}

				for ; (j - l) > -1; l++ {
					if forest[i][j] > forest[i][j-l] {
						continue
					} else {
						break
					}
				}
				if j == l-1 {
					l--
				}

				for ; (i - t) > -1; t++ {
					if forest[i][j] > forest[i-t][j] {
						continue
					} else {
						break
					}
				}
				if i == t-1 {
					t--
				}

				for ; (i + b) < len(forest); b++ {
					if forest[i][j] > forest[i+b][j] {
						continue
					} else {
						break
					}
				}
				if i+b == len(forest) {
					b--
				}

				if l*r*b*t > mMax {
					mMax = l * r * t * b
				}
			}
		}
	}

	return mMax
}
