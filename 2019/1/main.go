package main

import(
	"os"
	"fmt"
	"bufio"
	"strconv"
)

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	total := 0
// 	for scanner.Scan() {
// 		i, _ := strconv.Atoi(scanner.Text())
// 		total+= ((i/3) - 2)
// 	}

// 	if scanner.Err() != nil {
// 			// handle error.
// 	}
// 	fmt.Println(total)
// }

var mp map[int]int

func totalFuelForModule(i int) int {
	temp := (i/3) - 2
	if temp <= 0 {
		return 0
	}
	value, present := mp[i]
	if present {
		return value
	}
	temp += totalFuelForModule(temp)
	mp[i] = temp
	return temp
}

func main() {
	mp = make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		total += totalFuelForModule(i)
	}

	if scanner.Err() != nil {
		// handle error.
	}
	fmt.Println(total)
}
