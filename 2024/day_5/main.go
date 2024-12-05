package main

import (
	"fmt"
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

type problemInput struct {
	rules  []rule
	orders []order

	reverseGraph common.Graph[bool] // [x][y] = true => y depends on x === x is depended on by
}

type rule struct {
	a, z int
}

type order []int

var max int = 100

func parsePart1(inp []string) problemInput {
	n := len(inp)
	i := 0

	rules := make([]rule, 0, len(inp))
	for ; i < n; i++ {
		if inp[i] == "" {
			i++
			break
		}
		rule := &rule{}
		fmt.Sscanf(inp[i], "%d|%d", &(rule.a), &(rule.z))
		rules = append(rules, *rule)
	}

	orders := make([]order, 0, n-len(rules))
	for ; i < n; i++ {
		inp_split := strings.Split(inp[i], ",")
		order := make(order, 0, len(inp_split))
		for _, inp_split_s := range inp_split {
			x, _ := strconv.Atoi(inp_split_s)
			order = append(order, x)
		}
		orders = append(orders, order)
	}

	reverseGraph := make([][]bool, max+1)
	for i := 0; i < max+1; i++ {
		reverseGraph[i] = make([]bool, max+1)
	}
	for _, rule := range rules {
		reverseGraph[rule.z][rule.a] = true
	}

	return problemInput{rules: rules, orders: orders, reverseGraph: reverseGraph}
}

func parsePart2(inp []string) problemInput {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp problemInput) int {
	var sum int

	for _, order := range inp.orders {
		indegrees := createIndegrees(inp.reverseGraph, order)

		allGood := true

		for i, page := range order {
			if len(indegrees[page]) != 0 {
				allGood = false
				break
			}

			for _, others := range order[i+1:] {
				delete(indegrees[others], page)
				if len(indegrees[others]) == 0 {
					delete(indegrees, others)
				}
			}
		}

		if allGood {
			sum += order[len(order)/2]
		}
	}

	return sum
}

func Part2(inp problemInput) int {
	var sum int

	for _, order := range inp.orders {
		indegrees := createIndegrees(inp.reverseGraph, order)

		target := len(order) / 2
		for _, page := range order {
			if len(indegrees[page]) == target {
				sum += page
				break
			}
		}
	}

	return sum - Part1(inp)
}

func createIndegrees(grpah common.Graph[bool], order []int) map[int]map[int]bool {
	indegrees := make(map[int]map[int]bool)
	currentPages := make(map[int]bool)

	for _, page := range order {
		currentPages[page] = true
		indegrees[page] = make(map[int]bool)
	}

	for _, page := range order {
		for dependencyPage, blocks := range grpah[page] {
			if blocks && currentPages[dependencyPage] {
				indegrees[page][dependencyPage] = true
			}
		}
	}

	return indegrees
}
