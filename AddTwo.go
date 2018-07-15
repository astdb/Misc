/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

// To execute Go code, please declare a func main() in a package "main"

package main

import (
  "fmt"
  "math"
)

func main() {
  if ListToNum(addTwoNumbers(NumToList(243), NumToList(564))) == 807 {
	fmt.Println("PASS (243 + 564 == 807)")
  } else {
	fmt.Println("FAIL (243 + 564 != 807)")
  }

  if ListToNum(addTwoNumbers(NumToList(0), NumToList(0))) == 0 {
	fmt.Println("PASS (0 + 0 == 0)")
  } else {
	fmt.Println("FAIL (0 + 0 != 0)")
  }

  if ListToNum(addTwoNumbers(NumToList(1), NumToList(1))) == 2 {
	fmt.Println("PASS (1 + 1 == 2)")
  } else {
	fmt.Println("FAIL (1 + 1 != 2)")
  }

  if ListToNum(addTwoNumbers(NumToList(12564), NumToList(654983))) == 667547 {
	fmt.Println("PASS (12564 + 654983 == 667547)")
  } else {
	fmt.Println("FAIL (12564 + 654983 != 667547)")
  }

  if ListToNum(addTwoNumbers(NumToList(1000000), NumToList(1000000))) == 2000000 {
	fmt.Println("PASS (1000000 + 1000000 == 2000000)")
  } else {
	fmt.Println("FAIL (1000000 + 1000000 != 2000000)")
  }
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
      Val int
      Next *ListNode
  }

// function to create new *ListNode
func NewListNode(val int, next *ListNode) *ListNode {
	var ln ListNode	
	
	ln.Val = val
	ln.Next = next

	return &ln
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//   if list != nil {
//     for list.Next != nil {
//       digits = append(digits, list.Val)
//       list = list.Next
//     }
//   }
	return NumToList(ListToNum(l1) + ListToNum(l2))
}

// Print out the number a list represents
func PrintListNum(list *ListNode) {
	if list != nil {
		for list != nil {
			fmt.Printf("%d", list.Val)
			list = list.Next
		}
		
		fmt.Print("\n")
	  } else {
	   fmt.Println("0") 
	  }
}

// transform a linked list representing a number into int format (decimal)
func ListToNum(list *ListNode) int {
  digits := []int{}  // slice to store numeric digits 
  
  // go through list and collect digits
  if list != nil {
    for list != nil {
      digits = append(digits, list.Val)
      list = list.Next
    }
    
    total := 0  // final numeric value represented by the list
    k := 0  // slice index for iterating digits slice
    for i := len(digits)-1; i >= 0; i-- {
      total += int(math.Pow10(i)) * digits[k]
      k++
	}
	
	return total
    
  } else {
   return 0 
  }  
}

// transform an int to linked list format
func NumToList(n int) *ListNode {
	// get a digit slice of n
	ndigits := []int{}
	rem := 0
	for n > 0 {
		rem = n % 10
		n = n / 10
		ndigits = append(ndigits, rem)
	}

	if len(ndigits) > 0 {
		list := NewListNode(ndigits[len(ndigits)-1], nil)
		head := list

		for i := len(ndigits)-2; i >=0; i-- {
			list.Next = NewListNode(ndigits[i], nil)
			list = list.Next			
		}

		return head
	} else {
		return nil
	}
}
