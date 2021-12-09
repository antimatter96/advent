package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rawInput := takeInput()

	fmt.Println(Run(rawInput))
}

func takeInput() []string {
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

func parsePart1(inp []string) []int {
	timers := make([]int, 0)

	timers_string := strings.Split(inp[0], ",")
	for _, s := range timers_string {
		temp, _ := strconv.Atoi(s)
		timers = append(timers, temp)
	}

	return timers
}

func parsePart2(inp []string) []int {
	return parsePart1(inp)
}

func PreCompute(n int, days int) int {
	arr := []int{n}

	for day := 0; day < days; day++ {
		l := len(arr)

		for i := 0; i < l; i++ {
			arr[i]--
			if arr[i] == -1 {
				arr[i] = 6
				arr = append(arr, 8)
			}
		}
	}

	return len(arr)
}

func Run(inp []string) (int, int64) {
	numbers := parsePart1(inp)

	return Part1(numbers), int64(Part2(numbers))
}

func Part1(numbers []int) int {
	sort.Ints(numbers)
	median := numbers[len(numbers)/2]

	sum := 0

	for _, v := range numbers {
		sum += abs(v - median)
	}

	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func Part2(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += (v)
	}
	chance1 := math.Floor(float64(sum) / float64(len(numbers)))
	chance2 := math.Ceil(float64(sum) / float64(len(numbers)))

	sum1 := 0
	for _, v := range numbers {
		diff := abs(v - int(chance1))
		sum1 += diff * (diff + 1) / 2
	}

	sum2 := 0
	for _, v := range numbers {
		diff := abs(v - int(chance2))
		sum2 += diff * (diff + 1) / 2
	}

	if sum1 < sum2 {
		return sum1
	}

	return sum2
}
