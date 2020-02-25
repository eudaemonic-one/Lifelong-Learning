# Lecture 7 Design Patterns for Reuse

## Iterator Pattern

* Problem: Clients need uniform strategy to access all elements in a container, independent of the container type
* Solution: A strategy pattern for iteration
* Consequences:
  * Hides internal implementation of underlying container
  * Easy to change container type
  * Facilitates communication between parts of the program
* The default Collections implementations are mutable, but the Iterator implementations assume the collection does not change while the Iterator is being used

## Decorator Pattern

* Problems: You need arbitrary or dynamically composable extensions to individual objects
* Solution: Implement a common interface as the object you are extending, add functionality, but delegate primary responsibility to an underlying object
* Consequences:
  * More flexible than static inheritance
  * Customizable, cohesive extensions
  * Breaks object identity, self-references

## Design goals and design principles

* Design principles: heuristic to achieve design goals
  * Low coupling
    * Enhances understandability
    * Reduce cost of change
    * Eases reuse
  * Low representational gap
  * High cohesion
    * Each component should  have a small set of closely-related responsibilities
    * Facilitates reuse
    * Eases maintenance