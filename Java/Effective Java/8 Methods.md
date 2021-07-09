# Chapter 8. Methods

## Item 49: Check parameters for validity

* “Most methods and constructors have some restrictions on what values may be passed into their parameters.”
  * “For example, it is not uncommon that index values must be non-negative and object references must be non-null.”
  * “You should clearly document all such restrictions and enforce them with checks at the beginning of the method body.”
* “If an invalid parameter value is passed to a method and the method checks its parameters before execution, it will fail quickly and cleanly with an appropriate exception.”
* “For public and protected methods, use the Javadoc `@throws` tag to document the exception that will be thrown if a restriction on parameter values is violated (Item 74).”
  * “Typically, the resulting exception will be `IllegalArgumentException`, `IndexOutOfBoundsException`, or `NullPointerException` (Item 72). ”

```java
/**
 * Returns a BigInteger whose value is (this mod m). This method
 * differs from the remainder method in that it always returns a
 * non-negative BigInteger.
 *
 * @param m the modulus, which must be positive
 * @return this mod m
 * @throws ArithmeticException if m is less than or equal to 0
 */
public BigInteger mod(BigInteger m) {
    if (m.signum() <= 0)
        throw new ArithmeticException("Modulus <= 0: " + m);
    ... // Do the computation
}
```

* “The class-level comment applies to all parameters in all of the class’s public methods.”
  * “This is a good way to avoid the clutter of documenting every `NullPointerException` on every method individually.”
  * “It may be combined with the use of `@Nullable` or a similar annotation to indicate that a particular parameter may be null.”
* **“The `Objects.requireNonNull` method, added in Java 7, is flexible and convenient, so there’s no reason to perform null checks manually anymore.”**
* “Nonpublic methods can check their parameters using *assertions*.”
  * “Unlike normal validity checks, assertions throw `AssertionError` if they fail.”
  * “And unlike normal validity checks, they have no effect and essentially no cost unless you enable them, which you do by passing the `-ea` (or `-enableassertions`) flag to the `java` command.”


```java
// Private helper function for a recursive sort
private static void sort(long a[], int offset, int length) {
    assert a != null;
    assert offset >= 0 && offset <= a.length;
    assert length >= 0 && length <= a.length - offset;
    ... // Do the computation
}
```

* **“It is particularly important to check the validity of parameters that are not used by a method, but stored for later use.”**
  * “It is critical to check the validity of constructor parameters to prevent the construction of an object that violates its class invariants.”
* “There are exceptions to the rule that you should explicitly check a method’s parameters before performing its computation. ”
  * “An important exception is the case in which the validity check would be expensive or impractical *and* the check is performed implicitly in the process of doing the computation.”
* “Occasionally, a computation implicitly performs a required validity check but throws the wrong exception if the check fails. ”
  * “Under these circumstances, you should use the *exception translation* idiom, described in Item 73, to translate the natural exception into the correct one.”

* **“To summarize, each time you write a method or constructor, you should think about what restrictions exist on its parameters. You should document these restrictions and enforce them with explicit checks at the beginning of the method body. It is important to get into the habit of doing this. The modest work that it entails will be paid back with interest the first time a validity check fails.”**