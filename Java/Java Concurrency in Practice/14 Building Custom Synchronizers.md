# Chapter 14. Building Custom Synchronizers

* The class libraries include a number of *state-dependent* classes - those having operations with *state-based preconditions* - such as `FutureTask`, `Semaphore`, and `BlockingQueue`.
* You can build your own synchronizers using the low-level mechanisms provided by the language and libraries, including intrinsic *condition queues*, explicit `Condition` objects, and the `AbstractQueuedSynchronizer` framework.

## 14.1 Managing State Dependence

* State-dependent operations that *block* until the operation can proceed are more convenient and less error-prone than those that simply fail.
  * The built-in condition queue mechanism enables threads to block until an object has entered a state that allows progress and to wake blocked threads when they may be able to make further progress.
* Structure of blocking state-dependent actions.

![c0292-01](images/14 Building Custom Synchronizers/c0292-01.jpg)

* A bounded buffer provides `put` and `take` operations, each of which has preconditions.

![c0293-01](images/14 Building Custom Synchronizers/c0293-01.jpg)

### 14.1.1 Example: Propagating Precondition Failure to Callers

* `GrumpyBoundedBuffer`: Bounded buffer that balks when preconditions are not met.
  * Exceptions are supposed to be for exceptional conditions => the caller must be prepared to catch exceptions and possibly retry for every buffer operation.
  * However, "Buffer is full" is not an exceptional condition for a bounded buffer.

![c0294-01](images/14 Building Custom Synchronizers/c0294-01.jpg)

* Client logic for calling `GrumpyBoundedBuffer`.
  * spin waiting versus oversleep.

![c0294-02](images/14 Building Custom Synchronizers/c0294-02.jpg)

### 14.1.2 Example: Crude Blocking by Polling and Sleeping

* `SleepyBoundedBuffer`: Bounded buffer using crude blocking.
  * => encapsulating precondition management and simplifying using the buffer.
  * Choosing the sleep granularity => responsiveness versus CPU usage.
  * => the caller should deal with `InterruptedException` => to support cancellation mechanism.
* It would be nice to have a way of suspending a thread but ensuring that is awakened promptly when a certain condition becomes true =>  *condition queues*.

![c0296-01](images/14 Building Custom Synchronizers/c0296-01.jpg)

### 14.1.3 Condition Queues to the Rescue

* *condition queue*: a group of threads called *wait set* waiting for a specific condition to become true.
* Each object can act as a condition queue, and the `wait`, `notify`, and `notifyAll` methods in `Object` constitute the API for intrinsic condition queues.
  * `Object.wait` atomically releases the lock and asks the OS to suspend the current thread, allowing other threads to acquire the lock and therefore modify the object state. Upon waking, it reacquires the lock before returning.
* `BoundedBuffer`: Bounded buffer using condition queues.

![c0298-01](images/14 Building Custom Synchronizers/c0298-01.jpg)
