# Chapter 1. Introduction

* Developing => Testing => Debugging multithreaded programs
* Concurrency features offered by the platform *versus* How developers think about concurrency in their programs
* Fundamentals (concurrency and thread safety) => Structuring Concurrent Applications (exploit threads) => Liveness, Performance, and Testing => Advanced Topics

## 1.1 A (Very) Brief History of Concurrency

* Processes: isolated, independently executing programs.
  * OS: allocating resources such as memory, file handles, security credentials.
  * Processes communicate through sockets, signal handlers, shared memory, semaphores, and files.
* Motivation: Resource utilization, Fairness, Convenience.
* Threads: allow multiple control flow within a process, share process-wide resources, hold own program counter, stack, and local variables => *lightweight processes* => basic units of scheduling => execute simultaneously and asynchronously with respect to one another => finer-grained data sharing + data race.

## 1.2 Benefits of Threads

* **Benefits of threads:**
  * improve the performance of complex applications.
  * easy to model asynchronous workflows.
  * simplify code structure and thus easier to write, read, and maintain.
  * improve the responsiveness of GUI applications.
  * improve resource utilization and throughput of server applications.
  * simplify the implementation of the JVM.
* **Exploiting Multiple Processors.**
  * => better resource utilization and throughput on single/multiple processor systems.
* **Simplicity of Modeling.**
  * scheduling: interleaved operations, asynchronous I/O, and resource waits.
  * asynchronous workflow => a number of simpler, synchronous workflows running in separate threads.
* **Simplified Handling of Asynchronous Events.**
  * multiplexed I/O: Unix `select` and `poll` system calls, `java.nio` for nonblocking I/O.
* **More Responsive User Interfaces.**
  * main event loop => event dispatch thread (EDT).
  * {long-running tasks} => {separate threads} => free the event thread to process UI events => UI more responsive.

## 1.3 Risks of Threads

* **Safety Hazards.**
  * absence of synchronization => unpredictable ordering of operations => *race condition*.
  * `@ThreadSafe`, `@NotThreadSafe`, `@Immutable`.
  * Java provides synchronization => coordinate access to shared variables.
* **Liveness Hazards.**
  * liveness failure: unable to make forward progress.
  * e.g., inadvertent infinite loop, deadlock, starvation, livelock.
* **Performance Hazards.**
  * performance issues: poor service time, responsiveness, throughput, resource consumption, or scalability.
  * runtime overhead: context switches, synchronization mechanisms.