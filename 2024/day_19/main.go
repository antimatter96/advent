package main

import (
	"slices"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type input struct {
	patterns []string
	want     []string
}

func parsePart1(inp []string) input {
	patterns := strings.Split(strings.TrimSpace(inp[0]), ", ")

	return input{want: inp[2:], patterns: patterns}
}

func parsePart2(inp []string) input {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func sortByLen(a string, b string) int {
	if len(a) == len(b) {
		return strings.Compare(a, b)
	}
	return len(a) - len(b)
}

func Part1(inp input) int {
	var sum int

	slices.SortFunc(inp.patterns, sortByLen)
	slices.SortFunc(inp.want, sortByLen)

	memoizer := make(map[string]bool, len(inp.patterns))
	for _, pattern := range inp.patterns {
		memoizer[pattern] = true
	}

	for _, want := range inp.want {
		if possible(want, inp.patterns, memoizer) {
			sum += 1
		}
	}

	return sum
}

func Part2(inp input) int {
	var sum int

	slices.SortFunc(inp.patterns, sortByLen)
	slices.SortFunc(inp.want, sortByLen)

	memoizer := make(map[string]int, len(inp.patterns))

	for _, pattern := range inp.patterns {
		if len(pattern) == 1 {
			memoizer[pattern] = 1
		}
	}
	for _, pattern := range inp.patterns {
		if len(pattern) != 1 {
			_ = possibleCount(pattern, inp.patterns, memoizer)
		}
	}

	for _, want := range inp.want {
		sum += possibleCount(want, inp.patterns, memoizer)
	}

	return sum
}

func possible(want string, have []string, memoizer map[string]bool) bool {
	if want == "" {
		return true
	}
	if ans, has := memoizer[want]; has {
		return ans
	}

	for _, pattern := range have {
		if strings.HasPrefix(want, pattern) {
			left := want[len(pattern):]
			can := possible(left, have, memoizer)
			if can {
				memoizer[want] = true
				return can
			}
		}
	}
	memoizer[want] = false
	return false
}

func possibleCount(want string, have []string, memoizer map[string]int) int {
	if want == "" {
		return 1
	}
	if ans, has := memoizer[want]; has {
		return ans
	}

	for _, pattern := range have {
		if strings.HasPrefix(want, pattern) {
			left := want[len(pattern):]
			memoizer[want] += possibleCount(left, have, memoizer)
		}
	}

	return memoizer[want]
}
