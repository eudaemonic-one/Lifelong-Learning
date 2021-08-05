# Chapter 12. Testing Concurrent Programs

* Concurrent programs => nondeterminism => potential interactions, failure models.
* Tests of concurrent classes => safety + liveness.
  * Tests of safety => testing invariants such as assert invariants or execute test code atomically => can introduce timing or synchronization artifacts that can mask bugs that might otherwise manifest themselves.
  * Test of liveness => testing progress and nonprogress => hard to quantify.
* Performance measurement => Throughput + Responsiveness + Sacalability.

## 12.1 Testing for Correctness

* To build a set of test cases for a bounded buffer.
  * `BoundedBuffer` implements a fixed-length array-based queue with blocking `put` and `take` methods controlled by a pair of counting semaphores.
  * The `availableItems` semaphore represents the number of elements taht can be removed from the buffer, and is initially zero.
  * The `availableSpaces` represents how many items can be inserted into the buffer, and is initialized to the size of the buffer.
  * On exit from either the `put` or `take` methods, the sum of counts of both semaphores always equals the bound.

![c0249-01](images/12%20Testing%20Concurrent%20Programs/c0249-01.jpg)

### 12.1.1 Basic Unit Tests

* Sequential tests => disclose when a problem is *not* related to concurrency issues before you start looking for data races.

![c0250-01](images/12%20Testing%20Concurrent%20Programs/c0250-01.jpg)

### 12.1.2 Testing Blocking Operations

* If a method is supposed to block under certain conditions, then a test for that behavior should succeed only if the thread does *not* proceed.
  * => similar to testing throwing an exception.
* Once the method successfully blocks, you have to unblock it => via interruption => start a blocking activity in a separate thread, wait until the thread blocks, interrupt it, and then assert the blocking operation completed => requires the blocking methods to respond to interruption by returning early or throwing `InterruptedException`.
  * You have to make an arbitrary decision about how long the few instructions being executed could possibly take, and wait longer than that.
* An approach to testing blocking operations.
  * It creates a taker thread that attempts to `take` an element from an empty buffer.
  * If `take` succeeds, it registers failure; if correctly blocked, it will throw `InterruptedException`, and the `catch` block for this exception treats this as success and lets the thread exit.
  * The main test runner thread then attempts to `join` with the taker thread and verifies that the join returned successfully by calling `Thread.isAlive`; if the taker thread responded to the interrupt, the `join` should complete quickly.
  * The timed `join` ensures that the test completes even if `take` gets stuck in some unexpected way.
  * The same approach can be used to test that the taker thread unblocks after an element is placed in the queue by the main thread.

![c0252-01](images/12%20Testing%20Concurrent%20Programs/c0252-01.jpg)

* The result of `Thread.getState` should not be used for concurrency control, and is of limited usefulness for testing.
  * Its primary utility is as a source of debugging information.

### 12.1.3 Testing Safety

* To test a concurrent class performing correctly => multiple threads performing `put` and `take` operations over some amount of time and then test that nothing went wrong.
* Identify easily checked properties => checking the test property does not require any synchronization.

![c0255-01](images/12%20Testing%20Concurrent%20Programs/c0255-01.jpg)

![c0256-01](images/12%20Testing%20Concurrent%20Programs/c0256-01.jpg)

### 12.1.4 Testing Resource Management

* Test it does *not* do things it is *not* supposed to do.
* Undesirable memory retention can be easily tested with heap-inspection tools that measure application memory usage.

![c0258-01](images/12%20Testing%20Concurrent%20Programs/c0258-01.jpg)

### 12.1.5 Using Callbacks

* Callbacks to client-provided code can be helpful in constructing test cases; callbacks are often made at known points in an object's lifecycle that are good opportunities to assert invariants.
* Testing a thread pool => testing a number of elements of execution policy => additional threads are created and idle threads get reaped when they are supposed to.
  * Use a custom thread factory to instrument thread creation.
  * If the core pool size is smaller than the maximum size, the thread pool should grow as demand for execution increases.
  * Submitting long-running tasks to thez pool makes the number of executing tasks stay constant for long enough to make a few assertions.

![c0258-02](images/12%20Testing%20Concurrent%20Programs/c0258-02.jpg)

![c0259-01](images/12%20Testing%20Concurrent%20Programs/c0259-01.jpg)

### 12.1.6 Generating More Interleavings

* Increase the number of interleavings => use `Thread.yield` to encourage more switches during operations that access shared state.
  * => platform-specific because JVM is free to treat `Thread.yield` as no-op.
  * => may activate timing-sensitive bugs in code.

## 12.2 Testing for Performance

* It is always worthwhile to include some basic functionality testing within performance tests to ensure that you are not testing the performance of broken code.
* Performance test
  * => seek to measure end-to-end performance metrics for representative use cases.
  * => help select sizings empirically for various bounds.

### 12.2.1 Extending `PutTakeTest` to Add Timing

* Timing the entire run and dividing by the number of operations to get a per-operation time.
  * Use `CyclicBarrier` to start and stop the worker thread => use a barrier action measuring the start and end time.

![c0261-01](images/12%20Testing%20Concurrent%20Programs/c0261-01.jpg)

![c0262-01](images/12%20Testing%20Concurrent%20Programs/c0262-01.jpg)

* Test driver => running test for various combinations of parameters => throughput, how it scales with different numbers of threads, how we select the bound size.

![c0262-02](images/12%20Testing%20Concurrent%20Programs/c0262-02.jpg)

* Be careful about concluding from testing data.
  * The test is fairly artificial in how it simulates the application.

### 12.2.2 Comparing Multiple Algorithms

* `java.util.concurrent` => selected and tuned => more efficient than other implementations.

![ch12fig02](images/12%20Testing%20Concurrent%20Programs/ch12fig02.gif)

* The test suggests that `LinkedBlockingQueue` scales better than `ArrayBlockingQueue`.
  * Because it allows more concurrent access by `put`s and`take`s than an array-based queue because the best linked queue algorithms allow the head and tail to be updated independently.
  * Because allocation is usually thread-local, algorithms that can reduce contention by doing more allocation usually scale better.

### 12.2.3 Measuring Responsiveness

* We want to measure the *variance* of service time.
  * => allow us to estimate the answers to quality-of-service questions.
* Histograms of task completion times are normally the best way to visualize variance in service time.
  * => to keep track of per-task completion times in addition to aggregate completion time.
  * => measure the run time of small batches of operations => avoid error caused by timer granularity.
* Nonfair provides better throughput and fair provides lower variance.

## 12.3 Avoiding Performance Testing Pitfalls

### 12.3.1 Garbage Collection

* The timing of garbage collection is unpredictable.
  * a small variation in the size of the run could have a spurious effect on the measured time per iteration.
* Two strategies for preventing garbage collection from biasing your results:
  * => Ensure that garbage collection does not run at all during your test.
  * => Make sure that the garbage collector runs a number of times during your run so that the test program adequately reflects the cost of ongoing allocation and garbage collection.
  * The latter is better => it requires a longer test and is more likely to reflect real-world performance.

### 12.3.2 Dynamic Compilation

* Writing and interpreting performance benchmarks for dynamically compiled languages like Java is difficult.
* The timing of compilation is unpredictable.
  * Code may be decompiled and recompiled for various reasons.
* Allowing teh compiler to run during a measured test run can bias test results:
  * compilation comsumes CPU resources.
  * measuring the run time of a combination of interpreted and compiled code is not a meaningful performance metric.

![ch12fig05](images/12%20Testing%20Concurrent%20Programs/ch12fig05.gif)

* Prevent compilation from biasing your results
  * => run your program for a long time so that compilation and interpreted execution represent a small fraction of the total run time.
  * => use an unmeasured "warm-up" run.
    * The first group of results should be discarded as warm-up.
* When measuring multiple *unrelated* computationally intensive activities in a single run, it is a good idea to place explicit pauses between the measured trials to give the JVM a chance to catch up with background tasks with minimal interference from measured tasks.

### 12.3.3 Unrealistic Sampling of Code Paths

* Runtime compilers use profiling information to help optimize the code being compiled => different code paths.
* Run a mix of single-threaded and multi-threaded test => simulate realistic code paths.

### 12.3.4 Unrealistic Degrees of Contention

* Concurrent performance tests should try to approximate the thread-local computation done by a typical application in addition to the concurrent coordination under study.

### 12.3.5 Dead Code Elimination

* The optimizer prunes dead code from a program => you are measuring less execution than you think for a benchmark.
* Every computed result should be used somehow by your program => in a way that does not require synchronization or substantial computation.
* A trick if to compute the `hashCode` and compare it to an arbitrary value such as the current value of `System.nanoTime`.

![c0270-01](images/12%20Testing%20Concurrent%20Programs/c0270-01.jpg)

## 12.4 Complementary Testing Approaches

* The goal of testing is not so much to *find errors* as it is to *increase confidence* that the code works as expected.

### 12.4.1 Code Review

* Have concurrent code reviewed carefully by someone besides its author.
  * => finding subtle races.
  * => improving the quality of comments.

### 12.4.2 Static Analysis Tools

* *static analysis tools* => look for common *bug patterns*.
* FindBugs => detectors including concurrency-related bug patterns such as inconsistent synchronization.
* Invoking `Thread.run`.
  * It is always a mistake to call `Thread.run` directly; usually the programmer meant to call `Thread.start`.
* Unreleased lock.
  * The standard idiom is to release the lock from a `finally` block; otherwise the lock can remain unreleased in the event of an `Exception`.
* Empty `synchronized` block.
* Double-checked locking.
  * A broken idiom for reuding synchronization overhead in lazy initialization.
* Starting a thread from a constructor.
  * => introduce the risk of subclassing problems, and can allow the `this` reference to escape the constructor.
* Notification errors.
  * A `synchronized` block that calls `notify` or `notifyAll` but does not modify any state is likely to be an error.
* Condition wait errors.
  * When waiting on a condition queue, `Object.wait` or `Condition.await` should be called in a loop, with the appropriate lock held, after testing some state predicate.
* Misuse of `Lock` and `Condition`.
  * Using a `Lock` as the lock argument for a `synchronized` block is likely to be a typo, as is calling `Condition.wait` instead of `await`.
* Sleeping or waiting while holding a lock.
* Spin loops.
  * Code that does nothing but spin checking a field for an expected value can waste CPU time and, if the field is not volatile, is not guaranteed to terminate.
  * Latches or condition waits are often a better technique when waiting for a state transition to occur.

### 12.4.3 Aspect-oriented Testing Techniques

* Aspect-oriented programming (AOP) => assert invariants or some aspects of compliance with synchronization policies.

### 12.4.4 Profilers and Monitoring Tools

* Profiling tools support for threads.
  * => offer a display showing a timeline for each thread with different colors for the various thread states.
  * => show how effectively your program is utilizing the available CPU resources and where to look for the cause of doing badly.