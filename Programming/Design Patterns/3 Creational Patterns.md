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