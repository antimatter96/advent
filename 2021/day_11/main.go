package main

import (
	"bufio"
	"fmt"
	"os"
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
	numbers2 := parsePart1(inp)

	return Part1(numbers), Part2(numbers2)
}

func incrementSafe(i, j int, grid [][]int) {
	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[i]) {
		return
	}

	// only one flash per step
	if grid[i][j] == -1 {
		return
	}

	grid[i][j]++
}

func Part1(grid [][]int) int {
	days := 100
	n := 10
	sum := 0

	for day := 0; day < days; day++ {

		flashed := true

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				grid[i][j]++
			}
		}

		for flashed {
			//fmt.Println("run")
			flashed = false

			copied := copyArray(grid)

			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {

					if grid[i][j] > 9 {
						//fmt.Println("flash", i, j)
						flashed = true
						sum += 1
						copied[i][j] = -1

						for dx := -1; dx < 2; dx++ {
							for dy := -1; dy < 2; dy++ {
								if dx == 0 && dy == 0 {
									continue
								}
								incrementSafe(i+dx, j+dy, copied)
							}
						}
					}
				}
			}

			// exchange
			if flashed {
				grid = copied
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == -1 {
					grid[i][j] = 0
				}
			}
		}

		// fmt.Println("after day", day+1)
		// for i := 0; i < n; i++ {
		// 	for j := 0; j < n; j++ {
		// 		if grid[i][j] == 0 {
		// 			fmt.Print("\033[1;38;5;208m0\033[0m")
		// 		} else {
		// 			fmt.Print(grid[i][j])
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()

	}

	return sum
}

func Part2(grid [][]int) int {
	n := 10
	sum := 0

	for day := 0; ; day++ {

		flashed := true

		before := sum
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				grid[i][j]++
			}
		}

		for flashed {
			//fmt.Println("run")
			flashed = false

			copied := copyArray(grid)

			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {

					if grid[i][j] > 9 {
						flashed = true
						sum += 1
						copied[i][j] = -1

						for dx := -1; dx < 2; dx++ {
							for dy := -1; dy < 2; dy++ {
								if dx == 0 && dy == 0 {
									continue
								}
								incrementSafe(i+dx, j+dy, copied)
							}
						}
					}
				}
			}

			// exchange
			if flashed {
				grid = copied
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == -1 {
					grid[i][j] = 0
				}
			}
		}

		if sum-before == 100 {
			return day + 1
		}
	}
}

func copyArray(source [][]int) [][]int {
	copied := make([][]int, 0)
	for _, sourceRow := range source {
		row := make([]int, 0)
		row = append(row, sourceRow...)
		copied = append(copied, row)
	}
	return copied
}
