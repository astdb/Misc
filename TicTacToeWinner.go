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
}

// Tic-tac-toe board
type ttt_board struct {
	size   int
	xcount int
	ycount int
	lastX  int
	lastY  int
	board  [][]int
}

// produce blank nxn board
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

		B.board[i] = row
	}

	return &B
}

// mark x on board at given coordinate (0-indexed, x's are denoted by 1's)
func (b *ttt_board) markX(loc_x, loc_y int) {
	if loc_x < b.size && loc_y < b.size && b.board[loc_x][loc_y] == 0 {
		b.board[loc_x][loc_y] = 1
		b.lastX = loc_x
		b.lastY = loc_y
		b.xcount++
	}
}

// mark o on board at given coordinate (0-indexed, o's are denoted by 2's)
func (b *ttt_board) markY(loc_x, loc_y int) {
	if loc_x < b.size && loc_y < b.size && b.board[loc_x][loc_y] == 0 {
		b.board[loc_x][loc_y] = 2
		b.lastX = loc_x
		b.lastY = loc_y
		b.ycount++
	}
}

// determine winner
func (b *ttt_board) ttt_winner() int {

}
