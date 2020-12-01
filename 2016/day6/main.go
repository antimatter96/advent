package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Errorf(scanner.Err().Error())
	}

	day1(inp)
}

const length = 8

func day1(inp []string) {
	freqs := make([]map[rune]int, length)
	for i := 0; i < length; i++ {
		freqs[i] = make(map[rune]int)
	}

	for _, s := range inp {

		for i, r := range s {
			freqs[i][r]++
		}

	}

	//fmt.Println(freqs)
	for i := 0; i < len(inp[0]); i++ {
		xx := rankMapStringInt(freqs[i])
		fmt.Printf(string(xx[len(xx)-1]))
	}
	fmt.Println()

}

func day2(inp []string) {
	freqs := make([]map[rune]int, length)
	for i := 0; i < length; i++ {
		freqs[i] = make(map[rune]int)
	}

	for _, s := range inp {
		for i, r := range s {
			freqs[i][r]++
		}
	}

	for i := 0; i < len(inp[0]); i++ {
		xx := rankMapStringInt(freqs[i])
		fmt.Printf(string(xx[0]))
	}
	fmt.Println()
}

func rankMapStringInt(values map[rune]int) []rune {
	type kv struct {
		Key   rune
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value != ss[j].Value {
			return ss[i].Value > ss[j].Value
		}
		return ss[i].Key < ss[j].Key
	})
	ranked := make([]rune, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
