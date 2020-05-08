# Lecture 07 Machine-Level Programming III: Procedures

## Mechanisms in Procedures

* Passing control
  * To beginning of procedure code
  * Back to return point
* Passing data
  * Procedure arguments
  * Return value
* Memory management
  * Allocate during procedure execution
  * Deallocate upon return
* Mechanisms all implemented with machine instructions

## Stack Structure

* Region of memory managed with stack discipline
* Grows toward lower addresses
* **Stack pointer**: %rsp - address of top element

## Calling Conventions

### Passing control

* Use stack to support procedure call and return
* **Procedure call**: call label
  * Push return address on stack
  * Jump to label
* Return address
  * Address of the next instruction right after call
* **Procedure return**: ret
  * Pop address from stack
  * Jump to address

### Passing data

* Registers
  * First 6 arguments
    * %rdi
    * %rsi
    * %rdx
    * %rcx
    * %r8
    * %r9
  * Return value
    * %rax
  * Stack
    * Only allocate stack space when needed

### Manage local data

* Stack-based languages - Languages that support recursion
* Stack discipline
  * State for given procedure needed for limited time
  * Callee returns before caller does
* Stack allocated in **Frames**
  * Return information
  * Local storage
  * Temporary storage

### Stack Frame

* Current Stack Frame
  * “Argument build:”
    * Parameters for function about to call
  * Local variables
    * If can’t keep in registers
  * Saved register context
  * Old frame pointer (optional)
* Caller Stack Frame
  * Return address
    * Pushed by call instruction
  * Arguments for this call
* Registers Usage
  * %rax
    * return value
    * Also caller-saved
    * Can be modified by procedure
  * %rdi, ..., %r9
    * Arguments
    * Also caller-saved
    * Can be modified by procedure
  * %r10, %r11
    * Caller-saved
    * Can be modified by procedure
  * %rbx, %r12, %r13, %14
    * Callee-saved
    * Callee must save & restore
  * %rbp
    * Callee-saved
    * Callee must save & restore
    * May be used as frame pointer
    * Can mix & match
  * %rsp
    * Special form of callee save
    * Restored to original value upon exit from procedure