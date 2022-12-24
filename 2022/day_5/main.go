package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsString()

	fmt.Println(Run(rawInput))
}

type input struct {
	stacks []common.Stack[string]
	moves  []move
}

type move struct {
	from  int
	to    int
	count int
}

func parsePart1(inp string) input {
	stacks := make([]common.Stack[string], 0)
	groups := strings.Split(inp, "\n\n")

	currentStatusRaw := strings.Split(groups[0], "\n")

	labels := strings.Split(strings.TrimSpace(currentStatusRaw[len(currentStatusRaw)-1]), "")

	totalStacks, _ := strconv.Atoi(labels[len(labels)-1])
	for i := 0; i < totalStacks; i++ {
		stacks = append(stacks, common.Stack[string]{})
	}

	for _, s := range currentStatusRaw {
		row := strings.Split(s, "")

		for i := 1; i < len(row); i += 4 {
			if row[i] != " " {
				stacks[(i+1)/4].Push(row[i])
			}
		}
	}

	for _, stack := range stacks {
		stack.Reverse()
	}

	movesRaw := strings.Split(strings.TrimSpace(groups[1]), "\n")

	moves := make([]move, 0)

	var from, to, count int
	for _, s := range movesRaw {
		fmt.Sscanf(s, "move %d from %d to %d", &count, &from, &to)
		moves = append(moves, move{from: from, count: count, to: to})
	}

	return input{stacks, moves}
}

func parsePart2(inp string) input {
	return parsePart1(inp)
}

func Run(inp string) (string, string) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp input) string {
	top := ""

	var from, to int
	for _, move := range inp.moves {

		from = move.from - 1
		to = move.to - 1
		for i := 0; i < move.count; i++ {
			inp.stacks[to].Push(inp.stacks[from].Pop())
		}
	}

	for _, stack := range inp.stacks {
		top += stack.Top()
	}

	return top
}

func Part2(inp input) string {
	top := ""

	tempStack := common.Stack[string]{}

	var from, to int
	for _, move := range inp.moves {
		from = move.from - 1
		to = move.to - 1

		for i := 0; i < move.count; i++ {
			tempStack.Push(inp.stacks[from].Pop())
		}
		for i := 0; i < move.count; i++ {
			inp.stacks[to].Push(tempStack.Pop())
		}
	}

	for _, stack := range inp.stacks {
		top += stack.Top()
	}

	return top
}
