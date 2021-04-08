/*
Write a program that removes  duplicates from an unsorted linked list.

What if a temporary buffer was not allowed?

*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	testList := CreateList([]int{1, 2, 2, 3, 4, 5})
	PrintList(testList)
	// removeDups(testList)
	removeDups2(testList)
	PrintList(testList)
}

func CreateList(vals []int) *ListNode {
	var list *ListNode

	// i := 0
	if len(vals) > 0 {
		for k, v := range vals {
			if k == 0 {
				// populate head
				// list.Val = v
				// list.Next = nil
				list = getListNode(v, nil)
			} else {
				// create new list node and insert after head start
				node := getListNode(v, list.Next)
				// node.Val = v
				// node.Next = list.Next
				list.Next = node
			}
		}
	}

	return list
}

// iterate list from start
// for every node, scan the rest of the list in front and remove any similar nodes
func removeDups2(list *ListNode) *ListNode {
	// save reference to list head
	listRef := list

	for list != nil {
		// reference to current node
		curNode := list
		curVal := curNode.Val

		// iterate through the rest of the list
		prevNode := curNode
		curNode = curNode.Next
		for curNode != nil {
			if curNode.Val == curVal {
				// remove curNode
				prevNode.Next = curNode.Next
			}

			prevNode = curNode
			curNode = curNode.Next

		}

		list = list.Next
	}

	return listRef
}

func removeDups(list *ListNode) *ListNode {
	// set of seen values
	listVals := map[int]bool{}
	var prevNode *ListNode

	// save reference to list head
	listRef := list

	for list != nil {
		val := list.Val
		_, seen := listVals[val]
		if seen {
			// remove node
			if prevNode != nil {
				prevNode.Next = list.Next

			}
		} else {
			listVals[val] = true
			prevNode = list
		}

		// prevNode = list
		list = list.Next
	}

	return listRef
}

func PrintList(list *ListNode) {
	var sb strings.Builder
	sb.WriteString("[")

	i := 0
	for list != nil {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("%d", list.Val))
			i++
		} else {
			sb.WriteString(fmt.Sprintf(", %d", list.Val))
		}

		list = list.Next
	}

	sb.WriteString("]")
	log.Printf(sb.String())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListNode(val int, next *ListNode) *ListNode {
	var ln ListNode
	ln.Val = val
	ln.Next = next

	return &ln
}
