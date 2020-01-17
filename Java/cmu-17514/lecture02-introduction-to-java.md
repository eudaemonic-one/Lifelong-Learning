# Lecture 02 Introduction to Java

## "Hello World" explained

* must use a class even if not doing OO programming
* main must be public
* main must be static
* main must return void
* main must declare command line arguments even if unused
* `psvm` = `public static void main`

## Execution

* Compile
  * `javac HelloWorld.java`
* Launch
  * `java HelloWorld`

## Type System

* Primitives
  * int 32-bit signed integer
  * long 64-bits signed integer
  * byte 8-bits signed integer
  * short 16-bit signed integer
  * char 16-bit signed integer
  * float 32-bit IEEE 754 floating point number
  * double 64-bit IEEE 754 floating point number
  * Boolean Boolean value: true or false

* Object Reference Types
  * Have identity distinct from value
  * Some mutable, some immutable
  * On heap, garbase collected
  * Unity of expression with generics
  * More costly

#### Tips

* Minimize scope of local variable

* Initialize variables in declaration

* Prefer for-each loops to regular for-loops

* Use common idoms

* Watch out bas smell of code

### Objects

* a bundle of state and behavior
* State - fields of the object
* Behavior - methods

### Class

* Every object has a class
* Class defines both type and implementation
* Methods of a classs are its **Application Programming Interface (API)**

### The class hierarchy

* The root is Object
* All classes except Object have one parent class
* A class is an instance of all its superclasses

### Implementation inheritance

* A class
  * Inherits visible fields and methods from its superclasses
  * Can override methods to change their behavior
* Overriding method implementation must obey contracts of its superclasses
  * **Liskov Substitution Principle (LSP)**

### Interface types

* Defines a type without an implementation
* Much more flexible than class types

### Enum types

* Java has object-oriented enums

### Boxed primitives

* Immutable containers for primitive types
* Boolean, Integer, Short, Long, Character, Float, Double
* **Canonical use case is collections**
* **Do not use boxed primitives unless you have to**

### Comparing values

* x == y compares the contents of x and y
  * primitive values: returns true if x and y have the same value
  * objects refs: returns true if x and y **refer to** same object

* x.equals(y) compares the values of the objects referred to by x and y
  * **Always use .equals to compare object refs**
  * Except for enums, which are special

## Quick 'n' dirty I/O

* Output
  * Unformatted
  * Formatted

## A brief introduction to collections

Primary collection interfaces

* Collection<-(Set, List, Queue<-(Deque))

* Map

Primary collection implementations

* Set - HashSet
* List - ArrayList
* Queue - ArrayDeque
* Deque - ArrayDeque
* Stack - ArrayDeque
* Map - HashMap

Other noteworthy collection implementations

* Set - LinkedHashSet/TreeSet/EnumSet
* Queue - PriorityQueue
* Map - LinkedHashMap/TreeMap/EnumMap