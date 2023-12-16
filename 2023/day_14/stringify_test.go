package main

import (
	"testing"

	"github.com/antimatter96/advent/2023/common"
)

var stringImplementationTestArrays = []struct {
	name string
	fn   func(board common.Graph[string]) string
}{
	{
		name: "badStringIterate",
		fn:   badStringIterate,
	},
	{
		name: "badStringRange",
		fn:   badStringRange,
	},
	{
		name: "badStringIteratePre",
		fn:   badStringIteratePre,
	},
	{
		name: "badStringRangePre",
		fn:   badStringRangePre,
	},
	{
		name: "badStringIteratePreWhat",
		fn:   badStringIteratePreWhat,
	},
}

var str1 string

func BenchmarkStringImplementations(b *testing.B) {
	testInput := createInput()
	parsed := parsePart1(testInput)

	var a2 string
	for _, otherImp := range stringImplementationTestArrays {
		var a3 string
		b.Run(otherImp.name, func(subB *testing.B) {
			var a4 string
			for i := 0; i < subB.N; i++ {
				a4 = otherImp.fn(parsed)
			}
			a3 = a4
		})
		a2 = a3
	}
	str1 = a2
}
