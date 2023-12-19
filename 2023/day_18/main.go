package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

const (
	DIRECTION_UP    = "U"
	DIRECTION_DOWN  = "D"
	DIRECTION_LEFT  = "L"
	DIRECTION_RIGHT = "R"
)

type instruction struct {
	direction string
	length    int
}

func parsePart1(inpLines []string) []instruction {
	instructions := make([]instruction, 0, len(inpLines))

	var scannedDirection string
	for i := 0; i < len(inpLines); i++ {
		inpLines[i] = strings.TrimSpace(inpLines[i])

		if inpLines[i] == "" || inpLines[i] == " " {
			continue
		}

		instruction := instruction{}

		fmt.Sscanf(inpLines[i], "%s %d (#%s)", &scannedDirection, &instruction.length)

		switch scannedDirection {
		case "U":
			instruction.direction = DIRECTION_UP
		case "D":
			instruction.direction = DIRECTION_DOWN
		case "L":
			instruction.direction = DIRECTION_LEFT
		case "R":
			instruction.direction = DIRECTION_RIGHT
		}

		instructions = append(instructions, instruction)
	}

	for i := 0; i < len(instructions); i++ {
		common.Log.Debug().Str("d", instructions[i].direction).Int("l", instructions[i].length).Send()
	}

	return instructions
}

func parsePart2(inpLines []string) []instruction {
	instructions := make([]instruction, 0, len(inpLines))

	var scannedLength, dummy string
	for i := 0; i < len(inpLines); i++ {

		inpLines[i] = strings.TrimSpace(inpLines[i])

		if inpLines[i] == "" || inpLines[i] == " " {
			continue
		}

		instruction := instruction{}

		fmt.Sscanf(inpLines[i], "%s %s (#%s)", &dummy, &dummy, &scannedLength)

		switch scannedLength[5:6] {
		case "3":
			instruction.direction = DIRECTION_UP
		case "1":
			instruction.direction = DIRECTION_DOWN
		case "2":
			instruction.direction = DIRECTION_LEFT
		case "0":
			instruction.direction = DIRECTION_RIGHT
		}

		x, _ := strconv.ParseInt(scannedLength[0:5], 16, 0)
		instruction.length = int(x)

		instructions = append(instructions, instruction)
		common.Log.Debug().Str(".", inpLines[i]).Str(".", scannedLength[0:5]).Str("d", instructions[i].direction).Int("l", instructions[i].length).Send()
	}

	return instructions
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(instructions []instruction) int {
	x, y := 0, 0
	prevX, prevY := 0, 0

	area := 0
	permimeter := 0

	for _, insinstruction := range instructions {

		switch insinstruction.direction {
		case DIRECTION_DOWN:
			y += insinstruction.length
		case DIRECTION_UP:
			y -= insinstruction.length
		case DIRECTION_RIGHT:
			x += insinstruction.length
		case DIRECTION_LEFT:
			x -= insinstruction.length
		}

		area += (x - prevX) * (y + prevY)

		common.Log.Debug().Ints("f", []int{prevX, prevY}).Ints("t", []int{x, y}).Int("z", area).Send()

		permimeter += insinstruction.length

		prevX, prevY = x, y
	}

	area = int(math.Abs(float64(area)))

	common.Log.Debug().Int("Area", area/2).Int("Perimeter", permimeter/2).Send()

	totalArea := area/2 + permimeter/2 + 1

	return totalArea
}

func Part2(instructions []instruction) int {
	return Part1(instructions)
}
