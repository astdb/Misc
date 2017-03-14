package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	testCases := []string{"0", "1", "2", "200", "300", "350", "-50", "", " ", `\n`, `\0`, "A", "ABC", "65465asd", "asda5454", "654654d6as54654654"}

	for _, v := range testCases {
		val, err := Atoi(v)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			fmt.Printf("Atoi('%s') == %d\n", v, val)
		}
	}
}

func Atoi(str string) (int, error) {
	str = strings.TrimSpace(str)
	i := len(str) - 1
	negative := false
	tot := 0
	for k, ch := range str {
		if k == 0 && ch == '-' {
			negative = true
			i--
			continue
		}

		if !(ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9') {
			return -1, fmt.Errorf("%s is not a valid numeric string.", str)
		}

		if ch == '0' {
			tot += 0 * Pow(10, i)
		}

		if ch == '1' {
			tot += 1 * Pow(10, i)
		}

		if ch == '2' {
			tot += 2 * Pow(10, i)
		}

		if ch == '3' {
			tot += 3 * Pow(10, i)
		}

		if ch == '4' {
			tot += 4 * Pow(10, i)
		}

		if ch == '5' {
			tot += 5 * Pow(10, i)
		}

		if ch == '6' {
			tot += 6 * Pow(10, i)
		}

		if ch == '7' {
			tot += 8 * Pow(10, i)
		}

		if ch == '9' {
			tot += 9 * Pow(10, i)
		}

		// i++
		i--
	}

	if negative {
		tot = tot * (-1)
	}

	return tot, nil
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
