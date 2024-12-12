package main

import (
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) common.Graph[string] {
	graph := make([][]string, len(inp))
	for i := 0; i < len(inp); i++ {
		graph[i] = strings.Split(inp[i], "")
	}

	return graph
}

func parsePart2(inp []string) common.Graph[string] {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp common.Graph[string]) int {
	var sum int

	for _, points := range common.FloodFill(inp, false, ".") {
		perimter := 0
		for _, point := range points {
			perimter += addsToPerimeter(inp, point.X, point.Y)
		}

		common.Log.Debug().Str(string(inp[points[0].X][points[0].Y]), "").Int("a", len(points)).Int("p", perimter).Send()

		sum += perimter * len(points)
	}

	return sum
}

func Part2(inp common.Graph[string]) int {
	var sum int

	return sum
}

func addsToPerimeter[T comparable](inp common.Graph[T], i, j int) int {
	perimeter := 4

	y := j - 1
	if y > -1 && inp[i][j] == inp[i][y] {
		perimeter--
	}

	y = j + 1
	if y < len(inp[0]) && inp[i][j] == inp[i][y] {
		perimeter--
	}

	x := i - 1
	if x > -1 && inp[i][j] == inp[x][j] {
		perimeter--
	}

	x = i + 1
	if x < len(inp) && inp[i][j] == inp[x][j] {
		perimeter--
	}

	return perimeter
}
