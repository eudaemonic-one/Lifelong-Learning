# Learn Go Programming

## Overview

### Features of Go Programming

* Support for environment adopting patterns similar to dynamic languages. For example, type inference (x := 0 is valid declaration of a variable x of type int)
* Compilation time is fast.
* Inbuilt concurrency support: lightweight processes (via go routines), channels, select statement.
* Go programs are simple, concise, and safe.
* Support for Interfaces and Type embedding.
* Production of statically linked native binaries without external dependencies.

### Features Excluded Intentionally

* Support for type inheritance
* Support for method or operator overloading
* Support for circular dependencies among packages
* Support for pointer arithmetic
* Support for assertions
* Support for generic programming

## Program Structure

* Package Declaration
* Import Packages
* Functions
* Variables
* Statements and Expressions
* Comments

## Basic Syntax

### Tokens in Go

A token is either a keyword, an identifier, a constant, a string literal, or a symbol.

### Line Separator

The Go compiler internally places “;” as the statement terminator to indicate the end of one logical entity.

### Comments

Comments start with /* and terminates with the characters */

### Identifiers

identifier = letter { letter | unicode_digit }

## Data Types

### Integer Types

| uint8 | uint16 |  uint32 | uint64 | int8 | int16 | int32 | int64 |

### Floating Types

| float32 | float64 | complex64 (Complex numbers with float32 real and imaginary parts) | complex128 |

### Other Numeric Types

| byte | rune (same as int32) | unit | int | uintptr |

## Variables

Each variable in Go has a specific type, which determines the size and layout of the variable's memory, the range of values that can be stored within that memory, and the set of operations that can be applied to the variable.

Upper and lowercase letters are distinct because Go is case-sensitive.

### Variable Definition in Go

```go
var variable_list optional_data_type;
var i, j, k int;
i = 1, j = 2;
```

Variables with static storage duration are implicitly initialized with nil (all bytes have the value 0).

### Static Type Declaration in Go

A static type variable declaration provides assurance to the compiler that there is one variable available with the given type and name so that the compiler can proceed for further compilation without requiring the complete detail of the variable. A variable declaration has its meaning at the time of compilation only, the compiler needs the actual variable declaration at the time of linking of the program.

### Dynamic Type Declaration / Type Inference in Go

A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it. The compiler does not require a variable to have type statically as a necessary requirement.

```go
y := 42
```

### Mixed Variable Declaration in Go

```go
var a, b, c = 3, 5, "foo"
```

### The lvalues and the rvalues in Go

An *lvalue* may appear as either the left-hand or right-hand side of an assignment. An *rvalue* is an expression that cannot have a value assigned to it which means an rvalue may appear on the right- but not left-hand side of an assignment.
