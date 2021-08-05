# Lecture 14 Scheduling Computation I

## Scheduling

* Map collections of jobs onto set of machines
* Basic building blocks
  * Central scheduler: receives requests, tracks and allocates resources
  * Per-machine agent: runs jobs, enforces allocations, monitor usage
* Lots of differences among different schedulers

### Machine Checkout

* Each job want any one full machine
* User looks at list of available machines and picks one
  * Edits the list to indicate that it is no longer free
* Shared list -> Centralized server

### Extension #1: Scheduler Allocates for and Runs Jobs

* User submits job to system
  * might be a VM image or some executable script/program
* Schedular picks machine and runs the job
  * still requires free machine list
  * also requires ability to start the job on the chosen machine
    * e.g., send to VMM or to scheduling agent that executes on the machine
* When the job finishes
  * the machine frees itself, by telling the scheduler

### Extension #2: Packing Multiple onto a Machine

* User submits job plus resource request (parts of one machine)
  * e.g., RAM capacity and CPU fraction (in MHz or cores)
* Scheduler picks a machine with enough resources and runs job on it
  * must now track what portion of each machine is allocated versus free
  * picking machine is somewhat akin to memory allocation
    * options like first fit, best fit, etc. apply
* Assumptions
  * local machine agent ensures allocation fractions
  * interference among jobs on a machine can be ignored

### Extension #3: Packing with Uncertainty

* User's resource request can be imperfect
  * common to ask for more than needed
  * can often use more, if available, as well
* Overcommitting
  * monitor resource usage, identify under-utilization of allocation, and use it
  * assign more total allocation to a machine than would fit
  * biggest issue: deal with situations where resources run out
    * options: kill or migrate that job, kill or migrate a different job, shrink allocation
* Using slack resources
  * imaging that only 1/2 of the CPU has been allocated to jobs so far

### Extension #4: Informing Decisions re: Uncertainty

* User provides more information than just the resource request
  * scheduler and per-machine agent use it
* VMware extra information
  * Reservation: guaranteed minimum amount (say "no" if can't promise)
  * Limit: upper bound (don't use extra resources beyond certain amount)
  * Share: relative importance of different jobs (when sharing extra resources)

### Extension #5: Machines Not All The Same

* Few data centers / clouds have a single machine type
  * different amounts of RAM, different CPU speeds, core counts, etc
  * could be special features (e.g., GPU) only present on some of them
* Sceduler still works in largely the same way
  * special features require pruning set of options considered
* Interesting nuance: exposing vs. hiding machine difference

### Extension #6: Changing Previous Decisions

* Free resources can become fragmented or poorly distributed
  * as jobs finish at arbitrary times that often cannot be known
  * may be enough resources for a new job, but not all together
  * over-committing or slack usage may be improvable
* Changing decisions requires work
  * the job must be moved, somehow, from one machine to another
    * including tradeoff between short-term cost vs. long-term benefit
  * primary options: migration or "shoot-and-restart"
    * both take time and resource from doing real work

### Extension #7: Non-resource Constraints

* For some jobs, there are additional concerns to be addressed
  * e.g., being close to or not being close to another job
* VMware constriant examples
  * Affinity: identifies VMs that would benefit from being on same machine
    * to allow for faster communication
  * Anti-affinity: identifies VMs that must not be on same machine
    * to ensure that a machine crash does not disable both
* Constriants more generally
  * can be any machine attributes, though scheduler+user must understand
  * restricts the set of options that the scheduler can consider for a given job
  * also, affinity and anti-affinity can relate to more than just "same machine"

### Extension #578: Multi-machine Jobs

* It is not uncommon for a request to ask for several machines at once
  * e.g., to run Hadoop instance or a 3-tier web service
* Scheduler considers the request as a whole
  * most schedulers will wait until can schedule the entire thing
    * so, it needs nough free resources fitting constraints at the same time
    * some schedulers will give whatever subset it can, ASAP, rather than waiting
  * may also try to improve assigments based on knowing the full set
    * e.g., run them on same machine or rack
* Interesting nuance: to hoard or not to hoard
  * large requests may wait forever, if the scheduler just waits to get lucky
  * can "hold back" resources, as they become free, until enough are free