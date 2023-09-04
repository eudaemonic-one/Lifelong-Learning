# Chapter 3. Storage and Retrieval

## Data Structures That Power Your Database

log == an append-only sequence of records; might be binary and intended only for other programs to read

index == additional structure deriving from primary data; causing overhead

> An important tradeoff in storage systems: well-chosen indexes speed up read queries, but every index slows down writes.

### Hash Indexes

key-value stores == dictionary == hash table

The simple data storage consists only of appending to a file.

In order to avoid eventually running out of disk space due to over appending to a file, either break the log into segments (aka snapshot of current log file) of a certain size and/or perform compaction on segments. The merging and compaction can be done in a background thread. Each segment has its own in-memory hash table, mapping keys to file offsets.

Real issues in practice: file format, deleting records (tombstone), crash recovery, partially written records, concurrency control

Append-only (sequential write operations are faster than random writes; concurrency and crash recovery are mush simpler if segment files are append-only or immutable) vs. Override in-place (less heap memory space consumption; range queries are more efficient)

### SSTables and LSM-Trees

### Constructing and maintaining SSTables

> An SSTable is a key-sorted append-only key-value storage (each block is an immutable, sorted file of key-value pairs, each SSTable is chunks of data blocks plus an index of block ranges).

Sorted String Table (*SSTable*): always require the sequence of key-value pairs is *sorted by key*.

Maintaining a sorted structure in memory with B-Trees or Red-Black Trees or AVL Trees

* When a write comes in, add it to an in-memory balanced tree data structure (*memtable*).
* When the memtable gets bigger than some threshold - typically a few megabytes - write it out to disk as an SSTable file.
* In order to serve a read request, first try to find the key in the memtable, then in the most recent on-disk segment, then in next-older segment, etc.
* Also, range queries would be able to perform efficient scan on sorted key range.
* From time to time, run a merging and compaction process in the background to combine segment files and to discard overwritten or deleted values.

However, the scheme suffers from databases crashes, the most recent writes are lost. Thus, a separate log on disk to which every write is immediately appended is needed to be able to restore the memtable after a crash.

* When write comes in, the data is firstly added to a write-ahead log (*WAL*) file (*journal*) and flushed to the disk, which allows crash recovery even before persisting the in-memory data structure to disk. These writes could be fast as they are sequential writes and also batching or periodic rolling could apply here.
* Then, update the memtable.
* Once memtable reaches its size threshold, flush onto disk and create a new segment file.

### Making an LSM-tree out of SSTables

> An LSM-Tree is a two-layer data structure, new records are inserted into memory-resident component and sorted by a balanced tree; if exceeds a certain size threshold, flush memtable into disk as segment file.

Log-Structured Merge-Tree (*LSM-Tree*): provides high write throughput through performing only "in-memory" write operation (in contrast to SSTable updates are done to disk which can trigger expensive update to indexes).

Tail latencies use cases: when looking up keys that do not exist in the database, you have to check the memtable, then the segments all the way back to the oldest. Storage engines often use additional *Bloom filters* (a memory-efficient data structure for approximating the contents of a set), which can tell you if a key does no appear in the database.

> reference post 1: https://stackoverflow.com/questions/58168809/what-is-the-differences-between-the-term-sstable-and-lsm-tree
> reference post 2: https://medium.com/swlh/log-structured-merge-trees-9c8e2bea89e8
> reference post 3: https://rahulpradeep.medium.com/sstables-and-lsm-trees-5ba6c5529325

### B-Trees

B-trees break the database down into fixed-sized blocks or pages. Each page contains several keys (number of *branching factor*) and references to child pages. The leaf page contains the value inline or references to the pages where the values can be found.

When adding a new key, find the page whose range encompasses the new key and add it to that page. If there isn't enough free space, split it into two half-full pages, and the parent page is updated to account for the new subdivision of key ranges.

A four-level tree of 4KB pages with a branching factor of 500 can store up to 250 TB.

B-trees do override page in-place. Also, one update could cause several subsequent write operations for parent pages.Thus to make the database resilient to crashes, it is common to include addition data structure on disk: *write-ahead log* (WAL, aka *redo log*). It could restore the B-tree back to a consistent state once needed.

Also, *latches* (lightweight locks) is needed to avoid multiple threads overriding the same page and cause inconsistent state.

Locks (protect index logical content from other transactions; held for transaction duration; need to be able to rollback changes) vs. Latches (protect only a section of index's internal data structure; held for operation duration; do not need to be able to rollback changes)

Other B-trees optimizations:

* copy-on-write scheme: a modified page is written to a different location, and a new version of the parent pages in the tree is created, pointing at new location. multiple processes sharing the same page until the page is going to be modified by one modifier.
* abbreviating keys to save space in pages.
* lay out leaf pages in sequential order on disk even difficult to maintain order as the tree grows; by contrast, LSM-trees rewrite large segments of the storage in one go during merging and thus easy to keep sequential keys.
* additional pointers added to leaf page pointing to sibling pages, which allows keys in order without jumping back to parent pages.

### Comparing B-Trees and LSM-Trees

> LSM-trees are typically faster for writes, whereas B-trees are typically faster for read.

B-tree index must write every piece of data at least twice (WAL log, tree page, perhaps parent pages). Also, perhaps only a few bytes in that page is needed.

Log-structured indexes also rewrite data multiple times (*write amplification*) over the course of the database's lifetime. Write amplification has a direct performance cost if write-heavy loads become bottleneck.

LSM-trees sustain higher write throughput than B-trees because they sometimes have lower write amplification and sequence write compact files helps. LSM-trees can be compressed into smaller files on disks; while B-tree leaves fragmentation.

Log-structured storage compaction could interfere with the performance of ongoing reads and writes. The disk's finite write bandwidth is shared between initial write (log and memtable flushing) and the compaction threads.

B-trees's each key exists in exactly one place in the index, which provide more robust transaction isolation using locks on range of keys.

### Other Indexing Structures

*secondary indexes* can be constructed from a key-value index; indexed values might not be unique.

The "key" in an index is the thing that queries search for. The value could simply be a reference to the row stored elsewhere (*heap file*). The heap file could avoid duplicating data when multiple secondary indexes are present.

To store the indexed row value directly within an index is known as *clustered index*.

*covering index* or *index with included columns* could store some table columns so that queries can read directly using the index.

*multi-column indexes*: *concatenated index* combines several fields into one key by appending one column to another by appending one column to another.

*full-text search* usually consists of small dictionary-like in-memory index or sparse collection of some keys.

*fuzzy querying* searches text for words within a certain edit distance.

*in-memory databases* could keep datasets entirely in memory or potentially distributed across several machines. Meanwhile, append-only log in disk is only for durability. It also avoid the overheads of encoding in-memory data structures in a form that can be written to disk.

*in-memory data structures* offer database interface like interfaces for priority queues or sets.

*anti-caching* evicts the least recently used data from memory to disk when there is not enough memory, and loading it back when it it access in the future. It could support datasets larger than the availability memory.

## Transaction Processing or Analytics?

transaction == look up a small number of records by some key, using an index; records are inserted or updated => *online transaction processing* (OLTP)

*data analytics* == scan over a huge number of records, only reading a few columns per record, and calculates aggregate statistics => *online analytics processing* (OLAP)

### Data Warehousing

data warehouse: separate database that analysts can query without affecting OLTP operations.

Extract-Transform-Load (ETL): loading data from databases and transforming into analysis-friendly schema.

### Schemas for Analytics

*star schema* (*dimensional modeling*): *fact table* represents individual events; *dimension tables* represent who, what, where, when, how, and why of the event.

*snowflake schema*: dimensions are further broken own into sub-dimensions.

## Column-Oriented Storage

row-oriented vs. column-oriented

### Column Compression

> Column-oriented storage lends itself very well to compression.

*bitmap encoding*: using bits to represent *n* distinct values.

*Memory bandwidth and vectorized processing*: the query engine can take a chunk of compressed column data that fits in the CPU's L1 cache and iterate through it in a tight loop (with no function calls).

### Sort Order in Column Storage

sorted order => sort by date key, tie-break by another sort key => help with compression of columns further

several different sort orders => add data replicates while boosting other query patterns

### Writing to Column-Oriented Storage

LSM-tree is good fit for column writes, all writes go to an in-memory store, when enough writes have accumulated, merge with column files on disk.

### Aggregation: Data Cubes and Materialized Views

*materialized aggregates*: cache some counts or sums that queries use most often

*materialized view*: a table-like object whose contents are the results of some query.

*data cube* (*OLAP cube*): a grid of aggregates grouped by different dimensions => only serve as a performance boost for certain precomputed queries; no finer-granularity visibility into datasets.
