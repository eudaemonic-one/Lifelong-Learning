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
