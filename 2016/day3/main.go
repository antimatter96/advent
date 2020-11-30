package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Errorf(scanner.Err().Error())
	}

	day2(inp)
}

func day1(inp []string) {
	fmt.Println(len(inp))
	total := 0
	var x, y, z int
	for _, s := range inp {
		fmt.Sscanf(s, "  %d  %d  %d", &x, &y, &z)
		if isValid(x, y, z) {
			total++
		}
		//total += getReq(s)
	}
	fmt.Println(total)
}

var temp = []int{0, 0, 0}

func isValid(x, y, z int) bool {
	temp[0], temp[1], temp[2] = x, y, z
	sort.Ints(temp)

	return temp[0]+temp[1] > temp[2]
}

func day2(inp []string) {
	//lines := len(inp) / 3
	total := 0
	var x1, y1, z1 int
	var x2, y2, z2 int
	var x3, y3, z3 int

	var sss []string

	for i := 0; i < len(inp); i += 3 {
		var s = ""

		for j := 0; j < 3; j++ {
			s = s + inp[i+j] + "\n"
		}
		sss = append(sss, s)
	}

	for _, s := range sss {
		fmt.Sscanf(s, "  %d  %d  %d\n  %d  %d  %d\n  %d  %d  %d\n", &x1, &x2, &x3, &y1, &y2, &y3, &z1, &z2, &z3)
		if isValid(x1, y1, z1) {
			total++
		}
		if isValid(x2, y2, z2) {
			total++
		}
		if isValid(x3, y3, z3) {
			total++
		}
		//total += getReq(s)
	}
	fmt.Println(total)
}
