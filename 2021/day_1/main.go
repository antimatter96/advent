package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	rawInput := takeInput()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp []string) []int {
	arr := make([]int, len(inp))

	for i := 0; i < len(inp); i++ {
		arr[i], _ = strconv.Atoi(inp[i])
	}

	return arr
}

func parsePart2(inp []string) []int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp []int) int {
	count := 0

	for i := 1; i < len(inp); i++ {
		if inp[i] > inp[i-1] {
			count++
		}
	}

	return count
}

// Slowest
func Part2(inp []int) int {
	count := 0

	window1 := inp[0] + inp[1] + inp[2]
	window2 := inp[1] + inp[2] + inp[3]

	for i := 1; i < len(inp)-2; i++ {
		if window2 > window1 {
			count++
		}

		// This branch seems to slow stuff down
		if i < len(inp)-3 {

			window1 -= inp[i-1]
			window1 += inp[i+2]

			window2 -= inp[i]
			window2 += inp[i+3]
		}

	}

	return count
}

func Part2_1(inp []int) int {
	count := 0

	window1 := inp[0] + inp[1]
	window2 := inp[1] + inp[1+1]

	for i := 1; i < len(inp)-2; i++ {
		window1 += inp[i+1]
		window2 += inp[i+2]

		if window2 > window1 {
			count++
		}

		window1 -= inp[i-1]
		window2 -= inp[i]
	}

	return count
}

func Part2_2(inp []int) int {
	count := 0

	for i := 1; i < len(inp)-2; i++ {
		window1 := inp[i-1] + inp[i] + inp[i+1]
		window2 := inp[i] + inp[i+1] + inp[i+2]

		if window2 > window1 {
			count++
		}
	}

	return count
}
