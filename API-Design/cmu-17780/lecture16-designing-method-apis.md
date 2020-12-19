# Lecture 16 Designing Method APIs, Cont.

* Programmatic access to all data available as string
* Overload with care
* Use appropriate parameter and return types
* Avoid long parameter lists

## Return Values

* Avoid return values that demand exceptional processing
  * Client should not have to write extra code
  * e.g., return 0-length array or empty collection, not `null`
  * Think twice before returning `null`
    * A better solution
      * e.g., `Optional<T>` (aka maybe, option)
    * But sometimes null is the correct answer
      * For consistency or performance (e.g., `Map.get`)

## Input Params

* Don't make it **impossible** for user to benefit from the work you've done
  * Don't leave param list to be too long or obscure
* Use consistent parameter ordering across methods
  * Especially important if parameter types identical
* Avoid long parameter lists
  * Three of fewer parameters is ideal
    * More and users will have to refer to docs
  * Long lists of identical typed params very harmful
  * Techniques for shortening parameter lists
    * Break up method
    * Create a helper class to hold several parameter
    * Use builder pattern
* Use varargs only where truly beneficial
  * You know the benefits (`printf`, `String.format`)
  * Performance is also an issue (in Java)

## Behavior

* Don't make user play 20 questions with API
  * If there are only a few valid inputs, provide an easy way for users to get their hands on them
  * More generally, don't make callers invoke a method repeatedly until they get desired result
* Don't make caller do the work twice - provide fused read-modify-write where feasible
  * Required for concurrency, but good regardless
    * Improves usability, readability, and performance
  * Bad (`Map` prior to Java 8)
    * `Integer old = freq.get(word); int new = (old == null) ? 1 : old + 1; freq.put(word, new);`
  * Good (`Map` in Java 8)
    * `freq.merge(word, 1, Integer::sum)`
* Don't overspecify the behavior of methods
  * Don't specify internal details
  * All tuning parameters are suspect
    * Let client specify intended use, not internal detail
  * Do not let internal details leak into spec
  * Do not specify hash functions
* Specify methods entirely in terms of their inputs and outputs
  * Avoid accessing static state as if it were radioactive
  * Also external state, unless fundamental to abstraction

## Structure

* Create as few interface dependencies as possible, but no fewer
  * If a method takes a parameter or returns a value whose type if defined in another API, you've established an interface dependency
  * Remember, you get the transitive closure
  * It's a bad smell when a low-level API depends on a higher level API