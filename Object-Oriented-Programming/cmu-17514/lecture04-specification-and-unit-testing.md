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
    * Precondition: what method requires for correct operation
    * Postcondition: what method establishes on completion
    * Exceptional behavior: what it does if precondition violated
* Defines correctness of implementation

### Theoretical Approach - JML

```java
/*@ requires len >= 0 && array != null && array.length == len; @
	@ ensures \result ==
	@ (\sum int j; 0 <= j && j < len; array[j]);
	@*/
int total(int array[], int len);
```

* Advantages
  * Runtime checks generated automatically
  * Basis for formal verification
  * Automatic analysis tools
* Disadvantages
  * Requires a lot of work
  * Impractical in the large
  * Some aspects of behavior not amenable to formal specification

### Textual Specification - Javadoc

* Practical approach
* Document
  * Every parameter
  * Return value
  * Every exception (checked and unchecked)
  * What the method does
* Do **not** document implementation details

## Testing

### Semantic Correctness

* Compiler ensures types are correct (type-checking)
* Static analysis tools (e.g., SpotBugs) recognize many common problems (bug patterns)
  * e.g., Overriding `equals` without overriding `hashCode`

### Formal Verification

* Use mathematical methods to prove correctness with respect to the formal specification
* Formally prove that all possible executions of an implementation fulfill the specification
* Manual effort; partial automation; not automatically decidable

### Testing

* Execute the program with selected inputs in a controlled environment
* Goals
  * Reveal bugs, so they can be fixed (main goal)
  * Access quality
  * Clarify the specification, documentation
* Automate testing
  * **Execute test regularly**
    * After every change

#### Unit Test

* Unit tests for small units:; methods, classes, subsystems
* Typically written by developers
* Many small, fast-running, independent tests
* Few dependencies on other system parts or environment
* Insufficient, but a good starting point

#### Selecting Test Cases

* Read specification
* Write tests for
  * **Representative case**
  * **Invalid cases**
  * **Boundary conditions**
* Write stress tests
  * Automatically generate huge numbers of test cases
* Aim to cover the specification

#### Testable Code

* **Think about testing when writting code**
* Test-Driven Development (TDD)
  * Write tests before you write the code
  * Write tests can expose API weakness

#### Run Test Frequently

* You should only commit code that is passing all tests
* If test suite becomes too large
  * local package-level test (smoke test)
  * run all tests nightly

## Overriding Objects Methods

### Methods Common to All Objects

* The relevant methods are all present on all objects
  * `equals` \- returns true if the two objects are “equal”
  * `hashCode`- returns an int that must be equal for equal objects, and is likely to differ on unequal objects
  * `toString` \- returns a printable string representation

### Object Implementations

* Provide identity semantics
  * `equals(Object o)` - returns true if o refers to this object
  * `hashCode()` - returns a near-random int that never changes over the object lifetime
  * `toString()` - returns a nasty looking string consisting of the type and hash code
    • For example: `java.lang.Object@659e0bfd`

### Overriding Object Implementations

* **(nearly) Always override `toString`**
* **No need to override `equals` and `hashCode` if you want identity semantics**
  * **When in doubt, don't override them**

#### The `equals` Contract

* The equals method implements an **equivalence relation**
  * **Reflexive** – every object is equal to itself
  * **Symmetric** – if `a.equals(b)` then `b.equals(a)`
  * **Transitive** – if `a.equals(b)` and `b.equals(c)`, then `a.equals(c)`
  * **Consistent** – equal objects stay equal unless mutated
  * **“Non-null” **– `a.equals(null)` returns `false`

```java
public final class PhoneNumber {
  private final short areaCode;
  private final short prefix;
  private final short lineNumber;
  
  @Override public boolean equals(Object o) {
    if (!(o instanceof PhoneNumber))  // Does null check
      return false;
    PhoneNumber pn = (PhoneNumber) o; return pn.lineNumber == lineNumber
	}
	...
}
```

### The `hashCode` Contract

* Equal objects **must** have equal hash codes
* **If you override equals you must override `hashCode`**
* Unequal objects **should** have different hash codes
* Hash code must not change unless object mutated

```java
public final class PhoneNumber {
  private final short areaCode;
  private final short prefix;
  private final short lineNumber;
  
	@Override public int hashCode() {
    int result = 17; // Nonzero is good
    result = 31 * result + areaCode; // Constant must be odd
    result = 31 * result + prefix; // " " " "
    result = 31 * result + lineNumber; // " " " "
    return result;
	}
  
  // Or alternative hashCode override
  @Override public int hashCode() {
		return Objects.hash(areaCode, prefix, lineNumber);
	}
	...
}
```

### `==` vs. `equals`

* For primitives you must use `==`
* For object reference types
  * The `==` operator provides identity semantics
    * Exactly as implemented by `Object.equals`
    * Even if `Object.equals` has been overridden
  * **You should (almost) always use** `.equals`**