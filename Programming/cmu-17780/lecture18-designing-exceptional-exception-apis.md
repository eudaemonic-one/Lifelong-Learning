# Lecture 18 Designing Exceptional Exception API

* Use exceptions only for exceptional conditions
* Never use exceptions for normal control flow
* Prefer common idioms
* Beware the logical AND and OR operators

## Exception Design Principles

* Throw exceptions to indicate exceptional conditions
  * Don't force client to use exceptions for control flow
  * Conversely, don't fail silently
* Favor unchecked exceptions
  * Checked - client must take recovery action
  * Unchecked - generally a programming error
  * Overuse of checked exceptions causes boilerplate
* Include failure-capture information in exceptions
  * e.g., `IndexOutOfBoundsException` should include index and, ideally bounds of access
  * Eases diagnosis and repair or recovery
  * For unchecked exceptions, message suffices
  * For checked exceptions, provide accessors too