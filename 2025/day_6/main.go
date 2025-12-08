package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type problemInput struct {
	numbers    [][]int
	operations []rune
}

var op map[string]rune = map[string]rune{"+": '+', "-": '-', "*": '*', "/": '/'}

func parsePart1(inp []string) problemInput {
	n := len(inp)

	numbers := make([][]int, n-1)

	r := regexp.MustCompile(`\s+`)
	for i := 0; i < n-1; i++ {
		n_s := r.Split(strings.TrimSpace(inp[i]), -1)

		numbers[i] = make([]int, 0, len(n_s))
		for _, n_str := range n_s {
			x, _ := strconv.Atoi(n_str)
			numbers[i] = append(numbers[i], x)
		}

		common.Log.Debug().Strs("i", n_s).Ints("o", numbers[i]).Send()
	}

	ops_s := r.Split(strings.TrimSpace(inp[n-1]), -1)
	common.Log.Debug().Int("", len(ops_s)).Msg("Number of operations")
	operations := make([]rune, 0)
	for _, op_str := range ops_s {
		operations = append(operations, op[op_str])
	}

	return problemInput{numbers: numbers, operations: operations}
}

func parsePart2(inp []string) problemInput {
	n := len(inp)

	r := regexp.MustCompile(`\s+`)

	ops_s := r.Split(strings.TrimSpace(inp[n-1]), -1)
	operations := make([]rune, 0)
	for _, op_str := range ops_s {
		operations = append(operations, op[op_str])
	}
	common.Log.Debug().Int("", len(ops_s)).Msg("Number of operations")

	numbers := make([][]int, len(ops_s))
	for i := range ops_s {
		numbers[i] = make([]int, 0, n-1)
	}

	totalColumns := len(inp[0])
	currentNumberColumn := 0 // Which column are we filling currently

	for j := totalColumns - 1; j >= 0; j-- {
		allSpaces := true

		tempNumber := 0

		for i := 0; i < n-1; i++ {
			if inp[i][j] == ' ' {
				continue
			}
			allSpaces = false
			tempNumber *= 10
			tempNumber += int(inp[i][j] - '0')
		}

		if allSpaces {
			currentNumberColumn++
		} else {
			common.Log.Debug().Int("column", currentNumberColumn).Int("n", tempNumber).Msg("Number Parsed")
			numbers[currentNumberColumn] = append(numbers[currentNumberColumn], tempNumber)
		}
	}

	// We collected numbers rigt to left while operations were collected left to right
	slices.Reverse(operations)

	return problemInput{numbers: numbers, operations: operations}
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp problemInput) int {
	var sum int

	n_columns := len(inp.numbers[0])
	results := make([]int, n_columns)

	for c := 0; c < n_columns; c++ {
		results[c] = inp.numbers[0][c]

		if inp.operations[c] == '+' {
			for r := 1; r < len(inp.numbers); r++ {
				results[c] += inp.numbers[r][c]
			}
		} else if inp.operations[c] == '*' {
			for r := 1; r < len(inp.numbers); r++ {
				results[c] *= inp.numbers[r][c]
			}
		}

		sum += results[c]
	}

	return sum
}

func Part2(inp problemInput) int {
	/*
		inp.numbers[i] has all numbers of an operation
		operatio
	*/
	var sum int

	for i, op := range inp.operations {
		var result int

		if op == '+' {
			result = 0
			for _, num := range inp.numbers[i] {
				result += num
			}
		} else if op == '*' {
			result = 1
			for _, num := range inp.numbers[i] {
				result *= num
			}
		}
		sum += result
	}

	return sum
}
