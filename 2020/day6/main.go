package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	//inp := ""
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	inp += scanner.Text()
	// }

	// if scanner.Err() != nil {
	// 	// handle error.
	// }

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	day1(string(b))
}

func day1(inp string) {
	total := 0

	xx := extract(inp)

	for _, s := range xx {
		total += common(s)
	}

	fmt.Println(total)
}

func getUnique(arr []string) int {
	mp := make(map[rune]bool)

	for _, s := range arr {
		for _, r := range s {
			mp[r] = true
		}
	}

	//fmt.Println(arr)
	return len(mp)
}

func common(arr []string) int {
	mp := make(map[rune]int)

	for _, s := range arr {
		for _, r := range s {
			mp[r]++
		}
	}

	tots := 0
	for _, r := range mp {
		if r == len(arr) {
			tots++
		}
	}

	//fmt.Println(arr)
	return tots
}

func extract(inp string) [][]string {
	inp = strings.TrimSpace(inp)
	var m [][]string
	ss := strings.Split(inp, "\n\n")

	for _, s := range ss {
		sss := regexp.MustCompile(`[\s:]`).Split(s, -1)
		m = append(m, sss)
	}

	return m
}
