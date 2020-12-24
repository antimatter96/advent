package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func day1(inp []string) {
	for _, s := range inp {
		y := expand2(s)
		fmt.Println(len(y))
		//fmt.Println(len(ss))
	}
}

var rer strings.Builder
var next, repeat int
var c int

var i, j int
var sss []string
var n int

func expand(s string) string {
	rer.Reset()
	//fmt.Println(s)
	n = len(s)
	j = 0
	for i = 0; i < n; i++ {
		if s[i] == '(' {
			for j = i + 1; j < n; j++ {
				if s[j] == ')' {
					break
				}
			}
			//var next, repeat int
			sss = strings.Split(s[i+1:j], "x")
			next, _ = strconv.Atoi(sss[0])
			repeat, _ = strconv.Atoi(sss[1])
			//fmt.Println(s[i:j+1], next, repeat)

			//fmt.Println(s[j+1:j+1+next], repeat)
			for c = 0; c < repeat; c++ {
				rer.WriteString(s[j+1 : j+1+next])
			}
			i = j + next
			//fmt.Println(i, string(s[i]))
		} else {
			rer.WriteByte(s[i])
		}
	}
	return rer.String()
}

var re = regexp.MustCompile(`(\(\d+x\d+\))`)

func expand2(s string) string {
	if re.MatchString(s) {
		return expand2(expand(s))
	}
	return s
}
