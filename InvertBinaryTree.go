/*
Invert a binary tree.
Example:
Input:
    4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:
    4
   /   \
  7     2
 / \   / \
9   6 3   1
Trivia:
This problem was inspired by this original tweet by Max Howell:
Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

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
// "fmt"
)

func main() {

}

var f = true

func invertTree(root *TreeNode) *TreeNode {
	// Extract elements preorder (Root, LR)
	// insert back, in inverted order (larger to left, smaller to right)
	elements := []int{}

	// extract tree elements in preorder
	elements = preOrderExtract(root, elements)

	// insert elements back
	// root = nil
	
	// initialize to anything, maipulate within insert
	root = NewTreeNode(1)

	for _, v := range elements {
		insertTree(root, v)
	}

	return root
}

func insertTree(node *TreeNode, val int) {
	// if node == nil {
	if f {
		node = NewTreeNode(val)
		f = false
	} else {
		if val > node.Val {
			node = node.Left
			insertTree(node, val)

		} else {
			node = node.Right
			insertTree(node, val)
		}
	}
}

func preOrderExtract(node *TreeNode, elems []int) []int {
	if node != nil {
		// extract root
		elems = append(elems, node.Val)
		elems = preOrderExtract(node.Left, elems)
		elems = preOrderExtract(node.Right, elems)
	}

	return elems
}

func NewTreeNode(val int) *TreeNode {
	var t TreeNode
	t.Val = val
	t.Left = nil
	t.Right = nil

	return &t
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
