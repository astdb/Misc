/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?

Accepted
457,431
Submissions
900,819
*/

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */

public class ReverseLinkedList {
    public static void main(String[] args) {
        ListNode test1 = new ListNode(1);
        test1.next = new ListNode(2);
        test1.next.next = new ListNode(3);
        test1.next.next.next = new ListNode(4);
        test1.next.next.next.next = new ListNode(5);        

        printList(test1);
        printList(reverseList(test1));

        ListNode test2 = new ListNode(1);
        test2.next = new ListNode(2);
        // test2.next.next = new ListNode(3);

        printList(test2);
        printList(reverseList(test2));
    }

    public static ListNode reverseList(ListNode head) {
        ListNode reverse = null;
        ListNode node = null;

        while(head != null) {
            node = new ListNode(head.val);
            node.next = reverse;
            reverse = node;

            head = head.next;
        };
        
        return reverse;
    }

    public static void printList(ListNode list) {
        if(list == null) {
            System.out.println("NULL");
        } else {
            while(list != null){
                System.out.print(list.val + " ");
                list = list.next;
            }
            System.out.println();
        }
    }
}

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}
