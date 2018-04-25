package main

import (
	"fmt"
)

func main() {
	a := "1010"
	b := "1011"
	fmt.Printf("addBinary(%s, %s) = %s\n", a, b, addBinary(a, b))
}

func addBinary(a string, b string) string {
	// transform strings representing binary numbers to rune slices for easy traversal
	// note: these strings only contain 1/0 - byte slices would've done too. 
	a_run := []rune(a)
	b_run := []rune(b)

	// initialize slice to store addition result
	result := []rune{}

	// for readability, shorter number will be iterated within the longer number (length per the binary digits in each)
	var first []rune	// longer number
	var second []rune	// shorter number
	if len(b_run) > len(a_run) {
		// b is longer
		first = b_run
		second = a_run
	} else {
		// a is longer or both of same length
		first = a_run
		second = b_run
	}

	rem := false	// flag indicating if a carry exists

	// result of adding together two current digits -  this will be known as a digital result.
	// e.g. adding the rightmost 1's of a = 111 and b = 111 will result in 0 as the current digital result (with a carry of 1, i.e. rem == 0)
	var res rune

	second_index := len(second) - 1	// index to iterate over the 'shorter' number

	// iterate through the longer number, from right end
	for i := len(first)-1; i >= 0; i-- {
		v1 := first[i]		// get the first digit of longer number (from right)
		var v2 rune

		// read corresponding digit from the shorter number, if exists. Update index. 
		if second_index >= 0 {
			v2 = second[second_index]
			second_index--
		} else {
			v2 = '0'
		}

		if v1 == '1' && v2 == '1' {
			if rem {
				// adding 1 to 1 with a remainder results in 11, therefore current digital result is 1, and we keep the carry
				res = '1'

			} else {
				// adding 1to 1 without a carry results in 10, therefore current digital result is 0, with carry 1 (rem == true)
				res = '0'
				rem = true
			}
		} else if v1 == '0' && v2 == '0' {
			if rem {
				// adding 0 with 0 with a carry gives current digital result 1, with no carry
				res = '1'
				rem = false

			} else {
				// adding 0 with 0 with no carry gives current digital result 0, with no carry
				res = '0'
			}
			
		} else {
			if rem {
				// adding 1 to 0 (or vice versa) with carry gives current digital result 0, with carry
				res = '0'
			} else {
				// adding 1 to 0 (or vice versa) with no carry gives current digital result 1, with no carry
				res = '1'
			}
		}
		
		// append the digital result to the overall result
		result = append(result, res)
	}

	if rem {
		// at the end of digital additions, check if there's a carry remaining and if so append it to the front of the result
		res = '1'
		result = append(result, res)
	}

	// result contains the final result backwards - reverse it, and return as a string
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return string(result)
}
