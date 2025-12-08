package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	var p1, p2 int

	p1, p2 = Run(rawInput)

	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type Point struct {
	X, Y, Z int
	XYZ     string
}

type Points []Point

func parsePart1(inp []string) (Points, int) {
	var wiresToAdd int
	fmt.Sscanf(inp[0], "%d", &wiresToAdd)

	points := make([]Point, len(inp)-1)

	for i, row := range inp[1:] {
		split := strings.Split(row, ",")

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		points[i] = Point{XYZ: row, X: x, Y: y, Z: z}
	}

	return points, wiresToAdd
}

func parsePart2(inp []string) (Points, int) {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1, wiresToAdd1 := parsePart1(inp)
	parsedPart2, wiresToAdd2 := parsePart2(inp)

	return Part1(parsedPart1, wiresToAdd1), Part2(parsedPart2, wiresToAdd2)
}

type distance struct {
	p1, p2 *Point
	value  float64
}

func Part1(points Points, wiresToAdd int) int {
	distances := calculateDistances(points)

	union := common.Union[string]{}

	strs := make([]string, len(points))
	for i := range points {
		strs[i] = points[i].XYZ
	}

	union.Init(strs)
	for i := range wiresToAdd {
		common.Log.Debug().Str("a", distances[i].p1.XYZ).Str("b", distances[i].p2.XYZ).Msg("Adding")
		union.Add(distances[i].p1.XYZ, distances[i].p2.XYZ)
	}

	buckets := union.Bucketize()
	sizes := make([]int, 0)
	for p, kids := range buckets {
		common.Log.Debug().Str("Parent", p).Int("Kids", len(kids)).Send()

		if len(kids) > 1 {
			sizes = append(sizes, len(kids))
		}
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)

	return sizes[0] * sizes[1] * sizes[2]
}

func Part2(points Points, wiresToAdd int) int {
	var ans int

	distances := calculateDistances(points)

	strs := make([]string, len(points))
	for i := range points {
		strs[i] = points[i].XYZ
	}

	nPoints := len(points)
	union := common.Union[string]{}
	union.Init(strs)
	for i := range len(distances) {
		common.Log.Debug().Str("a", distances[i].p1.XYZ).Str("b", distances[i].p2.XYZ).Msg("Adding")
		union.Add(distances[i].p1.XYZ, distances[i].p2.XYZ)

		children := union.Size()
		if children == nPoints {
			ans = distances[i].p1.X * distances[i].p2.X
			break
		}
	}

	return ans
}

func calculateDistances(points []Point) []distance {
	distances := make([]distance, 0, len(points)*(len(points)-1)/2)

	var x_2, y_2, z_2, value float64

	for i := range points {
		point1 := &points[i]

		for j := i + 1; j < len(points); j++ {
			point2 := &points[j]

			if point1.XYZ == point2.XYZ {
				continue
			}

			x_2 = math.Pow(float64(point1.X-point2.X), 2)
			y_2 = math.Pow(float64(point1.Y-point2.Y), 2)
			z_2 = math.Pow(float64(point1.Z-point2.Z), 2)
			value = math.Sqrt(x_2 + y_2 + z_2)

			distances = append(distances, distance{p1: point1, p2: point2, value: value})
		}
	}

	slices.SortFunc(distances, func(a distance, b distance) int {
		return int(a.value - b.value)
	})

	return distances
}
