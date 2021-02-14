/*
Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number N on the chalkboard.  On each player's turn, that player makes a move consisting of:

Choosing any x with 0 < x < N and N % x == 0.
Replacing the number N on the chalkboard with N - x.
Also, if a player cannot make a move, they lose the game.

Return True if and only if Alice wins the game, assuming both players play optimally. 

Example 1:

Input: 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.

Example 2:

Input: 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
 

Note:

1 <= N <= 1000
*/

package main

import (
  "log"
)

func main() {
  tests := []int{2, 3}

  for _, test := range tests {
    log.Printf("divisorGame(%d) = %v\n", test, divisorGame(test))
  }
}

func divisorGame(n int) bool {
  // Alice's first move
  // divisorFound, divisor := getDiv(n, )

  log.Printf("N = %d\n", n)

  alice := true  // first turn Alice's
  for {
    if alice {
      log.Println("------------Alice's turn----------------")
    } else {
      log.Println("------------Bob's turn------------------")
    }

    validMove, x := getX(n)

    if validMove {
      log.Printf("Valid move - n (%d) is now set to x (%d)\n", n, x)
      n = x
    } else {
      log.Println("Could not make a valid move")
      if !alice {
        log.Println("Bob's turn (Alice won) - returning true")
        return true
      } else {
        log.Println("Alice's turn (Bob won) - returning false")
        return false
      }
    }

    if alice {
      alice = false
    } else {
      alice = true
    }
  }
}

func getX(n int) (bool, int) {
  x := 0
  valid := false
  
  for i := n-1; i > 0; i-- {
    if n % i == 0 {
      x = i
      valid = true

      return valid, x
    }
  }

  return valid, x
}
