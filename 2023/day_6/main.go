package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type race struct {
	duration int
	distance int
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) []race {

	races := make([]race, 0)
	durationStrings := strings.Split(strings.Split(inp[0], ":")[1], " ")
	for _, durationString := range durationStrings {
		if durationString == "" || durationString == " " {
			continue
		}
		duration, _ := strconv.Atoi(durationString)
		races = append(races, race{duration: int(duration)})
	}

	i := 0
	distanceStrings := strings.Split(strings.Split(inp[1], ":")[1], " ")
	for _, distanceString := range distanceStrings {
		if distanceString == "" || distanceString == " " {
			continue
		}
		distance, _ := strconv.Atoi(distanceString)
		races[i].distance = int(distance)

		i++
	}

	for _, r := range races {
		common.Log.Debug().Int("dur", r.duration).Int("dis", r.distance).Send()
	}

	return races
}

func parsePart2(inp []string) []race {
	races := make([]race, 0)
	durationString := strings.ReplaceAll(strings.Split(inp[0], ":")[1], " ", "")
	duration, _ := strconv.Atoi(durationString)
	races = append(races, race{duration: int(duration)})

	i := 0
	distanceString := strings.ReplaceAll(strings.Split(inp[1], ":")[1], " ", "")
	distance, _ := strconv.Atoi(distanceString)
	races[i].distance = int(distance)

	for _, r := range races {
		common.Log.Debug().Int("dur", r.duration).Int("dis", r.distance).Send()
	}

	return races
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(races []race) int {
	ways := int(1)

	for _, r := range races {
		common.Log.Debug().Str("", "").Send()

		// ONLY FOR DEBUGGING PURPOSES
		for i := int(0); i < r.duration+1; i++ {
			if distanceCovered(i, r.duration) > r.distance {
				common.Log.Debug().Int("pressed", i).Int("dis", distanceCovered(i, r.duration)).Int("Want", r.distance).Send()
				break
			}
		}

		// ONLY FOR DEBUGGING PURPOSES
		for i := r.duration; i > -1; i-- {
			if distanceCovered(i, r.duration) > r.distance {
				common.Log.Debug().Int("pressed", i).Int("dis", distanceCovered(i, r.duration)).Int("Want", r.distance).Send()
				break
			}
		}

		common.Log.Debug().Str("eq", fmt.Sprintf("-x^2 + %dx - %d = 0", r.duration, r.distance)).Send()

		a := int(-1)
		b := r.duration
		c := -r.distance

		D := (b * b) - (4 * a * c)
		sqrtD := math.Sqrt(float64(D))

		common.Log.Debug().Int("D", D).Float64("sqrt", sqrtD).Send()

		root1 := (float64(-b) + sqrtD) / (2.0 * float64(a))
		root2 := (float64(-b) - sqrtD) / (2.0 * float64(a))

		common.Log.Debug().Int("Root 1", int(root1)).Int("Root 2", int(root2)).Send()

		possibleRoot1 := int(root1)
		for _, i := range []int{possibleRoot1 - 1, possibleRoot1, possibleRoot1 + 1} {
			common.Log.Debug().Int("possibleRoot1", i).Int("dis", distanceCovered(i, r.duration)).Int("Want", r.distance).Int("diff", distanceCovered(i, r.duration)-r.distance).Send()
			if distanceCovered(i, r.duration) > r.distance {
				possibleRoot1 = i
				break
			}
		}

		possibleRoot2 := int(root2)
		for _, i := range []int{possibleRoot2 + 1, possibleRoot2, possibleRoot2 - 1} {
			common.Log.Debug().Int("possibleRoot2", i).Int("dis", distanceCovered(i, r.duration)).Int("WANT", r.distance).Int("diff", distanceCovered(i, r.duration)-r.distance).Send()
			if distanceCovered(i, r.duration) > r.distance {
				possibleRoot2 = i
				break
			}
		}

		common.Log.Debug().Int("ROOT 1", int(root1)).Int("Root 1 Final", possibleRoot1).Send()
		common.Log.Debug().Int("ROOT 2", int(root2)).Int("Root 2 Final", possibleRoot2).Send()

		ways *= (1 + possibleRoot2 - possibleRoot1)
	}

	return ways
}

func Part2(races []race) int {

	return Part1(races)
}

func distanceCovered(pressed, totalTime int) int {

	return pressed * (totalTime - pressed)
}
