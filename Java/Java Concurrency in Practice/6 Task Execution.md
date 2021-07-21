# Chapter 6. Task Execution

* divide the work into tasks => simplify program organization, facilitate error recovery, promote concurrency.

## 6.1 Executing Tasks in Threads

* Identify sensible *task boundaries* => improve flexibility in scheduling, facilitate concurrency.
* Server applications
  * should exhibit both *good throughput* and *good responsiveness*
  * should exhibit *graceful degradation* as they become overloaded
  * use individual client requests as task boundaries => usually appropriate task sizing.
* **Executing Tasks Sequentially**
  * The main thread alternates between accepting requests and processing the associated request.
  * Processing a web request => a mix of computation and I/O => may block due to network congestion or connectivity problems.
* **Explicitly Creating Threads for Tasks**
  * => Task processing is offloaded from the main thread, enabling the main loop to resume waiting for the next incoming connection more quickly => improve responsiveness.
  * => Tasks can be processed in parallel, enabling multiple requests to be serviced simultaneously => improve throughput.
  * => Task-handling code must be thread-safe, because it may be invoked conurrently for multiple tasks.
* **Disadvantages of Unbounded Thread Creation**
  * **Thread lifecycle overhead.**
    * Thread creation and teardown overhead => latency.
    * Even worse if requests are frequent and lightweight.
  * **Resource consumption.**
    * Threads consume memory.
    * If runnable threads more than available processors, threads sit idle.
  * **Stability**
    * There is a limit on how many threads can be created.
    * Hit the limit => `OutOfMemoryErro` => risky to recover it.
  * You need to place some bound on how many threads the application creates, and test thoroughly to ensure that.

## 6.2 The Executor Framework

* `java.util.concurrent` provides a flexible thread pool implementation as part of the `Executor` framework.
  * The primary abstraction for task execution is `Executor`.
  * => decoupling *task submission* from *task execution*, describing tasks with `Runnable`.
  * => let its implementations provide lifecycle support and hooks for statistics gathering, application management, and monitoring.
  * `Executor` is based on the producer-consumer pattern.

![c0117-01](images/6 Task Execution/c0117-01.jpg)

* **Example: Web Server Using Executor**

![c0118-01](images/6 Task Execution/c0118-01.jpg)

![c0118-02](images/6 Task Execution/c0118-02.jpg)

* **Execution Policies**
  * an execution policy := what, where, when, how of a task execution.
  * policy depends on available computing resources and quality-of-service requirements.
  * limiting the number of concurrent tasks => avoid failure due to resource exhaustion or contention for scarce resources.
* **Thread Pools**
  * => manage a homogeneous pool of worker threads => tightly bounded by a *work queue* holding tasks waiting to be executed.
  * worker thread: request the next task from the work queue, execute it, and go back to waiting for another task.
  * executing tasks in thread pool => reusing existing threads amortizes thread creation and teardown costs, can have enough threads to keep the processors busy, while not running out of memory or thrashes due to competition.
  * Create a thread pool through static factory methods in `Executors`:
    * `newFixedThreadPool`: A fixed size thread pool creates threads as tasks are submitted, up to the maximum pool size, and then attempts to keep the pool size constant.
    * `newCachedThreadPool`: A cached thread pool has more flexibility to reap idle threads when the current size of the pool exceeds the demand for processing, and to add new threads when demand increases, but places no bounds on the size of the pool.
    * `newSingleThreadExecutor`: A single-threaded executor creates a single worker thread to process tasks, replacing it if it dies unexpectedly. Tasks are guaranteed to be processed sequentialy according to the order imposed by the task queue (FIFO, LIFO, priority order).
    * `newScheduledThreadPool`: A fixed-size thread pool that supports delayed and periodic task execution, similar to `Timer`.
* **Executor Lifecycle**
  * JVM can't exit until all the threads have terminated => failing to shut down an `Executor` could prevent the JVM from exiting.
  * shutdown := graceful shutdown + abrupt shutdown.
  * `ExecutorService` extends `Executor` to provide methods for lifecycle management (running, shutting down, terminated).
    * Tasks submitted to an `ExecutorService` after it has been shut down are handled by the *rejected execution handler* => might silently discard the task, or might cause `execute` to throw the unchecked `RejectedExecutionException`.
    * It is common to follow `shutdown` immediately by `awaitTermination`.

![c0121-01](images/6 Task Execution/c0121-01.jpg)

![c0122-01](images/6 Task Execution/c0122-01.jpg)

* **Delayed and Periodic Tasks**
  * `ScheduledThreadPoolExecutor` should be thought of as replacement for `Timer`.
  * `Timer` has some drawbacks.
    * => creates only a single thread for executing timer tasks => if a timer task takes too long => the timing accuracy can suffer.
    * => `TimerTask` throws unchecked exception, while `Timer` thread doesn't catch the exception => `Timer` would assume the entire `Timer` was cancelled => scheduled but not yet executed `TimerTask`s are never run.
  * `DelayQueue`: a `BlockingQueue` implementation that provides the scheduling functionality of `ScheduledThreadPoolExecutor`.
    * => manages a collection of `Delayed` objects.
    * => let you `take` an element only if its delay has expired.
    * => objects are returned by the time associated with their delay.

![c0124-01](images/6 Task Execution/c0124-01.jpg)
