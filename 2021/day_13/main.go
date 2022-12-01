package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

type fold struct {
	along string
	k     int
}

type point struct {
	x, y int
}

func parsePart1(inp []string) ([]point, []fold) {
	folds := make([]fold, 0)
	points := make([]point, 0)

	for _, s := range inp {
		s = strings.TrimSpace(s)
		nums := strings.Split(s, ",")

		if len(nums) == 2 {
			var temp point
			fmt.Sscanf(s, "%d,%d", &temp.x, &temp.y)
			points = append(points, temp)
		} else if len(s) > 0 {
			s = strings.ReplaceAll(s, "=", " ")
			var temp fold
			fmt.Sscanf(s, "fold along %s %d", &temp.along, &temp.k)
			folds = append(folds, temp)
		}
	}

	return points, folds
}

func parsePart2(inp []string) ([]point, []fold) {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	p, f := parsePart1(inp)
	p2, f2 := parsePart1(inp)

	return Part1(p, f), Part2(p2, f2)
}

func (p *point) Match(fold fold) bool {
	if fold.along == "y" {
		return p.y > fold.k
	} else if fold.along == "x" {
		return p.x > fold.k
	}
	panic("unknown fold direction")
}

func (p *point) Mirror(fold fold) {
	if fold.along == "y" {
		p.y = fold.k - (p.y - fold.k)
	} else if fold.along == "x" {
		p.x = fold.k - (p.x - fold.k)
	}
}

func Part1(points []point, folds []fold) int {
	mp := make(map[point]struct{})

	for _, p := range points {
		mp[p] = struct{}{}
	}

	copied := copyMap(mp)

	for p := range mp {
		if p.Match(folds[0]) {
			delete(copied, p)
			p.Mirror(folds[0])
			copied[p] = struct{}{}
		}
	}

	return len(copied)
}

func Part2(points []point, folds []fold) int {
	mp := make(map[point]struct{})

	for _, p := range points {
		mp[p] = struct{}{}
	}

	for _, fold := range folds {
		copied := copyMap(mp)

		for p := range mp {
			if p.Match(fold) {
				delete(copied, p)
				p.Mirror(fold)
				copied[p] = struct{}{}
			}
		}

		mp = copied
	}

	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt

	for p := range mp {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}

	board := make([][]bool, 0)
	for i := 0; i < (maxX-minX)+1; i++ {
		row := make([]bool, maxY-minY+1)
		board = append(board, row)
	}

	for p := range mp {
		board[p.x][p.y] = true
	}

	for j := 0; j < maxY-minY+1; j++ {
		for i := 0; i < (maxX-minX)+1; i++ {
			if board[i][j] {
				fmt.Print("\033[1;47;49m#\033[0m")
			} else {
				fmt.Print(" ")
			}

			if (i+1)%5 == 0 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	return len(mp)
}

func copyMap(source map[point]struct{}) map[point]struct{} {
	copied := make(map[point]struct{})
	for p := range source {
		copied[p] = struct{}{}
	}
	return copied
}
