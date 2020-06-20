/*

Given a tic tac toe board, write a program to detect if any row, colummn, or diagonal has only X's or O's
Return true if it dows, false if not.

XXX
OXO
... would return true.



*/

package main

import (
	"log"
)

func main() {
	tests := [][][]rune{{{'X', 'O', 'O'}, {'X', 'O', 'X'}, {'X', 'X', 'O'}}, {{'X', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'X', 'X'}}, {{'X', 'O', 'X'}, {'O', 'X', 'O'}, {'X', 'O', 'X'}}}

	// tests := []rune{}

	for _, test := range tests {
		log.Printf("chekcWin(%v) == %v\n", test, checkWin(test))
	}
}

func checkWin(board [][]rune) bool {
	colWinner := []bool{}

	for i := 0; i < len(board); i++ {
		var rowStart rune
		// var colStart rune
		// colWinner := []bool{}
		rowWinner := true

		for j := 0; j < len(board[i]); j++ {
			thisCell := board[i][j]

			if i == 0 {
				// initialize row win flag set
				colWinner = append(colWinner, true)
			} else if thisCell != board[0][j] {
				// check if this cell is equal to the column header (board[0] holds the set of column headers)
				colWinner[j] = false
			}

			if j == 0 {
				rowStart = thisCell
			} else {
				if !(validCh(thisCell) && thisCell == rowStart) {
					rowWinner = false
				}
			}
		}

		if rowWinner {
			log.Printf("checkWin(): winning row\n")
			return true
		}
	}

	if colWin(colWinner) {
		log.Printf("checkWin(): winning column\n")
		return true
	}

	return false
}

func colWin(colWin []bool) bool {
	for _, v := range colWin {
		if v == true {
			return true
		}
	}

	return false
}

func validCh(ch rune) bool {
	if ch == 'X' || ch == 'O' {
		return true
	}

	return false
}
