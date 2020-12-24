package main

import (
	"bufio"
	"fmt"
	"os"
)

type cmd int

// RED
const (
	REC cmd = iota
	ROTR
	ROTC
)

type command struct {
	cmd
	op1, op2 int
}

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	day1(inp)
}

const length = 8

const rows = 6
const columns = 50

func day1(inp []string) {

	board := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]bool, columns)
	}
	//fmt.Println(board)

	for _, s := range inp {
		// HINT
		// DO SUCCESSIVE SAME OPERATION with 1st variable IN 1 GO
		// ROT 2,3
		// ROT 2,4
		// => DO ROT 2, something

		cmd := extract2(s)
		//fmt.Println(cmd)

		switch cmd.cmd {
		case REC:
			mark(board, cmd.op1, cmd.op2)
		case ROTR:
			rotRow(board, cmd.op1, cmd.op2)
		case ROTC:
			rotCol(board, cmd.op1, cmd.op2)
		default:
			panic("asd")
		}
		//fmt.Println(check(board))
	}
	fmt.Println(check(board))
}

func extract2(s string) *command {
	parsedCmd := &command{}

	if n, _ := fmt.Sscanf(s, "rect %dx%d", &parsedCmd.op1, &parsedCmd.op2); n == 2 {
		parsedCmd.cmd = REC
	} else if n, _ := fmt.Sscanf(s, "rotate row y=%d by %d", &parsedCmd.op1, &parsedCmd.op2); n == 2 {
		parsedCmd.cmd = ROTR
	} else if n, _ := fmt.Sscanf(s, "rotate column x=%d by %d", &parsedCmd.op1, &parsedCmd.op2); n == 2 {
		parsedCmd.cmd = ROTC
	}

	return parsedCmd
}

func mark(board [][]bool, wide, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < wide; j++ {
			board[i][j] = true
		}
	}
}
func rotRow(board [][]bool, rowNo, by int) {
	by = by % columns

	//fmt.Println(board[rowNo])
	r := len(board[rowNo]) - by
	board[rowNo] = append(board[rowNo][r:], board[rowNo][:r]...)
	//fmt.Println(board[rowNo])
}

func rotCol(board [][]bool, colNo, by int) {
	//fmt.Println("roating column", colNo, by, "\n")
	by = by % rows

	// umoptimised version
	// store fallen ones in an array
	// shift from the back

	temp := make([]bool, by)
	for i, j := rows-by, 0; i < rows; i, j = i+1, j+1 {
		temp[j] = board[i][colNo]
	}
	// for i, j := 0, 0; i < by; i, j = i+1, j+1 {
	// 	fmt.Printf("%d => %d\n", rows-by+i, i)
	// }
	//fmt.Println("\n", temp)

	// for i := 0; i < rows; i++ {
	// 	fmt.Printf("%v  ", board[i][colNo])
	// }
	// fmt.Println()

	for i := rows - 1; i-by > -1; i-- {
		//fmt.Printf("%d => %v\n", i, board[i-by][colNo])
		board[i][colNo] = board[i-by][colNo]
		//fmt.Printf("%d => %d\n", i, i-by)
	}

	for i := 0; i < by; i++ {
		board[i][colNo] = temp[i]
	}

	// for i := 0; i < rows; i++ {
	// 	fmt.Printf("%v  ", board[i][colNo])
	// }

	//fmt.Println("\n\nrotated column", colNo, by, "\n")
}

func check(board [][]bool) int {
	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] {
				fmt.Print("#")
				total++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return total
}
