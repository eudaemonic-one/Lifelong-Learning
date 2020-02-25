# Lecture 11 Introduction to Multi-threading and GUIs

## Introduction to Concurrency

* A thread is a thread of execution
* Multiple threads in the same program concurrently
* Threads share the same memory address space
  * Changes made by one thread may be read by others
* Threads vs. Processes
  * Threads are lightweight; processes are heavyweight
  * Threads share address space; processors don't
  * Threads require synchronization; processores don't
  * It's unsafe to kill threads; same to kill processes
* Reasons to use threads
  * Performance needed for blocking activities
  * Performance on multi-core processors
  * Natural concurrency in the real-world
  * Existing multi-threaded, managed run-time environment.

```java
Runnable greeter = new Runnable() {public void run() {doSomething}} // the same as Runnable greeter () -> doSomething()
new Thread(greeter).start()
```

* Synchronization
  * No shared mutable state
    * Don't mutate
    * Don't share
    * Synchronize properly
  * Challenge
    * Safety failure: Not enough synchronization
    * Liveness failure: Too much synchronization