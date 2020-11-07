/*
Given an integer number n, return the difference between the product of its digits and the sum of its digits.


Example 1:

Input: n = 234
Output: 15
Explanation:
Product of digits = 2 * 3 * 4 = 24
Sum of digits = 2 + 3 + 4 = 9
Result = 24 - 9 = 15

Example 2:

Input: n = 4421
Output: 21
Explanation:
Product of digits = 4 * 4 * 2 * 1 = 32
Sum of digits = 4 + 4 + 2 + 1 = 11
Result = 32 - 11 = 21


Constraints:

1 <= n <= 10^5

*/

package main

import (
	"log"
)

func main() {
	tests := []int{234, 4421}

	for _, test := range tests {
		log.Printf("subtractProductAndSum(%d) = %d\n", test, subtractProductAndSum(test))
	}
}

func subtractProductAndSum(n int) int {
	nDigits := getDecimalDigits(n)
  // log.Printf("subtractProdSum(): nDigits: %v\n", nDigits)

  // prod := getProduct(nDigits)
  // tot := getTotal(nDigits)

  // log.Printf("subtractProdSum(): tot: %d\n", tot)
  // log.Printf("subtractProdSum(): prod: %d\n", prod)

	return (getProduct(nDigits) - getTotal(nDigits))
}

func getDecimalDigits(n int) []int {
	res := []int{}

	rem := n % 10
	n = n / 10
	res = append(res, rem)

	for n > 0 {
		rem = n % 10
		n = n / 10

		res = append(res, rem)
	}

	return res
}

func getTotal(nums []int) int {
	var tot int

	for _, n := range nums {
		tot += n
	}

	return tot
}

func getProduct(nums []int) int {
	var prod int

  if len(nums) > 0 {
    prod = 1
  }

	for _, n := range nums {
		prod *= n
	}

	return prod
}
