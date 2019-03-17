// A binary tree is said to be height-balanced if for each node, the height difference of left and right
// subtrees is at most one (1).
// Write a program that takes the root of a BT as input and checks whether it's height-balanced

package main

import (
	"fmt"
)

func main() {
	fmt.Println("HeightBalanced")
}

func IsHeightBalanced(t *TreeNode) bool {
	if t != nil {
		leftBalanced := IsHeightBalanced(t.Left)

		// fmt.Println(t.Val)
		heightDiff := GetTreeHeight(t.Left) - GetTreeHeight(t.Right)
		if heightDiff < 0 {
			heightDiff *= (-1)
		}

		if heightDiff > 1 {
			return false
		}

		rightBalanced := IsHeightBalanced(t.Right)

		if leftBalanced && rightBalanced {
			return true
		}
	}

	// a nil tree is height balanced?
	return true
}

/*
Tree Height
1. If tree is empty then return 0
2. Else
     (a) Get the max depth of left subtree recursively  i.e.,
          call maxDepth( tree->left-subtree)
     (a) Get the max depth of right subtree recursively  i.e.,
          call maxDepth( tree->right-subtree)
     (c) Get the max of max depths of left and right
          subtrees and add 1 to it for the current node.
         max_depth = max(max dept of left subtree,
                             max depth of right subtree)
                             + 1
     (d) Return max_depth
*/
func GetTreeHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}

	leftSubTreeHeight := GetTreeHeight(t.Left)
	rightSubTreeHeight := GetTreeHeight(t.Right)

	var maxDepth int
	if leftSubTreeHeight >= rightSubTreeHeight {
		maxDepth = leftSubTreeHeight
	} else {
		maxDepth = rightSubTreeHeight
	}

	return (maxDepth + 1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	var t TreeNode
	t.Val = v
	t.Left = nil
	t.Right = nil

	return &t
}
