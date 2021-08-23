# Chapter 3: Functions

## Small

* Functions should be very small.

### Blocks and Indenting

* Blocks within `if` statements, `else` statements, `while` statements, and so on should be one line long.
* Functions should not be large enough to hold nested structures.

## Do One Thing

* Functions should do one thing. They should do it well. They should do it only.
  * If a function does only those steps that are one level below the stated name of the function, then the function is doing one thing.
  * Doing more than "one thing" is if you can extract another function from it with a name that is not merely a restatement of its implementation.

### Sections within Functions

* Functions that do one thing cannot be reasonably divided into sections.

## One Level of Abstraction per Function

* Once details are mixed with essential concepts, more and more details tend to accrete within the function.

### Reading Code from Top to Bottom: The Stepdown Rule

* We want every function to be followed by those at the next level of abstraction so that we can read the program.
