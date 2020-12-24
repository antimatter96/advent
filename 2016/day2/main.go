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

// LowLimit
var (
	LowLimit   = 0
	UpperLimit = 2
)

/*
2,0 => 1    2,1 => 2     2,2 => 3
1,0 => 4    1,1 => 5     1,2 => 6
0,0 => 7    0,1 => 8     0,2 => 9
*/

func digit(x, y int) int {
	z := 6 - (3 * x) + 1 + y
	return z
}

func day1(inp []string) {
	initialX, initialY := 1, 1
	for _, s := range inp {
		initialX, initialY = whereDoIEnd(initialX, initialY, s)
		fmt.Printf("%d", digit(initialX, initialY))
	}
	fmt.Println("")
}

func whereDoIEnd(initialX, initialY int, ins string) (int, int) {
	newX, newY := initialX, initialY
	var tempX, tempY int

	for _, r := range ins {
		tempX, tempY = newX, newY

		switch r {
		case 'U':
			tempX++
		case 'D':
			tempX--
		case 'L':
			tempY--
		case 'R':
			tempY++
		default:
			panic("err")
		}

		//fmt.Println("Initial", newX, newY, "Final", tempX, tempY)

		if tempX > UpperLimit || tempX < LowLimit ||
			tempY > UpperLimit || tempY < LowLimit {
			//fmt.Println("SKIPPING", newX, newY, tempX, tempY)
		} else {
			newX, newY = tempX, tempY
		}

	}

	//fmt.Println("FINAL", newX, newY)
	return newX, newY
}

/*
0,0    0,1    0,2    0,3    0,4
1,0    1,1    1,2    1,3    1,4
2,0    2,1    2,2    2,3    2,4
3,0    3,1    3,2    3,3    3,4
4,0    4,1    4,2    4,3    4,4
*/

var pos = map[string]string{
	"0,2": "1",
	"1,1": "2",
	"1,2": "3",
	"1,3": "4",
	"2,0": "5",
	"2,1": "6",
	"2,2": "7",
	"2,3": "8",
	"2,4": "9",
	"3,1": "A",
	"3,2": "B",
	"3,3": "C",
	"4,2": "D",
}

func whereDoIEnd2(initialX, initialY int, ins string) (int, int) {
	var (
		LowLimit2   = 0
		UpperLimit2 = 4
	)

	newX, newY := initialX, initialY
	var tempX, tempY int

	for _, r := range ins {
		tempX, tempY = newX, newY

		switch r {
		case 'U':
			tempX--
		case 'D':
			tempX++
		case 'L':
			tempY--
		case 'R':
			tempY++
		default:
			panic("err")
		}

		//fmt.Println("Initial", newX, newY, "Final", tempX, tempY)

		if tempX > UpperLimit2 || tempX < LowLimit2 ||
			tempY > UpperLimit2 || tempY < LowLimit2 {
			//fmt.Println("SKIPPING", newX, newY, tempX, tempY)
			continue
		}

		if tempX < lowLimit[tempY] || tempX > upperLimit[tempY] ||
			tempY < lowLimit[tempX] || tempY > upperLimit[tempX] {
			//fmt.Println("SKIPPING", newX, newY, tempX, tempY)
			continue
		}

		newX, newY = tempX, tempY
	}

	//fmt.Println("FINAL", newX, newY)
	return newX, newY
}

var lowLimit = map[int]int{
	0: 2,
	1: 1,
	2: 0,
	3: 1,
	4: 2,
}
var upperLimit = map[int]int{
	0: 2,
	1: 3,
	2: 4,
	3: 3,
	4: 2,
}

func day2(inp []string) {
	initialX, initialY := 2, 0
	for _, s := range inp {
		initialX, initialY = whereDoIEnd2(initialX, initialY, s)
		fmt.Printf("%s", pos[stringify(initialX, initialY)])
	}
	fmt.Println()
}

func stringify(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
