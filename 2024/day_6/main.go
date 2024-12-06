package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

var (
	HASHTAG string = "#"
	DOT     string = "."
	CROSS   string = "X"
)

func parsePart1(inp []string) common.Graph[*string] {
	graph := make([][]*string, len(inp))
	for i := 0; i < len(inp); i++ {
		graph[i] = make([]*string, len(inp[0]))
	}
	for i, row := range inp {
		for j, e := range strings.Split(row, "") {
			if e == "#" {
				graph[i][j] = &HASHTAG
			} else if e == "^" {
				graph[i][j] = &CROSS
			} else {
				graph[i][j] = &DOT
			}
		}
	}

	return graph
}

func parsePart2(inp []string) common.Graph[*string] {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp common.Graph[*string]) int {
	var sum int

	x, y := -1, -1

find:
	for i := range inp {
		for j := range inp[i] {
			if inp[i][j] == &CROSS {
				x, y = i, j
				break find
			}
		}
	}

	inp, _ = moveTillCompletion(inp, x, y, "UP")

	for i := range inp {
		for j := range inp[i] {
			if inp[i][j] == &CROSS {
				sum += 1
			}
		}
	}

	return sum
}

func Part2(inp common.Graph[*string]) int {
	var sum int

	x, y := -1, -1

find:
	for i := range inp {
		for j := range inp[i] {
			if inp[i][j] == &CROSS {
				x, y = i, j
				break find
			}
		}
	}

	inp, _ = moveTillCompletion(inp, x, y, "UP")

	for i := range inp {
		for j := range inp[i] {
			if inp[i][j] == &CROSS {
				newGraph := common.CopyPointerGraph(inp)
				newGraph[i][j] = &HASHTAG

				_, err := moveTillCompletion(newGraph, x, y, "UP")
				if err != nil {
					sum++
				}
			}
		}
	}

	return sum
}

func position(x, y int, direction string) string {
	return fmt.Sprintf("%d,%d,%s", x, y, direction)
}

func moveTillCompletion(inp common.Graph[*string], x, y int, direction string) (common.Graph[*string], error) {
	visited := make(map[string]bool, len(inp)*len(inp[0])*4)
	visited[position(x, y, direction)] = true

	for inp.InBounds(x, y) {
		nextI, nextJ := common.DirectionChanges[direction].NextI(x), common.DirectionChanges[direction].NextJ(y)

		if !inp.InBounds(nextI, nextJ) {
			break
		}

		if (inp)[nextI][nextJ] == &HASHTAG {
			direction = common.RotationRight[direction]
			if visited[position(x, y, direction)] {
				return nil, fmt.Errorf("format string")
			}
			visited[position(x, y, direction)] = true
		} else {
			if visited[position(nextI, nextJ, direction)] {
				return nil, fmt.Errorf("format string")
			}
			visited[position(nextI, nextJ, direction)] = true
			inp[nextI][nextJ] = &CROSS
			x = nextI
			y = nextJ
		}
	}
	return inp, nil
}
