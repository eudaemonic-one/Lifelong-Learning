# Chapter 11. Performance and Scalability

* Threads
  * => improve resource utilization
    * => exploit available processing capacity
  * => improve responsiveness
    * => begin processing new tasks immediately while existing still running tasks
* Techniques for analyzing, monitoring, improving concurrent performance => complexity => increase the likelihood of safety and liveness failures.
* Safety always comes first.

## 11.1 Thinking about Performance

* Improving performance means doing more work with fewer resources.
  * e.g., CPU cycles, memory, network bandwidth, I/O bandwidth, database requests, disk space, other resources.
  * Performance limited by availability of a resource => CPU-bound, database-bound.
* Using multiple threads => performance costs like:
  * Overhead associated with coordinating between threads (locking, signaling, memory synchronization).
  * Increased context switching.
  * Thread creation and teardown.
  * Scheduling overhead.
* Using concurreny to achieve better performance.
  * => utilize the processing resources we have more effectively.
  * => enable our program to exploit additional processing resources if they become available.

### 11.1.1 Performance Versus Scalability

* Application performance => service time, throughput, efficiency, scalability.
* *Scalability*: the ability to improve throughput or capacity when additional computing resources are added.
  * Tuning for scalability => do *more* work with *more* resources.
  * => often *increasing* the amount of work done to process each *individual* task.
    * such as dividing tasks into multiple pipelined subtasks.
* Three-tier application model : presentation, business logic, persistence.
* For server applications, "how much" (scalability, throughput, capacity) are of greater concern than "how fast" aspects.
* For interactive applications, latency tends to be more important.

### 11.1.2 Evaluating Performance Tradeoffs

* Avoid premature optimization => most optimizations are often undertaken before a clear set of requirements is available.
* The quest for performance is probably the single greatest source of concurrency bugs.
* Measure, don't guess.
