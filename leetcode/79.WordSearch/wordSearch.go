package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	if word == "" {
		return false
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				if found := visitCell(board, i, j, 0, word); found == true {
					return found
				}
			}
		}
	}
	fmt.Println(board)
	return false
}

func visitCell(board [][]byte, i, j, index int, word string) bool {
	if index == len(word)-1 {
		return true
	}

	if word[index] != board[i][j] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = 0
	if i > 0 && board[i-1][j] == word[index+1] {
		if visitCell(board, i-1, j, index+1, word) {
			return true
		}
	}
	if i < len(board)-1 && board[i+1][j] == word[index+1] {
		if visitCell(board, i+1, j, index+1, word) {
			return true
		}
	}
	if j > 0 && board[i][j-1] == word[index+1] {
		if visitCell(board, i, j-1, index+1, word) {
			return true
		}
	}

	if j < len(board[i])-1 && board[i][j+1] == word[index+1] {
		if visitCell(board, i, j+1, index+1, word) {
			return true
		}
	}
	board[i][j] = tmp

	return false
}
