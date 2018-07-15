/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3
Example 2:

Input: J = "z", S = "ZZ"
Output: 0
Note:

S and J will consist of letters and have length at most 50.
The characters in J are distinct.
*/

package main

import (
	"fmt"
	"os"
)

func main () {
	// read input strings
	if len(os.Args) < 3 {
		fmt.Println("Usage: $> go run JewelStones.go <jewelstring> <stonestring>")
		return
	}

	jewelStr := os.Args[1]
	stoneStr := os.Args[2]
	fmt.Println(numJewelsInStones(jewelStr, stoneStr))
}

func numJewelsInStones(J string, S string) int {
    // read in jewel/stone inputs as rune slices for easier manipulation
	jewels := []rune(J)
	stones := []rune(S)

	// check how many stones are jewels
	count := 0
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if stones[i] == jewels[j] {
				count++
			}
		}
	}

	return count
}
