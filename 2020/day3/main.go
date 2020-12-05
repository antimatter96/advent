package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Errorf(scanner.Err().Error())
	}

	day1(inp)
}

const length = 8

func day1(inp []string) {
	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2},
	}

	total := 1

	for _, rd := range slopes {
		total *= getTrees(inp, rd[0], rd[1])
	}

	fmt.Println(total)

}

func getTrees(inp []string, right, down int) int {
	total := 0

	maxY := len(inp)
	maxX := len(inp[0])
	x, y := 0, 0

	//fmt.Println(x, y, string(inp[x][y]))

	for y < maxY-1 {
		x, y = nextLocation(x, y, right, down, maxX, maxY)
		//fmt.Println(inp[y])
		//fmt.Println(x, string([y][x]), "\n")
		if inp[y][x] == '#' {
			total++
		}
	}

	return total
}

func nextLocation(x, y, right, down, maxX, maxY int) (int, int) {
	x += right
	x = (x % (maxX))

	y += down
	return x, y
}

func checkLimit(min, max int, b byte, pass string) bool {
	count := 0
	for i := 0; i < len(pass); i++ {
		if pass[i] == b {
			count++
			if count > max {
				//fmt.Println("count", count, min, max, pass)
				return false
			}
		}
	}

	//fmt.Println("count", count, min, max, pass)
	if count >= min {
		return true
	}

	return false
}

func checkLimit2(min, max int, b byte, pass string) bool {
	found := false
	min--
	max--

	if pass[min] == b {
		found = !found
	}
	if pass[max] == b {
		found = !found
	}

	return found
}
