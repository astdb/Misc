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
	// "log"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// get a TreeNode reference with a given numeric value
func NewTreeNode(nodeVal int) *TreeNode {
	var tn TreeNode

	tn.Val = nodeVal
	tn.Left = nil
	tn.Right = nil

	return &tn
}

// insert treeNode into given BST root
func InsertNode(root, treeNode *TreeNode) *TreeNode {
	var y *TreeNode
	var x *TreeNode

	x = root

	for x != nil {
		y = x
		if treeNode.Val < x.Val {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	// treeNode.p = y

	if y == nil {
		root = treeNode
	} else if treeNode.Val < y.Val {
		y.Left = treeNode
	} else {
		y.Right = treeNode
	}

	return root

	// if root == nil {
	// 	root = treeNode
	// 	return root
	// }

	// var curNode *TreeNode
	// var prevNode *TreeNode
	// curNode = root

	// for {
	// 	if
	// }
}

func getMinimumDifference(root *TreeNode) int {

	return 1
}
