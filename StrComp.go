/*
Implement a method of basic string compression by replacing repeated characters by their counts. e.g. aabcccccaaa would become a2b1c5a3.
If the compressed string doesn't become shorter, return the original string.

*/

package main

import (
	"log"
	"strconv"
)

func main() {
	tests := []string{"", "a", "aa", "aaa",  "aabcccccaaa"}

	for _, test := range tests {
		log.Printf("compress(%s) = %s\n", test, compress(test))
	}
}

func compress(s string) string {
	// turn s into a char slice
	sRunes := []rune(s)

	res := []rune{}

	var curCh rune
	var curChCount int

	for i := 0; i < len(sRunes); i++ {
		if i == 0 {
			curCh = sRunes[i]
			curChCount = 1

		} else {
			if sRunes[i] == curCh {
				curChCount++
			} else {
				res = append(res, curCh)
				res = append(res, []rune(strconv.Itoa(curChCount))...)

				curCh = sRunes[i]
				curChCount = 1
			}
		}
	}

  res = append(res, curCh)
  res = append(res, []rune(strconv.Itoa(curChCount))...)

	if len(res) >= len(sRunes) {
		return s
	}

	return string(res)
}
