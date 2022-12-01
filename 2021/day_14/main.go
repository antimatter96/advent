package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type rule struct {
	x, y   byte
	insert string
}

type rule2 struct {
	xy     string
	insert string
}

func parsePart1(inp []string) (string, []rule) {
	rules := make([]rule, 0)

	start := strings.TrimSpace(inp[0])
	for _, s := range inp[1:] {
		s = strings.TrimSpace(s)
		nums := strings.Split(s, "->")

		if len(nums) == 2 {
			var temp rule

			xy := strings.Split(nums[0], "")

			temp.x = xy[0][0]
			temp.y = xy[1][0]

			temp.insert = strings.TrimSpace(nums[1])

			rules = append(rules, temp)
		}
	}

	return start, rules
}

func parsePart2(inp []string) (string, []rule2) {
	rules := make([]rule2, 0)

	start := strings.TrimSpace(inp[0])
	for _, s := range inp[1:] {
		s = strings.TrimSpace(s)
		nums := strings.Split(s, "->")

		if len(nums) == 2 {
			var temp rule2

			temp.xy = strings.TrimSpace(nums[0])
			temp.insert = strings.TrimSpace(nums[1])

			rules = append(rules, temp)
		}
	}

	return start, rules
}

func Run(inp []string) (int, int) {
	p, f := parsePart1(inp)
	p2, f2 := parsePart2(inp)

	return Part1(p, f, 10), Part2(p2, f2, 40)
}

func Part1(start string, rules []rule, days int) int {
	mp := make(map[string]string)

	for _, v := range rules {
		mp[string(v.x)+","+string(v.y)] = v.insert
	}

	for k := 0; k < days; k++ {
		newStart := &strings.Builder{}
		newStart.WriteByte(start[0])

		for i := 0; i < len(start)-1; i++ {
			x := start[i]
			y := start[i+1]

			if insert, present := mp[string(x)+","+string(y)]; present {
				newStart.WriteString(insert)
			}
			newStart.WriteByte(y)
		}
		start = newStart.String()
	}

	cnt := make(map[rune]int)
	for _, r := range start {
		cnt[r]++
	}
	var cntArray []int
	for _, v := range cnt {
		cntArray = append(cntArray, v)
	}
	sort.Ints(cntArray)

	return cntArray[len(cnt)-1] - cntArray[0]
}

func Part2(start string, rules []rule2, days int) int {
	mp := make(map[string]string)

	for _, v := range rules {
		mp[v.xy] = v.insert
	}

	state := make(map[string]int)

	for i := 0; i < len(start)-1; i++ {
		state[string(start[i])+string(start[i+1])]++
	}

	for k := 0; k < days; k++ {
		newState := make(map[string]int)

		for pair, cnt := range state {
			if insert, present := mp[pair]; present {
				newState[string(pair[0])+string(insert[0])] += cnt
				newState[string(insert[0])+string(pair[1])] += cnt
			} else {
				newState[pair] += cnt
			}
		}

		state = newState
	}

	cnt := make(map[byte]int)
	for pair, r := range state {
		cnt[pair[0]] += r
		cnt[pair[1]] += r
	}
	var cntArray []int
	for _, v := range cnt {
		if v%2 == 0 {
			cntArray = append(cntArray, v/2)
		} else {
			cntArray = append(cntArray, (v+1)/2)
		}
	}
	sort.Ints(cntArray)

	return cntArray[len(cnt)-1] - cntArray[0]
}

func mapSum(source map[string]int) int {
	sum := 0
	for _, v := range source {
		sum += v
	}
	return sum
}
