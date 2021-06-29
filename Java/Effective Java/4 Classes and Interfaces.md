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
