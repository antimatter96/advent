package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsString()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp string) [][]int {
	arr := make([][]int, 0)

	inpTokens := strings.Split(strings.TrimLeft(inp, "\n"), "\n")

	var temp int
	tempArr := make([]int, 0)
	for _, inpToken := range inpTokens {
		if len(inpToken) == 0 {
			arr = append(arr, tempArr)
			tempArr = make([]int, 0)
			continue
		}

		temp, _ = strconv.Atoi(inpToken)
		tempArr = append(tempArr, temp)
	}

	return arr
}

func parsePart2(inp string) [][]int {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(elfs [][]int) int {
	max := 0

	var sum int

	for _, weights := range elfs {

		sum = 0

		for _, weight := range weights {
			sum += weight
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

// Slowest
func Part2(elfs [][]int) int {
	var sum int
	sums := make([]int, 0)

	for _, weights := range elfs {

		sum = 0

		for _, weight := range weights {
			sum += weight
		}

		sums = append(sums, sum)

	}

	sort.Ints(sums)

	l := len(sums)

	return sums[l-1] + sums[l-2] + sums[l-3]
}
