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
