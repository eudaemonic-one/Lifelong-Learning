# Lecture 09 Machine-Level Programming V: Advanced Topics

## Memory Layout

* Shared Libraries
* Stack
  * Runtime stack (8MB limit)
  * local variables
    * Functions store local data on in stack frame
    * Recursive functions cause deep nesting of frames
* Heap
  * Dynamically allocated as needed
* Data
  * Statically allocated data
* Text
  * Executable machine instructions
  * Read-only

## Buffer Overflow

* Buffer Overflow - when exceeding the memory size allocated for an array
* Overwrite normal return address A with address of some other code S
* When Q executes rest, will jump to other code
* Avoid overflow vulnerabilities
  * `fgets` instead of `gets`
  * `strncpy` instead of `strcpy`
* System-Level Protections can help
  * Randomized stack offsets
  * Nonexecutable code segments
* Stack canaries can help
  * place special value (canary) on stack just beyond buffer
  * check for corruption before exiting function
  * `gcc -fstack-protector`
* Return-Oriented Programming Attacks
  * Use existing code e.g. library code from `stdlib`
  * String together fragments to achieve overall desired outcome
  * Does not overcome stack canaries
  * Construct program from gadgets

## Compound Types (Structure and Union)

* Structure
  * Contiguously-allocated region of memory
  * Refer to members within structure by names
  * Members may be of different types
* Union
  * Allocate according to largest element
  * Can only use one field at a time