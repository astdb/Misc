// write code which computes the number of trailing zeroes in n-factorial.

package main

import (
	"fmt"
)

func main() {
	testcases := []int{-1, -2, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	// for k, v := range testcases {
	for _, v := range testcases {
		// nfac, err := factorial(v)

		// if err == nil {
		// 	fmt.Printf("Testcase %d\n\tn = %d\n\tn! = %d\n\tZeroes: %d\n", k, v, nfac, trailingZeroes(nfac))
		// } else {
		// 	fmt.Printf("Testcase %d\n\tn = %d\n\tn! = %v\n", k, v, err)
		// }
		fmt.Printf("%d! has %d trailing zeroes.\n", v, countFactorialZeroes(v))
	}
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("%d doens't have a defined factorial (negative value)", n)
	}

	if n == 0 {
		return 1, nil // 0! == 1 by definition
	}

	// return factorial(n-1)*n
	nfac, err := factorial(n - 1)

	if err == nil {
		return nfac * n, nil
	} else {
		return 0, fmt.Errorf("Error calculating factorial(%d)", n-1)
	}
}

// takes an integer and returns the number of trailing zeroes it has i.e. 101000 => 3
func trailingZeroes(n int) int {
	zeroes := 0
	rem := n % 10

	for rem == 0 {
		zeroes++
		n = n / 10
		rem = n % 10
	}

	return zeroes
}

// ------------- method without calculating n! ----------------------
func factorOfFive(i int) int {
	count := 0

	 for i % 5 == 0 {
		 count++
		 i /= 5
	 }

	 return count
}

func countFactorialZeroes(n int) int {
	count := 0

	if n < 0 {
		return -1
	}

	for i := 5; n/i > 0; i *= 5 {
		count += n/i
	}

	return count
}
