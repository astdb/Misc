  /*
Provided only the middle node of a linked list, delete that node from list.

*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	list := CreateList([]int{1, 2, 3, 4, 5})
	PrintList(list)

	mid := getMidNode(list)
	log.Println(mid.Val)

	delMiddleNode(mid)
	PrintList(list)
}

// copy the data from the middle->Next node to middle
func delMiddleNode(middle *ListNode) {
	if middle.Next != nil {
		middle.Val = middle.Next.Val
		middle.Next = middle.Next.Next
	}
}

// given a list, return middle node
func getMidNode(list *List) *ListNode {
	// initiate two list iterators
	// advance one by one, and the other by two
	// when the latter reaches the end of list, former will be at the start
	var res *ListNode

	if list != nil {
		it1 := list.Head
		it2 := list.Head

		for it1 != nil && it2 != nil && it2.Next != nil {
			it1 = it1.Next
			it2 = it2.Next.Next
		}

		res = it1
	}

	return res
}

func PrintList(list *List) {
	if list != nil {
		node := list.Head

		var sb strings.Builder
		sb.WriteString("[")

		first := true
		for node != nil {
			if first {
				sb.WriteString(fmt.Sprintf("%d", node.Val))
				first = false
			} else {
				sb.WriteString(fmt.Sprintf(", %d", node.Val))
			}

			node = node.Next
		}

		sb.WriteString("]")
		log.Println(sb.String())
		return
	}

	log.Println("[]")
	return
}

func CreateList(listVals []int) *List {
	var list *List
	if len(listVals) > 0 {
		list = NewList(listVals[0])

		for i := 1; i < len(listVals); i++ {
			list.Insert(listVals[i])
		}
	}

	return list
}

type List struct {
	Head *ListNode
}

// func (r *robot) Action(commandLine string) error {
func (l *List) Insert(v int) {
	newNode := NewNode(v, l.Head.Next)
	l.Head.Next = newNode
}

func NewList(head int) *List {
	var list List
	list.Head = NewNode(head, nil)
	return &list
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int, next *ListNode) *ListNode {
	var ln ListNode
	ln.Val = val
	ln.Next = next
	return &ln
}
