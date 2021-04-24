package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	// x := CreateList([]int{})
	// PrintList(x)
	// x.Insert(4)
	// PrintList(x)
	// return

	op1 := CreateList([]int{7, 6, 1}) // 612
	op2 := CreateList([]int{5, 2, 1}) // 215

	PrintList(op1)
	PrintList(op2)
	PrintList(add(op1, op2))

	op1 = CreateList([]int{9, 9})
	op2 = CreateList([]int{9, 9})
	PrintList(op1)
	PrintList(op2)
	PrintList(add(op1, op2))

	op1 = CreateList([]int{9, 9, 9})
	op2 = CreateList([]int{9, 9, 9})
	PrintList(op1)
	PrintList(op2)
	PrintList(add(op1, op2))

}

func add(list1 *List, list2 *List) *List {
	// 7->1->6 + 5->1->2 = 2->1->9
	var it1 *ListNode
	var it2 *ListNode
	it1 = list1.Head
	it2 = list2.Head

	// To store the result, we'll create a placeholder list with one element of value zero. When inserting the first element
	// value of this node will be overwritten instead of a new node being inserted.
	// Subsequent values will be inserted as per usual with new nodes created and added.
	res := CreateList([]int{0})
	var rem int
	firstInsert := true

	for it1 != nil || it2 != nil || rem != 0 {
		// thisAddRes := it1.Val + it2.Val + rem	// full addition result

		var thisAddRes int
		if it1 != nil {
			thisAddRes += it1.Val
			it1 = it1.Next
		}

		if it2 != nil {
			thisAddRes += it2.Val
			it2 = it2.Next
		}

		thisAddRes += rem

		var thisAddDig int // first significant digit from full add result (e.g. 3 if FAR is 253)
		if thisAddRes > 9 {
			thisAddDig = thisAddRes % 10
			rem = thisAddRes / 10

			// if
		} else {
			// no carry - reset rem
			rem = 0
			thisAddDig = thisAddRes
		}

		if firstInsert {
			// first insert - replace resulthead value
			res.Head.Val = thisAddDig
			firstInsert = false
		} else {
			// res.Insert(thisAddDig)
			res.InsertTail(thisAddDig)
		}
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

// insert new node at the tail of the list, and update tail
func (l *List) InsertTail(v int) {
	newNode := NewNode(v, nil)
	l.Tail.Next = newNode
	l.Tail = newNode
}

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
