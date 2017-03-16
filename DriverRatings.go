// simple program to simulate how driver ratings would fair with number of trips

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numTrips := 100 // number of trips
	minScore := 4   // minimum rating this driver would be given
	totalScore := 0 // rating total

	for i := 0; i <= numTrips; i++ {
		totalScore += getRandomScore(minScore)
	}

	avg := float64(totalScore) / float64(numTrips)
	fmt.Printf("Trips: %d\nAvg. rating: %.2f\n", numTrips, avg)
}

func getRandomScore(min int) int {
	score := rand.Intn(6)

	for score < min {
		score = rand.Intn(6)
	}

	return score
}
