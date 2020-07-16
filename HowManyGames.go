package main

import (
	"log"
)

func main() {
	//   var tests []int32
	tests := [][]int32{{20, 3, 6, 80}, {20, 3, 6, 85}}

	for testNo, test := range tests {
		log.Printf("Test #%d: howManyGames(%v) == %d\n", testNo, test, howManyGames(test[0], test[1], test[2], test[3]))
	}
}

// Complete the howManyGames function below.
func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var gameCount int32
	gameCount = 0

	var prevPrice int32 // price of last game bought
	// prevPrice = p

	// buy first game if wallet > starter price. else return 0 (gameCount initial value)
	if s >= p {
		// log.Printf("Buying first game for $%d\n", p)
		gameCount++
		s = s - p
		prevPrice = p
	} else {
		return gameCount
	}

	for s > 0 {
		thisGamePrice := prevPrice - d
		if thisGamePrice <= m {
			thisGamePrice = m
		}

		// log.Printf("Buying next game for $%d (prev game bought for $%d)\n", thisGamePrice, prevPrice)

		s = s - thisGamePrice
		prevPrice = thisGamePrice

		if s >= 0 {
			gameCount++
		}
		// gameCount++
	}

	return gameCount

}
