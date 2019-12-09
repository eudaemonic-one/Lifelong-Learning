# LeetCode

My solutions to LeetCode algorithms problems, mostly in Python3, Golang and C++.

## Solution Set

* [Solutions in Python3](./Python3/)
* [Solutions in Golang](./Golang/)

## Table of Contents

* [Frequently Used Data Strutures](#frequently-used-data-strutures)
  * [Array/String](#arraystring)
  * [Linked List](#linked-list)
  * [Stack](#stack)
  * [Queue](#queue)
  * [Double-ended Queue](#double-ended-queue)
  * [Tree](#tree)
* [Advanced Data Strutures](#advanced-data-strutures)
  * [Priority Queue](#priority-queue)
  * [Graph](#graph)
  * [Trie](#trie)
  * [Segment Tree](#segment-tree)
  * [Binary Indexed Tree](#binary-indexed-tree)
* [Sort](#sort)
  * [Bubble Sort](#bubble-sort)
  * [Insertion Sort](#insertion-sort)
  * [Merge Sort](#merge-sort)
  * [Quick Sort](#quick-sort)
  * [Topological Sort](#topological-sort)
  * [Counting Sort](#counting-sort)
  * [Radix Sort](#radix-sort)
  * [Buckets and The Pigeonhole Principle](#buckets-and-the-pigeonhole-principle)
* [Recursion &amp; Backtrack](#recursion--backtrack)
  * [Decoding Ways](#decoding-ways)
  * [Strobogrammatic Number](#strobogrammatic-number)
  * [Combinations](#combinations)
  * [N Queen](#n-queen)
* [Depth-first Search &amp; Breath-first Search](#depth-first-search--breath-first-search)
  * [Iterative Implementation of Depth-first Search &amp; Breath-first Search](#iterative-implementation-of-depth-first-search--breath-first-search)
  * [Recursive Implementation of Depth-first Search &amp; Breath-first Search](#recursive-implementation-of-depth-first-search--breath-first-search)
  * [Time Complexity of Depth-first Search &amp; Breath-first Search](#time-complexity-of-depth-first-search--breath-first-search)
  * [Shortest Path Problem](#shortest-path-problem)
* [Dynamic Programming](#dynamic-programming)
  * [Longest Increasing Subsequence](#longest-increasing-subsequence)
  * [Recursive Formula](#recursive-formula)
  * [Dynamic Programming (1): Linear Programming](#dynamic-programming-1-linear-programming)
  * [Unique Paths](#unique-paths)
  * [Dynamic Programming (2): Interval Programming](#dynamic-programming-2-interval-programming)
  * [Longest Palindromic Substring](#longest-palindromic-substring)
  * [Dynamic Programming (3): Constraint programming](#dynamic-programming-3-constraint-programming)
  * [Knapsack Problem](#knapsack-problem)
* [Binary Search &amp; Greedy](#binary-search--greedy)
  * [Binary Search (1): Deterministic Boundary](#binary-search-1-deterministic-boundary)
  * [Find First and Last Position of Element in Sorted Array](#find-first-and-last-position-of-element-in-sorted-array)
  * [Binary Search (2): Undeterministic Boundary](#binary-search-2-undeterministic-boundary)
  * [Lower Bound &amp; Upper Bound](#lower-bound--upper-bound)
  * [Binary Search (3): Rotated Sorted Array](#binary-search-3-rotated-sorted-array)
  * [Search in Rotated Sorted Array](#search-in-rotated-sorted-array)
  * [Binary Search (4): Indefinite Boundary](#binary-search-4-indefinite-boundary)
  * [Log Query](#log-query)
* [Special Topics](#special-topics)
  * [Two Pointers](#two-pointers)
  * [Sliding Window](#sliding-window)
  * [Math](#math)
  * [Bit Manipulation](#bit-manipulation)
  * [Union Find](#union-find)
* [Trending Interview Questions](#trending-interview-questions)
* [Mock Interview](#mock-interview)

## Frequently Used Data Strutures

### Array/String

#### Boyer-Moore Majority Voting Algorithm

The algorithm uses O(1) extra space and O(N) time. It requires exactly 2 passes over the input list.

In the first pass, we need 2 values:

1. A candidate value, initially set to any value.
2. A count, initially set to 0.

For each element in our input list, we first examine the count value. If the count is equal to 0, we set the candidate to the value at the current element. Next, first compare the element's value to the current candidate value. If they are the same, we increment count by 1. If they are different, we decrement count by 1.

At the end of all of the inputs, the candidate will be the majority value if a majority value exists. A second O(N) pass can verify that the candidate is the majority element.

e.g. [229. Majority Element II](https://leetcode.com/problems/majority-element-ii/)

#### Two Pass/Round Trip

To solve Array questions in O(n) time complexity and with constant space complexity. The best way is to iterate the array from left to right and update the status in the first pass. In the second pass, we can set a helper variable initially set to a specific value and then iterate from right to left and update the status correspondingly again.

e.g. [238. Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/)

### Linked List

#### Floyd's Tortoise and Hare

Floyd's algorithm is separated into two distinct phases. In the first phase, it determines whether a cycle is present in the list. If no cycle is present, it returns null immediately, as it is impossible to find the entrance to a nonexistant cycle. Otherwise, it uses the located "intersection node" to find the entrance to the cycle.

Implementation of Floyd's Tortoise and Hare in Python3:

```python
def detectCycle(self, head):
    if head == None:
        return None
    hare, turtle= head, head
    while hare != None:
        turtle = turtle.next
        hare = hare.next
        if hare == None:
            return None
        hare = hare.next
        if hare == turtle:
            turtle = head
            while turtle != hare:
                hare = hare.next
                turtle = turtle.next
            return hare
    return None
```

e.g. [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

e.g. [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

In above example, if we interpret nums such that for each pair of index i and value vi and the next value vj is at index vi, we can reduce this problem to cycle detection, which can be solved by (Floyd's Tortoise and Hare*.

### Stack

### Queue

### Double-ended Queue

*Deque* (*double-ended queue*) pops from/pushes to either side with the same O(1) performance.

It's more handy to **store in the deque indexes instead of elements** since both are used during an array parsing.

e.g. [239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/solution/)

In above example, we can see *deque* is a proper data structure to deal with *Sliding Window* problems (Also, in this case, we can use *Heap* to update the maximum value in the sliding window dynamically). While sliding through the array, each element is processed exactly twice - it's index added and then removed from the deque. Thus, the time complexity of using *deque* is O(n).

### Tree

## Advanced Data Strutures

### Priority Queue

A **heap** is a tree with the property that each node is the minimum-valued node in its subtree. The minimum element in the tree is the root, at index 0.

[Heap's implementation in Golang](https://golang.google.cn/pkg/container/heap/).

e.e. [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

### Graph

### Trie

A **Trie** is a special form of a **Nary tree**. Typically, a trie is used to store strings. Each Trie node represents a string (a prefix). Each node might have several children nodes while the paths to different children nodes represent different characters.

e.g. [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/)

Backtracking (or DFS) is the powerful way to exhaust every possible ways. However, we need to do **pruning** when the current condition shall not exist in the result set. Thus, Trie is a good way to do pruning to search for word/words.

e.g. [212. Word Search II](https://leetcode.com/problems/word-search-ii/)

#### How to represent a Trie

1. First Solution - Array.
2. Second Solution - Map.

#### Insertion in Trie

```text
1. Initialize: cur = root
2. for each char c in target string S:
3.      if cur does not have a child c:
4.          cur.children[c] = new Trie node
5.      cur = cur.children[c]
6. cur is the node which represents the string S
```

#### Search Prefix/Word

```text
1. Initialize: cur = root
2. for each char c in target string S:
3.      if cur does not have a child c:
4.          search fails
5.      cur = cur.children[c]
6. search successes
```

1. If search fails which means that no words start with the target word, the target word is definitely not in the Trie.
2. If search succeeds, we need to check if the target word is only a prefix of words in Trie or it is exactly a word. To solve this problem, you might want to modify the node structure a little bit. *Hint: A boolean flag in each node might work.*

### Segment Tree

### Binary Indexed Tree

## Sort

### Bubble Sort

### Insertion Sort

### Merge Sort

### Quick Sort

### Topological Sort

### Counting Sort

Counting sort operates by counting the number of objects that have each distinct key value, and using arithmetic on those tallies to determine the positions of each key value in the output sequence. Its running time is linear in the number of items and the difference between the maximum and minimum keys, so it is only suitable for direct use in situations where the variation in keys is not significantly greater than the number of items.

e.g. [274. H-Index](https://leetcode.com/problems/h-index/)

In above example, we can see that comparison sorting algorithm has a lower bound of O(nlogn). To achieve better performance, we need non-comparison based sorting algorithms.

### Radix Sort

### Buckets and The Pigeonhole Principle

## Recursion & Backtrack

### Decoding Ways

### Strobogrammatic Number

### Combinations

### N Queen

## Depth-first Search & Breath-first Search

### Iterative Implementation of Depth-first Search & Breath-first Search

### Recursive Implementation of Depth-first Search & Breath-first Search

### Time Complexity of Depth-first Search & Breath-first Search

### Shortest Path Problem

## Dynamic Programming

### Longest Increasing Subsequence

### Recursive Formula

### Dynamic Programming (1): Linear Programming

### Unique Paths

### Dynamic Programming (2): Interval Programming

### Longest Palindromic Substring

### Dynamic Programming (3): Constraint programming

### Knapsack Problem

## Binary Search & Greedy

### Binary Search (1): Deterministic Boundary

### Find First and Last Position of Element in Sorted Array

### Binary Search (2): Undeterministic Boundary

### Lower Bound & Upper Bound

### Binary Search (3): Rotated Sorted Array

### Search in Rotated Sorted Array

### Binary Search (4): Indefinite Boundary

### Log Query

## Special Topics

### Two Pointers

We could keep 2 pointers, one for the start and another for the end of the current subarray, and make optimal moves so as to keep the sum greater than s as well as maintain the lowest size possible.

e.g. [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)

In above example, we could also use Binary Search to crack the problem. We could create an array called sum, and define sums[i] as the sum of first i elements. Then, use Lower Bound method to find the index in sums such that value at that index is not lower than the s+sums[i-1]. If we find the value in sums, compare it with the minimum subarray size and update the result value.

### Sliding Window

### Math

### Bit Manipulation

### Union Find

**Union Find** algorithm is an algorithm that performs two useful operations on **Disjoint Set**:

* Find: Determine which subset a particular element is in. This can be used for determining if two elements are in the same subset.
* Union: Join two subsets into a single subset.

```text
Initialize:
1. set nums[i] as -1 for all elements

Find:
1. if nums[i] is -1 then return i
2. return find(nums, nums[i])

Union:
1. nums[x] = y
```

Union-Find Algorithm can be used to check whether an undirected graph contains cycle or not.

e.g. [261. Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/)

## Trending Interview Questions

### Longest Substring Without Repeating Characters

### Median of Two Sorted Arrays

### Merge k Sorted Lists

### Merge Intervals

### Non-overlapping Intervals

### Alien Dictionary

The key of this problem is that a **Topological Sorting** is possible if and only if the graph contains no directed cycles. Thus, we can *build a graph* (including vertices and edges) and perform a *Depth First Search* (including 4 states: NONEXIST, UNVISITED, VISITING, VISITED).

e.g. [269. Alien Dictionary](https://leetcode.com/problems/alien-dictionary/)

### Basic Calculator

### Regular Expression Matching

### Largest Rectangle in Histogram

### Implement strStr()

### Palindrome Pairs

### Longest Substring with At Least K Repeating Characters

### Trapping Rain Water II

## Mock Interview
