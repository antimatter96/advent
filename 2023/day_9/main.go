package main

import (
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

func parsePart1(inp []string) [][]int {

	for i := 0; i < len(inp); i++ {
		inp[i] = strings.TrimSpace(inp[i])
	}

	histories := make([][]int, 0)

	for _, historyString := range inp {
		historyStrings := strings.Split(historyString, " ")
		history := make([]int, 0)

		for _, s := range historyStrings {
			if s == "" || s == " " {
				continue
			}
			value, _ := strconv.Atoi(s)
			history = append(history, value)
		}

		histories = append(histories, history)
	}

	return histories
}

func parsePart2(inp []string) [][]int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	// parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart1)
}

func Part1(histories [][]int) int {
	sum := 0

	for _, history := range histories {

		selfHistories := makeSelfHistories(history)
		diff := 0

		for i, diffs := range selfHistories {
			common.Log.Debug().Int("Level", i+1).Ints("", diffs).Send()
		}

		for i := len(selfHistories) - 2; i > -1; i-- {
			n := len(selfHistories[i])

			diff = selfHistories[i][n-1] + diff

			common.Log.Debug().Ints("", selfHistories[i]).Int(",", diff).Send()
		}

		sum += diff

	}

	return sum
}

func Part2(histories [][]int) int {

	sum := 0

	for _, history := range histories {

		selfHistories := makeSelfHistories(history)
		diff := 0

		for i, diffs := range selfHistories {
			common.Log.Debug().Int("Level", i+1).Ints("", diffs).Send()
		}

		for i := len(selfHistories) - 2; i > -1; i-- {
			diff = selfHistories[i][0] - diff

			common.Log.Debug().Ints(",", selfHistories[i]).Int("", diff).Send()
		}

		sum += diff
	}

	return sum
}

func makeSelfHistories(history []int) [][]int {
	selfHistories := make([][]int, 0)

	selfHistories = append(selfHistories, history)

	currentSelfHistory := history

	for {
		diffArray := make([]int, len(currentSelfHistory)-1)

		for i := 1; i < len(currentSelfHistory); i++ {
			diffArray[i-1] = currentSelfHistory[i] - currentSelfHistory[i-1]
		}

		allZeros := true

		selfHistories = append(selfHistories, diffArray)

		for _, diff := range diffArray {
			if diff != 0 {
				allZeros = false
				break
			}
		}

		if allZeros {
			break
		}

		currentSelfHistory = diffArray
	}

	return selfHistories
}
