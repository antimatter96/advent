package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type cubeSets map[string]int

var limits cubeSets = cubeSets{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type game struct {
	id   int
	sets []cubeSets
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(gameRecords []string) []game {
	games := make([]game, 0)
	for _, gameRecordStr := range gameRecords {
		game := game{sets: make([]cubeSets, 0)}

		fmt.Sscanf(gameRecordStr, "Game %d:", &game.id)

		setStrings := strings.Split(strings.Split(gameRecordStr, ":")[1], ";")

		for _, setString := range setStrings {
			cubeSet := make(cubeSets)
			setString = strings.TrimSpace(setString)
			common.Log.Debug().Str("set string", setString).Send()

			cubeCounts := strings.Split(setString, ",")

			for _, cubeCount := range cubeCounts {
				cubeCount = strings.TrimSpace(cubeCount)

				var color string
				var count int

				fmt.Sscanf(cubeCount, "%d %s:", &count, &color)
				common.Log.Debug().Int(color, count).Send()

				cubeSet[color] = count
			}

			game.sets = append(game.sets, cubeSet)
		}

		games = append(games, game)
	}

	return games
}

func parsePart2(inp []string) []game {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(games []game) int {
	var sum int

	for _, game := range games {
		common.Log.Debug().Int("Id", game.id).Int("games", len(game.sets)).Send()

		sum += game.id

	mainLoop:
		for _, set := range game.sets {

			for color, count := range set {
				if limits[color] < count {
					sum -= game.id
					common.Log.Debug().Int("Id", game.id).Str("color", color).Int("limit", limits[color]).Int("actual", count).Send()
					break mainLoop
				}
			}
		}

	}

	return sum
}

func Part2(games []game) int {

	var sum int

	for _, game := range games {
		common.Log.Debug().Int("Id", game.id).Int("games", len(game.sets)).Send()

		maxPossible := make(cubeSets)

		for _, set := range game.sets {

			for color, count := range set {
				if max, present := maxPossible[color]; !present || max < count {
					maxPossible[color] = count
				}
			}
		}

		powerSet := 1

		tempLog := common.Log.Debug().Int("Id", game.id)
		for color, count := range maxPossible {
			tempLog.Int(color, count)
			powerSet *= count
		}
		tempLog.Send()

		sum += powerSet

	}

	return sum

}
