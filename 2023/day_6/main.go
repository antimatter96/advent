package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type race struct {
	duration int64
	distance int64
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int64("Answer 1", p1).Send()
	common.Log.Info().Int64("Answer 2", p2).Send()
}

func parsePart1(inp []string) []race {

	races := make([]race, 0)
	durationStrings := strings.Split(strings.Split(inp[0], ":")[1], " ")
	for _, durationString := range durationStrings {
		if durationString == "" || durationString == " " {
			continue
		}
		duration, _ := strconv.Atoi(durationString)
		races = append(races, race{duration: int64(duration)})
	}

	i := 0
	distanceStrings := strings.Split(strings.Split(inp[1], ":")[1], " ")
	for _, distanceString := range distanceStrings {
		if distanceString == "" || distanceString == " " {
			continue
		}
		distance, _ := strconv.Atoi(distanceString)
		races[i].distance = int64(distance)

		i++
	}

	for _, r := range races {
		common.Log.Debug().Int64("dur", r.duration).Int64("dis", r.distance).Send()
	}

	return races
}

func parsePart2(inp []string) []race {
	races := make([]race, 0)
	durationString := strings.ReplaceAll(strings.Split(inp[0], ":")[1], " ", "")
	duration, _ := strconv.Atoi(durationString)
	races = append(races, race{duration: int64(duration)})

	i := 0
	distanceString := strings.ReplaceAll(strings.Split(inp[1], ":")[1], " ", "")
	distance, _ := strconv.Atoi(distanceString)
	races[i].distance = int64(distance)

	for _, r := range races {
		common.Log.Debug().Int64("dur", r.duration).Int64("dis", r.distance).Send()
	}

	return races
}

func Run(inp []string) (int64, int64) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(races []race) int64 {
	ways := int64(1)

	for _, r := range races {
		common.Log.Debug().Str("", "").Send()

		// ONLY FOR DEBUGGING PURPOSES
		for i := int64(0); i < r.duration+1; i++ {
			if distanceCovered(i, r.duration) > r.distance {
				common.Log.Debug().Int64("pressed", i).Int64("dis", distanceCovered(i, r.duration)).Int64("Want", r.distance).Send()
				break
			}
		}

		// ONLY FOR DEBUGGING PURPOSES
		for i := r.duration; i > -1; i-- {
			if distanceCovered(i, r.duration) > r.distance {
				common.Log.Debug().Int64("pressed", i).Int64("dis", distanceCovered(i, r.duration)).Int64("Want", r.distance).Send()
				break
			}
		}

		common.Log.Debug().Str("eq", fmt.Sprintf("-x^2 + %dx - %d = 0", r.duration, r.distance)).Send()

		a := int64(-1)
		b := r.duration
		c := -r.distance

		D := (b * b) - (4 * a * c)
		sqrtD := math.Sqrt(float64(D))

		common.Log.Debug().Int64("D", D).Float64("sqrt", sqrtD).Send()

		root1 := (float64(-b) + sqrtD) / (2.0 * float64(a))
		root2 := (float64(-b) - sqrtD) / (2.0 * float64(a))

		common.Log.Debug().Int64("Root 1", int64(root1)).Int64("Root 2", int64(root2)).Send()

		possibleRoot1 := int64(root1)
		for _, i := range []int64{possibleRoot1 - 1, possibleRoot1, possibleRoot1 + 1} {
			common.Log.Debug().Int64("possibleRoot1", i).Int64("dis", distanceCovered(i, r.duration)).Int64("Want", r.distance).Int64("diff", distanceCovered(i, r.duration)-r.distance).Send()
			if distanceCovered(i, r.duration) > r.distance {
				possibleRoot1 = i
				break
			}
		}

		possibleRoot2 := int64(root2)
		for _, i := range []int64{possibleRoot2 + 1, possibleRoot2, possibleRoot2 - 1} {
			common.Log.Debug().Int64("possibleRoot2", i).Int64("dis", distanceCovered(i, r.duration)).Int64("WANT", r.distance).Int64("diff", distanceCovered(i, r.duration)-r.distance).Send()
			if distanceCovered(i, r.duration) > r.distance {
				possibleRoot2 = i
				break
			}
		}

		common.Log.Debug().Int64("ROOT 1", int64(root1)).Int64("Root 1 Final", possibleRoot1).Send()
		common.Log.Debug().Int64("ROOT 2", int64(root2)).Int64("Root 2 Final", possibleRoot2).Send()

		ways *= (1 + possibleRoot2 - possibleRoot1)
	}

	return ways
}

func Part2(races []race) int64 {

	return Part1(races)
}

func distanceCovered(pressed, totalTime int64) int64 {

	return pressed * (totalTime - pressed)
}
