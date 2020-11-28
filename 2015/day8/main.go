package main

import (
	"bufio"
	"fmt"
	"os"
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

	day2(inp)
}

func day1(inp []string) {
	total := 0

	for _, s := range inp {
		length := -2

		for i := 0; i < len(s); i++ {
			length++

			if s[i] != '\\' {
				continue
			}

			switch s[i+1] {
			case 'x':
				i += 3
			case '"', '\\', '\'':
				i++
			}
		}

		total += len(s) - length
	}

	fmt.Println(total)
}

func day2(inp []string) {
	total := 0

	for _, s := range inp {
		length := 6

		for i := 1; i < len(s)-1; i++ {
			length++

			switch s[i] {
			case '"', '\\', '\'':
				length++
			}
		}

		total += length - len(s)
	}

	fmt.Println(total)
}
