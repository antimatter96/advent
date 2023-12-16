package main

import (
	"github.com/antimatter96/advent/2023/common"
)

func shiftNorth(board common.Graph[string]) {

	for j := 0; j < len(board[0]); j++ {

		changed := true
		for changed {
			changed = false

			for i := 0; i < len(board)-1; i++ {
				if board[i][j] == "." && board[i+1][j] == "O" {
					board[i][j] = "O"
					board[i+1][j] = "."
					changed = true
				}
			}

		}
	}
}

func shiftSouth(board common.Graph[string]) {

	for j := 0; j < len(board[0]); j++ {

		changed := true
		for changed {
			changed = false

			for i := len(board) - 1; i > 0; i-- {
				if board[i][j] == "." && board[i-1][j] == "O" {
					board[i][j] = "O"
					board[i-1][j] = "."
					changed = true
				}
			}

		}
	}
}

func shiftWest(board common.Graph[string]) {

	for i := 0; i < len(board); i++ {

		changed := true
		for changed {
			changed = false

			for j := 0; j < len(board[0])-1; j++ {
				if board[i][j] == "." && board[i][j+1] == "O" {
					board[i][j] = "O"
					board[i][j+1] = "."
					changed = true
				}
			}

		}
	}
}

func shiftEast(board common.Graph[string]) {
	for i := 0; i < len(board); i++ {

		changed := true
		for changed {
			changed = false

			for j := len(board[0]) - 1; j > 0; j-- {
				if board[i][j] == "." && board[i][j-1] == "O" {
					board[i][j] = "O"
					board[i][j-1] = "."
					changed = true
				}
			}

		}
	}

}
