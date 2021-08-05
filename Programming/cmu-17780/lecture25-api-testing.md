# Lecture 25 API Testing

## General Principles

* General API principles apply: modular code with clear inputs, outputs, exceptions, and invariants in much easier to test (Avoid external state where possible)
* Don't degrade the quality of your API to make them more testable
  * Don't expose private API elements as public simply for testing
  * But feel free to make them package-private
* **Dependency Injection** is good practice (but a lousy term)
  * Have objects accept "dependencies" in their constructor rather than hard-coding
    * But provide defaults where appropriate
  * Functions can accept dependencies as arguments in the same way
  * Allows for "mocking" so you can test without the need for slow, expensive, or limited external resources (such as actual web services)
  * Using interface types instead of class types for arguments helps with this
  * Consider defining an interface for your objects so that users can easily test code that depends on them
* **Exhaustive testing** is great in the rare cases that it's practical
  * Test every possible combination of inputs to ensure correctness in all cases
  * Typically impractical (expecially for concurrent APIs)
* **Fuzzing** is a great alternative to exhaustive testing in more typical cases
  * An automated software testing techniqueu that involves providing invalida, unexpected, or random data as inputs to a computer program
  * Looking at the spec & testing all possible assertions
  * Exhaustive testing for small APIs
* How to ensure project is testable?
  * Implementation should not unduly impact interface
    * Examples of this relevance in web APIs: pagination, retry, failures (including timeouts)
      * Insulate the user from those
    * Private fields aren't part of the API - exclude them from API documentation
* How to test
  * Make sure always working on a specific version of dependencies
  * Use industry standard testing tools
* **Unit Tests** - Test each individual module independently
  * What to actually assert in testing - all testable assertions described in API
    * Normal cases - ideally with fuzzing or other stress testing
    * Edge cases (e.g., 0, 1, empty, etc)
    * Keep test cases simple - if you can't easily follow them then considering breaking out (private or package-private) helper functions and test those
    * No need to overdo it - this can bog down deployment pipelines
* **Integration** / **System Tests**
  * Good unit tests make these so much easier
  * Covering every edge/corner case is probably not practical
  * Make sure that 2+ reasonable inputs work properly (showing that the different smaller functions are integrated properly)
* **Mocking** - for web-based APIs, most tests should run without external dependencies
  * Unfortunate scenario: your API fails to build an pass unit tests because some server is down, and this stops you from deploying a critical patch
  * How to do mock: simulate known API queries (e.g., a GET with a certain path should return a certain result that you can process)
* **Static Analysis** - Use linters to catch simple mistakes with your code
  * Linters catch many tedious errors, such as violating language style guidelines
* **Continuous Integration** (CI) - manual testing does not scale
  * Leverage workflows to run sanity checks
  * Run the linter as well as a subset or entire test suite run on commits