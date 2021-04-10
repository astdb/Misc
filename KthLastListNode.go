/*
Find the k-th to last element in a given linked list

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

	n := kthToLast(list, 1)
	log.Println(n.Val)
}

func kthToLast(list *List, k int) *ListNode {
	// var res *ListNode

	// two iterators traverse list, second following k places behind the lead iterator.
	// both advance one node a time, second runner at k-nodes behind the first when first reaches end of list
	first := list.Head
	second := list.Head

	i := 0
	for first != nil {
		if i >= k && second != nil {
			second = second.Next
		}

		i++
		first = first.Next
	}

	return second
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
