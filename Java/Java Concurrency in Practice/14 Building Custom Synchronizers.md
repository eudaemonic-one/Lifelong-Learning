# Chapter 14. Building Custom Synchronizers

* The class libraries include a number of *state-dependent* classes - those having operations with *state-based preconditions* - such as `FutureTask`, `Semaphore`, and `BlockingQueue`.
* You can build your own synchronizers using the low-level mechanisms provided by the language and libraries, including intrinsic *condition queues*, explicit `Condition` objects, and the `AbstractQueuedSynchronizer` framework.
