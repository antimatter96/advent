package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inp := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}

	day1_2(inp)
}

func day1(inp string) {
	floor := 0
	for i := 0; i < len(inp); i++ {
		if inp[i] == '(' {
			floor++
		} else if inp[i] == ')' {
			floor--
		}
	}
	fmt.Println(floor)
}

func day1_2(inp string) {
	floor := 0
	printed := false
	for i := 0; i < len(inp); i++ {
		if inp[i] == '(' {
			floor++
		} else if inp[i] == ')' {
			floor--
		}
		if !printed && floor < 0 {
			fmt.Println(i + 1)
			printed = true
		}
	}
	fmt.Println(floor)
}
