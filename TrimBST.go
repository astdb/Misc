/*
Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of the trimmed binary search tree.

Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2
Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
*/

package main

import (
	"fmt"
)

func main() {

}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}

	// trim subtrees in post-order
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)

	// trim root node
	// i. root key < L
	if root.Val < L {
		rightCh := root.Right
		root = nil
		return rightCh
	}

	// ii. root key > R (entire right subtree can be eliminated with root)
	if root.Val > R {
		leftCh := root.Left
		root = nil
		return leftCh
	}

	// root in range
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
