
// given a linked list, remove all duplicate nodes
package main

import (
  "fmt"
)

func main () {
  // create a linkedlist with int data from 1 to 5
  list01 := createListNode(1)

  newnode := createListNode(2)
  list01 = addToList(newnode, list01)

  newnode = createListNode(3)
  list01 = addToList(newnode, list01)

  newnode = createListNode(4)
  list01 = addToList(newnode, list01)

  newnode = createListNode(5)
  list01 = addToList(newnode, list01)

  newnode = createListNode(3)
  list01 = addToList(newnode, list01)

  newnode = createListNode(6)
  list01 = addToList(newnode, list01)

  // for i := 2; i < 10; i++ {
  //   newnode := createListNode(i)
  //   list01 = addToList(newnode, list01)
  // }

  // var list02 *ListNode
  printList(list01)
  // list01 = removeDups(list01)
  list01 = removeDups2(list01)
  printList(list01)
}

func removeDups2(head *ListNode) *ListNode {
  current := head

  for current != nil {
    // remove all future nodes which have the same value
    runner := current

    for runner.Next != nil {
      if runner.Next.Data == current.Data {
        runner.Next = runner.Next.Next
      } else {
        runner = runner.Next
      }
    }
    
    current = current.Next
  }

  return head
}

// removeDups(list *ListNode) removes all duplicate nodes from a linkedlist with the given head node
func removeDups(list *ListNode) *ListNode {
  // algorithm: for each node, traverse the rest of the list to see if it's duplicated, and if so remove the duplicate occurences
  // and return the list
  fmt.Println("\n\n---------Removing duplicates------------------")
  head, previous := list, list
  var node *ListNode
  if list != nil {
    // node = list.Next
    node = list
  }

  for list != nil {
    comparisonnode := list
    fmt.Println("\nremoveDups(): setting " + printNode(comparisonnode) + " as comparison node..");

    // iterate the rest of the list to see if this node is replicated anywhere
    for node != nil && node.Next != nil {
      thisnode := node.Next
      fmt.Println("\n\tremoveDups(): comparing to " + printNode(thisnode));

      if comparisonnode.Data == thisnode.Data {
        // thisnode is a duplicate of comparisonnode - delete thinode
        fmt.Println("\n\tremoveDups(): " + printNode(thisnode) + " detected as a duplicate- removing..");
        previous.Next = thisnode.Next

        // if a node got deleted, previous shouldn't change, and we go on to the next
        node = thisnode.Next
        fmt.Println("\n\tremoveDups(): going on to " + printNode(node));
      } else {
        // if a node did not get deleted, both previous and next update as we go on to the next node
        fmt.Println("\n\tremoveDups(): " + printNode(thisnode) + " not a duplicate, going on to " + printNode(node.Next));

        previous = thisnode
        node = node.Next
      }
    }

    // done scanning the rest of the list for comparisonnode, onto the next
    fmt.Println("\nremoveDups(): done scanning the rest of the list for comparison node " + printNode(comparisonnode) + " moving on to " + printNode(list.Next));

    list = list.Next
  }

  return head
}

//removes a given list node (listnode) from a given list (list) and returns the list
// func removeNode(listnode *ListNode, list *ListNode) *ListNode {
//
// }

// print out a linkedlist with the given head node
func printList(list *ListNode) {
  // fmt.Println("\n\n---------Printing out list------------------")
  fmt.Print("[")

  listEmpty := true
  for list != nil {
    listEmpty = false
    fmt.Printf("%d ", list.Data)
    list = list.Next;
  }

  if listEmpty {
    fmt.Print("<empty list>")
  }
  fmt.Print("]")
  fmt.Print("\n")
}

// create a linkedlist with given initial data
func createListNode(initData int) *ListNode {
  // creates a new linked list node with data 'initData'
  fmt.Printf("\ncreateListNode(): creating new ListNode with %d", initData)
  var L ListNode
  L.Next = nil
  L.Data = initData
  return &L
}

// add a given node to a given linkedlist
func addToList(newnode *ListNode, list *ListNode) *ListNode {
  // add newnode to list (with the head list)
  fmt.Printf("\naddToList(): adding %d to list..", newnode.Data)
  temp := list
  for temp.Next != nil {
    temp = temp.Next
  }

  temp.Next = newnode
  return list
}

// a linkedlist node
type ListNode struct {
  Next *ListNode
  Data int
}

// Debug function - print a given linkedlist node
func printNode(node *ListNode) string {
  if node != nil {
      return fmt.Sprintf("\b[Data: %d, Next: %d ]", node.Data, node.Next)
  }

  return fmt.Sprintf("[<nil node>]")
}
