/*
Given two singly-linked, sorted lists, merge them into one.
Only field allowed to be modified is a node's next field.
*/

package main

import (
	"fmt"
)

func main() {
	// construct List x
	x := NewListNode(2)
	x.Next = NewListNode(5)
	x.Next.Next = NewListNode(7)

	// construct List y
	y := NewListNode(3)
	y.Next = NewListNode(11)

	PrintList(x)
	PrintList(y)

	PrintList(MergeLists(x, y))

}

func MergeLists(x, y *ListNode) *ListNode {
	// which node to start with
	var NewList *ListNode
	var NewListHead *ListNode

	firstRun := true
	for x != nil || y != nil {
		if firstRun {
			if x.Val > y.Val {
				// fmt.Printf("Setting Y as head (%d)\n", y.Val)
				NewList = NewListNode(y.Val)
				y = y.Next
			} else {
				// fmt.Printf("Setting X as head (%d)\n", x.Val)
				NewList = NewListNode(x.Val)
				x = x.Next
			}

			NewListHead = NewList

			firstRun = false
		}

		if x != nil && y != nil {
			if x.Val > y.Val {
				// NewList = NewListNode(y.Val)
				// fmt.Printf("\nInserting (%d) from Y\n", y.Val)
				NewList.Next = y
				NewList = NewList.Next
				y = y.Next
			} else {
				// NewList = NewListNode(x.Val)
				// fmt.Printf("\nInserting (%d) from X\n", x.Val)
				NewList.Next = x
				NewList = NewList.Next
				x = x.Next
			}

		} else if x != nil {
			NewList.Next = x
			NewList = NewList.Next
			x = x.Next

		} else if y != nil {
			NewList.Next = y
			NewList = NewList.Next
			y = y.Next
		}
	}

	return NewListHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int) *ListNode {
	var ln ListNode
	ln.Val = v
	ln.Next = nil

	return &ln
}

func PrintList(x *ListNode) {
	if x == nil {
		fmt.Println(" <empty>")
	} else {
		for x != nil {
			fmt.Printf(" %d ", x.Val)
			x = x.Next
		}

		fmt.Printf("\n")
	}
}
