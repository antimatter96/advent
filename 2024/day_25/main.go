package main

import (
	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type machine struct {
	lock    bool
	heights []int
}

const N int = 5

func parsePart1(inp []string) []*machine {
	machines := make([]*machine, 0, len(inp)/6)

	collected := []*string{}
	for i := 0; i < len(inp); i++ {
		if inp[i] == "" {
			machine := &machine{}

			machine.heights = []int{-1, -1, -1, -1, -1}
			ps := &machine.heights

			if *collected[0] == "....." {
				// KEY
				for j := 0; j < len(collected); j++ {
					for l := 0; l < N; l++ {
						if (*ps)[l] == -1 {
							if (*collected[j])[l] == '#' {
								(*ps)[l] = (len(collected) - j - 1)
							}
						}
					}
				}
			} else {
				machine.lock = true
				// LOCK
				for j := 0; j < len(collected); j++ {
					for l := 0; l < N; l++ {
						if (*collected[j])[l] == '#' {
							(*ps)[l]++
						}
					}
				}
			}

			machines = append(machines, machine)

			collected = collected[:0]
			continue
		}

		collected = append(collected, &inp[i])
	}

	return machines
}

func parsePart2(inp []string) []*machine {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(machines []*machine) int {
	var sum int

	locks := make([]*machine, 0)
	keys := make([]*machine, 0)

	for idx := range machines {
		if machines[idx].lock {
			locks = append(locks, machines[idx])
		} else {
			keys = append(keys, machines[idx])
		}
	}

	for _, lock := range locks {
	newKey:
		for _, key := range keys {

			for column := range N {
				if lock.heights[column]+key.heights[column] > N {
					continue newKey
				}
			}

			sum++
		}
	}

	return sum
}

func Part2(machines []*machine) int {
	return 0
	return Part1(machines)
}
