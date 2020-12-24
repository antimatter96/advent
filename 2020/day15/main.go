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
		panic(scanner.Err().Error())
	}

	day1(inp)
}

func day1(unparsed []string) {
	ss := unparsed[0]
	sss := strings.Split(ss, ",")

	lastSpoken := -1

	mostRecent := make(map[int]int)

	for i, s := range sss {
		v, _ := strconv.Atoi(s)
		mostRecent[v] = i
		lastSpoken = v
	}

	nextSpoken := -1

	var smr int
	var present bool

	for i := len(sss); i < 30000000; i++ {
		smr, present = mostRecent[lastSpoken]

		if !present {
			mostRecent[lastSpoken] = i - 1
			lastSpoken = 0
		} else {
			nextSpoken = i - 1 - smr
			mostRecent[lastSpoken] = i - 1
			lastSpoken = nextSpoken
		}
	}
	fmt.Println(lastSpoken)
}
