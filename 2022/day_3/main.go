package main

import (
	"fmt"

	"github.com/antimatter96/advent/2022/common"
)

func genMap() map[rune]int {
	score := make(map[rune]int)
	for i := 0; i < 26; i++ {
		score[rune(i+97)] = i + 1
		score[rune(i+65)] = i + 27
	}

	return score
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp []string) []string {
	return inp
}

func parsePart2(inp []string) [][]string {
	parsedInp := make([][]string, 0)

	intermediate := parsePart1(inp)

	var j int
	size := 3
	for i := 0; i < len(intermediate); i += size {
		j += size
		if j > len(inp) {
			j = len(inp)
		}

		parsedInp = append(parsedInp, intermediate[i:j])
	}

	return parsedInp
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp []string) int {
	score := genMap()
	total := 0
	for _, rucksack := range inp {
		firstHalf := rucksack[0 : len(rucksack)/2]
		secondHalf := rucksack[len(rucksack)/2:]

		firstMp := make(map[rune]int)

		// better to go through the smaller string
		for _, r := range firstHalf {
			firstMp[r]++
		}
		var common rune

		for _, r := range secondHalf {
			if _, ok := firstMp[r]; ok {
				common = r
			}
		}

		total += score[common]
	}
	return total
}

func Part2(inp [][]string) int {
	fmt.Println(inp)

	score := genMap()
	total := 0
	for _, group := range inp {
		maps := make([]map[rune]bool, 0)

		// better would be to start from the smallest

		for _, rucksack := range group {

			mp := make(map[rune]bool)
			for _, r := range rucksack {
				mp[r] = true
			}

			maps = append(maps, mp)
		}

		finalMap := make(map[rune]bool)

		// better would be to pick the smallest map
		for r, _ := range maps[0] {
			finalMap[r] = true
		}
		for _, mp := range maps {
			for r, _ := range maps[0] {
				finalMap[r] = finalMap[r] && mp[r]
			}
		}

		var common rune

		for r, ok := range finalMap {
			if ok {
				common = r
				break
			}
		}

		total += score[common]
	}
	return total
}
