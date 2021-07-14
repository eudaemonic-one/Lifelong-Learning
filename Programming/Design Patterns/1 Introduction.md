# Chapter 1. Introduction

* “One thing expert designers know not to do is solve every problem from first principles. Rather, they reuse solutions that have worked for them in the past.”
* “The purpose of this book is to record experience in designing object-oriented software as **design patterns**. Each design pattern systematically names, explains, and evaluates an important and recurring design in object-oriented systems.”


## 1.1 What Is a Design Pattern?

* pattern name: a handle we can use to describe a design problem.
* problem: when to apply the pattern.
* solution: elements that make up the design, thier relationships, responsibilities, and collaborations.
* Consequence: results and tradeoffs of applying the pattern.

## 1.2 Design Patterns in Smalltalk MVC

* Model/View/Controller (MVC)
* MVC decouples views and models by establishing a subscribe/notify protocol between them.
  * “A view must ensure that its appearance reflects the state of the model.”
  * “Whenever the model’s data changes, the model notifies views that depend on it.”
  * “In response, each view gets an opportunity to update itself. ”
* Views can be nested.
* MVC let you change the way a view responds to user input without changing its visual presentation.
* The View-Controller relationship is an example of the Strategy design pattern.
  * A Strategy is an object that represents an algorithm.
* MVC uses other design patterns, such as Factory Method to specify the default controller class for a view and Decorator to add scrolling to a view.
  * But the main relationships in MVC are given by the Observer, Composite, and Strategy design patterns.

## 1.3 Describing Design Patterns

* Pattern Name and Classification: essence and scheme.
* Intent: rationale and what design issue to address.
* Also Known As: other names.
* Motivation: a scenario that illustrate a design problem.
* Applicability: where to apply.
* Structure: a graphical representation.
* Participants: classes and/or objects.
* Collaborations: how participants carry out their responsibility.
* Consequences: trade-offs and results of using the pattern.
* Implementation: pitfalls, hints, or techniques of implementing the pattern.
* Sample Code: code fragments.
* Known Uses: examples found in real systems.
* Related Patterns: closedly related patterns.

## 1.4 The Catalog of Design Patterns

* **Abstract Factory**: Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
* **Adapter**: Convert the interface of a class into another interface clients expect. Adapter letsc classes work together that couldn't otherwise because of incompatible interfaces.
* **Bridge**: Decouple an abstraction from its implementation so that the two can vary independently.
* **Builder**: Separate the construction of a complex object from its representation so that the same construction process can create different representations.
* **Chain of Responsibility**: Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it.
* **Command**: Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.
* **Composite**: Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.
* **Decorator**: Attach additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.
* **Facade**: Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.
* **Factory Method**: Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.
* **Flyweight**: Use sharing to support large numbers of fine-grained objects efficiently.
* **Interpreter**: Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentences in the language.
* **Iterator**: Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
* **Mediator**: Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each otehr explicitly, and it lets you vary their interaction independently.
* **Memento**: Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.
* **Observer**: Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
* **Prototype**: Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.
* **Proxy**: Provide a surrogate or placeholder for another object to control access to it.
* **Singleton**: Ensure a class only has one instance, and provide a global point of access to it.
* **State**: Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.
* **Strategy**: Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.
* **Template Method**: Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.
* **Visitor**: Represent an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.

## 1.5 Organizing the Catalog

|           |            | Purpose          |                  |                         |
| --------- | ---------- | ---------------- | ---------------- | ----------------------- |
|           |            | **Creational**   | **Structural**   | **Behavioral**          |
| **Scope** | **Class**  | Factory Method   | Adapter (class)  | Interpreter             |
|           |            |                  |                  | Template Method         |
|           | **Object** | Abstract Factory | Adapter (object) | Chain of Responsibility |
|           |            | Builder          | Bridge           | Command                 |
|           |            | Prototype        | Composite        | Iterator                |
|           |            | Singleton        | Decorator        | Mediator                |
|           |            |                  | Facade           | Memento                 |
|           |            |                  | Flyweight        | Observer                |
|           |            |                  | Proxy            | State                   |
|           |            |                  |                  | Strategy                |
|           |            |                  |                  | Visitor                 |

* “Class patterns deal with relationships between classes and their subclasses. These relationships are established through inheritance, so they are static—fixed at compile-time.”
* “Object patterns deal with object relationships, which can be changed at run-time and are more dynamic.”
* “Creational class patterns defer some part of object creation to subclasses, while Creational object patterns defer it to another object.”
* “The Structural class patterns use inheritance to compose classes, while the Structural object patterns describe ways to assemble objects.”
* “The Behavioral class patterns use inheritance to describe algorithms and flow of control, whereas the Behavioral object patterns describe how a group of objects cooperate to perform a task that no single object can carry out alone.”
* “Clearly there are many ways to organize design patterns.”
  * “Some patterns are often used together.”
  * “Some patterns are alternatives.”
  * “Some patterns result in similar designs even though the patterns have different intents.”

## 1.6 How Design Patterns Solve Design Problems

* **Finding Appropriate Objects**
  * Decomposing a system into objects -> encapsulation, granularity, dependency, flexibility, performance, evolution, reusability, and on and on.
  * No counterparts in the real world -> flexible design.
  * Design patterns -> identify less-obvious abstractions -> more flexible and resusable.
* **Determining Object Granularity**
  * How do we decide what should be an object?
* **Specifying Object Interfaces**
  * Interface: the complete set of requests that can be sent to the object.
  * “We speak of an object as having the type “Window” if it accepts all requests for the operations defined in the interface named “Window.”

  * “An object may have many types, and widely different objects can share a type.”
  * Dynamic Binding: the run-time association of a request to an object.
  * Polymorphism: can substitute objects that have identical interfaces for each other at run-time.
  * Design patterns -> define interfaces by identifying their key elements and the kinds of data that get sent across an interface -> specify relationships between interfaces.
* **Specifying Object Implementation**
  * OMT-based notation
    * Class: rectangle with class name in the bold, operations, data comes after the operations. Return types and instance variable types are optional.
    * Instantiates: a dashed arrowhead points to the class of the instantiated objects.
    * Inheritance: a vertical line from a subclass and a triangle pointing to a parent class.
      * The names of abstract classes appear in slanted type.
  * “Programming to an Interface, not an Implementation”
    * “Class inheritance defines an object’s implementation in terms of another object’s implementation. In short, it’s a mechanism for code and representation sharing. In contrast, interface inheritance (or subtyping) describes when an object can be used in place of another.”
  * “Don’t declare variables to be instances of particular concrete classes. Instead, commit only to an interface defined by an abstract class.”

* **Putting Reuse Mechanisms to Work**
  * Inheritance versus Composition
    * Object Inheritance -> let you define the implementation of one class in terms of another's -> internals of parent classes are often visible to subclasses.
    * Object Composition -> assemble or compose objects to get more complex functionality -> no internal details of objects are visible.
    * Implementation dependencies -> problems when reuse a subclass (must rewrite parent class when inherited implementation not appropriate for new problem domains) -> limit flexibility and reusability -> cure if to inherit only from abstract classes.
    * “**Favor object composition over class inheritance** helps you keep each class encapsulated and focused on one task. Your classes and class hierarchies will remain small and will be less likely to grow into unmanageable monsters.”
  * Delegation
    * A receiving object delegates operations to its delegate.
    * Delegation -> easy to compose behaviors at run-time and to change the way they're composed -> dynamic, highly parameterized software -> efficiency depends on teh context.
  * Inheritance versus Parameterized Types
    * Generics / Templates: Define a type without specifying all other types it uses and the unspecified ones are supplied as parameters at the point of use.
    * Let you change the types that a class can use, not at run-time.
* **Relating Run-Time and Compile-Time Structure**
  * Run-time structure has little resemblance to its code structure.
  * Frozen code structure -> Compile-time.
  * Aggregation: one object owns or is responsible for another object and they have identical lifetimes.
  * Acquaintance: acquainted objects merely request operations of each other and suggest much looser coupling between objects.
* **Designing for Change**
  * “The key to maximizing reuse lies in anticipating new requirements and changes to existing requirements, and in designing your systems so that they can evolve accordingly.”
  * “Those changes might involve class redefinition and reimplementation, client modification, and retesting.”
  * Common causes of redesign:
    * “Creating an object by specifying a class explicitly.”
    * “Dependence on specific operations.”
    * “Dependence on hardware and software platform.”
    * “Dependence on object representations or implementations.”
    * “Algorithm dependencies.”
    * “Tight coupling.”
    * “Extending functionality by subclassing.”
    * “Inability to alter classes conveniently.”
* Design patterns play in the development of three broad classes of software: application programs, toolkits, and frameworks.

## 1.7 How to Select a Design Pattern

* “Consider how design patterns solve design problems.”
* “Scan Intent sections.”
* “Study how patterns interrelate.”
* “Study patterns of like purpose.”
* “Examine a cause of redesign.”
* “Consider what should be variable in your design.”
  * “Instead of considering what might *force* a change to a design, consider what you want to be *able* to change without redesign. ”


| Purpose    | Design Pattern          | Aspect(s) That Can Vary                                      |
| ---------- | ----------------------- | ------------------------------------------------------------ |
| Creational | Abstract Factory        | families of product objects                                  |
|            | Builder                 | how a composite object gets created                          |
|            | Factory Method          | subclass of object that is instantiated                      |
|            | Prototype               | class of object that is instantiated                         |
|            | Singleton               | the sole instance of a class                                 |
| Structural | Adapter                 | interface to an object                                       |
|            | Bridge                  | implementation of an object                                  |
|            | Composite               | structure and composition of an object                       |
|            | Decorator               | responsibilities of an object without subclassing            |
|            | Facade                  | interface to a subsystem                                     |
|            | Flyweight               | storage costs of objects                                     |
|            | Proxy                   | how an object is accessed; its location                      |
| Behavioral | Chain of Responsibility | object that can fulfill a request                            |
|            | Command                 | when and how a request is fulfilled                          |
|            | Interpreter             | grammar and interpretation of a language                     |
|            | Iterator                | how an aggregate's elements are accessed, traversed          |
|            | Mediator                | how and which objects interact with each other               |
|            | Memento                 | what private information is stored outside an object, and when |
|            | Observer                | number of objects that depend on another object; how the dependent objects stay up to date |
|            | State                   | states of an object                                          |
|            | Strategy                | an algorithm                                                 |
|            | Template Method         | steps of an algorithms                                       |
|            | Visitor                 | operations that can be applied to object(s) without changing their class(es) |

