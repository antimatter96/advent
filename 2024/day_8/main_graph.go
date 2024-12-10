package main

import (
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

type InputTwo struct {
	points map[string][]common.Point
	m, n   int

	graph common.Graph[string]
}

func parsePart1Two(inp []string) Input {
	mp := make(map[string][]common.Point)
	graph := make([][]string, len(inp))

	for i, row := range inp {
		graph[i] = make([]string, len(row))
		for j, e := range strings.Split(row, "") {
			graph[i][j] = e

			if e == "." {
				continue
			}
			if mp[e] == nil {
				mp[e] = make([]common.Point, 0)
			}
			mp[e] = append(mp[e], common.Point{X: i, Y: j})
		}
	}

	return Input{points: mp, m: len(inp), n: len(inp[0]), graph: graph}
}

func parsePart2Two(inp []string) Input {
	return parsePart1Two(inp)
}

func RunTwo(inp []string) (int, int) {
	parsedPart1 := parsePart1Two(inp)
	parsedPart2 := parsePart2Two(inp)

	return Part1Two(parsedPart1), Part2Two(parsedPart2)
}

func Part1Two(inp Input) int {
	var sum int

	for _, points := range inp.points {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				ext1, ext2 := getTwoExtremes(points[i], points[j])

				addToGraph_Two(inp.graph, ext1, inp.m, inp.n)
				addToGraph_Two(inp.graph, ext2, inp.m, inp.n)

				// inp.graph.Print()
			}
		}
	}

	for _, row := range inp.graph {
		for _, e := range row {
			if e == "#" {
				sum += 1
			}
		}
	}

	return sum
}

func Part2Two(inp Input) int {
	var sum int

	for _, points := range inp.points {
		for i := 0; i < len(points); i++ {
			// addToGraph(mp, points[i], inp.m, inp.n, inp.graph)

			for j := i + 1; j < len(points); j++ {
				// addToGraph(mp, points[j], inp.m, inp.n, inp.graph)

				newPoints := getExtremes(points[i], points[j], inp.m, inp.n)

				for _, newPoint := range newPoints {
					addToGraph_Two(inp.graph, newPoint, inp.m, inp.n)
				}
			}
		}
	}

	for _, row := range inp.graph {
		for _, e := range row {
			if e == "#" {
				sum += 1
			}
		}
	}

	return sum
}

func addToGraph_Two(graph common.Graph[string], point common.Point, m, n int) {
	if point.X >= m || point.X < 0 {
		return
	}
	if point.Y >= n || point.Y < 0 {
		return
	}

	graph[point.X][point.Y] = "#"
}
