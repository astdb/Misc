/*
For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.

 

Example 1:

Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
Example 2:

Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
Example 3:

Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
Example 4:

Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000
 

Note：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
If A.length > 1, then A[0] != 0

*/

package main

import (
	"log"
	"math"
)

func main() {
	// log.Println(int(math.Pow10(0)))

	// tests := [][]int{ {1,2,3}, {2,1}, {0} }
	// for _, test := range tests {
	// 	log.Printf("arrayToInt(%v) = %d\n", test, arrayToInt(test))
	// }

	// tests := []int{321, 0}
	// for _, test := range tests {
	// 	log.Printf("intToArray(%v) = %d\n", test, intToArray(test))
	// }

	tests := [][][]int{ {{1,2,0,0},{34}}, {{2,7,4}, {181}}, {{2,1,5}, {806}}, {{9,9,9,9,9,9,9,9,9,9},{1}}, {{0},{0}} }

	for _, test := range tests {
		// log.Printf("addToArrayForm(%v, %d) = %v\n", test[0], test[1][0], addToArrayForm(test[0], test[1][0]))
		log.Printf("addToArrayForm(%v, %d) = %v\n", test[0], test[1][0], addToArrayForm2(test[0], test[1][0]))
	}
}

func addToArrayForm2(A []int, K int) []int {
	i := 1

	x := true
	y := true

	K_arr := intToArray(K)

	var res int
	var rem int
	resArr := []int{}

	for x || y {
		x1 := 0
		x2 := 0

		if (len(A)-i) >= 0 {
			x1 = A[len(A)-i]
		} else {
			x = false
		}

		if (len(K_arr)-i) >= 0 {
			x2 = K_arr[len(K_arr)-i]
		} else {
			y = false
		}

		i++

		res = x1 + x2 + rem		
		if res >= 10 {
			rem = res % 10
			res = res / 10

		} else {
			rem = 0
		}

		resArr = append(resArr, res)
	}

	if rem > 0 {
		resArr = append(resArr, rem)
	}

	return reverse(resArr)
}



func addToArrayForm(A []int, K int) []int {
	
	// return A
	return intToArray(arrayToInt(A) + K)
}

func reverse(x []int) []int {
	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	}

	return x
}

func intToArray(x int) []int {
	res := []int{}

	var rem int

	if x == 0 {
		return []int{0}
	}
	
	for x  > 0 {
		 rem = x % 10
		 x = x / 10

		res = append(res, rem)
	 }

	 return reverse(res)
}

// turn an integer given as an array of its decimal digits to int form
func arrayToInt(x []int) int {
	res := 0

	for i := len(x) - 1; i >= 0; i-- {
		// log.Println("--------------------------")
		// log.Println(math.Pow10(x[len(x)-i-1]))
		// log.Println(res)
		// log.Println(x[i])
		res += x[i] * int(math.Pow10(len(x)-i-1))
		// log.Println(res)
	}

	return res
}
