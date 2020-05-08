# Lecture 20 Concurrency primitives, libraries, and design patterns

## Policies for Thread Safety

* Thread-confined state - mutate but don't share
  * Stack-confined
    * Primitive local variables are *never* shared between threads
    * Fast and cheap
  * Unshared object references
    * The thread that creates an object must take action to share (“publish”)
  * Thread-local variables
    * Shared object with a separate value for each thread
    * Rarely needed but invaluable (e.g., for user ID or transaction ID)
* Shared read-only state - share but don't mutate
  * Immutable data is always safe to share
* Shared thread-safe - object synchronizes itself internally
  * Thread-safe objects that perform internal synchronization
  * Better use ones from `java.util.concurrent`
* Shared guarded - client synchronizes objects externally
  * Shared objects that must be locked by user
  * Can be error prone: burden is on user
  * High concurrency can be difficult to achieve
    * Lock granularity is the entire object
  * You’re generally better off avoiding guarded objects

## A Primitive for Cooperation: wait/notify

* State (fields) are guarded by a lock

* Sometimes, a thread can’t proceed till state is right

  * So it waits with `wait`
  * Automatically drops lock while waiting

* Thread that makes state right wakes waiting thread(s) with `notify`

* **Never invoke wait outside a loop!**

* ```java
  synchronized (obj) {
    while (<condition does not hold>) {
      obj.wait();
    }
  ... // Perform action appropriate to condition
  }
  ```

* Why can a thread wake from a wait when condition does not hold?

  * Another thread can slip in between notify & wake
  * Another thread can invoke notify accidentally or maliciously when condition does not hold
  * Notifier can be liberal in waking threads
    * Using `notifyAll` is good practice, but can cause extra wakeups
  * Waiting thread can wake up without a notify(!)
    * Known as a spurious wakeup

## Defining Thread-safe Objects

* Identify variables that represent the object's state
* Identify invariants that constrain the state variables
* Establish a policy for maintaining invariants

### A Toy Example: Read-write Locks

```java
@ThreadSafe public class RwLock {
	/** Number of threads holding lock for read. */
  @GuardedBy("this") // Intrinsic lock on RwLock object
  private int numReaders = 0;
  
	/** Whether lock is held for write. */
	@GuardedBy("this")
	private boolean writeLocked = false;
	public synchronized void readLock() throws InterruptedException {
    while (writeLocked) {
			wait();
    }
    numReaders++;
  }
  
  public synchronized void writeLock() throws InterruptedException {
    while (numReaders != 0 || writeLocked) {
			wait();
    }
		writeLocked = true;
  }
  
	public synchronized void unlock() {
    if (numReaders > 0) {
			numReaders--;
		} else if (writeLocked) {
			writeLocked = false;
    } else {
			throw new IllegalStateException("Lock not held");
    }
    notifyAll(); // Wake any waiters
  }
}
```

### Advice for Building Thread-safe Objects

* **Do as little as possible in synchronized region: get in, get out**
  * Obtain lock
  * Examine shared data
  * Transform as necessary
  * Drop the lock
* If you must do something slow, move it outside the synchronized region
* Generally, avoid `wait`/`notify`
  * `java.util.concurrent` provides better alternatives

#### Documentation

* Document a class’s thread safety guarantees for its clients
* Document a class’s synchronization policy for its maintainers
* Use `@ThreadSafe`, `@GuardedBy` annotations

## Java libraries for concurrency (`java.util.concurrent`)

### `java.util.concurrent` is BIG

* Atomic variables: `java.util.concurrent.atomic`
  * Support various atomic read-modify-write ops
* Concurrent collections
  * Shared maps, sets, lists
* Data exchange collections
  * Blocking queues,deques,etc.
* Executor framework
  * Tasks, futures, thread pools, completion service, etc.
* Synchronizers
  * Semaphores, cyclic barriers, countdown latches, etc.
* Locks: `java.util.concurrent.locks`
  * Read-write locks, conditions, etc.

### Overview of `java.util.concurrent.atomic`

* `Atomic{Boolean,Integer,Long}`
  * Boxed primitives that can be updated atomically
* `AtomicReference<T>`
  * Object reference that can be updated atomically
* `Atomic{Integer,Long,Reference}Array`
  * Array whose elements may be updated atomically
* `Atomic{Integer,Long,Reference}FieldUpdater`
  * Reflection-based utility enabling atomic updates to volatile fields
* `LongAdder`,`DoubleAdder`
  * Highlyconcurrentsums
* `LongAccumulator`, `DoubleAccumulator`
* Generalization of adder to arbitrary functions (max, min, etc.)

### Concurrent collections

| Unsynchronized | Concurrent              |
| -------------- | ----------------------- |
| `HashMap`      | `ConcurrentHashMap`     |
| `HashSet`      | `ConcurrentHashSet`     |
| `TreeMap`      | `ConcurrentSkipListMap` |
| `TreeSet`      | `ConcurrentSkipListSet` |

* You **can’t** prevent concurrent use of a concurrent collection

* This works for synchronized collections...

  * ```java
    Map<String, String> syncMap = Collections.synchronizedMap(new HashMap<>()); synchronized(syncMap) {
    	if (!syncMap.containsKey("foo"))
        syncMap.put("foo", "bar");
    }
    ```

* But **not** for concurrent collections

  * They do their own internal synchronization
  * **Never synchronize on a concurrent collection!**

* Instead, use **atomic read-modify-write methods**

  * ```java
    V putIfAbsent(K key, V value);
    boolean remove(Object key, Object value);
    V replace(K key, V value);
    boolean replace(K key, V oldValue, V newValue);
    V compute(K key, BiFunction<...> remappingFn);
    V computeIfAbsent(K key, Function<...> mappingFn);
    V computeIfPresent (K key, BiFunction<...> remapFn);
    V merge(K key, V value, BiFunction<...> remapFn);
    ```

* Concurrent collection example: canonicalizing map

  * ```java
    private final ConcurrentMap<T,T> map = new ConcurrentHashMap<>();
    
    public T intern(T t) {
    	String previousValue = map.putIfAbsent(t, t);
      return previousValue == null ? t : previousValue;
    }
    ```

* java.util.concurrent.ConcurrentHashMap`

  * Uses **many** techniques used to achieve high concurrency
  * The simplest of these is *lock striping*
    * Multiple locks, each dedicated to a region of hash table

### Data Exchange Collections Summary

* Hold elements for processing by another thread (producer/consumer)
* `BlockingQueue` – Supports blocking ops
  * `ArrayBlockingQueue`, `LinkedBlockingQueue`
  * `PriorityBlockingQueue`, `DelayQueue`
  * `SynchronousQueue`
* `BlockingDeque` – Supports blocking ops
  * `LinkedBlockingDeque`
* `TransferQueue` – `BlockingQueue` in which producers may wait for consumers to receive elements
  * `LinkedTransferQueue`

#### Summary of `BlockingQueue` Methods

|         | Throws exception | Special value | Blocks | Times out            |
| ------- | ---------------- | ------------- | ------ | -------------------- |
| Insert  | add(e)           | offer(e)      | put(e) | offer(e, time, unit) |
| Remove  | remove()         | poll()        | take() | poll(time, unit)     |
| Examine | element()        | peek()        | n/a    | n/a                  |

### Executor Framework Overview

* Flexible interface-based task execution facility
* Key abstractions
  * `Runnable`, `Callable<T>` - kinds of tasks
* `Executor` – thing that executes tasks
* `Future<T>` – a promise to give you a `T`
* Executor service – Executor that
  * Lets you manage termination
  * Can produce Future instances
* Executors
  * `Executors.newSingleThreadExecutor()`
    * A single background thread
  * `newFixedThreadPool(int nThreads)`
    * A fixed number of background threads
  * `Executors.newCachedThreadPool()`
    * Grows in response to demand
* A very simple (but useful) executor service example
  * Background execution of a long-lived worker thread
    * To start the worker thread:
      * `ExecutorService executor = Executors.newSingleThreadExecutor();`
    * To submit a task for execution:
      * `executor.execute(runnable);`
    * To terminate gracefully:
      * `executor.shutdown();`
* Other things you can do with an executor service
  * Wait for a task to complete
    * `Foo foo = executorSvc.submit(callable).get();`
  * Wait for any or all of a collection of tasks to complete
    * `invoke{Any,All}(Collection<Callable<T>> tasks)`
  * Retrieve results as tasks complete
    * `ExecutorCompletionService`
  * Schedule tasks for execution a time in the future
    * `ScheduledThreadPoolExecutor`

### Overview of Synchronizers

* `CountDownLatch`
  * One or more threads to wait for others to count down
* `CyclicBarrier`
  * a set of threads wait for each other to be ready
* `Semaphore`
  * Like a lock with a maximum number of holders (“permits”)
* `Phaser` – Cyclic barrier on steroids
* `AbstractQueuedSynchronizer` – roll your own!

### Overview of `java.util.concurrency.locks`

* `ReentrantReadWriteLock`
  * Shared/Exclusive mode locks with tons of options
* `ReentrantLock`
* Condition
  * `wait`/`notify`/`notifyAll` with multiple wait sets per object
* `AbstractQueuedSynchronizer`
  * Skeletal implementation of locks relying on FIFO wait queue
* `AbstractOwnableSynchronizer`, `AbstractQueuedLongSynchronizer`
  * Fancier skeletal implementations

## The fork-join Pattern

```text
if (my portion of the work is small)
    do the work directly
else
    split my work into pieces
    recursively process the pieces
```

* Master Thread
  * Parallel Task1
  * Parallel Task2
  * Parallel TaskN

### ForkJoinPool: executor service for ForkJoinTask

* Dynamic, fine-grained parallelism with recursive task splitting

```java
class SumOfSquaresTask extends RecursiveAction {
  final long[] a;
  final int lo, hi;
  long sum;
  SumOfSquaresTask(long[] array, int low, int high) {
		a = array;
    lo = low;
    hi = high;
  }
  
	protected void compute() {
    if (hi - lo < THRESHOLD) {
			for (int i = l; i < h; ++i)
        sum += a[i] * a[i];
		} else {
			int mid = (lo + hi) >>> 1;
      SumOfSquaresTask left = new SumOfSquaresTask(a, lo, mid);
      left.fork(); // pushes task
      SumOfSquaresTask right = new SumOfSquaresTask(a, mid, hi);
			right.compute();
      right.join(); // pops/runs or helps or waits
      sum = left.sum + right.sum;
		}
  }
}
```