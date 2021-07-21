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
