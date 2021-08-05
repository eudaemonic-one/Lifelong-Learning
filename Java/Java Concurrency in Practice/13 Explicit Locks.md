# Chapter 13. Explicit Locks

* Before Java 5.0, `synchronized` and `volatile` are for coordinating access to shared data.
* Java 5.0, `ReentrantLock` => alternative with advanced features for when intrinsic locking proves too limited.

## 13.1 `Lock` and `ReentrantLock`

* `Lock` offers a choice of unconditional, polled, timed, and interruptible lock acquisition, and all lock and unlock operations are explicit.
  * => provide the same memory-visibility semantics as intrinsic locks.

![c0277-01](images/13%20Explicit%20Locks/c0277-01.jpg)

* `ReentrantLock` implements `Lock`, providing the same mutual exclusion and memory-visibility guarantees as `synchronized`.
  * => the same memory semantics as `synchronized` block, and offers reentrant locking semantics like `synchronized`.
* Intrinsic locking's limitations
  * => not possible to interrupt a thread waiting to acquire a lock
  * => not possible to attempt to acquire a lock without being willing to wait for it forever.
  * => must be released in the same block of code in which they are acquired => simplifying coding with exception handling, but impossible to use non-block-structured locking disciplines.
* The canonical form of using a `Lock`.
  * => the lock *must* be released in a `finally` block.
  * => additional `try-catch` or `try-finally` blocks for inconsistent object states.

![c0278-01](images/13%20Explicit%20Locks/c0278-01.jpg)

### 13.1.1 Polled and Timed Lock Acquisition

* Timed and polled locking => probabalistic deadlock avoidance.
  * => regain control => can try again.
* Avoiding lock-ordering deadlock using `tryLock`.
  * Use `tryLock` to attempt to acquire locks, but back off and retry if they cannot be acquired at the same time.

![c0280-01](images/13%20Explicit%20Locks/c0280-01.jpg)

* Locking with a time budget.
  * Supply a timeout corresponding to the remaining time in the budget when calling a blocking method => incorporate exclusive locking to a time-limited activity.

![c0281-01](images/13%20Explicit%20Locks/c0281-01.jpg)

### 13.1.2 Interruptible Lock Acquisition

* Interruptible lock acquisition.
  * => allow locking to be used within cancellable activities.
  * The timed `tryLock` is also responsive to interruption and so can be used when you need both timed and interruptible lock acquisition.

![c0281-02](images/13%20Explicit%20Locks/c0281-02.jpg)

### 13.1.3 Non-block-structured Locking

* Automatic lock release => simplifying analysis and prevents potential coding errors.
* Reducing lock granularity => enhance scalability.
  * Lock striping => allowing different hash chains in a hash-based collection to use different locks.
  * Use a separate lock for each link node, allowing different threads to operate independently on different portions of the list.

## 13.2 Performance Considerations

* A better lock implementation
  * => fewer system calls
  * => forces fewer context switches
  * => initiates less memory-synchronization traffic
* Performance is a moving target; yesterday's benchmark showing that X is faster than Y may already be out of date today.

## 13.3 Fairness

* The `ReentrantLock` constructor
  * => create a *nonfair* lock (the default) => *barging* => threads requesting a lock can jump ahead of the queue of waiting threads if the lock happens to be available when it is requested.
  * => create a *fair* lock
* In most cases, the performance benefits of nonfair locks outweigh the benefits of fair queueing.
  * => Don't pay for fairness if you don't need it.
* Fair locks tend to work best when they are held for a relatively long time or when the mean time between lock requests is relatively long.

## 13.4 Choosing Between Synchronized and ReentrantLock

* Intrinsic locks still have significant advantages over explicit locks.
  * The notation is familiar and compact, and many existing programs using intrinsic locking.
  * `ReentrantLock` is more dangerous if you forget to wrap the `unlock` call in a `finally` block.
* Prefer `synchronized` unless you need advanced features: timed, polled, or interruptible lock acquisition, fair queueing, or non-block-structured locking.
* Future performance improvements are likely to favor `synchronized` over `ReentrantLock`.

## 13.5 Read-write Locks

* `ReadWriteLock`: exposes two `Lock` objects, one for reading and one for writing.
  * => improve performance for frequently accessed read-mostly data structures on multiprocessor systems.

![c0286-01](images/13%20Explicit%20Locks/c0286-01.jpg)

* Implementation options for a `ReadWriteLock`:
  * Release preference.
    * Who should be given preference: readers, writers, or whoever asked first?
  * Reader barging.
    * Allowing readers to barge ahead => enhances concurrency, runs the risk of starving writers.
  * Reentrancy.
    * Are the read and write locks reentrant?
  * Downgrading.
    * If a thread holds the write lock, can it acquire the read lock without releasing the write lock?
  * Upgrading.
    * Can a read lock be upgraded to a write lock in preference to other waiting readers or writers?
    * Most implementations do not support upgrading.
* `ReentrantReadWriteLock` provides reentrant locking semantics for both locks.
* Wrapping a `Map` with a Read-write Lock.
  * `ConcurrentHashMap` provides much better performance.

![c0288-01](images/13%20Explicit%20Locks/c0288-01.jpg)
