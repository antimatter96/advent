package main

import (
	"bufio"
	"os"
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

	day(inp)
}

type point struct {
	x, y int
}

func day(arr []string) {

}
