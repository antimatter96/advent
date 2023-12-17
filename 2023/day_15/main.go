package main

import (
	"fmt"
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

func parsePart1(inp string) []string {
	inp = strings.TrimSpace(inp)

	inps := strings.Split(inp, ",")

	for _, i := range inps {
		common.Log.Debug().Str("", i).Int("len", len(i)).Send()
	}

	return inps
}

const (
	OPERATION_DASH = iota
	OPERATION_EQUALS
)

type step struct {
	label       string
	op          int
	focalLength int

	original string
}

func parsePart2(inp string) []step {
	inp = strings.TrimSpace(inp)

	stepStrs := strings.Split(inp, ",")

	steps := make([]step, 0, len(stepStrs))

	for _, stepStr := range stepStrs {
		var label string
		var op int
		var focalLength int

		if stepStr[len(stepStr)-1] == '-' {
			op = OPERATION_DASH

			label = stepStr[0 : len(stepStr)-1]
		} else {
			op = OPERATION_EQUALS

			split := strings.Split(stepStr, "=")

			label = split[0]
			focalLength, _ = strconv.Atoi(split[1])
		}

		steps = append(steps, step{
			label:       label,
			op:          op,
			focalLength: focalLength,

			original: stepStr,
		})

	}

	return steps
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(steps []string) int {
	sum := 0

	for _, step := range steps {
		h := hash(step)
		sum += h

		common.Log.Debug().Str("", step).Int("hash", h).Send()
	}

	return sum
}

type box struct {
	positions map[string]int

	arr []step
}

func (b *box) String() string {
	str := strings.Builder{}

	for i, lenes := range b.arr {
		str.WriteString(fmt.Sprintf("%d: [ %s %d ] ", i, lenes.label, lenes.focalLength))
	}

	return str.String()
}

func (b *box) add(s step) {
	positionInArray, present := b.positions[s.label]

	if present {
		b.arr[positionInArray].focalLength = s.focalLength

		return
	}

	b.positions[s.label] = len(b.arr)

	b.arr = append(b.arr, s)

	common.Log.Debug().Int("arr len", len(b.arr)).Send()
}

func (b *box) minus(s step) {
	positionInArray, present := b.positions[s.label]

	if !present {
		return
	}

	delete(b.positions, s.label)

	common.Log.Debug().Str("arr before", b.String()).Send()

	b.arr = append(b.arr[0:positionInArray], b.arr[positionInArray+1:]...)

	common.Log.Debug().Str("arr after", b.String()).Send()

	for i := positionInArray; i < len(b.arr); i++ {
		b.positions[b.arr[i].label] = i
	}
}

func Part2(steps []step) int {
	sum := 0

	boxes := make([]box, 256)

	for i := 0; i < len(boxes); i++ {
		boxes[i] = box{positions: map[string]int{}, arr: []step{}}
	}

	for stepNo, step := range steps {
		boxNo := hash(step.label)

		common.Log.Debug().Int("_", stepNo+1).Str("label", step.original).Int("boxNo", boxNo).Send()

		if step.op == OPERATION_EQUALS {
			boxes[boxNo].add(step)
		} else {
			boxes[boxNo].minus(step)
		}

		printBoxes(boxes)
	}

	sum = sumBoxes(boxes)

	return sum
}

func hash(inpString string) int {
	hash := 0

	for i := 0; i < len(inpString); i++ {
		hash += int(inpString[i])

		hash *= 17

		hash = hash % 256
	}

	return int(hash)
}

func printBoxes(boxes []box) {
	common.Log.Debug().Str("", "======").Send()

	for i, box := range boxes {
		if len(box.arr) == 0 {
			continue
		}

		common.Log.Debug().Int("Box", i).Str("Label", box.String()).Send()
	}
}

func sumBoxes(boxes []box) int {
	sum := 0

	for i, box := range boxes {
		if len(box.arr) == 0 {
			continue
		}

		for positionInArray, step := range box.arr {
			power := (i + 1) * (positionInArray + 1) * step.focalLength
			sum += power

			common.Log.Debug().Int("Box", i).Str("Label", step.label).Int("M", step.focalLength).Int("Power", power).Send()
		}
	}

	return sum
}
