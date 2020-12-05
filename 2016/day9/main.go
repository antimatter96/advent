package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Errorf(scanner.Err().Error())
	}

	day1(inp)
}

func day1(inp []string) {
	for _, s := range inp {
		_ = expand2(s)
		//fmt.Println(len(ss))
	}
}

func expand(s string) string {
	var rer strings.Builder
	//fmt.Println(s)
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			for j = i + 1; j < len(s); j++ {
				if s[j] == ')' {
					break
				}
			}
			var next, repeat int
			sss := strings.Split(s[i+1:j], "x")
			next, _ = strconv.Atoi(sss[0])
			repeat, _ = strconv.Atoi(sss[1])
			//fmt.Println(s[i:j+1], next, repeat)

			//fmt.Println(s[j+1:j+1+next], repeat)
			for c := 0; c < repeat; c++ {
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

func expand2(s string) string {
	//var rer strings.Builder
	fmt.Println(s)
	//j := 0
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == '(' {
	// 		for j = i + 1; j < len(s); j++ {
	// 			if s[j] == ')' {
	// 				break
	// 			}
	// 		}
	// 		var next, repeat int
	// 		sss := strings.Split(s[i+1:j], "x")
	// 		next, _ = strconv.Atoi(sss[0])
	// 		repeat, _ = strconv.Atoi(sss[1])
	// 		//fmt.Println(s[i:j+1], next, repeat)

	// 		//fmt.Println(s[j+1:j+1+next], repeat)
	// 		for c := 0; c < repeat; c++ {
	// 			rer.WriteString(s[j+1 : j+1+next])
	// 		}
	// 		i = j + next
	// 		//fmt.Println(i, string(s[i]))
	// 	} else {
	// 		rer.WriteByte(s[i])
	// 	}
	// }
	return "rer.String()"
}

// find i == (
// find mathcing j== )

// fmt.sscanf inke beech ka s[i:j] to get next x, repeat n

// use s[j:j+x] to get next x digits

// write in a string, n times

// i is now j + x
