package main

import (
	"fmt"
)

func main() {
	// root := NewTreeNode(4)
	// root.Left = NewTreeNode(2)
	// root.Right = NewTreeNode(7)
	
	var root *TreeNode
	root = nil
	PrepNode(root, 4, 2, 7)
	
	PrintTreeNode(root)
}

func PrepNode(t *TreeNode, v1, v2, v3 int) {
	t = NewTreeNode(v1)
	t.Left = NewTreeNode(v2)
	t.Right = NewTreeNode(v3)
}

func PrintTreeNode(t *TreeNode) {
	if t == nil {
		fmt.Println("Nil tree node.")
		return
	}

	fmt.Printf("Val: %d", t.Val)
	
	if t.Left != nil {
		fmt.Printf(", Left: %d", t.Left.Val)
	} else {
		fmt.Printf(", Left: nil")
	}
	
	if t.Right != nil {
		fmt.Printf(", Right: %d\n", t.Right.Val)
	} else {
		fmt.Printf(", Right: nil\n")
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

func (t *TreeNode) Insert(val int) {
	if t == nil {
		t = NewTreeNode(val)
	}

	
}
