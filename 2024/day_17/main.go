package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type input struct {
	a, b, c  int
	commands []int
}

var M, N int

func parsePart1(inp []string) input {
	info := input{}

	fmt.Sscanf(inp[0], "Register A: %d", &(info.a))
	fmt.Sscanf(inp[1], "Register B: %d", &(info.b))
	fmt.Sscanf(inp[2], "Register C: %d", &(info.c))

	programs := strings.Split(strings.TrimSpace(strings.Split(inp[4], ":")[1]), ",")
	info.commands = make([]int, 0, len(programs))

	for _, program := range programs {
		x, _ := strconv.Atoi(program)
		info.commands = append(info.commands, x)
	}

	return info
}

func parsePart2(inp []string) input {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

type computer struct {
	a, b, c int

	output []int
}

type inputBuffer struct {
	head int
	arr  []int
}

func (i *inputBuffer) getNext() int {
	toReturn := i.arr[i.head]
	i.head++
	return toReturn
}

func (i *inputBuffer) end() bool {
	return i.head >= len(i.arr)
}

func (i *inputBuffer) set(newPostion int) {
	i.head = newPostion
}
func (c *computer) getOperand(operand int) int {
	if operand <= 3 {
		return operand
	}

	if operand == 4 {
		return c.a
	}
	if operand == 5 {
		return c.b
	}
	if operand == 6 {
		return c.c
	}

	return -1
}

func (c *computer) addToOutput(value int) {
	c.output = append(c.output, value)
}

func (c *computer) proceed(programs *inputBuffer) int {
	opcode := programs.getNext()
	// fmt.Println("running", opcode, programs.arr[programs.head:], c.a, c.b, c.c)
	// fmt.Println("out", c.output)

	combo := programs.getNext()

	switch opcode {
	case 0:
		combo = c.getOperand(combo)
		c.a = c.a / (1 << combo)
	case 1:
		c.b = c.b ^ combo
	case 2:
		combo = c.getOperand(combo)
		c.b = combo & 7
	case 3:
		if c.a != 0 {
			programs.set(combo)
		}
	case 4:
		c.b = (c.b ^ c.c)
	case 5:
		combo = c.getOperand(combo)
		c.addToOutput(combo % 8)
	case 6:
		combo = c.getOperand(combo)
		c.b = c.a / (1 << combo)
	case 7:
		combo = c.getOperand(combo)
		c.c = c.a / (1 << combo)
	}

	return -1
}

func Part1(info input) int {
	c := &computer{a: info.a, b: info.b, c: info.c, output: make([]int, 0, len(info.commands))}
	inputBuffer := &inputBuffer{arr: info.commands, head: 0}

	for !inputBuffer.end() {
		c.proceed(inputBuffer)
	}

	fmt.Println(c.output)
	return 0
}

func Part2(info input) int {
	expectedOutput := make([]int, len(info.commands))
	copy(expectedOutput, info.commands)

	// These all produce same length things
	// HI := 281474976710656 // 2^48
	LO := 35184372088832 // 2^45

	// LO = 216122776687000
	// LO = 216133739272639
	HI := 236581645541055 // ONE POSSIBLE ANSWER
	LO = 211106232532991  // LAST DIGIT BECOMES SAME
	// inr := 2

	//	inFocus := 5

	bfr := &inputBuffer{arr: info.commands, head: 0}
	c := &computer{output: make([]int, 0, len(info.commands))}

	total := HI - LO
mainLoop:
	for i := LO; i <= HI; i += 1 {
		diff := i - LO
		if diff%1_000_000_000 == 0 {
			fmt.Println(i, 100*diff/total)
		}
		//
		bfr.head = 0
		clear(c.output)
		c.output = c.output[:0]

		c.a = i
		c.b = info.b + 0
		c.c = info.c + 0
		//
		for !bfr.end() {
			c.proceed(bfr)
			if !slices.Equal(expectedOutput[:len(c.output)], c.output[:len(c.output)]) {
				continue mainLoop
			}
		}
		//

		fmt.Println(i, (c.output), strconv.FormatInt(int64(i), 2))

		// binary := []rune(strconv.FormatInt(int64(i), 2))
		// slices.Reverse(binary)
	}

	return 0
}

func compareArrays(a, b []int) string {
	arr := strings.Builder{}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			arr.WriteRune('#')
		} else {
			arr.WriteRune('.')
		}
	}

	return arr.String()
}

func asString(a []int) string {
	arr := strings.Builder{}

	for i := 0; i < len(a); i++ {
		arr.WriteByte(byte(a[i]) + 48)
	}

	return arr.String()
}
