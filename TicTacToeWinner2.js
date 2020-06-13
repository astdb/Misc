/*
Tic-tac-toe is played by two players A and B on a 3 x 3 grid.

Here are the rules of Tic-Tac-Toe:

Players take turns placing characters into empty squares (" ").
The first player A always places "X" characters, while the second player B always places "O" characters.
"X" and "O" characters are always placed into empty squares, never on filled ones.
The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Given an array moves where each element is another array of size 2 corresponding to the row and column of the grid where they mark their respective character in the order in which A and B play.

Return the winner of the game if it exists (A or B), in case the game ends in a draw return "Draw", if there are still movements to play return "Pending".

You can assume that moves is valid (It follows the rules of Tic-Tac-Toe), the grid is initially empty and A will play first.

 

Example 1:

Input: moves = [[0,0],[2,0],[1,1],[2,1],[2,2]]
Output: "A"
Explanation: "A" wins, he always plays first.
"X  "    "X  "    "X  "    "X  "    "X  "
"   " -> "   " -> " X " -> " X " -> " X "
"   "    "O  "    "O  "    "OO "    "OOX"
Example 2:

Input: moves = [[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]]
Output: "B"
Explanation: "B" wins.
"X  "    "X  "    "XX "    "XXO"    "XXO"    "XXO"
"   " -> " O " -> " O " -> " O " -> "XO " -> "XO " 
"   "    "   "    "   "    "   "    "   "    "O  "
Example 3:

Input: moves = [[0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2]]
Output: "Draw"
Explanation: The game ends in a draw since there are no moves to make.
"XXO"
"OOX"
"XOX"
Example 4:

Input: moves = [[0,0],[1,1]]
Output: "Pending"
Explanation: The game has not finished yet.
"X  "
" O "
"   "
 

Constraints:

1 <= moves.length <= 9
moves[i].length == 2
0 <= moves[i][j] <= 2
There are no repeated elements on moves.
moves follow the rules of tic tac toe.
*/

'use strict';

let tests = [[[0,0],[2,0],[1,1],[2,1],[2,2]], [[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]], [[0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2]], [[0,0],[1,1]]]

for(let i = 0; i < tests.length; i++) {
  let test = tests[i];
  console.log("tictactoe(" + test + ") == " + tictactoe(test));
}

/**
 * @param {number[][]} moves
 * @return {string}
 */
var tictactoe = function(moves) {
  // initialize board
  let board = [['', '', ''], ['', '', ''], ['', '', '']];

  let moves = 9; // number of moves remaining

  let xMove = true; // flag denoting if current move is X or O

  for(let i = 0; i < moves.length; i++) {
    let filler = "";

    if(xMove) {
      filler = "X";
    } else {
      filler = "O";
    }

    // make move
    board[moves[0]][moves[1]] = filler;
    moves--;

    if(won(board, moves, xMove)) {
      if(xMove) {
        return "A";
      }

      return "B";
    }
  }

  if(moves == 0) {
    return "Draw";
  }

  return "Pending";
};

var won = function(board, moves, xMoves) {
  // if(board[0][0] == board[0][1] == board[0][2]) {
  //   // row 1

  // } else if(board[1][0] == board[1][1] == board[1][2]) {
  //   // row 2

  // } else if(board[2][0] == board[2][1] == board[2][2]) {
  //   // row 3
    
  // } else if(board[0][0] == board[1][0] == board[2][0]) {
  //   // col 1

  // } else if(board[0][1] == board[2][1] == board[2][2]) {

  // } 

  let boardSize = 3;

  let colStart = '';
  let colSame = false;

  let diag1Start = '';
  let diag1Same = true;

  for(let col = 0; col < boardSize; col++) {

    let rowStart = '';  // starting character of this row
    let rowSame = true; // flag indicating if row contains same char (X or O)

    for(let row = 0; row < boardSize; row++) {
      if(col === 0 && row === 0) {
        colStart = board[col][row];
        diag1Start = board[col][row];
        
      } else if(row === 0) {
        if(colStart !== board[col][row]) {
          colSame = fale;
        }
      }


      if(row == 0) {
        rowStart = board[col][row];
      } else {
        if(rowStart !== board[col][row]) {
          rowSame = false;
        }
      }
    }

    if(rowSame === true) {
      return true;
    }
  }
  
  if(colSame === true) {
    return true;
  }
}
