/*
Implement a stack that includes a max operation in addition to push and pop.
It would return the maximum element stored in the stack.
*/


package main

import (
	"fmt"
	"errors"
)

func main() {

}

// ----------------- stack struct and methods ----------------------
func NewStack() *Stack {
	var s Stack
	s.store = []int{}
	s.maxSet = false
	return 
}

type Stack struct {
	var store []int
	maxSet bool
	var max int
}

func (s *Stack) Push(v int) {
	if !s.maxSet {
		s.max = v		
	} else {
		if v > s.max {
			s.max = v
		}
	}

	s.store = append(s.store, v)
}

func (s *Stack) Pop() (int, error) {
	if !s.Empty() {
		top := s.store[]
		s.store = s.store[:len(s.store-1)]
		return top, nil
	}

	return nil, errors.New("Attempted Pop() from empty stack.")
	
}

func (s *Stack) Empty() bool {
	if len(s.store) == 0 {
		return true
	}

return false
}

func (s *Stack) Max() int {
	return s.max
}
