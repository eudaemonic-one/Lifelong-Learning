# Lecture 08 Machine-Level Programming IV: Data

## Switch Statement

* Jump table structure
  * `goto *JTab[x]`
* Jumping
  * Direct `jmp .L8`
  * Indirect `jmp *.L4(,%rdi,8)`

## Array Allocation

* Basic Principle
  * `T A[L]`
  * Array of data type `T` and length `L`
  * Contiguously allocated region of `L * sizeof(T)` bytes in memory
  * Identifier `A` can be used as a pointer to array element 0: Type `T`
* Multidimensional Arrays
  * `T A[R][C]`
  * `R` rows `C` columns
  * Row-Major Ordering
