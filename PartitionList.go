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
	partList := CreateList([]int{val})

	// save external reference to partition key node
	partKeyNode := partList.Head

	// iterate through list
	it := list.Head
	for it != nil {
		// if value < partition, insert at list head. 
		if it.Val < partKeyNode.Val {
			list.InsertHead(it.Val)
		} else {
			// if value > partition, insert at end of list.
			list.InsertTail(it.Val)
		}		

		it = it.Next
	}

	// return partitioned list
	return partList
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
	Tail *ListNode
}

func (l *List) Insert(v int) {
	newNode := NewNode(v, l.Head.Next)
	l.Head.Next = newNode
}

// insert new node at the head of the list, and update head
func (l *List) InsertHead(v int) {
	newNode := NewNode(v, l.Head)
	l.Head = newNode
}

// insert new node at the end of the list
func (l *List) InsertHead(v int) {
	newNode := NewNode(v, l.Head)
	l.Head = newNode

func NewList(head int) *List {
	var list List
	list.Head = NewNode(head, nil)
	list.Tail = list.Head
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
