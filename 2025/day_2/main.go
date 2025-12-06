package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsString()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type idRanges struct {
	start int
	end   int
}

func parsePart1(inp string) []idRanges {
	ranges := make([]idRanges, 0, len(inp))
	var start, end int

	split := strings.Split(inp, ",")
	for _, r := range split {
		fmt.Sscanf(r, "%d-%d", &start, &end)
		ranges = append(ranges, idRanges{start, end})
	}

	return ranges
}

func parsePart2(inp string) []idRanges {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(idRanges []idRanges) int {
	sum := 0

	for _, r := range idRanges {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			n := len(str)

			if n%2 != 0 {
				continue
			}

			if str[:n/2] == str[n/2:] {
				common.Log.Debug().Str("str", str).Send()
				sum += i
			}
		}
	}

	return sum
}

func Part2(idRanges []idRanges) int {
	fmt.Println("run cat input.txt | python main.py")

	return 0
}
