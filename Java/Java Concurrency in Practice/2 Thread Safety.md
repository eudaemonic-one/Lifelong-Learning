# Chapter 2. Thread Safety

* Writing thread-safe code => managing access to *shared, mutable state*.
  * An object's state: any data that can affect its externally visible behavior.
  * shared: could be accessed by multiple threads.
  * mutable: its value could change during its lifetime.
* Primary synchronization in Java: `synchronized`, `volatile` variables, explicit locks, atomic variables.
* Fix broken multi-threaded programs:
  * => Don't *share* the state variable across threads.
  * => Make the state variable *immutable*.
  * => Use *synchronization* whenever accessing the state variable.
* **It is far easier to design a class to be thread-safe than to retrofit it for thread safety later.**
* **It is always a good practice first to make your code right, and *then* make it fast.**
  * Pursue optimization only if your performance measurements and requirements tell you that you must, your optimizations actually made a difference under realistic conditions.