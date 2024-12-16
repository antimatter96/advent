package main

import (
	"fmt"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type machine struct {
	xA, yA int
	xB, yB int

	xPrize, yPrize int
}

func parsePart1(inp []string) []*machine {
	machines := make([]*machine, 0, len(inp)/3)
	for i := 0; i < len(inp); i += 3 {
		if inp[i] == "" {
			i -= 2
			continue
		}

		machine := &machine{}

		fmt.Sscanf(inp[i], "Button A: X+%d, Y+%d", &machine.xA, &machine.yA)
		fmt.Sscanf(inp[i+1], "Button B: X+%d, Y+%d", &machine.xB, &machine.yB)
		fmt.Sscanf(inp[i+2], "Prize: X=%d, Y=%d", &machine.xPrize, &machine.yPrize)

		machines = append(machines, machine)
	}

	return machines
}

func parsePart2(inp []string) []*machine {
	machines := parsePart1(inp)

	for _, machine := range machines {
		machine.xPrize += 10000000000000
		machine.yPrize += 10000000000000
	}

	return machines
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(machines []*machine) int {
	var sum int

	for _, machine := range machines {
		bDem := (machine.xB * machine.yA) - (machine.yB * machine.xA)
		bNum := (machine.xPrize * machine.yA) - (machine.yPrize * machine.xA)
		b := bNum / bDem
		a := (machine.xPrize - (b * machine.xB)) / machine.xA

		xCalc, yCalc := a*machine.xA+b*machine.xB, a*machine.yA+b*machine.yB
		if xCalc != machine.xPrize || yCalc != machine.yPrize {
			continue
		}

		sum += (3 * a) + b
	}

	return sum
}

func Part2(machines []*machine) int {
	return Part1(machines)
}
