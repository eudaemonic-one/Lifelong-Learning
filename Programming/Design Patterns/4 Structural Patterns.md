# Chapter 4. Structural Patterns

* Structural patterns: how classes and objects are composed to form larger structures.
* Structural class patterns: use inheritance to compose interfaces or implementations.
  * Making independently developed class libraries work together.
* Structural object patterns: describe how to compose objects to realize new functionality.
  * Change the composition at run-time.

## Class, Object Structural: Adapter

* **Intent**
  * Convert the interface of a class into another interface clients expect -> work together with otherwise incompatible interfaces.
* **Also Known As**
  * Wrapper
* **Motivation**
  * Sometimes a class designed for reuse isn't reusable only because its interface doesn't match the domain-specific interface an application requires.
  * Let existing and unrelated classes work in an application that expects a different interface.
* **Applicability**
  * Use when
    * you want to use an existing class, and its interface does not match the one you need.
    * you want to create a resuable class that cooperates with unrelated or unforeseen classes, that is, classes that don't necessarily have compatible interfaces.
    * (object adpater only) you need to use several existing subclasses, but it's impractical to adapt their interface by subclassing every one. An object adapter can adapt the interface of its parent class.
* **Structure**

![pg141fig01](images/4 Structural Patterns/pg141fig01.jpg)

![pg141fig02](images/4 Structural Patterns/pg141fig02.jpg)

* **Participants**
  * **Target**
    * defines the domain-specific interface that Client uses.
  * **Client**
    * collaborates with objects conforming to the Target interface.
  * **Adaptee**
    * deinfes an existing interface that needs adapting.
  * **Adapter**
    * adpats the interface of Adaptee to the Target interface.
* **Collaborations**
  * Clients call operations on an Adapter instance. In turn, the adapter calls Adaptee operations that carry out the request.
* **Consequences**
  * Class adapter:
    * adapts Adaptee to Target by committing to a concrete Adaptee class -> can not adapt a class *and* all its subclasses.
    * lets Adapter override some of Adaptee's behavior.
    * only one object -> no indirection to get to the adaptee.
  * Object adpater:
    * lets a single Adapter work with many Adaptees -> Adaptee itself and all subclasses.
    * harder to override Adaptee behavior.
  * How much adapting does Adapter do?
    * Similarity -> the amount of work.
  * Pluggable adapters.
    * Build interface adaption into classes.
  * Using two-way adapters to provide transparency.
    * An adapted object no longer conforms to the Adaptee interface.
    * The two-way class adapter conforms to both of the adapted classes and can work in either system -> two clients view an object differently ->  transparency.
* **Implementation**
  * Implementing class adapters in C++.
    * Adapter inherits publicly from Target and privately from Adaptee.
  * Pluggable adapters.
    * Find a narrow interface for Adaptee -> three implementation approaches:
      * Using abstract operations.
        * define abstract operations for the narrow Adaptee interface -> let subclasses implement -> subclasses specialize the narrow interface.
      * Using delegate objects.
        * forward requests to a delegate object.
        * subsitute delegate objects to use different strategies.
      * Parameterized adapters.
        * parameterize an adapter with blocks + each block construct supports adaptation to an individual request without subclassing.
* **Related Patterns**
  * Bridge separates an interface from its implementation, while an adapter changes the interface of an existing object.
  * Decorator enhances another object without changing its interface.

## Object Structural: Bridge

* **Intent**
  * Decouple an abstraction from its implementation so that the two can vary independently.
* **Also Known As**
  * Handle/Body
* **Motivation**
  * One abstraction has many implementations -> define the interface to the abstraction -> concrete subclasses implement it in different ways -> bind implementation to the abstraction permanently -> difficult to modify, extend, reuse abstractions and implementations independently.
  * Clients should be able to create an object without committing to a concrete implementation -> putting abstraction and its implementation in separate class hierarchies.
  * The relationship between abstraction and implementation is called a **bridge**.
* **Applicability**
  * Use when
    * you want to avoid a permanent binding between an abstraction and its implementation.
    * both the abstractions and their implementations should be extensible by subclassing.
    * changes in the implementation of an abstraction should have no impact on clients.
    * you have a proliferation of classes within a hierarchy and need to split an object into two parts.
    * you want to share an implementation among multiple objects, and this fact should be hidden from the client.
* **Structure**

![pg153fig01](images/4 Structural Patterns/pg153fig01.jpg)

* **Participants**
  * **Abstraction**
    * defines the abstraction's interface.
  * **RefinedAbstraction**
    * Extends the interface defined by Abstraction.
  * **Implementor**
    * defines the interface for implementation classes.
      * typically primitive operations.
  * **ConcreteImplementor**
    * implements the Implementor interface and defines its concrete implementation.
* **Collaborations**
  * Abstraction forwards client requests to its Implementor object.
* **Consequences**
  * Decoupling interface and implementation.
    * Eliminate compile-time dependencies on the implementation -> even possible for an object to change its implementation at run-time.
  * Improved extensibility.
  * Hiding implementation details from clients.
* **Implementation**
  * Only one Implementor.
    * where there's only one implementation -> degenerate into one-to-one relationship.
  * Creating the right Implementor object.
    * Decide between them based on parameters passed to its constructor.
    * Choose a default implementation initially and change it later according to usage.
    * Delegate the decision to another object altogether.
  * Sharing implementors.
  * Using multiple inheritance.
    * Use C++ to inherit publicly from Abstraction and privately from a ConcreteImplementor.
* **Related Patterns**
  * An Abstract Factory can create and configure a particular Bridge.
  * Bridge is used up-front in a design.

## Object Structural: Composite

* **Intent**
  * Compose objects into tree structures to represent part-whole hierarchies -> treat individual objects and compositions of objects uniformly.
* **Motivation**
  * Distinguish objects -> complex.
  * Composite pattern: recursive composition -> no distinction.
    * abstract class == primitives + containers.
* **Applicability**
  * Use when
    * you want to represent part-whole hierarchies of objects.
    * you want clients to be able to ignore the difference between compositions of objects and individual objects.
* **Structure**

![pg164fig02](images/4 Structural Patterns/pg164fig02.jpg)

![pg165fig01](images/4 Structural Patterns/pg165fig01.jpg)

* **Participants**
  * **Component**
    * declares the interface for objects in the composition.
    * implements default behavior for the interface common to all classes, as appropriate.
    * declares an interface for accessing and managing its child components.
    * (optional) defines an interface for accessing a component's parent in the recursive structure, and implements it if that's appropriate.
  * **Leaf**
    * represents leaf objects in the composition that have no children.
    * Defines behavior for primitive objects in the composition.
  * **Composite**
    * defines behavior for components having children.
    * stores child components.
    * implements child-related operations in the Component interface.
  * **Client**
    * manipulates objects in the composition through the Component interface.
* **Collaborations**
  * Clients use the Component class interface to interact with objects in the compositie structure.
    * If the recipient is a Leaf, then the request is handled directly.
    * If the recipient is a Composite, then it usually forwards requests to its child components, possibly performing addtional operations before and/or after forwarding.
* **Consequences**
  * defines class hierarchies consisting of primitive objects and composite objects.
  * makes the client simple.
  * makes it easier to add new kinds of components.
  * can make your design overly general.
    * can not restrict the components of a composite -> can not rely on the type system to enforce constraints -> run-time checks needed.
* **Implementation**
  * Explicit parent references.
    * Maintain references from child to parent -> simplify the traversal and management of a composite structure.
    * Parent references help support the Chain of Responsibility pattern.
    * Define the parent reference in the Component class.
    * Maintain the invariant that all children of a composite have as their parent the composite that in turn has them as children.
    * Ensure to change a component's parent *only* when it's being added or removed from a composite.
  * Sharing components.
    * The Flyweight pattern solves the problem.
  * Maximizing the Component interface.
    * Define as many common operations for Composite and Leaf classes as possible.
  * Declaring the child management operations.
    * Transparency > safety.
  * Should Component implement a list of Components?
    * Storing child incurs a space penalty.
  * Child ordering.
    * The Iterator pattern can guide you in managing the sequence of children.
  * Caching to improve performance.
    * Cache traversal or search information about children for efficiency.
    * Changes to a component -> invalidate the caches of its parent.
  * Who should delete components?
    * If without garbage collection -> Composite responsible for deleting its children when it's destroyed.
    * Leaf objects are immutable and thus can be shared.
  * What's the best data structure for storing components?
    * Not necessary use a general-purpose data structure.
* **Related Patterns**
  * Often the component-parent link is used for a Chain of Responsibility.
  * Decorator: often used with Composite, often have a common parent class if used together.
  * Flyweight lets you share components, but they can no longer refer to parents.
