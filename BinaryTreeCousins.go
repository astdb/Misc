/*
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
Return true if and only if the nodes corresponding to the values x and y are cousins.

Example 1:
Input: root = [1,2,3,4], x = 4, y = 3
Output: false

Example 2:
Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true

Example 3:
Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false

Constraints:

The number of nodes in the tree will be between 2 and 100.
Each node has a unique integer value from 1 to 100.
*/

package main

import (
	"log"
)

func main() {

}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	// return (root.Left.Val == x && root.Right.Val == y) || (root.Left.Val == y && root.Right.Val == x) || isCousins(root.Left, x, y) || isCousins(root.Right, x, y)
	return (nodeEquals(root.Left, x) && nodeEquals(root.Right, y)) || (nodeEquals(root.Left, y) && nodeEquals(root.Right, x)) || isCousins(root.Left, x, y) || isCousins(root.Right, x, y)

}

func nodeEquals(node *TreeNode, x int) bool {
	if node == nil {
		return false
	}

	return (node.Val == x)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
