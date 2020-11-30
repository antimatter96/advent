package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	north int = iota
	east  int = iota
	south int = iota
	west  int = iota
)

func main() {
	inp := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}

	day1(inp)
}

func day1(inp string) {
	ss := strings.Split(inp, ", ")

	currentDir := north
	x, y := 0, 0

	for _, command := range ss {
		dir := command[0:1]
		destSt := command[1:]
		dest, _ := strconv.Atoi(destSt)

		if dir == "L" {
			currentDir--
		} else {
			currentDir++
		}
		currentDir += 4
		currentDir = (currentDir % 4)

		fmt.Println("FROM", command, x, y, currentDir)
		switch currentDir {
		case north:
			y += dest
		case south:
			y -= dest
		case east:
			x += dest
		case west:
			x -= dest
		default:
			panic("v interface{}")
		}
		fmt.Println(" TO ", command, x, y)
	}

	fmt.Println(x+y, math.Abs(float64(x))+math.Abs(float64(y)))
}

func str(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func day1_2(inp string) {
	ss := strings.Split(inp, ", ")

	currentDir := north
	x, y := 0, 0

	set := make(map[string]bool)
	set[str(x, y)] = true

	for _, command := range ss {
		dir := command[0:1]
		destSt := command[1:]
		dest, _ := strconv.Atoi(destSt)

		if dir == "L" {
			currentDir--
		} else {
			currentDir++
		}
		currentDir += 4
		currentDir = (currentDir % 4)

		fmt.Println("FROM", command, x, y, currentDir)
		old_x, old_y := x, y
		found := false
		switch currentDir {
		case north:
			y += dest
			for j := old_y + 1; !found && j < y; j++ {
				if _, ok := set[str(x, j)]; ok {
					y = j
					found = true
				}

				set[str(x, j)] = true
			}
		case south:
			y -= dest
			for j := old_y - 1; !found && j > y; j-- {
				if _, ok := set[str(x, j)]; ok {
					y = j
					found = true
				}

				set[str(x, j)] = true
			}
		case east:
			x += dest
			for i := old_x + 1; !found && i < x; i++ {
				if _, ok := set[str(i, y)]; ok {
					x = i
					found = true
				}

				set[str(i, y)] = true
			}
		case west:
			x -= dest
			for i := old_x - 1; !found && i > x; i-- {
				if _, ok := set[str(i, y)]; ok {
					x = i
					found = true
				}

				set[str(i, y)] = true
			}
		default:
			panic("v interface{}")
		}

		if found {
			break
		}

		fmt.Println(" TO ", command, x, y)
	}

	fmt.Println(x+y, math.Abs(float64(x))+math.Abs(float64(y)))
}
