# Chapter 2. Creating and Destroying Objects

## Item 1: Consider static factory methods instead of constructors

* “A class can provide a public static factory method, which is simply a static method that returns an instance of the class.”
* “A class can provide its clients with static factory methods instead of, or in addition to, public constructors.”
* **“One advantage of static factory methods is that, unlike constructors, they have names.”**
  * “A static factory with a well-chosen name is easier to use and the resulting client code easier to read.”
* **“A second advantage of static factory methods is that, unlike constructors, they are not required to create a new object each time they’re invoked.”**
  * “This allows immutable classes (Item 17) to use preconstructed instances, or to cache instances as they’re constructed, and dispense them repeatedly to avoid creating unnecessary duplicate objects.”
* **“A third advantage of static factory methods is that, unlike constructors, they can return an object of any subtype of their return type.”**
  * “One application of this flexibility is that an API can return objects without making their classes public. Hiding implementation classes in this fashion leads to a very compact API.”
  * “Prior to Java 8, interfaces couldn’t have static methods. By convention, static factory methods for an interface named Type were put in a noninstantiable companion class (Item 4) named Types.”

* **“A fourth advantage of static factories is that the class of the returned object can vary from call to call as a function of the input parameters.”**
  * “The existence of implementation classes can be invisible to clients.”
  * “Clients neither know nor care about the class of the object they get back from the factory.”
* **“A fifth advantage of static factories is that the class of the returned object need not exist when the class containing the method is written.”**
* **“The main limitation of providing only static factory methods is that classes without public or protected constructors cannot be subclassed.”**
* **“A second shortcoming of static factory methods is that they are hard for programmers to find.”**
* **Common names for static factory methods:**
  * **from** - “A *type-conversion* method that takes a single parameter and returns a corresponding instance of this type.”
  * **of** - “An *aggregation* method that takes multiple parameters and returns an instance of this type that incorporates them.”
  * **valueOf** - “A more verbose alternative to `from` and `of`.”
  * **instance** or **getInstance** - “Returns an instance that is described by its parameters (if any) but cannot be said to have the same value.”
  * **create** or **newInstance** - “Like `instance` or `getInstance`, except that the method guarantees that each call returns a new instance.”
  * **get*Type*** - “Like `getInstance`, but used if the factory method is in a different class.”
  * **new*Type*** - “Like `newInstance`, but used if the factory method is in a different class.”
  * ***type*** - “A concise alternative to `getType` and `newType`.”

## Item 2: Consider a builder when faced with many constructor parameters

* “Static factories and constructors share a limitation: they do not scale well to large numbers of optional parameters.”
* Telescoping Constructor Pattern
  * "You provide a constructor with only the required parameters, another with a single optional parameter, a third with two optional parameters, and so on, culminating in a constructor with all the optional parameters.”
  * “The telescoping constructor pattern works, but it is hard to write client code when there are many parameters, and harder still to read it.”
* JavaBeans Pattern
  * “You call a parameterless constructor to create the object and then call setter methods to set each required parameter and each optional parameter of interest.”
  * “Because construction is split across multiple calls, a JavaBean may be in an inconsistent state partway through its construction.”
  * “A related disadvantage is that the JavaBeans pattern precludes the possibility of making a class immutable (Item 17) and requires added effort on the part of the programmer to ensure thread safety.”
* **Builder Pattern**
  * “Instead of making the desired object directly, the client calls a constructor (or static factory) with all of the required parameters and gets a builder object. Then the client calls setter-like methods on the builder object to set each optional parameter of interest. Finally, the client calls a parameterless build method to generate the object, which is typically immutable."
  * “The builder is typically a static member class (Item 24) of the class it builds. ”
  * “The builder’s setter methods return the builder itself so that invocations can be chained, resulting in a *fluent* API.”
  * **“The Builder pattern is well suited to class hierarchies.”**
  * “A minor advantage of builders over constructors is that builders can have multiple varargs parameters because each parameter is specified in its own method.”
  * “While the cost of creating this builder is unlikely to be noticeable in practice, it could be a problem in performance-critical situations.”
  * “Builder pattern is more verbose than the telescoping constructor pattern, so it should be used only if there are enough parameters to make it worthwhile, say four or more.”
  * “It’s often better to start with a builder in the first place.”
* **“In summary, the Builder pattern is a good choice when designing classes whose constructors or static factories would have more than a handful of parameters.”**

## Item 3: Enforce the singleton property with a private constructor or an enum type

* “A *singleton* is simply a class that is instantiated exactly once [Gamma95].”
* “Singletons typically represent either a stateless object such as a function (Item 24) or a system component that is intrinsically unique.”
* “**Making a class a singleton can make it difficult to test its clients** because it’s impossible to substitute a mock implementation for a singleton unless it implements an interface that serves as its type.”

```java
// Singleton with public final field
public class Elvis {
    public static final Elvis INSTANCE = new Elvis();
    private Elvis() { ... }

    public void leaveTheBuilding() { ... }
}
```

* “The main advantage of the public field approach is that the API makes it clear that the class is a singleton: the public static field is final, so it will always contain the same object reference. The second advantage is that it’s simpler.”

```java
// Singleton with static factory
public class Elvis {
    private static final Elvis INSTANCE = new Elvis();
    private Elvis() { ... }
    public static Elvis getInstance() { return INSTANCE; }

    public void leaveTheBuilding() { ... }
}
```

* “One advantage of the static factory approach is that it gives you the flexibility to change your mind about whether the class is a singleton without changing its API.”
* “A second advantage is that you can write a generic singleton factory if your application requires it (Item 30).”
* “A final advantage of using a static factory is that a method reference can be used as a supplier.”
* “To make a singleton class that uses either of these approaches *serializable* (Chapter 12), it is not sufficient merely to add `implements Serializable` to its declaration. To maintain the singleton guarantee, declare all instance fields transient and provide a readResolve method (Item 89). Otherwise, each time a serialized instance is deserialized, a new instance will be created.”

```java
// readResolve method to preserve singleton property
private Object readResolve() {
     // Return the one true Elvis and let the garbage collector
     // take care of the Elvis impersonator.
    return INSTANCE;
}
```

```java
// Enum singleton - the preferred approach
public enum Elvis {
    INSTANCE;

 		public void leaveTheBuilding() { ... }
}
```

* “This approach is similar to the public field approach, but it is more concise, provides the serialization machinery for free, and provides an ironclad guarantee against multiple instantiation, even in the face of sophisticated serialization or reflection attacks.”
* **“A single-element enum type is often the best way to implement a singleton.”**

## Item 4: Enforce noninstantiability with a private constructor

* “*Utility classes* were not designed to be instantiated”
* **“Attempting to enforce noninstantiability by making a class abstract does not work.”**
  * “The class can be subclassed and the subclass instantiated. Furthermore, it misleads the user into thinking the class was designed for inheritance (Item 19).”
* **“A class can be made noninstantiable by including a private constructor.”**

```java
// Noninstantiable utility class
public class UtilityClass {
    // Suppress default constructor for noninstantiability
    private UtilityClass() {
        throw new AssertionError();
    }
    ... // Remainder omitted
}
```

* “The `AssertionError` isn’t strictly required, but it provides insurance in case the constructor is accidentally invoked from within the class. It guarantees the class will never be instantiated under any circumstances.”
* “As a side effect, this idiom also prevents the class from being subclassed.”
  * “All constructors must invoke a superclass constructor, explicitly or implicitly, and a subclass would have no accessible superclass constructor to invoke.”


## Item 5: Prefer dependency injection to hardwiring resources

* **“Static utility classes and singletons are inappropriate for classes whose behavior is parameterized by an underlying resource.”**
* “A simple pattern that satisfies this requirement is to **pass the resource into the constructor when creating a new instance**.”


```java
// Dependency injection provides flexibility and testability
public class SpellChecker {
    private final Lexicon dictionary;

    public SpellChecker(Lexicon dictionary) {
        this.dictionary = Objects.requireNonNull(dictionary);
    }

    public boolean isValid(String word) { ... }
    public List<String> suggestions(String typo) { ... }
}
```

* “Dependency injection works with an arbitrary number of resources and arbitrary dependency graphs.”
* “It preserves immutability (Item 17), so multiple clients can share dependent objects.”
* “Dependency injection is equally applicable to constructors, static factories (Item 1), and builders (Item 2).”
* “A useful variant of the pattern is to pass a resource factory to the constructor. A factory is an object that can be called repeatedly to create instances of a type. Such factories embody the Factory Method pattern [Gamma95].”
  * “The `Supplier<T>` interface, introduced in Java 8, is perfect for representing factories.”
  * “Methods that take a Supplier<T> on input should typically constrain the factory’s type parameter using a bounded wildcard type (Item 31) to allow the client to pass in a factory that creates any subtype of a specified type.”

```java
Mosaic create(Supplier<? extends Tile> tileFactory) { ... }
```

* **“In summary, do not use a singleton or static utility class to implement a class that depends on one or more underlying resources whose behavior affects that of the class, and do not have the class create these resources directly. Instead, pass the resources, or factories to create them, into the constructor (or static factory or builder).”**

## Item 6: Avoid creating unnecessary objects

* **“It is often appropriate to reuse a single object instead of creating a new functionally equivalent object each time it is needed. ”**
  * “Reuse can be both faster and more stylish."
  * "An object can always be reused if it is immutable (Item 17).”
* “You can often avoid creating unnecessary objects by using static factory methods (Item 1) in preference to constructors on immutable classes that provide both.”
* “Some object creations are much more expensive than others. If you’re going to need such an “expensive object” repeatedly, it may be advisable to cache it for reuse.”
  * **“While String.matches is the easiest way to check if a string matches a regular expression, it’s not suitable for repeated use in performance-critical situations.”**
* “Another way to create unnecessary objects is autoboxing, which allows the programmer to mix primitive and boxed primitive types, boxing and unboxing automatically as needed.”
  * **“Autoboxing blurs but does not erase the distinction between primitive and boxed primitive types.”**
  * **“Prefer primitives to boxed primitives, and watch out for unintentional autoboxing.”**
* “Avoiding object creation by maintaining your own object pool is a bad idea unless the objects in the pool are extremely heavyweight.”
  * “Generally speaking, however, maintaining your own object pools clutters your code, increases memory footprint, and harms performance.”

* “The counterpoint to this item is Item 50 on defensive copying. The present item says, “Don’t create a new object when you should reuse an existing one,” while Item 50 says, “Don’t reuse an existing object when you should create a new one.”


## Item 7: Eliminate obsolete object references

* **“An obsolete reference is simply a reference that will never be dereferenced again.”**
* “If an object reference is unintentionally retained, not only is that object excluded from garbage collection, but so too are any objects referenced by that object, and so on.”
* “The fix for this sort of problem is simple: **null out references once they become obsolete**.”
  * “An added benefit of nulling out obsolete references is that if they are subsequently dereferenced by mistake, the program will immediately fail with a NullPointerException, rather than quietly doing the wrong thing.”
* **“Nulling out object references should be the exception rather than the norm.”**
  * “The best way to eliminate an obsolete reference is to let the variable that contained the reference fall out of scope. This occurs naturally if you define each variable in the narrowest possible scope (Item 57).”
* **“Whenever a class manages its own memory, the programmer should be alert for memory leaks”**
* **“Another common source of memory leaks is caches.”**
  * “The cache should occasionally be cleansed of entries that have fallen into disuse.”

* **“A third common source of memory leaks is listeners and other callbacks.”**
  * “One way to ensure that callbacks are garbage collected promptly is to store only weak references to them, for instance, by storing them only as keys in a WeakHashMap.”
* “Because memory leaks typically do not manifest themselves as obvious failures, they may remain present in a system for years. They are typically discovered only as a result of careful code inspection or with the aid of a debugging tool known as a heap profiler. Therefore, it is very desirable to learn to anticipate problems like this before they occur and prevent them from happening.”


## Item 8: Avoid finalizers and cleaners

* **“Finalizers are unpredictable, often dangerous, and generally unnecessary.”**
* **“Cleaners are less dangerous than finalizers, but still unpredictable, slow, and generally unnecessary.”**
* **“You should never do anything time-critical in a finalizer or cleaner.”**
* “Cleaners are a bit better than finalizers in this regard because class authors have control over their own cleaner threads, but cleaners still run in the background, under the control of the garbage collector, so there can be no guarantee of prompt cleaning.”
* **“You should never depend on a finalizer or cleaner to update persistent state.”**
* “Another problem with finalizers is that an uncaught exception thrown during finalization is ignored, and finalization of that object terminates.”
* **“There is a severe performance penalty for using finalizers and cleaners.”**
* **“Finalizers have a serious security problem: they open your class up to finalizer attacks.”**
* “So what should you do instead of writing a finalizer or cleaner for a class whose objects encapsulate resources that require termination, such as files or threads? Just **have your class implement `AutoCloseable`**, and require its clients to invoke the `close` method on each instance when it is no longer needed, typically using `try`-with-resources to ensure termination even in the face of exceptions (Item 9).”
  * “One detail worth mentioning is that the instance must keep track of whether it has been closed: the `close` method must record in a field that the object is no longer valid, and other methods must check this field and throw an `IllegalStateException` if they are called after the object has been closed.”
* **“In summary, don’t use cleaners, or in releases prior to Java 9, finalizers, except as a safety net or to terminate noncritical native resources.”**

## Item 9: Prefer `try`-with-resources to `try`-`finally`

* “Historically, a try-finally statement was the best way to guarantee that a resource would be closed properly, even in the face of an exception or return.”

```java
// try-finally - No longer the best way to close resources!
static String firstLineOfFile(String path) throws IOException {
    BufferedReader br = new BufferedReader(new FileReader(path));
    try {
        return br.readLine();
    } finally {
        br.close();
    }
}
```

* “To be usable with the `try`-with-resources statement, a resource must implement the AutoCloseable interface, which consists of a single void-returning `close` method.”

```java
// try-with-resources on multiple resources - short and sweet
static void copy(String src, String dst) throws IOException {
    try (InputStream   in = new FileInputStream(src);
         OutputStream out = new FileOutputStream(dst)) {
        byte[] buf = new byte[BUFFER_SIZE];
        int n;
        while ((n = in.read(buf)) >= 0)
            out.write(buf, 0, n);
    }
}
```

* **“Always use `try`-with-resources in preference to `try`-`finally` when working with resources that must be closed. The resulting code is shorter and clearer, and the exceptions that it generates are more useful.”**

