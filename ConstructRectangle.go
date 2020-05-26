/*
For a web developer, it is very important to know how to design a web page's size. So, given a specific rectangular web page’s area, your job by now is to design a rectangular web page, whose length L and width W satisfy the following requirements:

1. The area of the rectangular web page you designed must equal to the given target area.

2. The width W should not be larger than the length L, which means L >= W.

3. The difference between length L and width W should be as small as possible.
You need to output the length L and the width W of the web page you designed in sequence.
Example:
Input: 4
Output: [2, 2]
Explanation: The target area is 4, and all the possible ways to construct it are [1,4], [2,2], [4,1].
But according to requirement 2, [1,4] is illegal; according to requirement 3,  [4,1] is not optimal compared to [2,2]. So the length L is 2, and the width W is 2.
Note:
The given area won't exceed 10,000,000 and is a positive integer
The web page's width and length you designed must be positive integers.
*/

package main

import (
	"log"
	"math"
)

func main() {
	tests := []int{0, 1, 2, 3, 4, 19, 20, 25, 1000000, 1000000000, 1000000000000}

	for _, test := range tests {
		log.Printf("constructRectangle(%d) == %v\n", test, constructRectangle(test))
	}
}

func constructRectangle(area int) []int {
	var curDiff int // diff of above

	var leng int
	var wid int

	// iterate through factor pairs for area, and find the couple with least difference
	if area == 1 {
		return []int{1,1}
	}

	for i := 1; i <= area/2; i++ {
		if area%i == 0 {
			fac1 := i
			fac2 := area / i

			if i == 1 {
				// initialize
				curDiff = absDiff(fac1, fac2)

				if fac1 > fac2 {
					leng = fac1
					wid = fac2
				} else {
					leng = fac2
					wid = fac1
				}
			} else if curDiff > absDiff(fac1, fac2) {
				curDiff = absDiff(fac1, fac2)

				if fac1 > fac2 {
					leng = fac1
					wid = fac2
				} else {
					leng = fac2
					wid = fac1
				}

			}
		}
	}

	return []int{leng, wid}
}

// return the absolute value of the difference between x and y
func absDiff(x, y int) int {
	res := x - y

	if res < 0 {
		return (res * (-1))
	}

	return res
}
