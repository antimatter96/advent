package main

import (
	"slices"
	"strconv"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) []int {
	initialCodes := make([]int, 0, len(inp))

	for _, initialCode_s := range inp {
		initialCode, _ := strconv.Atoi(initialCode_s)
		initialCodes = append(initialCodes, initialCode)
	}

	return initialCodes
}

func parsePart2(inp []string) []int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

var memTable map[int]int

func Part1(inp []int) int {
	var sum int
	slices.Sort(inp)

	memTable = make(map[int]int)

	for _, secret := range inp {
		x := secret
		for i := 0; i < 2000; i++ {
			x = findNext(x)
		}
		sum += x
	}

	return sum
}

func Part2(inp []int) int {
	var sum int

	return sum
}

func findNext(in int) int {
	if next, ok := memTable[in]; ok {
		return next
	}

	out := in + 0

	intermediate := out * 64
	out = mix(out, intermediate)
	out = prune(out)

	intermediate = out / 32
	out = mix(out, intermediate)
	out = prune(out)

	intermediate = out * 2048
	out = mix(out, intermediate)
	out = prune(out)

	memTable[in] = out
	return out
}

func mix(in int, mixer int) int {
	return in ^ mixer
}

func prune(in int) int {
	return in % 16777216
}
