# Lecture 19 Concurrency: Java Primitives, continued

## A Conccurrency Bug: Cooperative Thread Termination

```java
public class StopThread {
  private static boolean stopRequested;
	public static void main(String[] args) throws Exception {
    Thread backgroundThread = new Thread(() -> {
      while (!stopRequested)
        /* Do something */ ;
		});
    backgroundThread.start();
		TimeUnit.SECONDS.sleep(1);
		stopRequested = true;
  }
}
```

* What went wrong？

  * In the absence of synchronization, there is no guarantee as to when, if ever, one thread will see changes made by another

  * **JVMs can and do perform this optimization (“hoisting”):**

    * ````java
      while (!done)
        /* do something */ ;
      ````

  * becomes:

    * ```java
      if (!done)
        while (true)
          /* do something */ ;
      ```

* **You must lock write and read**

```java
public class StopThread {
	private static boolean stopRequested;
	private static synchronized void requestStop() {
		stopRequested = true;
  }
	private static synchronized boolean stopRequested() {
    return stopRequested;
	}
	public static void main(String[] args) throws Exception {
    Thread backgroundThread = new Thread(() -> {
      while (!stopRequested())
        /* Do something */ ;
		});
    backgroundThread.start();
		TimeUnit.SECONDS.sleep(1);
    requestStop();
  }
}
```

## A Liveness Problem: Poor Performance

```java
public class BankAccount {
  private long balance;
	public BankAccount(long balance) {
    this.balance = balance;
	}
  static synchronized void transferFrom(BankAccount source, BankAccount dest, long amount) {
		source.balance -= amount;
		dest.balance += amount;
  }
  public synchronized long balance() {
    return balance;
	}
}
```

* Lock on class

```java
public class BankAccount {
  private long balance;
	public BankAccount(long balance) {
    this.balance = balance;
	}
	static void transferFrom(BankAccount source, BankAccount dest, long amount) {
    synchronized(BankAccount.class) {
			source.balance -= amount;
			dest.balance += amount;
    }
  }
  public synchronized long balance() {
    return balance;
  }
}
```

* A proposed fix: lock splitting

```java
public class BankAccount {
  private long balance;
	public BankAccount(long balance) {
    this.balance = balance;
	}
	static void transferFrom(BankAccount source, BankAccount dest, long amount) {
    synchronized(source) {
			synchronized(dest) {
        source.balance -= amount;
			}
    }
	}
}
```

* **A liveness problem: deadlock**
  * A possible interleaving of operations
* **Avoiding deadlock**
  * The waits-for graph represents dependencies between threads
  * Deadlock has occurred iff the waits-for graph contains a cycle
  * One way to avoid deadlock: locking protocols that avoid cycles
* Avoiding deadlock by ordering lock acquisition

```java
public class BankAccount {
	private long balance;
	private final long id = SerialNumber.generateSerialNumber();
	public BankAccount(long balance) {
    this.balance = balance;
	}
  static void transferFrom(BankAccount source, BankAccount dest, long amount) {
    BankAccount first = (source.id < dest.id) ? source : dest;
    BankAccount second = (first == source) ? dest : source;
		synchronized (first) {
    	synchronized (second) {
				source.balance -= amount;
				dest.balance += amount;
      }
    }
  }
}
```

* Another subtle problem: The lock object is exposed
* An easy fix: Use a private lock
* Encapsulate an object’s state – Easier to implement invariants

```java
public class BankAccount {
	private long balance;
	private final long id = SerialNumber.generateSerialNumber();
  private final Object lock = new Object();
	public BankAccount(long balance) { this.balance = balance; }
  static void transferFrom(BankAccount source, BankAccount dest, long amount) {
		BankAccount first = source.id < dest.id ? source : dest;
    BankAccount second = first == source ? dest : source;
    synchronized (first.lock) {
			synchronized (second.lock) {
        source.balance -= amount;
        dest.balance += amount;
			}
    }
	}
}
```

## Challenges of Concurrency

* A liveness problem: poor performance
  * A proposed fix: lock splitting
* A liveness problem: deadlock
  * A possible interleaving of operations
  * The waits-for graph represents dependencies between threads
  * Deadlock has occurred iff the waits-for graph contains a cycle
  * Avoiding deadlock by ordering lock acquisition
* Another subtle problem: The lock object is exposed

## Concurrency and Information Hiding

* Encapsulate an object's state - Easier to implement invariants
  * Encapsulate synchronization - Easier to implement synchronization policy
* Aside: @ThreadSafe @NotThreadSafe @GuardedBy("lock")
* @ThreadSafe
  * Place this annotation on methods that can safely be called from more than one thread concurrently
  * The method implementer must ensure thread safety using a variety of possible techniques including immutable data, synchronized shared data, or not using any shared data at all
* @GuardedBy
  * Denotes that the annotated method or field can only be accessed when holding the referenced lock
* @Immutable
  * Immutable objects are constructed once, in a consistent state, and can be safely shared
  * Immutable objects are naturally thread-safe and can therefore be safely shared among threads

## JUnit does not well-support concurrent tests

* Write JUnit test with a false sense of security
* **JUnit doesn’t see assertion failures in other threads**
* Concurrent clients beware

## Concurrent Programming can be hard to get right

* Invoke `Thread.start`, not `Thread.run`
  * Can be very difficult to diagnose
* **This is a severe API design bug**
* **Thread should not have implemented Runnable**
  * This confuses is-a and has-a relationships
  * Thread's `runnable` should have been private
* Thread violates the "Minimize accessibility" principle
