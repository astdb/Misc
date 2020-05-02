/*
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
Accepted
*/

package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	tests := [][]string{{"12", "8"}, {"1", "9"}}

	for _, test := range tests {
		log.Printf("addStrings(%v, %v) == %v\n", test[0], test[1], addStrings(test[0], test[1]))
	}
}

func addStrings(num1 string, num2 string) string {
	num1_runes := []rune(num1)
	num2_runes := []rune(num2)

	carry := 0
	num1Index := len(num1_runes) - 1
	num2Index := len(num2_runes) - 1
	var strRes strings.Builder

	for num1Index >= 0 || num2Index >= 0 || carry != 0 {
		// if num1Index >= 0 && num2Index >= 0 {
		// 	num1Digit, err := strings.Atoi(string(num1_runes[num1Index]))
		// 	if err != nil {
		// 		log.Fatal("Invalid input.")
		// 	}
		// } else if num1Index >= 0 {

		// } else if num2Index >= 0 {

		// }

		num1Int := 0
		num2Int := 0
		var err error

		if num1Index >= 0 {
			num1Int, err = strconv.Atoi(string(num1_runes[num1Index]))
			if err != nil {
				log.Fatalf("Invalid int rune in num1: %q\n", num1_runes[num1Index])
			}

			num1Index--

		}

		if num2Index >= 0 {
			num2Int, err = strconv.Atoi(string(num2_runes[num2Index]))
			if err != nil {
				log.Fatalf("Invalid int rune in num2: %q\n", num2_runes[num2Index])
			}

			num2Index--
		}

		log.Printf("Adding: %d + %d + %d\n", num1Int, num2Int, carry)
		thisRes := num1Int + num2Int + carry
		if thisRes >= 10 {
			carry = thisRes / 10
			thisRes = thisRes % 10

		} else {
			carry = 0

		}

		// result = append(result, rune(thisRes))
		strRes.WriteString(strconv.Itoa(thisRes))
	}

	// return string(result)
	result := []rune(strRes.String())
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	return string(result)
}
