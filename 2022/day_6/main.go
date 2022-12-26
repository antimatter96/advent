package main

import (
	"fmt"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsString()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp string) string {

	return inp
}

func parsePart2(inp string) string {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp string) int {
	return solve(4)(inp)
}

func Part2(inp string) int {
	return solve(14)(inp)
}

func solve(charCount int) func(string) int {

	fx := func(inp string) int {
		set := common.CountedSet[string]{}

		for i := 0; i < charCount; i++ {
			set.Add(string(inp[i]))
		}

		for i := charCount; i < len(inp); i++ {
			if set.UniqueCount() == charCount {
				return i
			}
			set.Remove(string(inp[i-charCount]))
			set.Add(string(inp[i]))
		}

		return -1
	}

	return fx
}
