# Lecture 10 The Memory Hierarchy

## Write & Read Memory

* Write
  * Transfer data from CPU to memory
  * Store
* Read
  * Transfer data from memory to CPU
  * Load

## Bus

* A **bus** is a collection of parallel wires that carry address, data, and control signals

## Random-Access Memory (RAM)

* Basic storage unit is normally a cell (one bit per cell)
* Static RAM
  * 6 transistors / bit
  * Holds state indefinitely
  * Scales with semiconductor technology
* Dynamic RAM
  * 1 Transistor + 1 Capacitor / bit
  * Must refresh state periodically
  * Limited by need for minimum capacitance (Aspect ratio how deep can make capacitor)
  * Operation of DRAM cell has not changed since its invention
  * DRAM cores with better interface logic and faster I/O
* Conventional DRAM Organization
  * $d \times w$ DRAM
  * $d ⋅ w$ total bits organized as d supercells of size w bits
  * Reading DRAM Supercell (2,1)
    * Row access strobe (RAS) selects row 2
    * Row 2 copied from DRAM array to row buffer
    * Column access strobe (CAS) selects column 1
    * Supercell (2,1) copied from buffer to data lines, and eventually
      back to the CPU
    * All data written back to row to provide refresh

## Locality

* **Principle of Locality**: Programs tend to use data and instructions with addresses near or equal to those they have used recently
* **Temporal locality**: Recently referenced items are likely
  to be referenced again in the near future
* **Spatial locality**: Items with nearby addresses tend
  to be referenced close together in time

## Memory Hierarchy

* L0: Regs
* L1: L1 cache (SRAM)
* L2: L2 cache (SRAM)
* L3: L3 cache (SRAM)
* L4: Main memory (DRAM)
* L5: Local secondary storage (local disks)
* L6: Remote secondary storage (e.g. Web servers)

## Caches

* **Cache**: A smaller, faster storage device that acts as a staging area for a subset of the data in a larger, slower device
* For each k, the faster, smaller device at level k serves as a cache for the larger, slower device at level k+1
* **Big Idea (Ideal)**: The memory hierarchy creates a large pool of storage that costs as much as the cheap storage near the bottom, but that serves data to programs at the rate of the fast storage near the top
* **3 Types of Cache Misses**
  * Cold (compulsory) miss
  * Capacity miss
  * Conflict miss

## Disk

* Disk Geometry
  * Disks consist of platters, each with two surfaces
  * Each surface consists of concentric rings called tracks
  * Each track consists of sectors separated by gaps
* Disk Capacity
  * maximum number of bits that can be stored
  * Recording density
  * Track density
  * Areal density
* Disk Access Time
  * Taccess = Tavg seek + Tavg rotation + Tavg transfer
  * Seek time: Time to position heads over cylinder containing target sector
  * Rotational latency: Time waiting for first bit of target sector to pass under r/w head
    * Tavg rotation = 1/2 × 1/RPMs × 60 secs / 1 min
  * Transfer time: Time to read the bits in the target sector
    * Tavg transfer = 1 / RPM × 1/(avg # sectors/track) × 60 secs / 1 min

