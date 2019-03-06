/*
Implement a stack that includes a max operation in addition to push and pop.
*/

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// create stack
	testStack := NewStack()

	testStack.push(1)
	testStack.push(2)
	testStack.push(3)
	testStack.push(4)
	testStack.push(5)
	testStack.print()

	m, err := testStack.Max()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Max: %d\n", m)

	testStack.pop()
	testStack.print()
	m, err = testStack.Max()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Max: %d\n", m)

}

type Stack struct {
	Store   []int
	MaxList []int
}

func NewStack() *Stack {
	var s Stack
	s.Store = []int{}
	s.MaxList = []int{}
	return &s
}

func (s *Stack) Max() (int, error) {
	if !s.isEmpty() {
		return s.MaxList[len(s.MaxList)-1], nil
	} else {
		return 0, errors.New("Max() called on empty stack.")
	}
}

func (s *Stack) push(x int) {
	if !s.isEmpty() {
		s.Store = append(s.Store, x)

		if x > s.MaxList[len(s.MaxList)-1] {
			s.MaxList = append(s.MaxList, x)
		} else {
			s.MaxList = append(s.MaxList, s.MaxList[len(s.MaxList)-1])
		}
	} else {
		// first insert
		s.Store = append(s.Store, x)
		s.MaxList = append(s.MaxList, x)
	}

}

func (s *Stack) pop() int {
	topVal := s.Store[len(s.Store)-1]
	s.Store = s.Store[:len(s.Store)-1]
	s.MaxList = s.MaxList[:len(s.MaxList)-1]

	return topVal
}

func (s *Stack) peek() (int, error) {
	if !s.isEmpty() {
		return s.Store[len(s.Store)-1], nil
	} else {
		return 0, errors.New("Peek called on empty stack")
	}
}

func (s *Stack) isEmpty() bool {
	if len(s.Store) <= 0 {
		return true
	}

	return false
}

func (s *Stack) print() {
	if !s.isEmpty() {
		for _, v := range s.Store {
			fmt.Printf("%d ", v)
		}

		fmt.Printf("\n")
	} else {
		fmt.Println("Empty stack")
	}
}
