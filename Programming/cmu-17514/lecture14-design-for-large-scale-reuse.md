# Lecture 14 Design for Large-Scale Reuse: Libraries & Frameworks

## The Adapter Pattern

* Problem: You have a client that expects one API for a service provider, and a service provider with a different API.
* Solution: Write a class that implements the expected API converting calls to the service provider's actual API.
* Consequences:
  * Easy interoperability of unrelated clients and libraries
    * Client can use unforeseen future libraries
  * Adapter class is occupied to concrete service provider, can make it harder to override service provider behavior

## Terminology

* Library: a set of classes and methods that provide reusable functionality
* Framework: reusable skeleton code that can be customized into an application
  * Framework calls back into client code
  * The Hollywood principle
* API: Application Programming Interface
* Client: the code that uses the API
* Plugin: Client code that customizes the API
* Protocol: The expected sequnce of interactions between the API and the client
* Callback: A plugin method that the framework will call to access customized functionality

## Libraries and Frameworks in Practice

* Defines key abstractions and their interfaces
* Defines object interactions & invariants
* Defines flow of control
* Provides architectural guidance
* Provides defaults

## Whitebox vs. Blackbox Framework

* Whitebox framework
  * Extension via subclassing and overriding methods
  * template methods pattern
  * Subclass has main method but gives control to framework
* Blackbox framework
  * Extension via implementing a plugin interface
  * Strategy or Observer
  * Plugin-loading mechansim loads plugins and gives control to the framework

## Framework Design

* Once designed there is little opportunity for change
* Separating common parts from variable parts
* Possible problems:
  * Too few extension points
  * Too many extension points: Hard to learn, slow
  * Too generic: Little reuse value