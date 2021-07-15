# Chapter 3. Creational Patterns

* “Creational design patterns abstract the instantiation process. ”
  * Independent of how its objects are created, composed, and represented.
  * Class creational pattern uses inheritance to vary the class that's instantiated.
  * Object creational pattern delegates instantiation to another object.
* Creational design patterns depend more on object composition than class inheritance.
* Creational patterns -> flexibility for *what* gets created, *who* creates it, *how* it gets created, and *when* -> objects vary widely in structure and functionality.
* Relationship among creational patterns:
  * Competitors: Either Prototype or Abstract Factory
  * Complementary: Builder can use one of the other patterns to implement which components get built, Prototype can use Singleton in its implementation.
* This chapter uses a common example - building a maze for a computer game - to illustrate their implementations.
  * Don't want hard-coding maze layout.
  * Changing the layout -> override and reimplement member function -> error-prone + doesn't promote reuse.
  * Creational patterns -> more flexible -> easy to change components of a maze.

![pg82fig01](images/3 Creational Patterns/pg82fig01.jpg)

## Object Creational: Abstract Factory

* **Intent**
  * Provide an interface for creating families of related or dependent objects without specifying their concret classes.
* **Also Known As**
  * Kit
* **Motivation**
  * Hard-code widgets for particular standard -> not portable -> hard to change the standard in the future.
  * Abstract factory class creates each basic kind of widget -> solve this problem.
    * Concrete subclass of abstract factory class for each standard and each subclass implements the operations.
    * Clients can ignore the classes implementing widgets for a particular standard.
* **Applicability**
  * Use when
    * “a system should be independent of how its products are created, composed, and represented.”
    * “a system should be configured with one of multiple families of products.”
    * “a family of related product objects is designed to be used together, and you need to enforce this constraint.”
    * “you want to provide a class library of products, and you want to reveal just their interfaces, not their implementations.”
* **Structure**

![pg88fig01](images/3 Creational Patterns/pg88fig01.jpg)

* **Participants**
  * **AbstractFactory**
    * declares an interface for operations that create abstract product objects.
  * **ConcreteFactory**
    * Implements the operations to craete concrete product objects.
  * **AbstractProduct**
    * declares an interface for a type of product object.
  * **ConcreteProduct**
    * defines a product object to be created by the corresponding concrete factory.
    * implements the AbstractProduct interface.
  * **Client**
    * uses only interfaces declared by AbstractFactory and AbstractProduct classes.
* **Collaborations**
  * One instance of a ConcreteFactory class is created at run-time.
  * Clients uses different concrete factories to create different product objects.
  * AbstractFactory defers creation of product objects to its ConcreteFactory subclass.
* **Consequences**
  * It isolates concrete classes.
  * It makes exchanging product families easy.
  * It promotes consistency among products.
  * Supporting new kinds of products is difficult.
    * Supporting new kinds of products -> extending the factory interface -> changing the AbstractFactory class and all of its subclasses.
* **Implementation**
  * Factories as singletons.
  * Creating the products.
    * Define a factory method for each product.
    * A concrete factory will specify its products by overriding the factory method for each.
    * If many product families are possible, the concrete factory can be implemented using the Prototype pattern.
      * The concrete factory is initialized with a prototypical instance of each product in the family, and it creates a new product by cloning its prototype.
  * Defining extensible factories.
    * Add a parameter to specify the kind of object to be created -> only need a single "Make" operation with a parameter indicating the kind -> more flexible, less safe.
* **Related Patterns**
  * Factory Method -> implement AbstractFactory classes.
  * A concrete factory -> always a Singleton.

## Object Creational: Builder

* **Intent**
  * Separate the construction of a complex object from its representation so that the same construction process can create different representations.
* **Motivation**
  * The problem with open-ended number of components to build one object -> the need to add a new component without modifying the object.
* **Applicability**
  * Use when
    * “the algorithm for creating a complex object should be independent of the parts that make up the object and how they’re assembled.”
    * “the construction process must allow different representations for the object that’s constructed.”

* **Structure**

![pg98fig01](images/3 Creational Patterns/pg98fig01.jpg)

* **Participants**
  * **Builder**
    * specifies an abstract interface for creating parts of a Product object.
  * **ConcreteBuilder**
    * constructs and assembles parts of the product by implementing the Builder interface.
    * defines and keeps track of the representation it creates.
    * provides an interface for retrieving the product.
  * **Director**
    * constructs an object using the Builder interface.
  * **Product**
    * represents the complex object under construction.
    * includes classes that define the constituent parts, including interfaces for assembling the parts into the final result.
* **Collaborations**
  * The client creates the Director object and configures it with the desired Builder object.
  * Director notifies the builder whenever a part of the product should be built.
  * Builder handles requests from the director and adds parts to the product.
  * The client retrieves the product from the builder.

![pg99fig01](images/3 Creational Patterns/pg99fig01-6379149.jpg)

* **Consequences**
  * It lets you vary a product's internal representation.
  * It isolates code for construction and representation.
  * It gives you finer control over the construction process.
* **Implementation**
  * Assembly and construction interface.
    * Builders construct their products in step-by-step fashion.
    * A model of appending requests is usually sufficient.
  * Why no abstract class for products?
    * The produced products differ greatly -> no need to have a common parent class.
  * Empty methods as default in Builder.
    * Empty methods -> let clients override the ones needed.
* **Related Patterns**
  * A Composite is what the builder often builds.