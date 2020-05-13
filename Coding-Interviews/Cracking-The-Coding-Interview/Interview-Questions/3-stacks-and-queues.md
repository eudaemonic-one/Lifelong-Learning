# Stacks and Queues

## Implementing a Stack

* A stack uses LIFO (last-in first-out) ordering
* It uses the following operations:
  * `pop()`: Remove the top item from the stack
  * `push(item)`: Add an item to the top of the stack
  * `peek()`: Return the top of the stack
  * `isEmpty()`: Return true if and only if the stack is empty
* One case where stacks are often useful is in certain **recursive algorithms**
  * Sometimes you need to push temporary data onto a stack as you recurse, but then remove them as you backtrack
* A stack can also be used to implement a recursive algorithm iteratively

```java
public class MyStack<T> {
  private static class StackNode<T> {
    private T data;
    private StackNode<T> next;
    
    public StackNode(T data) {
      this.data = data;
    }
  }
  
  private StackNode<T> top;
  
  public T pop() {
    if (top == null) {
      throw new EmptyStackException();
    }
    T item = top.data;
    top = top.next;
    return item;
  }
  
  public void push(T item) {
    StackNode<T> t = new StackNode<T>(item);
    t.next = top;
    top = t;
  }
  
  public T peek() {
    if (top == null) {
      throw new EmptyStackException();
    }
    return top.data;
  }
  
  public boolean isEmpty() {
    return top == null;
  }
}
```

## Implementing a Queue

* A queue implements FIFO (first-in first-out) ordering
* It uses the operations:
  * `add(item)`: Add an item to the end of the list
  * `remove()`: Remove the first item in the list
  * `peek()` : Return the top of the queue
  * `isEmpty()`: Return true if and only if the queue is empty
* One place where queues are often used is in **breadth-first search** or in implementing a **cache**
  * In breadth-first search, for example, we used a queue to store a list of the nodes that we need to process
  * Each time we process a node, we add its adjacent nodes to the back of the queue

```java
public class MyQueue<T> {
  private static class QueueNode<T> {
    private T data;
    private QueueNode<T> next;
    
    public QueueNode(T data) {
      this.data = data;
    }
  }
  
  private QueueNode<T> first;
  private QueueNode<T> last;
  
  public void add(T item) {
    QueueNode<T> t = new QueueNode<T>(item);
    if (last != null) {
      last.next = t;
    }
    last = t;
    if (first == null) {
      first = last;
    }
  }
  
  public T remove() {
    if (first == null) {
      throw new NoSuchElementException();
    }
    T data = first.data;
    first = first.next;
    if (first == null) {
      last = null;
    }
    return data;
  }
  
  public T peek() {
    if (first == null) {
      throw new NoSuchElementException();
    }
    return first.data;
  }
  
  public boolean isEmpty() {
    return first == null;
  }
}
```

## Interview Questions

* **3.1 Three in One:**
	* Describe how you could use a single array to implement three stacks.
* **3.2 Stack Min:**
	* How would you design a stack which, in addition to push and pop, has a function min which returns the minimum element? Push, pop and min should all operate in O(1) time.
* **3.3 Stack of Plates:**
	* Imagine a (literal) stack of plates. If the stack gets too high, it might topple. Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold. Implement a data structure SetOfStacks that mimics this. SetOfStack should be composed of several stacks and should create a new stack once the previous one exceeds capacity. SetOfStack.push() and SetOfStacks.pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack).
	* FOLLOW UP
	* Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.
* **3.4 Queue via Stacks:**
	* Implement a MyQueue class which implements a queue using two stacks.
* **3.5 Sort Stack:**
	* Write a program to sort a stack such that the smallest items are on the top. You can use an additional temporary stack, but you may not copy the elements into any other data structure (such as an array). The stack supports the following operations: push, pop, peek, and isEmpty.
* **3.6 Animal Shelter:**
	* An animal shelter, which holds only dogs and cats, operates on a strictly "first in, first out" basis. People must adopt either the"oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). They cannot select which specific animal they would like. Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCat. You may use the built-in Linked list data structure.
