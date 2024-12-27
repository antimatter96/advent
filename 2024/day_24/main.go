package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type operation string

const (
	XOR operation = "XOR"
	OR  operation = "OR"
	AND operation = "AND"
)

type wiring struct {
	x, y string
	z    string

	op operation
}

type system struct {
	wires map[string]bool

	ws []wiring
}

var M, N int

func parsePart1(inp []string) system {
	sys := system{wires: make(map[string]bool), ws: make([]wiring, 0)}

	i := 0

	var wire string
	var value int

	for ; i < len(inp); i++ {
		if inp[i] == "" {
			break
		}

		fmt.Sscanf(strings.ReplaceAll(inp[i], ":", ""), "%s %d", &wire, &value)
		sys.wires[wire] = (value == 1)
	}

	i++

	for ; i < len(inp); i++ {
		tokens := strings.Split(inp[i], " ")

		w := wiring{x: tokens[0], y: tokens[2], z: tokens[4], op: operation(tokens[1])}

		sys.ws = append(sys.ws, w)
	}

	return sys
}

func parsePart2(inp []string) system {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(sys system) int {
	var ans int

	indegrees := make(map[string][]string)

	for _, wire := range sys.ws {
		if indegrees[wire.z] == nil {
			indegrees[wire.z] = make([]string, 0)
		}
		indegrees[wire.z] = append(indegrees[wire.z], wire.x)
		indegrees[wire.z] = append(indegrees[wire.z], wire.y)
	}
	for wire := range sys.wires {
		if indegrees[wire] == nil {
			indegrees[wire] = make([]string, 0)
		}
	}

	fmt.Println(indegrees)

	forward := make(map[string][]*wiring)
	for i := 0; i < len(sys.ws); i++ {
		if forward[sys.ws[i].x] == nil {
			forward[sys.ws[i].x] = make([]*wiring, 0)
		}
		if forward[sys.ws[i].y] == nil {
			forward[sys.ws[i].y] = make([]*wiring, 0)
		}

		forward[sys.ws[i].x] = append(forward[sys.ws[i].x], &sys.ws[i])
		forward[sys.ws[i].y] = append(forward[sys.ws[i].y], &sys.ws[i])
	}

	updated := make(map[string]bool)

	for len(indegrees) > 0 {
		toDelete := make([]string, 0)

		for wire, dependencies := range indegrees {
			if len(dependencies) == 0 {
				updated[wire] = true
				toDelete = append(toDelete, wire)

				continue
			}

			allDone := true
			for _, upstream := range dependencies {
				if !updated[upstream] {
					allDone = false
					break
				}
			}

			if !allDone {
				continue
			}

			for _, upstream := range dependencies {
				var whichOp *wiring = nil

				for i := 0; i < len(forward[upstream]); i++ {
					if forward[upstream][i].z == wire {
						whichOp = forward[upstream][i]
						break
					}
				}

				x := sys.wires[whichOp.x]
				y := sys.wires[whichOp.y]
				sys.wires[wire] = OP(x, y, whichOp.op)

				fmt.Println(x, whichOp.op, y, "=>", sys.wires[wire], "->", wire)
			}

			updated[wire] = true
			toDelete = append(toDelete, wire)
		}

		for _, wire := range toDelete {
			delete(indegrees, wire)
		}
	}

	z_wires := make([]string, 0)
	for wire := range sys.wires {
		// Can probably do this while updating :shrug
		if strings.HasPrefix(wire, "z") {
			z_wires = append(z_wires, wire)
		}
	}
	slices.Sort(z_wires)

	binary := ""

	for _, wire := range z_wires {
		if sys.wires[wire] {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
	}

	fmt.Sscanf(binary, "%b", &ans)

	return ans
}

func Part2(sys system) int {
	var sum int

	return sum
}

func OP(x, y bool, op operation) bool {
	if op == AND {
		return x && y
	}

	if op == OR {
		return x || y
	}

	if op == XOR {
		return (x && !y) || (y && !x)
	}

	panic("")
}
