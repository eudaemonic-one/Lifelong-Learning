# Technical Questions

## How to Prepare

* Try to solve the problem on your own
* Write the code on paper
* Test your code - on paper
* Type your paper code as-is into a computer

## What You Need To Know

### Core Data Structures, Algorithms, and Concepts

| Data Structures        | Algorithms           | Concepts                |
| ---------------------- | -------------------- | ----------------------- |
| Linked List            | Breadth-First Search | Bit Manipulation        |
| Trees, Tries, & Graphs | Depth-First Search   | Memory (Stack vs. Heap) |
| Stacks & Queues        | Binary Search        | Recursion               |
| Heaps                  | Merge Sort           | Dynamic Programming     |
| Vectors / ArrayLists   | Quick Sort           | Big O Time & Space      |
| Hash Tables            |                      |                         |

#### Power of 2 Table

| Power of 2 | Exact Value (X)   | Approx. Value | X Bytes into MB, GB, etc |
| ---------- | ----------------- | ------------- | ------------------------ |
| 7          | 128               |               |                          |
| 8          | 256               |               |                          |
| 10         | 1024              | 1 thousand    | 1KB                      |
| 16         | 65,536            |               | 64KB                     |
| 20         | 1,048,576         | 1 million     | 1MB                      |
| 30         | 1,073,741,824     | 1 billion     | 1GS                      |
| 32         | 4,294,967,296     |               | 4GB                      |
| 40         | 1,099,511,627,776 | 1 trillion    | 1TB                      |

## Walking Through a Problem

* **Listen**
  * Pay very close attention to any information in the problem description
  * You probably need it all for an optimal algorithm
* **Example**
  * Most examples are too small or are special cases
* **Brute Force**
  * Get a brute-force solution as soon as possible
  * Don't worry about developing an efficient algorithm yet
  * State a naive algorithm and its runtime, then optimize from there
  * Don't code yet though!
* **Test**
  * Test in this order:
    * Conceptual test
    * Unusual or non-standard code
    * Hot spots, like **arithmetic** and **null** nodes
    * Small test cases
    * Special cases and edge cases
* **Optimize**
  * Walk through your brute force with BUD optimization or try some of these ideas:
    * Look for any unused info
    * Solve it manually on an example, then reverse engineer your thought process
    * Solve it "incorrectly" and then think about why the algorithm fails
      * Can you fix those issues?
    * Make a time vs. space tradeoff
      * Hash tables are especially useful!
  * **BUD Optimization**
    * Bottlenecks
    * Unnecessary Work
    * Duplicated Work
* **Implement**
  * Your goal is to write beautiful code
  * Modularize your code from the beginning and refactor to clean up anything that isn't beautiful
* **WalkThrough**
  * Now that you have an optimal solution, walk through your approach in detail
  * Make sure you understand each detail before you start coding
  * **Keep talking!**
    * Your interviewer wants to hear how you approach the problem

### Listen Carefully

* Listen carefully to the problem, and be sure that you've mentally recorded any **unique** information in the problem
* For example, suppose a question starts with one of the following lines
  * "Given two arrays that are sorted, find..."
    * You probably need to know that the data is sorted
    * The optimal algorithm for the sorted situation is probably different than the optimal algorithm for the unsorted situation
  * "Design an algorithm to be run repeatedly on a server that..."
    * The server/to-be-run-repeatedly situation is different from the run-once situation
    * Perhaps this means that you cache data
* Your first algorithm doesn't need to use the information
* But if you find yourself stuck, or you're still working to develop something more optimal, ask yourself if you've used all the information in the problem
* You might even find it useful to write the pertinent information on the whiteboard

### Draw an Example

* An example can dramatically improve your ability to solve an interview question
* You want to create an example that is:
  * Specific
    * It should use real numbers or strings (if applicable to the problem)
  * Sufficiently large
  * Not a special case

### State a Brute Force

* It's okay that this initial solution is terrible
* Explain what the space and time complexity is, and then dive into improvements
* It's a starting point for optimiza- tions, and it helps you wrap your head around the problem

### Optimize

* Look for any unused information
* Use a fresh example
  * Sometimes, just seeing a different example will unclog your mind or help you see a pattern in the problem
* Solve it "incorrectly."
* Make time vs. space tradeoff
* Precompute information
* Use a hash table
* Think about the best conceivable runtime

### Walk Through

* After you've nailed down an optimal algorithm, don't just dive into coding
* Whiteboard coding is slow - very slow
* So is testing your code and fixing it
* As a result, you need to make sure that you get it as close to "perfect" in the beginning as possible
* if you don't understand exactly what you're about to write, you'll struggle to code it

### Implement

* Start coding in the far top left corner of the whiteboard
* Avoid "line creep" (where each line of code is written an awkward slant
* Remember that you only have a short amount of code to demonstrate that you're a great developer
* Beautiful code means:
  * Modularized code
  * Error checks
  * Use other classes/structs where appropriate
  * Good variable names

### Test

* **Start with a "conceptual" test**
  * A conceptual test means just reading and analyzing what each line of code does
* **Weird looking code**
* **Hot spots**
  * Base cases in recursive code
  * Integer division
  * Null nodes in binary trees
  * The Start and end of iteration through a linked list
* **Small test cases**
* **Special cases**
  * Test your code against null or single element values, the extreme cases, and other special cases

### Optimize & Solve Technique #1: Look for BUD

* **BUD**
  * **Bottlenecks**
  * **Unnecessary work**
  * **Duplicated work**
* These are three of the most common things that an algorithm can "waste" time doing
* If it's still not optimal, you can repeat this approach on your current best algorithm

### Optimize & Solve Technique #2: DIY (Do It Yourself)

* Therefore, when you get a question, try just working it through **intuitively** on a real example
* Use a nice, big example and intuitively-manually, that is-solve it for the **specific example**
* Then, afterwards, think hard about how you solved it
* **Reverse engineer your own approach**
* Be particularly aware of any "optimizations" you intuitively or automatically made
* Example: Given a smaller strings and a bigger string b, design an algorithm to find all permuta­ tions of the shorter string within the longer one.Print the location of each permutation
  * If you're like most candidates, you probably thought of something like: Generate all permutations ofs and then look for each in b
    * Since there are $S!$ permutations, this will take $O(S ! * B)$ time, where $S$ is the length of s and $B$ is the length of b
  * Walk through b and look at sliding windows of 4 characters (since s has length 4) and check if each window is a permutation of s
  * Walk through b. Every time you see a character in s, check if the next four(the length ofs) characters are a permutation of s

### Optimize & Solve Technique #3: Simplify and Generalize

* First we simplify or tweak some **constraint**, such as the data type
* Then, we solve this new simplified version of the problem
* Finally, once we have an algorithm for the simplified problem, we try to adapt it for the more complex version
* Example: A ransom note can be formed by cutting words out of a magazine to form a new sentence. How would you figure out if a ransom note (represented as a string) can be formed from a given magazine (string)?
  * To simplify the problem, we can modify it so that we are cutting characters out of a magazine instead of whole words
  * We can solve the simplified ransom note problem with characters by simply creating an array and counting the characters
  * Each spot in the array corresponds to one letter
  * When we generalize the algorithm, we do a very similar thing
    * This time, rather than creating an array with character counts, we create a hash table that maps from a word to its frequency

### Optimize & Solve Technique #4: Base Case and Build

* With Base Case and Build, we solve the problem first for a base case (e.g., n = 1) and then try to build up from there
* When we get to more complex/interesting cases (often n = 3 or n = 4), we try to build those using the prior solutions
* **Base Case and Build algorithms often lead to natural recursive algorithms**
* Example: Design an algorithm to print all permutations of a string. For simplicity, assume all characters are unique

```text
Case "a" --> {"a"}
Case "ab" --> {"ab", "ba"}
Case "abc" --> ?

P("abc") = insert "c" into all locations of all strings in P("ab")  P("abc") = insert "c" into all locations of all strings in {"ab","ba"}
P("abc") = merge({"cab", ""acb", "abc"}, {"cba", abca", bac"})
P("abc") = {"cab", "acb", "abc", "cba", "bca", bac"}
```

### Optimize & Solve Technique #5: Data Structure Brainstorm

* This approach is certainly hacky, but it often works
* We can simply run through a list of data structures and try to apply each one
* Example: Numbers are randomly generated and stored into an (expanding) array. How would you keep track of the median?
* Our data structure brainstorm might look like the following:
  * Linked list? Probably not. Linked lists tend not to do very well with accessing and sorting numbers
  * Array? Maybe, but you already have an array. Could you somehow keep the elements sorted? That's probably expensive. Let's hold off on this and return to it if it's needed
  * Binary tree? This is possible, since binary trees do fairly well with ordering. In fact, if the binary search tree is perfectly balanced, the top might be the median. But, be careful-if there's an even number of elements, the median is actually the average of the middle two elements. The middle two elements can't both be at the top. This is probably a workable algorithm, but let's come back to it
  * Heap? A heap is really good at basic ordering and keeping track of max and mins. This is actually interesting-if you had two heaps, you could keep track of the bigger half and the smaller half of the elements. The bigger half is kept in a min heap, such that the smallest element in the bigger half is at the root. The smaller half is kept in a max heap, such that the biggest element of the smaller half is at the root. Now, with these data structures, you have the potential median elements at the roots. If the heaps are no longer the same size, you can quickly "rebalance" the heaps by popping an element off the one heap and pushing it onto the other

## Best Conceivable Runtime (BCR)

* The best conceivable runtime is, literally, the best runtime you could conceive of a solution to a problem having
* You can easily prove that there is no way you could beat the BCR
* The Best Conceivable Runtime is for a problem and is largely a function of the inputs and outputs
  * It has no particular connection to a specific algorithm
* BCR can be useful since **we can use the runtimes to get a "hint" for what we need to reduce**
  * If we imagine our current algorithm's runtime as $O(N \times N)$, then getting to $O(N)$ or $O(N \times log N)$ might mean reducing that second $O(N)$ in the equation to $O(1)$ or $O(log N)$
  * One of the tips there suggests **precomputing** or doing **upfront work**
  * Any upfront work we do in $O(N)$ time is a freebie
    * It won't impact our runtime
* This is another place where BCR can be useful
  * **Any work you do that's less than or equal to the BCR is "free"** in the sense that it won't impact your runtime
  * You might want to eliminate it even­ tually, but it's not a top priority just yet
* This is another place where BCR is useful
  * **It tells us that we're "done" in terms of optimizing the runtime**, and we should therefore turn our efforts to the **space complexity**
* This is another way we can useBCR
  * If you ever **reach the BCR and have $O(1)$ additional space**, then you know that **you can't optimize the big O time or space**

## Handling Incorrect Answers

* First, responses to interview questions shouldn't be thought of as "correct" or "incorrect"
  * Rather, it's about how optimal their final solution was, how long it took them to get there, how much help they needed, and how clean was their code
* Second, your performance is evaluated in comparison to other candidates
* Third, many - possibly most - questions are too difficult to expect even a strong candidate to immediately spit out the optimal algorithm

## When You've Heard a Question Before

* If you've heard a question before, admit this to your interviewer
* Additionally, your interviewer may find it highly dishonest if you don't reveal that you know the question

## The "Perfect" Language for Interviews

* At many of the top companies, interviewers aren't picky about languages
* They're more interested in how well you solve the problems than whether you know a specific language
* Other companies though are more tied to a language and are interested in seeing how well you can code in a particular language
* You should keep in mind the following:
  * Prevalence
  * Language Readability
  * Potential Problems
  * Verbosity
  * Ease of Use

## What Good Coding Looks Like

* Broadly speaking, good code has the following properties:
  * Correct
  * Efficient
  * Simple
  * Readble
  * Maintainable
* **Use Data Structures Generously**
  * Suppose you were asked to write a function to add two simple mathematical expressions which are of the form $Ax^a + Bx^b + \cdots$ (where the coefficients and exponents can be any positive or negative real number)
  * A bad implementation would be to store the expression as a single array of doubles, where the kth element corresponds to the coefficient of the $x^k$ term in the expression
    * This structure is problematic because it could not support expressions with negative or non-integer exponents
    * It would also require an array of 1000 elements to store just the expression $x^{1000}$
  * A slightly less bad implementation would be to store the expression as a set of two arrays, `coefficients` and `exponents`
    * Under this approach, the terms of the expression are stored in any order, but "matched" such that the ith term of the expression is represented by $coefficients[i] * x^{exponents[i]}$
  * A good implementation for this problem is to design your own data structure for the expression
* **Appropriate Code Reuse**
* **Modular**
  * Writing modular code means separating isolated chunks of code out into their own methods
  * This helps keep the code more maintainable, readable, and testable
* **Flexible and Robust**
  * Writing flexible, general-purpose code means using variables instead of hard-coded values or **using templates/ generics** to solve a problem
  * Of course, there is a **limit**
    * If the solution is much more complex for the general case, and it seems unneces­ sary at this point in time, it may be better just to implement the simple, expected case
* **Error Checking**
  * One sign of a careful coder is that she doesn't make assumptions about the input
  * Instead, she validates that the input is what it should be, either through ASSERT statements or if-statements
  * Checks like these are critical in production code and, therefore, in interview code as well

## Don't Give Up!

* Do you rise to a challenge, or do you shrink back in fear?
* It's important that you step up and eagerly meet a tricky problem head-on
* For extra "points", show **excitement** about solving hard problems
