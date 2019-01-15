/*
X is a good number if after rotating each digit individually by 180 degrees, we get a valid number that is different from X.  Each digit must be rotated - we cannot choose to leave it alone.

A number is valid if each digit remains a digit after rotation. 0, 1, and 8 rotate to themselves; 2 and 5 rotate to each other; 6 and 9 rotate to each other, and the rest of the numbers do not rotate to any other number and become invalid.

Now given a positive number N, how many numbers X from 1 to N are good?

Example:
Input: 10
Output: 4
Explanation:
There are four good numbers in the range [1, 10] : 2, 5, 6, 9.
Note that 1 and 10 are not good numbers, since they remain unchanged after rotating.
Note:

N  will be in range [1, 10000].
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: $> go run RotateDigits.go <input>")
		return
	}
	in, _ := strconv.Atoi(os.Args[1])
	fmt.Println(rotatedDigits(in))
}

func rotatedDigits(N int) int {
	goodNums := 0
	for i := 1; i <= N; i++ {
		if good(i) {
			// fmt.Printf("%d is good\n", i)
			goodNums++
		}
	}

	return goodNums
}

func good(n int) bool {
	nDigits := getDigits(n)
	nDigitsRotated := []int{}
	fmt.Printf("Digits of %d: %v\n", n, nDigits)

	for _, d := range nDigits {
		if rotatable(d) {
			nDigitsRotated = append(nDigitsRotated, rotate(d))
		} else {
			return false
		}
	}

	fmt.Printf("Digits of %d rotated: %v\n", n, nDigitsRotated)
	nDigitsRotatedNum := getNum(nDigitsRotated)
	fmt.Printf("Digits of %d rotated (numeric): %d\n\n", n, nDigitsRotatedNum)

	if n != nDigitsRotatedNum {
		fmt.Printf("%d != %d\n", n, nDigitsRotatedNum)
		return true
	}

	return false
}

func getNum(n []int) int {
	pow := 0
	tot := 0
	for _, i := range n {
		tot += i * int(math.Pow10(pow))
		pow++
	}

	return tot
}

func getDigits(n int) []int {
	digits := []int{}
	// rem := n / 10
	// n = n / 10
	// digits = append(digits,rem)
	var rem int
	for n > 0 {
		rem = n % 10
		n = n / 10
		digits = append(digits, rem)
	}

	return digits
}

func rotatable(n int) bool {
	if n == 3 || n == 4 || n == 7 || n > 9 {
		return false
	}

	return true
}

func rotate(n int) int {
	rotation := map[int]int{}
	rotation[0] = 0
	rotation[1] = 1
	rotation[8] = 8
	rotation[2] = 5
	rotation[5] = 2
	rotation[6] = 9
	rotation[9] = 6

	return rotation[n]
}
