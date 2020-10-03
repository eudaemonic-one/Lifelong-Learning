# Lecture 8 Disk Array Systems

## Reconstructing Failed Drive Data

### Reliability of Disk Arrays with No Repair

* $MTBF = \frac{\Sigma{t_{down}-t_{up}}}{Number of failures}$
* $MTBF_{MDS} = \frac{\Sigma{t_{down}-t_{up}}}{Number of failures}$
* Striped array of N disks
  * $MTBF_{striped-array} = MTBF_{drive} / N$
* Stripped+Mirrored array of N+N disks
  * $MTBF_{striped-mirrored-array} = MTBF_{pair} / N$
  * $MTBF_{pair} = (MTBF_{drive}/2) + MTBF_{drive} = 1.5 * MTBF_{drive}$
* 4N data disks and N parity disks
  * $MTBF_{p-array} = MTBF_{stripe} / N$
  * $MTBF_{stripe} = (MTBF_{stripe}/5) + (MTBF_{drive}/4) = 0.45 * MTBF_{drive}$

### Disk Rebuild and Disk Sparsing

* Goal: restore array redundancy after a failure
* After first failure, data still available for degraded access
* Second failure would result in **data loss**
* Trade-off: reliability vs. performance

### Reliability of Disk Arrays with Rebuild

* Mean Time To Rebuild (MTTR)
  * No data loss if repair completes before 2nd failure
* Mean Time To Data Loss (MTTDL)
  * Canonical way: solve Markov model of array states
* There is a huge difference between MTTR between with/without rebuild

### Three Modes of Operation

* Normal mode
  * everything working; maximum efficiency
* Degraded mode
  * some disk unavailable
  * must use degraded mode operations
* Rebuild mode
  * reconstructing lost disk's contents onto space
  * degraded mode operations plus competition with rebuild