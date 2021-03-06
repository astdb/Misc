/*
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
Note:
The range of node's value is in the range of 32-bit signed integer.
*/

package main

import (
// "fmt"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
ALGORITHM
---------

 - Push root node onto a queue. Then, while the queue is not empty:
	- Remove a node from queue.
	- Push all its children into a temp queue.
	- Keep on popping nodes from queue and pushing their child nodes onto temp queue until queue is empty.
	- Each time the queue becomes empty, a level has been considered.
	- Keep track of a total and number of nodes while pushing onto temp queue, and calculate an average each time main queue becomes empty.
	- Reinitialize main queue with temp queue
*/

func averageOfLevels(root *TreeNode) []float64 {
	levelAverages := []float64{}
	mainQueue := []*TreeNode{root}

	for len(mainQueue) > 0 {
		sum := 0
		count := 0.0

		tempQueue := []*TreeNode{}

		for len(mainQueue) > 0 {
			n := mainQueue[0]
			mainQueue = mainQueue[1:]
			sum += n.Val
			count++

			if n.Left != nil {
				tempQueue = append(tempQueue, n.Left)
			}

			if n.Right != nil {
				tempQueue = append(tempQueue, n.Right)
			}
		}

		mainQueue = tempQueue
		levelAverages = append(levelAverages, float64(sum)/count)
	}

	return levelAverages
}
