// check if an input integer's decimal form is a palindrome

package main

import (
    "fmt"
    "math"
)

func main() {
    tests := []int{456, 454, 498894, 2, 0, 1, 7, 11, 121, 333, 2147447412, -1, 12, 100, 2147483647}

    for _, test := range tests {
        fmt.Println(test, isPal(test))
        fmt.Println(test, palindrome(test))
    }
}

func palindrome(x int) bool {
    if x <= 0 {
        return x == 0
    }

    numDigits := int(math.Floor(math.Log10(float64(x)))) + 1
    msdMask := int(math.Pow10(numDigits - 1))

    for i := 0; i < (numDigits / 2); i++ {
        if x/msdMask != x%10 {
            return false
        }

        x %= msdMask // remove most significant digit of x
        x /= 10      // remove least significant digit of x
        msdMask /= 100
    }

    return true
}

func isPal(x int) bool {
    x_digits := []int{}
    var rem int

    if x < 0 {
        return false
    }

    for x > 0 {
        rem = x % 10
        x = x / 10
        x_digits = append(x_digits, rem)
    }

    startIndex := 0
    endIndex := len(x_digits) - 1

    for startIndex <= endIndex {
        if x_digits[startIndex] != x_digits[endIndex] {
            return false
        }

        startIndex++
        endIndex--
    }

    return true
}


