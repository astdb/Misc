
// write a program to detect is a given string has all unique chars
// what if additional data structures cannot be used?

package main

import (
	"fmt"
	"os"
)

func main() {
	// read input string from command line (as a slice of unicode characters)
	if len(os.Args) < 2 {
		fmt.Println("Usage: $ go run UniqueChars.go teststring")
		return
	}
	// testString := []rune(os.Args[1])
	testString := os.Args[1]

	// for each character, from the start
	/* for i, ch := range testString {
		// check if it exists in rest of string
		for j := i + 1; j < len(testString); j++ {
			if ch == testString[j] {
				// i'th char found repeated at position j
				fmt.Println("NONUNIQUE")
				return;
			}
		}
	}

	fmt.Println("UNIQUE") */

	if isUnique(testString) {
		fmt.Println("UNIQUE")
	} else {
		fmt.Println("NONUNIQUE")
	}
}

// declare boolean slice of charset size (e.g. 256 for ext. ASCII)
// set that charposition to true if not already set: if already
// set, nonunique char found
func isUnique(str string) bool {
	// setup array
	CHARSET_SIZE := 256

	// if string length > charset size, must have repeating chars
	if len(str) > CHARSET_SIZE {
		return false
	}

	charPos := make([]bool, CHARSET_SIZE, CHARSET_SIZE)
	
	for _,ch := range str {
		ch_int := int(ch)

		if ch_int < CHARSET_SIZE {
			if charPos[ch_int] == true {
				// char seen before
				return false
			} else {
				// new char
				charPos[ch_int] = true
			}
		} else {
			// char outside charset found - panic?
			panic("isUnique: char outside charset found")
		}
	}

	return true
}
