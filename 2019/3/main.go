package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func up(board [][]int, currentX, currentY, steps, mark int) {
	for i := 0; i < steps; i++ {
		board[currentX][currentY+1+i] = mark
	}
}

func down(board [][]int, currentX, currentY, steps, mark int) {
	for i := 0; i < steps; i++ {
		board[currentX][currentY-1-i] = mark
	}
}

func left(board [][]int, currentX, currentY, steps, mark int) {
	for i := 0; i < steps; i++ {
		board[currentX-1-i][currentY] = mark
	}
}

func right(board [][]int, currentX, currentY, steps, mark int) {
	for i := 0; i < steps; i++ {
		board[currentX+1+i][currentY] = mark
	}
}

func upCheck(board [][]int, currentX, currentY, steps, mark int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		if board[currentX][currentY+1+i] == mark {
			touched = append(touched, []int{currentX, currentY + 1 + i})
		}
	}
	return touched
}

func downCheck(board [][]int, currentX, currentY, steps, mark int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		if board[currentX][currentY-1-i] == mark {
			touched = append(touched, []int{currentX, currentY - 1 - i})
		}
	}
	return touched
}

func leftCheck(board [][]int, currentX, currentY, steps, mark int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		if board[currentX-1-i][currentY] == mark {
			touched = append(touched, []int{currentX - 1 - i, currentY})
		}
	}
	return touched
}

func rightCheck(board [][]int, currentX, currentY, steps, mark int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		if board[currentX+1+i][currentY] == mark {
			touched = append(touched, []int{currentX + 1 + i, currentY})
		}
	}
	return touched
}

func upCheck1(currentX, currentY, steps int, points [][]int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		for _, p := range points {
			if currentX == p[0] && currentY+1+i == p[1] {
				touched = append(touched, []int{p[0], p[1], i + 1})
			}
		}
	}
	return touched
}

func downCheck1(currentX, currentY, steps int, points [][]int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		for _, p := range points {
			if currentX == p[0] && currentY-1-i == p[1] {
				touched = append(touched, []int{p[0], p[1], i + 1})
			}
		}
	}
	return touched
}

func leftCheck1(currentX, currentY, steps int, points [][]int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		for _, p := range points {
			if currentX-1-i == p[0] && currentY == p[1] {
				touched = append(touched, []int{p[0], p[1], i + 1})
			}
		}

	}
	return touched
}

func rightCheck1(currentX, currentY, steps int, points [][]int) [][]int {
	var touched [][]int
	for i := 0; i < steps; i++ {
		for _, p := range points {
			if currentX+1+i == p[0] && currentY == p[1] {
				touched = append(touched, []int{p[0], p[1], i + 1})
			}
		}
	}
	return touched
}

func intAfterFirst(str string) int {
	x, _ := strconv.Atoi(str[1:])
	return x
}

// func main2() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	wire1 := ""
// 	wire2 := ""
// 	for scanner.Scan() {
// 		if wire1 == "" {
// 			wire1 = scanner.Text()
// 		} else if wire2 == "" {
// 			wire2 = scanner.Text()
// 		}

// 	}

// 	if scanner.Err() != nil {
// 		// handle error.
// 	}
// 	//fmt.Println(wire1)
// 	//fmt.Println(wire2)

// 	wire1Commands := strings.Split(wire1, ",")
// 	wire2Commands := strings.Split(wire2, ",")

// 	maxX := 0
// 	maxY := 0
// 	minX := 0
// 	minY := 0

// 	for _, commandList := range [][]string{wire1Commands, wire2Commands} {
// 		//fmt.Println(commandList)

// 		currentX := 0
// 		currentY := 0
// 		for _, command := range commandList {
// 			steps := intAfterFirst(command)
// 			switch command[0] {
// 			case 'U':
// 				currentY += steps
// 			case 'D':
// 				currentY -= steps
// 			case 'L':
// 				currentX -= steps
// 			case 'R':
// 				currentX += steps
// 			}

// 			if currentX > maxX {
// 				maxX = currentX
// 			} else if currentX < minX {
// 				minX = currentX
// 			}

// 			if currentY > maxY {
// 				maxY = currentY
// 			} else if currentY < minY {
// 				minY = currentY
// 			}
// 		}
// 	}

// 	//fmt.Println(maxX, minX, maxY, minY)

// 	var board [][]int

// 	board = make([][]int, maxX-minX+1)
// 	for i := 0; i < (maxX - minX + 1); i++ {
// 		board[i] = make([]int, maxY-minY+1)
// 	}

// 	originX := 0 - minX
// 	originY := 0 - minY

// 	for _, commandList := range [][]string{wire1Commands} {
// 		currentX := originX
// 		currentY := originY

// 		for _, command := range commandList {
// 			// fmt.Println(
// 			// 	fmt.Sprintf("%d of %d", i, len(commandList)),
// 			// 	fmt.Sprintf("x: %6d, y: %6d", currentX, currentY),
// 			// 	command)
// 			//fmt.Println(i, len(wire1Commands), currentX, currentY, command)
// 			steps := intAfterFirst(command)
// 			switch command[0] {
// 			case 'U':
// 				{
// 					up(board, currentX, currentY, steps, 1)
// 					currentY += steps
// 				}
// 			case 'D':
// 				{
// 					down(board, currentX, currentY, steps, 1)
// 					currentY -= steps
// 				}
// 			case 'L':
// 				{
// 					left(board, currentX, currentY, steps, 1)
// 					currentX -= steps
// 				}
// 			case 'R':
// 				{
// 					right(board, currentX, currentY, steps, 1)
// 					currentX += steps
// 				}
// 			}
// 		}

// 	}

// 	var points [][]int
// 	for _, commandList := range [][]string{wire2Commands} {
// 		currentX := originX
// 		currentY := originY

// 		for _, command := range commandList {
// 			// fmt.Println(
// 			// 	fmt.Sprintf("%d of %d", i, len(commandList)),
// 			// 	fmt.Sprintf("x: %6d, y: %6d", currentX, currentY),
// 			// 	command)
// 			//fmt.Println(i, len(wire1Commands), currentX, currentY, command)
// 			steps := intAfterFirst(command)
// 			switch command[0] {
// 			case 'U':
// 				{
// 					points = append(points, upCheck(board, currentX, currentY, steps, 1)...)
// 					currentY += steps
// 				}
// 			case 'D':
// 				{
// 					points = append(points, downCheck(board, currentX, currentY, steps, 1)...)
// 					currentY -= steps
// 				}
// 			case 'L':
// 				{
// 					points = append(points, leftCheck(board, currentX, currentY, steps, 1)...)
// 					currentX -= steps
// 				}
// 			case 'R':
// 				{
// 					points = append(points, rightCheck(board, currentX, currentY, steps, 1)...)
// 					currentX += steps
// 				}
// 			}
// 		}

// 	}

// 	//fmt.Println(points)
// 	minD := math.MaxFloat64

// 	for _, p := range points {
// 		d := math.Abs(float64(p[0]-originX)) + math.Abs(float64(p[1]-originY))
// 		if minD > d {
// 			minD = d
// 		}
// 	}

// 	fmt.Println(minD)
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wire1 := ""
	wire2 := ""
	for scanner.Scan() {
		if wire1 == "" {
			wire1 = scanner.Text()
		} else if wire2 == "" {
			wire2 = scanner.Text()
		}

	}

	if scanner.Err() != nil {
		// handle error.
	}
	//fmt.Println(wire1)
	//fmt.Println(wire2)

	wire1Commands := strings.Split(wire1, ",")
	wire2Commands := strings.Split(wire2, ",")

	maxX := 0
	maxY := 0
	minX := 0
	minY := 0

	for _, commandList := range [][]string{wire1Commands, wire2Commands} {

		currentX := 0
		currentY := 0
		for _, command := range commandList {
			steps := intAfterFirst(command)
			switch command[0] {
			case 'U':
				currentY += steps
			case 'D':
				currentY -= steps
			case 'L':
				currentX -= steps
			case 'R':
				currentX += steps
			}

			if currentX > maxX {
				maxX = currentX
			} else if currentX < minX {
				minX = currentX
			}

			if currentY > maxY {
				maxY = currentY
			} else if currentY < minY {
				minY = currentY
			}
		}
	}

	//fmt.Println(maxX, minX, maxY, minY)

	var board [][]int

	board = make([][]int, maxX-minX+1)
	for i := 0; i < (maxX - minX + 1); i++ {
		board[i] = make([]int, maxY-minY+1)
	}

	originX := 0 - minX
	originY := 0 - minY

	for _, commandList := range [][]string{wire1Commands} {
		currentX := originX
		currentY := originY

		for _, command := range commandList {
			// fmt.Println(
			// 	fmt.Sprintf("%d of %d", i, len(commandList)),
			// 	fmt.Sprintf("x: %6d, y: %6d", currentX, currentY),
			// 	command)
			//fmt.Println(i, len(wire1Commands), currentX, currentY, command)
			steps := intAfterFirst(command)
			switch command[0] {
			case 'U':
				{
					up(board, currentX, currentY, steps, 1)
					currentY += steps
				}
			case 'D':
				{
					down(board, currentX, currentY, steps, 1)
					currentY -= steps
				}
			case 'L':
				{
					left(board, currentX, currentY, steps, 1)
					currentX -= steps
				}
			case 'R':
				{
					right(board, currentX, currentY, steps, 1)
					currentX += steps
				}
			}
		}

	}

	var points [][]int
	for _, commandList := range [][]string{wire2Commands} {
		currentX := originX
		currentY := originY

		rr := 0
		for _, command := range commandList {
			// fmt.Println(
			// 	fmt.Sprintf("%d of %d", i, len(commandList)),
			// 	fmt.Sprintf("x: %6d, y: %6d", currentX, currentY),
			// 	command)
			//fmt.Println(i, len(wire1Commands), currentX, currentY, command)
			z := len(points)
			steps := intAfterFirst(command)
			switch command[0] {
			case 'U':
				{
					points = append(points, upCheck(board, currentX, currentY, steps, 1)...)
					currentY += steps
				}
			case 'D':
				{
					points = append(points, downCheck(board, currentX, currentY, steps, 1)...)
					currentY -= steps
				}
			case 'L':
				{
					points = append(points, leftCheck(board, currentX, currentY, steps, 1)...)
					currentX -= steps
				}
			case 'R':
				{
					points = append(points, rightCheck(board, currentX, currentY, steps, 1)...)
					currentX += steps
				}
			}
			rr += steps
			if len(points) > z {
				fmt.Println(points[z:], rr)
			}
		}

	}

	turns := make([]map[string]int, 2)

	for i := 0; i < len(turns); i++ {
		turns[i] = make(map[string]int)
	}

	fmt.Println(points)
	var temp [][]int
	for wire, commandList := range [][]string{wire1Commands, wire2Commands} {
		currentX := originX
		currentY := originY

		totsTillNow := 0
		for cn, command := range commandList {
			steps := intAfterFirst(command)
			switch command[0] {
			case 'U':
				{
					temp = upCheck1(currentX, currentY, steps, points)
					currentY += steps
				}
			case 'D':
				{
					temp = downCheck1(currentX, currentY, steps, points)
					currentY -= steps
				}
			case 'L':
				{
					temp = leftCheck1(currentX, currentY, steps, points)
					currentX -= steps
				}
			case 'R':
				{
					temp = rightCheck1(currentX, currentY, steps, points)
					currentX += steps
				}
			}
			totsTillNow += steps
			for _, p := range temp {
				ps := fmt.Sprintf("%d,%d", p[0], p[1])
				fmt.Println(wire, cn, ps, p[2], totsTillNow)
				earl, pre := turns[wire][ps]
				if !pre {
					turns[wire][ps] = totsTillNow + p[2]
				} else {
					fmt.Println(">>>", wire, cn, ps, p[2], totsTillNow, earl)
				}
			}
		}

	}

	minSteps := math.MaxInt32
	for _, p := range points {
		ps := fmt.Sprintf("%d,%d", p[0], p[1])
		x, pre := turns[0][ps]
		if !pre {
			fmt.Println("nono 1")
		}
		y, pre := turns[1][ps]
		if !pre {
			fmt.Println("nono 2")
		}
		fmt.Println(p, x, y)
		if minSteps > x+y {
			minSteps = x + y
			fmt.Println(">>", p, x, y)
		}
	}

	fmt.Println(minSteps)

}
