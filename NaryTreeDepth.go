
/*
Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

For example, given a 3-ary tree:

https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png

We should return its max depth, which is 3.
 

Note:

The depth of the tree is at most 1000.
The total number of nodes is at most 5000.
*/

/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};

class Solution {
public:
    int maxDepth(Node* root) {

    }
};
*/

package main

import (
	"fmt"
)

var depth int

func main() {
	fmt.Println("NTD starting..")
}

func MaxDepth(root *TreeNode) int {
	depths := []int{}
	GetDepths(root, depths, depth)

	maxDepth := 0
	for _, v := range depths {
		if v > depth {
			maxDepth = v
		}
	}

	return maxDepth
}

func GetDepths(root *TreeNode, depths []int, depth int) {
	if root != nil {
		depth++
		// getDepths(root.Left, depths, depth)
		// getDepths(root.Right, depths, depth)

		for _, child := range root.Children {
			GetDepths(child, depths, depth)
		}
	} else {
		// returning - at a leaf
		depths = append(depths, depth)
		depth = 0
	}
}

type TreeNode struct {
	Val int
	// Left *TreeNode
	// Right *TreeNode
	Children []*TreeNode
}

func NewTreeNode(val int) *TreeNode {
	var t TreeNode
	t.Val = val
	t.Children = []*TreeNode{}

	return &t
}
