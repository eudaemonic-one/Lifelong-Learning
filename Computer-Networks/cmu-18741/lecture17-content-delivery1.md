# Lecture 17 Content Delivery 1

## Web Content Delivery

* Typical workload -> Multiple (small) objects per page
  * File sizes heavily tailed
  * 3-way handshake
  * Lots of slow start
  * Extra connection state
* HTTP History (0.9/1.0->1.1->2.0)
* Web proxy caches
* HTTP Caching
  * Clients often cache documents
    * How to update documents?
      * `If-Modified-Since` requests to check
    * When/how often should the original be checked for changes?
      * Use `Expires` header
* Problems
  * Fraction of HTTP objects that are cacheable is dropping
  * Dynamic data

## Peer-to-Peer

* P2P system leverages the resource of client machines (peers)
* Why P2P?
  * Harness lots of spare capacity
  * Build very large-scale, self-managing systems

### Common P2P Framework

* Common primitives:
  * Join: how do I begin participating?
  * Publish: how do I advertise my file?
  * Search: how do I find a file?
  * Fetch: how do I retrieve a file?
* Search tends to be the most challenging

### Napster: Centralized Database

* Join: on startup, client contacts central server
* Publish: report list of files to central server
* Search: query the server -> return someone that stores the requested file
* Fetch: get the file directly from peer

### Gnutella: Flooding

* Join: on startup, client contacts a few other nodes; these become its neighbors
* Publish: no need
* Search: ask neighbors, who ask their neighbors, and so on (with TTL limit propagation) -> if found, reply to sender
* Fetch: get the file directly from peer

### KaZaA: Intelligent Query Flooding

* Join: on startup, client contacts a supernode - may at some point become one itself
* Publish: send list of files to supernode
* Search: send query to supernode, supernodes flood query amongst themselves
* Fetch: get the file directly from peers; can fetch simultaneously from multiple peers
  * Use UUHash to distinguish files with hash
* Why supernodes?
  * Query consolidation
  * Caching effect
  * Considers node heterogeneity
* Supernode selection is time-based

### BitTorrent: Swarming

* Focused on efficient fetching, not searching
* Join: contact central "tracker" server for list of peers
* Publish: Run a tracker server
* Search: Find a tracker out-of-band for a file
* Fetch: Download chunks of the file from your peers; upload chunks you have to them
* Sharing strategy
  * Employ Tit-for-tat sharing strategy
    * A is downloading from some other people
    * A will let the fastest N of those download from him
    * Be optimistic: occasionally let freeloaders download
      * Otherwise no one would ever start
      * Also allows you to discover better peers to download from when they reciprocate
  * Goal: Pareto efficiency

### Why Are P2P Useful

* Caching and soft-state data 
  * Works well: use peers as caches for hot data
* Finding read-only data
  * Limited flooding finds hay
  * DHTs find needles

### Writable, Persistent P2P

* Do you trust your data to 100,000 monkeys?
* Node availability hurts
  * When someone goes away, you must replicate the data they hold
  * Hard drives are huge, but cable modem upload bandwidth is tiny