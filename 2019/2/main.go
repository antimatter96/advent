package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(memory []int) {
	length := len(memory)
	for i := 0; i < length; i += 4 {
		if memory[i] == 99 {
			break
		}
		if memory[i] == 1 {
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		} else if memory[i] == 2 {
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		}

	}
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	str := ""
// 	for scanner.Scan() {
// 		str +=scanner.Text()
// 	}

// 	if scanner.Err() != nil {
// 		// handle error.
// 	}
// 	fmt.Println(str)
// 	strs := strings.Split(str, ",")
// 	memory := make([]int, len(strs))

// 	for i, stringNumber := range strs {
//     intNumber, err := strconv.Atoi(stringNumber)
//     if err != nil {
//       panic(err)
//     }
//     memory[i] = intNumber
// 	}

// 	memory[1] = 12
// 	memory[2] = 2

// 	calculate(memory)
// 	fmt.Println(memory[0])
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}
	fmt.Println(str)
	strs := strings.Split(str, ",")
	memory := make([]int, len(strs))
	temp := make([]int, len(strs))

	for i, stringNumber := range strs {
		intNumber, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}
		memory[i] = intNumber
	}

	found := false
	for i := 0; i < 99; i++ {
		memory[1] = i
		if found {
			break
		}
		for j := 0; j < 99; j++ {
			memory[2] = j
			copy(temp, memory)
			calculate(temp)
			if temp[0] == 19690720 {
				fmt.Println((100 * i) + j)
				found = true
				break
			}
		}
	}
}
