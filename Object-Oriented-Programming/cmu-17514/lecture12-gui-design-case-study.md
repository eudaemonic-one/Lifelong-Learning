# Lecture 12 Design Case Study: GUI

## The Composite Pattern

* Problem: Collection of objects has behavior similar to the individual objects
* Solution: Have collection of objects and individual objects implement the same interfaces
* Consequences:
  * Client code can treat collection as if it were an individual object
  * Easier to add new object types
  * Design might become too general, interface insufficiently useful

## The Chain of Responsibility Pattern

* Problem: You need to associate functionality within a deep nested or iterative structure, possibly with multiple objects
* Solution: Request for functionality, pass request along chain until some component handles it
* Consequences:
  * Decouples sender from receiver of request
  * Can simplify request-handling by handing requests near root of hierarchy
  * Handling of request not guaranteed