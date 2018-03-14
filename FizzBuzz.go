
/*
According to http://wiki.c2.com/?FizzBuzzTest: The "Fizz-Buzz test" is an interview question designed to help filter out the 99.5% of programming job candidates who "can't seem to program their way out of a wet paper bag". The text of the programming assignment is as follows:

Write a program that prints the numbers from 1 to 100. But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three and five print “FizzBuzz”.
*/

package main

import (
	"fmt"
)

func main() {
	// iterate through numbers 1-100 (inclusive)
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			// multiple of 3
			fmt.Print("Fizz")
		}

		if i % 5 == 0 {
			// multiple of 5
			fmt.Print("Buzz")
		}

		if i % 3 != 0 && i % 5 != 0 {
			// not a multiple of 3 or 5
			fmt.Print(i)
		}

		// go to next line
		fmt.Print("\n")
	}
}
