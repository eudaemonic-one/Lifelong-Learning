# Chapter 6. Partitioning

*partition* == *sharding* => each piece of data belongs to exactly one partition => scalability => dataset and query load can be distributed across many processors.

## Partitioning and Replication

> Partitioning is usually combined with replication so that copies of each partition are stored on multiple nodes.

## Partitioning of Key-Value Data

> The goal of partitioning is to spread the data and the query load evenly across nodes.

*skewed* == *hot spot*

### Partitioning by Key Range

One way of partitioning is to assign a continuous range of keys to each partition.

### Partitioning by Hash of Key

> A good hash function takes skewed data and makes it uniformly distributed.

partitioning => unable to do efficient range queries => concatenated index can support efficient rang scan over columns other than the partition key.

### Skewed Workloads and Relieving Hot Spots

How to compensate highly skewed workload? Maybe add a random prefix or suffix to the beginning or end of the key?

## Partitioning and Secondary Indexes

> A secondary index usually doesn’t identify a record uniquely but rather is a way of searching for occurrences of a particular value.

### Partitioning Secondary Indexes by Document

*document-partitioned index* (*local index*): each partition maintains its own secondary indexes, covering only the documents in that partition.

Querying a partitioned database is known as *scatter*/*gather* (tail latency amplification).

### Partitioning Secondary Indexes by Term

*term-partitioned index*: terms (full-text or numerics) determine the partition of the index; a global index being partitioned differently from the primary key index.

The downside of a global index is that writes are slower and more complicated, because a write to a single document may now affect multiple partitions of the index (every term in the document might be on a different partition, on a different node).

In practice, updates to global secondary indexes are often asynchronous.

## Rebalancing Partitions

*rebalancing*: moving load from one node in the cluster to another.

### Strategies for Rebalancing

*hash mod N*: nodes addition/removal cause frequent rebalancing and data moves.

*fixed number of partitions*: create many more partitions than there are nodes, and assign several partitions to each node. If a node is added to the cluster, simply *steal* a few partitions from every existing node until partitions are fairly distributed.

> By assigning more partitions to nodes that are more powerful, you can force those nodes to take a greater share of the load.

*dynamic partitioning*: when a partition grows to exceed a configured size, it is split into two partitions so that approximately half of the data ends up on each side of the split. Each partition is assigned to one node, and each node can handle multiple partitions. *pre-splitting* allows an initial set of partitions to be configured on an empty database.

*partitioning proportionally to nodes*: have a fixed number of partitions per node. Picking partition boundaries randomly requires that hash-based partitioning is used.

### Operations: Automatic or Manual Rebalancing

> Rebalancing is an expensive operation. This process can overload the network or the nodes and harm the performance of other requests while the rebalancing is in progress.

## Request Routing

*service discovery* have a few approaches:

1. Allow clients to contact any node.
2. Send all requests from clients to a routing tier first, which determines the node that should handle each request and forwards it accordingly.
3. Require that clients be aware of the partitioning and the assignment of partitions to nodes. In this case, a client can connect directly to the appropriate node, without any intermediary.

*separate coordination service* (such as ZooKeeper): keep track of cluster metadata and maintain the authoritative mapping of partitions to nodes. Whenever a partition changes ownership, or a node is added or removed, ZooKeeper notifies the routing tier.

*gossip protocol*: disseminate any changes among the nodes in cluster.

### Parallel Query Execution

*massively parallel processing* (MPP): query optimizer breaks complex query into a number of execution stages and partitions to be executed in parallel on different nodes of the database cluster.
