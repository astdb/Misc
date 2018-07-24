/*

Given an integer array with even length, where different numbers in this array represent different kinds of candies. Each number means one candy of the corresponding kind. You need to distribute these candies equally in number to brother and sister. Return the maximum number of kinds of candies the sister could gain.
Example 1:
Input: candies = [1,1,2,2,3,3]
Output: 3
Explanation:
There are three different kinds of candies (1, 2 and 3), and two candies for each kind.
Optimal distribution: The sister has candies [1,2,3] and the brother has candies [1,2,3], too. 
The sister has three different kinds of candies. 
Example 2:
Input: candies = [1,1,2,3]
Output: 2
Explanation: For example, the sister has candies [2,3] and the brother has candies [1,1]. 
The sister has two different kinds of candies, the brother has only one kind of candies. 
Note:

The length of the given array is in range [2, 10,000], and will be even.
The number in given array is in range [-100,000, 100,000].

*/

package main

import (
	"fmt"
)

func main () {
	fmt.Println(distributeCandies([]int{1,1,2,2,3,3}))
	fmt.Println(distributeCandies([]int{1,1,2,3}))
	fmt.Println(distributeCandies([]int{1,1,1,1,2,2,2,3,3,3}))
}

func distributeCandies(candies []int) int {
	// initializae map to count unique types of candy
	candyTypes := map[int]int{}
	
	// for each candy, record its type (if not recorded already)
	for _, candy := range candies {
		_, counted := candyTypes[candy]

		if !counted {
			candyTypes[candy] = 1
		}
	}

	// if number of types >= number of candies, max. no of unique types one can get is half the no. of candies
	if len(candyTypes) >= (len(candies)/2) {
		return (len(candies)/2)
	}

	// if number of types < number of candies, max. no of unique types one can get is no. of types
	return (len(candyTypes))
}
