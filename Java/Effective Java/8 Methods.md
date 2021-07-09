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

## Item 52: Use overloading judiciously

```java
// Broken! - What does this program print?
public class CollectionClassifier {
    public static String classify(Set<?> s) {
        return "Set";
    }

    public static String classify(List<?> lst) {
        return "List";
    }

    public static String classify(Collection<?> c) {
        return "Unknown Collection";
    }

    public static void main(String[] args) {
        Collection<?>[] collections = {
            new HashSet<String>(),
            new ArrayList<BigInteger>(),
            new HashMap<String, String>().values()
        };

        for (Collection<?> c : collections)
            System.out.println(classify(c));
    }
}
```

* “You might expect this program to print `Set`, followed by `List` and `Unknown Collection`, but it doesn’t. It prints `Unknown Collection` three times.”
* “Why does this happen? Because the `classify` method is *overloaded*, and **the choice of which overloading to invoke is made at compile time**.”
* “The behavior of this program is counterintuitive because **selection among overloaded methods is static, while selection among overridden methods is dynamic**.”
  * “The correct version of an *overridden* method is chosen at runtime, based on the runtime type of the object on which the method is invoked.”
  * “As a reminder, a method is overridden when a subclass contains a method declaration with the same signature as a method declaration in an ancestor.”
  * “If an instance method is overridden in a subclass and this method is invoked on an instance of the subclass, the subclass’s *overriding method* executes, regardless of the compile-time type of the subclass instance.”


```java
public static String classify(Collection<?> c) {
    return c instanceof Set  ? "Set" :
           c instanceof List ? "List" : "Unknown Collection";
}
```

* “If the typical user of an API does not know which of several method overloadings will get invoked for a given set of parameters, use of the API is likely to result in errors.”
  * “Therefore you should **avoid confusing uses of overloading**.”
* “Exactly what constitutes a confusing use of overloading is open to some debate.”
  * **“A safe, conservative policy is never to export two overloadings with the same number of parameters.”**
  * “If a method uses varargs, a conservative policy is not to overload it at all, except as described in Item 53.”
  * “These restrictions are not terribly onerous because **you can always give methods different names instead of overloading them**.”
* “For constructors, you don’t have the option of using different names: multiple constructors for a class are *always* overloaded. ”
  * “You do, in many cases, have the option of exporting static factories instead of constructors (Item 1).”
  * “Also, with constructors you don’t have to worry about interactions between overloading and overriding, because constructors can’t be overridden.”
* “Prior to Java 5, all primitive types were radically different from all reference types, but this is not true in the presence of autoboxing, and it has caused real trouble.”


```java
public class SetList {
    public static void main(String[] args) {
        Set<Integer> set = new TreeSet<>();
        List<Integer> list = new ArrayList<>();
        
        for (int i = -3; i < 3; i++) {
            set.add(i);
            list.add(i);
        }
        for (int i = 0; i < 3; i++) {
            set.remove(i);
            list.remove(i);
        }
        System.out.println(set + " " + list);
    }
}
```

* “First, the program adds the integers from −3 to 2, inclusive, to a sorted set and a list. Then, it makes three identical calls to `remove` on the set and the list. If you’re like most people, you’d expect the program to remove the non-negative values (0, 1, and 2) from the set and the list and to print `[-3, -2, -1] [-3, -2, -1]`. In fact, the program removes the non-negative values from the set and the odd values from the list and prints `[-3, -2, -1] [-2, 0, 2]`.”
* “Here’s what’s happening: The call to `set.remove(i)` selects the overloading `remove(E)`, where `E` is the element type of the set (`Integer`), and autoboxes `i` from `int` to `Integer`. This is the behavior you’d expect, so the program ends up removing the positive values from the set. The call to `list.remove(i)`, on the other hand, selects the overloading `remove(int i)`, which removes the element at the specified position in the list. If you start with the list `[-3, -2, -1, 0, 1, 2]` and remove the zeroth element, then the first, and then the second, you’re left with `[-2, 0, 2]`, and the mystery is solved.”
* “To fix the problem, cast `list.remove`’s argument to `Integer`, forcing the correct overloading to be selected. Alternatively, you could invoke `Integer.valueOf` on i and pass the result to `list.remove`.”
* “The confusing behavior demonstrated by the previous example came about because the `List<E>` interface has two overloadings of the remove method: `remove(E)` and `remove(int)`.”
* “The addition of lambdas and method references in Java 8 further increased the potential for confusion in overloading.”

```java
new Thread(System.out::println).start();

ExecutorService exec = Executors.newCachedThreadPool();
exec.submit(System.out::println);
```

* “While the `Thread` constructor invocation and the `submit` method invocation look similar, the former compiles while the latter does not. The arguments are identical (`System.out::println`), and both the constructor and the method have an overloading that takes a `Runnable`.”
* “What’s going on here? The surprising answer is that the `submit` method has an overloading that takes a `Callable<T>`, while the `Thread` constructor does not. You might think that this shouldn’t make any difference because all overloadings of `println` return `void`, so the method reference couldn’t possibly be a `Callable`. This makes perfect sense, but it’s not the way the overload resolution algorithm works. Perhaps equally surprising is that the `submit` method invocation would be legal if the `println` method weren’t also overloaded.”
* “It is the combination of the overloading of the referenced method (`println`) and the invoked method (`submit`) that prevents the overload resolution algorithm from behaving as you’d expect.”
* “Therefore, **do not overload methods to take different functional interfaces in the same argument position**.”
* “Array types and class types other than `Object` are radically different. Also, array types and interface types other than `Serializable` and `Cloneable` are radically different. Two distinct classes are said to be *unrelated* if neither class is a descendant of the other [JLS, 5.5].”


```java
// Ensuring that 2 methods have identical behavior by forwarding
public boolean contentEquals(StringBuffer sb) {
    return contentEquals((CharSequence) sb);
}
```

* **“To summarize, just because you can overload methods doesn’t mean you should. It is generally best to refrain from overloading methods with multiple signatures that have the same number of parameters. In some cases, especially where constructors are involved, it may be impossible to follow this advice. In these cases, you should at least avoid situations where the same set of parameters can be passed to different overloadings by the addition of casts. If this cannot be avoided, for example, because you are retrofitting an existing class to implement a new interface, you should ensure that all overloadings behave identically when passed the same parameters. If you fail to do this, programmers will be hard pressed to make effective use of the overloaded method or constructor, and they won’t understand why it doesn’t work.”**

## Item 53: Use varargs judiciously

* “Varargs methods, formally known as *variable arity* methods [JLS, 8.4.1], accept zero or more arguments of a specified type. The varargs facility works by first creating an array whose size is the number of arguments passed at the call site, then putting the argument values into the array, and finally passing the array to the method.”
* “ometimes it’s appropriate to write a method that requires *one* or more arguments of some type, rather than *zero* or more.”


```java
// The WRONG way to use varargs to pass one or more arguments!
static int min(int... args) {
    if (args.length == 0)
        throw new IllegalArgumentException("Too few arguments");
    int min = args[0];
    for (int i = 1; i < args.length; i++)
        if (args[i] < min)
            min = args[i];
    return min;
}
```

```java
// The right way to use varargs to pass one or more arguments
static int min(int firstArg, int... remainingArgs) {
    int min = firstArg;
    for (int arg : remainingArgs)
        if (arg < min)
            min = arg;
    return min;
}
```

* “Exercise care when using varargs in performance-critical situations. Every invocation of a varargs method causes an array allocation and initialization.”
* “The static factories for `EnumSet` use this technique to reduce the cost of creating enum sets to a minimum. This was appropriate because it was critical that enum sets provide a performance-competitive replacement for bit fields (Item 36).”
* **“In summary, varargs are invaluable when you need to define methods with a variable number of arguments. Precede the varargs parameter with any required parameters, and be aware of the performance consequences of using varargs.”**

## Item 54: Return empty collections or arrays, not nulls

```java
// Returns null to indicate an empty collection. Don't do this!
private final List<Cheese> cheesesInStock = ...;

/**
 * @return a list containing all of the cheeses in the shop,
 *     or null if no cheeses are available for purchase.
 */
public List<Cheese> getCheeses() {
    return cheesesInStock.isEmpty() ? null
        : new ArrayList<>(cheesesInStock);
}
```

* “There is no reason to special-case the situation where no cheeses are available for purchase. Doing so requires extra code in the client to handle the possibly null return value, for example:”

```java
List<Cheese> cheeses = shop.getCheeses();
if (cheeses != null && cheeses.contains(Cheese.STILTON))
    System.out.println("Jolly good, just the thing.");
```

* “Here is the typical code to return a possibly empty collection.”


```java
//The right way to return a possibly empty collection
public List<Cheese> getCheeses() {
    return new ArrayList<>(cheesesInStock);
}
```

* “In the unlikely event that you have evidence suggesting that allocating empty collections is harming performance, you can avoid the allocations by returning the same *immutable* empty collection repeatedly, as immutable objects may be shared freely (Item 17).”


```java
// Optimization - avoids allocating empty collections
public List<Cheese> getCheeses() {
    return cheesesInStock.isEmpty() ? Collections.emptyList()
        : new ArrayList<>(cheesesInStock);
}
```

* “The situation for arrays is identical to that for collections. ”
  * “Never return null instead of a zero-length array.”
  * “Normally, you should simply return an array of the correct length, which may be zero.”
  * “Note that we’re passing a zero-length array into the toArray method to indicate the desired return type”

```java
//The right way to return a possibly empty array
public Cheese[] getCheeses() {
    return cheesesInStock.toArray(new Cheese[0]);
}
```

* “Do *not* preallocate the array passed to `toArray` in hopes of improving performance. ”


```java
// Don’t do this - preallocating the array harms performance!
return cheesesInStock.toArray(new Cheese[cheesesInStock.size()]);
```

* “In summary, **never return `null` in place of an empty array or collection**. It makes your API more difficult to use and more prone to error, and it has no performance advantages.”


## Item 55: Return optionals judiciously

* “Prior to Java 8, there were two approaches you could take when writing a method that was unable to return a value under certain circumstances. Either you could throw an exception, or you could return `null` (assuming the return type was an object reference type). Neither of these approaches is perfect.”
  * “Exceptions should be reserved for exceptional conditions (Item 69), and throwing an exception is expensive because the entire stack trace is captured when an exception is created.”
  * “Returning `null` doesn’t have these shortcomings, but it has its own. If a method returns `null`, clients must contain special-case code to deal with the possibility of a null return, unless the programmer can prove that a null return is impossible.”
* “In Java 8, there is a third approach to writing methods that may not be able to return a value. The `Optional<T>` class represents an immutable container that can hold either a single non-null `T` reference or nothing at all.”
  * “An optional that contains nothing is said to be *empty*. A value is said to be *present* in an optional that is not empty.”


```java
// Returns maximum value in collection as an Optional<E>
public static <E extends Comparable<E>>
        Optional<E> max(Collection<E> c) {
    if (c.isEmpty())
        return Optional.empty();
        
    E result = null;
    for (E e : c)
        if (result == null || e.compareTo(result) > 0)
            result = Objects.requireNonNull(e);

    return Optional.of(result);
}
```

* “**Never return a null value from an `Optional`-returning method**: it defeats the entire purpose of the facility.”
* “Many terminal operations on streams return optionals.”

```java
// Returns max val in collection as Optional<E> - uses stream
public static <E extends Comparable<E>>
        Optional<E> max(Collection<E> c) {
    return c.stream().max(Comparator.naturalOrder());
}
```

* “So how do you choose to return an optional instead of returning a `null` or throwing an exception? **Optionals are similar in spirit to checked exceptions** (Item 71), in that they *force* the user of an API to confront the fact that there may be no value returned.”
  * “Throwing an unchecked exception or returning a `null` allows the user to ignore this eventuality, with potentially dire consequences. However, throwing a checked exception requires additional boilerplate code in the client.”
* “If a method returns an optional, the client gets to choose what action to take if the method can’t return a value.”

```java
// Using an optional to provide a chosen default value
String lastWordInLexicon = max(words).orElse("No words...");
```

* “or you can throw any exception that is appropriate.”

```java
// Using an optional to throw a chosen exception
Toy myToy = max(toys).orElseThrow(TemperTantrumException::new);
```

* “If you can *prove* that an optional is nonempty, you can get the value from the optional without specifying an action to take if the optional is empty, but if you’re wrong, your code will throw a `NoSuchElementException`:”


```java
// Using optional when you know there’s a return value
Element lastNobleGas = max(Elements.NOBLE_GASES).get();
```

* “Occasionally you may be faced with a situation where it’s expensive to get the default value, and you want to avoid that cost unless it’s necessary. For these situations, `Optional` provides a method that takes a `Supplier<T>` and invokes it only when necessary. This method is called `orElseGet`, but perhaps it should have been called `orElseCompute` because it is closely related to the three Map methods whose names begin with `compute`.”
* “There are several `Optional` methods for dealing with more specialized use cases: `filter`, `map`, `flatMap`, and `ifPresent`. In Java 9, two more of these methods were added: `or` and `ifPresentOrElse`.”
* “In case none of these methods meets your needs, `Optional` provides the `isPresent()` method, which may be viewed as a safety valve. It returns `true` if the optional contains a value, `false` if it’s empty. You can use this method to perform any processing you like on an optional result, but make sure to use it wisely.”

```java
Optional<ProcessHandle> parentProcess = ph.parent();
System.out.println("Parent PID: " + (parentProcess.isPresent() ?
    String.valueOf(parentProcess.get().pid()) : "N/A"));
```

* “The code snippet above can be replaced by this one, which uses `Optional`’s map function:”


```java
System.out.println("Parent PID: " +
  ph.parent().map(h -> String.valueOf(h.pid())).orElse("N/A"));
```

* “When programming with streams, it is not uncommon to find yourself with a `Stream<Optional<T>>` and to require a `Stream<T>` containing all the elements in the nonempty optionals in order to proceed. If you’re using Java 8, here’s how to bridge the gap:”


```java
streamOfOptionals
    .filter(Optional::isPresent)
    .map(Optional::get)
```

* “Not all return types benefit from the optional treatment. **Container types, including collections, maps, streams, arrays, and optionals should not be wrapped in optionals.**”
* “So when should you declare a method to return `Optional<T>` rather than `T`? As a rule, you should **declare a method to return `Optional<T>` if it might not be able to return a result *and* clients will have to perform special processing if no result is returned**.”
* “Returning an optional that contains a boxed primitive type is prohibitively expensive compared to returning a primitive type because the optional has two levels of boxing instead of zero. ”
  * “Therefore, the library designers saw fit to provide analogues of `Optional<T>` for the primitive types `int`, `long`, and `double`. These optional types are `OptionalInt`, `OptionalLong`, and `OptionalDouble`. They contain most, but not all, of the methods on `Optional<T>`.”
  * “Therefore, **you should never return an optional of a boxed primitive type**, with the possible exception of the “minor primitive types,” `Boolean`, `Byte`, `Character`, `Short`, and `Float`.”
* “More generally, **it is almost never appropriate to use an optional as a key, value, or element in a collection or array**.”
  * “Is it ever appropriate to store an optional in an instance field? Often it’s a “bad smell”: it suggests that perhaps you should have a subclass containing the optional fields. But sometimes it may be justified.”
* **“In summary, if you find yourself writing a method that can’t always return a value and you believe it is important that users of the method consider this possibility every time they call it, then you should probably return an optional. You should, however, be aware that there are real performance consequences associated with returning optionals; for performance-critical methods, it may be better to return a `null` or throw an exception. Finally, you should rarely use an optional in any other capacity than as a return value.”**

## Item 56: Write doc comments for all exposed API elements

* **“If an API is to be usable, it must be documented.”**
  * “Traditionally, API documentation was generated manually, and keeping it in sync with code was a chore.”
  * “The Java programming environment eases this task with the Javadoc utility. *Javadoc* generates API documentation automatically from source code with specially formatted *documentation comments*, more commonly known as *doc comments*.”
* **“To document your API properly, you must precede every exported class, interface, constructor, method, and field declaration with a doc comment.”**
  * “If a class is serializable, you should also document its serialized form (Item 87).”
  * “Public classes should not use default constructors because there is no way to provide doc comments for them.”
* **“The doc comment for a method should describe succinctly the contract between the method and its client.”**
  * “With the exception of methods in classes designed for inheritance (Item 19), the contract should say *what* the method does rather than *how* it does its job.”
  * “The doc comment should enumerate all of the method’s *preconditions*, which are the things that have to be true in order for a client to invoke it, and its *postconditions*, which are the things that will be true after the invocation has completed successfully.”
    * “Typically, preconditions are described implicitly by the `@throws` tags for unchecked exceptions; each unchecked exception corresponds to a precondition violation. Also, preconditions can be specified along with the affected parameters in their `@param` tags.”
  * “In addition to preconditions and postconditions, **methods should document any *side effects***.”
    * “For example, if a method starts a background thread, the documentation should make note of it.”
  * “To describe a method’s contract fully, the doc comment should have an `@param` tag for every parameter, an `@return` tag unless the method has a void return type, and an `@throws` tag for every exception thrown by the method, whether checked or unchecked (Item 74).”
    * “By convention, the text following an `@param` tag or `@return` tag should be a noun phrase describing the value represented by the parameter or return value.”
    * “Rarely, arithmetic expressions are used in place of noun phrases; see `BigInteger` for examples. ”
    * “The text following an `@throws` tag should consist of the word “if,” followed by a clause describing the conditions under which the exception is thrown.”
    * “By convention, the phrase or clause following an `@param`, `@return`, or `@throws` tag is not terminated by a period.”
    * “Also notice the use of the Javadoc `{@code}` tag around the code fragment in the `@throws` clause. This tag serves two purposes: it causes the code fragment to be rendered in `code font`, and it suppresses processing of HTML markup and nested Javadoc tags in the code fragment. The latter property is what allows us to use the less-than sign (`<`) in the code fragment even though it’s an HTML metacharacter.”
    * “To include a multiline code example in a doc comment, use a Javadoc `{@code}` tag wrapped inside an HTML `<pre>` tag. In other words, precede the code example with the characters `<pre>{@code and follow it with }</pre>`. This preserves line breaks in the code, and eliminates the need to escape HTML metacharacters, but *not* the at sign (`@`), which must be escaped if the code sample uses annotations.”
    * “Finally, notice the use of the words “this list” in the doc comment. By convention, the word “this” refers to the object on which a method is invoked when it is used in the doc comment for an instance method.”


```java
/**
 * Returns the element at the specified position in this list.
 *
 * <p>This method is <i>not</i> guaranteed to run in constant
 * time. In some implementations it may run in time proportional
 * to the element position.
 *
 * @param  index index of element to return; must be
 *         non-negative and less than the size of this list
 * @return the element at the specified position in this list
 * @throws IndexOutOfBoundsException if the index is out of range
 *         ({@code index < 0 || index >= this.size()})
 */
E get(int index);
```

* “As mentioned in Item 15, when you design a class for inheritance, you must document its self-use patterns, so programmers know the semantics of overriding its methods. These self-use patterns should be documented using the `@implSpec` tag, added in Java 8.”
  * “Recall that ordinary doc comments describe the contract between a method and its client; `@implSpec` comments, by contrast, describe the contract between a method and its subclass, allowing subclasses to rely on implementation behavior if they inherit the method or call it via `super`.”


```java
/**
 * Returns true if this collection is empty.
 *
 * @implSpec
 * This implementation returns {@code this.size() == 0}.
 *
 * @return true if this collection is empty
 */
public boolean isEmpty() { ... }
```

* “Don’t forget that you must take special action to generate documentation that contains HTML metacharacters, such as the less-than sign (`<`), the greater-than sign (`>`), and the ampersand (`&`).”
  * “The best way to get these characters into documentation is to surround them with the `{@literal}` tag, which suppress processing of HTML markup and nested Javadoc tags. It is like the `{@code}` tag, except that it doesn’t render the text in code font.” 

* **“Doc comments should be readable both in the source code and in the generated documentation.”**
* “The first “sentence” of each doc comment (as defined below) becomes the *summary description* of the element to which the comment pertains.”
  * “To avoid confusion, **no two members or constructors in a class or interface should have the same summary description**.”
    * “Pay particular attention to overloadings, for which it is often natural to use the same first sentence (but unacceptable in doc comments).”
* “For methods and constructors, the summary description should be a verb phrase (including any object) describing the action performed by the method.”
  * “`ArrayList(int initialCapacity)`—Constructs an empty list with the specified initial capacity.”
  * “`Collection.size()`—Returns the number of elements in this collection.”
* “For classes, interfaces, and fields, the summary description should be a noun phrase describing the thing represented by an instance of the class or interface or by the field itself.”
  * “`Instant`—An instantaneous point on the time-line.”
  * “ `Math.PI`—The `double` value that is closer than any other to pi, the ratio of the circumference of a circle to its diameter.”
* “In Java 9, a client-side index was added to the HTML generated by Javadoc. This index, which eases the task of navigating large API documentation sets, takes the form of a search box in the upper-right corner of the page.”

```java
* This method complies with the {@index IEEE 754} standard.
```

* “Generics, enums, and annotations require special care in doc comments. **When documenting a generic type or method, be sure to document all type parameters**:”


```java
/**
 * An object that maps keys to values.  A map cannot contain
 * duplicate keys; each key can map to at most one value.
 *
 * (Remainder omitted)
 *
 * @param <K> the type of keys maintained by this map
 * @param <V> the type of mapped values
 */
public interface Map<K, V> { ... }
```

* “**When documenting an enum type, be sure to document the constants** as well as the type and any public methods.”


```java
/**
 * An instrument section of a symphony orchestra.
 */
public enum OrchestraSection {
    /** Woodwinds, such as flute, clarinet, and oboe. */
    WOODWIND,

    /** Brass instruments, such as french horn and trumpet. */
    BRASS,

    /** Percussion instruments, such as timpani and cymbals. */
    PERCUSSION,

    /** Stringed instruments, such as violin and cello. */
    STRING;
}
```

* “When documenting an annotation type, be sure to document any members as well as the type itself.”
  * “Document members with noun phrases, as if they were fields.”

  * “For the summary description of the type, use a verb phrase that says what it means when a program element has an annotation of this type:”


```java
/**
 * Indicates that the annotated method is a test method that
 * must throw the designated exception to pass.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface ExceptionTest {
     /**
      * The exception that the annotated test method must throw
      * in order to pass. (The test is permitted to throw any
      * subtype of the type described by this class object.)
      */
    Class<? extends Throwable> value();
}
```

* “Package-level doc comments should be placed in a file named `package-info.java`. In addition to these comments, `package-info.java` must contain a package declaration and may contain annotations on this declaration.”
* “Similarly, if you elect to use the module system (Item 15), module-level comments should be placed in the `module-info.java` file.”
* “Two aspects of APIs that are often neglected in documentation are thread-safety and serializability.”
  * “**Whether or not a class or static method is thread-safe, you should document its thread-safety level**, as described in Item 82.”
  * “If a class is serializable, you should document its serialized form, as described in Item 87.”
* “Javadoc has the ability to “inherit” method comments.”
  * “You can also inherit parts of doc comments from supertypes using the `{@inheritDoc}` tag. This means, among other things, that classes can reuse doc comments from interfaces they implement, rather than copying these comments.”
* “For complex APIs consisting of multiple interrelated classes, it is often necessary to supplement the documentation comments with an external document describing the overall architecture of the API. If such a document exists, the relevant class or package documentation comments should include a link to it.”
* “If you adhere to the guidelines in this item, the generated documentation should provide a clear description of your API. The only way to know for sure, however, is to **read the web pages generated by the Javadoc utility**.”
* “To summarize, documentation comments are the best, most effective way to document your API. Their use should be considered mandatory for all exported API elements. Adopt a consistent style that adheres to standard conventions. Remember that arbitrary HTML is permissible in documentation comments and that HTML metacharacters must be escaped.”
