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
	// tests := [][][]rune{{{'X','O','O'},{'O','X','O'},{'X','O','X'}}, {{'X', 'O', 'O'}, {'X', 'O', 'X'}, {'X', 'X', 'O'}}, {{'X', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'X', 'X'}}, {{'X', 'O', 'X'}, {'O', 'X', 'O'}, {'X', 'O', 'X'}}}

  tests := [][][]rune{{{'X','O','O'},{'O','X','O'},{'X','O','X'}}}

  // tests := [][][]rune{{{'X','X','O'},{'O','O','X'},{'O','O','X'}}}

	for _, test := range tests {
		log.Printf("chekcWin(%v) == %v\n", test, checkWin(test))
	}
}

func checkWin(board [][]rune) bool {
	colWinner := []bool{}

  var diag1Start rune
  var diag2Start rune
  diag1Winner := true
  prevRowD1 := 0
  diag2Winner := true
  diag1RowIndex := 0
  prevRowD2 := 0
  var diag2RowIndex int

	for i := 0; i < len(board); i++ {
		var rowStart rune
		// var colStart rune
		// colWinner := []bool{}
		rowWinner := true

		for j := 0; j < len(board[i]); j++ {
			thisCell := board[i][j]

      if i == 0 && j == 0 {
        log.Printf("checkWin(): Setting diag1Start to %c\n", thisCell)
        diag1Start = thisCell
        diag1RowIndex++
        prevRowD1 = i
      } else {
        if i == 1 && j == 1 {
          log.Printf("\tcheckWin(): diag1RowIndex = %d, thisCell = %c, diag1Start = %c, prevRowD1 = %d, i =  %d, j = %d\n", diag1RowIndex, thisCell, diag1Start, prevRowD1, i, j)
        }

        if j == diag1RowIndex && thisCell != diag1Start && i > prevRowD1 {
          log.Printf("checkWin(): diagonal(1) char mismatch detected\n")
          diag1Winner = false
          prevRowD1 = i
          diag1RowIndex++
        }

        // diag1RowIndex++
      }

      if i == 0 && j == len(board[i])-1 {
        diag2RowIndex = len(board[i])-1
        diag2Start = thisCell
        // log.Printf("checkWin(): Setting diag2Start to %c\n", thisCell)
        diag2RowIndex--
        prevRowD2 = i
      } else {
        if j == diag2RowIndex && thisCell != diag2Start && i > prevRowD2{
          // log.Printf("checkWin(): diagonal char mismatch detected\n")
          diag2Winner = false
          prevRowD2 = i
          diag2RowIndex--
        }

        // diag2RowIndex--
      }

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
			log.Printf("\tcheckWin(): winning row\n")
			return true
		}
	}

	if colWin(colWinner) {
		log.Printf("\tcheckWin(): winning column\n")
		return true
	}

  if diag1Winner {
    log.Printf("\tcheckWin(): winning diagonal (1)")
    return true
  }

  if diag2Winner {
    log.Printf("\tcheckWin(): winning diagonal (2)")
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
