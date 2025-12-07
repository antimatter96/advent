package main

import (
	"strings"

	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

var strToInt map[string]int = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

func parsePart1(inp []string) [][]int {
	common.Log.Debug().Int("Input Lengtg", len(inp)).Send()

	numbers := make([][]int, len(inp))

	for i, line := range inp {
		numbers[i] = make([]int, len(line))
		split := strings.Split(line, "")

		for j, char := range split {
			numbers[i][j] = strToInt[char]
		}
	}
	return numbers
}

func parsePart2(inp []string) [][]int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(nums [][]int) int {
	var sum int

	for _, num := range nums {
		sum += bruteForce(num)
	}

	return sum
}

func Part2(nums [][]int) int {
	var sum int

	return sum
}

func bruteForce(num []int) int {
	var max int
	var tens int
	max = 0
	for i := range num {
		tens = num[i] * 10
		for j := i + 1; j < len(num); j++ {
			if tens+num[j] > max {
				max = tens + num[j]
			}
		}
	}
	return max
}
