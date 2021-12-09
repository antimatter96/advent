package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

func parsePart1(inp []string) [][]int {
	timers := make([][]int, 0)

	for _, s := range inp {
		nums := strings.Split(s, "")

		row := make([]int, 0)

		for _, nums := range nums {
			temp, _ := strconv.Atoi(nums)
			row = append(row, temp)
		}

		timers = append(timers, row)
	}

	return timers
}

func parsePart2(inp []string) [][]int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	numbers := parsePart1(inp)

	return Part1(numbers), Part2(numbers)
}

func Part1(grid [][]int) int {
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isLowest(i, j, grid) {
				sum += grid[i][j] + 1
			}
		}
	}
	return sum
}

func isLowest(i, j int, grid [][]int) bool {
	if getSafe(i, j-1, grid, math.MaxInt) <= grid[i][j] {
		return false
	}
	if getSafe(i, j+1, grid, math.MaxInt) <= grid[i][j] {
		return false
	}
	if getSafe(i-1, j, grid, math.MaxInt) <= grid[i][j] {
		return false
	}
	if getSafe(i+1, j, grid, math.MaxInt) <= grid[i][j] {
		return false
	}

	return true
}

func getSafe(i, j int, grid [][]int, fallback int) int {
	if i < 0 || i >= len(grid) {
		return fallback
	}

	if j < 0 || j >= len(grid[i]) {
		return fallback
	}

	return grid[i][j]
}

func Part2(grid [][]int) int {
	lowPoints := make([][]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isLowest(i, j, grid) {
				lowPoints = append(lowPoints, []int{i, j})
			}
		}
	}

	for i, lowPoint := range lowPoints {
		floodFill(lowPoint[0], lowPoint[1], grid, -(i + 1))
	}

	cnt := make(map[int]int)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cnt[grid[i][j]]++
		}
	}

	basinSizes := make([]int, 0)

	for k, v := range cnt {
		if k < 0 {
			basinSizes = append(basinSizes, v)
		}
	}

	sort.Ints(basinSizes)
	z := len(basinSizes)

	return basinSizes[z-1] * basinSizes[z-2] * basinSizes[z-3]
}

func floodFill(i, j int, grid [][]int, toFill int) {
	was := grid[i][j]
	grid[i][j] = toFill

	if val := getSafe(i, j-1, grid, math.MinInt); val != 9 && val >= was {
		floodFill(i, j-1, grid, toFill)
	}
	if val := getSafe(i, j+1, grid, math.MinInt); val != 9 && val >= was {
		floodFill(i, j+1, grid, toFill)
	}
	if val := getSafe(i-1, j, grid, math.MinInt); val != 9 && val >= was {
		floodFill(i-1, j, grid, toFill)
	}
	if val := getSafe(i+1, j, grid, math.MinInt); val != 9 && val >= was {
		floodFill(i+1, j, grid, toFill)
	}

}
