# Chapter 5. Generics

## Item 26: Don’t use raw types

* “A class or interface whose declaration has one or more *type parameters* is a *generic* class or interface [JLS, 8.1.2, 9.1.2].”
* “Each generic type defines a set of *parameterized types*, which consist of the class or interface name followed by an angle-bracketed list of *actual type parameters* corresponding to the generic type’s *formal type parameters* [JLS, 4.4, 4.5].”
* “Each generic type defines a *raw type*, which is the name of the generic type used without any accompanying type parameters [JLS, 4.8].”
  * “Raw types behave as if all of the generic type information were erased from the type declaration. They exist primarily for compatibility with pre-generics code.”

```java
// Raw collection type - don't do this!

// My stamp collection. Contains only Stamp instances.
private final Collection stamps = ... ;

// Erroneous insertion of coin into stamp collection
stamps.add(new Coin( ... )); // Emits "unchecked call" warning

// Raw iterator type - don't do this!
for (Iterator i = stamps.iterator(); i.hasNext(); )
    Stamp stamp = (Stamp) i.next(); // Throws ClassCastException
        stamp.cancel();
```

```java
// Parameterized collection type - typesafe
private final Collection<Stamp> stamps = ... ;
```

* “From this declaration, the compiler knows that `stamps` should contain only `Stamp` instances and *guarantees* it to be true, assuming your entire codebase compiles without emitting (or suppressing; see Item 27) any warnings.”
* **“If you use raw types, you lose all the safety and expressiveness benefits of generics.”**
* “While you shouldn’t use raw types such as `List`, it is fine to use types that are parameterized to allow insertion of arbitrary objects, such as `List<Object>`.”


```java
// Use of raw type for unknown element type - don't do this!
static int numElementsInCommon(Set s1, Set s2) {
    int result = 0;
    for (Object o1 : s1)
        if (s2.contains(o1))
            result++;
    return result;
}
```

* “This method works but it uses raw types, which are dangerous. The safe alternative is to use *unbounded wildcard types*. If you want to use a generic type but you don’t know or care what the actual type parameter is, you can use a question mark instead. ”
* “If these restrictions are unacceptable, you can use *generic methods* (Item 30) or *bounded wildcard types* (Item 31).”
* “There are a few minor exceptions to the rule that you should not use raw types.”
  * **“You must use raw types in class literals.”**
    * “In other words, `List.class`, `String[].class`, and `int.class` are all legal, but `List<String>.class` and `List<?>.class` are not.”
  * “A second exception to the rule concerns the instanceof operator. Because generic type information is erased at runtime, it is illegal to use the `instanceof` operator on parameterized types other than unbounded wildcard types.”

```java
// Legitimate use of raw type - instanceof operator
if (o instanceof Set) {       // Raw type
    Set<?> s = (Set<?>) o;    // Wildcard type
    ...
}
```

* **“In summary, using raw types can lead to exceptions at runtime, so don’t use them. They are provided only for compatibility and interoperability with legacy code that predates the introduction of generics. As a quick review, `Set<Object>` is a parameterized type representing a set that can contain objects of any type, `Set<?>` is a wildcard type representing a set that can contain only objects of some unknown type, and `Set` is a raw type, which opts out of the generic type system. The first two are safe, and the last is not.”**

## Item 27: Eliminate unchecked warnings

* “When you program with generics, you will see many compiler warnings: unchecked cast warnings, unchecked method invocation warnings, unchecked parameterized vararg type warnings, and unchecked conversion warnings.”

```java
Set<Lark> exaltation = new HashSet();

Venery.java:4: warning: [unchecked] unchecked conversion
        Set<Lark> exaltation = new HashSet();
                               ^
  required: Set<Lark>
  found:    HashSet
```

* “You can then make the indicated correction, causing the warning to disappear. Note that you don’t actually have to specify the type parameter, merely to indicate that it’s present with the diamond operator (<>), introduced in Java 7.”
* **“Eliminate every unchecked warning that you can.”**
  * “If you eliminate all warnings, you are assured that your code is typesafe, which is a very good thing. It means that you won’t get a `ClassCastException` at runtime, and it increases your confidence that your program will behave as you intended.”
* **“If you can’t eliminate a warning, but you can prove that the code that provoked the warning is typesafe, then (and only then) suppress the warning with an `@SuppressWarnings("unchecked")` annotation.”**
  * **“Always use the `SuppressWarnings` annotation on the smallest scope possible.”**
  * **“Every time you use a `@SuppressWarnings("unchecked")` annotation, add a comment saying why it is safe to do so.”**

```java
// Adding local variable to reduce scope of @SuppressWarnings
public <T> T[] toArray(T[] a) {
    if (a.length < size) {
        // This cast is correct because the array we're creating
        // is of the same type as the one passed in, which is T[].
        @SuppressWarnings("unchecked") T[] result =
            (T[]) Arrays.copyOf(elements, size, a.getClass());
        return result;
    }
    System.arraycopy(elements, 0, a, 0, size);
    if (a.length > size)
        a[size] = null;
    return a;
}
```

