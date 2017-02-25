/**
  Design an API to which receives, stores and returns top 10 of all orders placed on a large e-commerce site (presumably largest site on the internet for purchased-goods, and by purchased goods I mean, um, books..).
  For simplicity's sake assume that you're able to store all order data on a single machine. An order consists of an order ID and a price.
  The API must be able to return top 10 orders sorted by price to be returned. Try to make both storage and top 10 retrieval functions
  as fast as possible. The top 10 should consider upto the latest of orders placed.

  Solution Design
  ---------------
   - Use a linkedlist for storage, as it would allow linear-time ordered insertion
   - Make the list have ordered sorted from highest to lowest, with the head being the highest order and the tail being lowest.
   -
*/

import java.util.*;

// driver class for OrderAPI
public class OrderAPIDriver {
  public static void main(String[] args){
    // initialize API
    OrderAPI _theOrderAPI = OrderAPI.createOrderAPI();

    // add some orders
    double rangeMin = 0.01;
    double rangeMax = 9999.0;
    for(int i = 0; i < 15; i++){
      Random r = new Random();
      double randPrice = rangeMin + (rangeMax - rangeMin) * r.nextDouble();
      _theOrderAPI.receiveOrder((i+1), randPrice);
    }

    // System.out.println("Orders loaded.");
    // System.exit(1);

    _theOrderAPI.printOrders();

    OrderNode[] topTen = _theOrderAPI.returnTopTen();

    System.out.print("TOP TEN\n-----------\n");
    for(OrderNode _order : topTen){
      if(_order != null){
        // System.out.println("(" + _order.getID() + ", " + _order.getPrice() + ")");
        System.out.printf("(%d, $%.2f)\n", _order.getID(), _order.getPrice());
      } else {
        System.out.println("(<null order object>)");
      }
    }
  }

}

// OrderAPI would provide storage structure for orders, and functions to receive and return orders - should ideally be a singleton
class OrderAPI {
  // Linked List of
  private OrderNode orderList = null;
  private static OrderAPI _singleinstance = null;

  protected OrderAPI(){
    //..
  }

  public static OrderAPI createOrderAPI(){
    if( _singleinstance == null ){
      _singleinstance = new OrderAPI();
    }

    return _singleinstance;
  }

  public void printOrders(){
    OrderNode iterator = this.orderList;
    System.out.print("ALL ORDERS\n-----------\n");

    while( iterator != null ){
      // System.out.println("(" + iterator.getID() + ", " + iterator.getPrice() + ")");
      System.out.printf("(%d, %.2f)\n", iterator.getID(), iterator.getPrice());
      iterator = iterator.getNext();
    }
    System.out.println();
  }

  // API endpoint to return the top 10 orders
  public OrderNode[] returnTopTen(){
    OrderNode[] topTen = new OrderNode[10];

    OrderNode iterator = this.orderList;

    int i = 0;
    while( iterator != null && i < 10 ){
      topTen[i] = iterator;
      i++;

      iterator = iterator.getNext();
    }

    return topTen;
  }

  // API endpoint to receive placed orders and store on orderList
  public boolean receiveOrder(int theID, double price){
    // create an order node
    OrderNode thisOrd = new OrderNode(theID, price);

    if( this.orderList == null ) {
      // very first order
      this.orderList = thisOrd;
      return true;
    }

    // we have some orders stored already, insert this one on the right position based on price
    OrderNode iterator = this.orderList;
    OrderNode previous = null;

    boolean inserted = false;
    while(iterator != null) {
      if( iterator.getPrice() <= thisOrd.getPrice()) {
        if ( previous != null ) {
          // inserting somewhere in the middle of the list
          previous.setNext(thisOrd);
          thisOrd.setNext(iterator);
          inserted = true;
          return true;
        } else {
          // inserting at head
          thisOrd.setNext(iterator);
          this.orderList = thisOrd;
          inserted = true;
          return true;
        }
      }

      // no insertion, moving onto next
      previous = iterator;
      iterator = iterator.getNext();
    }

    if ( !inserted ) {
      // insert at the end
      previous.setNext(thisOrd);
      return true;
    }

    return inserted;
  }
}

// This class would define a node of the linkedlist or orders
class OrderNode {
  private int orderID = -1;
  private double orderPrice = -1.0;
  private OrderNode next = null;

  public OrderNode(int id, double price){
    this.orderID = id;
    this.orderPrice = price;
  }

  public double getPrice(){
    return this.orderPrice;
  }

  public int getID(){
    return this.orderID;
  }

  public OrderNode getNext(){
    return this.next;
  }

  public void setNext(OrderNode next){
    this.next = next;
  }
}
