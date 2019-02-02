/*
Given a tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

Example 1:
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9
Note:

The number of nodes in the given tree will be between 1 and 100.
Each node will have a unique integer value from 0 to 1000.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("IncOrderBST starting...")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	var NewTree *TreeNode
	var currentNode *TreeNode

	inOrderVals(root, vals)

	for k, v := range vals {
		if k == 0 {
			NewTree = NewTreeNode(v)
			NewTree.Right = currentNode
		} else {
			currentNode = NewTreeNode(v)
			currentNode = currentNode.Right
		}
	}

	return NewTree
}

func inOrderVals(root *TreeNode, vals []int) {
	if root != nil {
		inOrderVals(root.Left, vals)
		vals = append(vals, root.Val)
		inOrderVals(root.Right, vals)
	}
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
