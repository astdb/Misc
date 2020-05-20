/*
Given scores of N athletes, find their relative ranks and the people with the top three highest scores, who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

Example 1:
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal". 
For the left two athletes, you just need to output their relative ranks according to their scores.
Note:
N is a positive integer and won't exceed 10,000.
All the scores of athletes are guaranteed to be unique.
*/

package main

import (
    "log"
    "sort"
    "strconv"
)

func main() {
    tests := [][]int{{5,4,3,2,1}}

    for _, test := range tests {
        log.Printf("findRelativeRanks(%v) == %v\n", test, findRelativeRanks(test))
    }
}

func findRelativeRanks(nums []int) []string {
    sort.Ints(nums)
    res := []string{}

    for i, v := range nums {
        if i == 0 {
            res = append(res, "Gold Medal")
        }

        if i == 1 {
            res = append(res, "Silver Medal")
        }

        if i == 2 {
            res = append(res, "Bronze Medal")
        }

        if i > 2 {
            res = append(res, strconv.Itoa(v))
        }
    }

    return res
}
