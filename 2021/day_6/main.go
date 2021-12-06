package main

import (
	"bufio"
	"fmt"
	"os"
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

func Run(inp []string) (int64, int64) {
	numbers := parsePart1(inp)

	return Part1(numbers), int64(Part2(numbers))
}

func Part1(numbers []int) int64 {
	mp := make(map[int]int)
	for i := 1; i <= 5; i++ {
		mp[i] = PreCompute(i, 80)
	}

	fmt.Println(mp)
	sum := 0

	for _, n := range numbers {
		sum += mp[n]
	}

	return int64(sum)
}

func Part2(numbers []int) int {
	mp := make(map[int]int)
	for i := 1; i <= 5; i++ {
		mp[i] = PreCompute2(i, 256)
	}

	sum := 0

	for _, n := range numbers {
		sum += mp[n]
	}

	return sum
}

func PreCompute2(n int, days int) int {
	states := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	states[n] = 1

	for day := 1; day <= days; day++ {
		z := states[0]

		states = states[1:]

		states[6] += z

		states = append(states, z)
	}

	sum := 0
	for _, v := range states {
		sum += v
	}

	return sum
}

func printState(states []int) string {
	s := strings.Builder{}

	for i := 0; i < len(states); i++ {
		s.WriteString(fmt.Sprintf("%3d", i))
		s.WriteString(" ")
	}

	s.WriteString(strings.Repeat("----", len(states)))
	s.WriteByte('\n')

	for i := 0; i < len(states); i++ {
		s.WriteString(fmt.Sprintf("%3d", states[i]))
		s.WriteString(" ")
	}

	return s.String()
}
