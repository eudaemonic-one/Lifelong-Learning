# Lecture 16 API Design 2

## General Principles

* API Should Do One Thing and Do it Well
  * Functionality should be easy to explain
    * If it's hard to name thats generally a bad sign
    * Be amenable to splitting and merging modules
* API Should be as small as possible but no smaller
  * API must satisfy its requirements
    * Generalizing an API can make it smaller
  * **When in doubt, leave it out**
    * Functionality, classes, methods, parameters, etc
    * You can always add, but you can never remove
  * Conceptual weight (a.k.a. conceptual surace area)
  * e.g. List<T> sublist<T>()
* **Don't make users do anything library could do for them**
  * Reduce need for boilerplate code
* Monitor complexity constantly
* Implementation should not impact API
  * Implementation constraints may change; API won't
* API should coexist peacefully with platform
  * Do what if customary
  * Take advantage of API-friendly features
* Consider the performance consequences of API design decisions
  * **Bad API decisions can limit performance forever**
  * But do not warp API to gain performance

## Class Design

* Don't expose a new type that lacks meaningful contractual refinements on an existing supertype
  * Reduces conceptual surface area
  * Increases flexibility
* Minimize Mutability
  * Parameters should be immutable that eliminates need for defensive copying
  * Classes should be immutable unless there's a good reason to do otherwise
  * If mutable, keep state-space small, well-defined
* Minimize accessibility of everything
  * Make classes, members as private as possible
  * Public classes should have no public fields
  * Maximizes information hiding
  * Minimizes coupling
* Subclass only when an is-a relationship exists
  * Subclassing implies substitutability (Liskov)
  * Never subclass just to reuse implementation
  * Ask yourself "Is every Foo really a Bar?"
* Design & document for inheritance or else prohibit it
  * Inheritance violates encapsulation
  * If you allow subclassing, document self-use
  * Conservative policy: all concrete classes uninheritable

## Method Design

* "Fail Fast" - prevent failure, or fail quickly, predictably, and informatively
  * API should make it impossible to do what's wrong
    * Fail at complie time or sooner
  * Misuse that's statically detectable is second best
    * Fail at build time, with proper tooling
  * Misuse leading to prompt runtime failure is third best
    * Fail when first erroneous call is made
    * Method should be failure-atomic
  * Misuse that can lie undetected is what nightmares are made of
    * Fail at undetermined place and time in the future
* Use appropriate parameter and return types
  * Favor interface types over classes for input
  * **Don't use String if a better type exists**
* Use consistent parameter ordering across methods
* Avoid long parameter lists
  * Three or fewer parameters is idea
  * Long lists of identically typed params are very harmful
    * Programmers transpose parameters by mistake
    * Programs still compile and run, but misbehave
* Avoid return values that demand exceptional processing
  * Clients should not have to write extra code
* Handle boundary conditions (edge cases, corner cases) gracefully
* Do not overspecify the behavior of methods
  * Don't specify internal details
* Provide programmatic access to all data available in string form
* Overload with care
  * Avoid ambiguous overloading
  * Often better to use a different name

## Exception Design

* Throw exceptions to indicate exceptional conditions
  * Don't force client to use exceptions for control flow
  * Conversely, don't fail silently
* Favor unchecked exceptions
  * Checked - client must take recovery action
  * Unchecked - generally a programming error
  * Overuse of checked exceptions causes boilerplate
* Favor the reuse of existing exception types
  * Especially IllegalArgumentException and IllegalStateException
* Include failure-capture information in exceptions
  * e.g. IndexOutOfBoundsException should include index and ideally, bounds of access

## Documentation

* API documentation is critical
* Document religiously
  * Document every class, interface, method, constructor, parameter, and exception
  * Document thread safety
  * If class is mutable, document state space
  * If API spans packages, JavaDoc is not sufficient