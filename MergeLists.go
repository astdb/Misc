
// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

package main

import (
	"fmt"
)

func main () {
	l1 := makeList([]int{1,2,4})
	printList(l1)

	l2 := makeList([]int{1,3,4})
	printList(l2)

	printList(mergeTwoLists(l1,l2))
}

 func makeList(vals []int) *ListNode {
	 if len(vals) == 0 {
		 return nil
	 }

	//  head := makeListNode(vals[1])
	//  run := head
	var run, head *ListNode

	 for i,v := range vals {
		if i == 0 {
			head = makeListNode(v)
	 		run = head
		} else {
			next := makeListNode(v)
			run.Next = next
			run = run.Next
		}
	 }

	 return head
 }

 func printList(list *ListNode) {
	if list == nil {
		fmt.Println("[]")
		return
	}

	fmt.Printf("[%d", list.Val)
	list = list.Next

	for list != nil {
		fmt.Printf(", %d", list.Val)
		list = list.Next
	}

	fmt.Print("]\n")
 }

 func makeListNode(i int) *ListNode {
	 var node ListNode
	 node.Val = i
	 node.Next = nil
	 return &node
 }

 // Definition for singly-linked list.
 type ListNode struct {
	     Val int
	     Next *ListNode
	}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	head = nil

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	
	merged := head
	runner1 := l1
	runner2 := l2

	i := 0
	for {
		if runner1.Val < runner2.Val {
			if i == 0 {
				head = runner1
				merged = head
				runner1 = runner1.Next
			} else {
				merged.Next = runner1
				merged = merged.Next
				runner1 = runner1.Next
			}
			
		} else {
			if i == 0 {
				head = runner2
				merged = head
				runner2 = runner2.Next
			} else {
				merged.Next = runner2
				merged = merged.Next
				runner2 = runner2.Next
			}
			
		}

		// if runner1 == nil && runner2 == nil {
		// 	return head
		// }

		if runner1 == nil {
			merged.Next = runner2
			// runner2 = runner2.Next
			return head
		}

		if runner2 == nil {
			merged.Next = runner1
			// runner1 = runner1.Next
			return head
		}

		i++
	}
}
