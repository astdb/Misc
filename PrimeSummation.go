/*
	The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
	Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	tot := 0 // prime total

	// start looking for primes from 2, until we've found a prime >= 2 000 000
	for j := 2; j < 10; j++ {
		if isPrime(j) {
			tot += j
		}
	}

	fmt.Println(tot)
}

func isPrime(j int) bool {
	if j == 2 || j == 3 {
		// 2 is prime by definition - only 'even' prime
		return true
	} else {
		// test j for divisability upto sqrt(j)
		intSqrtJ := int(math.Sqrt(float64(j)))
		for k := 2; k <= intSqrtJ; k++ {
			if j%k == 0 {
				// j isn't prime - goto next j
				return false
			}

			if k == intSqrtJ {
				return true
			}
		}
	}

	return false
}
