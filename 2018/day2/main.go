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

	day2(inp)
}

func day1(arr []string) {
	twos := 0
	threes := 0

	for _, s := range arr {
		x, y := present(s)

		if x {
			twos++
		}
		if y {
			threes++
		}
	}

	fmt.Println(twos * threes)
}

func present(s string) (bool, bool) {
	mp := make(map[rune]int)

	for _, r := range s {
		mp[r]++
	}

	twos := false
	threes := false

	for _, cnt := range mp {
		if cnt == 3 {
			threes = true
		}
		if cnt == 2 {
			twos = true
		}
	}

	return twos, threes
}

func diff(s1, s2 string) int {
	cnt := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			cnt++
		}
	}
	return cnt
}

func findCommon(s1, s2 string) string {
	var common []byte
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			common = append(common, s1[i])
		}
	}

	return string(common)
}

func day2(arr []string) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if diff(arr[i], arr[j]) == 1 {
				fmt.Println(findCommon(arr[i], arr[j]))
			}
		}
	}

}
