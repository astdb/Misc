/*
Find all x prime numbers from 1 to 50,000. A prime number is a number which does not have any integer factors other than 1 and itself. 1 is not a prime number by definition.

01. Print out number of all primes upto 50,000 without optimization.
02. divide i by numbers only upto sqrt(i)
03. divide i only by prime factors

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(primes(50000))
	// fmt.Println(primesSqrt(50000))
	fmt.Println(primesPrimes(50000))
}

func primes(limit int) int {
	primeCount := 0
	for i := 2; i <= limit; i++ {
		// assume i is prime
		prime := true

		// try dividing i by all valid numbers less i
		for j := 2; j < i; j++ {
			if i % j == 0 {
				// found a factor for i - i is not prime
				prime = false
				break
			}
		}

		if prime {
			primeCount++
		}
	}

	return primeCount
}

func primesSqrt(limit int) int {
	primeCount := 0
	for i := 2; i <= limit; i++ {
		// assume i is prime
		prime := true

		// try dividing i by all valid numbers less than square root of i
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i % j == 0 {
				// found a factor for i - i is not prime
				prime = false
				break
			}
		}

		if prime {
			primeCount++
		}
	}

	return primeCount
}

func primesPrimes(limit int) int {
	primeCount := 0
	primeList := []int{}

	for i := 2; i <= limit; i++ {
		// assume i is prime
		prime := true

		// try dividing i by all valid primes less than i
		for _, v := range primeList {
			if i % v == 0 {
				// found a factor for i - i is not prime
				prime = false
				break
			}
		}

		if prime {
			primeCount++
			primeList = append(primeList, i)
		}
	}

	return primeCount
}
