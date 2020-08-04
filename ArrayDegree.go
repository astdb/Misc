/*
Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

 

Example 1:

Input: nums = [1,2,2,3,1]
Output: 2
Explanation: 
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.
Example 2:

Input: nums = [1,2,2,3,1,4,2]
Output: 6
Explanation: 
The degree is 3 because the element 2 is repeated 3 times.
So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.
 

Constraints:

nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.

*/

// find degree of given array
// find all subarrays, find degree of each, keep track of smallest length seen with prev array

func findShortestSubArray(nums []int) int {
   
}

// return the max frequency of any of arr's elements
func arrayDegree(arr []int) {

}

// getSubArrays() recursively computes all sub arrays of a given array, and stores results in result struct property (passed in by reference)
func getSubArrays(arr []int32, start, end int, res *Res) {
	if end >= len(arr) {
		return

	} else if start > end {
		getSubArrays(arr, 0, end + 1, res)

	} else {
		thisSubArr := []int32{}

		for i := start; i < end; i++ {
			thisSubArr = append(thisSubArr, arr[i])
		}

		thisSubArr = append(thisSubArr, arr[end])
		res.Result = append(res.Result, thisSubArr)

		getSubArrays(arr, start + 1, end, res)
	}
}

// Res provides a wrapper for the result storing all subarrays, to be passed by reference to subarray computing function
type Res struct {
	Result [][]int32
}

// NewRes() creates and returns a pointer to an instance of subarray results wrapper struct
func NewRes() *Res {
	var x Res
	return &x
}
