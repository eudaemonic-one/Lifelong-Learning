# Cache Memories

## Cache Concepts

* Hit: Block b is in cache
* Miss: Block b is not in cache
  * Cold (compulsory) miss
    * Occurs because the cache starts empty
  * Capcity miss
    * Occurs when the set of active cache blocks (working set) is larger than the cache
  * Conflict miss
    * Occur when the level k cache is large enough, but multiple data objects all map to the same level k block
* **Cache memories** are small, fast SRAM‐based memories managed automatically in hardware
* CPU looks first for data in cache

## General Cache Organization (S, E, B)

* Cache size = $S * E * B$ data bytes
* S = $2^s$ sets; E = $2^e$ lines per set; $B = 2^b$ bytes per cache block
* In cache line
  * valid bit v
  * tag
  * Block ($2^b$ bytes per cache block)
* **Cache Read**
  * Locate set
  * Check if any line in set has matching tag
  * Yes + line valid: hit
  * Locate data starting at offset
* **Address of word**
  * t bits tag
  * s bits set index
  * b bits block offset
* Valid and match (tag) = hit
* Not valid or no match = miss
  * One line in set is selected for eviction and replacement
  * Replacement policies: random, least recently used (LRU)
* **Cache Write**
  * valid bit + dirty bit + tag + B = $2^b$ bytes
  * Write Hit
    * Write through (write immediately to memory)
    * Write back (defer write to memory until replacement of line)
      * Each cache line needs a dirty bit (set if data differs from memory)
  * Write miss
    * Write allocate (load into cache, update line in cache)
      * Good if more writes to the location will follow
    * No write allocate (writes straight to memory, does not load into cache)
  * Typical
    * Write through + No write allocate
    * **Write back + Write allocate**
  * Indexing Approaches
    * Middle Bit Indexing TTSSBB
      * Makes good use of spatial locality
    * Hight Bit Indexing SSTTBB
      * Program with high spatial locality would generate lots of conflicts

## Intel Core i7 Cache Hierarchy

* Core
  * Regs
  * L1 d-cache (for data) L1 i-cache (for instruction)
  * L2 unified cache
  * L3 unified cache (shared by all cores)
  * Main memory

## Cache Performance Metrics

* Miss rate
  * Fraction of memory references not found in cache (misses / accesses) = 1 – hit rate
  * Typical numbers: 3-10% for L1, <1% for L2
* Hit time
  * Time to deliver a line in the cache to the processor
  * Typical numbers: 4 clock cycle for L1, 10 clock cycles for L2
* Miss Penalty
  * Additional time required because of a miss
  * Typically 50-200 cycles for main memory
* **“miss rate” is used instead of “hit rate”**
* Writing Cache Friendly Code
  * Make the common case go fast
  * Minimize the misses in the inner loops

## The Memory Mountain

* Read Throughput (read bandwidth)
* **Memory mountain** measures read throughput as a function of spatial and temporal locality
* Call test() once to warm up the caches and call it again and measure the read throughput (MB/s)

