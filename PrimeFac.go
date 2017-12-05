/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// num := 600851475143
	// num := 13195
	num := 40
	largestFac := 0

	for num % 2 == 0 {
		fmt.Printf("%d ", 2)
		num = num/2;
		largestFac = 2	
	}

	for i := 3; float64(i) < math.Sqrt(float64(num)); i += 2 {
		for num % i == 0 {
			if i > largestFac {
				largestFac = i
			}

			fmt.Printf("%d ", i)			
			num /= i
		}
	}

	if num > 2 {
		if num > largestFac {
			largestFac = num
		}
		fmt.Printf("%d ", num)
	}

	fmt.Printf("\nLargest factor: %d\n", largestFac)
}
