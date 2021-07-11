# Chapter 10. Exceptions

## Item 69: Use exceptions only for exceptional conditions

* **“Exceptions are, as their name implies, to be used only for exceptional conditions; they should never be used for ordinary control flow.”**
  * “More generally, use standard, easily recognizable idioms in preference to overly clever techniques that purport to offer better performance.”
  * “Even if the performance advantage is real, it may not remain in the face of steadily improving platform implementations.”
  * “The subtle bugs and maintenance headaches that come from overly clever techniques, however, are sure to remain.”
* “A well-designed API must not force its clients to use exceptions for ordinary control flow.”
  * “A class with a “state-dependent” method that can be invoked only under certain unpredictable conditions should generally have a separate “state-testing” method indicating whether it is appropriate to invoke the state-dependent method.”
    * “For example, the `Iterator` interface has the state-dependent method `next` and the corresponding state-testing method `hasNext`.”
  * “An alternative to providing a separate state-testing method is to have the state-dependent method return an empty optional (Item 55) or a distinguished value such as `null` if it cannot perform the desired computation.”
  * “If an object is to be accessed concurrently without external synchronization or is subject to externally induced state transitions, you must use an optional or distinguished return value, as the object’s state could change in the interval between the invocation of a state-testing method and its state-dependent method. Performance concerns may dictate that an optional or distinguished return value be used if a separate state-testing method would duplicate the work of the state-dependent method.”
  * “All other things being equal, a state-testing method is mildly preferable to a distinguished return value. It offers slightly better readability, and incorrect use may be easier to detect: if you forget to call a state-testing method, the state-dependent method will throw an exception, making the bug obvious; if you forget to check for a distinguished return value, the bug may be subtle. ”
* **“In summary, exceptions are designed for exceptional conditions. Don’t use them for ordinary control flow, and don’t write APIs that force others to do so.”**