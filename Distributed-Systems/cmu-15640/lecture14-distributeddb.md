# Lecture 14 Distributed Databases: Case Study

## Consistent Distributed Database

![consistent_distributed_database](images/lecture14-distributeddb/consistent_distributed_database.png)

![fault_tolerant_distributed_database_ii](images/lecture14-distributeddb/fault_tolerant_distributed_database_ii.png)

|                   | Use Case                                         | Problems                     |
| ----------------- | ------------------------------------------------ | ---------------------------- |
| Distributed Mutex | Distributed KV without transactions              | Failures + Slow              |
| 2PC               | Disrtibuted DB with transactions (e.g., Spanner) | Failures                     |
| Primary-Backup    | Cost-efficient fault tolerance (e.g., FaRM)      | Correlated failures          |
| Paxos             | Staying up no matter the cost (e.g., Spanner)    | Delay and huge cost overhead |
| RAID, Checksums   | Every system                                     | Node failures                |

### Practical Constraints

* When availability over consistency
  * Challenge: Version reconcilation (parallel writes)
  * Practical approach (Dynamo): Vector Clocks
* Trend: Stronger-than-sequential consistency
  * Resurgence of consistent distributed DBs
  * Workloads are read heavy
* Cannot guarantee 100% availability

### Reading from Single Machine

* **Read lock**
  * Block all writes until read has finished
* **Snapshot**
  * Read from DB-copy, writes continue to original DB
* **Multi-version Concurrency Control (MVCC)**
  * New commit -> add as (timestamp, value)
  * Keep old (timestamp, value) tuples
  * Snapshot: read latest tuples with timestamp < now
* When reading from mutiple machines, must create distributed snapshots at exactly the same time
* e.g., PostgreSQL
  * Need synchronized clocks across all nodes
  * Need highly accurate time synchronization
  * Time sync error proportional to RTT
  * Global Internet RTTs in 100s of milliseconds

### Spanner: Google's Globally-Distributed Database

* Feature: **Lock-free** distributed read transactions
* Property: **External consistency** of distributed transactions
* Implementation: WAL + 2PC + Paxos + Snapshots
* **TrueTime: Interval-based global time**

![interval_based_global_time](images/lecture14-distributeddb/interval_based_global_time.png)

* Strict two-phase locking for write transactions
* Assign timestamp while locks are held

![how_spanner_do_time_sync](images/lecture14-distributeddb/how_spanner_do_time_sync.png)

* Challenge: time sync errors even with GPS/atomic clocks
* Conceptually must wait until all write transactions visible (their timestamps have passed)

![spanner_truetime_concept](images/lecture14-distributeddb/spanner_truetime_concept.png)

* Global wall-clock time with bounded uncertainty

![spanner_external_consistency](images/lecture14-distributeddb/spanner_external_consistency.png)

* If a transaction T1 commits before another transaction T2 starts, then T1's commit timestamp is smaller than T2
* Similar to how we reason with wall-clock time
