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
		panic(scanner.Err().Error())
	}

	day2(inp)
}

func day1(inp []string) {
	var max = 0

	for _, s := range inp {
		x, y := getSeat(s)

		seatID := getSeatID(x, y)

		if seatID > max {
			max = seatID
		}
		fmt.Println(x, y)
	}

	fmt.Println(max)
}

func day2(inp []string) {
	//var ints []int
	//var temp int

	var max = 0
	var min = 100000

	mp := make(map[int]bool)
	for _, s := range inp {
		x, y := getSeat(s)
		fmt.Println(x, y)
		seatID := getSeatID(x, y)

		if seatID > max {
			max = seatID
		}
		if seatID < min {
			min = seatID
		}
		mp[seatID] = true
	}

	// for i := min; i < max; i++ {
	// 	if _, present := mp[i]; !present {
	// 		fmt.Println(i/8, i%8)

	// 		fmt.Println(getSeatID(i/8, i%8))
	// 	}
	// }

	fmt.Println(max, min)
}

const bitsForRow int = 7
const bitsForSeat int = 3

func getSeat(s string) (int, int) {

	lowRow, highRow := 0, (1<<(bitsForRow))-1

	for i := 0; i < bitsForRow; i++ {
		//fmt.Println(lowRow, highRow)
		if s[i] == 'F' {
			highRow = (lowRow + highRow) / 2
		} else if s[i] == 'B' {
			lowRow = 1 + ((lowRow + highRow) / 2)
		}
	}

	//fmt.Println(lowRow, highRow)

	lowCol, highCol := 0, (1<<(bitsForSeat))-1
	for i := 0; i < bitsForSeat; i++ {
		//fmt.Println(lowCol, highCol)
		if s[i+7] == 'L' {
			highCol = (lowCol + highCol) / 2
		} else if s[i+7] == 'R' {
			lowCol = 1 + ((lowCol + highCol) / 2)
		}
	}

	//fmt.Println(lowCol, highCol)
	return lowRow, lowCol
}

const multiplier = (1 << bitsForSeat)

func getSeatID(row, col int) int {
	return (row * multiplier) + col
}
