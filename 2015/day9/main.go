package main

import (
	"bufio"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat/combin"
)

type edge struct {
	dist     int
	from, to string
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

var countryToIndex = map[string]int{}
var indexToCountry = []string{}
var distances = [][]int{}

func day1(inp []string) {
	countryToIndex = make(map[string]int)

	var graph []edge

	i := 0
	for _, s := range inp {
		ed := &edge{}
		fmt.Sscanf(s, "%s to %s = %d", &ed.from, &ed.to, &ed.dist)

		graph = append(graph, *ed)

		if _, ok := countryToIndex[ed.from]; !ok {
			countryToIndex[ed.from] = i
			indexToCountry = append(indexToCountry, ed.from)
			i++
		}

		if _, ok := countryToIndex[ed.to]; !ok {
			countryToIndex[ed.to] = i
			indexToCountry = append(indexToCountry, ed.to)
			i++
		}
	}

	distances = make([][]int, len(indexToCountry))
	for i := 0; i < len(indexToCountry); i++ {
		distances[i] = make([]int, len(indexToCountry))
	}

	for _, ed := range graph {
		distances[countryToIndex[ed.from]][countryToIndex[ed.to]] = ed.dist
		distances[countryToIndex[ed.to]][countryToIndex[ed.from]] = ed.dist
	}

	n := len(distances)
	k := len(distances)
	gen := combin.NewPermutationGenerator(n, k)

	minDist := 1 << 62
	maxDist := 0

	for gen.Next() {
		arr := gen.Permutation(nil)

		dist := 0
		for i := 0; i < len(arr)-1; i++ {
			dist += distances[arr[i]][arr[i+1]]
		}

		if dist < minDist {
			minDist = dist
		}
		if dist > maxDist {
			maxDist = dist
		}
	}

	fmt.Println("==>", minDist)
	fmt.Println("==>", maxDist)
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
