package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	fmt.Println(Run(rawInput))
}

type direction int

const (
	direction_up    direction = iota
	direction_down  direction = iota
	direction_left  direction = iota
	direction_right direction = iota
)

var dirMp = map[string]direction{
	"U": direction_up,
	"D": direction_down,
	"L": direction_left,
	"R": direction_right,
}

type commandStruct struct {
	dir   direction
	count int
}

type xCor int
type yCor int

type knot struct {
	x xCor
	y yCor
}

type rope [10]knot

func parsePart1(inp []string) []commandStruct {
	commands := make([]commandStruct, 0)
	for _, row := range inp {
		split := strings.Split(row, " ")

		dir := split[0]
		count, _ := strconv.Atoi(split[1])

		commands = append(commands, commandStruct{
			dir:   dirMp[dir],
			count: count,
		})
	}

	return commands
}

func parsePart2(inp []string) []commandStruct {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(commands []commandStruct) int {
	headX, headY := xCor(0), yCor(0)
	tailX, tailY := xCor(0), yCor(0)

	mp := make(map[string]int)

	for _, command := range commands {
		for i := 0; i < command.count; i++ {

			switch command.dir {
			case direction_up:
				{
					headY++
				}
			case direction_down:
				{
					headY--
				}
			case direction_left:
				{
					headX--
				}
			case direction_right:
				{
					headX++
				}
			}

			tailX, tailY = moveTail(headX, headY, tailX, tailY)
			mp[fmt.Sprintf("%d,%d", tailX, tailY)]++
		}
	}

	return len(mp)
}

func moveTail(headX xCor, headY yCor, tailX xCor, tailY yCor) (newTailX xCor, newTailY yCor) {
	newTailX, newTailY = tailX, tailY
	if isTouching(headX, headY, tailX, tailY) {
		return
	}

	if headX == tailX {
		if headY-tailY > 1 {
			newTailY++
			return
		} else if tailY-headY > 1 {
			newTailY--
			return
		}
	}

	if headY == tailY {
		if headX-tailX > 1 {
			newTailX++
			return
		} else if tailX-headX > 1 {
			newTailX--
			return
		}
	}

	if headX > tailX {
		newTailX++

		if headY > tailY {
			newTailY++
		} else {
			newTailY--
		}

		return
	}

	if headX < tailX {
		newTailX--

		if headY > tailY {
			newTailY++
		} else {
			newTailY--
		}

		return
	}

	if headY > tailY {
		newTailY++

		if headX > tailX {
			newTailX++
		} else {
			newTailX--
		}

		return
	}

	if headY < tailY {
		newTailY--

		if headX > tailX {
			newTailX++
		} else {
			newTailX--
		}

		return
	}

	/*
		+-----------------------+
		| .    .    .    .    . |
		+-----------------------+
		| .    .    .    .    . |
		+-----------------------+
		| .    .    x    .    . |
		+-----------------------+
		| .    .    .    .    . |
		+-----------------------+
		| .    .    .    .    . |
		+-----------------------+
	*/

	return
}

func isTouching(headX xCor, headY yCor, tailX xCor, tailY yCor) bool {
	if headX == tailX && headY == tailY {
		return true
	}

	if headX == tailX && (headY-tailY == 1 || headY-tailY == -1) {
		return true
	}

	if headY == tailY && (headX-tailX == 1 || headX-tailX == -1) {
		return true
	}

	if (headX-tailX == 1 || headX-tailX == -1) && (headY-tailY == 1 || headY-tailY == -1) {
		return true
	}

	/*
		+-------------+
		| O    O    O |
		+-------------+
		| O    H    O |
		+-------------+
		| O    O    O |
		+-------------+
	*/

	return false
}

func Part2(commands []commandStruct) int {
	var rp rope
	mp := make(map[string]int)

	for _, command := range commands {
		for i := 0; i < command.count; i++ {

			switch command.dir {
			case direction_up:
				{
					rp[0].y++
				}
			case direction_down:
				{
					rp[0].y--
				}
			case direction_left:
				{
					rp[0].x--
				}
			case direction_right:
				{
					rp[0].x++
				}
			}

			for i := 1; i < 10; i++ {
				newX, newY := moveTail(rp[i-1].x, rp[i-1].y, rp[i].x, rp[i].y)

				if newX == rp[i].x && newY == rp[i].y {
					break
				}

				rp[i].x = newX
				rp[i].y = newY
			}

			mp[fmt.Sprintf("%d,%d", rp[9].x, rp[9].y)]++
		}
	}

	return len(mp)
}
