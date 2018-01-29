
// wrapper class for a LinkedList

public class LinkedList {
    private ListNode head = null;
    
    public LinkedList(int data){
        System.out.println("Creating new LinkedList..");
        this.head = new ListNode(data);
    }

    // public getListHead() {
    //     return this.head;
    // }

    public void addNode(int data){
        if(this.head == null) {
            this.head = new ListNode(data);
            return;
        }

        ListNode list = this.head;
        while(list.getNext() != null) {
            list = list.getNext();
        }

        // list.next = new ListNode(data);
        ListNode newNode = new ListNode(data);
        list.setNext(newNode);
    }

    public void deleteNode(int data) {
        if(head != null){
            if(head.getData() == data) {
                // move head
                this.head = head.getNext();
                return;
            }

            ListNode list = head;
            while(list.getNext() != null) {
                if(list.getNext().getData() == data) {
                    // if the next node is the node to be deleted, set this node's next attribute to the next attribute of the node to be deleted
                    list.setNext(list.getNext().getNext());
                    return;
                }

                list = list.getNext();
            }
        }
    }

    public void printList() {
        if(this.head == null){
            System.out.println("<empty list>");
            return;
        }

        ListNode list = this.head;
        while(list != null) {
            System.out.println(list.getData());
            list = list.getNext();
        }
    }
}


// class representing a linked list node
class ListNode {
    private int data;
    private ListNode next;

    protected ListNode(int data) {
        System.out.println("Creating new ListNode..");
        this.data = data;
        this.next = null;
    }

    protected int getData() {
        return this.data;
    }

    protected ListNode getNext() {
        return this.next;
    }

    protected void setNext(ListNode node){
        this.next = node;
    }

    protected void setNext(int data) {
        this.data = data;
    }
}
