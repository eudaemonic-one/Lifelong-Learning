# Chapter 9. General Programming

## Item 57: Minimize the scope of local variables

* “By minimizing the scope of local variables, you increase the readability and maintainability of your code and reduce the likelihood of error.”
* **“The most powerful technique for minimizing the scope of a local variable is to declare it where it is first used.”**
  * “Declaring a local variable prematurely can cause its scope not only to begin too early but also to end too late.”
* **“Nearly every local variable declaration should contain an initializer.”**
  * “If you don’t yet have enough information to initialize a variable sensibly, you should postpone the declaration until you do.”
    * “One exception to this rule concerns `try`-`catch` statements. If a variable is initialized to an expression whose evaluation can throw a checked exception, the variable must be initialized inside a `try` block (unless the enclosing method can propagate the exception).”
    * “If the value must be used outside of the `try` block, then it must be declared before the `try` block, where it cannot yet be “sensibly initialized.”
* “Loops present a special opportunity to minimize the scope of variables. The `for` loop, in both its traditional and for-each forms, allows you to declare *loop variables*, limiting their scope to the exact region where they’re needed.”
  * “Therefore, **prefer `for` loops to `while` loops**, assuming the contents of the loop variable aren’t needed after the loop terminates.”
* “If you need access to the iterator, perhaps to call its `remove` method, the preferred idiom uses a traditional `for` loop in place of the for-each loop:”


```java
// Idiom for iterating when you need the iterator
for (Iterator<Element> i = c.iterator(); i.hasNext(); ) {
    Element e = i.next();
    ... // Do something with e and i
}
```

* “Here is another loop idiom that minimizes the scope of local variables:”

```java
for (int i = 0, n = expensiveComputation(); i < n; i++) {
    ... // Do something with i;
}
```

* “The important thing to notice about this idiom is that it has *two* loop variables, `i` and `n`, both of which have exactly the right scope. The second variable, `n`, is used to store the limit of the first, thus avoiding the cost of a redundant computation in every iteration.”
* “A final technique to minimize the scope of local variables is to **keep methods small and focused**.”
  * “If you combine two activities in the same method, local variables relevant to one activity may be in the scope of the code performing the other activity. To prevent this from happening, simply separate the method into two: one for each activity.”
