# Linked List

* A **linked list** is a data structure that represents a sequence of nodes
* In a singly linked list, each node points to the next node in the linked list
* A doubly linked list gives each node pointers to both the next node and the previous node.
* Unlike an array, a linked list does not provide constant time access to a particular "index" within the list
* The benefit of a linked list is that you can add and remove items from the beginning of the list in constant time

## Creating a Linked List

```java
class Node {
	Node next = null;
	int data;
	
	public Node(int d) {
		data = d;
	}
	
	void appendToTail(int d) {
		Node end = new Node(d);
		Node n = this;
		while (n.next != null) {
			n = n.next;
		}
		n.next = end;
	}
}
```

## Deleting a Node from a Singly Linked List

```java
	Node deleteNode(Node head, int d) {
		Node n = head;
		if (n.data == d) {
			return head.next;
		}
		while (n.next != null) {
			if (n.next.data == d) {
				n.next = n.next.next;
				retrun head;
			}
			n = n.next;
		}
		return head;
	}
```

## The "Runner" Technique

* The runner technique means that you iterate through the linked list with two pointers simultaneously, with one ahead of the other
* The "fast" node might be ahead by a fixed amount, or it might be hopping multiple nodes for each one node that the "slow" node iterates through
* For example, suppose you had a linked list $a_1 -> a_2 -> \cdots -> a_n -> b_1 -> b_2 -> \cdots -> b_n$ and you wanted to rearrage it into $a_1 -> b_1 -> a_2 -> b_2 -> \cdots -> a_n -> b_n$
	* You could have one pointer pl (the fast pointer) move every two elements for every one move that p2 makes
	* When pl hits the end of the linked list, p2 will be at the midpoint
	* Then, move pl back to the front and begin "weaving" the elements
	* On each iteration, p2 selects an element and inserts it after pl

## Recursive Problems

* A number of linked list problems rely on recursion
* Recursive algorithms take at least $O(n)$ space, where $n$ is the depth of the recursive call
* All recursive algorithms can be implemented iteratively, although they may be much more complex

## Interview Questions

* **2.1 Remove Dups:**
  * Write code to remove duplicates from an unsorted linked list.
  * FOLLOW UP
  	* How would you solve this problem if a temporary buffer is not allowed?
* **2.2 Return Kth to Last:**
	* Implement an algorithm to find the kth to last element of a singly linked list.
* **2.3 Delete Middle Node:**
	* Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.
	* EXAMPLE
		* Input: the node c from the linked list a->b->c->d->e->f
		* Result: nothing is returned, but the new linked list looks like a->b->d->e->f
* **2.4 Partition:**
	* Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x. If x is contained within the list, the values of x only need to be after the elements less than x (see below). The partition element x can appear anywhere in the "right partition"; it does not need to appear between the left and right partitions.
	* EXAMPLE
		* Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition=5]
		* Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
* **2.5 Sum Lists:**
	* You have two numbers represented by a linked list, where each node contains a single digit.The digits are stored in reverse order, such that the 1's digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list.
	* EXAMPLE
		* Input: (7-> 1 -> 6) + (5 -> 9 -> 2).That is, 617 + 295.
		* Output: 2 -> 1 -> 9. That is, 912.
  * FOLLOW UP
  	* Suppose the digits are stored in forward order. Repeat the above problem.
  * EXAMPLE
  	* lnput: (6 -> 1 -> 7) + (2 -> 9 -> 5).That is, 617 + 295.
  	* Output: 9 -> 1 -> 2. That is, 912.
* **2.6 Palindrome:**
	* Implement a function to check if a linked list is a palindrome.
* **2.7 Intersection:**
	* Given two (singly) linked lists, determine if the two lists intersect. Return the intersecting node. Note that the intersection is defined based on reference, not value.That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second linked list, then they are intersecting.
* **2.8 Loop Detection:**
	* Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.
  * DEFINITION
  	* Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so as to make a loop in the linked list.
  * EXAMPLE
  	* Input: A -> B -> C -> D -> E -> C [the same C as earlier]
  	* Output: C
