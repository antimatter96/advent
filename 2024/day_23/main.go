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

type connection struct {
	a, b string
}

func parsePart1(inp []string) []connection {
	network := make([]connection, 0, len(inp))

	for _, connection_s := range inp {
		nodes := strings.Split(connection_s, "-")
		network = append(network, connection{a: nodes[0], b: nodes[1]})
	}

	return network
}

func parsePart2(inp []string) []connection {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp []connection) int {
	var sum int

	g := make(map[string]map[string]bool)

	for _, conn := range inp {
		if g[conn.a] == nil {
			g[conn.a] = make(map[string]bool)
		}
		if g[conn.b] == nil {
			g[conn.b] = make(map[string]bool)
		}

		g[conn.a][conn.b] = true
		g[conn.b][conn.a] = true
	}

	foundNodes := make(map[string]bool, 0)

	for a, connectionsOfA := range g {
		for b := range connectionsOfA {
			for connectionsOfC := range g[b] {
				if a == connectionsOfC {
					continue
				}

				if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(connectionsOfC, "t") {
					if connectionsOfA[connectionsOfC] {
						x := []string{a, b, connectionsOfC}
						slices.Sort(x)
						foundNodes[x[0]+","+x[1]+","+x[2]] = true
					}
				}
			}
		}
	}

	sum = len(foundNodes)

	return sum
}

func Part2(inp []connection) int {
	var sum int

	return sum
}
