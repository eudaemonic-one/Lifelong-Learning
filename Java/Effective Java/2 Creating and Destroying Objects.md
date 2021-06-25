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

