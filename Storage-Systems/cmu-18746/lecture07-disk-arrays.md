# Lecture 7 Multi-disk Systems and Disk Array Data Organizations

## Multi-disk Systems

* **Storage capacity**
  * Problem: cost, data growth
  * Solution: use multiple disks
  * Goal: Data growth at least 2x-3x faster than areal density
    * Approach: combine capacity of multiple disks
* **Performance**
  * Problem: load balancing
  * Solution: dynamic placement, striping
  * Goal: Bandwidth (MB/second)
    * Approach: stream data from multiple disks in parallel
  * Goal: Throughput (IOs/second)
    * Approach: concurrent requests to multiple disks
* **Reliability**
  * Problem: guarantee fault tolerance
  * Solution: replication, parity
  * Goal: Tolerate partial and full disk failures
    * Approach: store data copies across different disks
* Popular solution: **R**edundant **A**rray of **I**ndependent **D**isks

### Disk Subsystem Load Balancing

* Assumption: uniform data placement policy
  * Writing performance: excellent
  * Reading performance: **depends**
* Problem: some data more popular than other data
  * Distribution depends on the apps, usage, time
* Load Imbalances: fixed vs. migrating
  * **Fixed**: some data is always hot
  * **Migrating**: hot data chanegs over time
* Goal: find the right data placement policy
  * Common approach: Fixed data placement
  * Better approach: Dynamic data placement
    * Popular files labelled as hot, separated across multiple disks
  * Practical appraoch: **Disk striping**
    * Distribute chunks of data across all disks

### Disk Striping

* Data interleaved across multiple disks
  * Large file streaming benefits from parallel transfers
  * Thorough load balancing ideal for high-throughput requests
* How disk striping works
  * Break up LBN space into fixed-size stripe units
  * Distribute stripe units among disk in round-robin fashion
  * Straight-forward to compute location of block #B
    * Disk # = B % N, where N = number of disks
    * Disk block # = B / N (computes block offset on given disk)
* Key design decision: picking the stripe unit size
  * Assist alignment: choose multiple of file system block size