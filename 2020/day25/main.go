package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	day22(inp)
}

func day22(inps []string) {
	p1String := inps[0]
	p1, _ := strconv.Atoi(p1String)

	fmt.Println(p1)

	loop1 := (process(p1, 7))

	p2String := inps[1]
	p2, _ := strconv.Atoi(p2String)

	fmt.Println(p2)

	loop2 := (process(p2, 7))

	fmt.Println(loop1, loop2)

	fmt.Println(get(p1, loop2))
	fmt.Println(get(p2, loop1))

}
func get(subject int, loops int) int {
	value := 1

	for i := 0; i < loops; i++ {
		value *= subject
		value = value % 20201227

		// fmt.Println(value)
	}

	return value

}

func process(m int, subject int) int {
	value := 1

	for i := 0; i < 1000000000000; i++ {
		value *= subject
		value = value % 20201227

		if value == m {
			return i + 1
		}
		// fmt.Println(value)
	}

	return 0
}
