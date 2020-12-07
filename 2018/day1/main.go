package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func day1(arr []string) {
	final := 0
	for _, val := range arr {
		num, _ := strconv.Atoi(val)
		final += num
	}

	fmt.Println(final)
}

func day2(arr []string) {
	reached := make(map[int]int)
	reached[0] = 1
	final := 0

	found := false

	for !found {
		for _, val := range arr {
			num, _ := strconv.Atoi(val)
			final += num

			reached[final]++

			if reached[final] > 1 {
				found = true
				break
			}
		}
	}

	fmt.Println(final)
}
