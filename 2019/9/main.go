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
	opcode     int64
	noOfparams int64
	modes      []mode
}

var paramsForOpcode = map[int64]int64{
	99: 0,
	3:  1,
	4:  1,
	1:  3,
	2:  3,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
	9:  1,
}

func getInstruction(memory []int64, i int64) instructionSet {
	var ins instructionSet
	instruction := memory[i]
	ins.opcode = instruction % 100
	ins.noOfparams = paramsForOpcode[ins.opcode]
	if ins.noOfparams > 0 {
		ins.modes = make([]mode, 3)
		instruction /= 100

		for i := int64(0); i < ins.noOfparams; i++ {
			//fmt.Println(instruction)
			ins.modes[i] = mode(instruction % 10)
			instruction /= 10
		}
	}

	return ins
}

func main() {
	//timer1 := time.NewTimer(10 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	if scanner.Err() != nil {
	}
	strs := strings.Split(str, ",")
	//memory := make([]int, len(strs))

	zero := make([]int64, len(strs))

	//fmt.Println(strs)
	for i, stringNumber := range strs {
		intNumber, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}
		zero[i] = int64(intNumber)
	}

	for i := 0; i < 400; i++ {
		c := computer{baseRegister: 0, ip: 0, in: make(chan int64, 1), out: make(chan int64)}
		//fmt.Println(c.memory)
		c.memory = make([]int64, 10*len(zero))
		copy(c.memory, zero)
		//fmt.Println(c.memory)
		c.in <- int64(i)
		go c.Calculate()
		for outp := range c.out {
			fmt.Println(">>>", outp)
		}
		close(c.in)
	}
	//fmt.Println(zero)

}

type computer struct {
	baseRegister int64
	memory       []int64
	ip           int64
	in           chan int64
	out          chan int64
}

func (c *computer) getData(i int64, mode mode) int64 {

	next := c.memory[i]

	//fmt.Println(mode, next, c.baseRegister+next)

	switch mode {
	case POS:
		return c.memory[0+next]
	case IMM:
		return next
	case REL:
		return c.baseRegister + next
	}
	panic("Unkwon mode")
	return 0
}

func (c *computer) Calculate() {
	//fmt.Println(c.ip, c.memory)
	for c.ip < int64(len(c.memory)) {

		ins := getInstruction(c.memory, c.ip)
		//fmt.Println(c.ip, fmt.Sprintf("%05d", c.memory[c.ip]), fmt.Sprintf("%2d", ins.opcode))
		switch ins.opcode {
		case 99:
			close(c.out)
			return
		case 3:
			d := c.getData(c.ip+1, 1)
			switch ins.modes[0] {
			case POS:
				c.memory[d] = <-c.in
			case REL:
				//fmt.Println("Use")
				c.memory[c.baseRegister+d] = <-c.in
			}
		case 4:
			d := c.getData(c.ip+1, ins.modes[0])
			c.out <- d
		case 9:
			c.baseRegister += c.memory[c.ip+1]
		case 1:
			op1 := c.getData(c.ip+1, ins.modes[0])
			op2 := c.getData(c.ip+2, ins.modes[1])
			loc := c.memory[c.ip+3]
			c.memory[loc] = op1 + op2
		case 2:
			op1 := c.getData(c.ip+1, ins.modes[0])
			op2 := c.getData(c.ip+2, ins.modes[1])
			loc := c.memory[c.ip+3]
			c.memory[loc] = op1 * op2
		case 5:
			test := c.getData(c.ip+1, ins.modes[0])
			loc := c.getData(c.ip+2, ins.modes[1])
			if test != 0 {
				c.ip = loc
				c.ip -= (ins.noOfparams + 1)
			}
		case 6:
			test := c.getData(c.ip+1, ins.modes[0])
			loc := c.getData(c.ip+2, ins.modes[1])
			if test == 0 {
				c.ip = loc
				c.ip -= (ins.noOfparams + 1)
			}
		case 7:
			op1 := c.getData(c.ip+1, ins.modes[0])
			op2 := c.getData(c.ip+2, ins.modes[1])
			loc := c.memory[c.ip+3]
			if op1 < op2 {
				c.memory[loc] = 1
			} else {
				c.memory[loc] = 0
			}
		case 8:
			op1 := c.getData(c.ip+1, ins.modes[0])
			op2 := c.getData(c.ip+2, ins.modes[1])
			loc := c.memory[c.ip+3]
			if op1 == op2 {
				c.memory[loc] = 1
			} else {
				c.memory[loc] = 0
			}
		}
		c.ip += ins.noOfparams + 1
	}
}
