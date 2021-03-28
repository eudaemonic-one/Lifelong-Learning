# Guest Lecture 11 Edge Computing: Where Mobile Meets the Cloud

## Why Edge Computing?

### Diverse Cloud-based Mobile APps

* Many are small front-end applets
* Push beyond capabilities of devices alone
  * Leverage computing in cloud, big data
* Interactive applications
* Need processing capacity of clouds, but with strict timing requirements
* May require rethinking cloud architecture
  * -> Edge Computing

### Do We Really Need to Offload?

* Today's smartphone are getting more powerful
* Advanced algorithms running on phones already
* But, computation varies dramatically depend on operational conditions

### Energy Limits on Mobiles

* Even if mobile fast enough, may not want to spend the energy
* Battery technology advancing slowly compared to CPU, networks, storage
* Performance of mobile often limited by thermals

## Cloud Location Really Matters

* Proximity to the data center is essential
  * Response time & Energy consumption
* Local Cloud infrastructure on LAN /WiFi
  * Better attainable proximity
* Contradictory requirements
  * Local clouds are valuable for mobile computing
  * However, it needs many data centers at the edges the Internet

## Alternative Cloud Architecture

### Cloudlet

* Local cloud infrastructure
  * On LAN / WLAN
  * Hight bandwidth, low latency relative to cloud
* Provides IaaS - like a mini AWS datacenter
* Amenale to decentralized, incremental deployment - like WiFi

### Alternative Cloud Architecture

* Level 1
  * Today's unmodified cloud
* Level 2
  * Stateless data centers at the edge
  * Cloudlet
  * Appliance-like deployment model
* Common requirements
  * Strong isolation between untrusted user-level computations
  * Mechanism for authentications, access control and metering
* Physical realizations

## Edge Computing vs. Cloud

* Edge is best for local use
  * So need to find the best cloudlet for a particular user
* Cloud vs. Cloudlet provisioning
  * One time cost in Cloud
  * Can't have all applications running in all cloudlets all the time
  * Dynamic, on-use provisioning is needed at edge
* User mobility
  * Not critical at granularity of large data center service area
  * May need to migrate services between cloudlets as user moves

## VM Synthesis

* VM synthesis: dividing a custom VM into two pieces
  * Base VM: Vanilla OS that contains kernel and basic libraries
  * VM overlay: A binary patch that contains customized parts
* Deduplication
  * Remove redundancy in the VM overlay
* Reducing Semantic Gap
  * Include only the state that actually matters to the guest OS
  * Use TRIM support for disk and inspect free page lists for memory
* Pipelining
  * Transfer VM overlay
  * Decompress
  * Apply delta
* Early Start
  * Start executing VM before synthesis completes
* Shown to be effective for rapid provisioning of cloudlet resources
* Overcomes one of the key problems with a VM-based offload solution

### Live VM Migration in Data Centers

* Downtime is the ultimate performance metric
* Convential live migration is optimized for datacenter environment
* Problem: Often works poorly across WANs

### What About Containers

* Containers are glorified processes
  * Private file system, with custom environment
  * Shared kernel, drivers with host environment
  * Some resource isolation through cgroups, namespaces in Linux kernel
* Docker
  * Uses layered filesystem model (like overlays)
* Cons:
  * Not as general (OS, kernel dependency)
  * Less isolation than VMs
  * Migration is hard - process state within kernel

### Bandwidth Issues

* WAN bandwidth to remote datacenter can be limiting factor
* Shared edge computing resources can help

## Edge Computing

* e.g., Process live frames
  * Focus on DNN, computer vision algorithms
  * Store at edge, forward only as needed to cloud
* Compute resources limited at Edge
* "Mainstream" project - exploit common practice of fine-tuning pretrained DNNs
  * Share compute when DNNs use same base model, leave some layers unmodified
  * Build training, runtime system to maximize sharing
* "FilterForward" project
  * Perform expensive analysis in Cloud
  * Filter out all but most relevant frames for particular task