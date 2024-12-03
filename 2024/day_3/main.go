package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2024/common"

	"regexp"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) string {
	common.Log.Debug().Int("Input Lengtg", len(inp)).Send()

	return strings.Join(inp, "")
}

func parsePart2(inp []string) string {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(nums string) int {
	var sum int

	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := regex.FindAllString(nums, -1)

	var x, y int

	for _, expression := range matches {
		common.Log.Debug().Str("exp", expression)
		fmt.Sscanf(expression, "mul(%d,%d)", &x, &y)
		sum += x * y
	}

	return sum
}

func Part2(nums string) int {

	var sum int

	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := regex.FindAllStringIndex(nums, -1)

	regexSwitches := regexp.MustCompile(`do(n't)?\(\)`)
	matchesSwitches := regexSwitches.FindAllStringIndex(nums, -1)
	matchesSwitches = append(matchesSwitches, []int{len(nums) + 1, 0}) // This acts a gaurd

	switchIndex := 0
	on := true

	var x, y int

	for _, expressionEndPoints := range matches {
		currentStart := startingIndex(expressionEndPoints)

		// Check if there was any preceeding switches
		for startingIndex(matchesSwitches[switchIndex]) < currentStart {
			if matchesSwitches[switchIndex][1]-matchesSwitches[switchIndex][0] == 7 {
				on = false
			} else {
				on = true
			}
			switchIndex++
		}

		if on {
			expression := expressionEndPoints
			common.Log.Debug().Str("exp", nums[expression[0]:expression[1]])
			fmt.Sscanf(nums[expression[0]:expression[1]], "mul(%d,%d)", &x, &y)
			sum += x * y
		}

	}

	return sum
}

func startingIndex(match []int) int {
	return match[0]
}
