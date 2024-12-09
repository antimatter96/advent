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

type block struct {
	occupied [][]int
	free     int
}

func parsePart1(inp []string) []block {
	blocks := make([]block, len(inp[0]))

	for i, n_s := range inp[0] {
		n := int(n_s) - 48

		if i%2 == 0 {
			blocks[i].occupied = make([][]int, 1)

			blocks[i].occupied[0] = make([]int, n)
			for j := 0; j < n; j++ {
				blocks[i].occupied[0][j] = i / 2
			}
		} else {
			blocks[i].free = n
		}
	}

	return blocks
}

func parsePart2(inp []string) []block {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp []block) int {
	var sum int

	first := 0
	last := len(inp) - 1

	for first <= last {
		for first < len(inp) && (inp[first].free == 0) {
			first++
		}
		for last > -1 && (len(inp[last].occupied) == 0 || len(inp[last].occupied[0]) == 0) {
			last--
		}

		inp[first].free--
		if inp[first].occupied == nil {
			inp[first].occupied = make([][]int, 1)
			inp[first].occupied[0] = make([]int, 0)
		}

		inp[first].occupied[0] = append(inp[first].occupied[0], inp[last].occupied[0][0])
		inp[last].occupied[0] = inp[last].occupied[0][1:]
	}

	n := 0
	for _, y := range inp {
		for _, occ := range y.occupied {
			for _, e := range occ {
				sum += n * e
				n++
			}
		}

	}
	return sum
}

func Part2(inp []block) int {
	var sum int

	return sum
}
