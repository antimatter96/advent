package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

const length = 8

func day1(inp []string) {
	total := 0
	for _, s := range inp {
		if extract2(s) {
			total++
		}
	}
	fmt.Println(total)
}

func extract(s string) {
	//fmt.Println(s)

	//ss := regexp.MustCompile(`([a-z\-]+)-`)
	ss := regexp.MustCompile(`(?:(?:([^\[][a-z]+[^\]])|(\[[a-z]+\])))`)
	//ss := regexp.MustCompile(`([a-z\-]+)-([\d]+)\[([a-z]+)\]`)

	sss := ss.FindAllStringSubmatch(s, -1)
	//fmt.Sscanf(s, "%s-%d[%s]", frequencyString, id, checksum)
	for _, a1 := range sss {
		for _, a2 := range a1 {
			fmt.Printf("|%s|%s|\n", s, a2)
		}
	}
}

func extract2(s string) bool {
	ss := strings.ReplaceAll(s, "[", "--")
	sss := strings.ReplaceAll(ss, "]", "--")

	ssss := strings.Split(sss, "--")

	hasAPalindrome := false
	hasEnclosedPalindrome := false

	insideBrackets := false
	for _, re := range ssss {
		//fmt.Println("|", re, "|", insideBrackets)

		if re == "" {
			insideBrackets = !insideBrackets
			continue
		}
		isp := containsABBA(re)
		if isp {
			hasAPalindrome = true
			if insideBrackets {
				hasEnclosedPalindrome = true
				break
			}
		}
		insideBrackets = !insideBrackets
		//fmt.Println("|", re, "|", isp, insideBrackets)
	}

	//fmt.Println(hasAPalindrome, hasEnclosedPalindrome)
	return hasAPalindrome && !hasEnclosedPalindrome
}

func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}
