package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rawInput := takeInput()

	fmt.Println(Run(rawInput))
}

func takeInput() []string {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return inp
}

type command struct {
	dir string
	mag int
}

func parsePart1(inp []string) []command {
	arr := make([]command, len(inp))

	for i := 0; i < len(inp); i++ {
		fmt.Sscanf(inp[i], "%s %d", &arr[i].dir, &arr[i].mag)
	}

	return arr
}

func parsePart2(inp []string) []command {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

type position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *position) move(cmd command) {
	switch cmd.dir {
	case "forward":
		p.horizontal += cmd.mag
	case "down":
		p.depth += cmd.mag
	case "up":
		p.depth -= cmd.mag
	}
}

func Part1(inp []command) int {
	initial := position{0, 0, 0}

	for _, cmd := range inp {
		initial.move(cmd)
	}

	ans := initial.depth * initial.horizontal

	return ans
}

func (p *position) move2(cmd command) {
	switch cmd.dir {
	case "forward":
		{
			p.horizontal += cmd.mag
			p.depth += (p.aim * cmd.mag)
		}
	case "down":
		p.aim += cmd.mag
	case "up":
		p.aim -= cmd.mag
	}
}

func Part2(inp []command) int {
	initial := position{0, 0, 0}

	for _, cmd := range inp {
		initial.move2(cmd)
	}

	ans := initial.depth * initial.horizontal

	return ans
}
