package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 26397 {
		t.Errorf("expected %d but got %d", 26397, p1)
	}

	if p2 != 288957 {
		t.Errorf("expected %d but got %d", 288957, p2)
	}
}
