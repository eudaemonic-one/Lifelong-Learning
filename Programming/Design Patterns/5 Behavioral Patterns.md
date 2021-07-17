# Chapter 5. Behavioral Patterns

* Behavioral patterns: algorithms and the assignment of responsibilities between objects + communication between them.
* Behavioral class patterns: use inheritance to distribute behavioral between classes.
* Behavioral object patterns: use object composition rather than inheritance.
  * How do peer objects know who to cooperate to perform some tasks with?

## Object Behavioral: Chain of Responsibility

* **Intent**
  * Avoid coupling the sender of a request to its receiver by giving more than one object change to handle the request.
  * Chain the receiving objects and pass the request along the chain until an object handles it.
* **Motivation**
  * Decouple the object that initiates the request from the objects that might handle the request.
  * The request has an implicit receiver.
  * Each object on the chain shares a common interface for handling requests and for accessing its **successor** on the chain.
* **Applicability**
  * Use then
    * more than one object may handle a request, and the handler isn't known *a priori*.
    * you want to issue a request to one of several objects without specifying the receiver explicitly.
    * the set of objects that can handle a request should be specified dynamically.
* **Structure**

![pg225fig01](images/5 Behavioral Patterns/pg225fig01.jpg)

![pg225fig02](images/5 Behavioral Patterns/pg225fig02.jpg)

* **Participants**
  * **Handler**
    * defines an interface for handling requests.
    * (optional) implements the successor link.
  * **ConcreteHandler**
    * handles requests it is responsible for.
    * can access its successor.
    * if the ConcreteHandler can handle the request, it does so; otherwise it forwards the request to its successor.
  * **Client**
    * initiates the request to a ConcreteHandler object on the chain.
* **Collaborations**
  * When a client issues a request, the request propagates along the chain until a ConcreteHandler object takes responsibility for handling it.
* **Consequences**
  * Reduced coupling.
    * Keep a single reference to their successor -> simplify object inter-connections.
  * Added flexibility in assigning responsibilities to objects.
    * Add or change responsibilities by adding or changing the chain at run-time.
  * Receipt isn't guaranteed.
    * The chain should be configured properly.
* **Implementation**
  * Implementing the successor chain.
    * Define new links.
    * Use existing links.
  * Connecting successors.
    * The Handler might provide a default implementation to forward the request to the successor -> ConcreteHandler doesn't have to override the operation if not interested in.
  * Representing requests.
    * Hardcode operation invocation -> convenient, safe.
    * A single handler function that takes a request code -> open-ended set of requests.
    * Separate request objects that bundle request parameters -> safer parameter-passing.
      * Define request kinds and parameters by subclassing.
* **Related Patterns**
  * Often applied in conjunction with Composite.