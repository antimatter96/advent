package common

import (
	"bufio"
	"os"
)

func TakeInputAsString() string {
	var inp string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text() + "\n"
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return inp
}

func TakeInputAsStringArray() []string {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return inp
}
