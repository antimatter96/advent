package main

import (
	"bufio"
	"fmt"
	"os"
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

const length = 8

func day1(inp []string) {
	total := 0
	var max, min int
	var char, pass string
	for _, s := range inp {
		fmt.Sscanf(s, "%d-%d %s %s", &min, &max, &char, &pass)
		if checkLimit2(min, max, char[0], pass) {
			total++
		}
	}
	fmt.Println(total)
}

func checkLimit(min, max int, b byte, pass string) bool {
	count := 0
	for i := 0; i < len(pass); i++ {
		if pass[i] == b {
			count++
			if count > max {
				//fmt.Println("count", count, min, max, pass)
				return false
			}
		}
	}

	//fmt.Println("count", count, min, max, pass)
	if count >= min {
		return true
	}

	return false
}

func checkLimit2(min, max int, b byte, pass string) bool {
	found := false
	min--
	max--

	if pass[min] == b {
		found = !found
	}
	if pass[max] == b {
		found = !found
	}

	return found
}
