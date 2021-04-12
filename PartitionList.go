/*
Given a node value and a list, write code to partition the list around the given value such that all nodes with lower values
than the given value come before all nodes of equal or greater value than the given.

*/

package main

import (
	"strings"
	"log"
)

func main() {

}

func partitionList(val int, list *List) *List {
	// create new list, with partition key as head node
	partList := CreateList()

	// save external reference to partition key node

	// iterate through list
	{
		// if value < partition, insert at list head. 

		// if value > partition, insert at end of list.

	}

	// return partitioned list
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
