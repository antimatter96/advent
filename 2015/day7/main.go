package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		fmt.Println(scanner.Err().Error())
	}

	day1(inp)
}

type command struct {
	operation string
	output    string
	operand1  string
	operand2  string
}

func day1(inp []string) {
	cmnds := make([]command, len(inp))
	sort.Strings(inp)
	for _, s := range inp {
		cmnds = append(cmnds, parseCommand(s))
		//fmt.Printf("%+v %s\n", parseCommand(s), s)
	}

	for _, cmnd := range cmnds {
		wires[cmnd.output] = cmnd
	}

	a := value("a", "")

	mem = make(map[string]uint16)

	mem["b"] = a

	fmt.Println(value("a", ""))
	//fmt.Println(a)

}

var (
	wires = make(map[string]command)
	mem   = make(map[string]uint16)
)

func value(wire string, parent string) uint16 {
	//fmt.Println(parent, ">", "FINDING", wire)
	if v, err := strconv.ParseUint(wire, 10, 16); err == nil {
		//fmt.Println(parent, ">", "found", "converted", wire, " is ", uint16(v))
		return uint16(v)
	}

	if val, ok := mem[wire]; ok {
		//fmt.Println(parent, ">", "found", "memoized", wire, " is ", val)
		return val
	}

	cmnd := (wires)[wire]

	//fmt.Println(parent, ">", "finding", cmnd.operand1, cmnd.operation, cmnd.operand2)
	var x uint16
	switch cmnd.operation {
	case "NOT":
		x = ^value(cmnd.operand1, parent+" > "+cmnd.operand1)
	case "SET":
		x = value(cmnd.operand1, parent+" > "+cmnd.operand1)
	case "AND":
		x = value(cmnd.operand1, parent+" > "+cmnd.operand1) & value(cmnd.operand2, parent+" > "+cmnd.operand2)
	case "OR":
		x = value(cmnd.operand1, parent+" > "+cmnd.operand1) | value(cmnd.operand2, parent+" > "+cmnd.operand2)
	case "LSHIFT":
		x = value(cmnd.operand1, parent+" > "+cmnd.operand1) << value(cmnd.operand2, parent+" > "+cmnd.operand2)
	case "RSHIFT":
		x = value(cmnd.operand1, parent+" > "+cmnd.operand1) >> value(cmnd.operand2, parent+" > "+cmnd.operand2)
	}

	//fmt.Println(parent, ">", "found", wire, " == ", x)
	mem[wire] = x
	return x
}

func parseCommand(s string) command {
	cmnd := &command{}
	ss := strings.Split(s, " -> ")
	cmnd.output = ss[1]

	sss := strings.Split(ss[0], " ")

	// Possibly a SET command
	if len(sss) == 1 {
		cmnd.operation = "SET"
		cmnd.operand1 = sss[0]
		return *cmnd
	}

	// Possibly a NOT command
	if len(sss) == 2 {
		cmnd.operation = "NOT"
		cmnd.operand1 = sss[1]
		return *cmnd
	}

	// The Remaining stuff
	if len(sss) == 3 {
		cmnd.operation = sss[1]
		cmnd.operand1 = sss[0]
		cmnd.operand2 = sss[2]

		return *cmnd
	}

	return *cmnd
}
