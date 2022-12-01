package main

import (
	"fmt"
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 1588 {
		t.Errorf("expected %d but got %d", 1588, p1)
	}

	if p2 != 2188189693529 {
		t.Errorf("expected %d but got %d", 2188189693529, p2)
	}
}

var temp int

func BenchmarkSlower(b *testing.B) {
	var start1, rules1 = parsePart1(strings.Split(testInput, "\n"))

	var x int

	for day := 1; day < 26; day += 2 {
		b.Run(fmt.Sprintf("Bad__%02d", day), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x = Part1(start1, rules1, day)
			}
		})
	}

	temp = x
}

func BenchmarkFaster(b *testing.B) {
	var start2, rules2 = parsePart2(strings.Split(testInput, "\n"))

	var x int

	for day := 1; day < 26; day += 2 {
		b.Run(fmt.Sprintf("Good_%02d", day), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x = Part2(start2, rules2, day)
			}
		})
	}

	temp = x
}

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
