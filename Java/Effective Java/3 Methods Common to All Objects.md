# Chapter 3. Methods Common to All Objects

* “All of `Object`s nonfinal methods (`equals`, `hashCode`, `toString`, `clone`, and `finalize`) have explicit general contracts because they are designed to be overridden.”

## Item 10: Obey the general contract when overriding equals

* “The easiest way to avoid problems is not to override the equals method, in which case each instance of the class is equal only to itself.”

  * “Each instance of the class is inherently unique.”
    * e.g., `Thread`
  * “There is no need for the class to provide a “logical equality” test.”
    * e.g., `java.util.regex.Pattern`
  * “A superclass has already overridden equals, and the superclass behavior is appropriate for this class.”
    * e.g., `Set` inherits `equals` from `AbstractSet`
  * “The class is private or package-private, and you are certain that its equals method will never be invoked.”
* “So when is it appropriate to override `equals`? It is when a class has a notion of *logical equality* that differs from mere object identity and a superclass has not already overridden `equals`. ”
* “When you override the `equals` method, you must adhere to its **general contract**.”

  * “*Reflexive*: For any non-null reference value `x`, `x.equals(x)` must return `true`.”
  * “*Symmetric*: For any non-null reference values `x` and `y`, `x.equals(y)` must return `true` if and only if `y.equals(x)` returns `true`.”
  * “*Transitive*: For any non-null reference values `x`, `y`, `z`, if `x.equals(y)` returns `true` and `y.equals(z)` returns `true`, then `x.equals(z)` must return `true`.”
  * “*Consistent*: For any non-null reference values `x` and `y`, multiple invocations of `x.equals(y)` must consistently return true or consistently return `false`, provided no information used in equals comparisons is modified.”
  * “For any non-null reference value `x`, `x.equals(null)` must return `false`.”
* **“Once you’ve violated the equals contract, you simply don’t know how other objects will behave when confronted with your object.”**
* “Putting it all together, here’s a recipe for a high-quality equals method:”
  * **“Use the `==` operator to check if the argument is a reference to this object.”**
  * **“Use the `instanceof` operator to check if the argument has the correct type.”**
  * **“Cast the argument to the correct type.”**
  * **“For each “significant” field in the class, check if that field of the argument matches the corresponding field of this object.”**
* **“When you are finished writing your equals method, ask yourself three questions: Is it symmetric? Is it transitive? Is it consistent?”**
  * **“Always override hashCode when you override equals (Item 11).”**
  * **“Don’t try to be too clever.”** 
  * **“Don’t substitute another type for Object in the equals declaration.”**
* **“In summary, don’t override the `equals` method unless you have to: in many cases, the implementation inherited from `Object` does exactly what you want. If you do override `equals`, make sure to compare all of the class’s significant fields and to compare them in a manner that preserves all five provisions of the `equals` contract.”**

## Item 11: Always override `hashCode` when you override `equals`

* **“You must override `hashCode` in every class that overrides `equals`.”**
* **“The key provision that is violated when you fail to override `hashCode` is the second one: equal objects must have equal hash codes.”**
* A simple recipe:
  * “1. Declare an `int` variable named `result`, and initialize it to the hash code `c` for the first significant field in your object, as computed in step 2.a. (Recall from Item 10 that a significant field is a field that affects equals comparisons.)”
  * “2. For every remaining significant field f in your object, do the following:”
    * “a. Compute an `int` hash code `c` for the field:”
      * “If the field is of a primitive type, compute `Type.hashCode(f)`, where `Type` is the boxed primitive class corresponding to `f`’s type.”
      * “ii. If the field is an object reference and this class’s `equals` method compares the field by recursively invoking `equals`, recursively invoke `hashCode` on the field. If a more complex comparison is required, compute a “canonical representation” for this field and invoke `hashCode` on the canonical representation. If the value of the field is `null`, use `0` (or some other constant, but 0 is traditional).”
      * “iii. If the field is an array, treat it as if each significant element were a separate field. That is, compute a hash code for each significant element by applying these rules recursively, and combine the values per step 2.b. If the array has no significant elements, use a constant, preferably not `0`. If all elements are significant, use `Arrays.hashCode`.”
    * “b. Combine the hash code `c` computed in step 2.a into result as follows: `result = 31 * result + c;`”
  * 3. Return result.

```java
// Typical hashCode method
@Override public int hashCode() {
    int result = Short.hashCode(areaCode);
    result = 31 * result + Short.hashCode(prefix);
    result = 31 * result + Short.hashCode(lineNum);
    return result;
}
```

* “The `Objects` class has a static method that takes an arbitrary number of objects and returns a hash code for them. This method, named `hash`, lets you write one-line `hashCode` methods whose quality is comparable to those written according to the recipe in this item. Unfortunately, they run more slowly because they entail array creation to pass a variable number of arguments, as well as boxing and unboxing if any of the arguments are of primitive type. This style of hash function is recommended for use only in situations where performance is not critical.”

```java
// One-line hashCode method - mediocre performance
@Override public int hashCode() {
   return Objects.hash(lineNum, prefix, areaCode);
}
```

* “If a class is immutable and the cost of computing the hash code is significant, you might consider caching the hash code in the object rather than recalculating it each time it is requested.”

```java
// hashCode method with lazily initialized cached hash code
private int hashCode; // Automatically initialized to 0

@Override public int hashCode() {
    int result = hashCode;
    if (result == 0) {
        result = Short.hashCode(areaCode);
        result = 31 * result + Short.hashCode(prefix);
        result = 31 * result + Short.hashCode(lineNum);
        hashCode = result;
    }
    return result;
}
```

* **“Do not be tempted to exclude significant fields from the hash code computation to improve performance.”**
* **“Don’t provide a detailed specification for the value returned by `hashCode`, so clients can’t reasonably depend on it; this gives you the flexibility to change it.”**

## Item 12: Always override `toString`

* “While `Object` provides an implementation of the `toString` method, the string that it returns is generally not what the user of your class wants to see. It consists of the class name followed by an “at” sign (`@`) and the unsigned hexadecimal representation of the hash code, for example, `PhoneNumber@163b91`.”
* **“Providing a good `toString` implementation makes your class much more pleasant to use and makes systems using the class easier to debug.”**
  * “Even if you never call toString on an object, others may.”
* **“When practical, the toString method should return all of the interesting information contained in the object.”**
* **“One important decision you’ll have to make when implementing a `toString` method is whether to specify the format of the return value in the documentation.”**
  * “It is recommended that you do this for value classes, such as phone number or matrix. The advantage of specifying the format is that it serves as a standard, unambiguous, human-readable representation of the object. This representation can be used for input and output and in persistent human-readable data objects, such as CSV files.”
  * “The disadvantage of specifying the format of the `toString` return value is that once you’ve specified it, you’re stuck with it for life, assuming your class is widely used. Programmers will write code to parse the representation, to generate it, and to embed it into persistent data. If you change the representation in a future release, you’ll break their code and data, and they will yowl.”

  * “Whether or not you decide to specify the format, **you should clearly document your intentions**.”
  * “Whether or not you specify the format, **provide programmatic access to the information contained in the value returned by `toString`**.”
    * “By failing to provide accessors, you turn the string format into a de facto API, even if you’ve specified that it’s subject to change.”