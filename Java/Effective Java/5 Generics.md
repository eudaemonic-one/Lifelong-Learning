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

## Item 28: Prefer lists to arrays

* “Arrays differ from generic types in two important ways.”
* “First, arrays are *covariant*.”
  * “This scary-sounding word means simply that if `Sub` is a subtype of `Super`, then the array type `Sub[]` is a subtype of the array type `Super[]`.”
  * “Generics, by contrast, are invariant: for any two distinct types `Type1` and `Type2`, `List<Type1>` is neither a subtype nor a supertype of `List<Type2>` [JLS, 4.10; Naftalin07, 2.5].”

```java
// Fails at runtime!
Object[] objectArray = new Long[1];
objectArray[0] = "I don't fit in"; // Throws ArrayStoreException

// Won't compile!
List<Object> ol = new ArrayList<Long>(); // Incompatible types
ol.add("I don't fit in");
```

* “The second major difference between arrays and generics is that arrays are *reified* [JLS, 4.7]. This means that arrays know and enforce their element type at runtime.”
  * “As noted earlier, if you try to put a `String` into an array of `Long`, you’ll get an `ArrayStoreException`.”
  * “Generics, by contrast, are implemented by *erasure* [JLS, 4.6]. This means that they enforce their type constraints only at compile time and discard (or *erase*) their element type information at runtime.”
* “Because of these fundamental differences, arrays and generics do not mix well.”
  * “For example, it is illegal to create an array of a generic type, a parameterized type, or a type parameter.”
* “Why is it illegal to create a generic array? Because it isn’t typesafe. If it were legal, casts generated by the compiler in an otherwise correct program could fail at runtime with a `ClassCastException`.”
* “**When you get a generic array creation error or an unchecked cast warning on a cast to an array type, the best solution is often to use the collection type `List<E>` in preference to the array type `E[]`.** You might sacrifice some conciseness or performance, but in exchange you get better type safety and interoperability.”


## Item 29: Favor generic types

```java
// Object-based collection - a prime candidate for generics
public class Stack {
    private Object[] elements;
    private int size = 0;
    private static final int DEFAULT_INITIAL_CAPACITY = 16;

    public Stack() {
        elements = new Object[DEFAULT_INITIAL_CAPACITY];
    }

    public void push(Object e) {
        ensureCapacity();
        elements[size++] = e;
    }

    public Object pop() {
        if (size == 0)
            throw new EmptyStackException();
        Object result = elements[--size];
        elements[size] = null; // Eliminate obsolete reference
        return result;
    }

    public boolean isEmpty() {
        return size == 0;
    }

    private void ensureCapacity() {
        if (elements.length == size)
            elements = Arrays.copyOf(elements, 2 * size + 1);
    }
} 
```

* “This class should have been parameterized to begin with, but since it wasn’t, we can generify it after the fact. In other words, we can parameterize it without harming clients of the original non-parameterized version.”

* “As explained in Item 28, you can’t create an array of a non-reifiable type, such as `E`. This problem arises every time you write a generic type that is backed by an array.”

```java
Stack.java:8: warning: [unchecked] unchecked cast
found: Object[], required: E[]
        elements = (E[]) new Object[DEFAULT_INITIAL_CAPACITY];
                       ^
```

* “The compiler may not be able to prove that your program is typesafe, but you can. You must convince yourself that the unchecked cast will not compromise the type safety of the program.”
* “Once you’ve proved that an unchecked cast is safe, suppress the warning in as narrow a scope as possible (Item 27).”


```java
// The elements array will contain only E instances from push(E).
// This is sufficient to ensure type safety, but the runtime
// type of the array won't be E[]; it will always be Object[]!
@SuppressWarnings("unchecked")
public Stack() {
    elements = (E[]) new Object[DEFAULT_INITIAL_CAPACITY];
}
```

* **“In summary, generic types are safer and easier to use than types that require casts in client code. When you design new types, make sure that they can be used without such casts. This will often mean making the types generic. If you have any existing types that should be generic but aren’t, generify them. This will make life easier for new users of these types without breaking existing clients (Item 26).”**

## Item 30: Favor generic methods

* “Just as classes can be generic, so can methods. Static utility methods that operate on parameterized types are usually generic.”

```java
// Uses raw types - unacceptable! (Item 26)
public static Set union(Set s1, Set s2) {
    Set result = new HashSet(s1);
    result.addAll(s2);
    return result;
}

Union.java:5: warning: [unchecked] unchecked call to
HashSet(Collection<? extends E>) as a member of raw type HashSet
        Set result = new HashSet(s1);
                     ^
Union.java:6: warning: [unchecked] unchecked call to
addAll(Collection<? extends E>) as a member of raw type Set
        result.addAll(s2);
                     ^
```

* “To fix these warnings and make the method typesafe, modify its declaration to declare a *type parameter* representing the element type for the three sets (the two arguments and the return value) and use this type parameter throughout the method.”
* “The type parameter list, which declares the type parameters, goes between a method’s modifiers and its return type.”


```java
// Generic method
public static <E> Set<E> union(Set<E> s1, Set<E> s2) {
    Set<E> result = new HashSet<>(s1);
    result.addAll(s2);
    return result;
}
```

* “A limitation of the `union` method is that the types of all three sets (both input parameters and the return value) have to be exactly the same. You can make the method more flexible by using *bounded wildcard* types (Item 31).”
* “On occasion, you will need to create an object that is immutable but applicable to many different types. Because generics are implemented by erasure (Item 28), you can use a single object for all required type parameterizations, but you need to write a static factory method to repeatedly dole out the object for each requested type parameterization. This pattern, called the *generic singleton factory*, is used for function objects (Item 42) such as `Collections.reverseOrder`, and occasionally for collections such as `Collections.emptySet`.”


```java
// Generic singleton factory pattern
private static UnaryOperator<Object> IDENTITY_FN = (t) -> t;

@SuppressWarnings("unchecked")
public static <T> UnaryOperator<T> identityFunction() {
    return (UnaryOperator<T>) IDENTITY_FN;
}
```

* “The cast of `IDENTITY_FN` to (`UnaryFunction<T>`) generates an unchecked cast warning, as `UnaryOperator<Object>` is not a `UnaryOperator<T>` for every `T`. But the identity function is special: it returns its argument unmodified, so we know that it is typesafe to use it as a `UnaryFunction<T>`, whatever the value of `T`. Therefore, we can confidently suppress the unchecked cast warning generated by this cast. Once we’ve done this, the code compiles without error or warning.”


```java
public interface Comparable<T> {
    int compareTo(T o);
}
```

* “Many methods take a collection of elements implementing `Comparable` to sort it, search within it, calculate its minimum or maximum, and the like. To do these things, it is required that every element in the collection be comparable to every other element in it, in other words, that the elements of the list be *mutually comparable*.”


```java
// Using a recursive type bound to express mutual comparability
public static <E extends Comparable<E>> E max(Collection<E> c);
```

* “Recursive type bounds can get much more complex, but luckily they rarely do. If you understand this idiom, its wildcard variant (Item 31), and the *simulated self-type* idiom (Item 2), you’ll be able to deal with most of the recursive type bounds you encounter in practice.”
* **“In summary, generic methods, like generic types, are safer and easier to use than methods requiring their clients to put explicit casts on input parameters and return values. Like types, you should make sure that your methods can be used without casts, which often means making them generic. And like types, you should generify existing methods whose use requires casts. This makes life easier for new users without breaking existing clients (Item 26).”**
