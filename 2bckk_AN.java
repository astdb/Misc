code from http://collabedit.com/2bckk

You are given a BST, print the numbers in a sorted order.

public static printTree(Tree tree) {
    if (tree != null) {
        printTree(tree.left);        
        System.out.println(tree.getData());        
        printTree(tree.right)
    }
    
    return; 
}

Linked lists - finding loops, finding middle node

class Node {
    protected int data;
    protected Node next;
    protected node previous;
    
        // denotes if this node is a head node of a list
        protected bool head = false;
        
        // the number of nodes in this list, if the node is a head node
        protected int count =  0;
    
}

// returns true if the list has a loop
public static bool hasLoop(Node node) {
    // increment first pointer by one and second by two at a time
    // if there's a loop, they'll meet at some point, if not, one will reach the end
    if(node == null) {
        return node;
    }
    
    Node firstPointer = node;
    Node secondPointer = node;
    
    while(true) {
        if(secondPointer.next != null) {
            secondPointer = secondPointer.next.next;
            firstPointer = firstPointer.next;
            
            if(firstPointer == secondPointer) {
                // pointers have met, loop
                return true;
            }
        } else {
            // end reached, no loop
            return false;
        }
    }
}

public static Node findMiddleNode2(Node node) {
    if(node == null) {
        return node;
    }
    
    Node firstPointer = node;
    Node secondPointer = node;
    
    while(node != null) {
        // increment forst pointer by one and second by two at a time
        // when the second pointer reaches the end, the first would be at the middle
        if(secondPointer.next != null) {
            secondPointer = secondPointer.next.next;
            firstPointer = firstPointer.next;
        } else {
            return firstPointer;
        }
    }
}

public static Node findMiddleNode1(Node node) {
    if(node == null) {
        return node;
    }

    int listCount = 0;
    Node start = node;
    while (node != null) {
        listCount++;
        node = node.next;
    }
    
    int middleIndex;    
    if(listCount %  2 == 0) {
        middleIndex = listCount /2;
    } else {
        middleIndex = (listCount + 1) / 2;
    }
    
    listCount = 0;
    while (node != null) {
        listCount++;
        node = node.next;
        
        if(listCount == middleIndex) {
            return node;
        }
    }
    
}
