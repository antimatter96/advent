package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func main() {
	inp := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}

	day1(inp)
}

func day1(inp string) {
	for i := 0; i < 99999999; i++ {
		data := []byte(fmt.Sprintf("%s%d", inp, i))
		ss := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(ss, "000000") {
			fmt.Println(i)
			break
		}
	}
}

func day1_2(inp string) {
	floor := 0
	printed := false
	for i := 0; i < len(inp); i++ {
		if inp[i] == '(' {
			floor++
		} else if inp[i] == ')' {
			floor--
		}
		if !printed && floor < 0 {
			fmt.Println(i + 1)
			printed = true
		}
	}
	fmt.Println(floor)
}
