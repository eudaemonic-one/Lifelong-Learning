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