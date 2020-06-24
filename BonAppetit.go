
/*
Function Description

Complete the bonAppetit function in the editor below. It should print Bon Appetit if the bill is fairly split. Otherwise, it should print the integer amount of money that Brian owes Anna.

bonAppetit has the following parameter(s):

bill: an array of integers representing the cost of each item ordered
k: an integer representing the zero-based index of the item Anna doesn't eat
b: the amount of money that Anna contributed to the bill
*/

package main

import (
  "log"
)

func main() {
  bonAppetit([]int32{2,4,6}, 2, 3)
  bonAppetit([]int32{2,4,6}, 2, 6)
}

// Complete the bonAppetit function below.
func bonAppetit(bill []int32, k int32, b int32) {
  var total int32
  total = 0

  for key, val := range bill {
    if int32(key) != k {
      total += val
    }
  }

  var sharedHalf int32
  sharedHalf = total/2
  if sharedHalf == b {
    log.Printf("Bon Appetit")
    return
  }

  log.Printf("%d", (b-sharedHalf))
}
