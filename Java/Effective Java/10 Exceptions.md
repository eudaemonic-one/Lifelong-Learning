# Chapter 10. Exceptions

## Item 69: Use exceptions only for exceptional conditions

* **“Exceptions are, as their name implies, to be used only for exceptional conditions; they should never be used for ordinary control flow.”**
  * “More generally, use standard, easily recognizable idioms in preference to overly clever techniques that purport to offer better performance.”
  * “Even if the performance advantage is real, it may not remain in the face of steadily improving platform implementations.”
  * “The subtle bugs and maintenance headaches that come from overly clever techniques, however, are sure to remain.”
* “A well-designed API must not force its clients to use exceptions for ordinary control flow.”
  * “A class with a “state-dependent” method that can be invoked only under certain unpredictable conditions should generally have a separate “state-testing” method indicating whether it is appropriate to invoke the state-dependent method.”
    * “For example, the `Iterator` interface has the state-dependent method `next` and the corresponding state-testing method `hasNext`.”
  * “An alternative to providing a separate state-testing method is to have the state-dependent method return an empty optional (Item 55) or a distinguished value such as `null` if it cannot perform the desired computation.”
  * “If an object is to be accessed concurrently without external synchronization or is subject to externally induced state transitions, you must use an optional or distinguished return value, as the object’s state could change in the interval between the invocation of a state-testing method and its state-dependent method. Performance concerns may dictate that an optional or distinguished return value be used if a separate state-testing method would duplicate the work of the state-dependent method.”
  * “All other things being equal, a state-testing method is mildly preferable to a distinguished return value. It offers slightly better readability, and incorrect use may be easier to detect: if you forget to call a state-testing method, the state-dependent method will throw an exception, making the bug obvious; if you forget to check for a distinguished return value, the bug may be subtle. ”
* **“In summary, exceptions are designed for exceptional conditions. Don’t use them for ordinary control flow, and don’t write APIs that force others to do so.”**

## Item 70: Use checked exceptions for recoverable conditions and runtime exceptions for programming errors

* “Java provides three kinds of throwables: *checked exceptions*, *runtime exceptions*, and *errors*.”
* “The cardinal rule in deciding whether to use a checked or an unchecked exception is this: **use checked exceptions for conditions from which the caller can reasonably be expected to recover**.”
  * “By throwing a checked exception, you force the caller to handle the exception in a catch clause or to propagate it outward.”
  * “Each checked exception that a method is declared to throw is therefore a potent indication to the API user that the associated condition is a possible outcome of invoking the method.”
* “There are two kinds of unchecked throwables: runtime exceptions and errors.”
  * “They are identical in their behavior: both are throwables that needn’t, and generally shouldn’t, be caught.”
  * “If a program does not catch such a throwable, it will cause the current thread to halt with an appropriate error message.”
* **“Use runtime exceptions to indicate programming errors.”**
  * “The great majority of runtime exceptions indicate precondition violations. A *precondition violation* is simply a failure by the client of an API to adhere to the contract established by the API specification.”
* “While the Java Language Specification does not require it, there is a strong convention that *errors* are reserved for use by the JVM to indicate resource deficiencies, invariant failures, or other conditions that make it impossible to continue execution. ”
  * “Therefore, all of the unchecked throwables you implement should subclass `RuntimeException` (directly or indirectly).”
  * “Not only shouldn’t you define `Error` subclasses, but with the exception of `AssertionError`, you shouldn’t throw them either.”
* “API designers often forget that exceptions are full-fledged objects on which arbitrary methods can be defined. The primary use of such methods is to provide code that catches the exception with additional information concerning the condition that caused the exception to be thrown.”
  * “In the absence of such methods, programmers have been known to parse the string representation of an exception to ferret out additional information. ”

  * “This is extremely bad practice (Item 12).”
  * “Because checked exceptions generally indicate recoverable conditions, it’s especially important for them to provide methods that furnish information to help the caller recover from the exceptional condition.”
* **“To summarize, throw checked exceptions for recoverable conditions and unchecked exceptions for programming errors. When in doubt, throw unchecked exceptions. Don’t define any throwables that are neither checked exceptions nor runtime exceptions. Provide methods on your checked exceptions to aid in recovery.”**

## Item 71: Avoid unnecessary use of checked exceptions

* “Many Java programmers dislike checked exceptions, but used properly, they can improve APIs and programs. Unlike return codes and unchecked exceptions, they *force* programmers to deal with problems, enhancing reliability. That said, overuse of checked exceptions in APIs can make them far less pleasant to use.”
* “If a method throws a single checked exception, this exception is the sole reason the method must appear in a `try` block and can’t be used directly in streams.”
* “The easiest way to eliminate a checked exception is to return an *optional* of the desired result type (Item 55).”
  * “Instead of throwing a checked exception, the method simply returns an empty optional.”
  * “The disadvantage of this technique is that the method can’t return any additional information detailing its inability to perform the desired computation.”
  * “Exceptions, by contrast, have descriptive types, and can export methods to provide additional information (Item 70).”
* “You can also turn a checked exception into an unchecked exception by breaking the method that throws the exception into two methods, the first of which returns a `boolean` indicating whether the exception would be thrown.”
  * “This refactoring is not always appropriate, but where it is, it can make an API more pleasant to use.”

* **“In summary, when used sparingly, checked exceptions can increase the reliability of programs; when overused, they make APIs painful to use. If callers won’t be able to recover from failures, throw unchecked exceptions. If recovery may be possible and you want to *force* callers to handle exceptional conditions, first consider returning an optional. Only if this would provide insufficient information in the case of failure should you throw a checked exception.”**

## Item 72: Favor the use of standard exceptions

* “Reusing standard exceptions has several benefits.”
  * “Chief among them is that it makes your API easier to learn and use because it matches the established conventions that programmers are already familiar with.”
  * “A close second is that programs using your API are easier to read because they aren’t cluttered with unfamiliar exceptions.”
  * “Last (and least), fewer exception classes means a smaller memory footprint and less time spent loading classes.”

* “The most commonly reused exception type is `IllegalArgumentException` (Item 49). This is generally the exception to throw when the caller passes in an argument whose value is inappropriate.”
* “Another commonly reused exception is `IllegalStateException`. This is generally the exception to throw if the invocation is illegal because of the state of the receiving object.”
* “If a caller passes `null` in some parameter for which null values are prohibited, convention dictates that `NullPointerException` be thrown rather than `IllegalArgumentException`.”
* “Similarly, if a caller passes an out-of-range value in a parameter representing an index into a sequence, `IndexOutOfBoundsException` should be thrown rather than `IllegalArgumentException`.”
* “Another reusable exception is `ConcurrentModificationException`. It should be thrown if an object that was designed for use by a single thread (or with external synchronization) detects that it is being modified concurrently.”
* “A last standard exception of note is `UnsupportedOperationException`. This is the exception to throw if an object does not support an attempted operation.”
* **“Do not reuse `Exception`, `RuntimeException`, `Throwable`, or Error directly.”**
* “If an exception fits your needs, go ahead and use it, but only if the conditions under which you would throw it are consistent with the exception’s documentation: reuse must be based on documented semantics, not just on name.”
* “Also, feel free to subclass a standard exception if you want to add more detail (Item 75), but remember that exceptions are serializable (Chapter 12). That alone is reason not to write your own exception class without good reason.”
* “Throw `IllegalStateException` if no argument values would have worked, otherwise throw `IllegalArgumentException`.”

## Item 73: Throw exceptions appropriate to the abstraction

* “It is disconcerting when a method throws an exception that has no apparent connection to the task that it performs.”
  * “This often happens when a method propagates an exception thrown by a lower-level abstraction.”
  * “Not only is it disconcerting, but it pollutes the API of the higher layer with implementation details.”
  * “If the implementation of the higher layer changes in a later release, the exceptions it throws will change too, potentially breaking existing client programs.”
* “To avoid this problem, **higher layers should catch lower-level exceptions and, in their place, throw exceptions that can be explained in terms of the higher-level abstraction**.”


```java
// Exception Translation
try {
    ... // Use lower-level abstraction to do our bidding
} catch (LowerLevelException e) {
    throw new HigherLevelException(...);
}
```

* “A special form of exception translation called *exception chaining* is called for in cases where the lower-level exception might be helpful to someone debugging the problem that caused the higher-level exception. The lower-level exception (the cause) is passed to the higher-level exception, which provides an accessor method (`Throwable`’s `getCause` method) to retrieve the lower-level exception:”

```java
// Exception Chaining
try {
    ... // Use lower-level abstraction to do our bidding
} catch (LowerLevelException cause) {
    throw new HigherLevelException(cause);
} 
```

* “The higher-level exception’s constructor passes the cause to a *chaining-aware* superclass constructor, so it is ultimately passed to one of `Throwable`’s chaining-aware constructors, such as `Throwable(Throwable)`:”


```java
// Exception with chaining-aware constructor
class HigherLevelException extends Exception {
    HigherLevelException(Throwable cause) {
        super(cause);
    }
}
```

* **“While exception translation is superior to mindless propagation of exceptions from lower layers, it should not be overused.”**
  * “Where possible, the best way to deal with exceptions from lower layers is to avoid them, by ensuring that lower-level methods succeed.”
* “If it is impossible to prevent exceptions from lower layers, the next best thing is to have the higher layer silently work around these exceptions, insulating the caller of the higher-level method from lower-level problems.”
* **“In summary, if it isn’t feasible to prevent or to handle exceptions from lower layers, use exception translation, unless the lower-level method happens to guarantee that all of its exceptions are appropriate to the higher level. Chaining provides the best of both worlds: it allows you to throw an appropriate higher-level exception, while capturing the underlying cause for failure analysis (Item 75).”**
