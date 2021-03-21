/*
Given an array of characters with sufficient spare space at the end (and 'true length' of the string provided), convert
all spaces into %20.

e.g. input:  "Mr John F. Smith      "
     output: "Mr%20John%20F.%20Smith"
*/

package main

import (
	"log"
)

func main() {
	tests := []*TestCase{
		&TestCase{Str: []rune{'M', 'r', ' ', 'J', 'o', 'h', 'n', ' ', 'F', '.', ' ', 'S', 'm', 'i', 't', 'h', ' ', ' ', ' ', ' ', ' ', ' '}, StrLen: 16},
		&TestCase{Str: []rune{'A', 'B', 'C'}, StrLen: 3},
		&TestCase{Str: []rune{}, StrLen: 0},
	}

	for _, test := range tests {
		log.Printf("urlify(\"%s\", %d) = \"%s\"\n", string(test.Str), test.StrLen, string(urlify(test.Str, test.StrLen)))
	}
}

func urlify(str []rune, strlen int) []rune {
	// index of the last char of the true string
	els := strlen - 1

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			// - shift everything from i+1 to two positions to the right
			// - replace str[i] str[i+1] and str[i+2] with '%', '2', and '0'
			// - increment i by 2
			for j := els; j > i; j-- {
				str[j+2] = str[j]
			}

			// update end index of true string
			els += 2

			// replace whitespace with %20
			str[i] = '%'
			str[i+1] = '2'
			str[i+2] = '0'

			// update str iteration index
			i += 2
		}
	}

	return str
}

type TestCase struct {
	Str    []rune
	StrLen int
}
