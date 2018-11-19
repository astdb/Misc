
*/

package main

import (
        "fmt"
)

func main() {
        tests := [][]int{[]int{3, 2, 3}, []int{2, 2, 1, 1, 1, 2, 2}}

        for _, test := range tests {
                fmt.Println(majorityElement(test))
        }
}

func majorityElement(nums []int) int {
        maj := len(nums) / 2
        counts := map[int]int{}

        for _, d := range nums {
                // check if d is already in count
                _, exists := counts[d]
                if exists {
                        counts[d]++
                } else {
                        counts[d] = 1
                }

                if counts[d] > maj {
                        return d
                }
        }

        return 5
}
