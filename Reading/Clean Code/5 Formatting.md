# Chapter 5: Formatting

## The Purpose of Formatting

* The coding style and readability set precedents that continue to affect maintainability and extensibility long after the original code has been changed beyond recognition.

## Vertical Formatting

* In Java, file size is closely related to class size.
  * It appears to be possible to build significant systems out of files that are typically 200 lines long, with an upper limit of 500.

### The Newspaper Metaphor

* We would like a source file to be like a newspaper article. The name should be simple but explanatory. The name, by itself, should be sufficient to tell us whether we are in the right module or not. The topmost parts of the source file should provide the high-level concepts and algorithms. Detail should increase as we move downward, until at the end we find the lowest level functions and details in the source file.

### Vertical Openness Between Concepts

* Each expression or a clause, and each group of lines should be separated from each other with blank lines.

### Vertical Density

* Lines of code that are tightly related should appear vertically dense.

### Vertical Distance

* Concepts that are closely related should be kept vertically close to each other.
* Closely related concepts should not be separated into different files unless you have a very good reason.
* **Variable Declarations**
  * Variables should be declared as close to their usage as possible.
  * Control variables for loops should usually be declared within the loop statement.
* **Instance variables** should be declared at the top of the class.
* **Dependent Functions**
  * If one function calls another, they should be vertically close, and the caller should be above the callee. if at all possible. This gives the program a natural flow.
* **Conceptual Affinity**
  * Affinity might be caused because a group of functions perform a similar operation.

### Vertical Ordering

* A function that is called should be below a function that does the calling
