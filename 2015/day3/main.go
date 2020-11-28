package main

import (
	"bufio"
	"fmt"
	"os"
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

	day1_2(inp)
}

func day1(inp string) {
	fmt.Println(inp)
	mp := make(map[string]int)
	x, y := 0, 0
	mp[stringify(x, y)]++
	for _, b := range inp {
		x, y = nextPosition(x, y, b)
		mp[stringify(x, y)]++
	}
	fmt.Println(len(mp))
}

func nextPosition(x, y int, r rune) (int, int) {
	switch r {
	case '^':
		y++
	case 'v':
		y--
	case '>':
		x++
	case '<':
		x--
	}
	return x, y
}

func stringify(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func day1_2(inp string) {
	fmt.Println(inp)
	mp := make(map[string]int)
	x1, y1 := 0, 0
	x2, y2 := 0, 0
	mp[stringify(x1, y1)]++
	mp[stringify(x2, y2)]++
	for i, b := range inp {
		if i%2 == 0 {
			x1, y1 = nextPosition(x1, y1, b)
			mp[stringify(x1, y1)]++
		} else {
			x2, y2 = nextPosition(x2, y2, b)
			mp[stringify(x2, y2)]++
		}

	}
	fmt.Println(len(mp))
}
