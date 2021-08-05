# Lecture 8 Cloud Storage 2

## Cloud Storage Options

### Provide a Tradtional Filesystem

* The OS running in each VM mounts file service
* e.g., NFS, AFS, GFS, HDFS

### Provide Block Stores (Virtual Disks)

* A common option in VM-based environments
  * Guest OS running in a VM  has code for FSs on disks
  * So give it a disk to use
* Virtual disk looks to guest OS just like real disk
  * Same interface: read/write of fixed-size blocks, ID'd by block number
  * Guest OS can format it, implement FS atop it
    * VMM makes guest OS disk operations access the right content
* Most cloud infrastructures have this option
  * e.g., AWS Elastic Block Store (EBS), OpenStack Cinder
* Guest OS may or may not know whether virtual disk is local
  * Non-local interface: network-disk interface
  * Local interface: VMM translates to other protocol as needed
* Virtual Disks often implemented as files (sequences of fixed-sized blocks)
* Thin provisioning
  * Promise more space than you have
  * Allocate physical space only for blocks that get written
  * Can benefit from TRIM
* Performance interference
  * Each VM may have a virtual disk
  * We expect time-sharing to have fairness/QoS

### Provide "Union" Filesystems

* A common option in container-based environments
  * Container runs atop OS
  * Container is given access to file system (thinks it has entire FS via chroot)
  * Needs some "system-wide" and some "private" files
* Concept: Make a single FS view from multiple FSs
  * Show contents of a directory as merge of several
    * Some read-only (system-wide) and some not (private)
  * Implemented by a layer atop the individual FSs
    * For read-only: Look in first FS first, then second if needed
    * For create/write: put into first non-read-only FS

### Provide "Object" Store

* A common option in large clouds
  * A simplified, generic "file" storage system
    * Like files, objects are sequences of bytes
    * Named by non-hierarchical alphanumerical objects IDs
  * e.g., AWS S3, Box, iCloud
* Usually limited interface and semantics
  * e.g., CRUD API: Create, Read, Update, Delete
  * Often assume single writer, sequential (or all-at-once)
    * No promises regarding: sharing/concurrency, interrupted writes