# Chapter 12. Testing Concurrent Programs

* Concurrent programs => nondeterminism => potential interactions, failure models.
* Tests of concurrent classes => safety + liveness.
  * Tests of safety => testing invariants such as assert invariants or execute test code atomically => can introduce timing or synchronization artifacts that can mask bugs that might otherwise manifest themselves.
  * Test of liveness => testing progress and nonprogress => hard to quantify.
* Performance measurement => Throughput + Responsiveness + Sacalability.