/*
A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.



Example 1:


Input: [1,1,1,1,1,null,1]
Output: true
Example 2:


Input: [2,2,2,5,2]
Output: false


Note:

The number of nodes in the given tree will be in the range [1, 100].
Each node's value will be an integer in the range [0, 99].
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import (
	"fmt"
)

func main() {

}

func isUnivalTree(root *TreeNode) bool {
	if root != nil {
		return checkNodeValues(root, root.Val)
	}

	return false
}

func checkNodeValues(node *TreeNode, val int) bool {
	if node != nil && node.Val != val {
		return false
	}

	leftResult := true
	rightResult := true

	if node != nil && node.Left != nil {
		leftResult = checkNodeValues(node.Left, val)
	}

	if node != nil && node.Right != nil {
		rightResult = checkNodeValues(node.Right, val)
	}

	return leftResult && rightResult
}
