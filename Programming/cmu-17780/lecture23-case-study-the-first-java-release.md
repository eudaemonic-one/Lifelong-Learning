# Lecture 23 Case Study \- JDK 1.02 Libraries (The First Java Release)

## The Successful Part

### Java Platform

* Safe language / managed runtime
  * No segment faults, memory corruption bugs, etc
* Tightly specified primitive types, expression evaluation order, etc
  * Greatly facilitates program portability
  * A natural accompaniment to a managed runtime
* Dynamic linking
  * In bad old days, changing a library required recompiling all clients
* Superficial similarity to C/C++
  * Appealed immediately to lots of C and C++ programmers

### Type System

* Object Oriented Language
  * Encapsulation is necessary to prove components correct in isolation
  * Class inheriatance was, if nothing else, a marketing necessity
* Multiple interface inheritance
  * Enables easy pluggability of components
  * Avoids the pain of multiple implementation inheritance
* Static typing
  * Detects bugs at compile time, increasing program reliability
  * Enables IDE to help programmer write correct code quickly (auto-complete)

### Features

* Threads
  * Twilight of uniprocessor era
  * Concurrency was increasingly important
* Garbage Collection
  * Eliminates pain and bugs that go with manual memory management
* Exceptions
  * Error code are error prone
* Unicode
  * It was the twilight of the ASCII era

## What You Leave Out Can Be as Important as What You Put In

* Lexical macros
  * Made all Java programmers look similar
  * Enabled programmer protability
  * Enabled world-class toolability
* Multiple implementation inheritance
* Untrammeled operator overloading
* **Java omitted support for header files** (external declarations)
* **Javadoc largely eliminated external API documentation**
* In combination, these decisions meant
  * Declaration, documentation, and implementation collocated
  * Good API documentation became part of culture from start

## The Bad and The Ugly Part

### Expression Evaluation

* Silent widening conversions from `int` to `float` and `long` to `double` are lossy
* Compound assignment operators (e.g., `+=`, `-=`) can cause a silent narrowing cast
* Operators `==` and `!=` do reference comparisons even if `equals` is overridden
* Constant variables are inclined where they are used

### Constructors

* Default constructor should not exist
  * And they certainly shouldn't be `public`
  * Lead to unintentional instantiability and sloppy API documentation
* Invoking overridden method from constructor should be illegal
  * It has no valid uses, and leads to subtle bugs

### Concurrency

* **All** objects have a lock associated with them
* The lock associated with an object is publicly visible
  * `synchronized(Thread.class)`
* All locks have exactly one associated condition variable
  * You can't associate multiple "wait-sets" with a single locks
  * Results in unnecessary context switches

### Miscellaneous

* Signed `byte` type
* Lack of unsigned `int` and `long` types
* Switch statement is not structured
* Arrays should have overridden `toString` to be informative
* Exceptions obliterate pending exceptions
* Guaranteed `String` constant interning

## Conclusion

* The good parts: key design decisions
* The bad and the ugly: largely confined to details

