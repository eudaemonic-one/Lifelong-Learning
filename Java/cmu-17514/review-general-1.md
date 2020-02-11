Review (General) Phase I

## Introduction

Software engineering entails making **decisions under constraints** of limited time, knowledge, and resources.

| Metrics of Software quality | Design Goals                                                 |
| --------------------------- | ------------------------------------------------------------ |
| Functional Correctness      | Adherence of implementation to the specifications            |
| Robustness                  | Ability to handle anomalous events                           |
| Flexibility                 | Ability to accommodate changes in specifications             |
| Reusability                 | Ability to be reused in another application                  |
| Efficiency                  | Statisfaction of speed and storage requirement               |
| Scalability                 | Ability to serve as the basis of a larger version of the application |
| Security                    | Level of consideration of application security               |

* **Design Goals** enable evaluation of designs
* **Design Principles** are heuristics that describe best practice
* **Design Patterns** codify repeated experiences, common solutions

## Introduction to Java

### HelloWorld

```java
class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello world!");
    }
}
```

* you must use a class even if you aren’t doing OO programming
* public static void main (**PSVM**)
* standard I/O requires use of static field of System

### Execution of Java Program

* First compile source file by java HelloWorld.java which produces class file HelloWorld.class
* Then launch the program by java HelloWorld which will execute main method on Java Virtual Machine(JVM)
  * Managed runtime enables features such as safe, flexible, garbage collection

### Type System

| Primitives                                           | Object Reference Types                          |
| ---------------------------------------------------- | ----------------------------------------------- |
| int, long, byte, short, char, float, double, boolean | Classes, interfaces, arrays, enums, annotations |
| No identity except their value                       | Have identity distinct from value               |
| Immutable                                            | Some mutable, some immutable                    |
| On stack, exist only when in use                     | On heap, garbage collected                      |
| Can’t achieve unity of expression                    | Unity of expression with generics               |
| Dirt cheap                                           | More costly                                     |

### Objects and Classes

* An **object** is a bundle of state and behavior
* Every object has a class
  * A class defines methods and fields (collectively known as **Members**)
* Class defines both type and implementation
* Loosely speaking, the methods of a class are its **Application Programming Interface (API)**

### Inheritance

* The root of class hierarchy is Object (**all non-primitives are objects**)
* A class is an instance of all its superclasses
* A class
  * Inherits visible fields and methods from its superclasses
  * Can override methods to change their behavior
* Overriding method implementation must obey contract(s) of its superclass(es)
  * Ensures subclass can be used anywhere superclass can
  * **Liskov Substitution Principle (LSP)**
    * If S is a subtype of T, an object of type T may be substituted with any object of a subtype S without altering any of the desirable properties of the program (correctness, task performed, etc.)

### Interface

* Defines a type without an implementation
* Much more flexible than class types

### Miscs

* **Minimize scope of local variables** [EJ Item 57]
* Initialize variables in declaration
* Prefer for-each loops to regular for-loops
* **Always use .equals to compare object references**
* **Comparing values**
  * x == y compares the **contents** of x and y
    * **primitive values**: returns true if x and y **have the same value**
    * **objects references**: returns true if x and y **refer to same object**
  * x.equals(y) compares the **values of the objects referred to** by x and y

### Collections in Java

| Interface | Implementation                        |
| --------- | ------------------------------------- |
| Set       | HashSet/LinkedHashSet/TreeSet/EnumSet |
| List      | ArrayList                             |
| Queue     | ArrayDeque/PriorityQueue              |
| Deque     | ArrayDeque                            |
| Stack     | ArrayDeque                            |
| Map       | HashMap/LinkedHashMap/TreeMap/EnumMap |

## Object-Oriented Programming in Java

### Objects / Classes / Interfaces/ Implementations

* An **object** is a bundle of **state ** (fields) and **behavior** (methods), which are collectively known as **memebers**
* Class defines both **type** and **implementation**
* Loosely speaking, the methods of a class are its **Application Programming Interface (API)**
* Multiple implementations of API can coexist
  * **Multiple classes can implement the same API with different performance and behavior**
* In Java, an API is specified by *interface* or *class*
  * Interface provides only an API
  * Class provides an API and an implementation
  * A class can implement multiple interfaces
* Prefer interfaces to classes as types

### Information Hiding

* Information Hiding or Encapsulation
  * Cleanly separates API from implementation
  * Modules communicate *only* through APIs
* Benefits of Information Hiding
  * **Decouples** the classes that comprise a system
  * **Speeds up system development**
  * **Eases burden of maintenance**
  * **Enables effective performance tuning**
  * **Increases software reuse**
* Information hiding with interfaces
  * Declare variables using interface types
  * Client can use only interface methods
  * Fields not accessible from client code
* **Visibility modifiers for members**
  * private – Accessible *only* from declaring class
  * package-private – Accessible from any class in the package where it is declared
  * protected – Accessible from subclasses of declaring class (and within package)
  * public – Accessible from anywhere

### Exceptions

* Inform caller of problem by transfer of control
* Semantics
  * Propagates up stack until main method is reached (terminates program), or exception is caught
* Sources
  * Program can throw explicitly
  * Underlying virtual machine (JVM) can generate
* **Checked vs. Unchecked Exceptions**
  * **Checked exception**
    * Must be caught or propagated, or program won’t compile
    * Exceptional condition that programmer must deal with
  * **Unchecked exception**
    * No action is required for program to compile
    * Usually indicates a programming error
  * **Error**
    * Special unchecked exception thrown by JVM
    * Recovery is impossible
* Design choice
  * try-with-resources that automatically closes resources

### References

* https://www.geeksforgeeks.org/classes-objects-java/
* https://www.geeksforgeeks.org/inheritance-in-java/
* https://www.geeksforgeeks.org/interfaces-in-java/
* https://www.geeksforgeeks.org/difference-between-abstract-class-and-interface-in-java/
* https://www.geeksforgeeks.org/access-modifiers-java/
* https://www.geeksforgeeks.org/checked-vs-unchecked-exceptions-in-java/

## Testing and Object Methods in Java

### Contracts

* Agreement between an object and its user
* Includes
  * Method signature (type specifications)
  * Functionality and correctness expectations
  * Performance expectations
* **What** the method does, not **how** it does it
  * **Interface** (API), not **implementation**
* Method contract structure
  * Preconditions
  * Postconditions
  * Exceptional behavior
* Formal contract specification - **Java Modelling Language (JML)**
* Textual specification - **Javadoc**

### Testing Correctness - Junit and Friends

* Semantic correctness
  * Compiler ensures types are correct
  * Static analysis tools recognize many common problems
* Formal verification (such as mathematical proof)
* Testing
  * Executing the program with selected inputs in a controlled environment
  * Reveal bugs and assess quality
  * **Executes tests regularly - After every change**
* Unit tests
  * Unit tests for small units: methods, classes, subsystems
  * JUnit - Popular unit-testing framework for Java
* Write tests based on the specification
  * Representative cases
  * Invalid cases
  * Boundary conditions
* Stress tests - automatically generate huge numbers of test cases
* **Think about testing when writing code**
* Test-Driven Development (TDD)
  * write tests before you write the code

### Overriding Object Methods

* **equals** - returns true if the two objects are equal
  * Reflexive
  * Symmetric
  * Transitive
  * Consistent
  * Non-null
* **hashCode** - returns an int that must be equal for equal objects, and is likely to differ on unequal objects
  * Equal objects must have equal hash codes
  * If override equals must override hashCode
  * Unequal objects should have different hash codes
  * Hash code must not change unless object mutated
* **toString** - returns a printable string representation
* Always override toString
* No need to override equals and hashCode if you want identity semantics
* For primitives must use ==
* For object reference types, the == operator provides *identity semantics*, exactly as implemented by Object.equals