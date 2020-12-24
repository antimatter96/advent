package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func day1(inp []string) {
	fmt.Println(inp)
	total := 0
	for _, s := range inp {
		fmt.Println(s)
		total += getReq(s)
	}
	fmt.Println(total)
}

func getReq(inp string) int {
	spl := strings.Split(inp, "x")
	fmt.Println(spl)
	splInt := make([]int, 3)
	for i := 0; i < 3; i++ {
		splInt[i], _ = strconv.Atoi(spl[i])
	}
	sort.Ints(splInt)
	fmt.Println(splInt)
	return getSome_2(splInt[0], splInt[1], splInt[2])
}

func getSome(a, b, c int) int {
	return (2 * a * b) + (2 * b * c) + (2 * c * a) + (a * b)
}

func getSome_2(a, b, c int) int {
	return (a * b * c) + 2*(a+b)
}
