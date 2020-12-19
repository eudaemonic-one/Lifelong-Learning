# Lecture 4 API Documentation

* If you build it but don't document it properly, they will not come

## What Consititute API

* Type specifications
* Method specifications
* Not implementation (Any details)

## How to Document API

* API documentation is **critical**
* A potent pair of Java design decisions that are often overlooked
  * Decision 1: Eliminate header files
  * Decision 2: Support the Javadoc tool
  * Result: declaration, documentation, and implementation of an API are collocated
* Document **every**
  * type (class or interface)
    * Say exactly what an instance represents
    * Generally a simple noun phrase
  * method
  * constructor
  * parameter
  * exception

### Type Doc

* Things that apply to every method in the type
* How to obtain an instance of the type
  * This is a classic **discoverability problem**
* Important relationships to other types
* Thread-safety - any restrictions on concurrent use
* Sample code involving multiole methods

#### Example Code Must Be Exemplary

* **Perhaps the most important of all code**
* It forms the basis of real code that uses API

#### Mutable Types are Hard to Document

* When is it legal to call which method?
* What happens when multiple threads use object?
* **You must document the state space**

### Method Doc

* **Describe contract between method & caller**
  * **Preconditions** - what must be true before the call
  * **Postconditions** - what will be true after the call
  * **Side-effects** - Any state mutation that is not clearly implied by the name of the method being called
  * **Thread-safety** - If specific to method
* Sample code, if specific to method

### Parameter or Field Doc

* Say **precisely** what it represents
  * Typically a short noun phrase will do
* Don't forget to include, where appropriate
  * **Units**
  * **Form**
  * **Restrictions**

### Exception Doc

* **Say precisely when the method throws it**
* **If it's not a precondition violation, say what the caller can do in response, if appropriate**
* Document all exceptions, unchecked and checked

### Documentation Generators

* Documentation generators are great, but they're not a panacea
* Large APIs require big picture documentation
* In some cases, package level doc will suffice
* In others, you'll need external documentation

### Test the Doc

* **Have a colleague use your doc to write to the API**
* Don' make your customer do this
* Testing documentation is similar to testing code
  * Start early, and test often
* If someone else is documenting your API
  * Respect the documenter and establish a dialogue

