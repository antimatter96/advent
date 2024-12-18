package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	approach := common.InitApproachFlags()

	var p1, p2 int

	if *approach == 1 {
		p1, p2 = Run(rawInput)
	} else {
		p1, p2 = RunTwo(rawInput)
	}

	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type Input struct {
	points map[string][]common.Point
	m, n   int

	graph common.Graph[string]
}

func parsePart1(inp []string) Input {
	mp := make(map[string][]common.Point)

	for i, row := range inp {
		for j, e := range strings.Split(row, "") {

			if e == "." {
				continue
			}
			if mp[e] == nil {
				mp[e] = make([]common.Point, 0)
			}
			mp[e] = append(mp[e], common.Point{X: i, Y: j})

		}
	}

	return Input{points: mp, m: len(inp), n: len(inp[0])}
}

func parsePart2(inp []string) Input {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp Input) int {
	mp := make(map[string]bool)

	for _, points := range inp.points {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				ext1, ext2 := getTwoExtremes(points[i], points[j])

				addToGraph(mp, ext1, inp.m, inp.n)
				addToGraph(mp, ext2, inp.m, inp.n)

			}
		}
	}

	return len(mp)
}

func Part2(inp Input) int {

	mp := make(map[string]bool)

	for _, points := range inp.points {
		for i := 0; i < len(points); i++ {
			addToGraph(mp, points[i], inp.m, inp.n)

			for j := i + 1; j < len(points); j++ {
				// addToGraph(mp, points[j], inp.m, inp.n, inp.graph)

				newPoints := getExtremes(points[i], points[j], inp.m, inp.n)

				for _, newPoint := range newPoints {
					addToGraph(mp, newPoint, inp.m, inp.n)
				}
			}
		}
	}

	return len(mp)
}

func position(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getTwoExtremes(p1, p2 common.Point) (common.Point, common.Point) {
	common.Log.Debug().Str("msg", "Points bewteen").Interface("x", p1).Interface("y", p2).Send()

	return common.Point{X: (2 * p1.X) - p2.X, Y: (2 * p1.Y) - p2.Y}, common.Point{X: (2 * p2.X) - p1.X, Y: (2 * p2.Y) - p1.Y}
}

func getExtremes(p1, p2 common.Point, m, n int) []common.Point {
	common.Log.Debug().Str("msg", "Points bewteen").Interface("x", p1).Interface("y", p2).Send()

	points := make([]common.Point, 0)

	if p1.X == p2.X {
		for y := 0; y < n; y++ {
			points = append(points, common.Point{X: p1.X, Y: y})
		}
		return points
	} else if p1.Y == p2.Y {
		for x := 0; x < m; x++ {
			points = append(points, common.Point{X: x, Y: p1.Y})
		}
		return points
	}

	nF := float64(n)

	x := p1.X + 1
	for ; x < m; x++ {
		y := float64(p1.Y) + (float64(p2.Y-p1.Y)*float64(x-p1.X))/float64(p2.X-p1.X)

		if y > -1 && y < nF {
			if math.Trunc(y) == y {
				points = append(points, common.Point{X: x, Y: int(math.Trunc(y))})
			}
		}
	}

	x = p1.X - 1
	for ; x > -1; x-- {
		y := float64(p1.Y) + (float64(p2.Y-p1.Y)*float64(x-p1.X))/float64(p2.X-p1.X)
		if y > -1 && y < nF {
			if math.Trunc(y) == y {
				points = append(points, common.Point{X: x, Y: int(math.Trunc(y))})
			}
		}
	}

	return points
}

func addToGraph(antinodes map[string]bool, point common.Point, m, n int) {
	if point.X >= m || point.X < 0 {
		return
	}
	if point.Y >= n || point.Y < 0 {
		return
	}

	antinodes[position(point.X, point.Y)] = true
}
