# Lecture 15 API Design 1: Process and Naming

## Library vs. Framework

* Library: Your code calls library API functions
* Framework: Framework calls your code

## API

* Application Programming Interface
* Component specification in terms of operations, inputs, & outputs
* Allows implementation to vary without compromising clients

## API Design Process

* true requirements gathering
* **if the problem can't be fixed, fail fast**
* start with short spec - 1 page is ideal
  * Agility trumps completeness
  * bounce spec off as many people as possible
  * flesh it out, you gain confidence
* write to your API early and often
  * start before implemented the API & before specified it properly
  * continue writing to API as you flesh it out
  * code lives on as examples, unit tests
* try API on at least 3 use cases before release (**RULE OF THREE**)
* maintain realistic expectations
  * don't please everyone
  * come up with a unified, coherent design that represents a compromise
  * expect to make mistakes
* issue tracking
  * throughtout the process, maintain a list of design issues
    * individual decisions such as what input format to accept
      * write down all the options
      * say which were ruled out and why
      * when you decide, say which was chosen and why
  * prevent wasting time on solve issues
  * provides rationale for the resulting API

### Naming

* Names are important
  * action verbs for mutation
  * Prepositions, linking verbs, nouns, or adjectives for pure functions
* Primary goals
  * easy to read
  * hard to misread
  * easy to write
  * be largely self-explanatory
  * leverage existing knowledge
  * Interact harmoniously with language and each other
* **Choose key nouns carefully**
  * find good abstraction
* Names can be literal or metaphorical
  * Matrix -> inverse, determinant, eigenvalue
  * Publication, Subscriber -> publish, subscribe, cancel, issue
* Consider **copying** vocabulary and structure from another API
* **The API talks back to you. Listen!**
* Use words consistently throughout your API
* Vocubulary consistency as it relates to scope
  * Within API, consistency is critical
  * If forced to choose between local & platform consistency, choose local
* Avoid abbreviations ecept where customary
* Grammar is a part of naming too
  * Nouns for classes
  * Nouns or adjectives for interfaces
* Names should be regular - strive for symmetry
  * if API has 2 verbs and 2nouns, support all 4 combinations
* Don't mislead your user
  * don't violate **the principle of least astonishment**
* Don't lie to your user
  * name method for what it does, not what you with it did
* Good naming takes time, but it's worth it