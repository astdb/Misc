/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

func sumOfLeftLeaves(root *TreeNode) int {
  res := 0

  if root != nil {
    if isLeaf(root.Left) {
      res += root.Left.Val
    } else {
      res += sumOfLeftLeaves(root.Left)
    }

    res += sumOfLeftLeaves(root.Right)
  }

  return res
}

func isLeaf(node *TreeNode) bool {
  if node == nil {
    return false
  }

  if node.Left == nil && node.Right == nil {
    return true
  }

  return false
}


type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func sumOfLeftLeaves1(root *TreeNode) int {
    if root == nil {
      return 0
    }

    if root.Left != nil {
      return root.Left.Val + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
    } else if root.Right != nil {
      return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
    } else {
      return 0
    }
}
