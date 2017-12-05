
	// Given an array of integers, return indices of the two numbers such that they add up to a specific target.


package main

import (
	"fmt"
)

func main () {
	// input := []int{2, 7, 11, 15}
	// target := 9

	input := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(input,target))
}

func twoSum(nums []int, target int) []int {    
    nums_len := len(nums)
    for i := 0; i < nums_len; i++ {        
        for j := i+1; j < nums_len; j++ {            
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    
    return []int{-1,-1}
}
