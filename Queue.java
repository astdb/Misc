
// simple queue implementation with a doubly-linked list. 

public class Queue {
    private QueueElement head;
    private QueueElement tail;
    int size;

    public Queue(int i) {
        QueueElement start = new QueueElement(i);
        start.setNext(null);
        start.setPrev(null);
        this.head = start;
        this.tail = start;
        this.size++;
        return;
    }

    public void enQueue(int i) {
        // insert at tail
        QueueElement elem = new QueueElement(i);
        elem.setPrev(this.tail);
        this.tail.setNext(elem);
        this.size++;
        return;
    }

    public QueueElement deQueue(int i) {
        // remove from head
        if(queue.size <= 0) {
            // throw new EmptyQueueException()
            return null;
        }

        this.head
    }

    public boolean isEmpty() {
        if(this.size <= 0) {
            return true;
        }

        return false;
    }
}

class QueueElement {
    private int data;
    private QueueElement next;
    private QueueElement prev;

    protected QueueElement(int data) {
        this.data = data;
    }

    protected void setNext(QueueElement next) {
        this.next = next;
    }

    protected void setPrev(QueueElement prev) {
        this.prev = prev;
    }

    protected int data() {
        return this.data;
    }

    protected QueueElement next() {
        return this.next;
    }

    protected QueueElement prev() {
        return this.prev;
    }
}
