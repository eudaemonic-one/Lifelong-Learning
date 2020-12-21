# Lecture 11 Introduction to Software Architecture and Documentation

## Software Architecture

* Comprise software elements, the externally visible properties of those elements, and the relationships among them
* Aids in communicationc with stakeholders
* Defines constraints on implementation
* Dictates organizational structure
* Inhibits or enables quality attributes
* Supports predicting cost, quality, and schedule
* Aids in software evolution and prototyping
* Quality matters
  * Performance, Availability, Modifiability, Portability, Scalability, Security, Testability, Usability, Cost to build, Cots to operate

## Twitter Case Study

* Architectural decisions affect entire systems, not only individual modules
* Abstract, different abstractions for different scenarios
* Reason about quality attributes early
* Make architectural decisions explicit

## Architecture vs. Object-level Design

* Requirements
* Architecture
* OO-Design
* Code

| Design Questions                       | Architectural Questions                            |
| -------------------------------------- | -------------------------------------------------- |
| How do I add a menu item in Eclipse?   | How do I extend Eclipse with a plugin?             |
| How lock protects this data?           | What threads exists and how do they coordinate?    |
| How does Google rank pages?            | How does Google scale to billions of hits per day? |
| What is the interface between objects? | What is the interface between subsystems?          |

## Architecture Documentation & Views

* Blueprint for the system
  * Artifact for early analysis
* Documentation speaks for the architect
* Support traceability
* Static View
  * Modules (subsystems, structures) and their relations (dependencies)
* Dynamic View
  * Components (processes, runnable entities) and connectors (messages, data flow)
* Physical View (Deployment)
  * Hardware structures and their connections
* Selecting a Notation
  * Suitable for purpose
  * UML possible (semi-formal), but possibly constraining
  * Always include a legend
  * Supplement graphics with explanation
  * Do not try to do too much in one diagram

