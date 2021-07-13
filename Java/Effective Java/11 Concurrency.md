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
