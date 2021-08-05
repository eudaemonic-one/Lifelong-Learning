# Lecture 9 Tail Latency

## Tail Latency

### Service Times Vary

* Interference within shared infrastructure
* Lots of causes
  * Background/maintainenance activities
    * Garbage collection, log compaction, virus scanning, backup, etc
  * Dynamic and static hardware variations
  * Complex queuing and caching policies

### Slow Operations Matter

* Very slow responses make for angry users
  * e.g., below 100ms is plenty fast for humans
* Big jobs often wait for the last task to finish
  * better to have them all be a little slower, than to have few very slow

### Tail Latency & Fan-out

* Matters to applications with large fan-outs of leaf tasks
  * Fan-out parallelizes work to lower latency perhaps

### Reduce Service Time Variation

* Great option, but really difficult to do comprehensively
* Some approaches
  * Prioritize
    * e.g., do the stuff that is being waited for first (before background stuff)
    * e.g., do the stuff that is "falling behind" first
  * Manage background activities
    * e.g., synchronize schedulable maintenance stuff among machines
    * e.g., try to do background stuff only when not busy with stuff

## Tail Tolerance Techniques

* "Hedged" requests (or "speculative" redundant requests)
  * Ask more than one server to do the work (e.g., read replica, compute map)
    * usually only after the initial server doesn't finish quickly (to reduce waste)
  * Take the first response to arrive and ignore the slow one
* "Tied" requests (aggressive hedging with cancelation)
  * Ask more than one server immediately, but let them know you did
  * When one finishes (or starts), it "cancels" the other/redundant request
  * Addresses infrequent slow responses faster with less redundant work
* "Micro-partitioning"
  * Migrate (replicate) 5% of partitions when imbalanced
* "Probation"
  * Elimination of slow nodes from datapath until their specs get better

### Special App-specific Tail Tolerance Techniques

* e.g., large information-retrieval (IR) apps like fuzzy search
* Positive example: an IR service can answer without all leaves' responses
* Negative example: some queries can sometimes cause deterministic failure