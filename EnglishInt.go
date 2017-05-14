// given an integer, print an English phrase describing that number e.g. "One thousand", "Two hundred thirty four"

/*
Synopsis
--------
We consider integer numbers of base ten in this program.

For numbers from 0 through to 19, each one takes a unique English form (e.g. zero, one, two, ..., ten, eleven, ... etc). 

For numbers 20 and up, the English name can be built from the individual names of the number's component digits and their placements. 
For example, all numbers where 10^1th position is 2, the English name would include <prefix>twenty<postfix>. Note that prefix and
postfix could be zero-length i.e. 20 would simply turn into 'twenty'.

A mapping needs to be created which would return the correct English name component based on a digit and its placement in the number we're
trying to anglicanize. In functional form, this would be a routine that takes in a digit and it's position in the input number, and returns a 
strng fragment that would slot into the appropriate position of the output English name. 
*/

package main

import (
	"fmt"
	"math/rand"
)

var smalls = []string{"ZERO", "ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE", "TEN", "ELEVEN", "TWELVE", "THRITEEN", "FOURTEEN", "FIFTEEN", "SIXTEEN", "SEVENTEEN", "EIGHTEEN", "NINETEEN"}
var tens = []string{"", "", "TWENTY", "THIRTY", "FORTY", "FIFTY", "SIXTY", "SEVENTY", "EIGHTY", "NINETY"}
var bigs = []string{"", "THOUSAND", "MILLION", "BILLION", "TRILLION"}
var hundred = "HUNDRED"
var negative = "NEGATIVE"

func main() {
	for i:= 0; i < 20; i++ {
		n := rand.Intn(10000000)
		fmt.Printf("%d:\t%s\n", n, convert(n))
	}
}

func convert(num int) string {
	if num == 0 {
		return smalls[num]
	} else if num < 0 {
		return negative + " " + convert(-1 * num)
	}

	// parts := []string{}
	parts := ""
	chunkCount := 0

	for num > 0 {
		if num % 1000 != 0 {
			chunk := convertChunk(num % 1000) + " " + bigs[chunkCount]
			parts = chunk + " " + parts
		}

		num /= 1000
		chunkCount++
	}

	return parts
}

func convertChunk(number int) string {
	parts := ""

	// convert hundreds place
	if number >= 100 {
		parts = parts + " " + smalls[number/100]
		parts = parts + " " + hundred
		number %= 100
	}

	// convert tens place
	if number >= 10 && number <= 19 {
		parts = parts + " " + smalls[number]
	} else if number >= 20 {
		parts = parts + " " + tens[number/10]
		number %= 10
	}

	// convert ones place
	if number >= 1 && number <= 9 {
		parts = parts + " " + smalls[number]
	}

	return parts
}
