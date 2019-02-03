// Find the number of set bits in an int

package main

import (
        "fmt"
)

func main() {
        tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

        for _, test := range tests {
                fmt.Printf("countBits(%d) == %d\n", test, countBits(test))
        }
}

func countBits(x int) int {
        numBits := 0

        for x != 0 {
                numBits += (x & 1)
                x >>= 1
        }

        return numBits
}
