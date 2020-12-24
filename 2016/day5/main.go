package main

import (
	"bufio"
	"crypto/md5"
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
		panic(scanner.Err().Error())
	}

	day1(inp)
}

func day1(inp []string) {
	s := inp[0]

	toDo := 8
	found := 0

	//ticker := time.NewTicker(50 * time.Millisecond)

	for i := 0; i < 99999999; i++ {
		// if i%1000000 == 0 {
		// 	fmt.Println(i)
		// }
		//<-ticker.C
		check2(fmt.Sprintf("%s%d", s, i))
	}

	fmt.Println(toDo, found)
}

func check(s string) {
	//fmt.Println("a ...interface{}", s)
	ss := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	if strings.HasPrefix(ss, "00000") {
		fmt.Println(string(ss[5]))
	}
}

var lowerLimit = int('0') - '0'
var upperLimit = int('7') - '0'

func check2(s string) {
	//fmt.Println("a ...interface{}", s)
	ss := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	if strings.HasPrefix(ss, "00000") {
		if int(ss[5])-'0' >= lowerLimit && int(ss[5])-'0' <= upperLimit {
			fmt.Println(string(ss[5]), string(ss[6]))
		}
	}
}
