package main

import (
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

func main() {
	rawInput := common.TakeInputAsString()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp string) common.Graph[rune] {
	inp = strings.TrimSpace(inp)

	inpLines := strings.Split(inp, "\n")

	schematic := make(common.Graph[rune], len(inpLines))

	for i := 0; i < len(inpLines); i++ {
		schematic[i] = []rune(strings.TrimSpace(inpLines[i]))
	}

	// common.Log.Debug().Int("lines", len(inpLines)).Send()

	return schematic
}

func parsePart2(inp string) common.Graph[rune] {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	for k := range symbols {
		whenToStop[k] = struct{}{}
	}

	return Part1(parsedPart1), Part2(parsedPart2)
}

var whenToProcees map[rune]struct{} = map[rune]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
}

var symbols map[rune]struct{} = map[rune]struct{}{
	'-': {},
	'@': {},
	'*': {},
	'/': {},
	'&': {},
	'#': {},
	'%': {},
	'+': {},
	'=': {},
	'$': {},
}

var whenToStop map[rune]struct{} = map[rune]struct{}{
	'.': {},
}

func Part1(schematic common.Graph[rune]) int {

	sumGraph1 := sumGraph(schematic)
	common.Log.Debug().Int("initial", sumGraph1).Send()

	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[0]); j++ {
			if _, ok := symbols[schematic[i][j]]; ok {
				common.StartFloodFill[rune](schematic, i, j, schematic[i][j], whenToProcees, whenToStop)

				// schematic.Print()

			}
		}

	}

	sumGraph2 := sumGraph(schematic)

	common.Log.Debug().Int("final", sumGraph2).Send()

	common.Log.Debug().Int("answer", sumGraph1-sumGraph2).Send()

	return sumGraph1 - sumGraph2
}

func Part2(schematic common.Graph[rune]) int {
	return 0
}

func sumGraph(schematic common.Graph[rune]) int {
	strs := make([]string, 0)

	for i := 0; i < len(schematic); i++ {
		strB := strings.Builder{}
		for j := 0; j < len(schematic[0]); j++ {

			if _, ok := whenToStop[schematic[i][j]]; ok {
				strB.WriteRune(' ')
				continue

			}
			strB.WriteRune(schematic[i][j])

		}

		strs = append(strs, strings.TrimSpace(strB.String()))
	}

	sum := 0

	for i := 0; i < len(strs); i++ {

		if strs[i] == "" {
			continue
		}
		numbers := strings.Split(strs[i], " ")

		for _, numberStr := range numbers {
			number, _ := strconv.Atoi(numberStr)

			common.Log.Debug().Int("", number).Send()

			sum += number
		}

	}

	return sum
	return sum
}
