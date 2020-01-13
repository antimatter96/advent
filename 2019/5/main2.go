package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mode int

//type opcode int

const (
	POS mode = iota
	IMM
	REL
)

type instructionSet struct {
	opcode     int
	noOfparams int
	modes      []mode
}

var paramsForOpcode = map[int]int{
	99: 0,
	3:  1,
	4:  1,
	1:  3,
	2:  3,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
}

func getInstruction(memory []int, i int) instructionSet {
	var ins instructionSet
	instruction := memory[i]
	ins.opcode = instruction % 100
	ins.noOfparams = paramsForOpcode[ins.opcode]
	if ins.noOfparams > 0 {
		ins.modes = make([]mode, 3)
		instruction /= 100

		for i := 0; i < ins.noOfparams; i++ {
			//fmt.Println(instruction)
			ins.modes[i] = mode(instruction % 10)
			instruction /= 10
		}
	}

	return ins
}

func getData(memory []int, i int, mode mode) int {
	next := memory[i]
	switch mode {
	case POS:
		return memory[next]
	case IMM:
		return next
	}
	panic("Unkwon mode")
	return 0
}

func Calculate(memory []int, inputs chan int, out chan int) {
	length := len(memory)
	for i := 0; i < length; {

		ins := getInstruction(memory, i)
		//		fmt.Println(i, fmt.Sprintf("%05d", memory[i]), fmt.Sprintf("%2d", ins.opcode))
		switch ins.opcode {
		case 99:
			close(out)
			return
		case 3:
			loc := memory[i+1]
			memory[loc] = <-inputs
		case 4:
			d := getData(memory, i+1, ins.modes[0])
			out <- d
		case 1:
			op1 := getData(memory, i+1, ins.modes[0])
			op2 := getData(memory, i+2, ins.modes[1])
			loc := memory[i+3]
			memory[loc] = op1 + op2
		case 2:
			op1 := getData(memory, i+1, ins.modes[0])
			op2 := getData(memory, i+2, ins.modes[1])
			loc := memory[i+3]
			memory[loc] = op1 * op2
		case 5:
			test := getData(memory, i+1, ins.modes[0])
			loc := getData(memory, i+2, ins.modes[1])
			if test != 0 {
				i = loc
				i -= (ins.noOfparams + 1)
			}
		case 6:
			test := getData(memory, i+1, ins.modes[0])
			loc := getData(memory, i+2, ins.modes[1])
			if test == 0 {
				i = loc
				i -= (ins.noOfparams + 1)
			}
		case 7:
			op1 := getData(memory, i+1, ins.modes[0])
			op2 := getData(memory, i+2, ins.modes[1])
			loc := memory[i+3]
			if op1 < op2 {
				memory[loc] = 1
			} else {
				memory[loc] = 0
			}
		case 8:
			op1 := getData(memory, i+1, ins.modes[0])
			op2 := getData(memory, i+2, ins.modes[1])
			loc := memory[i+3]
			if op1 == op2 {
				memory[loc] = 1
			} else {
				memory[loc] = 0
			}
		}
		i += ins.noOfparams + 1
	}
	fmt.Println("Done")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}
	//fmt.Println(str)
	strs := strings.Split(str, ",")
	memory := make([]int, len(strs))

	for i, stringNumber := range strs {
		intNumber, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}
		memory[i] = intNumber
	}

	in := make(chan int, 1)
	out := make(chan int)
	in <- 5
	go Calculate(memory, in, out)
	for outp := range out {
		fmt.Println("output", outp)
	}
	close(in)
}
