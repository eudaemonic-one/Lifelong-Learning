# Chapter 3. Sharing Objects

* Synchronization := *atomicity* ("critical sections") + *memory visibility*.
  * Prevent one thread from modifying the object's state when another is using it.
  * Ensure other threads can *see* the changes that were made by one thread.
* Sharing and publishing objects safely => objects are safely accessed by multiple threads.
