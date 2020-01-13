package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input_1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(reachedTwiceFirst(string(file)))
}

func getFinalFrequency(s string) int {
	final := 0

	split := strings.Split(s, "\n")

	for _, val := range split {
		num, _ := strconv.Atoi(val)
		final += num
	}

	return final
}

func reachedTwiceFirst(s string) int {
	reached := make(map[int]int)
	reached[0] = 1
	final := 0

	split := strings.Split(s, "\n")

	for true {
		fmt.Println(1)
		for _, val := range split {
			num, _ := strconv.Atoi(val)
			final += num
			reached[final]++
			if reached[final] > 1 {
				return final
			}
		}
	}

	return final
}
