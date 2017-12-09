
/*
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	tests := []string{"I", "XIV", "MMM", "CXLVII"}

	for _,v := range tests {
		fmt.Printf("%s -> %d\n", v, romanToInt(v))
	}
}

func romanToInt(s string) int {
   // trim any space from input and convert to all-uppercase rune slice
   s_run := []rune(strings.ToUpper(strings.TrimSpace(s)))

	// build map of values for roman numerals
	romanNums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// iterate through the slice representing the roman numeral and build up the integer value
	int_val := 0
	s_len := len(s_run)
	for i := 0; i < s_len; i++ {
		if i + 1 < s_len {
			if romanNums[s_run[i]] >= romanNums[s_run[i+1]] {
				// the numeral in ith position can be singly added to the int total
				int_val += romanNums[s_run[i]]
			} else {
				// i'th and i+1th positions make a single next value e.g. 'iv', 'ix', 'XL'
				int_val += (romanNums[s_run[i+1]] - romanNums[s_run[i]])
				i++
			}
		} else {
			int_val += romanNums[s_run[i]]
		}		
	}

	return int_val
}
