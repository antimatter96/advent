package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type node struct {
	left  string
	right string
}

type networkMap struct {
	instructions string
	nodes        map[string]node
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput, 0)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) *networkMap {
	networkMap := &networkMap{}

	for i := 0; i < len(inp); i++ {
		inp[i] = strings.TrimSpace(inp[i])
	}

	networkMap.instructions = inp[0]

	networkMap.nodes = make(map[string]node, 0)

	for i := 1; i < len(inp); i++ {
		if inp[i] == "" || inp[i] == " " {
			continue
		}

		node := node{}
		var parent string

		inp[i] = strings.ReplaceAll(inp[i], ")", "")
		fmt.Sscanf(strings.ReplaceAll(inp[i], ",", ""), `%s = (%s %s)`, &parent, &(node.left), &(node.right))

		networkMap.nodes[parent] = node

	}

	return networkMap
}

func parsePart2(inp []string) *networkMap {
	return parsePart1(inp)
}

func Run(inp []string, mode int) (int, int) {
	parsedPart1 := parsePart1(inp)
	// parsedPart2 := parsePart2(inp)

	switch mode {
	case 0:
		return Part1(parsedPart1), Part2(parsedPart1)
	case 1:
		return Part1(parsedPart1), 0
	case 2:
		return 0, Part2(parsedPart1)
	}

	return 0, 0
}

func Part1(networkMap *networkMap) int {
	for parent, node := range networkMap.nodes {
		common.Log.Debug().Str("A", parent).Str("L", node.left).Str("R", node.right).Send()
	}

	start := "AAA"

	count := 0
	for start != "ZZZ" {

		common.Log.Debug().Str("start", start).Send()

		for _, r := range networkMap.instructions {
			if r == 'L' {
				start = networkMap.nodes[start].left
			} else {
				start = networkMap.nodes[start].right
			}

			count++
		}
	}

	return count
}

func Part2(networkMap *networkMap) int {
	for parent, node := range networkMap.nodes {
		common.Log.Debug().Str("A", parent).Str("L", node.left).Str("R", node.right).Send()
	}

	starts := []string{}

	for k := range networkMap.nodes {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

	EndWithZ := func(start string) bool {
		return start[2] != 'Z'
	}

	counts := []int{}

	for i := 0; i < len(starts); i++ {
		start := starts[i]
		count := 0

		for EndWithZ(start) {
			common.Log.Debug().Str("start", start).Send()

			for _, r := range networkMap.instructions {
				if r == 'L' {
					start = networkMap.nodes[start].left
				} else {
					start = networkMap.nodes[start].right
				}
				count++
			}
		}

		common.Log.Debug().Str("start", starts[i]).Str("end", start).Int("count", count).Send()
		counts = append(counts, count)
	}

	common.Log.Debug().Ints("individual counts", counts).Send()

	lcm := common.LCMs(counts, 0)

	return lcm
}

// func DetermineLast() int {

// }
