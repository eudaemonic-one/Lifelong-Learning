# Lecture 18 Fault Tolerance

## Failures in the Field

* Failures are expensive
* Disk failures in practice
  * HDDs seem to fail often
    * 1-13% replaced annually
  * SDDs seem more reliable
    * 1-3% replaced annually
  * Studies look at replacement events only
  * Definitions of failed hard drive are important

## Failure Types

* Media errors
  * Hard errors (Permanent failures)
    * Extrinsic failures (Manufacturing defects, etc.)
    * Intrinsic failures (Wear-out)
  * Soft errors (Transient errors)
* Faulty component behaviors
  * Byzantine behavior: may send arbitrary messages
  * Fail-stop behavior: stops and does not send messages
* Examples of fail-stop failures
  * Media failure: I/O device code bugs, disk HW failures -> Loss of durable data
  * System failure: DB bug, OS fault, HW failure -> Loss of volatile data but durable memory (disk) survives
* Transaction failure
  * Code aborts, based on input/database inconsistency
  * Mechanical aborts caused by concurrency control solutions to isolation
  * Frequent events, "instant" recovery needed

## Handling Failures

### Handling Hard Failures

* **Storage redundancy + Repeatable computation**
* Short, non-shared, deterministic programs
  * OS, framework, or user destroys partial changes then reruns program
  * Builds on external storage independently protected (RAID/Replicas)
* Long running, non-shared, deterministic programs
  * Examples: Extract/Transform/Load jobs
  * Periodic state checkpoints to durable, independent, protected storage
  * On failure: isolate failed component/system, restart from checkpoint
    * Dependent components can trigger failure detection

### Handling Soft Failures

* **Micro-reboot** restarts individual long-running system software components
  * Fine-grained, transactional components: restart and reinitialize fast
  * State segregation: prevent corruption by storing important state externally
  * Loosely coupled component: well-defined, well-enforced boundaries
  * Retry-able requests: inter-component interactions use timeouts
  * Expiring locks (leases): clean-up simplification
* Concurrent, shared data (database) multi-app systems
  * Shared state interaction through ACID transactions and write-ahead logging
  * External state independently protected
* Concurrent, shared-nothing replicated systems (may be no external state)
  * Replicated state machines, driven by coordinating replica changes

#### Transactions

* Goal: multiple users manipulate shared data safely
* ACID properties of a transaction
  * Atomicity: an operation is done all-or-nothing
  * Consistency: user-specified constraints applied before commit
  * Isolation: partial changes not visible to other users' code (less complex)
  * Durability: changes survive subsequent failures (storage and process redundancy)
* "AID" provided by database system, "C" (mostly) by programmer
  * Database is consistent if and only if contents result only from successful transactions
  * Integrity constraints (partial consistency) may be enforced by DB

#### Isolation: Two-Phase Locking (2PL)

* Goal: isolation mechanism guaranteeing serializability
* Assuming well-formed/consistent transactions seeking isolation
  * Simple locking fails to provide isolation if transactions interleave mutation/locking
* 2PL: acquire no lock after releasing any
  * Strict 2PL: release no lock before committing, avoids cascading aborts
  * Locks held a long time increase blocking; decrease concurrency
* Optimistic methods don't lock but may abort and retry
  * Faster if conflict is rare, but risks livelock if not

#### Recoverable Database System Model

* Log changes durably before database changes durable
  * Write-ahead logging
* REDO: repeat completed transaction on old DB data
  * Partial system or total media failure
* UNDO: rollback aborted transaction
  * Transaction or system failure
  * Only if uncommited transaction allowed to change durable media

#### Replicated State Machines

* State machine: code + data + input command = deterministic output
* Assuming non-faulty replicas: same initial state + same input = same final state + same output
  * Fail-stop failures: 1 surviving replica is sufficient
  * Byzantine failures: non-faulty survivors must win a vote
    * Need 2t+1 replicas to survive t malicious failures
* Common tools: PAXOS, Apache ZooKeeper
* Agreement: deilver every request to all non-faulty machines
  * A coordinator/client specifies a request and the rest agree
* Ordering: ensure the same order of execution at all non-faulty machines
  * Assign identifier to requests and execute in identifier order
  * Use a clock: Logical clock, Real-time clock, Replica-generated clock

#### Concurrency and "Happens Before"

* Two events are not concurrent if one "happens before" the other
* Replicated state machine wants same order of changes at all replicas