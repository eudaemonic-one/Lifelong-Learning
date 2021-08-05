# Storage Evolution @ Google

## Major Challenges

* A few petabyte of free space
  * An emergency low disk space condition
* Can I deliver bytes affordably?
* Can I deliver I/O affordably?
* Will they come in the right ratio?
* Giant organization: hard to understand where demand comes from
* Lots of developers: how to make storage easy while making the costs comprehensible
* Huge demand: can't pretend drive industry supply is infinite

### I/O

* Indexing, logs: big users of batch disk I/O
* Overall dominated by BigTable, Spanner

## Storage Devices

* DRAM
  * Pros: Fast (profiling like CPU), Small accesses
  * Cons: Not persistent, Expensive, Lack of storage API
* 3D XPoint
  * Pros: Fast (profiling like CPU), Small accesses, Persistent, Denser than DRAM
  * Cons: Intel-only, Expensive, Lack of storage API
* Flash
  * Pros: Fast (read latency ~100usec, 100k of reads/sec), Lower power when there's no I/O
  * Cons: Write wearout, Write amplificatio from GC, Surprising performance anomalies
* Disk
  * 7200RPM "Nearline" drive (80\*64KB reads/sec, 50\*1MB reads/sec, 150MB/sec)
  * 16TB drives are common
  * Reading 64KB data averages 13ms
  * But 99.99% can be seconds
    * Caching, sector remapping, read errors
    * Previous write delaying read
    * Non-data commands
    * Background task in the firmware
* Tape
  * Pros: Media is cheap, Keeps getting better, Very difficult for software to quickly wipe all data
  * Cons: Operationally difficult, Three parts: tape, drive, library, Time to first byte is really long

## Total Cost of Ownership

* Major components of disk TCO:
  * Disk acquisition cost
  * Disk power cost
  * Disk connection overhead (disk slot, compute, and network)
  * Repairs and maintenance overhead
* Most importantly, we care about storage TCO, nots disk TCO
* We minimize total storage TCO if we keep the disk **full** and **busy**

### What Disk Should I Buy?

* Have a mix because we're growing
* Have an overall goal for IOPS (I/O per second) and capacity
* Select disks to bring the cluster and fleet closer to our overall mix
* Buy flash for caching to bring IOPS/GB into disk range
* Buy disks for capacity and fill them up

### Filling Up Disks

* Filesystem doesn't work when 100% full
* Can't remove capacity for repairs or upgrades without empty space
* Individual groups don't want to run near 100% of quota
* Administrators are uncomfortable with statistical overcommit
* Worries about supply chain uncertainty

## GFS

* Designed and implemented in 2002
* Location independent namespace
* Initially 100s of TB, scaled to low 10s of PB
* Completely userspace implementation, no POSIX semantics
* Simple design, good for large batch work

### GFS Feature Accretion

* Shadow masters: read-only lagging replicas
* Multiple GFS cells per chunkserver
* Automatic master election, consistent metadata replication
* Archival Reed-Solomon encodings
  * Must be first written replicated
  * Might have long pauses when reading

### GFSv2

* More predictable performance
* More predictable **tail latency**
* GFS Master replcaed by **Colossus**
* GFS chunkserver replcaed by **D**

### Solve An Easier Problem

* A file system for BigTable
  * Append-only
  * Single writer (occasional multi-reader)
  * Rename only user to indicate finished file
  * No snapshots
  * Directories unnecessary

### Storage Options Back Then

* Sharded MySQL with local disk & replication
  * Poor load balancing
  * Complicated operations
* Local key-value stores with Paxos replication
  * Authentication database
  * Chubby
  * Doesn't scale
* BigTable (sorted key-value store on GFS)

### BigTable

* Automatically shards data across tablets
* Locates tablets via metadata lookup
* Easy to use semantics
* Efficient point lookups and scans
* In-memory locality groups

