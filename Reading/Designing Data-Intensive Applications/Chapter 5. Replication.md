# Chapter 5. Replication

*Replication*: keeping a copy of the same data on multiple machines that are connected via a network. Several reasons may apply:

* To keep data geographically close to your users (and thus reduce access latency).
* To allow the system to continue working even if some of its parts have failed (and thus increase availability).
* To scale out the number of machines that can serve read queries (and thus increase read throughput).

## Leaders and Followers

*leader-based replication* (*active/passive* or *master-slave replication*):

* One of the replicas is designated the *leader* (*master* or *primary*). When clients want to write to the database, they must send their requests to the leader, which first writes the new data to its local storage.
* The other replicas are known as *followers* (*read replicas*, *slaves*, *secondaries*, or *hot standbys*). Whenever the leader writes new data to its local storage, it also sends the data change to all of its followers as part of a replication log or change stream. Each follower takes the log from the leader and updates its local copy of the database accordingly, by applying all writes in the same order as they were processed on the leader.
* When a client wants to read from the database, it can query either the leader or any of the followers.

### Synchronous Versus Asynchronous Replication

In practice, if you enable synchronous replication on a database, it usually means that one of the followers is synchronous, and the others are asynchronous. If the synchronous follower becomes unavailable or slow, one of the asynchronous followers is made synchronous. This guarantees that you have an up-to-date copy of the data on at least two nodes: the leader and one synchronous follower. This configuration is sometimes also called *semi-synchronous*.

Often, leader-based replication is configured to be completely asynchronous. In this case, if the leader fails and is not recoverable, any writes that have not yet been replicated to followers are lost.

*chain replication*:  a variant of synchronous replication that only applies between two replicas at one time.

### Setting Up New Followers

Setting up a follower can be done without downtime.

1. Take a consistent snapshot of the leader's database (similar to a backup).
2. Copy the snapshot to the new follower node.
3. The follower connects to the leader and requests all the data changes that have happened since the snapshot was taken.
4. When the follower has processed the backlog of data changes since the snapshot, we say it has caught up.

### Handling Node Outages

Follower failure: Catch-up recovery => If a follower crashes and is restarted, or if the network between the leader and the follower is temporarily interrupted, the follower can recover from its log of data changes it has received from the leader.

Leader failure: Failover => one of the followers needs to be promoted to be the new leader, clients need to be reconfigured to sent their writes to the new leader, and the other followers need to start consuming data changes from the new leader.

1. Determining that the leader has failed => timeout, heartbeat.
2. Choosing a new leader => election or appointed by a previously elected controller node.
3. Reconfiguring the system to use the new leader => the system needs to ensure that the old leader becomes a follower and recognizes the new leader.

Failover is fraught with things that can go wrong:

* If asynchronous replication is used, the new leader may not have received all the writes from the old leader before it failed.
* Discarding writes is especially dangerous if other storage systems outside of the database need to be coordinated with the database contents.
* It could happen that two nodes both believe that they are the leader (*split brain*). If both leaders accept writes, and there is no process for resolving conflicts, data is likely to be lost or corrupted. As a safety catch, some systems have a mechanism to shut down one node if two leaders are detected.
* What is the right timeout before the leader is declared dead?

### Implementation of Replication Logs

Statement-based replication: the leader logs every write request (*statement*) that it executes and sends that statement log to its followers.

* Any statement that calls a nondeterministic function, is likely to generate a different value on each replica.
* If statements use an auto-incrementing column, or if they depend on the existing data in the database, they must be executed in exactly the same order on each replica, or else they may have a different effect.

*Write-ahead log (WAL) shipping*: use exact same log to build a replica on another node; besides writing the log to disk, the leader also sends it across tne network to its followers.

If the replication protocol allows the follower to use a newer software version than the leader, you can perform a zero-downtime upgrade of the database software by first upgrading the followers and then performing a failover to make one of the upgraded nodes the new leader.

*Logical (row-based) log replication*: a sequence of records describing writes to database tables at the granularity of a row. A transaction that modifies several rows generates several such log records, followed by a record indicating that the transaction was committed. The logical log format is easier for *change data capture* and for external applications to parse.

*Trigger-based replication*: a trigger lets you register custom application code that is automatically executed when a data change occurs in a database system. The trigger has the opportunity to log this change into a separate table, from which it can be read by an external process.

## Problems with Replication Lag

> In *read-scaling* architecture, you can increase the capacity for serving read-only requests simply by adding more followers. Also, this approach only works with asynchronous replication.

*eventual consistency* vs. *replication lag*

### Reading Your Own Writes

*read-after-write consistency* (*read-your-writes consistency*): if the user reloads the page, they will always see any updates they submitted themselves.

* When reading something that the user may have modified, read it from the leader; otherwise, read it from a follower.
* You could track the time of the last update and, for one minute after the last update, make all reads from the leader.
* The client can remember the timestamp of its most recent write—then the system can ensure that the replica serving any reads for that user reflects updates at least until that timestamp.
* Another complication arises when the same user is accessing your service from multiple devices, for example a desktop web browser and a mobile app. In this case you may want to provide *cross-device read-after-write consistency*.
  * The code running on one device doesn’t know what updates have happened on the other device. This metadata will need to be centralized.
  * If your replicas are distributed across different datacenters, there is no guarantee that connections from different devices will be routed to the same datacenter.

### Monotonic Reads

When reading from asynchronous followers, it’s possible for a user to see things moving backward in time.

*monotonic reads*: less guarantee than strong consistency, but a stronger guarantee than eventual consistency. One way is to make sure each user always makes their reads from the same replica.

### Consistent Prefix Reads

*consistent prefix reads*: guarantees *causality* or say *causal dependencies*, if a sequence of writes happens in a certain order, then anyone reading those writes will see them appear in the same order.

Make sure that any writes that are causally related to each other are written to the same partition could avoid the problem of no global ordering of writes across partitions in distributed (sharded) databases.

### Solutions for Replication Log

*transaction* => provide stronger guarantees so that application can be simpler.

> In the move to distributed (replicated and partitioned) databases, many systems have abandoned transactions, claiming that transactions are too expensive in terms of performance and availability, and asserting that eventual consistency is inevitable in a scalable system.
