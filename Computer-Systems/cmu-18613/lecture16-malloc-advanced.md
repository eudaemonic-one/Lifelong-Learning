# Lecture 16 Malloc Advanced

## Explicit Free Lists

* Maintain list(s) of free blocks, not all blocks
* The “next” free block could be anywhere
  * So we need to store forward/back pointers, not just sizes
* Still need boundary tags for coalescing
  * To find adjacent blocks according to memory order
* Insertion policy
  * LIFO
  * FIFO
  * Address-ordered policy
    * Insert freed blocks so that free list blocks are always in address order: addr(prev) < addr(curr) < addr(next)
* Use circular, doubly-linked list
* Comparing to implicit list
  * Allocate is linear time in number of free blocks instead of all blocks
  * Much faster when most of the memory is full
  * Slightly more complicated allocate and free because need to splice blocks in and out of the list
  * Some extra space for the links (2 extra words needed for each block)

## Segregated Free Lists

* Each size class of blocks has its own free list
* Often have separate classes for each small size
* For larger sizes: One class for each size $[2^i+1, 2^{i+1}]$
* To allocate a block of size n:
  * Search appropriate free list for block of size $m > n$ (i.e., first fit)
  * If an appropriate block is found:
    * Split block and place fragment on appropriate list
    * If no block is found, try next larger class
  * Repeat until block is found
* If no block is found
  * Request additional heap memory from OS (using `sbrk()`)
  * Allocate block of n bytes from this new memory
  * Place remainder as a single free block in appropriate size class
* To free a block:
  * Coalesce and place on appropriate list

### Garbage Collection

* **Garbage collection** automatic reclamation of heap-allocated storage—application never has to explicitly free memory
* How does the memory manager know when memory can be freed?
  * We can tell that certain blocks cannot be used if there are no pointers to them
* Must make certain assumptions about pointers
  * Memory manager can distinguish pointers from non-pointers
  * All pointers point to the start of a block
  * Cannot hide pointers
* GC Algorithms
  * Mark-and-sweep collection
  * Reference counting
  * Copying collection
  * Generational Collectors
    * Collection based on lifetimes
    * Most allocations become garbage very soon
    * So focus reclamation work on zones of memory recently allocated

### Memory as a Graph

* We view memory as a directed graph
  * Each block is a node in the graph
  * Each pointer is an edge in the graph
  * Locations not in the heap that contain pointers into the heap are called root nodes

## Memory-Related Perils and Pitfalls

* Dereferencing bad pointers
* Reading uninitialized memory
* Overwriting memory
* Referencing nonexistent variables
* Freeing blocks multiple times
* Referencing freed blocks
* Failing to free blocks