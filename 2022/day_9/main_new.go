package main

import (
	"fmt"
)

func RunNew(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1New(parsedPart1), Part2New(parsedPart2)
}

func Part1New(commands []commandStruct) int {
	head := knot{x: xCor(0), y: yCor(0)}
	tail := knot{x: xCor(0), y: yCor(0)}

	mp := make(map[string]int)

	for _, command := range commands {
		for i := 0; i < command.count; i++ {

			switch command.dir {
			case direction_up:
				{
					head.y++
				}
			case direction_down:
				{
					head.y--
				}
			case direction_left:
				{
					head.x--
				}
			case direction_right:
				{
					head.x++
				}
			}

			tail = moveTailNew(head, tail)
			mp[fmt.Sprintf("%d,%d", tail.x, tail.y)]++
		}
	}

	return len(mp)
}

func moveTailNew(head, tail knot) (newTail knot) {
	newTail = knot{tail.x, tail.y}

	if isTouchingNew(head, tail) {
		return
	}

	if head.x == tail.x {
		if head.y-tail.y > 1 {
			newTail.y++
			return
		} else if tail.y-head.y > 1 {
			newTail.y--
			return
		}
	}

	if head.y == tail.y {
		if head.x-tail.x > 1 {
			newTail.x++
			return
		} else if tail.x-head.x > 1 {
			newTail.x--
			return
		}
	}

	if head.x > tail.x {
		newTail.x++

		if head.y > tail.y {
			newTail.y++
		} else {
			newTail.y--
		}

		return
	}

	if head.x < tail.x {
		newTail.x--

		if head.y > tail.y {
			newTail.y++
		} else {
			newTail.y--
		}

		return
	}

	if head.y > tail.y {
		newTail.y++

		if head.x > tail.x {
			newTail.x++
		} else {
			newTail.x--
		}

		return
	}

	if head.y < tail.y {
		newTail.y--

		if head.x > tail.x {
			newTail.x++
		} else {
			newTail.x--
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

func isTouchingNew(head, tail knot) bool {
	if head.x == tail.x && head.y == tail.y {
		return true
	}

	if head.x == tail.x && (head.y-tail.y == 1 || head.y-tail.y == -1) {
		return true
	}

	if head.y == tail.y && (head.x-tail.x == 1 || head.x-tail.x == -1) {
		return true
	}

	if (head.x-tail.x == 1 || head.x-tail.x == -1) && (head.y-tail.y == 1 || head.y-tail.y == -1) {
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

func Part2New(commands []commandStruct) int {
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
				newTail := moveTailNew(rp[i-1], rp[i])

				if newTail.x == rp[i].x && newTail.y == rp[i].y {
					break
				}

				rp[i].x = newTail.x
				rp[i].y = newTail.y
			}

			mp[fmt.Sprintf("%d,%d", rp[9].x, rp[9].y)]++
		}
	}

	return len(mp)
}
