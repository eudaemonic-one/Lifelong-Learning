# Lecture 18 Virtual Memory: Systems

## Simple Memory System Example

* Addressing
  * 14-bit VA
  * 12-bit PA
  * Page size = 64 bytes
  * TLB
    * 16-entries
    * 4-way associative

![image-20200410012451123](images/lecture18-vm-systems/simple_memory_systen_tlb.png)

![image-20200410014740492](images/lecture18-vm-systems/address_translation_example_tlb_cache_miss.png)

## Case Study: Core i&/Linux Memory System

### End-to-end Core i7 Address Translation

![image-20200410015153897](images/lecture18-vm-systems/end_to_end_core_i7_address_translation.png)

* **Core i7 Level 1-3 Page Table Entries**

![image-20200410015100235](images/lecture18-vm-systems/core_i7_level_1_3_page_table_entries.png)

### Linux Organizes VM as Collection of Areas

![image-20200410015724913](images/lecture18-vm-systems/linux_organizes_vm_as_collection_of_areas.png)

### Linux Page Fault Handling

![image-20200410015847807](images/lecture18-vm-systems/linux_page_fault_handling.png)

## Memory Mapping

* VM areas initialized by associating them with disk objects
  * Called **memory mapping**
* Area can be backed by (i.e., get its initial values from) :
  * Regular file on disk (e.g., an executable object file)
  * Anonymous file (e.g., nothing)
    * First fault will allocate a physical page full of 0's (demand-zero page)
    * Once the page is written to (dirtied), it is like any other page
* Dirty pages are copied back and forth between memory and a special **swap file**

### Private Copy-on-write (COW) Objects

* Two processes mapping a private **copy-on-write (COW)** object
* Area flagged as private copy-on- write
* PTEs in private areas are flagged as read-only
* Instruction writing to private page triggers protection fault
* Handler creates new R/W page
* Instruction restarts upon handler return
* Copying deferred as long as possible

### Finding Shareable Pages

* **Kernel Same-Page Merging**
  * OS scans through all of physical memory, looking for duplicate pages
  * When found, merge into single copy, marked as copy-on-write
  * Especially useful when processor running many virtual machines

### User-Level Memory Mapping

* Uses of mmap
  * Reading big files
* Shared data structures (when called with `MAP_SHARED` flag)
  * File-based data structures
