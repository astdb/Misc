// bruteforce method to find the size of the longest palindromic substring of a given string

package main

import (
	"fmt"
	"unicode"
)

func main() {
	TestStrings := []string{"ABBA", "BOBCAT", "ACYCLIC", "RACECAR"}

    for _, str := range TestStrings {
        print_all_substrings(str)
	    fmt.Println(longest_palindromic_substring(str))
    }
}

func longest_palindromic_substring(str string) int {
	var str_r []rune
	for _, r := range str {
		if unicode.IsLetter(r) {
			str_r = append(str_r, unicode.ToUpper(r))
		}
	}

	MaxPalLen := 0
	strrlen := len(str_r)
	for i := 1; i <= strrlen; i++ {
		// finding all substrings of i-length
		j := 0
		for k := j + i; k <= strrlen; j++ {
			substr := string(str_r[j:k])

			// check if palindrome
		    pal := true
			for m := range substr {
				if substr[m] != substr[len(substr)-1-m] {
					pal = false
				}
			}

			if pal == true {
				if len(substr) > MaxPalLen {
					MaxPalLen = len(substr)
				}
			}
			k++
		}
	}
	return MaxPalLen
}

//!-
