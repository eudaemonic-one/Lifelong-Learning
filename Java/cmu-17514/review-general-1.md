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

## Design for Reuse: Delegation and Inheritance

### Miscs

| Primitive             | Latency: ns | Latency: us | Latency: ms |
| --------------------- | ----------- | ----------- | ----------- |
| L1 cache reference    | 0.5         |             |             |
| L2 cache reference    | 7           |             |             |
| Main memory reference | 100         |             |             |
| Send packet           | 150,000,000 | 150,000     | 150         |

### Behavioral Subtyping

* **Liskov Substitution Principle**
  * **Let q(x) be a property provable about objects x of type T. Then q(y) should be provable for objects y of type S where S is a subtype of T.**
* Compiler-enforced rules in Java
  * Subtypes can add, but not remove methods
  * Concrete class must implement all undefined methods
  * Overriding method must return same type or subtype
  * Overriding method must accept the same parameter types
  * Overriding method may not throw additional exceptions
* **Subtypes must have**:
  * **Same or stronger invariants**
  * **Same or stronger postconditions for all methods**
  * **Same or weaker preconditions for all methods**

### Delegation

* Delegation is simply when one object relies on another object for some subset of its functionality

* Judicious delegation enables code reuse

* Delegation and design

  * Small interfaces with clear contracts
  * Classes to encapsulate algorithms, behaviors

* **Subtype Polymorphism**

  * Different kinds of objects can be treated uniformly by client code
  * Each object behaves according to its type

* Benefits of inheritance

  * Reuse of code
  * Modeling flexibility

* **Inherance and subtyping**

  * **Inheritance is for polymorphism and code reuse**

    * Write code once and only once

    * Superclass features implicitly available in

      subclass

  * **Subtyping is for polymorphism**

    * Accessing objects the same way, but getting different behavior
    * Subtype is substitutable for supertype

* **Interfaces and Classes**

  * An interface defines expectations / commitments for clients
  * A class fulfills the expectations of an interface

* **Delegation vs. Inheritance**

  * Inheritance can improve modeling flexibility
  * Favor composition/delegation over inheritance\
    * Inheritance violates information hiding
    * Delegation supports information hiding
  * Design and document for inheritance, or prohibit it
    * Document requirements for overriding any method

## Introduction to Design Patterns Part 1: Designing Classes

### Dynamic method dispatch

* (Compile time) Determine which class to look in
* (Compile time) Determine method signature to be executed
* (Runtime) Determine dynamic class of the receiver
* (Runtime) From dynamic class, determine method to invoke

### UML

* Interfaces vs. Classes
* Fields vs. Methods
* Relationships:
  * **Extends** (inheritance) - triangle (connects superclasses) with solid line
  * **Implement** (Realization) - triangle (connects interface) with dot line
  * **Has a** (aggregation) - diamond (connects host) with solid line
  * **non-specific association** - solid line
* Visibility: **+** (public) **-** (private) **#** (protected)

### Strategy Pattern

* Problem: Clients need different variants of an algorithm
* Solution: Create an interface for the algorithm, with an implementing class for each variant of the algorithm
* Consequences:
  * Easily extensible for new algorithm implementations
  * Separates algorithm from client context
  * Introduces an extra interface and many classes

### Command Pattern

* Problem: Clients need to execute some (possibly flexible) operation without knowing the details of the operation
* Solution: Create an interface for the operation, with a class (or classes) that actually executes the operation
* Consequences:
  * Separates operation from client context
  * Can specify, queue, and execute commands at different times
  * Introduces an extra interface and classes

### Template Method Pattern

* Problem: An algorithm consists of customizable parts and invariant parts
* Solution: Implement the invariant parts of the algorithm in an abstract class, with abstract (unimplemented) primitive operations representing the customizable parts of the algorithms. Subclasses customize the primitive operations
* Consequences:
  * Code reuse for the invariant parts of algorithm
  * Customization is restricted to the primitive operations
  * Inverted (Hollywood-style) control for customization
* Template method uses inheritance to vary part of an algorithm
  * Template method implemented in supertype, primitive operations implemented in subtypes
* Strategy pattern uses delegation to vary the entire algorithm
  * Strategy objects are reusable across multiple classes
  * Multiple strategy objects are possible per class

## Design Patterns for Reuse Part 2

### Iterator Pattern

* Problem: Clients need uniform strategy to access all elements in a container, independent of the container type
* Solution: A strategy pattern for iteration
* Consequences:
  * Hides internal implementation of underlying container
  * Easy to change container type
  * Facilitates communication between parts of the program

### Decorator Pattern

* Problem: You need arbitrary or dynamically composable extensions to individual objects
* Solutions: Implement a common interface as the object you are extending, add functionality, but delegate primary responsibility to an underlying object
* Consequences:
  * More flexible than static inheritance
  * Customizable, cohesive extension
  * Breaks object identity, self references
* Composition and forwarding

### Design Principles: Heuristics to achieve design goals

* Low coupling
  * each component should depend on as few other components as possible
  * to increase understandability, reuse
* Low representational gap
  * to increase understandability, maintainability
* High cohesion
  * each component should have a small set of closely-related responsibilities
  * increase understandability

## A Formal Design Process: Domain Modeling

### Visualizing Dynamic Behavior: Interaction Diagrams

* An *interaction diagram* is a picture that shows, for a single scenario of use, the events that occur across the system’s boundary or between subsystems
* Clarifies interactions:
  * Between the program and its environment
  * Between major parts of the program
* Sequence Diagram

### Artifacts of Design Process

* Model / diagram the problem, define objects
  * Domain model (aka. conceptual model)
* Define system behaviors
  * System sequence diagram
    * A *system sequence diagram* is a model that shows, for one scenario of use, the sequence of events that occur on the system’s boundary
    * Design goal: Identify and define the interface of the system
    * Input: Domain description and one use case
    * Output: A sequence diagram of system-level operations
* Assign object responsibilities, define interactions
  * Object interaction diagrams
* Model / diagram a potential solution
  * Object model

### Modeling a Problem Domain

* Identify key concepts of the domain description
  * Identify nouns, verbs, and relationships between concepts
  * Distinguish operations and concepts
  * Brainstorm with a domain expert
* Visualize as a UML class diagram, a *domain model*
  * Show class and attribute concepts
  * No operations/methods