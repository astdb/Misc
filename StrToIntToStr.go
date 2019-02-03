/*
Write functions that returns the int version of a string and a string version of an int
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	tests_str := []string{"0", "123", "-3216851", "+54658"}

	for
}

// returns a string representation of a given base 10 integer
func Atoi(x int) (string, error) {
	nums := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
	digits := []int{}

neg := false	
if x < 0 {
	neg = true
}

	var rem int
	for x > 0 {
		rem = x % 10
	x = x /10
	digits = append(digits, rem)
	}

	var str strings.Builder

	if neg {
		str.WriteString("-")
	}

	for i := len(digits)-1; i >= 0; i-- {
		i, valid := nums[digits[i]]
		if !valid {
			return nil, errors.New("Invalid input int.")
		}

		str.WriteString(i)
	}

	return str.String()
}

// returns base 10 numeric version of a given string representing a valid number
func Itoa(s string) (int,error) {
	s = strings.TrimSpace(s)
	
nums := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

	x_runes := []rune(x)
	tot := 0
	neg := false
	for i := len(x)-1; i >= 0; i-- {
		if i == 0 {
			if x_runes[i] == '-' {
				neg = true
break
}

if x_runes[i] == '+' {
				break
}

		} 
			n, valid := nums[x_runes[i]]
		if !valid {
			return nil, errors.New('Invalid numeric string.')
		}

		tot += n * int(math.Pow10(i))		
	}

	if neg {
		return ((-1) * tot), nil
	}

	return tot, nil
}
