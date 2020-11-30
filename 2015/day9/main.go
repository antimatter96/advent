package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	dist     int
	from, to string
}

type edges []edge

func (s edges) Len() int {
	return len(s)
}
func (s edges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s edges) Less(i, j int) bool {
	return s[i].dist < s[i].dist
}

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}

	day1(inp)
}

func day1(inp []string) {
	total := 0

	var graph []edge
	for _, s := range inp {
		ed := &edge{}
		fmt.Sscanf(s, "%s to %s = %d", &ed.from, &ed.to, &ed.dist)

		graph = append(graph, *ed)
	}

	fmt.Println(total, graph)
}

func day2(inp []string) {
	total := 0

	for _, s := range inp {
		length := 6

		for i := 1; i < len(s)-1; i++ {
			length++

			switch s[i] {
			case '"', '\\', '\'':
				length++
			}
		}

		total += length - len(s)
	}

	fmt.Println(total)
}
