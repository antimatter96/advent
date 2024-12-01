package main

import (
	"fmt"

	"github.com/antimatter96/advent/2024/common"

	"sort"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type input struct {
	l1 []int
	l2 []int
}

func parsePart1(inp []string) input {
	common.Log.Debug().Int("Input Lengtg", len(inp)).Send()

	l1, l2 := make([]int, len(inp)), make([]int, len(inp))

	for i, s := range inp {
		fmt.Sscanf(s, "%d %d", &(l1[i]), &(l2[i]))
	}

	return input{l1, l2}
}

func parsePart2(inp []string) input {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(nums input) int {
	var sum int

	sort.Ints(nums.l1)
	sort.Ints(nums.l2)

	for i := 0; i < len(nums.l1); i++ {
		if nums.l1[i]-nums.l2[i] > 0 {
			sum += nums.l1[i] - nums.l2[i]
		} else {
			sum += nums.l2[i] - nums.l1[i]
		}
	}

	return sum
}

func Part2(nums input) int {
	var sum int

	count := make(map[int]int)
	for i := 0; i < len(nums.l2); i++ {
		count[nums.l2[i]]++
	}

	for i := 0; i < len(nums.l2); i++ {
		sum += nums.l1[i] * count[nums.l1[i]]
	}

	return sum
}
