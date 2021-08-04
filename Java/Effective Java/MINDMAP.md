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

## Enums and Annotations

* Use **enums** instead of `int` constants
  * Motivation
    * `int` enum pattern => severaly deficient.
      * => no type safety, little expressive power.
      * => clients must recompile if the value associated with an `int` enum is changed.
      * => no way to translate `int` enum constants into printable strings.
  * Consequences
    * *enum type*
      * => instance-controlled => export one instance for each enumeration constant via a public static final field => a generalization of singletons.
      * => flexible to add arbitrary methods and fields and implement arbitrary interfaces.
      * => be their nature immutable.
    * Use enums any time you need a set of constants whose memebers are known at compile time.
  * Implementation
    * To associate data with enum constants, declare instance fields and write a constructor that takes the data and stores it in the fields.
    * Declare an abstract `apply` method in the enum type, and override it with a concrete method for each constant in a *constant-specific class body*.
      * => consider the strategy enum pattern if some, but not all, enum constants share common behaviors.
    * Consider writing a `fromString` method to translate the string representation back to the corresponding enum.
* Use instance fields instead of **ordinals**
  * Motivation
    * `ordinal` method => returns the numerical position of each enum constant in its type.
    * Abuse of ordinal to derive an associated value.
  * Consequences
    * Never derive a value associated with an enum from its ordinal; store it in an instance field instead.
    * Most programmers will have no use for the `ordinal` method.
* Use **`EnumSet`** instead of bit fields
  * Motivation
    * *bit field enumeration constants* - OBSOLETE!
      * => allow bitwise `OR` operation to combine constants into a set.
      * => hard to interpret a bit field and iterate over all of the elements.
      * => you have to predict the maximum number of bits and choose a proper type for the bit field.
  * Consequences
    * `EnumSet`
      * => conciseness, performance.
      * => a rich set of static factories for easy set creation => e.g., `EnumSet.of`.
  * Implementation
    * Take a `Set` rather than an `EnumSet` => accept the interface type rather than the implementation type.
* Use **`EnumMap`** instead of ordinal indexing
  * Consequences
    * Use `EnumMap` to associate data with an enum.
      * => very fast `Map` implementation designed for use with enum keys.
    * Do not use the `ordinal` method to index array of arrays.
  * Implementation
    * e.g., `Map<Plant.LifeCycle, Set<Plant>> plantsByLifeCycle = new EnumMap<>(Plant.LifeCycle.class);`.
    * the `EnumMap` constructor takes the `Class` object of the key type: a *bounded type token*, which provides runtime generic type information.
    * Use a stream and an `EnumMap` to associate data with an enum.
      * e.g., `Arrays.stream(garden).collect(groupingBy(p -> p.lifeCycle, () -> new EnumMap<>(LifeCycle.class), toSet()))`.
    * Use a nested `EnumMap` to associate data with enum pairs.
* **Emulate extensible enums with interfaces**
  * Motivation
    * It is confusing that elements of an extension type are instances of the base type and not vice versa => no good way to enumerate over all of a base type and its extensions.
    * *operation codes* (*opcodes*) => extensible enumerated types.
  * Consequences
    * Emulate extensible enum type by writing an interface.
      * => clients implement the interface to extend their own enums.
      * => implementations cannot be inherited from one enum type to another.
      * => encapsulate shared functionality in a helper class or a static helper method.
    * Implementations
      * `<T extends Enum<T> & Operation>` ensures that the `Class` object (`Class<T>`) represents both an enum and a subtype of `Operation`.
* Prefer **annotations** to naming patterns
  * Motivation
    * *naming patterns* => indicate some program elements demanded special treatment.
      * e.g., JUnit testing framework requires test methods by begining their names with characters `test`.
      * => typographical errors result in silent failures.
      * => no way to ensure that they are used only on appropriate program elements.
      * => provide no good way to associate parameter values with program elements.
  * Consequences
    * Annotations
      * *marker annotation*: has no parameters but simply marks the annotated element.
      * class literals can be used as the values for annotation parameter.
      * *repeatable annotation type*
    * There is simply no reason to use naming patterns when you can use annotations instead.
    * All programmers should use the predefined annotation types that Java provides.
* Consistently use the **`Override`** annotation
  * Consequences
    * Use the `Override` annotation on every method declaration that you believe to override a superclass declaration.
    * It is good practice to use `Override` on concrete implementations of interface methods to ensure that the signature is correct.
    * In concrete classes, you need not annotate methods that you believe to override abstract method declarations.
* Use **marker interfaces** to define types
  * Motivation
    * *marker interface*: contains no method declarations but merely designates a class that implements the interface as having some property.
  * Consequences
    * marker interfaces
      * => define a type that is implemented by instance of the marked class
        * => catch errors at compile time although some APIs do not take advantage of the interface => e.g., `ObjectOutputStream.write` takes type `Object` instead of `Serializable`.
      * => they can be targeted more precisely.
    * marker annotations
      * => they are part of the larger annotation facility => consistency in annotation-based frameworks.
      * => can be applied to any program element.

## Lambdas and Streams

* Prefer **lambdas** to anonymous classes
  * Motivation
    * *function objects*: represent functions or actions.
    * *anonymous class*: creates a function object.
      * => adequate for *Strategy* pattern.
    * *functional interfaces*: with a single abstract method, deserve special treatment.
      * => allows to create instances of these interfaces using *lambda expression*.
  * Consequences
    * Omit the types of all lambda parameters unless their presence makes your program clearer.
    * Lambdas lack names and documentation; if a computation isn't self-explanatory, or exceeds a few lines, don't put it in a lambda.
    * You should rarely, if ever, serialize a lambda.
      * => instead, using an instance of a private static nested class.
    * Don't use anonymous classes for function objects unless you have to create instances of types that aren't functional interfaces.
* Prefer **method references** to lambdas
  * Consequences
    * *method references* => more succinct than lambdas.
    * Where method references are shorter and clearer, use them; where they aren't, stick with lambdas equivalent.
  * Implementation
    * e.g., `map.merge(key, 1, Integer::sum);` rather than `map.merge(key, 1, (count, incr) -> count + incr);`
    * Method Reference Type: Static, Bound, Unbound, Class Constructor, Array Constructor.
* Favor the use of **standard functional interfaces**
  * Consequences
    * `java.util.function` package provides a large collection of standard functional interfaces for your use.
    * If one of the standard functional interfaces does the job, you should generally use it in preference to a purpose-built functional interface.
    * Don't be tempted to use basic functional interfaces with boxed primitives instead of primitive functional interfaces.
    * You should seriously consider writing a purpose-built functional interface if
      * It will be commonly used and could benefit from a descriptive name.
      * It has a strong contract associated with it.
      * It would benefit from custom default methods.
    * Always annotate your functional interfaces with the `@FunctionalInterface` annotation.
  * Implementations
    * The six basic functional interfaces:
      * `UnaryOperator<T>` => `T apply(T t)`
      * `BinaryOperator<T>` => `T apply(T t1, T t2)`
      * `Predicate<T>` => `boolean test(T t)`
      * `Function<T,R>` => `R apply(T t)`
      * `Supplier<T>` => `T get()`
      * `Consumer<T>` => `void accept(T t)`
    * Three variants of each  of the six basic interfaces to operate on the primitive types `int`, `long`, `double`.
    * Nine variants of the `Function` for use when the result type is primitive or `Object` (Obj) => prefix `Function` with `SrcToResult`.
    * Two-argument versions: `BiPredicate<T,U>`, `BiFunction<T,U,R>`, `BiConsumer<T,U>`.
    * `BiFunction` variants returning the three primitive types: `ToIntBiFunction<T,U>`, `ToLongBiFunction<T,U>`, `ToDoubleBiFunction<T,U>`
    * `BooleanSupplier`
* Use **streams** judiciously
  * Motivation
    * iterative code using code blocks, stream pipelines using function objects.
    * streams API => east the task of performing bulk operations, sequentially or in parallel.
  * Consequences
    * stream pipelines are evaluated *lazily*.
      * => evaluation doesn't start until the terminal operation is invoked.
      * => data elements that aren't required in order to complete the terminal operation are never computed.
    * streams API is *fluent* => allow calls to be chained into a single expression.
    * Overusing streams makes programs hard to read and maintain.
    * In the absence of explicit types, careful naming of lambda parameters is essential to the readability of stream pipelines.
    * Using helper methods is even more important for readability in stream pipelines than in iterative code.
    * Refactor existing code to use streams and use them in new code only where it makes sense to do so.
    * Good matches for using stream technique:
      * Uniformly transform sequences of elements
      * Filter sequences of elements
      * Combine sequences of elements using a single operation
      * Accumulate sequences of elements into a collection, perhaps grouping them by some common attribute
      * Search a sequence of elements for an element satisfying some criterion
  * Implementation
    * common stream source: collections, arrays, files, regular expression pattern matchers, pseudorandom number generators, and other streams.
    * you should refrain from using streams to process `char` values.
* **Prefer side-effect-free functions in streams**
  * Motivation
    * *pure function*: one whose result depends only on its input => streams paradigm.
  * Implementation
    * The `forEach` operation should be used only to report the result of a stream computation, not to perform the computation.
    * Use *collectors* to gather the elements of a stream into a true `Collection`.
      * e.g., `comparing` method takes key extraction function.
      * e.g., `toList()`, `toSet()`, `toCollection(collectionFactory)`.
      * e.g., `toMap(keyMapper, valueMapper)` and three-argument form supports dealing with key collisions.
      * e.g., `groupingBy` returns collectors to produce maps that group elements into categories based on a *classifier function*.
      * *downstream collector*: produces a value from a stream containing all the elements in a category => e.g., `toSet()` results in sets rather than lists as the values in a map.
      * e.g., `joining` returns a collector that simply concatenates the elements.
* **Prefer Collection to Stream as a return type**
  * Motivation
    * `Stream` fails to extend` Iterable` => workaround to iterate over a stream is to adapt from `Stream<E>` to `Iterable<E>`.
  * Consequences
    * `Collection` interface is a subtype of `Iterable` and has a `stream` method => provides both iteration abd stream access.
    * `Collection` or an appropriate subtype is generally the best return type fore a public, sequence-returning method.
      * Arrays also provide `Array.asList` and `Stream.of` methods.
    * Do not store a large sequence in memory just to return it as a collection.
      * `Collection` has an `int`-returning `size` method, which limits the length of the returned sequence to `Integer.MAX_VALUE`.
      * => consider implementing a custom collection.
* Use caution when **making streams parallel**
  * Consequences
    * Do not parallel stream pipelines indiscriminately.
      * Parallelizing a pipeline is unlikely to increase its performance if the source is from `Stream.iterate`, or the intermediate operation `limit` is used.
    * Performance gains from parallelism are best on streams over `ArrayList`, `HashMap`, `HashSet`, and `ConcurrentHashMap` instances; arrays; `int` ranges; and `long` ranges.
      * they can be accurately and cheaply split into subranges of any desired sizes => *spliterator* => `Stream.spliterator` or `Iterable.spliterator`.
      * they provide good-to-excellent *locality of reference* when processed sequentially => sequential element references are stored together in memory.
    * The stream pipeline's terminal operation affects the effectiveness of parallel execution.
    * Not only can parallelizing a stream lead to poor performance, including liveness failures; it can lead to incorrect results and unpredictable behavior (*safety failures*).
      * Override the `spliterator` method and test the performance extensively.
    * Parallelzing a stream => strictly a performance optimization.
      * It is possible to achieve near-linear speedup in the number of processor cores.
  * Implementation
    * The best terminal operations for parralelism are *reductions*: `Stream`'s `reduce`, or prepackaged `min`, `max`, `count`, and `sum`, or *short-circuiting* operations `anyMatch`, `allMatch`, and `noneMatch`.
    * If you are going to parallelize a stream of random numbers, start with a `SplittableRandom` instance.

## Methods

* **Check parameters for validity**
  * Consequences
    * check invalid parameter values => fail quickly and cleanly with an appropriate exception.
  * Implementation
    * For public and protected methods, use the Javadoc `@throws` tag to document the exception that will be thrown if a restriction on parameter values is violated.
      * e.g., `IllegalArgumentException`, `IndexOutOfBoundsException`, `NullPointerException`.
    * The class-level comment applies to all parameters in all of the class's public method.
    * The `Objects.requireNonNull` method is flexible and convenient.
    * Non public methods can check their parameters using *assertion* => throw `AssertionError` if they fail.
* Make **defensive copies** when needed
  * Consequences
    * => make a *defensive copy* of each mutable parameter to the method or constructor => use the copy in place.
    * => prevent from violating the invariants.
    * => performance penalty.
  * Implementation
    * Defensive copies are made *before* checking the validity of the parameters, and the validity check is performed on the copies rather than on the originals => prevent from *time-of-check*/*time-of-use* attack.
    * Return defensive copies of mutable internal fields.
    * Do not use the `clone` method to make a defensive copy of a parameter whose type is subclassable by untrusted parties.
* Design **method signatures** carefully
  * Consequences
    * Choose method names carefully.
      * => obey naming conventions.
      * => be consistent with the broader consensus.
      * => avoid long method names.
    * Don't go overboard in providing convenience methods.
      * => make a class difficult to learn, use, document, test, and maintain.
      * When in doubt, leave it out.
    * Avoid long parameter lists.
      * Aim for four parameters or fewer.
      * Long sequences of identically typed parameters are especially harmful.
      * => break the method up into multiple methods.
      * => create *helper classes* to hold groups of parameters.
      * => adapt the Builder pattern from object construction to method invocation.
    * For parameter types, favor interfaces over classes.
    * Prefer two-element enum types to `boolean` parameters.
      * => make code easier to read and write, also easy to add more options.
* Use **overloading** judiciously
  * Motivation
    * The choice of which overloading to invoke is made at compile time.
    * Selection among overloaded methods is *static*, while selection among overridden methods is *dynamic*.
  * Consequences
    * Avoid confusing uses of overloading.
    * A safe, conservation policy if never to export two overloadings with the same number of parameters.
    * You can always give methods different names instead of overloading them.
    * Do not overload methods to take different functional interfaces in the same argument position.
    * If you are retrofitting an existing class to implement a new interface, you should ensure that all overloadings behave identically when passed the same parameters.
* Use **varargs** judiciously
  * Motivation
    * *varargs* methods: *variable arity* methods.
      * => firstly creating an array, then putting the argument values into the array, and finally passing the array to the method.
  * Consequences
    * => every invocation causes an array allocation and initialization => performance cost.
  * Implementation
    * To take one or more arguments, declare the method to take two parameters, one normal parameter of the specified type and one varargs parameter of this type.
      * e.g., `static int min(int firstArg, int... remainingArgs);`
* **Return empty collections or arrays, not nulls**
  * Motivation
    * returning `null` for special-case => requires extra code in the client to handle the possibly null return value.
  * Consequences
    * Never return `null` in place of an empty array or collection.
      * => difficult to use and more prone to error, no performance advantages.
  * Implementation
    * Returning the same *immutable* empty collection repeatedly => avoids allocating empty collection that harms performance.
    * Return a zero-length array instead of `null`.
    * Do *not* preallocate the array passed to `toArray` in hopes of improving performance.
* Return **optionals** judiciously
  * Motivation
    * When unable to return a value
      * => throw an exception => exceptions should be reserved for exceptional conditions, and is expensive.
      * => return `null` => clients must contain special-case code.
      * => `Optional<T>` => an immutable container for object reference that is *empty* or *present*.
  * Consequences
    * => similar in spirit to check exceptions, they *force* the user to confront the fact there may be no value returned.
    * => an `Optional` requires allocated and initialized => performance cost.
  * Implementation
    * Never return a `null` value from an `Optional`-returning method.
    * To create optionals => `Optional.empty()`, `Optional.of(value)`, `Optional.ofNullable(value)`.
    * Container types, including collections, maps, streams, arrays, and optionals should not be wrapped in optionals.
    * To choose what action to take if the method can't return a value => `orElse`, `orElseThrow`, `get`.
    * Never return an optional of a boxed primitive type.
      * Instead, use `OptionalInt`, `OptionalLong`, `OptionalDouble`.
    * Never appropriate to use an optional as a key, value, or element in a collection or array.
* Write **doc comments** for all exposed API elements
  * Motivation
    * *Javadoc* generates API documentation automatically from source code => *doc comments* => the API to be usable.
  * Consequences
    * You must precede *every* exported class, interface, constructor, method, and field declaration with a doc comment.
    * The doc comment for a method should describe succinctly the contract between the method and its client.
      * with exception of methods designed for inheritance, the contract should say *what* the method does rather than *how* it does its job.
      * preconditions => `@throws` tags for unchecked exceptions, `@param` for affected parameters.
      * postconditions => after invocation has completed successfully.
      * *side effect* => observable change in state.
    * Doc comments should be readable both in the source code and in the generated documentation.
    * No two members or constructors in a class or interface should have the same summary description.
    * Remember to document thread-safety (level) and serialization (form).
  * Implementation
    * `@param` for every parameter, followed by a noun phrase.
    * `@return` unless the method has a void return type, followed by a noun phrase.
    * `@throws` for every exception thrown by the method, whether checked or unchecked, should consist of the work "if", followed by a clause describing the exception conditions.
    * `{@code}` around the code fragment.
    * `this` refers to the object on which a method is invoked when it is used in the doc comment for an instance method.
    * `@implSpec`  for *self-use patterns* for inheritance.
    * When documenting a generic type or method, be sure to document all type parameters.
    * When documenting an enum type, be sure to document the contants.
    * When documenting an annotation type, be sure to document any members.

## General Programming

* **Minimize the scope of local variables**
  * Consequences
    * => increase the readability and maintainability.
    * => reduce the likelihood of error.
  * Implementation
    * Declare it where it is first used.
    * Nearly every local variable declaration should contain an initializer.
      * One exception concerns `try-catch` statements.
    * Prefer `for` loops to `while` loops.
      * Also works for iterating when you need the iterator.
    * Keep methods small and focused.
* Prefer **for-each** loops to traditional `for` loops
  * Consequences
    * => eliminates chances to use the wrong variable.
    * => no performance penalty.
  * Implementation
    * where you can't use for-each:
      * Destructive filtering => an explicit iterator and call `remove` method, or `Collections.removeIf`.
      * Tranforming => the list iterator or array index to replace the value of an element.
      * Parallel iteration => all iterators or index variables advanced in lockstep.
    * you can iterate every object that implements `Iterable`.
* **Know and use the libraries**
  * Consequences
    * => you take the advantage of the knowledge of the experts who wrote it and the experience of those who used it before you.
    * => you don't have to waste your time writing ad hoc solutions, even if they are marginally related to your work.
    * => their performance tends to improve over time, with no effort on your part.
    * => they tend to gain functionality over time.
    * => you place your code in the mainstream.
  * Implementation
    * It pays to keep abreast of additions in every major release.
    * Every programmar should be familiar with the basics of `java.lang`, `java.util`, and `java.io`, and their subpackages.
* **Avoid `float` and `double` if exact answers are required**
  * Consequences
    * `float` and `double` types are particularly ill-suited for monetary calculations.
    * Use `BigDecimal`, `int`, `long` for monetary calculations.
      * => less convenient than primitive arithmetic type.
      * => slower.
* Prefer primitive types to **boxed primitives**
  * Consequences
    * primitives have only their values, whereas boxed primitives have identities distinct from their values.
    * primitives have only fully functional values, whereas each boxed primitive type has one nonfunctional value, which is `null`.
    * primitives are more time- and space-efficient than boxed primitives.
  * Implementation
    * Applying the `==` operator to boxed primitives is almost always wrong.
      * => use comparator or static compare methods.
    * When you mix primitives and boxed primitives in an operation, the boxed primitive is auto-unboxed.
      * => reduces the verbosity, but not the danger.
      * => unboxing can throw a `NullPointerException`.
* **Avoid strings where other types are more appropriate**
  * Consequences
    * Strings are poor substitutes for
      * other value types.
      * enum types.
      * aggregate types.
      * capabilities.
    * Used inappropriately
      * => more cumbersome, less flexible, slower, more error-prone.
* Beware the performance of **string concatenation**
  * Consequences
    * `+`
      * => convenient for generating a single line of output, or for a small, fixed-size object.
      * strings are immutable => the contents of both are copied => time quadratic => this technique does not scale.
    * `StringBuilder`
      * => store the statement under construction using `append`.
      * => much faster.
* **Refer to objects by their interfaces**
  * Consequences
    * You should favor the use of interfaces over classes to refer to objects.
      * e.g., parameters, return values, fields.
      * => flexible to switch implementations.
    * Refer to an object by a class if no appropriate interface exists.
      * *value classes* such as `String` and `BigInteger`.
      * fundamental types in *class-based framework*.
      * classes that provide extra methods not found in the interface.
* Prefer interfaces to **reflection**
  * Motivation
    * `java.lang.reflect`
      * Given a `Class` object, you can obtain `Constructor`, `Method`, and `Field` instances => let you manipulate their undering counterparts *reflectively*,
      * => You lose all the benefits of compile-time type checking, including exception checking.
      * => The code required to perform reflective access is clumsy and verbose.
      * => Performance suffers.
  * Consequences
    * There are a few sophisticated applications that require reflection.
      * e.g., code analysis tools, dependency injection frameworks.
    * If you have any doubts as to whether your application requires reflection, it probably doesn't.
    * If you must use a class that is unavailable at compile time, you can create instances reflectively and access them normally via their interface or superclass.
* Use **native methods** judiciously
  * Motivation
    * Java Native Interface (JNI) => call *native methods* written in *native programming languages*.
      * => access platform-specific facilities.
      * => seldom necessary.
  * Consequences
    * It is rarely advisable to use native methods for improved performance.
    * A single bug in the native code can corrupt your entire application.
* **Optimize** judiciously
  * Motivation
    * Strive to write good programs rather than fast ones.
      * information hiding => localize design decisions => can be changed in the future.
    * Strive to avoid design decisions that limit performance.
      * e.g., APIs, wire-level protocols, and persistent data formats.
    * Consider the performance consequences of your API design decisions.
      * e.g., making a public type mutable => defensive copying.
      * e.g., using inheritance rather than composition => artificial limits on the performance of the subclass.
      * e.g., using an implementation type rather than interface => ties to a specific implementation.
  * Implementation
    * It is a very bad idea to wrap an API to achieve good performance.
      * good API => good performance.
    * Measure performance before and after each attempted optimization.
      * profiling tools, microbenchmarking framework, performance model.
* Adhere to generally accepted **naming conventions**
  * Motivation
    * *naming conventions*: typographical and grammatical.
  * Implementation
    * Package or Module => `org.junit.jupiter.api`, `com.google.common.collect`.
    * Class or Interface => `Stream`, `FutureTask`, `LinkedHashMap`, `HttpClient`.
    * Method or Field => `remove`, `groupingBy`, `getCrc`.
    * Constant Field => `MIN_VALUE`, `NEGATIVE_INFINITY`.
    * Local Variable => `i`, `denom`, `houseNum`.
    * Type Parameter => `T`, `E`, `K`, `V`, `X`, `R`, `U`, `V`, `T1`, `T2`.
