package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	common "github.com/antimatter96/advent/2021/common"
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

func parsePart1(inp []string) []string {
	return inp
}

func parsePart2(inp []string) []string {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	numbers := parsePart1(inp)

	return Part1(numbers), Part2(numbers)
}

var errorPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func Part1(grid []string) int {
	sum := 0

	for _, line := range grid {
		stk := &common.RuneStack{}

		for _, r := range line {
			if r == ')' || r == '>' || r == ']' || r == '}' {
				k := stk.Pop()

				if areMatching(k, r) {
					continue
				} else {
					sum += errorPoints[r]
					break
				}
			} else {
				stk.Push(r)
			}
		}
	}

	return sum
}

var addPoints = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func Part2(grid []string) int {
	var scores []int

	for _, line := range grid {
		stk := &common.RuneStack{}
		points := 0
		allGood := true

		for _, r := range line {
			if r == ')' || r == '>' || r == ']' || r == '}' {
				k := stk.Pop()
				if areMatching(k, r) {
					continue
				} else {
					allGood = false
					break
				}
			} else {
				stk.Push(r)
			}
		}

		if allGood {
			for !stk.Empty() {
				k := stk.Pop()
				points *= 5
				points += addPoints[k]
			}

			scores = append(scores, points)

			fmt.Println("Points", points)
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func areMatching(open, close rune) bool {
	if (open == '(' && close == ')') ||
		(open == '{' && close == '}') ||
		(open == '[' && close == ']') ||
		(open == '<' && close == '>') {
		return true
	}

	return false
}
