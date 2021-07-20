# Chapter 4. Composing Objects

* Structuring classes => make them thread-safe, maintain them without accidentally undermining their safety guarantees.

## 4.1 Designing a Thread-safe Class

* The design process for a thread-safe class:
  * Identify the variables that form the object's state;
    * An object's state := its fields.
  * Identify the invariants that constrain the state variables;
  * Establish a policy for managing concurrent access to the object's state.
    * *synchronization policy* := what combination of immutability, thread confinement, locking used to maintain thread safety, which variables are guarded by which locks.
    * => defines how an object coordinates access to its state without violating its invariants or postconditions.
    * remember to document the synchronize policy.

![c0056-01](images/4 Composing Objects/c0056-01.jpg)

* **Gathering Synchronization Requirements**
  * state space := the range of possible states.
  * smaller state space => easier to reason.
  * using final fields or immutable objects => simpler to analyze the state space.
  * invariants and postconditions => additional synchronization or encapsulation requirements applied on state or state transitions.
* **State-dependent Operations**
  * use existing library classes => wait for state-based preconditions.
    * e.g., blocking queues, semaphores, synchronizers.
* **State Ownership**
  * publish a mutable object => no longer have exclusive control => shared ownership.
  * collection classes owns the state of the collection infrastructure, but client codes owns the objects stored in the collection.
