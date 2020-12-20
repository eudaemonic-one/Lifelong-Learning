# Lecture 6 Requirements 1: Overview and Concepts

* Requirements say what the system will do (and not how it will do it)

## Why is This Hard?

* Communication problem
  * Goal: figure out what should be built
  * Express those ideas so that the correct thing is built
* Overall problems
  * Involved subproblems?
  * Required functionality?
  * Nice to have functionality?
  * Expected qualities?
  * How fast to devliver at what quality for what price?

## Requirement Engineering

* Knowledge acquisition - how to capture relevant detail about a system?
* Knowledge representation - once captured, how do we express it most effectively?

### Functional Requirements

* What the machine should do
  * Input
  * Output
  * Interface
  * Response to events
* Criteria
  * Completeness: All requirements are documented
  * Consistency: No conflcts between requirements
  * Precision: No ambiguity in requirements

### Quality/Non-functional Requirements

* Specify not the functionality of the system, but the quality with which it delivers that functionality
* Can be more critical than functional requirements
* Design criteria to help choose between alternative implementations
  * Confidentiality
  * Privacy
  * Integrity
  * Availability
* Requirements serve as contracts: they should be testable/falsifiable
* Informal goal: a general intention, such as ease of use
* Verifiable non-functional requirements: A statement using some measure that can be objectively tested

### Domain Knowledge

* Refinement is the act of translating requirements into specifications (bridging the gap)
* Requirements: desired behavior (effect on the environment) to be realized by the proposed system
* Assumptions or domain knowledge: existing behavior that is unchanegd by the proposed system
* Some gaps must remain
  * Unshared actions cannot be accurately expressed in the machine
  * Future requirements are also not directly implementable

### Avoiding Implementation Bias

* Requirements describe what is observable at the environment-machine interfac
* Indicative mood describes the environment (as-is)
* Optative mood to describe the environment with the machine (to-be)

## Activities of Requirements Engineering

* Identify stakeholders
* Understand the domain
  * Analyze artifacts, interact with stakeholders
* Discover the real needs
  * Interview stakeholders
* Explore alternatives to address needs
* Who is the system for?
* Stakeholders:
  * End users
  * System administrators
  * Engineers maintaining the system
  * Business managers