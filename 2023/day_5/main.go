package main

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type almanac struct {
	seedRanges   [][]int
	instructions [][]*instruction
}

type instruction struct {
	destStart   int
	sourceStart int
	rangeLength int
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) *almanac {
	alamanac := &almanac{}

	for i := 0; i < len(inp); i++ {
		inp[i] = strings.TrimSpace(inp[i])
	}

	seedStrings := strings.Split(strings.Split(inp[0], ":")[1], " ")
	for _, s := range seedStrings {
		if s == "" || s == " " {
			continue
		}
		seed, _ := strconv.Atoi(s)

		alamanac.seedRanges = append(alamanac.seedRanges, []int{seed, 0})

	}

	var currentArray []*instruction
	for i := 1; i < len(inp); i++ {
		if inp[i] == "" || inp[i] == " " {
			continue
		}
		if len(strings.Split(inp[i], "map")) > 1 {
			if len(currentArray) > 0 {
				alamanac.instructions = append(alamanac.instructions, currentArray)

			}
			currentArray = make([]*instruction, 0)
		} else {

			instructionStrings := strings.Split(inp[i], " ")
			instructions := make([]int, 0)

			for _, s := range instructionStrings {
				instruction, _ := strconv.Atoi(s)
				instructions = append(instructions, instruction)
			}

			instruction := &instruction{
				sourceStart: instructions[1],
				destStart:   instructions[0],
				rangeLength: instructions[2],
			}

			currentArray = append(currentArray, instruction)

		}

	}
	alamanac.instructions = append(alamanac.instructions, currentArray)

	return alamanac
}

func parsePart2(inp []string) *almanac {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	// parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart1)
}

func Part1(almanac *almanac) int {
	var minLocation int = math.MaxInt

	for i := 0; i < len(almanac.instructions); i++ {

		sort.Slice(almanac.instructions[i], func(x, y int) bool {
			return almanac.instructions[i][x].sourceStart < almanac.instructions[i][y].sourceStart
		})
	}

	for seedRangeInt, seedRange := range almanac.seedRanges {

		common.Log.Debug().Int("seedRangeInt", seedRangeInt).Int("length", seedRange[1]).Send()

		for seed := seedRange[0]; seed < seedRange[0]+seedRange[1]+1; seed++ {

			start := seed

			for _, instructionSet := range almanac.instructions {

				for _, instruction := range instructionSet {
					if start >= instruction.sourceStart && start < (instruction.sourceStart+instruction.rangeLength) {

						start = instruction.destStart + (start - instruction.sourceStart)
						break
					}
				}

				//common.Log.Debug().Int("seed", seed).Int("step", start).Send()

				// New Mapping

			}

			if start < minLocation {
				minLocation = start
			}
		}

	}

	return minLocation
}

func Part2(almanac *almanac) int {

	newSeedRanges := make([][]int, 0)

	for i := 0; i < len(almanac.seedRanges); i += 2 {
		newSeedRanges = append(newSeedRanges, []int{almanac.seedRanges[i][0], almanac.seedRanges[i+1][0]})
	}

	almanac.seedRanges = newSeedRanges

	return Part1(almanac)
}

// func DetermineLast() int {

// }
