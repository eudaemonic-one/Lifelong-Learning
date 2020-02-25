# Lecture 04 Specification and Unit Testing

## Specifying Program Behavior - contracts

### Contract

* Agreement between an object and its user
* Include
  * Method signature (type specification)
  * Functionality and correctness expectations
  * Performance expectation
* **What** the method does, not **how** it does it

* Method contract details
  * Defines method and caller's responsibilities
  * structure
    * Precondition
    * Postcondition
    * Exceptional behavior

### Textual specification - Javadoc

* Practical approach
* Document
  * Every parameter
  * Return value
  * Every exception (checked and unchecked)
  * What the method does

## Testing

### Semantic correctness

* type checking
* static analysis tools (e.g. SpotBugs) recognize bug patterns

### Testing

* Execute the program with selected inputs in a controlled environment
* Goals
  * Reveal bugs
  * Access quality
  * Clarify the specification, documentation
* Automate testing
  * Execute test regularly
    * After every change

#### Unit Test

* Unit tests for small units:; methods, classes, subsystems
* Typically written by developers
* Many small, fast-running, independent tests
* Few dependencies on other system parts or environment
* Insufficient, but a good starting point

#### JUnit

* Unit testing framework in Java

#### Selecting test cases

* Read specification
* Write tests for
  * Representative case
  * Invalid cases
  * Boundary conditions
* Write stress tests
  * Automatically generate huge numbers of test cases
* Aim to cover the specification

#### Testable code

* Think about testing when writting code
* Test-Driven Development (TDD)
  * Write tests before you write the code
  * Write tests can expose API weakness

#### Run test frequently

* You should only commit code that is passing all tests
* If test suite becomes too large
  * local package-level test (smoke test)
  * run all tests nightly

## Overriding objects methods

* The relevant methods are all present on all Objects
  * equals(Object o)
  * hashCode
  * toString