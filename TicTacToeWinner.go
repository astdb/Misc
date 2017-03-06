
// Design an algorithm to decide if someone has won a game of tic tac toe

// ALGORITHM
// ---------
// # assuming a 3x3 board
// initialize board
// initialize coordinate lists and mark counts for both players

// for each mark placed
// 	third mark for this player?
// 		yes: check for consecutive row/diag. for this player
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


func ttt_winner(){

}
