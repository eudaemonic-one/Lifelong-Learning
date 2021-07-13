# Chapter 11. Concurrency

## Item 78: Synchronize access to shared mutable data

* “The `synchronized` keyword ensures that only a single thread can execute a method or block at one time.”
  * “Many programmers think of synchronization solely as a means of *mutual exclusion*, to prevent an object from being seen in an inconsistent state by one thread while it’s being modified by another. In this view, an object is created in a consistent state (Item 17) and locked by the methods that access it.”
  * “These methods observe the state and optionally cause a *state transition*, transforming the object from one consistent state to another.”
* “Without synchronization, one thread’s changes might not be visible to other threads. Not only does synchronization prevent threads from observing an object in an inconsistent state, but it ensures that each thread entering a synchronized method or block sees the effects of all previous modifications that were guarded by the same lock.”
* “The language specification guarantees that reading or writing a variable is *atomic* unless the variable is of type `long` or `double` [JLS, 17.4, 17.7].”
* **“Synchronization is required for reliable communication between threads as well as for mutual exclusion.”**
* “Consider the task of stopping one thread from another. The libraries provide the `Thread.stop` method, but this method was deprecated long ago because it is inherently unsafe—its use can result in data corruption.”
  * “Do not use `Thread.stop`.”
  * “A recommended way to stop one thread from another is to have the first thread poll a `boolean` field that is initially `false` but can be set to `true` by the second thread to indicate that the first thread is to stop itself. Because reading and writing a `boolean` field is atomic, some programmers dispense with synchronization when accessing the field:”

```java
// Broken! - How long would you expect this program to run?
public class StopThread {
    private static boolean stopRequested;

    public static void main(String[] args)
            throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i = 0;
            while (!stopRequested)
                i++;
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        stopRequested = true;
    }
}
```

* “In the absence of synchronization, it’s quite acceptable for the virtual machine to transform this code:”

```java
if (!stopRequested)
    while (true)
        i++;
```

* “The result is a *liveness failure*: the program fails to make progress. One way to fix the problem is to synchronize access to the `stopRequested` field.”


```java
// Properly synchronized cooperative thread termination
public class StopThread {
    private static boolean stopRequested;

    private static synchronized void requestStop() {
        stopRequested = true;
    }

    private static synchronized boolean stopRequested() {
        return stopRequested;
    }

    public static void main(String[] args)
            throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i = 0;
            while (!stopRequested())
                i++;
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        requestStop();
    }
}
```

* “Note that both the write method (`requestStop`) and the read method (`stop-Requested`) are synchronized. It is not sufficient to synchronize only the write method!”
* “Synchronization is not guaranteed to work unless both read and write operations are synchronized.”
* “The locking in the second version of `StopThread` can be omitted if `stopRequested` is declared volatile. ”
  * “While the `volatile` modifier performs no mutual exclusion, it guarantees that any thread that reads the field will see the most recently written value:”

```java
// Cooperative thread termination with a volatile field
public class StopThread {
    private static volatile boolean stopRequested;

    public static void main(String[] args)
            throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i = 0;
            while (!stopRequested)
                i++;
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        stopRequested = true;
    }
}
```

* “You do have to be careful when using volatile.”

```java
// Broken - requires synchronization!
private static volatile int nextSerialNumber = 0;

public static int generateSerialNumber() {
    return nextSerialNumber++;
}
```

* “The problem is that the increment operator (`++`) is not atomic. It performs two operations on the `nextSerialNumber` field: first it reads the value, and then it writes back a new value, equal to the old value plus one.”
* “One way to fix `generateSerialNumber` is to add the `synchronized` modifier to its declaration. This ensures that multiple invocations won’t be interleaved and that each invocation of the method will see the effects of all previous invocations.”
  * “To bulletproof the method, use `long` instead of `int`, or throw an exception if `nextSerialNumber` is about to wrap.”
* “Better still, follow the advice in Item 59 and use the class `AtomicLong`, which is part of `java.util.concurrent.atomic`.”
  * “This package provides primitives for lock-free, thread-safe programming on single variables.”


```java
// Lock-free synchronization with java.util.concurrent.atomic
private static final AtomicLong nextSerialNum = new AtomicLong();

public static long generateSerialNumber() {
    return nextSerialNum.getAndIncrement();
}
```

* “The best way to avoid the problems discussed in this item is not to share mutable data.”
  * “In other words, **confine mutable data to a single thread**.”
* “It is acceptable for one thread to modify a data object for a while and then to share it with other threads, synchronizing only the act of sharing the object reference. Other threads can then read the object without further synchronization, so long as it isn’t modified again.”
  * “Such objects are said to be *effectively immutable* [Goetz06, 3.5.4]. Transferring such an object reference from one thread to others is called *safe publication* [Goetz06, 3.5.3].”
  * “There are many ways to safely publish an object reference: you can store it in a static field as part of class initialization; you can store it in a volatile field, a final field, or a field that is accessed with normal locking; or you can put it into a concurrent collection (Item 81).”
* “In summary, **when multiple threads share mutable data, each thread that reads or writes the data must perform synchronization**.”
  * “In the absence of synchronization, there is no guarantee that one thread’s changes will be visible to another thread. The penalties for failing to synchronize shared mutable data are liveness and safety failures. These failures are among the most difficult to debug. They can be intermittent and timing-dependent, and program behavior can vary radically from one VM to another. If you need only inter-thread communication, and not mutual exclusion, the volatile modifier is an acceptable form of synchronization, but it can be tricky to use correctly.”


## Item 79: Avoid excessive synchronization

* “Depending on the situation, excessive synchronization can cause reduced performance, deadlock, or even nondeterministic behavior.”
* **“To avoid liveness and safety failures, never cede control to the client within a synchronized method or block.”**
  * “In other words, inside a synchronized region, do not invoke a method that is designed to be overridden, or one provided by a client in the form of a function object (Item 24).”

```java
// Broken - invokes alien method from synchronized block!
public class ObservableSet<E> extends ForwardingSet<E> {
    public ObservableSet(Set<E> set) { super(set); }

    private final List<SetObserver<E>> observers
            = new ArrayList<>();

    public void addObserver(SetObserver<E> observer) {
        synchronized(observers) {
            observers.add(observer);
        }
    }

    public boolean removeObserver(SetObserver<E> observer) {
        synchronized(observers) {
            return observers.remove(observer);
        }
    }

    private void notifyElementAdded(E element) {
        synchronized(observers) {
            for (SetObserver<E> observer : observers)
                observer.added(this, element);
        }
    }

    @Override public boolean add(E element) {
        boolean added = super.add(element);
        if (added)
            notifyElementAdded(element);
        return added;
    }

    @Override public boolean addAll(Collection<? extends E> c) {
        boolean result = false;
        for (E element : c)
            result |= add(element);  // Calls notifyElementAdded
        return result;
    }
}
```

* “Observers subscribe to notifications by invoking the `addObserver` method and unsubscribe by invoking the `removeObserver` method. In both cases, an instance of this callback interface is passed to the method.”

```java
@FunctionalInterface public interface SetObserver<E> {
    // Invoked when an element is added to the observable set
    void added(ObservableSet<E> set, E element);
}
```

* “Suppose we replace the `addObserver` call with one that passes an observer that prints the `Integer` value that was added to the set and removes itself if the value is `23`:”


```java
set.addObserver(new SetObserver<>() {
    public void added(ObservableSet<Integer> s, Integer e) {
        System.out.println(e);
        if (e == 23)
            s.removeObserver(this);
    }
});
```

* “You might expect the program to print the numbers `0` through `23`, after which the observer would unsubscribe and the program would terminate silently. In fact, it prints these numbers and then throws a `ConcurrentModificationException`. The problem is that `notifyElementAdded` is in the process of iterating over the `observers` list when it invokes the observer’s `added` method. The `added` method calls the observable set’s `removeObserver` method, which in turn calls the method `observers.remove`.”
* “Now let’s try something odd: let’s write an observer that tries to unsubscribe, but instead of calling `removeObserver` directly, it engages the services of another thread to do the deed. This observer uses an *executor service* (Item 80):”


```java
// Observer that uses a background thread needlessly
set.addObserver(new SetObserver<>() {
   public void added(ObservableSet<Integer> s, Integer e) {
      System.out.println(e);
      if (e == 23) {
         ExecutorService exec =
               Executors.newSingleThreadExecutor();
         try {
            exec.submit(() -> s.removeObserver(this)).get();
         } catch (ExecutionException | InterruptedException ex) {
            throw new AssertionError(ex);
         } finally {
            exec.shutdown();
         }
      }
   }
});
```

* “When we run this program, we don’t get an exception; we get a deadlock. The background thread calls `s.removeObserver`, which attempts to lock `observers`, but it can’t acquire the lock, because the main thread already has the lock. All the while, the main thread is waiting for the background thread to finish removing the observer, which explains the deadlock.”
* “Suppose you were to invoke an alien method from a synchronized region while the invariant protected by the synchronized region was temporarily invalid. Because locks in the Java programming language are *reentrant*, such calls won’t deadlock.”
* “Reentrant locks simplify the construction of multithreaded object-oriented programs, but they can turn liveness failures into safety failures.”
* “Luckily, it is usually not too hard to fix this sort of problem by moving alien method invocations out of synchronized blocks. For the `notifyElementAdded` method, this involves taking a “snapshot” of the `observers` list that can then be safely traversed without a lock.”

```java
// Alien method moved outside of synchronized block - open calls
private void notifyElementAdded(E element) {
    List<SetObserver<E>> snapshot = null;
    synchronized(observers) {
        snapshot = new ArrayList<>(observers);
    }
    for (SetObserver<E> observer : snapshot)
        observer.added(this, element);
}
```

* “In fact, there’s a better way to move the alien method invocations out of the synchronized block. The libraries provide a *concurrent collection* (Item 81) known as `CopyOnWriteArrayList` that is tailor-made for this purpose.”
  * “This `List` implementation is a variant of `ArrayList` in which all modification operations are implemented by making a fresh copy of the entire underlying array. Because the internal array is never modified, iteration requires no locking and is very fast.”


```java
// Thread-safe observable set with CopyOnWriteArrayList
private final List<SetObserver<E>> observers =
        new CopyOnWriteArrayList<>();

public void addObserver(SetObserver<E> observer) {
    observers.add(observer);
}

public boolean removeObserver(SetObserver<E> observer) {
    return observers.remove(observer);
}

private void notifyElementAdded(E element) {
    for (SetObserver<E> observer : observers)
        observer.added(this, element);
} 
```

* “An alien method invoked outside of a synchronized region is known as an *open call* [Goetz06, 10.1.4]. Besides preventing failures, open calls can greatly increase concurrency. An alien method might run for an arbitrarily long period. If the alien method were invoked from a synchronized region, other threads would be denied access to the protected resource unnecessarily.”
* “As a rule, you should do as little work as possible inside synchronized regions.”
  * “Obtain the lock, examine the shared data, transform it as necessary, and drop the lock.”
  * “If you must perform some time-consuming activity, find a way to move it out of the synchronized region without violating the guidelines in Item 78.”
* “In a multicore world, the real cost of excessive synchronization is not the CPU time spent getting locks; it is *contention*: the lost opportunities for parallelism and the delays imposed by the need to ensure that every core has a consistent view of memory. Another hidden cost of oversynchronization is that it can limit the VM’s ability to optimize code execution.”
* “If you are writing a mutable class, you have two options: you can omit all synchronization and allow the client to synchronize externally if concurrent use is desired, or you can synchronize internally, making the class *thread-safe* (Item 82).”
  * “The collections in `java.util` (with the exception of the obsolete Vector and Hashtable) take the former approach, while those in `java.util.concurrent` take the latter (Item 81).”
  * “For example, `StringBuffer` instances are almost always used by a single thread, yet they perform internal synchronization. It is for this reason that `StringBuffer` was supplanted by `StringBuilder`, which is just an unsynchronized `StringBuffer`.”
  * “Similarly, it’s a large part of the reason that the thread-safe pseudorandom number generator in `java.util.Random` was supplanted by the unsynchronized implementation in `java.util.concurrent.ThreadLocalRandom`.”
  * **“When in doubt, do *not* synchronize your class, but document that it is not thread-safe.”**
* “If a method modifies a static field and there is any possibility that the method will be called from multiple threads, you must synchronize access to the field internally (unless the class can tolerate nondeterministic behavior).”
  * “It is not possible for a multithreaded client to perform external synchronization on such a method, because unrelated clients can invoke the method without synchronization.”
* “In summary, to avoid deadlock and data corruption, never call an alien method from within a synchronized region. More generally, keep the amount of work that you do from within synchronized regions to a minimum. When you are designing a mutable class, think about whether it should do its own synchronization. In the multicore era, it is more important than ever not to oversynchronize. Synchronize your class internally only if there is a good reason to do so, and document your decision clearly (Item 82).”


## Item 80: Prefer executors, tasks, and streams to threads

```java
ExecutorService exec = Executors.newSingleThreadExecutor();
exec.execute(runnable);
exec.shutdown();
```

* “You can do *many* more things with an executor service.”
  * “you can wait for a particular task to complete (with the `get` method)”
  * “you can wait for any or all of a collection of tasks to complete (using the `invokeAny` or `invokeAll` methods)”
  * “you can wait for the executor service to terminate (using the `awaitTermination` method)”
  * “you can retrieve the results of tasks one by one as they complete (using an `ExecutorCompletionService`)”
  * “you can schedule tasks to run at a particular time or to run periodically (using a `ScheduledThreadPoolExecutor`)”
* “If you want more than one thread to process requests from the queue, simply call a different static factory that creates a different kind of executor service called a *thread pool*.”
  * “You can create a thread pool with a fixed or variable number of threads.”
  * “The `java.util.concurrent.Executors` class contains static factories that provide most of the executors you’ll ever need.”
  * “If, however, you want something out of the ordinary, you can use the `ThreadPoolExecutor` class directly. This class lets you configure nearly every aspect of a thread pool’s operation.”
* “Choosing the executor service for a particular application can be tricky.”
  * “For a small program, or a lightly loaded server, `Executors.newCachedThreadPool` is generally a good choice because it demands no configuration and generally “does the right thing.”
  * “In a heavily loaded production server, you are much better off using `Executors.newFixedThreadPool`, which gives you a pool with a fixed number of threads, or using the `ThreadPoolExecutor` class directly, for maximum control.”
* “When you work directly with threads, a `Thread` serves as both a unit of work and the mechanism for executing it.”
* “In the executor framework, the unit of work and the execution mechanism are separate. The key abstraction is the unit of work, which is the *task*.”
  * “There are two kinds of tasks: `Runnable` and its close cousin, `Callable` (which is like `Runnable`, except that it returns a value and can throw arbitrary exceptions).”
  * “In essence, the Executor Framework does for execution what the Collections Framework did for aggregation.”
* “In Java 7, the Executor Framework was extended to support fork-join tasks, which are run by a special kind of executor service known as a fork-join pool.”
  * “A fork-join task, represented by a `ForkJoinTask` instance, may be split up into smaller subtasks, and the threads comprising a `ForkJoinPool` not only process these tasks but “steal” tasks from one another to ensure that all threads remain busy, resulting in higher CPU utilization, higher throughput, and lower latency.”
  * “Parallel streams (Item 48) are written atop fork join pools and allow you to take advantage of their performance benefits with little effort, assuming they are appropriate for the task at hand.”


## Item 81: Prefer concurrency utilities to `wait` and `notify`

* **“Given the difficulty of using `wait` and `notify` correctly, you should use the higher-level concurrency utilities instead.”**
* “The higher-level utilities in `java.util.concurrent` fall into three categories: the Executor Framework, which was covered briefly in Item 80; concurrent collections; and synchronizers.”
* “The concurrent collections are high-performance concurrent implementations of standard collection interfaces such as `List`, `Queue`, and `Map`. To provide high concurrency, these implementations manage their own synchronization internally (Item 79). ”
  * “Therefore, **it is impossible to exclude concurrent activity from a concurrent collection; locking it will only slow the program**.”

```java
// Concurrent canonicalizing map atop ConcurrentMap - not optimal
private static final ConcurrentMap<String, String> map =
        new ConcurrentHashMap<>();

public static String intern(String s) {
    String previousValue = map.putIfAbsent(s, s);
    return previousValue == null ? s : previousValue;
}
```

```java
// Concurrent canonicalizing map atop ConcurrentMap - faster!
public static String intern(String s) {
    String result = map.get(s);
    if (result == null) {
        result = map.putIfAbsent(s, s);
        if (result == null)
            result = s;
    }
    return result;
}
```

* **“Use `ConcurrentHashMap` in preference to `Collections.synchronizedMap`.”**
  * “Simply replacing synchronized maps with concurrent maps can dramatically increase the performance of concurrent applications.”
* “Some of the collection interfaces were extended with *blocking operations*, which wait (or *block*) until they can be successfully performed.”
  * “For example, `BlockingQueue` extends `Queue` and adds several methods, including `take`, which removes and returns the head element from the queue, waiting if the queue is empty.”
  * “This allows blocking queues to be used for *work queues* (also known as *producer-consumer queues*), to which one or more *producer threads* enqueue work items and from which one or more *consumer threads* dequeue and process items as they become available.”
* “*Synchronizers* are objects that enable threads to wait for one another, allowing them to coordinate their activities.”
  * “The most commonly used synchronizers are `CountDownLatch` and `Semaphore`. Less commonly used are `CyclicBarrier` and `Exchanger`. The most powerful synchronizer is `Phaser`.”

* “Countdown latches are single-use barriers that allow one or more threads to wait for one or more other threads to do something.”
  * “The sole constructor for `CountDownLatch` takes an int that is the number of times the `countDown` method must be invoked on the latch before all waiting threads are allowed to proceed.”
  * “For example, suppose you want to build a simple framework for timing the concurrent execution of an action. This framework consists of a single method that takes an executor to execute the action, a concurrency level representing the number of actions to be executed concurrently, and a runnable representing the action. All of the worker threads ready themselves to run the action before the timer thread starts the clock. When the last worker thread is ready to run the action, the timer thread “fires the starting gun,” allowing the worker threads to perform the action. As soon as the last worker thread finishes performing the action, the timer thread stops the clock. Implementing this logic directly on top of `wait` and `notify` would be messy to say the least, but it is surprisingly straightforward on top of `CountDownLatch`:”

```java
// Simple framework for timing concurrent execution
public static long time(Executor executor, int concurrency,
            Runnable action) throws InterruptedException {
    CountDownLatch ready = new CountDownLatch(concurrency);
    CountDownLatch start = new CountDownLatch(1);
    CountDownLatch done  = new CountDownLatch(concurrency);

    for (int i = 0; i < concurrency; i++) {
        executor.execute(() -> {
            ready.countDown(); // Tell timer we're ready
            try {
                start.await(); // Wait till peers are ready
                action.run();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            } finally {
                done.countDown();  // Tell timer we're done
            }
        });
    }

    ready.await();     // Wait for all workers to be ready
    long startNanos = System.nanoTime();
    start.countDown(); // And they're off!
    done.await();      // Wait for all workers to finish
    return System.nanoTime() - startNanos;
}
```

* “The executor passed to the `time` method must allow for the creation of at least as many threads as the given concurrency level, or the test will never complete. This is known as a *thread starvation deadlock* [Goetz06, 8.1.1].”
* “If a worker thread catches an `InterruptedException`, it reasserts the interrupt using the idiom `Thread.currentThread().interrupt()` and returns from its run method. This allows the executor to deal with the interrupt as it sees fit. ”
* “Note that `System.nanoTime` is used to time the activity.”
  * “For interval timing, always use `System.nanoTime` rather than `System.currentTimeMillis`.”
* “The `wait` method is used to make a thread wait for some condition. It must be invoked inside a synchronized region that locks the object on which it is invoked. Here is the standard idiom for using the `wait` method:”

```java
// The standard idiom for using the wait method
synchronized (obj) {
    while (<condition does not hold>)
        obj.wait(); // (Releases lock, and reacquires on wakeup)
    ... // Perform action appropriate to condition
}
```

* **“Always use the wait loop idiom to invoke the `wait` method; never invoke it outside of a loop.”**
* “A related issue is whether to use `notify` or `notifyAll` to wake waiting threads. (Recall that `notify` wakes a single waiting thread, assuming such a thread exists, and `notifyAll` wakes all waiting threads.)”
  * “It is sometimes said that you should ***always* use `notifyAll`**. This is reasonable, conservative advice. ”

  * “It will always yield correct results because it guarantees that you’ll wake the threads that need to be awakened. You may wake some other threads, too, but this won’t affect the correctness of your program.”
* “Even if these preconditions are satisfied, there may be cause to use `notifyAll` in place of `notify`. Just as placing the `wait` invocation in a loop protects against accidental or malicious notifications on a publicly accessible object, using `notifyAll` in place of `notify` protects against accidental or malicious waits by an unrelated thread.”
* **“There is seldom, if ever, a reason to use `wait` and `notify` in new code.”**

## Item 82: Document thread safety

* **“The presence of the `synchronized` modifier in a method declaration is an implementation detail, not a part of its API.”**
  * “It does not reliably indicate that a method is thread-safe.”
* **“To enable safe concurrent use, a class must clearly document what level of thread safety it supports.”**
  * “Immutable—Instances of this class appear constant. No external synchronization is necessary.”
    * “Examples include `String`, `Long`, and `BigInteger` (Item 17).”
  * “Unconditionally thread-safe—Instances of this class are mutable, but the class has sufficient internal synchronization that its instances can be used concurrently without the need for any external synchronization.”
    * “Examples include `AtomicLong` and `ConcurrentHashMap`.”

  * “Conditionally thread-safe—Like unconditionally thread-safe, except that some methods require external synchronization for safe concurrent use.”
    * “Examples include the collections returned by the `Collections.synchronized` wrappers, whose iterators require external synchronization.”
  * “Not thread-safe—Instances of this class are mutable. To use them concurrently, clients must surround each method invocation (or invocation sequence) with external synchronization of the clients’ choosing.”
    * “Examples include the general-purpose collection implementations, such as `ArrayList` and `HashMap`.”
  * “Thread-hostile—This class is unsafe for concurrent use even if every method invocation is surrounded by external synchronization.”
    * “Thread hostility usually results from modifying static data without synchronization.”
    * “When a class or method is found to be thread-hostile, it is typically fixed or deprecated.”
    * “The `generateSerialNumber` method in Item 78 would be thread-hostile in the absence of internal synchronization, as discussed on page 322.”
  * “These categories (apart from thread-hostile) correspond roughly to the *thread safety annotations* in *Java Concurrency in Practice*, which are `Immutable`, `ThreadSafe`, and `NotThreadSafe` [Goetz06, Appendix A].”
* “Documenting a conditionally thread-safe class requires care. You must indicate which invocation sequences require external synchronization, and which lock (or in rare cases, locks) must be acquired to execute these sequences.”
* “The description of a class’s thread safety generally belongs in the class’s doc comment, but methods with special thread safety properties should describe these properties in their own documentation comments.”
* “It is not necessary to document the immutability of enum types.”
* “Unless it is obvious from the return type, static factories must document the thread safety of the returned object, as demonstrated by `Collections.synchronizedMap` (above).”
* “To prevent this denial-of-service attack, you can use a *private lock object* instead of using synchronized methods (which imply a publicly accessible lock):”


```java
// Private lock object idiom - thwarts denial-of-service attack
private final Object lock = new Object();

public void foo() {
    synchronized(lock) {
        ...
    }
}
```

* “We are applying the advice of Item 17, by minimizing the mutability of the lock field.”
  * **“Lock fields should always be declared final.”**
* “The private lock object idiom can be used only on *unconditionally* thread-safe classes. Conditionally thread-safe classes can’t use this idiom because they must document which lock their clients are to acquire when performing certain method invocation sequences.”
* **“To summarize, every class should clearly document its thread safety properties with a carefully worded prose description or a thread safety annotation. The `synchronized` modifier plays no part in this documentation. Conditionally thread-safe classes must document which method invocation sequences require external synchronization and which lock to acquire when executing these sequences. If you write an unconditionally thread-safe class, consider using a private lock object in place of synchronized methods. This protects you against synchronization interference by clients and subclasses and gives you more flexibility to adopt a sophisticated approach to concurrency control in a later release.”**

## Item 83: Use lazy initialization judiciously

* “*Lazy initialization* is the act of delaying the initialization of a field until its value is needed. If the value is never needed, the field is never initialized. This technique is applicable to both static and instance fields.”
* “If a field is accessed only on a fraction of the instances of a class *and* it is costly to initialize the field, then lazy initialization may be worthwhile.”
  * **“The only way to know for sure is to measure the performance of the class with and without lazy initialization.”**
* “In the presence of multiple threads, lazy initialization is tricky. ”
  * “If two or more threads share a lazily initialized field, it is critical that some form of synchronization be employed, or severe bugs can result (Item 78).”
* **“Under most circumstances, normal initialization is preferable to lazy initialization.”**
* “**If you use lazy initialization to break an initialization circularity, use a synchronized accessor** because it is the simplest, clearest alternative:”


```java
// Lazy initialization of instance field - synchronized accessor
private FieldType field;

private synchronized FieldType getField() {
    if (field == null)
        field = computeFieldValue();
    return field;
}
```

* **“If you need to use lazy initialization for performance on a static field, use the *lazy initialization holder class idiom*.”**

```java
// Lazy initialization holder class idiom for static fields
private static class FieldHolder {
    static final FieldType field = computeFieldValue();
}

private static FieldType getField() { return FieldHolder.field; }
```

* “When `getField` is invoked for the first time, it reads `FieldHolder.field` for the first time, causing the initialization of the `FieldHolder` class.”
* **“If you need to use lazy initialization for performance on an instance field, use the *double-check idiom*.”**
  * “This idiom avoids the cost of locking when accessing the field after initialization (Item 79).”

```java
// Double-check idiom for lazy initialization of instance fields
private volatile FieldType field;

private FieldType getField() {
    FieldType result = field;
    if (result == null) {  // First check (no locking)
        synchronized(this) {
            if (field == null)  // Second check (with locking)
                field = result = computeFieldValue();
        }
    }
    return result;
}
```

* “Occasionally, you may need to lazily initialize an instance field that can tolerate repeated initialization. If you find yourself in this situation, you can use a variant of the double-check idiom that dispenses with the second check. It is, not surprisingly, known as the *single-check idiom*.”

```java
// Single-check idiom - can cause repeated initialization!
private volatile FieldType field;

private FieldType getField() {
    FieldType result = field;
    if (result == null)
        field = result = computeFieldValue();
    return result;
}
```

* “All of the initialization techniques discussed in this item apply to primitive fields as well as object reference fields.”
* **“In summary, you should initialize most fields normally, not lazily. If you must initialize a field lazily in order to achieve your performance goals or to break a harmful initialization circularity, then use the appropriate lazy initialization technique. For instance fields, it is the double-check idiom; for static fields, the lazy initialization holder class idiom. For instance fields that can tolerate repeated initialization, you may also consider the single-check idiom.”**