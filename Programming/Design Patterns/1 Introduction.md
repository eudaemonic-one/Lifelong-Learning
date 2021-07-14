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
