package main

import (
	"fmt"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

var M int
var LOOPS int

var emptyArray []string
var fullArray []string

func parsePart1(inp []string) []common.Point {
	fmt.Sscanf(inp[0], "%d", &M)
	fmt.Sscanf(inp[1], "%d", &LOOPS)

	points := make([]common.Point, len(inp)-2)
	for i := 2; i < len(inp); i++ {
		points[i-2] = common.Point{}

		fmt.Sscanf(inp[i], "%d,%d", &points[i-2].X, &points[i-2].Y)
	}

	return points
}

func parsePart2(inp []string) []common.Point {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {

	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	emptyArray = make([]string, M+1)
	fullArray = make([]string, M+1)
	for i := 0; i < M+1; i++ {
		fullArray[i] = "."
	}

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(points []common.Point) int {
	graph := make(common.Graph[string], M+1)
	for i := 0; i < M+1; i++ {
		graph[i] = make([]string, M+1)
		copy(graph[i], fullArray)
	}

	var sum int

	graph.Print()

	for i := 0; i < LOOPS; i++ {
		point := points[i]
		graph[point.Y][point.X] = "#"
	}

	graph.Print()

	queue := common.QueueSet[string]{}
	queue.Push((&common.Point{X: 0, Y: 0}).String())

	dist := make(map[string]int, 0)
	dist["0,0"] = 0
	goDiagnol := false
bfs:
	for !queue.Empty() {
		p := queue.Pop()

		pp := common.Point{}
		pp.FromString(p)

		for incX := -1; incX < 2; incX++ {
			x := pp.X + incX
			for incY := -1; incY < 2; incY++ {
				y := pp.Y + incY

				if !goDiagnol && incX*incY != 0 {
					continue
				}

				if x == 0 && y == 0 {
					continue
				}

				if graph.InBounds(x, y) && graph[x][y] == "." {
					str := (&common.Point{X: x, Y: y}).String()
					if dist[str] == 0 {
						queue.Push(str)
						dist[(&common.Point{X: x, Y: y}).String()] = dist[p] + 1
					}

					if x == M && y == M {
						fmt.Println(dist[str])
						break bfs
					}
				}
			}
		}

	}

	fmt.Println(dist)

	return sum
}

func Part2(points []common.Point) int {
	return 0
}
