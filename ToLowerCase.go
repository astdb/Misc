/*
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: $> go run ToLowerCase.go <inputstring>")
		return
	}

	fmt.Println(toLowerCase(os.Args[1]))
}

func toLowerCase(str string) string {
	/*// convert input to editable rune slice
	str_runes := []rune(str)

	// iterate through, replacing uppercase chars with lowercase ones
	for i := 0; i < len(str_runes); i++ {
		if str_runes[i] >= 65 || str_runes[i] <= 90 {
			// uppercase - replace with lowercase equivalent
			ch := str_runes[i]
			str_runes[i] = ch + 32
		}
	}

	return string(str_runes)*/

	// storage for lowercase string
	lower := []rune{}

	// parse string, per rune
	for _, ch := range str {
		// check if ch is uppercase
		// fmt.Println(string(ch), " => ", ch)

		if ch >= 65 && ch <= 90 {
			// ASCII uppercase
			lower = append(lower, ch+32)
		} else {
			lower = append(lower, ch)
		}
	}

	return string(lower)
}
