/*
Write functions that
take the string version of an int and returns the int value
take tan integer and returns the string version of it
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	tests_int := []int{0, -1, 300, 654985, +65}
	tests_str := []string{"0", "123", "-3216851", "+54658"}

	for _, test := range tests_int {
		res, err := Itoa(test)
		if err != nil {
			fmt.Printf("Itoa(%d) = %v\n", test, err)
		} else {
			fmt.Printf("Itoa(%d) = %s\n", test, res)
		}
	}

	fmt.Println("\n")

	for _, test := range tests_str {
		res, err := Atoi(test)
		if err != nil {
			fmt.Printf("Atoi(%s) = %v\n", test, err)
		} else {
			fmt.Printf("Atoi(%s) = %d\n", test, res)
		}
	}

}

// returns a string representation of a given base 10 integer
func Itoa(x int) (string, error) {
	nums := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
	digits := []int{}

	neg := false
	if x < 0 {
		x = x * (-1)
		neg = true
	}

	if x == 0 {
		return nums[x], nil
	}

	var rem int
	for x > 0 {
		rem = x % 10
		x = x / 10
		digits = append(digits, rem)
	}

	var str strings.Builder

	if neg {
		str.WriteString("-")
	}

	for i := len(digits) - 1; i >= 0; i-- {
		i, valid := nums[digits[i]]
		if !valid {
			return "", errors.New("Invalid input int.")
		}

		str.WriteString(i)
	}

	return str.String(), nil
}

// returns base 10 numeric version of a given string representing a valid number
func Atoi(s string) (int, error) {
	s = strings.TrimSpace(s)

	nums := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

	s_runes := []rune(s)
	tot := 0
	pow := 0
	neg := false
	for i := len(s_runes) - 1; i >= 0; i-- {
		// for i := 0; i < len(s_runes); i++ {
		if i == 0 {
			if s_runes[i] == '-' {
				neg = true
				break
			}

			if s_runes[i] == '+' {
				break
			}

		}

		n, valid := nums[s_runes[i]]
		if !valid {
			return 0, errors.New("Invalid numeric string.")
		}

		tot += n * int(math.Pow10(pow))
		pow++
	}

	if neg {
		return ((-1) * tot), nil
	}

	return tot, nil
}
