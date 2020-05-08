# Lecture 01 Overview

## Big Picture

### Systems Knowledge

* How hardware (processors, memories, disk drives, network infrastructure) plus software (operating systems, compilers, libraries, network protocols) combine to support the execution of application programs
* How you as a programmer can best use these resources

### Useful outcomes from taking 213/513/613

* Become more effective programmers
  * Able to find and eliminate bugs efficiently
  * Able to understand and tune for program performance

## Great Realities

### Ints are not Integers, Floats are not Reals

* Does not generate random values
* Cannot assume all “usual” mathematical properties
* Observation

### You’ve Got to Know Assembly

* Chances are, you’ll never write programs in assembly
* But: Understanding assembly is key to machine-level execution model

### Memory Matters

Random Access Memory Is an Unphysical Abstraction

* Memory is not unbounded
* Memory referencing bugs especially pernicious
* Memory performance is not uniform

### There’s more to performance than asymptotic complexity

* Constant factors matter too!
* And even exact op count does not predict performance
* Must understand system to optimize performance

### Computers do more than execute programs

* They need to get data in and out
* hey communicate with each other over networks
