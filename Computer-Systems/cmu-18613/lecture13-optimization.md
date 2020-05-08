# Lecture 13 Code Optimization

## Performance Realities

* There is more to performance than asymptotic complexity
* constant factors matter too

## Code Motion

* Reduce frequency with which computation performed
  * if it will always produce same result
  * especially moving code out of loop

## Reduction in Strength

* Replace costly operation with simpler one
* Shift, add instead of multiply or divide

## Share Common Subexpressions

* Reuse portions of expressions (implemented in GCC -O1)

## Optimization Blockers

* Operate under fundamental constraint
  * Must not cause any change in program behavior
* Behavior obvious to the programmer is not obvious to compiler
* Most analysis is only within a procedure
  * Whole-program analysis is usually too expensive
* Most analysis is based only on static information
* When in doubt, the compiler must be conservative

### Memory Aliasing

* There might be two different memory references specify single location
* Introduce local variables -> telling compiler not to check for aliasing

## Exploiting Instruction-Level Parrallelism

* Cycles Per Element (CPE)
  * Cycles per operation
  * T = CPE * n + Overhead
* Superscalar Processor
  * can issue multiple instructions in one cycle
  * instruction level parallelism
* Pipelined function units
  * Multiple instructions can execute in parallel
  * Some instructions take > 1 cycle but can be pipelined