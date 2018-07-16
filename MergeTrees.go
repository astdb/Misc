/*
Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Example 1:
Input:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
Merged tree:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
Note: The merging process must start from the root nodes of both trees.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("BOO!")
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

func NewTreeNode() *TreeNode {
	var tn TreeNode
	tn.Val = 0
	tn.Left = nil
	tn.Right = nil
	return &tn
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	t3 := NewTreeNode()

	if t1 != nil && t2 != nil {
		t3.Val = t1.Val + t2.Val

		t3.Left = mergeTrees(t1.Left, t2.Left)
		t3.Right = mergeTrees(t1.Right, t2.Right)

		return t3
	} else if t1 != nil {
		t3.Val = t1.Val

		t3.Left = mergeTrees(t1.Left, nil)
		t3.Right = mergeTrees(t1.Right, nil)

		return t3
	} else if t2 != nil {
		t3.Val = t2.Val

		t3.Left = mergeTrees(nil, t2.Left)
		t3.Right = mergeTrees(nil, t2.Right)

		return t3
	} else {
		return nil
	}

	// return nil
}
