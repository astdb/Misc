// uniform pseudorandom-number generator within a given range

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println("urd(1.6) == ", urd(1, 6))
	}
}

func urd(x, y int) int {
	// number of flips required
	flipCount := int(math.Pow(2.0, math.Ceil(math.Log2(float64(y-x+1)))))

	// flip n number of times, record outcomes
	flips := []int{}
	for i := 0; i < flipCount; i++ {
		flips = append(flips, rand.Intn(2))
	}

	// is the outcome within range
	for getDec(flips) > (y - x + 1) {
		flips = []int{}
		for i := 0; i < flipCount; i++ {
			flips = append(flips, rand.Intn(2))
		}
	}

	return (x - 1) + getDec(flips)
}

// get decimal value for a binary slice
func getDec(x []int) int {
	tot := 0
	for i := 0; i < len(x); i++ {
		tot += int(math.Pow(2.0, float64(i)))
	}

	return tot
}
