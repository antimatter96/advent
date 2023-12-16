package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/antimatter96/advent/2023/common"
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

var implementationTestArrays = []struct {
	name    string
	oldFunc func(board common.Graph[string])
	newFunc func(board common.Graph[string])
}{
	{
		name:    "north",
		oldFunc: shiftNorth,
		newFunc: shiftNorthNew,
	},
	{
		name:    "south",
		oldFunc: shiftSouth,
		newFunc: shiftSouthNew,
	},
	{
		name:    "west",
		oldFunc: shiftWest,
		newFunc: shiftWestNew,
	},
	{
		name:    "east",
		oldFunc: shiftEast,
		newFunc: shiftEastNew,
	},
}

func createInput() []string {
	m := 10 + rand.Intn(91)
	n := 10 + rand.Intn(91)

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

func TestDifferentImplementations(t *testing.T) {
	testInput := createInput()

	for _, testCase := range implementationTestArrays {
		parsed1 := parsePart1(testInput)
		parsed2 := parsePart1(testInput)

		testCase.oldFunc(parsed1)
		testCase.newFunc(parsed2)

		expected := badString(parsed1)
		got := badString(parsed2)

		if got != expected {
			t.Errorf("failed %s", testCase.name)
		}
	}

}

func BenchmarkImplementations(b *testing.B) {
	testInput := createInput()

	for _, implementations := range implementationTestArrays {

		parsed1 := parsePart1(testInput)
		parsed2 := parsePart1(testInput)

		b.Run(implementations.name+" old", func(subB *testing.B) {
			for i := 0; i < subB.N; i++ {
				implementations.oldFunc(parsed1)
			}
		})

		b.Run(implementations.name+" new", func(subB *testing.B) {
			for i := 0; i < subB.N; i++ {
				implementations.newFunc(parsed2)
			}
		})
	}

}
