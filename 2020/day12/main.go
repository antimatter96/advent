package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Errorf(scanner.Err().Error())
	}

	day1(inp)
}

const (
	north int = iota
	east  int = iota
	south int = iota
	west  int = iota
)

func day1(inp []string) {
	ss := inp

	shipX, shipY := 0, 0
	waypointX, waypointY := 10, 1

	for _, command := range ss {
		// fmt.Printf("")
		fmt.Printf("FROM %3s %3d %3d || %3d %3d", command, waypointX, waypointY, shipX, shipY)

		dir := command[0:1]
		destSt := command[1:]
		dest, _ := strconv.Atoi(destSt)

		switch dir {
		case "F":
			shipX += (dest * waypointX)
			shipY += (dest * waypointY)
		case "N":
			waypointY += dest
		case "S":
			waypointY -= dest
		case "E":
			waypointX += dest
		case "W":
			waypointX -= dest
		default:
			waypointX, waypointY = rotate2(waypointX, waypointY, dest, dir)
		}

		fmt.Printf("  TO %3s %3d %3d || %3d %3d\n", command, waypointX, waypointY, shipX, shipY)
	}

	fmt.Println(shipX+shipY, math.Abs(float64(shipX))+math.Abs(float64(shipY)))
}

func rotate(currentDir int, degrees int, dir string) int {
	degrees = degrees / 90

	if dir == "L" {
		currentDir -= degrees
	} else if dir == "R" {
		currentDir += degrees
	}

	currentDir += 4
	currentDir = currentDir % 4

	return currentDir
}

var temp int

func rotate2(x, y, degrees int, dir string) (int, int) {
	degrees = degrees / 90
	degrees = degrees % 4

	if degrees == 2 {
		x = -x
		y = -y

		return x, y
	}

	temp = x
	x = y
	y = temp

	if (dir == "R" && degrees == 1) || (dir == "L" && degrees == 3) {
		y = -y
	} else if (dir == "R" && degrees == 3) || (dir == "L" && degrees == 1) {
		x = -x
	}

	return x, y
}

func day2(inp []string) {
	ss := inp

	currentDir := east
	x, y := 0, 0

	for _, command := range ss {
		// fmt.Printf("FROM %4s %4d %4d DIR : %d ", command, x, y, currentDir)

		dir := command[0:1]
		destSt := command[1:]
		dest, _ := strconv.Atoi(destSt)

		if dir == "F" {
			switch currentDir {
			case north:
				y += dest
			case south:
				y -= dest
			case east:
				x += dest
			case west:
				x -= dest
			}
		} else {
			switch dir {
			case "N":
				y += dest
			case "S":
				y -= dest
			case "E":
				x += dest
			case "W":
				x -= dest
			case "L":
				currentDir = rotate(currentDir, dest, dir)
			case "R":
				currentDir = rotate(currentDir, dest, dir)
			}
		}

		// fmt.Printf("  To %4d %4d DIR : %d\n", x, y, currentDir)
	}

	fmt.Println(x+y, math.Abs(float64(x))+math.Abs(float64(y)))
}
