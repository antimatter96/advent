package main

import (
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

func badStringRange(board common.Graph[string]) string {
	var str strings.Builder

	for _, row := range board {
		for _, cell := range row {
			str.WriteString(cell)
		}
	}

	res := str.String()

	return res
}

func badStringIterate(board common.Graph[string]) string {
	var str strings.Builder

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			str.WriteString(board[i][j])
		}
	}

	res := str.String()

	return res
}

func badStringIteratePre(board common.Graph[string]) string {
	var str strings.Builder
	str.Grow(len(board) * len(board[0]))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			str.WriteString(board[i][j])
		}
	}

	res := str.String()

	return res
}

// badStringIteratePreWhat
// bottom of the barrel optimisation [not useful]
func badStringIteratePreWhat(board common.Graph[string]) string {
	var str strings.Builder
	str.Grow(len(board) * len(board[0]))

	var i, j int
	m := len(board)
	n := len(board[0])
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			str.WriteString(board[i][j])
		}
	}

	res := str.String()

	return res
}

func badStringRangePre(board common.Graph[string]) string {
	var str strings.Builder

	str.Grow(len(board) * len(board[0]))

	for _, row := range board {
		for _, cell := range row {
			str.WriteString(cell)
		}
	}

	res := str.String()

	return res
}
