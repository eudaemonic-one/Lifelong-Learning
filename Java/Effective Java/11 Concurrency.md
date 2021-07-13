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
