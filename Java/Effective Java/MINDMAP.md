# The Mind Map of Effective Java

* Types: interfaces (annotations), classes, arrays, primitives.
* A class's members: fields, methods, member classes, member interfaces.
* A method's signature: name and the type of its formal parameters.

## Creating and Destroying Objects

* Consider ***static factory methods*** instead of constructors
  * Consequences
    * => they have name
    * => instance control (not required to create a new object, can return an object of any subtype of their return type, the returned object can vary from call to call as a function of the input parameters, the class of the returned object need not exist when the class containing the method is written)
    * => classes providing only static factory methods cannot be subclassed.
    * => hard for programmers to find.
  * Implementation
    * common names: `from`, `of`, `valueOf`, `instance` or `getInstance`, `create` or `newInstance`, `getType`, `newType`, `type`
  * Applicability
    * often preferable.
* Consider a ***builder*** when faced with many constructor parameters
  * Motivation
    * static factories, constructors => do not scale well to large numbers of optional parameters.
    * *Telescoping constructor pattern* => hard to read and write, do not scale well.
    * *JavaBeans Pattern* => allows inconsistency (through its construction), mandates mutability.
  * Consequences
    * => simulate named optional parameters => easy to read and write.
    * => check invariants in the builder's constructor.
    * => well suited to class hierarchies => use a parallel hierarchy of builders, each nested in the corresponding class.
    * => creating the builder could be a problem in performance-critical situations.
  * Applicability
    * use only if there are four or more parameters.
    * better to start with a builder in the first place.
* Enforce the ***singleton*** property with a private constructor or an enum type
  * Motivation
    * represent either a stateless object or a intrinsically unique system component
  * Consequences
    * => a class is instantiated exactly once.
    * => make it difficult to test its clients.
  * Implementation
    * singleton with public final field.
      * => error-prone to accessibility attack.
    * singleton with static factory.
      * => difficult to `implements serializable` => must declare all fields `transient` and provide a `readResolve` method.
    * `Enum` singleton is the preferred approach.
      * => consice, provides the serialization machinery for free.
      * => cannot extend a superclass other than `Enum`.
* Enforce **noninstantiability** with a private constructor
  * Motivation
    * *utility classes* were not designed to be instantiated.
  * Consequences
    * a class can be made noninstantiable by including a private constructor => suppress default constructor.
  * Implementation
    * Throw `AssertionError` optionally.
* Prefer ***dependency injection*** over hardwiring resources
  * Motivation
    * many classes depend on one or more underlying resources.
  * Consequences
    * => provides flexibility and testability.
    * => preserves immutability of shared dependent objects.
  * Implementation
    * pass the resource into the constructor when creating a new instance.
    * pass a resource *factory* (the *Factory Method* pattern) to the constructor.
      * `Supplier<T>` interface => using a *bounded wildcard type* `Supplier<T extends XXX>` as the factory's type parameter => can create any subtype of a specified type.
* **Avoid creating unnecessary objects**
  * Motivation
    * reuse => faster and more stylish.
  * Implementation
    * using *static factory methods*.
    * cache "expensive object".
      * e.g., `String.matches` => `Pattern`
    * prefer primitives to boxed primitives, and watch out for unintentional *autoboxing*.
  * Consequences
    * maintaining your own *object pool* => clutters your code, increases memory footprint, harms performance.
    * failing to make *defensive copying* => bugs and security holes.
* **Eliminate obsolete object references**
  * Implementation
    * null out references once they become obsolete.
      * => immediately fail with a `NullPointerException` rather than failing quietly.
      * => should be the exception rather than the norm.
    * let the variable that contained the reference fall out of scope.
      * => define each variable in the narrowest possible scope.
    * *heap profiler* => aid code inspection and debugging.
  * Consequences
    * managing its own memory => memory leaks.
      * => null out references when an element is freed.
    * caches => memory leaks.
      * => clean entries that have fallen into disuse.
    * listeners and callbacks => memory leaks
      * => storing only *weak references*.
* **Avoid finalizers and cleaners**
  * Motivation
    * finalizers => unpredictable, dangerous, unnecessary.
    * clearners => less dangerous than finalizers, but still unpredictable, slow, unnecessary.
  * Applicability
    * never do anything time-critical in a finalizer or cleaner.
    * never depend on a finalizer or cleaner to update persistent state.
  * Consequences
    * finalizer => an uncaught exception thrown during finalization is ignored, and finalization of that object terminates.
    * => *severe* performance penalty.
    * => open your class up to *finalizer attacks*.
      * => must write a final `finalize` method to protect nonfinal classes.
  * Implementation
    * Just have your class implement `AutoCloseable`.
      * => invoke `close` on each no longer needed instance.
      * => use `try`-with-resource clause to ensure termination even in the face of exceptions.
      * => keep track of whether it has been closed and throw `IllegalStateException` exceptionally.
* Prefer **`try`-with-resources** to `try`-`finally`
  * Motivation
    * many resource must be closed by invoking a `close` method.
    * `try`-`finally` is ugly when used with more than one resource.
  * Implementation
    * Implement the `AutoCloseable` interface => a single `void`-returning `close` method.

## Methods Common to All Objects

* Obey the general contract when **overriding `equals`**
  * Applicability
    * No need to override `equal` if each instance is equal only to itself.
      * Each instance of the class is inherently unique.
      * No need for the class to provide a "logical equality" test.
      * A superclass has already overridden `equals`, and the superclass behavior is appropriate for this class.
      * The class is private or package-private, and you are certain that its `equals` method will never be invoked.
  * Consequences
    * General contract => Reflexive, Symmetric, Transitive, Consistent, Non-nullity.
  * Implementation
    * Use the `==` operator to check if the argument is a reference to this object.
    * Use the `instanceof` operator to check if the argument has the correct type.
    * Cast the argument to the correct type.
    * For each "significant" field in the class, check if that field of the argument matches the corresponding field of this object.
    * Always override `hashCode` when you override `equals`.
    * Don't substitute another type for `Object` in the `equals` declaration.
      * => parameter type must be `Object`.
* **Always override `hashCode` when you override `equals`**
  * Motivation
    * equal objects must have equal hash codes.
  * Implementation
    * `result = 31 * result + c;`
    * `Objects.hash` => run more slowly.
    * might consider caching the hash code in the object if the cost is significant.
    * do not be tempted to exclude siginifcant fields from the hash code computation to improve performance.
    * don't provide a detailed specification for the value returned by `hashCode` => so clients can't reasonably depend on it => flexible to change it.
* **Always override `toString`**
  * Motivation
    * default implementation: "at sign" (`@`) and the unsigned hexadecimal representation of the hash code => not what user expect.
  * Consequences
    * => makes your class much more pleasant to use and makes systems using the class easier to debug.
    * => `toString` returns all of the interesting information contained in the object.
    * Failing to provide accessors => turn the string format into a de facto API.
  * Implementation
    * `Returns the string representation of this XXX.`
    * Whether or not you decide to specify the format, you should clearly document your intentions.
      * Specifying the format => standard, unambiguous, human-readable representation, unable to change.
      * Not specifying the format => subject to change.
* **Override `clone` judiciously**
  * Motivation
    * `Cloneable`: a *mixin interface* for classes to advertise that they permit cloning.
      * It lacks a `clone` method, and `Object`'s `clone` method is protected.
      * Even a reflective invocation may fail, because there is no guarantee that the object has an accessible `clone` method.
      * `Object`'s `clone` method returns a field-by-field copy or throws `CloneNotSupportedException`.
  * Consequences
    * A class implementing `Cloneable` is expected to provide a properly functioning public `clone` method.
    * Immutable classes should never provide a `clone` method.
    * In effect, the `clone` method functions as a constructor.
      * => ensure that it does no harm to the original object and it properly establishes invariants on the clone.
      * Calling `clone` on an array returns an array whose runtime and compile-time types are identical to those of the array being cloned.
    * Like serialization, the `Cloneable` architecture is incompatible with normal use of final fields referring to mutable objects.
      * => must support a "deep copy" method.
  * Implementation
    * Override `clone` with a public method whose return type is the class itself.
      * First call `super.clone`, then fix any fields that need fixing.
        * Copying any mutable objects and replacing the clone's references to these objects with references to their copies.
        * Serial number or unique ID need to be fixed.
    * Public `clone` method should omit the `throws` clause.
      * => need not throw `CloneNotSupportedException`.
      * => methods don't throw checked exceptions are easier to use.
    * When designing a class for inheritance
      * => implementing a properly functioning protected `clone` method that is declared to throw `CloneNotSupportedException`.
      * => `@Override protected final Object clone() throws CloneNotSupportedException { ... }`
    * When designing a thread-safe class
      * => `clone` method must be properly synchronized.
    * A better approach to object copying is to provide a *copy constructor* or *copy factory*.
      * `public Yum(Yum yum) { ... };`
      * `public static Yum newInstance(Yum yum) { ... };`
      * => don't rely on a risk-prone extralinguistic object creation mechanism.
      * => don't demand unenforceable adherence to thinly documented conventions.
      * => don't conflict with the proper use of final fields.
      * => don't throw unnecessary checked exceptions.
      * => don't require casts.
    * Interface-based *conversion constructor* and *conversion factories*
      * => allow the client to choose the implementation type of the copy.
        * e.g., `HashSet cs = new TreeSet<>(s)`
* **Consider implementing `Comparable`**
  * Consequences
    * => *natural ordering* => easy to search, compute extreme values, and maintain automatically sorted collections.
  * Implementation
    * `compareTo` method: Compares this object with the specified object for order. Returns a negative integer, zero, or a positive as this object is less than, equal to, or greater than the specified object. Throws `ClassCastException` if the specified object's type prevents it from being compared to this object.
    * Invoke the `compareTo` method recursively.
      * Use of the relational operators `<` and `>` in `compareTo` method is verbose and error-prone and no longer recommended.
    * *comparator construction methods* => using static imported comparator construction methods => simple names for clarity and brevity.
      * e.g., `comparingInt`, `thenComparingInt`, `Integer.compare`.

## Classes and Interfaces

* Minimize the **accessibility of classes and members**

  * Motivation
    * *information hiding* or *encapsulation* => decoupling API from implementation => allow components to be developed, tested, optimized, used, understood, and modified in isolation.
    * make each class or member as inaccessible as possible.
  * Consequences
    * *accessibility* (*access control* in Java)
      * e.g., `private`, `protected`, `public`, and package-private.
      * make it package-private => implementation.
      * make it public => must maintain compatibility forever.
    * => overridden methods cannot have a more restrictive access level in the subclass.
    * => should not raise the accessibility any higher than package-private to facilitate testing your code.
    * *module system* => *module declaration*
  * Implementation
    * Instance fields of public classes should rarely be public.
      * public mutable fields => not thread-safe.
      * expose public static final fields is an exception.
    * It is wrong to have a public static final array field, or an accessor that returns such a field.
      * => make the array private, and
        * => add a public immutable list through `Collections.unmodifiableList`.
        * => or, add a public method tha returns a copy of private array.

* In public classes, **use accessor methods**, not public fields

  * Consequences
    * => *accessor methods* for private fields and *mutators* for mutable classes.
    * If a class is package-private or private nested class, there is nothing inherently wrong with exposing its data field.
    * It is questionable for public classes to expose immutable fields.
      * => it precludes changing the internal representation in a later release.

* Minimize **mutability**

  * Applicability
    * Classes should be immutable *unless* there's a very good reason to make them mutable.
    * If a class cannot be made immutable, limit its mutability *as much as possible*.
    * Constructors should create fully initialized objects with all of their invariants established.
  * Consequences
    * => information contained in each instance is fixed for its lifetime.
    * => simple => clear state space, precise state transitions.
    * => inherently thread-safe; require no synchronization.
    * => can be shared freely => never have to make *defensive copies* => need not provide a `clone` method or *copy constructor*.
    * => they can share their internals.
    * => make greater building blocks for other objects.
    * => provide failure atomicity for free.
    * => they require a separate object for each distinct value => extra cost.
  * Implementation
    * To make a class immutable
      * => Don't provide methods that modify the object's state (*mutators*).
      * => Ensure that the class can't be extended.
      * => Make all fields final.
      * => Make all fields private.
      * => Ensure exclusive access to any mutable components.
        * Make *defensive copies* in constructors, accessors, and `readObject` methods.
    * To fix performance problem
      * => guess multi-step operations will be commonly required and to provide them as primitives.
      * => provide a *public* mutable companion class.
    * To make immutability flexible
      * => make all of its constructors private or package-private and add public static factories in place of the public constructors.
        * => effectively final outside its package.
        * => allow to tune the performance through object-caching.
      * => only guarantee that no method may produce an *externally visible* change in the object's state.
        * e.g., have nonfinal fields caching the result of expensive computations the first time they are needed.

* **Favor composition over inheritance**

  * Motivation

    * inheritance
      * => achieve code reuse, also override methods.
      *  => a subclass depends on the implementation details of its superclass for its proper function => violates encapsulation
        * superclass can acquire new methods in subsequent releases.
      * => dangerous to inherit across package boundaries.
      * => propagate any flaws in the superclass's API.
    * *composition*: declare a private field that references an instance of the existing class.

  * Consequences

    * => *forwarding* => no dependencies on the implementation of the existing class.
      * => write forwarding methods only once.

    * the combination of composition and forwarding => *delegation*.
    * => not suited for use in *callback frameworks*.
      * wrapper class objects pass self-references (`this`) to other objects for subsequent callback invocations and it doesn't know of its wrapper => *SELF problem*.

  * Implementation

    * Reusable forwarding class + Wrapper class extends forwarding class.

* **Design and document for inheritance or else prohibit it**

  * Implementation
    * First, the class must document precisely the effects of overriding any method.
      * The class must document its *self-use* of overridable methods.
      * For each public or protected method, the documentation must indicate which overridable methods the method invokes, in what sequences, and how the results of each invocation affect subsequent processing.
      * `@implSpec`: inner working of the method.
    * To allow efficient subclasses, a class may have to provide hooks into its internal workings in the form of judiciously chosen protected methods.
      * e.g., `removeRange` method from `java.util.AbstractList` => provided solely to make it easy for subclasses to provide a fast helper.
    * Constructors must not invoke overridable methods, directly or indirectly.
      * The superclass constructor runs before the subclass constructor => the overriding method in the subclass will get invoked before the subclass constructor has run.
      * Neither `clone` nor `readObject` may invoke an overridable method, directly or indirectly.
      * When implementing `Serializable` => make `readResolve` or `writeReplace` method protected => avoid them to be ignored by subclasses.
    * Eliminate a class's self-use of overridable methods mechanically.
      * => move the body to a private helper method.
  * Consequences
    * => The *only* way to test a class designed for inheritance is to write subclasses.
    * => You must test your class by writing subclasses *before* you release it.
    * => The best solution is to prohibit subclassing in classes that are not designed and documented to be safely subclassed.
      * => declare the class final.
      * => make all the constructors private or package-private and to add public static factories in place of the constructors.

* **Prefer interfaces to abstract classes**

  * Motivation
    * interfaces *versus* abstract classes.
      * both can provide implementations for instance methods.
      * abstract class => single inheritance => must place common abstract class high up => damage to the type hierarchy.
      * Existing classes can easily be retrofitted to implement a new interface.
      * Interfaces are ideal for defining mixins.
        * => provide optional behavior.
      * Interfaces allow for the construction of nonhierarchical type frameworks.
      * Interfaces enable safe, powerful functionality enhancements via the *wrapper class* idiom.
  * Consequences
    * abstract *skeletal implementation* class
      * => combine the advantages of interfaces and abstract classes .
      * => the interface defines the type, providing default methods.
      * => the skeletal implementation class implements the remaining non-primitive interface mthods atop the primitive interface methods.
      * => extending a skeletal implementation.
      * skeletal implementation classes are called `AbstractInterface`.
    * *simple implementation*
      * it isn't abstract; it is the simplest possible working implementation.
      * e.g., `AbstractMap.SimpleEntry`.

* **Design interfaces for posterity**

  * Consequences
    * default method => a *default implementation* can be used by all classes that implement the interface.
    * => not always possible to write a default method that maintains all invariants of every conceivable implementation.
    * => existing implementations of an interface may compile without error or warning but fail at runtime.
  * Applicability
    * It is of the utmost importance to design interfaces with great care.
    * You cannot count on correcting interface flaws after an interface is released.

* **Use interfaces only to define types**

  * Applicability
    * the interface serves as a *type* that can be used to refer to instances of the class.
  * Consequences
    * *constant interface* antipattern
      * => consists solely of static final fields, each exporting a constant.
      * => leak implementation detail into the class's exported API.
      * => if the constants are strongly tied to an existing class or interface, you should add them to the class or interface.
        * => export them with an *enum type*.
        * => export the constants with a noninstantiable *utility class*,
          * => use *static import* facility.

* **Prefer class hierarchies to tagged classes**

  * Motivation
    * tagged classes: contain a *tag* field indicating the flavor of the instance.
      * => verbose, error-prone, and inefficient.
  * Implementation
    * First, define an abstract class containing an abstract method for each method in the tagged class whose behavior depends on the tag vlaue.
    * Next, define a concrete subclass of the root class for each flavof of the original tagged class.
    * Also include in each subclass the appropriate implementation of each abstract method in the root class.

* **Favor static member classes over nonstatic**

  * Applicability
    * A nested class should exist only to serve its enclosing class.
    * If you declare a member class that does not require access to an enclosing instance, *always* put the `static` modifier in its declaration.
  * Consequences
    * *static member classes*
      * => can function as public helper class, useful only in conjunction with its outer class.
        * e.g., Clients of `Calculator` could refer to operations using public static member enum class `Calculator.Operation`.
    * *private static member classes*
      * => can represent components of the object represented by their enclosing classes.
        * e.g., `Map`'s internal `Entry` object.
          * => while each entry is associated with a map, the methods on the entry do not need to access to the map.
    * *nonstatic member class*
      * => implicitly associated with an *enclosing instance* of its containing class => takes up space in the nonstatic member class instance and adds time to its construction.
      * => you can invoke methods on the enclosing instance or obtain a reference to the enclosing instance using the *qualified* `this` construct.
      * => can be used to define an *Adapter*.
        * e.g., `Map`'s *collection views* methods `keySet`, `entrySet`, and `values`.
        * e.g., `Set` and `List` typically use nonstatic classes to implement their iterators.
    * *anonymous class*
      * => has no name.
      * => not a member of its enclosing class.
      * => cannot have any static members other than *constant variables*, which are final primitive or string fields.
      * => can create small *function objects* and *process objects* on the fly.
        * => but lambdas are now preferred.
      * => can use in the implementation of static factory methods.
    * *local class*
      * => declared anywhere a local variable can be declared.
      * => should be kept short so as not to harm readability.

* **Limit source files to a single top-level class**

  * Motivation
    * Defining multiple top-level classes in a source file => possible to provide multiple definitions for a class => affected by the order in which the source files are passed to the compiler.
  * Implementation
    * Use static member classes instead of multiple top-level classes.
  * Applicability
    * Never put multiple top-level classes or interfaces in a single source file.

## Generics

* **Don't use raw types**
  * Applicability
    * *generic*: a class or interface whose declaration has one or more *type parameters*.
    * *parameterized types*: e.g., `List<String>`: a list of parameterized type `String`.
    * *ray type*: the name of the generic type used without accompanying type parameters.
  * Consequences
    * Each generic type defines a *raw type*.
      * => behave as if all of the generic type information were erased from the type declaration.
      * => exist primarily for compatibility with pre-generics code.
    * If you use raw types, you lose all the safety and expressiveness benefits of generics.
      * Using raw types can lead to exceptions at runtime; parameterized collection type ensures compile-time check.
  * Implementation
    * To allow insertion of arbitrary objects, use *unbounded wildcard types*.
      * If you want to use a generic type but you don't know or care what the actual type parameter is, use a question mark instead.
      * You can't put any element (other than `null`) into a `Collection<?>`.
      * e.g., `Set<E>` => `Set<?>` containing only objects of some unknown type.
      * use *generic methods* or *bounded wildcard types* if you can assume about the type of the objects.
    * Exceptions to the rule that you should not use raw types.
      * You must use raw types in class literals.
        * e.g., `List.class` instead of `List<String>.class`.
      * It is legal to use the `instanceof` operator on parameterized types.
* **Eliminate unchecked warnings**
  * Applicability
    * Eliminate every unchecked warning that you can.
  * Consequences
    * => ensure that your code is type safe => you won't get a `ClassCastException` at runtime.
  * Implementation
    * If you can prove that the code that provoked the warning is type safe, then (and only then) suppress the warning with an `@SuppressWarnings("unchecked")` annotation.
      * Always use the `SuppressWarnings` annotation on the smallest scope possible.
      * Every time you use a `@SuppressWarnings("unchecked")` annotation, add a comment saying why it is safe to do so.
* **Prefer lists to arrays**
  * Motivation
    * Arrays are *covariant* => `Sub[]` is a subtype of the array `Super[]`.
    * Generics are *invariant* => `List<Type1> is neither a subtype nor a supertype of `List<Type2`.
    * Arrays are *reified*. => arrays know and enforce their element type at runtime.
      * e.g., get an `ArrayStoreException` if putting a `String` into an array of `Long`.
    * Generics are implemented by *erasure* => they enforce their type constraints only at compile time and discard their element type at runtime.
    * => arrays and generics do not mix well.
      * It is illegal to create an array of a generic type, a parameterized type, or a type parameter.
      * e.g., illegal creation expressions: `new List<E>[]`, `new List<String>[]`, `new E[]`.
  * Consequences
    * use the collection type `List<E>` in preference to the array type `E[]`.
      * => might sacrifice some conciseness or performance, in exchange for better type safety and interoperability.
* **Favor generic types**
  * Motivation
    * We can often *generify* programs without harming clients of the original non-parameterized version.
    * Generic types are safer and easier to use than types that require casts in client code.
  * Implementation
    * some technique for eliminating the generic array creation may cause *heap pollution* => the rumtime type of the array does not match its compile-time type.
    * some generic types that restrict the permissible values of their type parameters => *bounded type parameter*.
      * e.g., `class DelayQueue<E extends Delayed> implements BlockingQueue<E>`
* **Favor generic methods**
  * Consequences
    * declare a *type parameter* => make the method typesafe.
  * Implementation
    * the type parameter list goes between a method's modifiers and its return type.
      * e.g., `public static <E> Set<E> union(Set<E> s1,  Set<E> s2);`.
    * *generic singleton factory*: a static factory method to repeatedly dole out the object for each requested type parameterization.
      * e.g., `Collections.reverseOrder`, `Collections.emptySet`.
    * *recursive type bound* => wildcard variant, the *simulated self-type* idiom.
      * e.g., `public static <E extends Comparable<E>> E max(Collection<E> c);`.
* Use **bounded wildcards** to increase API flexibility
  * Consequences
    * PECS stands for producer-`extends`, consumer-`super`.
    * Do not use bounded wildcard types as return types.
  * Implementation
    * *bounded wildcard type*
      * `<? extends E>`: wildcard type for a parameter that serves as an `E` producer.
      * `<? super E>`: wildcard type for a parameter that serves as an `E` consumer.
    * Comparables are always consumers.
      * generally use `Comparable<? super T>` and `Comparator<? super T>`.
      * e.g., `public static <T extends Comparable<? super T>> T max(List<? extends T> list);`.
    * If a type parameter appears only once in a method declaration, replace it with a wildcard.
    * Write a private helper method to *capture* the wildcard type.
* **Combine generics and varargs judiciously**
  * Motivation
    * varargs => *leasy abstraction* => an array is created to hold the varargs parameters and it is visible.
    * If a method declares its varargs parameter a generic or parameterized types => warning on the declaration => possible *heap pollution*.
  * Consequences
    * It is unsafe to store a value in a generic varargs array parameter.
    * It is unsafe to give another method access to a generic varargs parameter array.
    * Use `@SafeVarargs` on every method with a varargs parameters of a generic or parameterized type => *never* write unsafe varargs methods.
    * An alternative is to replace the varargs parameter with a `List` parameter.
      * => the compiler can *prove* that the method is typesafe.
      * => the author does not have to vouch for its safety with a `SafeVarargs` annotation.
      * => the client code is a bit verbose and may be a bit slower.
  * Implementation
    * `@SafeVarargs` annotation => the author promise it is typesafe => allow the author of a method with a generic varargs parameter to suppress client warnings automatically.
      * If the method doesn't store anything into the array and doesn't allow a reference to the array to escape, then it's safe.
* **Consider typesafe heterogeneous containers**
  * Motivation
    * Limitations to fixed numbers of type parameters per container.
  * Consequences
    * *type token*: a class literal passed among methods to communicate both compile-time and runtime type information.
    * *bounded type token*: a type token that places a bound on type.
    * placing the type parameter on the key rather than the container => customer key type => unfixed number of type parameters per contains.
  * Implementation
    * typesafe heterogeneous container => e.g., `public <T> void putFavorite(Class<T> type, T instance);`
    * dynamic cast => e.g., `favorites.put(Objects.requireNonNull(type), type.cast(instance));`
