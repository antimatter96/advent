package main

import (
	"fmt"

	"github.com/antimatter96/advent/2025/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type rotation struct {
	direction byte
	distance  int
}

func parsePart1(inp []string) []rotation {
	rotations := make([]rotation, len(inp))

	var direction byte
	var distance int

	for i, s := range inp {

		direction = 'L'
		n, _ := fmt.Sscanf(s, "L%d", &distance)
		if n == 0 {
			direction = 'R'
			fmt.Sscanf(s, "R%d", &distance)
		}

		rotations[i].direction = direction
		rotations[i].distance = distance
	}

	return rotations
}

func parsePart2(inp []string) []rotation {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(rotations []rotation) int {
	count := 0
	position := 50

	for _, rotation := range rotations {
		common.Log.Debug().Int("Distance", rotation.distance).Str("Direction", string(rotation.direction)).Send()

		rotation.distance %= 100
		if rotation.direction == 'R' {
			position += rotation.distance
		} else {
			position -= rotation.distance
		}

		position = (position + 100) % 100
		if position == 0 {
			count++
		}
		common.Log.Debug().Int("Position", position).Send()
	}

	return count
}

func Part2(rotations []rotation) int {
	count := 0
	position := 50

	for _, rotation := range rotations {
		common.Log.Debug().Int("Distance", rotation.distance).Str("Direction", string(rotation.direction)).Send()
		initial := position

		count += (rotation.distance / 100) // complete rotations

		rotation.distance %= 100
		if rotation.direction == 'R' {
			position += rotation.distance
		} else {
			position -= rotation.distance
		}

		common.Log.Debug().Int("Intial", initial).Int("Final", position).Send()
		common.Log.Debug().Int("Count", count).Send()
		common.Log.Debug().Int("New Count", count).Send()

		position = (position + 100) % 100

		if position == initial {
			common.Log.Debug().Str("skipping", "same as final").Send()
			continue
		}
		if initial == 0 {
			common.Log.Debug().Str("skipping", "started at 0").Send()
			continue
		}

		var distanceToZero int
		if rotation.direction == 'L' {
			distanceToZero = initial
		} else {
			distanceToZero = 100 - initial
		}

		common.Log.Debug().Int("distanceToZero", distanceToZero).Int("Moved", rotation.distance).Send()
		if rotation.distance >= distanceToZero {
			count++
		}

		common.Log.Debug().Int("New Count", count).Send()
	}

	return count
}
