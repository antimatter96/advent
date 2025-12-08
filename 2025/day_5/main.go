package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type problemInput struct {
	freshIDSets  []freshIDSet
	availableIDs []int
}

type freshIDSet struct {
	low  int
	high int
}

func parsePart1(inp []string) problemInput {
	n := len(inp)
	i := 0

	freshIDSets := make([]freshIDSet, 0, len(inp))
	for ; i < n; i++ {
		if inp[i] == "" {
			i++
			break
		}
		freshIDSet := freshIDSet{}
		fmt.Sscanf(inp[i], "%d-%d", &(freshIDSet.low), &(freshIDSet.high))
		freshIDSets = append(freshIDSets, freshIDSet)
	}

	availableIDs := make([]int, 0, n-len(freshIDSets))
	for ; i < n; i++ {
		x, _ := strconv.Atoi(inp[i])
		availableIDs = append(availableIDs, x)
	}

	return problemInput{freshIDSets: freshIDSets, availableIDs: availableIDs}
}

func parsePart2(inp []string) problemInput {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp problemInput) int {
	var sum int

	inp.freshIDSets = compressIDSets(inp.freshIDSets)

	for _, id := range inp.availableIDs {

		for _, freshIDSet := range inp.freshIDSets {
			if id >= freshIDSet.low && id <= freshIDSet.high {
				sum++
				break
			}
		}
	}

	return sum
}

func Part2(inp problemInput) int {
	var sum int

	inp.freshIDSets = compressIDSets(inp.freshIDSets)

	for _, freshIDSet := range inp.freshIDSets {
		sum += (freshIDSet.high - freshIDSet.low + 1)
	}
	return sum
}

func compressIDSetsBruteForce(sets []freshIDSet) []freshIDSet {
	sort.Slice(sets, func(i int, j int) bool {
		if sets[i].low == sets[j].low {
			return sets[i].high > sets[j].high
		}
		return sets[i].low < sets[j].low
	})

	compressed := make([]freshIDSet, 0, len(sets))
	compressed = append(compressed, sets...)

	return compressed
}

func compressIDSets(sets []freshIDSet) []freshIDSet {
	sort.Slice(sets, func(i int, j int) bool {
		if sets[i].low == sets[j].low {
			return sets[i].high < sets[j].high
		}
		return sets[i].low < sets[j].low
	})

	compressed := make([]freshIDSet, 0, len(sets))
	for i := 0; i < len(sets)-1; i++ {
		if sets[i].high < sets[i+1].low {
			compressed = append(compressed, sets[i])
			continue
		}

		if sets[i].high >= sets[i+1].low {
			common.Log.Debug().Str("from", fmt.Sprintf("[%d - %d] [%d - %d]", sets[i].low, sets[i].high, sets[i+1].low, sets[i+1].high)).Str("to", fmt.Sprintf("[%d - %d]", sets[i].low, sets[i+1].high)).Msg("Compressing")
			sets[i+1].low = sets[i].low

			if sets[i].high >= sets[i+1].high {
				sets[i+1].high = sets[i].high
			}
		}
	}

	compressed = append(compressed, sets[len(sets)-1])

	common.Log.Debug().Int("from", len(sets)).Int("to", len(compressed)).Msg("Compressed")
	return compressed
}
