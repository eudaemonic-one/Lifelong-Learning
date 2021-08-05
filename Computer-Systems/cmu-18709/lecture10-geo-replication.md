# Lecture 10 Geo-replication

## Why Geo-replication

* Reliability: disaster survival
* Compliance: legal and political requirements
* Client responsiveness: latency (lower round-trip times)

## Geo-replicated Storage

* Storage tier Dimensions
  * Datacenter shards data across many nodes
  * Data geo-replicated in multiple datacenters
* Common geo-replicated storage goals
  * Serve client requests quickly
  * Scale out nodes/datacenters
  * Interact with data coherently
* **CAP Theorem** - You cannot always have Consistency, Availability, and Partition tolerance
  * Reality is that partition is rare, but during partition you have to pick between consistency (stop & wait) or availability (access stale data)

## ALPS Properties

* Many systems provide "**ALPS Properties**"
  * Availability
  * Low Latency (Local RTT)
  * Partition Tolerance
  * Scalability
* In ALPS-oriented systems, each replica "independent"
  * Any request can be serviced by any data center
  * Updates propagated to other data centers in the background
    * Essentially, updates are logged and streamed to other sites
    * Often via protocol that ensures "**eventual consistency**"
      * No guarantees on when
* ALPS-oriented geo-replicated storage cannot guarantee that users can interact with data coherently
* Consistency
  * Guarantees on the shared view across the system
  * Stronger consistency makes programming easier and makes user experience better

* Consistency options with ALPS
  * Linearizability -> Impossible
  * Serializability -> Impossible
  * Causal -> Don't settle for eventual consistency
  * Eventual -> e.g., Amazon Dynamo, Facebook/Apache Cassandra
* Causality simplifies programming
* Causal Consistency vs. Eventual Consistency
  * Causal Consistency requires all values returned by reads to be consistent with all potential causality relationships (partial ordering)
    * No potential causality means logically concurrent
  * Eventual consistency does not strive to maximize potential causality
* Achieving good consistency given ALPS (high-level)
  * Assign version IDs (logical clock or physical clock)
    * version ordering is consistent with potential causality, and
    * resolves conflicts deterministically
  * Replicas log an order
    * Record version info
  * Replicas converge
    * Exchange logs, select a deterministic ordering, and apply it