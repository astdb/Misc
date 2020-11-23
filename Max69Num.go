/*
https://leetcode.com/problems/maximum-69-number/
Given a positive integer num consisting only of digits 6 and 9.

Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).



Example 1:

Input: num = 9669
Output: 9969
Explanation:
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.
Example 2:

Input: num = 9996
Output: 9999
Explanation: Changing the last digit 6 to 9 results in the maximum number.
Example 3:

Input: num = 9999
Output: 9999
Explanation: It is better not to apply any change.


Constraints:

1 <= num <= 10^4
num's digits are 6 or 9.

*/

package main

import (
	"log"
	"math"
)

func main() {
	tests := []int{9669, 9996, 9999}

	for _, test := range tests {
		log.Printf("maximum69Number(%d) = %d\n", test, maximum69Number(test))
	}
}

// iterate the digits of num from left to right -flip the first most significant 6
func maximum69Number(num int) int {
	numDigits := getDigits(num)

	log.Printf("%v\n", numDigits)

	for i := len(numDigits) - 1; i >= 0; i-- {
		if numDigits[i] == 6 {
			// most significat 6 - flip to 9
			numDigits[i] = 9
			break
		}
	}

	return getVal(numDigits)

	// return 0
}

func getVal(digits []int) int {
	tot := 0
	// for i := len(digits)-1; i >= 0; i-- {
	for i := 0; i < len(digits); i++ {
		tot = tot + (digits[i] * int(math.Pow10(i)))
	}

	return tot
}

func getDigits(n int) []int {
	rem := n % 10
	n = n / 10
	res := []int{rem}

	for n > 0 {
		rem = n % 10
		n = n / 10

		res = append(res, rem)
	}

	return res
}
