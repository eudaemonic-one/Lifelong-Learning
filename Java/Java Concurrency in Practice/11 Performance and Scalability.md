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

## 11.2 Amdahl's Law

* *Amdahl's law*: how much a program can theoretically be sped up by additional computing resources, based on the proportion of parallelizable and serial components.
* We can achieve a speedup of at most:
  * $$Speedup \le \frac{1}{F+\frac{(1-F)}{N}}$$
  * F := the fraction of the calculation that must be executed serially
  * N := processor number
* Identify the sources of serialization:
  * synchronization to maintain the work queue's integrity in the face of concurrency access
  * accessing any shared data structure => serialization
  * result handling => final merge is a source of serialization

### 11.2.1 Example: Serialization Hidden in Framework

* The synchronized `LinkedList` guards the entire queue state with a single lock that is held for the duration of the `offer` or `remove` call; `ConcurrentLinkedQueue` uses a sophisticated nonblocking queue algorithm that uses atomic references to update individual link pointers.

![ch11fig02](images/11 Performance and Scalability/ch11fig02.gif)

### 11.2.2 Applying Amdahl's Law Qualitatively

* Amdahl's law => quantifies the possible speedup when more computing resources are available.
* Reducing lock granularity: lock splitting (splitting one lock into two) and lock striping (splitting one lock into many).
  * lock striping seems much more promising.
