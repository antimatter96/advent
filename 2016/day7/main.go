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

func extract(s string) bool {
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

func extract2(s string) bool {
	ss := strings.ReplaceAll(s, "[", "--")
	sss := strings.ReplaceAll(ss, "]", "--")

	ssss := strings.Split(sss, "--")

	hasAPalindrome := false
	hasEnclosedPalindrome := false

	insideBrackets := false
	mpOut := make(map[string]bool)
	mpIn := make(map[string]bool)

	for _, re := range ssss {
		//fmt.Println("|", re, "|", insideBrackets)

		if re == "" {
			insideBrackets = !insideBrackets
			continue
		}

		if insideBrackets {
			containsABA(re, mpIn, false)
		} else {
			containsABA(re, mpOut, true)
		}

		//fmt.Printf("%+v\n====\n%+v\n\n", mpIn, mpOut)

		if len(mpIn) > 0 && len(mpOut) > 0 {
			if haveCommon(mpOut, mpIn) {
				return true
			}
		}

		insideBrackets = !insideBrackets
		//fmt.Println("|", re, "|", isp, insideBrackets)
	}

	//fmt.Println(hasAPalindrome, hasEnclosedPalindrome)
	return hasAPalindrome && !hasEnclosedPalindrome
}

func containsABA(s string, mp map[string]bool, invert bool) {
	var aba string
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			if invert {
				aba = fmt.Sprintf("%c%c%c", s[i+1], s[i+2], s[i+1])
			} else {
				aba = s[i : i+3]

			}
			mp[aba] = true
		}
	}
}

func haveCommon(mpIn, mpOut map[string]bool) bool {
	bigger := &mpIn
	smaller := &mpOut

	if len(mpOut) > len(mpIn) {
		bigger = &mpOut
		smaller = &mpIn
	}

	for pattern, _ := range *smaller {
		if _, ok := (*bigger)[pattern]; ok {
			return true
		}
	}

	return false

}
