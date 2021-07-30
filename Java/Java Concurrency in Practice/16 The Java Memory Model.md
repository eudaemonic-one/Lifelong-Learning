# Chapter 16. The Java Memory Model

* Java Memory Model (JMM) => low-level details.

## 16.1 What is a Memory Model, and Why would I Want One?

* In a multithreaded environment, sequentiality => significant performance cost.
* JMM => minimal guarantees for JVM.

### 16.1.1 Platform Memory Models

* Processor provides varying degrees of *cache coherence*.
* *memory model* => what guarantees they can expect from the memory system, and specifies the special instructions required (called *memory barriers* or *fences*) to get the additional memory cordination guarantees required when sharing data.
  * JVM handles all differences.
* *sequential consistency* => JMM does not offer it.

### 16.1.2 Reordering

* *reordering* => operations might be delayed or appear to execute out of order.

### 16.1.3 The Java Memory Model in 500 Words or Less

* *actions*: reads and writes to variables, locks and unlocks of monitors, and starting and joining with threads.
* *happens-before* => a partial ordering => to guarantee that the thread executing action B can see the results of action A, there must be a *happens-before* relationship between A and B.
* *data race* => when a variable is read by more than one thread, and written by at least one thread, but the reads and writes are not ordered by *happens-before*.
* *correctly synchronized program* => with no data races, and exhibit sequential consistency => all actions happen in a fixed, global order.
* The rules for *happen-before* are:
  * **Program order rule.** Each action in a thread *happens-before* every action in that thread that comes later in the program order.
  * **Monitor lock rule.** An unlock on a monitor lock *happens-before* every subsequent lock on that same monitor lock.
  * **Volatile variable rule.** A write to a volatile field *happens-before* every subsequent read of that same field.
  * **Thread start rule.** A call to `Thread.start` on a thread *happens-before* every action in the started thread.
  * **Thread termination rule.** Any action in a thread *happens-before* any other thread detects that thread has terminated, either by successfully return from `Thread.join` or by `Thread.isAlive` returning false.
  * **Interruption rule.** A thread calling interrupt on another thread *happens-before* the interrupted thread detects the interrupt (either by having `InterruptedException` thrown, or invoking `isInterrupted` or `interrupted`.
  * **Finalizer rule.** The end of a constructor for an object *happens-before* the start of the finalizer for that object.
* *transitivity* => If A *happens-before* B, and B *happens-before* C, then A *happens-before* C.
* Even though actions are only partially ordered, synchronization actions - lock acquisition and release, and reads and writes of volatile variables - are totally ordered.

### 16.1.4 Piggybacking on Synchronization

* *happens-before* => you can piggyback on the visibility properties of an existing synchronization.
* Other *happens-before* orderings guaranteed by the class library include:
  * Placing an item in a thread-safe collection *happens-before* another thread retrieves that item from the collection.
  * Counting down on a `CountDownLatch` *happens-before* a thread returns from `await` on that latch.
  * Releasing a premit to a `Semaphore` *happens-before* acquiring a permit from that same `Semaphore`.
  * Actions taken by the task represented by a `Future` *happens-before* another thread successfully returns from `Future.get`.
  * Submitting a `Runnable` or `Callable` to an `Executor` *happens-before* the task begins execution.
  * A thread arriving at a `CyclicBarrier` or `Exchanger` *happens-before* the other threads are released from the same barrier or exchange point.
