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

type entry struct {
	Input  [10]string
	Output [4]string
}

func parsePart1(inp []string) []entry {
	entries := make([]entry, len(inp))

	for i, entryString := range inp {
		split := strings.Split(entryString, "|")

		input := strings.Split(strings.TrimSpace(split[0]), " ")
		output := strings.Split(strings.TrimSpace(split[1]), " ")

		inputSorted := make([]string, 0)
		outputSorted := make([]string, 0)

		for _, s := range input {
			str := strings.Split(s, "")
			sort.Strings(str)
			inputSorted = append(inputSorted, strings.Join(str, ""))
		}
		for _, s := range output {
			str := strings.Split(s, "")
			sort.Strings(str)
			outputSorted = append(outputSorted, strings.Join(str, ""))
		}

		copy(entries[i].Input[:], inputSorted)
		copy(entries[i].Output[:], outputSorted)
	}

	return entries
}

func parsePart2(inp []string) []entry {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	numbers := parsePart1(inp)

	return Part1(numbers), Part2(numbers)
}

var segmentsIn = []int{
	6,
	2,
	5,
	5,
	4,
	5,
	6,
	3,
	7,
	6,
}

func Part1(entries []entry) int {
	needCoundtOf := []int{1, 4, 7, 8}
	mp := make(map[int]int)

	for _, e := range needCoundtOf {
		mp[segmentsIn[e]] = e
	}

	count := 0

	for _, entry := range entries {
		for _, output := range entry.Output {
			if _, present := mp[len(output)]; present {
				count++
			}
		}
	}

	return count
}

type magic struct {
	PreComputed []simple
	Left        int
}

type simple struct {
	Known          int
	IntersectsWith int
	HasCommon      int
}

func Part2(entries []entry) int {
	needCoundtOf := []int{1, 4, 7, 8}
	mp := make(map[int]int)

	for _, e := range needCoundtOf {
		mp[segmentsIn[e]] = e
	}

	sum := 0

	for _, entry := range entries {
		thisIs := make(map[string]int)
		knowns := make(map[int]string)

		unknown := make([]string, 0)

		for _, input := range entry.Input {
			if digit, present := mp[len(input)]; present {
				thisIs[input] = digit
				knowns[digit] = input
			}
		}

		for _, output := range entry.Output {
			if digit, present := mp[len(output)]; present {
				thisIs[output] = digit
			} else {
				unknown = append(unknown, output)
			}
		}

		// 	1 ∩ 6 => 1
		// 	4 ∩ 9 => 4
		// 	7 ∩ 6 => 2
		// 	1 ∩ 3 => 2
		// 	4 ∩ 2 => 2
		// 	7 ∩ 3 => 3

		magics := map[int]magic{
			6: {
				PreComputed: []simple{
					{Known: 1, IntersectsWith: 6, HasCommon: 1},
					{Known: 4, IntersectsWith: 9, HasCommon: 4},
					{Known: 7, IntersectsWith: 6, HasCommon: 2},
				},
				Left: 0,
			},
			5: {
				PreComputed: []simple{
					{Known: 1, IntersectsWith: 3, HasCommon: 2},
					{Known: 4, IntersectsWith: 2, HasCommon: 2},
					{Known: 7, IntersectsWith: 3, HasCommon: 3},
				},
				Left: 5,
			},
		}

		for _, s := range unknown {
			for _, preComputed := range magics[len(s)].PreComputed {
				known := knowns[preComputed.Known]
				canFind, target := preComputed.IntersectsWith, preComputed.HasCommon

				if digitsCommon(known, s) == target {
					thisIs[s] = canFind
				}
			}

			if _, present := thisIs[s]; !present {
				thisIs[s] = magics[len(s)].Left
			}
		}

		actual := ""
		for _, output := range entry.Output {
			actual += strconv.Itoa(thisIs[output])
		}

		value, _ := strconv.Atoi(actual)
		sum += value
	}

	return sum
}

func digitsCommon(s1, s2 string) int {
	n := 0

	mp := make(map[rune]struct{})
	for _, r := range s1 {
		mp[r] = struct{}{}
	}

	for _, r := range s2 {
		if _, present := mp[r]; present {
			n++
		}
	}

	return n
}
