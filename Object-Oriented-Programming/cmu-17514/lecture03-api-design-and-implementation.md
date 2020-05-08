# Lecture 03 API Design and Implementation

## Object-Oriented Programming basics

### Object

* An **object** is a bundle of state and behavior
* State - **fields**
* Behavior - **methods**

### Classes

* Every object has a class
* Class defines both **type and implementation**
* the methods of a class are its **Application Programming Interface (API)**
  * Defines how users interact with instances

### Interfaces and implementations

* Multiple implementations of API can coexist
* In Java, an API is specified by *interface* or *class*
  * Interface - only API
  * Class - API and implementation
  * An interface defines but does not implement API
  * Interface decouples client from implementations
* Why multiple implementations
  * Often **performance and behavior** both vary
  * Provides a **functionality-perfomance tradeoff**

## Information hiding

* Hides internal data and implementation details from other modules
* Well-designed code hides all implementation details

### Benifits of information hiding

* Decouples the classes that comprise a system
  * Allows them to be developed, tested, optimized, used, understood, and modified in isolation
* Speeds up system development
  * Classes can be developed in parallel
* Eases burden of maintenance
* Enables effective performance tuning
  * Hot classes can be optimized in isolation
* Increases software reuse
  * Loosely-coupled classes often prove useful in other context

### Information hiding with interfaces

* Declare variables using interface types
* Client can use only interface methods
* Fields not accessible from client code

### Mandatory Information Hiding

* private - Accessible only from declaring class
* package-private - Accessible from any class in same package
* protected - Accessible from subclasses of declaring class (and within package)
* public - Accessible from anywhere

## Exceptions

* Inform caller of problem by transfer of control
* Semantics
  * Propagate up stack until main method is reached (terminate program), or exception is caught
* Sources
  * Program can throw explicitly
  * Underlying virtual machine can generate

### Benefits of Exceptions

* You can't forget to handle common failure modes
  * Compare: using a flag or special return value
* Provide high-level summary of error, and stack trace
  * Compare: core dump
* Improve code structure
* Ease task of writing robust code

### Checked vs. Unchecked Exceptions

* Checked exception
  * Must be caught or propagated, or program won't compile
  * Exceptional condition that programmer must deal with
* Unchecked exception
  * No action is required for program to compile
  * Usuallt\y indicates a programming error
* Error
  * Special unchecked exception thrown by JVM
  * Recovery is impossible

* Java's exception hierarchy
  * object -> Throwable -> (Exception, Error)
* Design choice: checked exceptions, unchecked exceptions, and error codes
  * Unchecked exception - Programming error, other unrecoverable failure
  * Checked exception - An error that every caller should be aware of and handle
  * Special value
    * Common but atypical result
    * Ex: null from Map.get
  * Do not return null for zero length result
  * Guidelines for using exceptions - refer to Effective Java
  * Try with resources - Automatically closes resource