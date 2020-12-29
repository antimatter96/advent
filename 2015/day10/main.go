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
		fmt.Println(scanner.Err().Error())
	}

	day1(inp)
}

const loops = 50

func day1(inps []string) {
	inp := inps[0]

	for i := 0; i < loops; i++ {
		inp = expand(inp)
	}

	fmt.Println(len(inp))
}

func expand(inp string) string {
	var s strings.Builder
	s.Grow(len(inp) * 2)

	runningLength := 0
	var runeWas rune

	for i, v := range inp {
		if i == 0 {
			runeWas = v
			runningLength++
			continue
		}

		if v != runeWas {
			writeThis(&s, runeWas, runningLength)
			// write runewas runningLength times
			runeWas = v
			runningLength = 1
		} else {
			runningLength++
		}
	}
	writeThis(&s, runeWas, runningLength)

	return s.String()
}

func writeThisNtimes(s *strings.Builder, r rune, n int) {
	for i := 0; i < n; i++ {
		s.WriteRune(r)
	}
}

func writeThis(s *strings.Builder, r rune, n int) {
	s.WriteString(fmt.Sprintf("%d%d", n, r-48))
}
