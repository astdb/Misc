// Design an algorithm to decide if someone has won a game of tic tac toe

// ALGORITHM
// ---------
// # assuming a 3x3 board
// initialize board
// initialize coordinate lists and mark counts for both players

// for each
//	place mark
// 	third mark for this player?
// 		yes: check for consecutive row/diag. from last mark for this player
// 			win: declare win and exit
// 			no: other player has three marks?
// 				yes: exit - draw
// 				no: add mark to player's coord list and inc. their mark count: continue to next mark (from other player)

// 		no: add mark to player's coord list and inc. their mark count: continue to next mark (from other player)

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Tic-Tac-Toe...")

	_game := makeBoard(3)

	_game.markX(1,0)
	_game.print()

	_game.markY(1,1)
	_game.print()

	_game.markX(0,2)
	_game.print()

	_game.markY(0,0)
	_game.print()

	_game.markX(2,1)
	_game.print()

	_game.markY(2,2)
	_game.print()

	// game_over := false
	// i := 0	// simple alternatig flag indicating if an X or O should be placed
	// for {
	// 	// play Tic tac toe
	// 	if game_over {
	// 		break
	// 	}

	// 	if i == 0 {
	// 		// place X
	// 		_game.markX(1,0)
	// 		i = 1
	// 	} else {
	// 		// place Y
	// 		_game.markY(1,1)
	// 		i = 0
	// 	}

	// }
}

// Struct representing Tic-tac-toe board
type ttt_board struct {
	size   int     // size x size board will be initialized
	xcount int     // number of X's placed on board
	ycount int     // number of O's placed on board
	lastX  int     // coordinate of the last X drawn
	lastY  int     // coordinate of the last Y drawn
	board  [][]int // slice of int slices holding board
}

// produce blank N x N board
func makeBoard(n int) *ttt_board {
	var B ttt_board
	B.size = n
	B.xcount = 0
	B.ycount = 0

	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}

		// B.board[i] = row
		B.board = append(B.board, row)
	}

	return &B
}

// mark X on board at given coordinate (0-indexed, x's are denoted by 1's)
func (b *ttt_board) markX(loc_x, loc_y int) {
	if loc_x < b.size && loc_y < b.size && b.board[loc_x][loc_y] == 0 {
		b.board[loc_x][loc_y] = 1
		b.lastX = loc_x
		b.lastY = loc_y
		b.xcount++
	}
}

// mark O on board at given coordinate (0-indexed, o's are denoted by 2's)
func (b *ttt_board) markY(loc_x, loc_y int) {
	if loc_x < b.size && loc_y < b.size && b.board[loc_x][loc_y] == 0 {
		b.board[loc_x][loc_y] = 2
		b.lastX = loc_x
		b.lastY = loc_y
		b.ycount++
	}
}

// print board to console
func (b *ttt_board) print() {
	for _, row := range b.board {
		for _, col := range row {
			// fmt.Printf("  %d", col)
			if col == 0 {
				fmt.Printf(" _")
			}

			if col == 1 {
				fmt.Printf(" O")
			}

			if col == 2 {
				fmt.Printf(" X")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	return
}

// determine winner
func (b *ttt_board) ttt_winner(loc_x, loc_y int) bool {
	// when the maximum number of X's or O's are placed, this function will be called
	// with the coordinates of the last placed X or O to determine if that was a winning
	// mark.
	if loc_x < b.size && loc_y < b.size {
		mark := b.board[loc_x][loc_y] // could be an X or Y

		// this is a winner if the whole row, column or diagonal (if the coordinate is on a diagonal) consist of mark's

		// check row
		all_marks := true
		for _, v := range b.board[loc_x] {
			if v != mark {
				all_marks = false
			}
		}

		if all_marks {
			// we had a row of all X's or Y's
			return true
		}

		// check column
		all_marks = true
		for _, v := range b.board {
			if v[loc_y] != mark {
				all_marks = false
			}
		}

		if all_marks {
			// we had a column of all X's or Y's
			return true
		}

		// check diagonal (if the coordinate is on a diagonal)
		if loc_x == loc_y {
			// point is on a diagonal
			x, y := 0, 0
			all_marks = true
			for {
				if x >= b.size || y >= b.size {
					break
				}

				if b.board[x][y] != mark {
					all_marks = false
				}

				x++
				y++
			}

			if all_marks {
				// we had a column of all X's or Y's
				return true
			}

			x, y = 0, b.size-1
			all_marks = true
			for {
				if x >= b.size || y < 0 {
					break
				}

				if b.board[x][y] != mark {
					all_marks = false
				}

				x++
				y--
			}

			if all_marks {
				// we had a column of all X's or Y's
				return true
			}
		}
	}

	return false
}
