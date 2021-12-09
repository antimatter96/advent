package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 26 {
		t.Errorf("expected %d but got %d", 26, p1)
	}

	if p2 != 66582 {
		t.Errorf("expected %d but got %d", 66582, p2)
	}
}

func TestPrecompute(t *testing.T) {
	var mplength6 = map[string]int{
		"abcdef": 9,
		"bcdefg": 6,
		"abcdeg": 0,
	}
	length6 := reflect.ValueOf(mplength6).MapKeys()

	var known = map[string]int{
		"abd":     7,
		"ab":      1,
		"abef":    4,
		"acedgfb": 8, // useless
	}

	var mplength5 = map[string]int{
		"bcdef": 5,
		"acdfg": 2,
		"abcdf": 3,
	}
	length5 := reflect.ValueOf(mplength5).MapKeys()

	for ks, k := range known {

		fmt.Println("\033[1;38;5;208m_____\033[0m")
		for _, s := range length5 {
			fmt.Println(k, "∩", mplength5[s.String()], "=>", digitsCommon(s.String(), ks))
		}
		fmt.Println("\033[1;38;5;208m==========\033[0m")

		for _, s := range length6 {
			fmt.Println(k, "∩", mplength6[s.String()], "=>", digitsCommon(s.String(), ks))
		}

	}
}
