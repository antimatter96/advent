package main

import (
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) common.Graph[int] {
	graph := make([][]int, len(inp))
	for i, row := range inp {
		graph[i] = make([]int, 0, len(inp[0]))
		for _, e := range strings.Split(row, "") {
			x, _ := strconv.Atoi(e)
			graph[i] = append(graph[i], x)
		}
	}

	return graph
}

func parsePart2(inp []string) common.Graph[int] {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp common.Graph[int]) int {
	var sum int

	for x := range inp {
		for y := range inp[x] {
			if inp[x][y] == 0 {
				points := reachNine(inp, x, y)

				mp := make(map[string]struct{})
				for _, point := range points {
					mp[point] = struct{}{}
				}

				sum += len(mp)
			}
		}
	}

	return sum
}

func Part2(inp common.Graph[int]) int {
	var sum int

	for x := range inp {
		for y := range inp[x] {
			if inp[x][y] == 0 {
				points := reachNine(inp, x, y)
				sum += len(points)
			}
		}
	}

	return sum
}

func reachNine(g common.Graph[int], x, y int) []string {
	current := g[x][y]
	if current == 9 {
		return []string{common.FormatAsPointString(x, y)}
	}

	possible := make([]string, 0)
	if g.InBounds(x+1, y) && g[x+1][y] == current+1 {
		possible = append(possible, reachNine(g, x+1, y)...)
	}
	if g.InBounds(x, y+1) && g[x][y+1] == current+1 {
		possible = append(possible, reachNine(g, x, y+1)...)
	}
	if g.InBounds(x-1, y) && g[x-1][y] == current+1 {
		possible = append(possible, reachNine(g, x-1, y)...)
	}
	if g.InBounds(x, y-1) && g[x][y-1] == current+1 {
		possible = append(possible, reachNine(g, x, y-1)...)
	}

	return possible
}
