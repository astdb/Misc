
// A stack provides last-in first-out access to data
// this class represents a simple stack data structure implemented using an underlying linked-list

public class Stack {
    private StackItem top;
    private int size;

    public Stack(int firstItem) {
        this.top = new StackItem(firstItem);
        this.size++;
    }

    public void push(int item) {
        if(this.top == null) {
            this.top = new StackItem(item);
            this.size++;
            return;
        }

        StackItem temp = this.top;
        StackItem newItem = new StackItem(item);
        newItem.setNext(temp);
        this.top = newItem;
        this.size++;
        return;
    }
    
    public int peek() {
        return top.value();
    }

    public int pop()  {
        if(this.top == null) {
            // TODO: implement and throw proper exception here instead of quitting the whole JVM
            System.out.println("Exception: empty stack;");
            System.exit(1);
        }

        int val = this.top.value();
        this.top = this.top.next();
        this.size--;
        return val;
    }

    public int size() {
        return this.size;
    }
    
    public boolean isEmpty() {
        if(this.size > 0) {
            return false;
        }

        return true;
    }

    public void print() {
        if(isEmpty()) {
            System.out.println("\t<empty stack>");
            return;
        }

        System.out.println("-----------------------");
        System.out.println("\tSize: " + this.size());
        StackItem i = this.top;
        while(i != null) {
            System.out.println("\t" + i.value());
            i = i.next();
        }
    }
}

class StackItem {
    private int data;
    private StackItem next;

    protected StackItem(int data) {
        this.data = data;
        this.next = null;
    }

    protected int value() {
        return this.data;
    }

    protected StackItem next() {
        return this.next;
    }

    protected void setNext(StackItem next) {
        this.next = next;
    }
}
