/*
Given an integer n, add a dot (".") as the thousands separator and return it in string format.



Example 1:

Input: n = 987
Output: "987"
Example 2:

Input: n = 1234
Output: "1.234"
Example 3:

Input: n = 123456789
Output: "123.456.789"
Example 4:

Input: n = 0
Output: "0"


Constraints:

0 <= n < 2^31

*/

package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	tests := []int{987, 1234, 123456789, 0}

	for _, test := range tests {
		// log.Printf("getDigits(%d) == %v\n", test, getDigits(test))
		log.Printf("thousandSeparator(%d) == %v\n", test, thousandSeparator(test))
	}

	/*arr := []int{1,2,3,4}
	  log.Println(arr)
	  for i := len(arr)-1; i >= 0; i-- {
	    remLen := i

	    log.Printf("Currently at element arr[%d] (%d), length of remaining array is %d\n", i, arr[i], remLen)
	  }*/
}

func thousandSeparator(n int) string {
	digs := getDigits(n)

	var res strings.Builder

	// digs contains n's decimal digits in reverse order - iterate from the end
	for i := len(digs) - 1; i >= 0; i-- {
		res.WriteString(strconv.Itoa(digs[i]))
		// log.Printf("thousandSeparator(): string(digs[%d]) = %s\n", i, string(digs[i]))
		// res.WriteString("i")

		// if the remaining length of the array to iterate is evenly divisable by 3, add a comma after the current element
		// remLen := len(digs)-(i+1)
		remLen := i
		if (remLen > 0) && ((remLen % 3) == 0) {
			res.WriteString(".")
		}
	}

	return res.String()
}

// return decimal digits of a given integer
func getDigits(n int) []int {
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
