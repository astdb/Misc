/*
Given an integer n. Each number from 1 to n is grouped according to the sum of its digits.

Return how many groups have the largest size.



Example 1:

Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.
Example 2:

Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.
Example 3:

Input: n = 15
Output: 6
Example 4:

Input: n = 24
Output: 5


Constraints:

1 <= n <= 10^4

*/

package main

import (
	"log"
)

func main() {
	tests := []int{13, 2, 15, 24}

	for _, test := range tests {
		log.Printf("countLargestGroup(%d) = %d\n", test, countLargestGroup(test))
	}

	for _, test := range tests {
		log.Printf("tot(%d) = %d\n", test, digitTotal(test))
	}
}

func countLargestGroup(n int) int {
	// mapping of totals to digit counts
	digTots := map[int]int{}

	for i := 1; i <= n; i++ {
		tot := digitTotal(i)

		_, counted := digTots[tot]
		if counted {
			digTots[tot]++
		} else {
			digTots[tot] = 1
		}
	}

	// find largest count
	var largestCount int

	i := 0
	for _, count := range digTots {
		if i == 0 {
			largestCount = count
		} else if count > largestCount {
			largestCount = count
		}
	}

	// count number of groups with largest count
	groupCount := 0

	for _, count := range digTots {
		if count == largestCount {
			groupCount++
		}
	}

	return groupCount

	// return -1
}

// return the decimal total of the decimal digits forming x
func digitTotal(x int) int {
	tot := 0

	var rem int

	for x > 0 {
		rem = x % 10
		x = x / 10

		tot = tot + rem
	}

	return tot
}
