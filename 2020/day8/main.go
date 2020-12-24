package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	day1(inp)
}

type instruction struct {
	opcode  string
	operand int
}

func day1(unparsed []string) {
	var ins []instruction

	for _, u := range unparsed {
		ins = append(ins, *parse(u))
	}

	bruteForce(ins)

	//fmt.Println(acc)
}

func bruteForce(ins []instruction) {
	for i := 0; i < len(ins); i++ {
		if ins[i].opcode == "jmp" || ins[i].opcode == "nop" {
			existing := ins[i].opcode
			if existing == "jmp" {
				ins[i].opcode = "nop"
			} else {
				ins[i].opcode = "jmp"
			}

			acc, loops := seeIfItLoops(ins)

			if !loops {
				fmt.Println(acc)
				break
			} else {
				ins[i].opcode = existing
			}
		}

	}
}

func seeIfItLoops(ins []instruction) (int, bool) {
	visited := make(map[int]bool)

	n := len(ins)

	i := 0
	acc := 0
	for i < n {
		if _, ok := visited[i]; ok {
			return acc, true
		}

		visited[i] = true

		switch ins[i].opcode {
		case "nop":
			i++
		case "acc":
			acc += ins[i].operand
			i++
		case "jmp":
			i += ins[i].operand
		}

	}
	return acc, false
}

func parse(unparsed string) *instruction {
	temp := strings.Split(unparsed, " ")
	operand, _ := strconv.Atoi(temp[1])

	return &instruction{temp[0], operand}

}
