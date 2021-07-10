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


## Item 58: Prefer for-each loops to traditional `for` loops

```java
// Not the best way to iterate over a collection!
for (Iterator<Element> i = c.iterator(); i.hasNext(); ) {
    Element e = i.next();
    ... // Do something with e
}
```

```java
// Not the best way to iterate over an array!
for (int i = 0; i < a.length; i++) {
    ... // Do something with a[i]
}
```

* “These idioms are better than `while` loops (Item 57), but they aren’t perfect. The iterator and the index variables are both just clutter—all you need are the elements.”
* “Furthermore, they represent opportunities for error. The iterator occurs three times in each loop and the index variable four, which gives you many chances to use the wrong variable.”


```java
// The preferred idiom for iterating over collections and arrays
for (Element e : elements) {
    ... // Do something with e
}
```

* “When you see the colon (`:`), read it as “in.” Thus, the loop above reads as “for each element `e` in `elements`.”
* “There is no performance penalty for using for-each loops, even for arrays.”


```java
// Can you spot the bug?
enum Suit { CLUB, DIAMOND, HEART, SPADE }
enum Rank { ACE, DEUCE, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT,
            NINE, TEN, JACK, QUEEN, KING }
...
static Collection<Suit> suits = Arrays.asList(Suit.values());
static Collection<Rank> ranks = Arrays.asList(Rank.values());

List<Card> deck = new ArrayList<>();
for (Iterator<Suit> i = suits.iterator(); i.hasNext(); )
    for (Iterator<Rank> j = ranks.iterator(); j.hasNext(); )
        deck.add(new Card(i.next(), j.next()));
```

* “The problem is that the `next` method is called too many times on the iterator for the outer collection (`suits`). It should be called from the outer loop so that it is called once per suit, but instead it is called from the inner loop, so it is called once per card. After you run out of suits, the loop throws a `NoSuchElementException`.”
* “If instead you use a nested for-each loop, the problem simply disappears.”

```java
// Preferred idiom for nested iteration on collections and arrays
for (Suit suit : suits)
    for (Rank rank : ranks)
        deck.add(new Card(suit, rank));
```

* “Unfortunately, there are three common situations where you *can’t* use for-each:”
  * “**Destructive filtering**—If you need to traverse a collection removing selected elements, then you need to use an explicit iterator so that you can call its remove method. You can often avoid explicit traversal by using `Collection`’s `removeIf` method, added in Java 8.”
  * “**Transforming**—If you need to traverse a list or array and replace some or all of the values of its elements, then you need the list iterator or array index in order to replace the value of an element.”
  * “**Parallel iteration**—If you need to traverse multiple collections in parallel, then you need explicit control over the iterator or index variable so that all iterators or index variables can be advanced in lockstep.”
* “Not only does the for-each loop let you iterate over collections and arrays, it lets you iterate over any object that implements the `Iterable` interface, which consists of a single method. ”

```java
public interface Iterable<E> {
    // Returns an iterator over the elements in this iterable
    Iterator<E> iterator();
}
```

* **“In summary, the for-each loop provides compelling advantages over the traditional `for` loop in clarity, flexibility, and bug prevention, with no performance penalty. Use for-each loops in preference to `for` loops wherever you can.”**

## Item 59: Know and use the libraries

* “Suppose you want to generate random integers between zero and some upper bound. Faced with this common task, many programmers would write a little method that looks something like this:”

```java
// Common but deeply flawed!
static Random rnd = new Random();

static int random(int n) {
    return Math.abs(rnd.nextInt()) % n;
}
```

* “This method may look good, but it has three flaws.”
  * “The first is that if `n` is a small power of two, the sequence of random numbers will repeat itself after a fairly short period.”
  * “The second flaw is that if `n` is not a power of two, some numbers will, on average, be returned more frequently than others. If n is large, this effect can be quite pronounced.”
  * “The third flaw in the `random` method is that it can, on rare occasions, fail catastrophically, returning a number outside the specified range.”
* “To write a version of the `random` method that corrects these flaws, you’d have to know a fair amount about pseudorandom number generators, number theory, and two’s complement arithmetic. Luckily, you don’t have to do this—it’s been done for you. It’s called `Random.nextInt(int)`.”
* **“By using a standard library, you take advantage of the knowledge of the experts who wrote it and the experience of those who used it before you.”**
* “As of Java 7, you should no longer use `Random`.”
  * “For most uses, **the random number generator of choice is now `ThreadLocalRandom`**.”
  * “It produces higher quality random numbers, and it’s very fast.”

  * “For fork join pools and parallel streams, use `SplittableRandom`.”
* “A second advantage of using the libraries is that you don’t have to waste your time writing ad hoc solutions to problems that are only marginally related to your work.”
* “A third advantage of using standard libraries is that their performance tends to improve over time, with no effort on your part.”
* “A fourth advantage of using libraries is that they tend to gain functionality over time.”
  * “If a library is missing something, the developer community will make it known, and the missing functionality may get added in a subsequent release.”

* “A final advantage of using the standard libraries is that you place your code in the mainstream.”
  * “Such code is more easily readable, maintainable, and reusable by the multitude of developers.”
* **“Numerous features are added to the libraries in every major release, and it pays to keep abreast of these additions.”**
* “The libraries are too big to study all the documentation [Java9-api], but **every programmer should be familiar with the basics of `java.lang`, `java.util`, and `java.io`, and their subpackages**.”
* “Several libraries bear special mention. The collections framework and the streams library (Items 45–48) should be part of every programmer’s basic toolkit, as should parts of the concurrency utilities in `java.util.concurrent`. This package contains both high-level utilities to simplify the task of multithreaded programming and low-level primitives to allow experts to write their own higher-level concurrent abstractions.”
* **“To summarize, don’t reinvent the wheel.”**

## Item 60: Avoid `float` and `double` if exact answers are required

* “The `float` and `double` types are designed primarily for scientific and engineering calculations. They perform *binary floating-point arithmetic*, which was carefully designed to furnish accurate approximations quickly over a broad range of magnitudes. They do not, however, provide exact results and should not be used where exact results are required.”
* “**The `float` and `double` types are particularly ill-suited for monetary calculations** because it is impossible to represent 0.1 (or any other negative power of ten) as a float or double exactly.”
  * “The right way to solve this problem is to **use `BigDecimal`, `int`, or `long` for monetary calculations**.”
* “Here’s a straightforward transformation of the previous program to use the `BigDecimal` type in place of `double`. Note that `BigDecimal`’s `String` constructor is used rather than its `double` constructor. This is required in order to avoid introducing inaccurate values into the computation [Bloch05, Puzzle 2]:”


```java
public static void main(String[] args) {
    final BigDecimal TEN_CENTS = new BigDecimal(".10");
    int itemsBought = 0;
    BigDecimal funds = new BigDecimal("1.00");
    for (BigDecimal price = TEN_CENTS;
            funds.compareTo(price) >= 0;
            price = price.add(TEN_CENTS)) {
        funds = funds.subtract(price);
        itemsBought++;
    }
    System.out.println(itemsBought + " items bought.");
    System.out.println("Money left over: $" + funds);
}
```

* “There are, however, two disadvantages to using `BigDecimal`: it’s a lot less convenient than using a primitive arithmetic type, and it’s a lot slower. The latter disadvantage is irrelevant if you’re solving a single short problem, but the former may annoy you.”
* **“In summary, don’t use float or double for any calculations that require an exact answer.”**
  * “Use `BigDecimal` if you want the system to keep track of the decimal point and you don’t mind the inconvenience and cost of not using a primitive type.”
  * “Using `BigDecimal` has the added advantage that it gives you full control over rounding, letting you select from eight rounding modes whenever an operation that entails rounding is performed. This comes in handy if you’re performing business calculations with legally mandated rounding behavior.”
  * “Using `BigDecimal` has the added advantage that it gives you full control over rounding, letting you select from eight rounding modes whenever an operation that entails rounding is performed. This comes in handy if you’re performing business calculations with legally mandated rounding behavior.”
  * “If performance is of the essence, you don’t mind keeping track of the decimal point yourself, and the quantities aren’t too big, use `int` or `long`.”
  * “If the quantities don’t exceed nine decimal digits, you can use `int`; if they don’t exceed eighteen digits, you can use `long`. If the quantities might exceed eighteen digits, use `BigDecimal`.”


## Item 61: Prefer primitive types to boxed primitives

* “Java has a two-part type system, consisting of *primitives*, such as `int`, `double`, and `boolean`, and *reference types*, such as `String` and `List`. Every primitive type has a corresponding reference type, called a boxed primitive. The boxed primitives corresponding to `int`, `double`, and `boolean` are `Integer`, `Double`, and `Boolean`.”
* “There are three major differences between primitives and boxed primitives.”
  * “First, primitives have only their values, whereas boxed primitives have identities distinct from their values.”
  * “Second, primitive types have only fully functional values, whereas each boxed primitive type has one nonfunctional value, which is `null`, in addition to all the functional values of the corresponding primitive type.”
  * “Last, primitives are more time- and space-efficient than boxed primitives.”
* **“Applying the `==` operator to boxed primitives is almost always wrong.”**
* “In practice, if you need a comparator to describe a type’s natural order, you should simply call `Comparator.naturalOrder()`, and if you write a comparator yourself, you should use the comparator construction methods, or the static compare methods on primitive types (Item 14).”


```java
Comparator<Integer> naturalOrder = (iBoxed, jBoxed) -> {
    int i = iBoxed, j = jBoxed; // Auto-unboxing
    return i < j ? -1 : (i == j ? 0 : 1);
};
```

* “In nearly every case **when you mix primitives and boxed primitives in an operation, the boxed primitive is auto-unboxed**.”
  * “If a null object reference is auto-unboxed, you get a `NullPointerException`.”
* “So when should you use boxed primitives? They have several legitimate uses.”
  * “The first is as elements, keys, and values in collections. You can’t put primitives in collections, so you’re forced to use boxed primitives.”
  * “You must use boxed primitives as type parameters in parameterized types and methods (Chapter 5), because the language does not permit you to use primitives.”

  * “Finally, you must use boxed primitives when making reflective method invocations (Item 65).”
* **“In summary, use primitives in preference to boxed primitives whenever you have the choice. Primitive types are simpler and faster. If you must use boxed primitives, be careful! Autoboxing reduces the verbosity, but not the danger, of using boxed primitives. When your program compares two boxed primitives with the `==` operator, it does an identity comparison, which is almost certainly not what you want. When your program does mixed-type computations involving boxed and unboxed primitives, it does unboxing, and when your program does unboxing, it can throw a `NullPointerException`. Finally, when your program boxes primitive values, it can result in costly and unnecessary object creations.”**

## Item 62: Avoid strings where other types are more appropriate

* **“Strings are poor substitutes for other value types.”**
* **“Strings are poor substitutes for enum types.”**
* **“Strings are poor substitutes for aggregate types.”**
  * “A better approach is simply to write a class to represent the aggregate, often a private static member class (Item 24).”
* **“Strings are poor substitutes for capabilities.”**

```java
// Broken - inappropriate use of string as capability!
public class ThreadLocal {
    private ThreadLocal() { } // Noninstantiable

    // Sets the current thread's value for the named variable.
    public static void set(String key, Object value);

    // Returns the current thread's value for the named variable.
    public static Object get(String key);
}
```

* “The problem with this approach is that the string keys represent a shared global namespace for thread-local variables. In order for the approach to work, the client-provided string keys have to be unique: if two clients independently decide to use the same name for their thread-local variable, they unintentionally share a single variable, which will generally cause both clients to fail. Also, the security is poor. A malicious client could intentionally use the same string key as another client to gain illicit access to the other client’s data.”
* “This API can be fixed by replacing the string with an unforgeable key (sometimes called a *capability*):”


```java
public class ThreadLocal {
    private ThreadLocal() { }    // Noninstantiable

    public static class Key {    // (Capability)
        Key() { }
    }

    // Generates a unique, unforgeable key
    public static Key getKey() {
        return new Key();
    }

    public static void set(Key key, Object value);
    public static Object get(Key key);
}
```

* “While this solves both of the problems with the string-based API, you can do much better. You don’t really need the static methods anymore. They can instead become instance methods on the key, at which point the key is no longer a key for a thread-local variable: it is a thread-local variable.”

```java
public final class ThreadLocal {
    public ThreadLocal();
    public void set(Object value);
    public Object get();
}
```

* “This API isn’t typesafe, because you have to cast the value from `Object` to its actual type when you retrieve it from a thread-local variable. It is impossible to make the original `String`-based API typesafe and difficult to make the `Key`-based API typesafe, but it is a simple matter to make this API typesafe by making `ThreadLocal` a parameterized class (Item 29):”

```java
public final class ThreadLocal<T> {
    public ThreadLocal();
    public void set(T value);
    public T get();
}
```

* **“To summarize, avoid the natural tendency to represent objects as strings when better data types exist or can be written. Used inappropriately, strings are more cumbersome, less flexible, slower, and more error-prone than other types. Types for which strings are commonly misused include primitive types, enums, and aggregate types.”**

## Item 63: Beware the performance of string concatenation

* “The string concatenation operator (`+`) is a convenient way to combine a few strings into one. It is fine for generating a single line of output or constructing the string representation of a small, fixed-size object, but it does not scale.”
  * “Using the string concatenation operator repeatedly to concatenate n strings requires time quadratic in n.”
  * “This is an unfortunate consequence of the fact that strings are *immutable* (Item 17). When two strings are concatenated, the contents of both are copied.”
* “**To achieve acceptable performance, use a `StringBuilder` in place of a `String`** to store the statement under construction:”

```java
public String statement() {
    StringBuilder b = new StringBuilder(numItems() * LINE_WIDTH);
    for (int i = 0; i < numItems(); i++)
        b.append(lineForItem(i));
    return b.toString();
}
```

* **“The moral is simple: Don’t use the string concatenation operator to combine more than a few strings unless performance is irrelevant. Use `StringBuilder`’s `append` method instead. Alternatively, use a character array, or process the strings one at a time instead of combining them.”**

## Item 64: Refer to objects by their interfaces

* “You should favor the use of interfaces over classes to refer to objects.”
  * **“If appropriate interface types exist, then parameters, return values, variables, and fields should all be declared using interface types.”**
  * “The only time you really need to refer to an object’s class is when you’re creating it with a constructor.”

```java
// Good - uses interface as type
Set<Son> sonSet = new LinkedHashSet<>();

// Bad - uses class as type!
LinkedHashSet<Son> sonSet = new LinkedHashSet<>();
```

* **“If you get into the habit of using interfaces as types, your program will be much more flexible.”**
  * “If you decide that you want to switch implementations, all you have to do is change the class name in the constructor (or use a different static factory). ”
* **“It is entirely appropriate to refer to an object by a class rather than an interface if no appropriate interface exists.”**
  * “Value classes are rarely written with multiple implementations in mind. They are often final and rarely have corresponding interfaces.”
  * “A second case in which there is no appropriate interface type is that of objects belonging to a framework whose fundamental types are classes rather than interfaces.”
  * “A final case in which there is no appropriate interface type is that of classes that implement an interface but also provide extra methods not found in the interface”

* **“If there is no appropriate interface, just use the least specific class in the class hierarchy that provides the required functionality.”**

## Item 65: Prefer interfaces to reflection

* “The *core reflection facility*, `java.lang.reflect`, offers programmatic access to arbitrary classes. Given a `Class` object, you can obtain `Constructor`, `Method`, and `Field` instances representing the constructors, methods, and fields of the class represented by the `Class` instance. These objects provide programmatic access to the class’s member names, field types, method signatures, and so on.”
* “Moreover, `Constructor`, `Method`, and `Field` instances let you manipulate their underlying counterparts *reflectively*: you can construct instances, invoke methods, and access fields of the underlying class by invoking methods on the `Constructor`, `Method`, and `Field` instances.”
  * “For example, `Method.invoke` lets you invoke any method on any object of any class (subject to the usual security constraints).”
  * “Reflection allows one class to use another, even if the latter class did not exist when the former was compiled.”
* “This power, however, comes at a price:”
  * “**You lose all the benefits of compile-time type checking**, including exception checking.”
    * “If a program attempts to invoke a nonexistent or inaccessible method reflectively, it will fail at runtime unless you’ve taken special precautions.”
  * **“The code required to perform reflective access is clumsy and verbose.”**
  * **“Performance suffers.”**
* “There are a few sophisticated applications that require reflection. Examples include code analysis tools and dependency injection frameworks. Even such tools have been moving away from reflection of late, as its disadvantages become clearer.”
  * **“If you have any doubts as to whether your application requires reflection, it probably doesn’t.”**
* **“You can obtain many of the benefits of reflection while incurring few of its costs by using it only in a very limited form.”**
  * “For many programs that must use a class that is unavailable at compile time, there exists at compile time an appropriate interface or superclass by which to refer to the class (Item 64).”
  * “If this is the case, you can **create instances reflectively and access them normally via their interface or superclass**.”


```java
// Reflective instantiation with interface access
public static void main(String[] args) {
    // Translate the class name into a Class object
    Class<? extends Set<String>> cl = null;
    try {
        cl = (Class<? extends Set<String>>)  // Unchecked cast!
                Class.forName(args[0]);
    } catch (ClassNotFoundException e) {
        fatalError("Class not found.");
    }
    // Get the constructor
    Constructor<? extends Set<String>> cons = null;
    try {
        cons = cl.getDeclaredConstructor();
    } catch (NoSuchMethodException e) {
        fatalError("No parameterless constructor");
    }
    // Instantiate the set
    Set<String> s = null;
    try {
        s = cons.newInstance();
    } catch (IllegalAccessException e) {
        fatalError("Constructor not accessible");
    } catch (InstantiationException e) {
        fatalError("Class not instantiable.");
    } catch (InvocationTargetException e) {
        fatalError("Constructor threw " + e.getCause());
    } catch (ClassCastException e) {
        fatalError("Class doesn't implement Set");
    }
    // Exercise the set
    s.addAll(Arrays.asList(args).subList(1, args.length));
    System.out.println(s);
}
private static void fatalError(String msg) {
    System.err.println(msg);
    System.exit(1);
}
```

* “The toy program could easily be turned into a generic set tester that validates the specified `Set` implementation by aggressively manipulating one or more instances and checking that they obey the `Set` contract.”
* “Similarly, it could be turned into a generic set performance analysis tool.”
* “In fact, this technique is sufficiently powerful to implement a full-blown *service provider framework* (Item 1).”
* “This example demonstrates two disadvantages of reflection. ”
  * “First, the example can generate six different exceptions at runtime, all of which would have been compile-time errors if reflective instantiation were not used.”
  * “The second disadvantage is that it takes twenty-five lines of tedious code to generate an instance of the class from its name, whereas a constructor invocation would fit neatly on a single line.”
  * “The length of the program could be reduced by catching `ReflectiveOperationException`, a superclass of the various reflective exceptions that was introduced in Java 7.”
  * “Once instantiated, the set is indistinguishable from any other `Set` instance.”
* “If you compile this program, you’ll get an unchecked cast warning. This warning is legitimate, in that the cast to `Class<? extends Set<String>>` will succeed even if the named class is not a `Set` implementation, in which case the program with throw a `ClassCastException` when it instantiates the class. ”
* “A legitimate, if rare, use of reflection is to manage a class’s dependencies on other classes, methods, or fields that may be absent at runtime.”
  * “This can be useful if you are writing a package that must run against multiple versions of some other package.”
  * “The technique is to compile your package against the minimal environment required to support it, typically the oldest version, and to access any newer classes or methods reflectively. ”
  * “To make this work, you have to take appropriate action if a newer class or method that you are attempting to access does not exist at runtime. Appropriate action might consist of using some alternate means to accomplish the same goal or operating with reduced functionality.”
* **“In summary, reflection is a powerful facility that is required for certain sophisticated system programming tasks, but it has many disadvantages. If you are writing a program that has to work with classes unknown at compile time, you should, if at all possible, use reflection only to instantiate objects, and access the objects using some interface or superclass that is known at compile time.”**

## Item 66: Use native methods judiciously

* “The Java Native Interface (JNI) allows Java programs to call *native methods*, which are methods written in *native programming languages* such as C or C++.”
* “Historically, native methods have had three main uses. They provide access to platform-specific facilities such as registries. They provide access to existing libraries of native code, including legacy libraries that provide access to legacy data. Finally, native methods are used to write performance-critical parts of applications in native languages for improved performance.”
* “It is legitimate to use native methods to access platform-specific facilities, but it is seldom necessary: as the Java platform matured, it provided access to many features previously found only in host platforms.”
* **“It is rarely advisable to use native methods for improved performance.”**
* **“In summary, think twice before using native methods. It is rare that you need to use them for improved performance. If you must use native methods to access low-level resources or native libraries, use as little native code as possible and test it thoroughly. A single bug in the native code can corrupt your entire application.”**

## Item 67: Optimize judiciously

* “There are three aphorisms concerning optimization that everyone should know:”
  * “More computing sins are committed in the name of efficiency (without necessarily achieving it) than for any other single reason—including blind stupidity.
    —William A. Wulf [Wulf72]”
  * “We should forget about small efficiencies, say about 97% of the time: premature optimization is the root of all evil.
    —Donald E. Knuth [Knuth74]”
  * “We follow two rules in the matter of optimization:
    Rule 1. Don’t do it.
    Rule 2 (for experts only). Don’t do it yet—that is, not until you have a perfectly clear and unoptimized solution.
    —M. A. Jackson [Jackson75]”
* **“Strive to write good programs rather than fast ones.”**
  * “Don’t sacrifice sound architectural principles for performance.”
  * “Good programs embody the principle of *information hiding*: where possible, they localize design decisions within individual components, so individual decisions can be changed without affecting the remainder of the system (Item 15).”
  * “You must think about performance during the design process.”

* **“Strive to avoid design decisions that limit performance.”**
  * “The components of a design that are most difficult to change after the fact are those specifying interactions between components and with the outside world.”
    * “Chief among these design components are APIs, wire-level protocols, and persistent data formats.”
* **“Consider the performance consequences of your API design decisions.”**
  * “Making a public type mutable may require a lot of needless defensive copying (Item 50).”
  * “Similarly, using inheritance in a public class where composition would have been appropriate ties the class forever to its superclass, which can place artificial limits on the performance of the subclass (Item 18).”
  * “As a final example, using an implementation type rather than an interface in an API ties you to a specific implementation, even though faster implementations may be written in the future (Item 64).”
* “Luckily, it is generally the case that good API design is consistent with good performance. **It is a very bad idea to warp an API to achieve good performance.**”
* **“Measure performance before and after each attempted optimization.”**
  * “Often, attempted optimizations have no measurable effect on performance; sometimes, they make it worse.”
  * “Profiling tools can help you decide where to focus your optimization efforts”
  * “Another tool that deserves special mention is jmh, which is not a profiler but a *microbenchmarking framework* that provides unparalleled visibility into the detailed performance of Java code [JMH].”
* “Java has a weaker *performance model*: The relative cost of the various primitive operations is less well defined.”
  * “The “abstraction gap” between what the programmer writes and what the CPU executes is greater, which makes it even more difficult to reliably predict the performance consequences of optimizations.”
* “Not only is Java’s performance model ill-defined, but it varies from implementation to implementation, from release to release, and from processor to processor.”

* **“To summarize, do not strive to write fast programs—strive to write good ones; speed will follow. But do think about performance while you’re designing systems, especially while you’re designing APIs, wire-level protocols, and persistent data formats. When you’ve finished building the system, measure its performance. If it’s fast enough, you’re done. If not, locate the source of the problem with the aid of a profiler and go to work optimizing the relevant parts of the system. The first step is to examine your choice of algorithms: no amount of low-level optimization can make up for a poor choice of algorithm. Repeat this process as necessary, measuring the performance after every change, until you’re satisfied.”**

## Item 68: Adhere to generally accepted naming conventions

* “The Java platform has a well-established set of *naming conventions*, many of which are contained in *The Java Language Specification* [JLS, 6.1].”
* “Loosely speaking, naming conventions fall into two categories: typographical and grammatical.”
* “There are only a handful of typographical naming conventions, covering packages, classes, interfaces, methods, fields, and type variables.”
  * “Package and module names should be hierarchical with the components separated by periods.”
    * “Components should consist of lowercase alphabetic characters and, rarely, digits.”
    * “The name of any package that will be used outside your organization should begin with your organization’s Internet domain name with the components reversed, for example, `edu.cmu`, `com.google`, `org.eff`.”
    * “Components should be short, generally eight or fewer characters.”
      * “Meaningful abbreviations are encouraged, for example, `util` rather than `utilities`. Acronyms are acceptable, for example, `awt`. Components should generally consist of a single word or abbreviation.”
    * “Additional components are appropriate for large facilities whose size demands that they be broken up into an informal hierarchy.”
      * “For example, the `javax.util` package has a rich hierarchy of packages with names such as `java.util.concurrent.atomic`. Such packages are known as *subpackages*, although there is almost no linguistic support for package hierarchies.”
  * “Class and interface names, including enum and annotation type names, should consist of one or more words, with the first letter of each word capitalized, for example, `List` or `FutureTask`.”
    * “Abbreviations are to be avoided, except for acronyms and certain common abbreviations like `max` and `min`.”
  * “Method and field names follow the same typographical conventions as class and interface names, except that the first letter of a method or field name should be lowercase, for example, `remove` or `ensureCapacity`.”
    * “If an acronym occurs as the first word of a method or field name, it should be lowercase.”
  * “The sole exception to the previous rule concerns “constant fields,” whose names should consist of one or more uppercase words separated by the underscore character, for example, `VALUES` or `NEGATIVE_INFINITY`.”
    * “A constant field is a static final field whose value is immutable.”
    * “For example, enum constants are constant fields. ”
  * “Local variable names have similar typographical naming conventions to member names, except that abbreviations are permitted, as are individual characters and short sequences of characters whose meaning depends on the context in which they occur, for example, `i`, `denom`, `houseNum`.”
  * “Input parameters are a special kind of local variable. They should be named much more carefully than ordinary local variables, as their names are an integral part of their method’s documentation.”
  * “Type parameter names usually consist of a single letter.”
    * “Most commonly it is one of these five: `T` for an arbitrary type, `E` for the element type of a collection, `K` and `V` for the key and value types of a map, and `X` for an exception.”
    * “The return type of a function is usually `R`.”
    * “A sequence of arbitrary types can be `T`, `U`, `V` or `T1`, `T2`, `T3`.”


| Identifier Type    | Examples                                              |
| ------------------ | ----------------------------------------------------- |
| Package or Module  | `org.junit.jupiter.api`, `com.google.common.collect`  |
| Class of Interface | `Stream`, `FutureTask`, `LinkedHashMap`, `HttpClient` |
| Method or Field    | `remove`, `groupingBy`, `getCrc`                      |
| Constant Field     | `MIN_VALUE`, `NEGATIVE_INFINITY`                      |
| Local Variable     | `i`, `denom`, `houseNum`                              |
| Type Parameter     | `T`, `E`, `K`, `V`, `X`, `R`, `U`, `V`, `T1`, `T2`    |

* “Grammatical naming conventions are more flexible and more controversial than typographical conventions.”
  * “Instantiable classes, including enum types, are generally named with a singular noun or noun phrase, such as `Thread`, `PriorityQueue`, or `ChessPiece`.”
  * “Non-instantiable utility classes (Item 4) are often named with a plural noun, such as `Collectors` or `Collections`.”
  * “Interfaces are named like classes, for example, `Collection` or `Comparator`, or with an adjective ending in able or ible, for example, `Runnable`, `Iterable`, or `Accessible`.”
  * “Because annotation types have so many uses, no part of speech predominates. Nouns, verbs, prepositions, and adjectives are all common, for example, `BindingAnnotation`, `Inject`, `ImplementedBy`, or `Singleton`.”
  * “Methods that perform some action are generally named with a verb or verb phrase (including object), for example, `append` or `drawImage`.”
  * “Methods that return a `boolean` value usually have names that begin with the word `is` or, less commonly, `has`, followed by a noun, noun phrase, or any word or phrase that functions as an adjective, for example, `isDigit`, `isProbablePrime`, `isEmpty`, `isEnabled`, or `hasSiblings`.”
  * “Methods that return a non-`boolean` function or attribute of the object on which they’re invoked are usually named with a noun, a noun phrase, or a verb phrase beginning with the verb `get`, for example, `size`, `hashCode`, or `getTime`.”
  * “There is also a strong precedent for following this naming convention if a class contains both a setter and a getter for the same attribute. In this case, the two methods are typically named `getAttribute` and `setAttribute`.”
  * “A few method names deserve special mention.”
    * “Instance methods that convert the type of an object, returning an independent object of a different type, are often called `toType`, for example, `toString` or `toArray`.”
    * “Methods that return a view (Item 6) whose type differs from that of the receiving object are often called `asType`, for example, `asList`.”
    * “Methods that return a primitive with the same value as the object on which they’re invoked are often called `typeValue`, for example, `intValue`. ”
    * “Common names for static factories include from, `of`, `valueOf`, `instance`, `getInstance`, `newInstance`, `getType`, and `newType` (Item 1, page 9).”
  * “Grammatical conventions for field names are less well established and less important than those for class, interface, and method names because well-designed APIs contain few if any exposed fields.”
    * “Fields of type boolean are often named like boolean accessor methods with the initial is omitted, for example, `initialized`, `composite`.”
    * “Fields of other types are usually named with nouns or noun phrases, such as `height`, `digits`, or `bodyStyle`.”
