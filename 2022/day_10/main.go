package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	ans1, ans2 := Run(rawInput)
	fmt.Println(ans1)
	fmt.Println(ans2)
}

type opcode int

const (
	noop opcode = iota
	addx opcode = iota
)

type command struct {
	opcode opcode
	data   int
}

func parsePart1(inp []string) []command {
	commands := make([]command, 0)

	for _, ss := range inp {
		split := strings.Split(ss, " ")

		if split[0] == "noop" {
			commands = append(commands, command{
				opcode: noop,
			})
		} else if split[0] == "addx" {
			data, _ := strconv.Atoi(split[1])
			commands = append(commands, command{
				opcode: addx,
				data:   data,
			})
		}
	}

	return commands
}

func parsePart2(inp []string) []command {
	return parsePart1(inp)
}

func Run(inp []string) (int, string) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp []command) int {
	checkpoints := common.CountedSet[int]{}
	for i := 20; i < 221; i += 40 {
		checkpoints.Add(i)
	}
	verifier := checker{checkpoints: checkpoints}

	x := 1
	cyclesCompleted := 0

	for _, cmd := range inp {
		if cmd.opcode == noop {
			cyclesCompleted += 1

			verifier.Process(cyclesCompleted, x)
			continue
		}

		if cmd.opcode == addx {
			cyclesCompleted += 1
			verifier.Process(cyclesCompleted, x)

			cyclesCompleted += 1
			verifier.Process(cyclesCompleted, x)

			x += cmd.data
		}

	}

	return verifier.data
}

type checker struct {
	data int

	str         string
	checkpoints common.CountedSet[int]
}

func (c *checker) Process(cyclesCompleted, x int) {
	if c.checkpoints.Has(cyclesCompleted) {
		c.data += (cyclesCompleted * x)
	}

	spriteMidPoint := x
	pixelDrawn := (cyclesCompleted - 1) % 40

	diff := pixelDrawn - spriteMidPoint

	if diff == 0 || diff == -1 || diff == 1 {
		c.str += "#"
	} else {
		c.str += "."
	}

}

func Part2(inp []command) string {
	checkpoints := common.CountedSet[int]{}
	for i := 20; i < 221; i += 40 {
		checkpoints.Add(i)
	}
	verifier := checker{checkpoints: checkpoints}

	x := 1
	cyclesCompleted := 0

	for _, cmd := range inp {
		if cmd.opcode == noop {
			cyclesCompleted += 1

			verifier.Process(cyclesCompleted, x)
			continue
		}

		if cmd.opcode == addx {
			cyclesCompleted += 1
			verifier.Process(cyclesCompleted, x)

			cyclesCompleted += 1
			verifier.Process(cyclesCompleted, x)

			x += cmd.data
		}

	}

	drawn := ""
	for i := 0; i < 6; i++ {
		drawn += (verifier.str[(i * 40):((i + 1) * 40)]) + "\n"
	}

	return drawn
}
