package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(content), "\n")
	n := noOfValids2(input)
	fmt.Println(n)
}

func noOfValids(lines []string) int {
	count := 0

	for _, pass := range lines {
		valid := true
		present := make(map[string]bool)
		passes := strings.Split(string(pass), " ")
		for _, p := range passes {
			if present[p] {
				valid = false
				break
			} else {
				present[p] = true
			}
		}

		if valid {
			count++
		}

	}

	return count
}

func noOfValids2(lines []string) int {
	count := 0

	for _, pass := range lines {
		valid := true
		present := make(map[string]bool)
		passes := strings.Split(string(pass), " ")
		for _, p := range passes {
			if present[p] {
				valid = false
				break
			} else {
				present[p] = true
			}
		}

		if !valid {
			continue
		}

		for i := 0; i < len(passes); i++ {
			for j := i + 1; j < len(passes); j++ {
				if isAnagram(passes[i], passes[j]) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			count++
		}

	}

	return count
}

func isAnagram(s1, s2 string) bool {
	s1Map := make(map[rune]int)
	for _, v := range s1 {
		s1Map[v]++
	}

	s2Map := make(map[rune]int)
	for _, v := range s2 {
		s2Map[v]++
	}

	return reflect.DeepEqual(s1Map, s2Map)

}
