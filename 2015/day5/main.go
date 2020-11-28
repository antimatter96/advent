package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	day1(inp)
}

func day1(inp []string) {
	//fmt.Println(inp)
	total := 0
	for _, s := range inp {
		//fmt.Println(s)
		if isNice2(s) {
			total++
		}
	}
	fmt.Println(total)
}

func isNice(s string) bool {
	return threeVowels(s) && repeat(s) && noBad(s)
}

func threeVowels(s string) bool {
	count := 0
	count += strings.Count(s, "a")
	count += strings.Count(s, "e")
	count += strings.Count(s, "i")
	count += strings.Count(s, "o")
	count += strings.Count(s, "u")

	return count >= 3
}

func repeat(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func noBad(s string) bool {
	return !(strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy"))
}

func isNice2(s string) bool {
	return repeatWithMiddle(s) && repeatTwice(s)
}

func repeatTwice(s string) bool {
	fmt.Println(s)
	for i := 0; i < len(s)-1; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func repeatWithMiddle(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
