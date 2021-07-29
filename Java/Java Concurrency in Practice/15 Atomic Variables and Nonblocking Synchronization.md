# Chapter 15. Atomic Variables and Nonblocking Synchronization

* Nonblocking algorithms => use low-level atomic machine instructions instead of lcoks.
  * => are used extensively in OS and JVMs for thread and process scheduling, garbage collection, and to implement locks and other concurrent data structures.
  * => offer significant scalability and liveness.
  * => reduce scheduling overhead because they don't block when multiple threads content for the same data.
  * => immune to deadlock and other liveness.
* Atomic variable classes
  * such as `AtomicInteger` and `AtomicReference`.
  * => better volatile variables.
  * => ideal for counters, sequence generators, and statistics gathering while offering better scalability than lock-based alternatives.

## 15.1 Disadvantages of Locking

* JVMs can *only* optimize uncontended lock acquisition and release fairly effectively.
* Volatile variables are lighter-weight and can provide similar visibility guarantees but they cannot be used to construct atomoic compound actions.
* *priority inversion*: a high-priority thread is blocked and waiting for a low-priority thread holding the lock.
* Locking => heavyweight for fine-grained operations.
