package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type equation struct {
	ans     int
	numbers []int
}

func parsePart1(inp []string) []equation {
	equations := make([]equation, 0, len(inp))

	for _, equation_s := range inp {
		nums_s := strings.Split(strings.Replace(strings.Replace(equation_s, ":", " ", -1), "  ", " ", -1), " ")

		numbers := make([]int, 0, len(nums_s)-1)
		for _, number_s := range nums_s[1:] {
			x, _ := strconv.Atoi(number_s)
			numbers = append(numbers, x)
		}

		ans, _ := strconv.Atoi(nums_s[0])

		equations = append(equations, equation{ans: ans, numbers: numbers})

	}

	return equations
}

func parsePart2(inp []string) []equation {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(equations []equation) int {
	var sum int

	for _, equation := range equations {
		slices.Reverse(equation.numbers)
		mainStack := common.NewStackFromSlice(equation.numbers)

		all := generateAll(equation.ans, mainStack)

		for _, stack := range all {
			if stack.Top() == equation.ans {
				sum += equation.ans
				break
			}
		}
	}

	return sum
}

func Part2(equations []equation) int {
	var sum int

	for _, equation := range equations {
		slices.Reverse(equation.numbers)
		mainStack := common.NewStackFromSlice(equation.numbers)

		all := generateAllWithConcate(equation.ans, mainStack)

		for _, stack := range all {
			if stack.Top() == equation.ans {
				sum += equation.ans
				break
			}
		}
	}

	return sum
}

func generateAll(ans int, stack common.Stack[int]) []common.Stack[int] {
	all := make([]common.Stack[int], 0)

	if stack.Top() > ans {
		return all
	}

	if stack.Length() == 1 {
		all = append(all, stack)
		return all
	}

	var x, y int
	multiplyArray := common.CopyStack(stack)
	x, y = multiplyArray.Pop(), multiplyArray.Pop()
	multiplyArray.Push(x * y)

	all = append(all, generateAll(ans, multiplyArray)...)

	additionArray := common.CopyStack(stack)
	x, y = additionArray.Pop(), additionArray.Pop()
	additionArray.Push(x + y)

	all = append(all, generateAll(ans, additionArray)...)

	return all
}

func generateAllWithConcate(ans int, stack common.Stack[int]) []common.Stack[int] {
	all := make([]common.Stack[int], 0)

	if stack.Top() > ans {
		return all
	}

	if stack.Length() == 1 {
		all = append(all, stack)
		return all
	}

	var x, y int
	multiplyArray := common.CopyStack(stack)
	x, y = multiplyArray.Pop(), multiplyArray.Pop()
	multiplyArray.Push(x * y)

	all = append(all, generateAllWithConcate(ans, multiplyArray)...)

	additionArray := common.CopyStack(stack)
	x, y = additionArray.Pop(), additionArray.Pop()
	additionArray.Push(x + y)

	all = append(all, generateAllWithConcate(ans, additionArray)...)

	concateArray := common.CopyStack(stack)
	x, y = concateArray.Pop(), concateArray.Pop()
	concateArray.Push(concate(x, y))

	all = append(all, generateAllWithConcate(ans, concateArray)...)

	return all
}

func concate(x, y int) int {
	var z int
	fmt.Sscanf(fmt.Sprintf("%d%d", x, y), "%d", &z)

	return z
}
