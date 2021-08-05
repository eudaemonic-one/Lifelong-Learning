# Lecture 1 Overview and Flash SSD Operation

## Storage Systems

* **Memory/storage hierarchy**
  * Balancing performance with cost
    * Small memories are fast but expensive
    * Large memories are slow but cheap
  * Exploit locality to get the best of both worlds
    * locality = re-use/nearness of accesses
    * allows most accesses to use small, fast memory
  * Locality is a general concept
    * power, BW, management
* **Persistence**
  * Storing data for lengthy periods of time
    * despite component failures, SW bugs/upgrades, human error
  * To be useful, it must also be possible to find it again later
    * this brings in data organization, consistency, and management issues
  * This is where the serious action is
    * and it does relate to the memory/storage hierarchy

## What is a Storage System

* Software and Hardware that collectively provide application writers with an infrastructure that retains persistent data and provides it on request
* Logical: Program <--> File System <- -> Controller Firmware
* Physical: Client <--> Server <--> I/O Controller <--> Device ASIC
* Storage Software Interfaces
  * Program <--> Physical Media
  * Naming and address mapping (e.g., relocation)
  * Caching and request transformations (e.g., combining)

|                        |                           |
| ---------------------- | ------------------------- |
| Program                | <File Name, Offset>       |
| File System (Database) | <Partition, Block #>      |
| Device Driver          | <Device #, LBN #>         |
| I/O Controller         | <Cylinder, Track, Sector> |
| Disk Media             |                           |

