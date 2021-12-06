package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	rawInput := takeInput()

	fmt.Println(Run(rawInput))
}

func takeInput() []string {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return inp
}

type line struct {
	x1, y1 int
	x2, y2 int
}

func parsePart1(inp []string) []line {
	lines := make([]line, len(inp))

	for i, l := range inp {
		fmt.Sscanf(l, "%d,%d -> %d,%d", &lines[i].x1, &lines[i].y1, &lines[i].x2, &lines[i].y2)
	}

	return lines
}

// func parsePart2(inp []string) ([]int, []board) {
// 	return parsePart1(inp)
// }

func Run(inp []string) (int, int) {
	// for i, l := range inp {
	// 	fmt.Println(i, l)
	// }
	lines := parsePart1(inp)
	//parsedPart2 := parsePart2(inp)

	return Part1(lines), Part2(lines)
	//	return 0, 0
}

// type point struct {
// 	x, y int
// }

func point(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getPoints(line line) []string {
	var points []string
	if line.x1 == line.x2 {
		for j := line.y1; j <= line.y2; j++ {
			points = append(points, point(line.x1, j))
		}
		for j := line.y2; j <= line.y1; j++ {
			points = append(points, point(line.x1, j))
		}
	} else if line.y1 == line.y2 {
		for i := line.x1; i <= line.x2; i++ {
			points = append(points, point(i, line.y1))
		}
		for i := line.x2; i <= line.x1; i++ {
			points = append(points, point(i, line.y1))
		}
	}
	return points
}

func getPoints2(line line) []string {
	var points []string
	if line.x1 == line.x2 {
		start, end := line.y1, line.y2
		if end < start {
			start, end = line.y2, line.y1
		}
		for j := start; j <= end; j++ {
			points = append(points, point(line.x1, j))
		}
	} else if line.y1 == line.y2 {
		start, end := line.x1, line.x2
		if end < start {
			start, end = line.x2, line.x1
		}

		for i := start; i <= end; i++ {
			points = append(points, point(i, line.y1))
		}
	} else {

		diffX := line.x2 - line.x1
		diffY := line.y2 - line.y1

		steps := int(math.Abs(float64(diffX)))

		diffX /= steps
		diffY /= steps

		for i, k := line.x1, 0; k <= steps; i, k = i+diffX, k+1 {
			points = append(points, point(i, line.y1+(k*diffY)))
		}

	}
	return points
}

func Part1(lines []line) int {
	mp := make(map[string]int)

	for _, l := range lines {
		points := getPoints(l)

		for _, point := range points {
			mp[point]++
		}
	}

	count := 0

	for _, v := range mp {
		if v > 1 {
			count++
		}
	}

	return count
}

func Part2(lines []line) int {
	mp := make(map[string]int)

	for _, l := range lines {
		points := getPoints2(l)

		for _, point := range points {
			mp[point]++
		}
	}

	count := 0

	for _, v := range mp {
		if v > 1 {
			count++
		}
	}

	return count
}
