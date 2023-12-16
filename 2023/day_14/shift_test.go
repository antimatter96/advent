package main

import (
	"testing"

	"github.com/antimatter96/advent/2023/common"
)

type shiftFuncImp struct {
	fn    func(board common.Graph[string])
	label string
}

var shiftImplementationTestArrays = []struct {
	name string
	fns  []shiftFuncImp
}{
	{
		name: "north",
		fns: []shiftFuncImp{
			{
				label: "shiftNorth",
				fn:    shiftNorth,
			},
			{
				label: "shiftNorthNew",
				fn:    shiftNorthNew,
			},
		},
	},
	{
		name: "south",
		fns: []shiftFuncImp{
			{
				label: "shiftSouth",
				fn:    shiftSouth,
			},
			{
				label: "shiftSouthNew",
				fn:    shiftSouthNew,
			},
		},
	},
	{
		name: "east",
		fns: []shiftFuncImp{
			{
				label: "shiftEast",
				fn:    shiftEast,
			},
			{
				label: "shiftEastNew",
				fn:    shiftEastNew,
			},
		},
	},
	{
		name: "west",
		fns: []shiftFuncImp{
			{
				label: "shiftWest",
				fn:    shiftWest,
			},
			{
				label: "shiftWestNew",
				fn:    shiftWestNew,
			},
		},
	},
}

func TestDifferentShiftImplementations(t *testing.T) {
	testInput := createInput()

	for _, testCase := range shiftImplementationTestArrays {
		parsed1 := parsePart1(testInput)

		testCase.fns[0].fn(parsed1)
		expected := badString(parsed1)

		for i := 1; i < len(testCase.fns); i++ {
			parsed2 := parsePart1(testInput)

			testCase.fns[i].fn(parsed2)

			got := badString(parsed2)

			if got != expected {
				t.Errorf("failed %s - %s", testCase.name, testCase.fns[i].label)
			}

		}
	}

}

func BenchmarkShiftImplementations(b *testing.B) {
	testInput := createInput()

	for _, implementations := range shiftImplementationTestArrays {

		for _, imp := range implementations.fns {

			b.Run(implementations.name+" "+imp.label, func(subB *testing.B) {
				parsed1 := parsePart1(testInput)

				for i := 0; i < subB.N; i++ {
					imp.fn(parsed1)
				}
			})

		}
	}

}
