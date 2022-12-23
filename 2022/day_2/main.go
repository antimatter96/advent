package main

import (
	"fmt"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	fmt.Println(Run(rawInput))
}

const (
	ROCK     int = iota
	PAPER    int = iota
	SCISSORS int = iota

	WIN  int = iota
	DRAW int = iota
	LOSS int = iota
)

var oppMap = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var myMap = map[string]int{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var resultMap = map[string]int{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}

var score = map[int]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,

	WIN:  6,
	DRAW: 3,
	LOSS: 0,
}

type command struct {
	opp    int
	player int
	result int
}

func parsePart1(inp []string) []command {
	arr := make([]command, len(inp))

	var x string
	var y string
	for i := 0; i < len(inp); i++ {
		fmt.Sscanf(inp[i], "%s %s", &x, &y)

		arr[i].opp = oppMap[x]
		arr[i].player = myMap[y]
	}

	return arr
}

func parsePart2(inp []string) []command {
	arr := make([]command, len(inp))

	var x string
	var y string
	for i := 0; i < len(inp); i++ {
		fmt.Sscanf(inp[i], "%s %s", &x, &y)

		arr[i].opp = oppMap[x]
		arr[i].result = resultMap[y]
	}

	return arr
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func score1(cmd command) int {
	result := 0

	if cmd.opp == cmd.player {
		result = 3
	}

	if cmd.player == SCISSORS {
		if cmd.opp == PAPER {
			result = 6
		}
	}

	if cmd.player == PAPER {
		if cmd.opp == ROCK {
			result = 6
		}

	}

	if cmd.player == ROCK {

		if cmd.opp == SCISSORS {
			result = 6
		}
	}

	return result + score[cmd.player]
}

func Part1(inp []command) int {
	sum := 0
	for _, cmd := range inp {
		sum += score1(cmd)
	}
	return sum
}

func score2(cmd command) int {
	myScore := 0

	if cmd.result == DRAW {
		myScore = score[cmd.opp]
	}

	if cmd.result == WIN {
		switch cmd.opp {
		case ROCK:
			myScore = score[PAPER]
		case PAPER:
			myScore = score[SCISSORS]
		case SCISSORS:
			myScore = score[ROCK]
		}

	} else if cmd.result == LOSS {
		switch cmd.opp {
		case ROCK:
			myScore = score[SCISSORS]
		case PAPER:
			myScore = score[ROCK]
		case SCISSORS:
			myScore = score[PAPER]
		}
	}

	return myScore + score[cmd.result]
}

func Part2(inp []command) int {
	sum := 0
	for _, cmd := range inp {
		sum += score2(cmd)
	}
	return sum
}
