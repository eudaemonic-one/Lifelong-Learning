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

## Item 50: Make defensive copies when needed

* **“You must program defensively, with the assumption that clients of your class will do their best to destroy its invariants.”**

```java
// Broken "immutable" time period class
public final class Period {
    private final Date start;
    private final Date end;

    /**
     * @param  start the beginning of the period
     * @param  end the end of the period; must not precede start
     * @throws IllegalArgumentException if start is after end
     * @throws NullPointerException if start or end is null
     */
    public Period(Date start, Date end) {
        if (start.compareTo(end) > 0)
            throw new IllegalArgumentException(
                start + " after " + end);
        this.start = start;
        this.end   = end;
    }

    public Date start() {
        return start;
    }

    public Date end() {
        return end;
    }

    ...    // Remainder omitted
}
```

* “It is, however, easy to violate this invariant by exploiting the fact that `Date` is mutable.”

```java
// Attack the internals of a Period instance
Date start = new Date();
Date end = new Date();
Period p = new Period(start, end);
end.setYear(78);  // Modifies internals of p!
```

* “As of Java 8, the obvious way to fix this problem is to use `Instant` (or `LocalDateTime` or `ZonedDateTime`) in place of a `Date` because `Instant` (and the other `java.time` classes) are immutable (Item 17).”
  * **“`Date` is obsolete and should no longer be used in new code.”**
* “To protect the internals of a `Period` instance from this sort of attack, **it is essential to make a defensive copy of each mutable parameter to the constructor** and to use the copies as components of the `Period` instance in place of the originals.”


```java
// Repaired constructor - makes defensive copies of parameters
public Period(Date start, Date end) {
    this.start = new Date(start.getTime());
    this.end   = new Date(end.getTime());

    if (this.start.compareTo(this.end) > 0)
      throw new IllegalArgumentException(
          this.start + " after " + this.end);
}
```

* “Note that **defensive copies are made before checking the validity of the parameters (Item 49), and the validity check is performed on the copies rather than on the originals**.”
  * “It protects the class against changes to the parameters from another thread during the *window of vulnerability* between the time the parameters are checked and the time they are copied. In the computer security community, this is known as a *time-of-check*/*time-of-use* or *TOCTOU* attack [Viega01].”
* **“Do not use the `clone` method to make a defensive copy of a parameter whose type is subclassable by untrusted parties.”**
  * “That said, you are generally better off using a constructor or static factory to copy an instance, for reasons outlined in Item 13.”


```java
// Second attack on the internals of a Period instance
Date start = new Date();
Date end = new Date();
Period p = new Period(start, end);
p.end().setYear(78);  // Modifies internals of p!
```

* “To defend against the second attack, merely modify the accessors to **return defensive copies of mutable internal fields**:”


```java
// Repaired accessors - make defensive copies of internal fields
public Date start() {
    return new Date(start.getTime());
}

public Date end() {
    return new Date(end.getTime());
}
```

* “Any time you write a method or constructor that stores a reference to a client-provided object in an internal data structure, think about whether the client-provided object is potentially mutable. If it is, think about whether your class could tolerate a change in the object after it was entered into the data structure. If the answer is no, you must defensively copy the object and enter the copy into the data structure in place of the original.”
* “The same is true for defensive copying of internal components prior to returning them to clients.”
  * “Remember that nonzero-length arrays are always mutable. Therefore, you should always make a defensive copy of an internal array before returning it to a client. Alternatively, you could return an immutable view of the array. ”
* “There may be a performance penalty associated with defensive copying and it isn’t always justified.”
  * “If a class trusts its caller not to modify an internal component, perhaps because the class and its client are both part of the same package, then it may be appropriate to dispense with defensive copying.”
  * “Under these circumstances, the class documentation should make it clear that the caller must not modify the affected parameters or return values.”

* **“In summary, if a class has mutable components that it gets from or returns to its clients, the class must defensively copy these components. If the cost of the copy would be prohibitive *and* the class trusts its clients not to modify the components inappropriately, then the defensive copy may be replaced by documentation outlining the client’s responsibility not to modify the affected components.”**

## Item 51: Design method signatures carefully

* **“Choose method names carefully.”**
  * “Names should always obey the standard naming conventions (Item 68).”
  * “Your primary goal should be to choose names that are understandable and consistent with other names in the same package.”
  * “Your secondary goal should be to choose names consistent with the broader consensus, where it exists”
  * “Avoid long method names.”
  * “When in doubt, look to the Java library APIs for guidance.”
* **“Don’t go overboard in providing convenience methods.”**
  * “Too many methods make a class difficult to learn, use, document, test, and maintain.”
  * “This is doubly true for interfaces, where too many methods complicate life for implementors as well as users.”
  * “Consider providing a “shorthand” only if it will be used often.”
  * **“When in doubt, leave it out.”**
* **“Avoid long parameter lists.”**
  * “Aim for four parameters or fewer.”
  * **“Long sequences of identically typed parameters are especially harmful.”**
    * “Not only won’t users be able to remember the order of the parameters, but when they transpose parameters accidentally, their programs will still compile and run.”
* “There are three techniques for shortening overly long parameter lists.”
  * “One is to break the method up into multiple methods, each of which requires only a subset of the parameters.”
  * “A second technique for shortening long parameter lists is to create *helper classes* to hold groups of parameters.”
    * “Typically these helper classes are static member classes (Item 24). This technique is recommended if a frequently occurring sequence of parameters is seen to represent some distinct entity.”

  * “A third technique that combines aspects of the first two is to adapt the Builder pattern (Item 2) from object construction to method invocation.”
* “**For parameter types, favor interfaces over classes** (Item 64).”
  * “By using a class instead of an interface, you restrict your client to a particular implementation and force an unnecessary and potentially expensive copy operation if the input data happens to exist in some other form.”
* “**Prefer two-element enum types to `boolean` parameters**, unless the meaning of the boolean is clear from the method name.”
  * “Enums make your code easier to read and to write.”
  * “Also, they make it easy to add more options later.”

```java
public enum TemperatureScale { FAHRENHEIT, CELSIUS }
```

