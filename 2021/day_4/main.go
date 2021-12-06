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

type board [5][5]int

func (b *board) Check() bool {
	for i := 0; i < 5; i++ {
		any := true
		for j := 0; j < 5; j++ {
			if b[i][j] != -1 {
				any = false
				break
			}
		}

		if any {
			return true
		}

		any = true
		for j := 0; j < 5; j++ {
			if b[j][i] != -1 {
				any = false
				break
			}
		}

		if any {
			return true
		}
	}

	return false
}

func (b *board) Mark(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == n {
				b[i][j] = -1
				break
			}
		}
	}
}

func (b *board) Sum() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] != -1 {
				sum += b[i][j]
			}
		}
	}
	return sum
}

func parsePart1(inp []string) ([]int, []board) {
	drawn := make([]int, 0)

	drawn_string := strings.Split(inp[0], ",")
	for _, s := range drawn_string {
		temp, _ := strconv.Atoi(s)
		drawn = append(drawn, temp)
	}

	boards := make([]board, ((len(inp) - 1) / 6))

	for i := 2; i < len(inp); i += 6 {
		targetBoard := &boards[(i-2)/6]
		for j := 0; j < 5; j++ {
			fmt.Sscanf(inp[j+i], "%d %d %d %d %d", &targetBoard[j][0], &targetBoard[j][1], &targetBoard[j][2], &targetBoard[j][3], &targetBoard[j][4])
		}
	}

	return drawn, boards
}

func parsePart2(inp []string) ([]int, []board) {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	numbers, boards := parsePart1(inp)
	//parsedPart2 := parsePart2(inp)

	return Part1(numbers, boards), Part2(numbers, boards)
}

func Part1(numbers []int, boards []board) int {
	for _, number := range numbers {
		solved := false
		k := -1
		for i := 0; i < len(boards); i++ {
			boards[i].Mark(number)
			if boards[i].Check() {
				solved = true
				k = i
				break
			}
		}

		if solved {
			return boards[k].Sum() * number
		}
	}

	return 0
}

func Part2(numbers []int, boards []board) int {
	solved := make(map[int]bool)

	for _, number := range numbers {
		last := 0

		for i := 0; i < len(boards); i++ {
			if !solved[i] {

				boards[i].Mark(number)
				if boards[i].Check() {
					solved[i] = true
					last = i
				}

			}
		}

		if len(solved) == len(boards) {
			return boards[last].Sum() * number
		}
	}

	return 0
}

/*
for i := 0; i < 5; i++ {
				for _, b := range boards {
					fmt.Print(b.SingleLine(i), "   ")
				}
				fmt.Println()
			}
			fmt.Println('\n')
*/

func (b *board) SingleLine(i int) string {
	s := strings.Builder{}
	for k := 0; k < 5; k++ {
		if b[i][k] == -1 {
			s.WriteString("\033[1;38;5;208m-1\033[0m")
		} else {
			s.WriteString(fmt.Sprintf("%2d", b[i][k]))
		}
		s.WriteString("  ")
	}
	return s.String()
}
