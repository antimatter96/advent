package main

import (
	"testing"
)

const p1Answer int = 2024
const p2Answer int = 0

var testInput1 = []string{
	"x00: 1",
	"x01: 0",
	"x02: 1",
	"x03: 1",
	"x04: 0",
	"y00: 1",
	"y01: 1",
	"y02: 1",
	"y03: 1",
	"y04: 1",
	"",
	"ntg XOR fgs -> mjb",
	"y02 OR x01 -> tnw",
	"kwq OR kpj -> z05",
	"x00 OR x03 -> fst",
	"tgd XOR rvg -> z01",
	"vdt OR tnw -> bfw",
	"bfw AND frj -> z10",
	"ffh OR nrd -> bqk",
	"y00 AND y03 -> djm",
	"y03 OR y00 -> psh",
	"bqk OR frj -> z08",
	"tnw OR fst -> frj",
	"gnj AND tgd -> z11",
	"bfw XOR mjb -> z00",
	"x03 OR x00 -> vdt",
	"gnj AND wpb -> z02",
	"x04 AND y00 -> kjc",
	"djm OR pbm -> qhw",
	"nrd AND vdt -> hwm",
	"kjc AND fst -> rvg",
	"y04 OR y02 -> fgs",
	"y01 AND x02 -> pbm",
	"ntg OR kjc -> kwq",
	"psh XOR fgs -> tgd",
	"qhw XOR tgd -> z09",
	"pbm OR djm -> kpj",
	"x03 XOR y03 -> ffh",
	"x00 XOR y04 -> ntg",
	"bfw OR bqk -> z06",
	"nrd XOR fgs -> wpb",
	"frj XOR qhw -> z04",
	"bqk OR frj -> z07",
	"y03 OR x01 -> nrd",
	"hwm AND bqk -> z03",
	"tgd XOR rvg -> z12",
	"tnw OR pbm -> gnj",
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
