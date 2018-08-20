/*
Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.



For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.



Note:

Both of the given trees will have between 1 and 100 nodes.
*/

package main

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// check if the leaf nodes (in order from left to right) are the same for two given binary trees
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// get leaf node lists for two trees
	r1Leafs := getLeafNodes(root1)
	r2Leafs := getLeafNodes(root2)

	// return comparison result for the two leaf node lists of two trees
	return compareSlices(r1Leafs, r2Leafs)
}

// helper function to compare two int slices
func compareSlices(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// return a list of leaf nodes (left to right) inthe given TreeNode
func getLeafNodes(node *TreeNode) []int {
	leafNodes := []int{} // list of leaf nodes

	if node.Left == nil && node.Right == nil {
		// leaf node - add to list and return
		leafNodes = append(leafNodes, node.Val)
		return leafNodes
	}

	// not a leaf node - collect leaf node lists for left and right subtrees, append them and return
	leftLeafs := []int{}
	if node.Left != nil {
		leftLeafs = getLeafNodes(node.Left)
	}

	rightLeafs := []int{}
	if node.Right != nil {
		rightLeafs = getLeafNodes(node.Right)
	}

	leafNodes = append(leftLeafs, rightLeafs...)
	return leafNodes
}
