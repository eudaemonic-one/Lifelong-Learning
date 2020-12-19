# Lecture 17 Designing Class APIs

## Class Design Principles

* In public classes, prefer accessors to public fields
  * Cleaner
    * uniform access principle says user should be unaware of computation vs. storage
  * More flexible (information hiding)
    * Can eliminate private field & compute result in later release
    * Can compute lazily to speed class initialization
  * More powerful
    * Can take action (e.g., synchronization) on access
  * Only valid use for public fields: true constants
    * Typically primitive
      * e.g., `Math.PI`, `Integer.MAX_VALUE`
    * Immutable reference types also fine
      * e.g., `BigInteger.TWO`
    * Functional interfaces (e.g., `Comparator`) are not fine
* Don't expose a new type that lacks meaningful contractual refinements on an existing super type
  * Just use the supertype
  * Reduces conceptual surface area
  * Increases flexbility
  * Resist the urge to expose type just because it's there
* Classes should be immutable unless there's a good reason to make them mutable
  * Advantages: simple, thread-safe, reusable
  * Disadvantages: separate object for each value
  * If mutable, keep state-space small, well-defined
* Prefer interfaces to abstract classes
  * Client can implement multiple interfaces
  * Client can implement interfaces & choose superclass
  * You can implement an interface without accepting any state or implementation code
  * You can still provide implementation assistance
* When choosing abstractions, favor well-defined mathematical entities
  * e.g., `BigDecimal` rounding modes
  * Leverages unreasonable effectiveness of math
    * Places API on sound theoretical footing
  * Your user needn't be aware to benefit
    * Don't inflict mathematical terminology on a lay audience
  * Bolsters claim that math skills aid in API design
* Avoid reliance on extralinguistic mechanisms
  * e.g., reflection, cloning, native methods, bytecode rewriting
  * Resulting APIs tend to be unsafe, brittle
    * You lose basic guarantees made by the language
  * Exhibit A" Java serialization
    * API cannot be implemented without magic
    * Objects created without calling constructors
      * Establishing invariants in constructors no longer sufficient
    * Private representation becomes public API
      * Information hiding goes out the window
  * Exhibit B: `numpy.empty`
* Design & document for inheritaance or else prohibit it
  * Inheritance violates encapsulation (Snyder, '86)
  * If you allow subclassing, document self-use
  * Conservative policy: all concrete classes uninheritable
* Subclass only when an is-a relationship exists
  * Subclassing implies substitutability (Liskov)
    * Makes it possible to pass an instance of subclass wherever superclass is called for
  * If not is-a but you subclass anyway, all hell breaks loose
  * Never subclass just to reuse implementation
  * Ask yourself "Is every `Foo` really a `Bar`?"
* Don't put too many methods in a class
  * Too many methods make class difficult to learn, use
    * A dozen or fewer methods is ideal
  * Most important for core interfaces
    * Good: `Set`, `List`, `Map`
    * Bad: `Stream.Collectors`, `ByteBuffer`
  * Less important for static utilities

## Function Interface Design

* Criteria for writing a purpose-built functional interface
  * Likely to be commonly used
  * Has a good descriptive name
  * Has a strong contract associated with it
  * Would benefit from default methods