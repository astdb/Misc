// find the middle node of a linked list

package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
// declare two pointers to iterate through the list
	// advance first by one and second by two
	// when the second pointer reaches the end of the list,
	// first will be at middle
	if head == nil {
		return head
	}

	firstIterator := head
	secondIterator := head

	for secondIterator != nil && secondIterator.Next != nil {
		firstIterator = firstIterator.Next		
		secondIterator = secondIterator.Next.Next
	}

	return firstIterator
}

