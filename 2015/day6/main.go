package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	day1_2(inp)
}

func day1(inp []string) {
	board := make([][]bool, 1000)
	for i := 0; i < len(board); i++ {
		board[i] = make([]bool, 1000)
	}
	//fmt.Println(inp)
	total := 0
	for _, s := range inp {
		x1, y1, x2, y2, state := parseInstructions(s)
		processInstruction(&board, x1, y1, x2, y2, state)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] {
				total++
			}
		}
	}
	fmt.Println(total)
}

func processInstruction(board *[][]bool, x1, y1, x2, y2 int, state string) {
	fmt.Println(state, x1, y1, x2, y2)
	for i := x1; i < x2+1; i++ {
		for j := y1; j < y2+1; j++ {
			if state == "on" {
				(*board)[i][j] = true
			} else if state == "off" {
				(*board)[i][j] = false
			} else {
				(*board)[i][j] = !(*board)[i][j]
			}
		}
	}
}

func parseInstructions(s string) (int, int, int, int, string) {
	spl := strings.Split(s, " ")

	x1y1Cordinate := 2
	x2y2Cordinate := 4

	var state string
	if spl[1] == "on" {
		state = "on"
	} else if spl[1] == "off" {
		state = "off"
	} else {
		state = "toggle"
		x1y1Cordinate--
		x2y2Cordinate--
	}

	var x1, y1, x2, y2 int

	x1y1 := strings.Split(spl[x1y1Cordinate], ",")
	x1, _ = strconv.Atoi(x1y1[0])
	y1, _ = strconv.Atoi(x1y1[1])

	x2y2 := strings.Split(spl[x2y2Cordinate], ",")
	x2, _ = strconv.Atoi(x2y2[0])
	y2, _ = strconv.Atoi(x2y2[1])

	return x1, y1, x2, y2, state
}

func day1_2(inp []string) {
	board := make([][]int, 1000)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, 1000)
	}
	//fmt.Println(inp)
	total := 0
	for _, s := range inp {
		x1, y1, x2, y2, state := parseInstructions(s)
		processInstruction2(&board, x1, y1, x2, y2, state)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			total += board[i][j]
		}
	}
	fmt.Println(total)
}

func processInstruction2(board *[][]int, x1, y1, x2, y2 int, state string) {
	fmt.Println(state, x1, y1, x2, y2)
	for i := x1; i < x2+1; i++ {
		for j := y1; j < y2+1; j++ {
			if state == "on" {
				(*board)[i][j]++
			} else if state == "off" {
				(*board)[i][j]--
				if (*board)[i][j] < 0 {
					(*board)[i][j] = 0
				}
			} else {
				(*board)[i][j] += 2
			}
		}
	}
}
