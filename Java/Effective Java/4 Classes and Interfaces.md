# Chapter 4. Classes and Interfaces

## Item 15: Minimize the accessibility of classes and members

* “A well-designed component hides all its implementation details, cleanly separating its API from its implementation. Components then communicate only through their APIs and are oblivious to each others’ inner workings. This concept, known as *information hiding* or *encapsulation*, is a fundamental tenet of software design [Parnas72].”
* “Information hiding is important for many reasons, most of which stem from the fact that it *decouples* the components that comprise a system, allowing them to be developed, tested, optimized, used, understood, and modified in isolation.”
* “Java has many facilities to aid in information hiding. The *access control* mechanism [JLS, 6.6] specifies the accessibility of classes, interfaces, and members. The *accessibility* of an entity is determined by the location of its declaration and by which, if any, of the access modifiers (`private`, `protected`, and `public`) is present on the declaration.”
* “The rule of thumb is simple: **make each class or member as inaccessible as possible**.”
  * “By making it package-private, you make it part of the implementation rather than the exported API, and you can modify it, replace it, or eliminate it in a subsequent release without fear of harming existing clients.”
  * “If you make it public, you are obligated to support it forever to maintain compatibility.”
  * “If a package-private top-level class or interface is used by only one class, consider making the top-level class a private static nested class of the sole class that uses it (Item 24).”

* “For members (fields, methods, nested classes, and nested interfaces), there are four possible access levels, listed here in order of increasing accessibility:”
  * “**private**—The member is accessible only from the top-level class where it is declared.”
  * “**package-private**—The member is accessible from any class in the package where it is declared. Technically known as default access, this is the access level you get if no access modifier is specified (except for interface members, which are public by default).”
  * “**protected**—The member is accessible from subclasses of the class where it is declared (subject to a few restrictions [JLS, 6.6.2]) and from any class in the package where it is declared.”
  * “**public**—The member is accessible from anywhere.”
* “That said, both private and package-private members are part of a class’s implementation and do not normally impact its exported API. These fields can, however, “leak” into the exported API if the class implements `Serializable` (Items 86 and 87).”
* “For members of public classes, a huge increase in accessibility occurs when the access level goes from package-private to protected. A protected member is part of the class’s exported API and must be supported forever. Also, a protected member of an exported class represents a public commitment to an implementation detail (Item 19). The need for protected members should be relatively rare.”
* “**If a method overrides a superclass method, it cannot have a more restrictive access level in the subclass than in the superclass** [JLS, 8.4.8.3]. This is necessary to ensure that an instance of the subclass is usable anywhere that an instance of the superclass is usable (the Liskov substitution principle, see Item 15).”
* “To facilitate testing your code, you may be tempted to make a class, interface, or member more accessible than otherwise necessary. This is fine up to a point. It is acceptable to make a private member of a public class package-private in order to test it, but it is not acceptable to raise the accessibility any higher.”
* “**Instance fields of public classes should rarely be public** (Item 16).”
  * **“Classes with public mutable fields are not generally thread-safe.”**
* “You can expose constants via public static final fields, assuming the constants form an integral part of the abstraction provided by the class. By convention, such fields have names consisting of capital letters, with words separated by underscores (Item 68). It is critical that these fields contain either primitive values or references to immutable objects (Item 17).”
* **“It is wrong for a class to have a public static final array field, or an accessor that returns such a field.”**
  * “You can make the public array private and add a public immutable list.”
  * “Alternatively, you can make the array private and add a public method that returns a copy of a private array.”
* “As of Java 9, there are two additional, implicit access levels introduced as part of the *module system*.”
  * “A module is a grouping of packages, like a package is a grouping of classes.”
  * “A module may explicitly export some of its packages via *export declarations* in its *module declaration* (which is by convention contained in a source file named module-info.java).”
  * “Using the module system allows you to share classes among packages within a module without making them visible to the entire world.”


## Item 16: In public classes, use accessor methods, not public fields

```java
// Degenerate classes like this should not be public!
class Point {
    public double x;
    public double y;
}
```

* “Hard-line object-oriented programmers feel that such classes are anathema and should always be replaced by classes with private fields and public *accessor methods* (getters) and, for mutable classes, *mutators* (setters).”


```java
// Encapsulation of data by accessor methods and mutators
class Point {
    private double x;
    private double y;

    public Point(double x, double y) {
        this.x = x;
        this.y = y;
    }

    public double getX() { return x; }
    public double getY() { return y; }

    public void setX(double x) { this.x = x; }
    public void setY(double y) { this.y = y; }
}
```

* **“If a class is accessible outside its package, provide accessor methods to preserve the flexibility to change the class’s internal representation.”**
* **“If a class is package-private or is a private nested class, there is nothing inherently wrong with exposing its data fields.”**
* **“In summary, public classes should never expose mutable fields. It is less harmful, though still questionable, for public classes to expose immutable fields. It is, however, sometimes desirable for package-private or private nested classes to expose fields, whether mutable or immutable.”**

## Item 17: Minimize mutability

* “An immutable class is simply a class whose instances cannot be modified. All of the information contained in each instance is fixed for the lifetime of the object, so no changes can ever be observed.”
* “To make a class immutable, follow these five rules:”

  * **“Don’t provide methods that modify the object’s state** (known as *mutators*).”
  * **“Ensure that the class can’t be extended.”**
  * **“Make all fields final.”**
  * **“Make all fields private.”**
    * “While it is technically permissible for immutable classes to have public final fields containing primitive values or references to immutable objects, it is not recommended because it precludes changing the internal representation in a later release (Items 15 and 16).”

  * **“Ensure exclusive access to any mutable components.”**
    * “Make *defensive copies* (Item 50) in constructors, accessors, and `readObject` methods (Item 88).”
* **“Immutable objects are simple.”**
  * “Mutable objects, on the other hand, can have arbitrarily complex state spaces. If the documentation does not provide a precise description of the state transitions performed by mutator methods, it can be difficult or impossible to use a mutable class reliably.”

* **“Immutable objects are inherently thread-safe; they require no synchronization.”**
  * **“Immutable objects can be shared freely.”**
  * “An immutable class can provide static factories (Item 1) that cache frequently requested instances to avoid creating new instances when existing ones would do.”
  * “A consequence of the fact that immutable objects can be shared freely is that you never have to make *defensive copies* of them (Item 50).”
    * “Therefore, you need not and should not provide a `clone` method or *copy constructor* (Item 13) on an immutable class.”
* **“Not only can you share immutable objects, but they can share their internals.”**
  * “**Immutable objects make great building blocks for other objects**, whether mutable or immutable.”
* **“Immutable objects provide failure atomicity for free (Item 76).”**
  * “Their state never changes, so there is no possibility of a temporary inconsistency.”
* **“The major disadvantage of immutable classes is that they require a separate object for each distinct value.”**
  * “Creating these objects can be costly, especially if they are large.”
  * “The performance problem is magnified if you perform a multistep operation that generates a new object at every step, eventually discarding all objects except the final result. There are two approaches to coping with this problem.”
    * “The first is to guess which multistep operations will be commonly required and to provide them as primitives.”
    * “If not applicable, then your best bet is to provide a public mutable companion class.”
* “Instead of making an immutable class final, you can make all of its constructors private or package-private and add public static factories in place of the public constructors (Item 1).”

```java
// Immutable class with static factories instead of constructors
public class Complex {
    private final double re;
    private final double im;

    private Complex(double re, double im) {
        this.re = re;
        this.im = im;
    }

    public static Complex valueOf(double re, double im) {
        return new Complex(re, im);
    }

    ... // Remainder unchanged
}
```

* “However, some immutable classes have one or more nonfinal fields in which they cache the results of expensive computations the first time they are needed. If the same value is requested again, the cached value is returned, saving the cost of recalculation. This trick works precisely because the object is immutable, which guarantees that the computation would yield the same result if it were repeated.”
* “One caveat should be added concerning serializability. If you choose to have your immutable class implement `Serializable` and it contains one or more fields that refer to mutable objects, you must provide an explicit `readObject` or `readResolve` method, or use the `ObjectOutputStream.writeUnshared` and `ObjectInputStream.readUnshared` methods, even if the default serialized form is acceptable. Otherwise an attacker could create a mutable instance of your class.”
* **“Classes should be immutable unless there’s a very good reason to make them mutable.”**
* **“If a class cannot be made immutable, limit its mutability as much as possible.”**
  * “Combining the advice of this item with that of Item 15, your natural inclination should be to **declare every field private final unless there’s a good reason to do otherwise**.”
* **“Constructors should create fully initialized objects with all of their invariants established.”**

## Item 18: Favor composition over inheritance

* “Inheritance is a powerful way to achieve code reuse, but it is not always the best tool for the job.”
* “It is safe to use inheritance within a package, where the subclass and the superclass implementations are under the control of the same programmers. It is also safe to use inheritance when extending classes specifically designed and documented for extension (Item 19). Inheriting from ordinary concrete classes across package boundaries, however, is dangerous.”
* “**Unlike method invocation, inheritance violates encapsulation** [Snyder86].”
  * “In other words, a subclass depends on the implementation details of its superclass for its proper function. ”
  * “A related cause of fragility in subclasses is that their superclass can acquire new methods in subsequent releases.”
  * “Both of these problems stem from overriding methods.”

* **“Instead of extending an existing class, give your new class a private field that references an instance of the existing class. This design is called *composition* because the existing class becomes a component of the new one.”**
  * “Each instance method in the new class invokes the corresponding method on the contained instance of the existing class and returns the results. This is known as *forwarding*, and the methods in the new class are known as *forwarding methods*.”
  * “The resulting class will be rock solid, with no dependencies on the implementation details of the existing class. Even adding new methods to the existing class will have no impact on the new class.”

```java
// Wrapper class - uses composition in place of inheritance
public class InstrumentedSet<E> extends ForwardingSet<E> {
    private int addCount = 0;

    public InstrumentedSet(Set<E> s) {
        super(s);
    }

    @Override public boolean add(E e) {
        addCount++;
        return super.add(e);
     }
     @Override public boolean addAll(Collection<? extends E> c) {
         addCount += c.size();
         return super.addAll(c);
     }
     public int getAddCount() {
         return addCount;
     }
}

// Reusable forwarding class
public class ForwardingSet<E> implements Set<E> {
    private final Set<E> s;
    public ForwardingSet(Set<E> s) { this.s = s; }

    public void clear()               { s.clear();            }
    public boolean contains(Object o) { return s.contains(o); }
    public boolean isEmpty()          { return s.isEmpty();   }
    public int size()                 { return s.size();      }
    public Iterator<E> iterator()     { return s.iterator();  }
    public boolean add(E e)           { return s.add(e);      }
    public boolean remove(Object o)   { return s.remove(o);   }
    public boolean containsAll(Collection<?> c)
                                   { return s.containsAll(c); }
    public boolean addAll(Collection<? extends E> c)
                                   { return s.addAll(c);      }
    public boolean removeAll(Collection<?> c)
                                   { return s.removeAll(c);   }
    public boolean retainAll(Collection<?> c)
                                   { return s.retainAll(c);   }
    public Object[] toArray()          { return s.toArray();  }
    public <T> T[] toArray(T[] a)      { return s.toArray(a); }
    @Override public boolean equals(Object o)
                                       { return s.equals(o);  }
    @Override public int hashCode()    { return s.hashCode(); }
    @Override public String toString() { return s.toString(); }
}
```

* “This is also known as the *Decorator pattern* [Gamma95] because the `InstrumentedSet` class “decorates” a set by adding instrumentation. Sometimes the combination of composition and forwarding is loosely referred to as *delegation*. Technically it’s not delegation unless the wrapper object passes itself to the wrapped object [Lieberman86; Gamma95].”
* “The disadvantages of wrapper classes are few. One caveat is that wrapper classes are not suited for use in *callback frameworks*, wherein objects pass self-references to other objects for subsequent invocations (“callbacks”). Because a wrapped object doesn’t know of its wrapper, it passes a reference to itself (`this`) and callbacks elude the wrapper. This is known as the *SELF problem* [Lieberman86].”
* “It’s tedious to write forwarding methods, but you have to write the reusable forwarding class for each interface only once, and forwarding classes may be provided for you.”
* “Inheritance is appropriate only in circumstances where the subclass really is a *subtype* of the superclass. In other words, **a class B should extend a class A only if an “is-a” relationship exists between the two classes**.”
* **“Inheritance propagates any flaws in the superclass’s API, while composition lets you design a new API that hides these flaws.”

## Item 19: Design and document for inheritance or else prohibit it

* “First, the class must document precisely the effects of overriding any method. In other words, **the class must document its self-use of overridable methods**.”
  * “For each public or protected method, the documentation must indicate which overridable methods the method invokes, in what sequence, and how the results of each invocation affect subsequent processing.”
  * “More generally, a class must document any circumstances under which it might invoke an overridable method. For example, invocations might come from background threads or static initializers.”
* “A method that invokes overridable methods contains a description of these invocations at the end of its documentation comment. The description is in a special section of the specification, labeled “Implementation Requirements,” which is generated by the Javadoc tag `@implSpec`. This section describes the inner workings of the method.”
* “Designing for inheritance involves more than just documenting patterns of self-use. To allow programmers to write efficient subclasses without undue pain, **a class may have to provide hooks into its internal workings in the form of judiciously chosen protected methods** or, in rare instances, protected fields.”
* **“The only way to test a class designed for inheritance is to write subclasses.”**
  * “If you omit a crucial protected member, trying to write a subclass will make the omission painfully obvious. ”

  * “Experience shows that three subclasses are usually sufficient to test an extendable class.”
* “When you design for inheritance a class that is likely to achieve wide use, realize that you are committing forever to the self-use patterns that you document and to the implementation decisions implicit in its protected methods and fields. These commitments can make it difficult or impossible to improve the performance or functionality of the class in a subsequent release. ”
* **“You must test your class by writing subclasses before you release it.”**
* “**Constructors must not invoke overridable methods**, directly or indirectly.”
  * “The superclass constructor runs before the subclass constructor, so the overriding method in the subclass will get invoked before the subclass constructor has run. If the overriding method depends on any initialization performed by the subclass constructor, the method will not behave as expected.”
  * “Note that it is safe to invoke private methods, final methods, and static methods, none of which are overridable, from a constructor.”
* **“The best solution to this problem is to prohibit subclassing in classes that are not designed and documented to be safely subclassed.”**
  * “There are two ways to prohibit subclassing. The easier of the two is to declare the class final. The alternative is to make all the constructors private or package-private and to add public static factories in place of the constructors.”
* “If a concrete class does not implement a standard interface, then you may inconvenience some programmers by prohibiting inheritance. If you feel that you must allow inheritance from such a class, one reasonable approach is to ensure that the class never invokes any of its overridable methods and to document this fact.”
  * “You can eliminate a class’s self-use of overridable methods mechanically, without changing its behavior. Move the body of each overridable method to a private “helper method” and have each overridable method invoke its private helper method. Then replace each self-use of an overridable method with a direct invocation of the overridable method’s private helper method.”


## Item 20: Prefer interfaces to abstract classes

* “Java has two mechanisms to define a type that permits multiple implementations: interfaces and abstract classes. Since the introduction of *default methods* for interfaces in Java 8 [JLS 9.4.3], both mechanisms allow you to provide implementations for some instance methods.”
  * “A major difference is that to implement the type defined by an abstract class, a class must be a subclass of the abstract class. Because Java permits only single inheritance, this restriction on abstract classes severely constrains their use as type definitions.”
* **“Existing classes can easily be retrofitted to implement a new interface.”**
  * “All you have to do is to add the required methods, if they don’t yet exist, and to add an `implements` clause to the class declaration.”
  * “If you want to have two classes extend the same abstract class, you have to place it high up in the type hierarchy where it is an ancestor of both classes. Unfortunately, this can cause great collateral damage to the type hierarchy, forcing all descendants of the new abstract class to subclass it, whether or not it is appropriate.”

* **“Interfaces are ideal for defining mixins.”**
  * “Loosely speaking, a mixin is a type that a class can implement in addition to its “primary type,” to declare that it provides some optional behavior. ”
* **“Interfaces allow for the construction of nonhierarchical type frameworks.”**
* “**Interfaces enable safe, powerful functionality enhancements** via the *wrapper class* idiom (Item 18).”
* “When there is an obvious implementation of an interface method in terms of other interface methods, consider providing implementation assistance to programmers in the form of a default method.”
  * “You can’t add default methods to an interface that you don’t control.”

* **“You can, however, combine the advantages of interfaces and abstract classes by providing an abstract *skeletal implementation class* to go with an interface.”**
  * “The interface defines the type, perhaps providing some default methods, while the skeletal implementation class implements the remaining non-primitive interface methods atop the primitive interface methods. Extending a skeletal implementation takes most of the work out of implementing an interface. This is the *Template Method* pattern [Gamma95].”
  * “By convention, skeletal implementation classes are called `Abstract`*Interface*, where *Interface* is the name of the interface they implement.”
  * “Because skeletal implementations are designed for inheritance, you should follow all of the design and documentation guidelines in Item 19.”
* “A minor variant on the skeletal implementation is the *simple implementation*, exemplified by `AbstractMap.SimpleEntry`. A simple implementation is like a skeletal implementation in that it implements an interface and is designed for inheritance, but it differs in that it isn’t abstract: it is the simplest possible working implementation. You can use it as it stands or subclass it as circumstances warrant.”


## Item 21: Design interfaces for posterity

* “Prior to Java 8, it was impossible to add methods to interfaces without breaking existing implementations.”
  * “If you added a new method to an interface, existing implementations would, in general, lack the method, resulting in a compile-time error.”
  * “In Java 8, the default method construct was added [JLS 9.4], with the intent of allowing the addition of methods to existing interfaces. ”
  * “But adding new methods to existing interfaces is fraught with risk.”
* “The declaration for a default method includes a *default implementation* that is used by all classes that implement the interface but do not implement the default method.”
  * “Default methods are “injected” into existing implementations without the knowledge or consent of their implementors.”
  * **“It is not always possible to write a default method that maintains all invariants of every conceivable implementation.”**
* **“In the presence of default methods, existing implementations of an interface may compile without error or warning but fail at runtime.”**
* “It is also worth noting that default methods were not designed to support removing methods from interfaces or changing the signatures of existing methods. Neither of these interface changes is possible without breaking existing clients.”
* “The moral is clear. Even though default methods are now a part of the Java platform, **it is still of the utmost importance to design interfaces with great care**.”
* **“While it may be possible to correct some interface flaws after an interface is released, you cannot count on it.”**

## Item 22: Use interfaces only to define types

* “When a class implements an interface, the interface serves as a type that can be used to refer to instances of the class. That a class implements an interface should therefore say something about what a client can do with instances of the class. It is inappropriate to define an interface for any other purpose.”
* “One kind of interface that fails this test is the so-called *constant interface*. Such an interface contains no methods; it consists solely of static final fields, each exporting a constant. Classes using these constants implement the interface to avoid the need to qualify constant names with a class name.”
  * **“The constant interface pattern is a poor use of interfaces.”**
  * “That a class uses some constants internally is an implementation detail. Implementing a constant interface causes this implementation detail to leak into the class’s exported API.”
  * “If in a future release the class is modified so that it no longer needs to use the constants, it still must implement the interface to ensure binary compatibility.”
  * “If the constants are strongly tied to an existing class or interface, you should add them to the class or interface.”
  * “If the constants are best viewed as members of an enumerated type, you should export them with an enum type (Item 34). Otherwise, you should export the constants with a *noninstantiable utility* class (Item 4).”


```java
// Constant utility class
package com.effectivejava.science;

public class PhysicalConstants {
  private PhysicalConstants() { }  // Prevents instantiation

  public static final double AVOGADROS_NUMBER = 6.022_140_857e23;
  public static final double BOLTZMANN_CONST  = 1.380_648_52e-23;
  public static final double ELECTRON_MASS    = 9.109_383_56e-31;
}
```

* “Incidentally, note the use of the underscore character (`_`) in the numeric literals. Underscores, which have been legal since Java 7, have no effect on the values of numeric literals, but can make them much easier to read if used with discretion. Consider adding underscores to numeric literals, whether fixed of floating point, if they contain five or more consecutive digits.”
* “If you make heavy use of the constants exported by a utility class, you can avoid the need for qualifying the constants with the class name by making use of the *static import* facility.”


```java
// Use of static import to avoid qualifying constants
import static com.effectivejava.science.PhysicalConstants.*;
```

* **“In summary, interfaces should be used only to define types. They should not be used merely to export constants.”**

## Item 23: Prefer class hierarchies to tagged classes

```java
// Tagged class - vastly inferior to a class hierarchy!
class Figure {
    enum Shape { RECTANGLE, CIRCLE };

    // Tag field - the shape of this figure
    final Shape shape;

    // These fields are used only if shape is RECTANGLE
    double length;
    double width;

    // This field is used only if shape is CIRCLE
    double radius;

    // Constructor for circle
    Figure(double radius) {
        shape = Shape.CIRCLE;
        this.radius = radius;
    }

    // Constructor for rectangle
    Figure(double length, double width) {
        shape = Shape.RECTANGLE;
        this.length = length;
        this.width = width;
    }

    double area() {
        switch(shape) {
          case RECTANGLE:
            return length * width;
          case CIRCLE:
            return Math.PI * (radius * radius);
          default:
            throw new AssertionError(shape);
        }
    }
}
```

* “Occasionally you may run across a class whose instances come in two or more flavors and contain a tag field indicating the flavor of the instance.”
  * “They are cluttered with boilerplate, including enum declarations, tag fields, and switch statements.”
  * “Readability is further harmed because multiple implementations are jumbled together in a single class.”
  * “Memory footprint is increased because instances are burdened with irrelevant fields belonging to other flavors.”
  * “Fields can’t be made final unless constructors initialize irrelevant fields, resulting in more boilerplate. Constructors must set the tag field and initialize the right data fields with no help from the compiler: if you initialize the wrong fields, the program will fail at runtime.”
  * “You can’t add a flavor to a tagged class unless you can modify its source file. If you do add a flavor, you must remember to add a case to every switch statement, or the class will fail at runtime.”
  * “Finally, the data type of an instance gives no clue as to its flavor.”
* “In short, **tagged classes are verbose, error-prone, and inefficient**.”
* **“A tagged class is just a pallid imitation of a class hierarchy.”**
  * “To transform a tagged class into a class hierarchy, first define an abstract class containing an abstract method for each method in the tagged class whose behavior depends on the tag value.”
    * “If there are any methods whose behavior does not depend on the value of the tag, put them in this class. Similarly, if there are any data fields used by all the flavors, put them in this class.”
  * “Next, define a concrete subclass of the root class for each flavor of the original tagged class.”
  * “Also include in each subclass the appropriate implementation of each abstract method in the root class.”
* “Another advantage of class hierarchies is that they can be made to reflect natural hierarchical relationships among types, allowing for increased flexibility and better compile-time type checking.”


## Item 24: Favor static member classes over nonstatic

* “A nested class should exist only to serve its enclosing class.”
* “There are four kinds of nested classes: *static member classes*, *nonstatic member classes*, *anonymous classes*, and *local classes*. All but the first kind are known as *inner classes*.”
* “A static member class is the simplest kind of nested class. It is best thought of as an ordinary class that happens to be declared inside another class and has access to all of the enclosing class’s members, even those declared private.”
* “A static member class is a static member of its enclosing class and obeys the same accessibility rules as other static members. If it is declared private, it is accessible only within the enclosing class, and so forth.”
* **“One common use of a static member class is as a public helper class, useful only in conjunction with its outer class.”**
* “Each instance of a nonstatic member class is implicitly associated with an enclosing instance of its containing class. Within instance methods of a nonstatic member class, you can invoke methods on the enclosing instance or obtain a reference to the enclosing instance using the *qualified this* construct [JLS, 15.8.4].”
* “If an instance of a nested class can exist in isolation from an instance of its enclosing class, then the nested class must be a static member class: it is impossible to create an instance of a nonstatic member class without an enclosing instance.”
* **“One common use of a nonstatic member class is to define an *Adapter* [Gamma95] that allows an instance of the outer class to be viewed as an instance of some unrelated class.”**
* “**If you declare a member class that does not require access to an enclosing instance, always put the `static` modifier in its declaration**, making it a static rather than a nonstatic member class.”
  * “If you omit this modifier, each instance will have a hidden extraneous reference to its enclosing instance. As previously mentioned, storing this reference takes time and space.”
  * “More seriously, it can result in the enclosing instance being retained when it would otherwise be eligible for garbage collection (Item 7). The resulting memory leak can be catastrophic. It is often difficult to detect because the reference is invisible.”
* **“A common use of private static member classes is to represent components of the object represented by their enclosing class. ”**
* “As you would expect, an anonymous class has no name. It is not a member of its enclosing class. Rather than being declared along with other members, it is simultaneously declared and instantiated at the point of use. Anonymous classes are permitted at any point in the code where an expression is legal.”
* “Anonymous classes have enclosing instances if and only if they occur in a nonstatic context. But even if they occur in a static context, they cannot have any static members other than *constant variables*, which are final primitive or string fields initialized to constant expressions [JLS, 4.12.4].”
* “Before lambdas were added to Java (Chapter 6), anonymous classes were the preferred means of creating small *function objects* and *process objects* on the fly, but lambdas are now preferred (Item 42). Another common use of anonymous classes is in the implementation of static factory methods (see `intArrayAsList` in Item 20).”
* “Local classes are the least frequently used of the four kinds of nested classes. A local class can be declared practically anywhere a local variable can be declared and obeys the same scoping rules. Local classes have attributes in common with each of the other kinds of nested classes. Like member classes, they have names and can be used repeatedly. Like anonymous classes, they have enclosing instances only if they are defined in a nonstatic context, and they cannot contain static members. And like anonymous classes, they should be kept short so as not to harm readability.”


## Item 25: Limit source files to a single top-level class

* “While the Java compiler lets you define multiple top-level classes in a single source file, there are no benefits associated with doing so, and there are significant risks. The risks stem from the fact that defining multiple top-level classes in a source file makes it possible to provide multiple definitions for a class. Which definition gets used is affected by the order in which the source files are passed to the compiler.”


```java
// Two classes defined in one file. Don't ever do this!
class Utensil {
    static final String NAME = "pan";
}

class Dessert {
    static final String NAME = "cake";
}
```

* “If you are tempted to put multiple top-level classes into a single source file, consider using static member classes (Item 24) as an alternative to splitting the classes into separate source files.”
* “If the classes are subservient to another class, making them into static member classes is generally the better alternative because it enhances readability and makes it possible to reduce the accessibility of the classes by declaring them private (Item 15).”

```java
// Static member classes instead of multiple top-level classes
public class Test {
    public static void main(String[] args) {
        System.out.println(Utensil.NAME + Dessert.NAME);
    }

    private static class Utensil {
        static final String NAME = "pan";
    }

    private static class Dessert {
        static final String NAME = "cake";
    }
}
```

* **“Never put multiple top-level classes or interfaces in a single source file.”**
