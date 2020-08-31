package main

import (
	"log"
	"unicode"
)

func main() {
	tests := []string{"oneTwoThree", "saveChangesInTheEditor"}
	for _, test := range tests {
		log.Printf("camelcase(%s) == %d\n", test, camelcase(test))
	}
}

func camelcase(s string) int32 {
	var res int32
	res = 0

	if len(s) > 0 {
		res++
	}

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			res++
		}
	}

	return res
}
