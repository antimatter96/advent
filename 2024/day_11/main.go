package main

import (
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsString()
	rawInput = strings.TrimSpace(rawInput)

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp string) map[int]int {
	mp := make(map[int]int, len(inp))

	for _, nums_s := range strings.Split(inp, " ") {
		n, _ := strconv.Atoi(nums_s)

		mp[n] += 1
	}

	return mp
}

func parsePart2(inp string) map[int]int {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp map[int]int) int {
	var sum int

	inp = blinkNTimes(inp, 25)

	for _, count := range inp {
		sum += count
	}

	return sum
}

func Part2(inp map[int]int) int {
	var sum int

	inp = blinkNTimes(inp, 75)

	for _, count := range inp {
		sum += count
	}

	return sum
}

func blinkNTimes(initial map[int]int, N int) map[int]int {
	mp := initial
	for i := 0; i < N; i++ {
		latest := make(map[int]int, len(mp))

		for n, count := range mp {

			if n == 0 {
				latest[1] += count
				continue
			}

			s := strconv.Itoa(n)
			if len(s)%2 == 0 {
				x1, _ := strconv.Atoi(s[:len(s)/2])
				x2, _ := strconv.Atoi(s[len(s)/2:])

				latest[x1] += count
				latest[x2] += count
				continue
			}

			latest[n*2024] += count

		}

		mp = latest
	}

	return mp
}
