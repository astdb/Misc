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
	for j := 2; j < 2000000; j++ {
		if j == 2 || j == 3 {
			// 2 is prime by definition - only 'even' prime
			tot += j
		} else {
			// test j for divisability upto sqrt(j)
			for k := 2; k <= int(math.Sqrt(float64(j))); k++ {
				if j%k == 0 {
					// j isn't prime - goto next j
					break
				}

				if k == int(math.Sqrt(float64(j))) {
					tot += j
				}
			}
		}
	}

	fmt.Println(tot)
}
