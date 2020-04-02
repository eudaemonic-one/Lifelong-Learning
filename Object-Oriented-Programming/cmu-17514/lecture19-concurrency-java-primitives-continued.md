# Lecture 19 Concurrency: Java Primitives, continued

## Challenges of Concurrency

* A liveness problem: poor performance
  * A proposed fix: lock splitting
* A liveness problem: deadlock
  * A possible interleaving of operations
  * The waits-for graph represents dependencies between threads
  * Deadlock has occurred iff the waits-for graph contains a cycle
  * Avoiding deadlock by ordering lock acquisition
* Another subtle problem: The lock object is exposed

## Concurrency and Information Hiding

* Encapsulate an object's state - Easier to implement invariants
  * Encapsulate synchronization - Easier to implement synchronization policy
* Aside: @ThreadSafe @NotThreadSafe @GuardedBy("lock")

## JUnit does not well-support concurrent tests

* Write JUnit test with a false sense of security
* Concurrent clients beware

## Concurrent Programming can be hard to get right

* Invoke `Thread.start`, not `Thread.run`
  * Can be very difficult to diagnose
* This is a severe API design bug
* Thread should not have implemented Runnable
  * This confuses is-a and has-a relationships
  * Thread's `runnable` should have been private
* Thread violates the "Minimize accessibility" principle

