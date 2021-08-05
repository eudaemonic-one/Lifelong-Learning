# Lecture 06 Design Patterns
## UML class diagrams

### UML class diagrams

* Interfaces vs. classes
* Fields vs. methods
* Relationaships
  * extends (inheritance)
  * implements (realization)
  * has a (addregation)
  * non-specific association
* Visibility
  * \+ (public)
  * \- (Private)
  * \# (protected)

### UML interaction diagrams

## Introduction to design patterns

### Strategy pattern

* Problem: Clients need different variants of an algorithm
* Solution: Create an interface for the algorithm, with an  implementing class for each variant of the algorithm
* Consequences:
  * Easily extensible for new algorithm implementation
  * Separates algorithm from client context
* Share almost the same struction with command pattern

### Command pattern

* Problem: Clients needs to execute some operation without knowing the details of the operation
* Solution: Create an interface for the operation, with a class that actually executes the operation
* Consequences
  * Separates operation from client context
  * Can specify, queue, and execute commands at different times

### Template Method pattern

* Problem: An algorithm consists of customizable parts and invariant parts
* Solution: Implement the invariant parts in an abstract class, with abstract primitive operations representing the customizable parts of the algorithm. Subclasses customize the primitive operations
* Consequnces
  * Code reuse for the invariant parts
  * Customization is restricted to the primitive operations