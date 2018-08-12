/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

package main

import (
  "fmt"
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
func maxDepth(root *TreeNode) int {
    if root != nil {
      // tree is empty
      return 0
    } else {
      // compute subtree depths
      left_depth := maxDepth(root.Left)
      right_depth := maxDepth(root.Right)

      // return larger depth, adding one to accomodate current
      if left_depth > right_depth {
        return left_depth +1 
      } else {
        return right_depth + 1
      }
    }
}
