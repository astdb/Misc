/*
Implement a stack whihc provides an minimum() operation in addition to push()/pop()/peek() (all constant-time runtime complexity)

*/

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(-1)
	stack.Push(-2)
	stack.Push(4)
	stack.Push(5)
	stack.Push(1)
	stack.Print()

	min, err := stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}

	pop, err := stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Popped %d off.\n", pop.Val)
	}
	stack.Print()
	min, err = stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}

	pop, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Popped %d off.\n", pop.Val)
	}
	stack.Print()
	min, err = stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}

	pop, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Popped %d off.\n", pop.Val)
	}
	stack.Print()
	min, err = stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}

	pop, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Popped %d off.\n", pop.Val)
	}
	stack.Print()
	min, err = stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}

	pop, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Popped %d off.\n", pop.Val)
	}
	stack.Print()
	min, err = stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}

	pop, err = stack.Pop()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Popped %d off.\n", pop.Val)
	}
	stack.Print()
	min, err = stack.Min()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Min: %d\n", min)
	}
}

type Stack struct {
	Size int
	Mins []int
	Top  *StackNode // head of linked list containing stack structure
}

func NewStack() *Stack {
	var x Stack
	x.Size = 0

	var node *StackNode
	x.Top = node

	x.Mins = []int{}

	return &x
}

func (x *Stack) Print() {
	fmt.Printf("[")

	i := 0

	var stackIterator *StackNode
	stackIterator = x.Top
	for stackIterator != nil {
		if i == 0 {
			fmt.Printf("%d", stackIterator.Val)
			i++
		} else {
			fmt.Printf(", %d", stackIterator.Val)
		}

		stackIterator = stackIterator.Next
	}

	fmt.Printf("]\n")
}

func (x *Stack) Min() (int, error) {
	var min int
	var err error

	if len(x.Mins) == 0 {
		err = errors.New("Empty stack.")
		return min, err
	}

	min = x.Mins[len(x.Mins)-1]
	return min, err
}

// insert at head of linked list containing stack
func (x *Stack) Push(i int) {
	node := NewStackNode(i)
	node.Next = x.Top
	x.Top = node

	x.Size++

	if len(x.Mins) == 0 {
		x.Mins = append(x.Mins, i)
	} else {
		if x.Mins[len(x.Mins)-1] < i {
			// prev min still the smallest
			x.Mins = append(x.Mins, x.Mins[len(x.Mins)-1])
		} else {
			x.Mins = append(x.Mins, i)
		}
	}
}

// remove and return head node of list containing stack
func (x *Stack) Pop() (*StackNode, error) {
	var res *StackNode
	var err error
	if x.Top == nil {
		err = errors.New("Empty stack.")
		return res, err
	}

	x.Size--

	x.Mins = x.Mins[:len(x.Mins)-1]
	res = x.Top
	x.Top = x.Top.Next
	return res, err
}

type StackNode struct {
	Val  int
	Next *StackNode
}

func NewStackNode(v int) *StackNode {
	var sn StackNode
	sn.Val = v
	sn.Next = nil

	return &sn
}
