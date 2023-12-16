package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const p1Answer int = 136
const p2Answer int = 64

var testInput1 = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

var testInput2 = testInput1

func TestRun(t *testing.T) {
	p1, _ := Run(testInput1)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	_, p2 := Run(testInput2)
	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}

func createInput() []string {
	_ = 10 + rand.Intn(91)
	_ = 10 + rand.Intn(91)

	m := 100
	n := 100

	options := []string{"#", ".", "O"}
	testInputNew := make([]string, 0, m)

	fmt.Println(m, n)

	for i := 0; i < m; i++ {
		row := strings.Builder{}
		row.Grow(n)

		for j := 0; j < n; j++ {
			row.WriteString(options[rand.Intn(len(options))])
		}

		testInputNew = append(testInputNew, row.String())
	}

	return testInputNew
}
