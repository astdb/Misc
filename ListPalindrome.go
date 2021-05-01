/*
Check if a given linked list is a palindrome.

Use the list size (if known) or single/double hop iterator method (if size not known), to push the first half of the list onto a stack.
Iterate forward from middle node, and pop/compare items from stack to each node from  the middle.
*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
  tests := []*List{CreateList([]int{1,2,3,2,1}), CreateList([]int{1, 2, 3, 3, 2, 1})}

  for _, test := range tests {
    PrintList(test)
    log.Println(ListPalindrome(test))
  }


	/*list1 := CreateList([]int{1, 2, 3, 2, 1})
	PrintList(list1)
  log.Println(ListPalindrome(list1))

	list2 := CreateList([]int{1, 2, 3, 3, 2, 1})
	PrintList(list2)
  log.Println(ListPalindrome(list1))*/

  /*list3 := CreateList([]int{1,2,3})
  PrintList(list3)
  log.Println(ListPalindrome(list3))
  */
}

func ListPalindrome(list *List) bool {
	if list == nil {
		return false
	}

	stack := []int{}

	var it1 *ListNode
	var it2 *ListNode

	it1 = list.Head
	it2 = list.Head

	for it1 != nil && it2 != nil && it2.Next != nil {
		stack = append(stack, it1.Val)

		it1 = it1.Next
		it2 = it2.Next.Next
	}

  if it1 != nil {
    stack = append(stack, it1.Val)
  }

  log.Printf("ListPalindrome(): stack = %v\n", stack)
  log.Printf("ListPalindrome(): it1.Val = %v\n", it1.Val) 

	// check if list length is odd or even
	if it2 != nil && it2.Next != nil {
		// list size is odd - middle element is already loaded to stack - pop it without comparison
    log.Println("ListPalindrome(): list size is odd.")
		stack = stack[:len(stack)-1]
	}

	// iterate forward single-step, popping and comparing with stack contents as advancing
	for it1 != nil {
		val := it1.Val
		stackVal := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
    it1 = it1.Next

		if val != stackVal {
			return false
		}
	}

	return true
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
			list.InsertTail(listVals[i])
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
