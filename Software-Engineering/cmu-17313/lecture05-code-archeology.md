# Lecture 5 Code Archeology

## Understand Software Systems

* You cannot understand the entire system
* Develop and test a working model or set of working hypotheses about how (some part of) a system works
  * Working model: an understanding of the pieces of the system (components), and the way they interact (connections)
  * It is common in practice to consult documentation, experts
  * Prior knowledge/experience is also useful
* Software constantly changes -> Software is easy to change
* Software is a big redundant mess -> there's always something to copy as a starting point

## Cognitive Biases

* anchoring
* confirmation bias
* congruence bias: the tendency to test hypotheses exclusively through direct testing, instead of testing possible alternative hypotheses

### Information Gathering

* Basic needs:
  * Code/file search and navigation
  * Code editing
  * Execution of code, tests
  * Observation of output
* Static information gathering: use tools to help manage complexity
* Consider documentation and tutorials judiciously
* Dynamic information gathering: high-level principles
  * Change is a useful primitive to inform mental models about a software system
  * Systems almost provide some kind of starting point
  * Put simply
    * Build it
    * Run it
    * Change it
    * Run it again
* Sanity check confirms that you can build and run the code you built