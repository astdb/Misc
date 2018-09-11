package main

import (
    "fmt"
)

func main() {

}

type Node struct {
    Val int
    Children []*Node
}

func postOrder(root *Node, res []int) []int {
if root != nil {
	// res := []int{}
		for _, v := range root.Children {
			res = append(res, postOrder(v, res)...)
			res = append(res, v.Val)
		}

		res = append(res, res.Val)
    return res
} else {
	return res
}
}

