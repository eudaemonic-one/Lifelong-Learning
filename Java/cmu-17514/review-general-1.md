# Review (General) Phase I

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