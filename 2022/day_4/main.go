package main

import (
	"fmt"

	"github.com/antimatter96/advent/2022/common"
)

type assignmentPair struct {
	l1, r1 int
	l2, r2 int
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp []string) []assignmentPair {

	assignments := make([]assignmentPair, 0, len(inp))

	var l1, r1, r2, l2 int
	for _, s := range inp {
		fmt.Sscanf(s, "%d-%d,%d-%d", &l1, &r1, &l2, &r2)

		assignments = append(assignments, assignmentPair{
			l1, r1,
			l2, r2,
		})
	}

	return assignments
}

func parsePart2(inp []string) []assignmentPair {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(assignments []assignmentPair) int {
	count := 0

	for _, pair := range assignments {
		if pair.l1 >= pair.l2 && pair.r1 <= pair.r2 {
			count++
		} else if pair.l2 >= pair.l1 && pair.r2 <= pair.r1 {
			count++
		}
	}
	return count

}

func Part2(assignments []assignmentPair) int {
	count := 0

	for _, pair := range assignments {
		if pair.l1 >= pair.l2 && pair.l1 <= pair.r2 {
			count++
		} else if pair.l2 >= pair.l1 && pair.l2 <= pair.r1 {
			count++
		}
	}
	return count
}
