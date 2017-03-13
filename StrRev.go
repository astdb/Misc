package main

import (
    "fmt"
)

func main() {
	testCases := []string{"ABC", "abc", "s_runes_len := len(s_runes)", "Package", " ", "        ", "", "\n", "Nirmal Jayasinghe", "Rob Pike", "Hello, 世界!"}

	for k, v := range testCases {
		fmt.Printf("Testcase No. %d: \n\tOriginal: %s\n\tReversed: %s\n", k, v, revereseString2(v))
	}
}

// return a reversed version of the input string
func reverseString1(s string) string {
	s_new := ""	// create new string to hold reveresed version (go strings are immutable)
	s_runes := []rune(s)	// translate string into a slice of Unicode characters

	for i := len(s_runes)-1; i >= 0; i-- {
		s_new += string(s_runes[i])	// append characters from  the end of the unicode slice to the new string
	}

	return s_new	// return reveresed string
}

// return a reversed version of the input string
func reverseString2(s string) string {
	s_runes := []rune(s)	// translate string into a slice of Unicode characters

	j := 0
	for i := len(s_runes)-1; i >= 0; i-- {
		// s_runes[i], s_runes[j] = s_runes[j], s_runes[i]
		temp := s_runes[i]
		s_runes[i] = s_runes[j]
		s_runes[j] = temp
		j++
	}

	return string(s_runes)	// return reversed string
}
