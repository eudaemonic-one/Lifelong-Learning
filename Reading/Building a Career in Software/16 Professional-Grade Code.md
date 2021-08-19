# Chapter 16: Professional-Grade Code

## Master Language Conventions

* Favor adherence to language conventions.

## Naming

* Good naming has three objectives:
  * **Explicitness**: The name captures the role of the entity.
  * **Brevity**: The name is short and easy to read.
  * **Consistency**: The name follows the conventions of other similar entities in the codebase.

## Match the Codebase

* You should always match the coding conventions of that codebase even if you disagree with them; going off on your own will likely leave a confusing mess for future readers even if it does get past code reviewers, which it shouldn't.

## Commenting

* Comments exist to help future readers (including yourself) gain a holistic sense of your code as quickly as possible.
* Four guidelines:
  * **Motivation**: The motivation for your decision lives only in your own mind unless you write it down, and you can be certain that future readers will be curious.
  * **Design**: Give the reader a hand to the overall design because the reader's eyes can only see one piece of code at a time.
  * **Assumptions or invariants**: preconditions, even if asserted at runtime.
  * **Anything else nonobvious a reader might wonder about.**

## Commit Messages Are Underrated

* A commit message is our one chance to capture our goals, design, and trade-offs at the time of our work for eternal posterity, even if the code later changes.
* Template:

```text
<One-sentence summary of purpose>
<Blank line>
<Detailed explanation of purpose> <Design> <Tradeoffs, quirks, caveats>
```

## Testing

* Code isn't highly quality without tests, whatever the form your team may favor (unit, integration, end to end).
* If in doubt, start with true unit tests (i.e., tests that examine a single subcomponent in isolation, like a function, package, or class).
* You should measure code coverage, that is, what lines and branches of your program are exercised by your tests.
* There's no magic number for code coverage, but I prefer 100% when at all possible.

## Cleverness

* Favor explicitness and simplicity over cleverness whenever possible.

## Duplication vs. Reuse

* Generalizing a shared codepath can sometimes be higher risk than creating a new, similar function.
  * Good tests reduce that risk.
  * You first impulse should always be toward reuse, but if it's difficult, if your confidence in tests is poor, and if the codepath is critical, you might consider duplicating.

## Don't Check Nonsensical Conditions

* How to treat `NULL`, `nil`, or `None` depends on context.

## Open Source Conservatism

* Taking on an open source dependency is not much different from installing and frequently running a program written by some rando on the Internet.

## Tests Should Be Fast (Never Sleep)

* A slow test is an infuriating obstacle to development velocity.
  * You can inject a clock interface everywhere, so you mock that clock and accelerate time for tests => not waiting for a condition like a timeout to occur.

## Unit vs. Integration Testing

* Integeration tests can be slow; unit tests can be incomplete.
  * Unit tests let us itearate quickly on changes in a component, shaping our work with frequent testing.
  * Integration tests run longer bug give us a higher degree of confidence that a component works with adjacent systems before we send it to customers.
* A healthy test suite should have both.

## Inject Dependencies

* Dependency injection => decouple components, making it easy to change the implementation of a dependency.
* If we inject an interface, we can swap in a mocked implementation, which is a lifesaver for testing.

## Performance Last or a List of Priorities for Your Code

* Default priorities for a codebase should be
  * **Correctness**: The code does what it's supposed to do and is easy to verfiy.
  * **Maintainability**: The code is simple, easy for ourselves and others to change, and is likely to work under reasonable changes in operating conditions.
  * **Cost of development**: The code is fast to write, that is, minimizes engineering costs.
  * **Performance**: The code run fast.
* We should only pay readability for performance when performance really matters and we have data to suggest that the specific optimization in question will help.

## Code Review

* All production code changes are read by one or more other engineers to verify their correctness and quality.
* It usually proceeds in cycles, with the reviewer asking questions and making suggestions, then the author answering those questions, accepting some suggestions, and pushing back on others.
* **Receiving Code Reviews**
  * Slow down today is better than writing a postmortem.
  * Acknowledge and address every comment and never ignore any comment.
  * Ask follow-up questiosn on feedback:
    * Does the reviewer feel strongly about that suggestion?
    * Do they think it's okay to defer some changes for later?
    * Why do they think that using this pattern is better than the other?
* **Reviewing Code**
  * Ask myself about any diff:
    * **Is it correct?**
      * First understand exactly what a change is intended to do.
      * Second understand the system.
      * Finally read in detail, checking edge cases, error handling, threading, language usage.
    * **Is it clear?**
      * e.g., naming, commenting, simplicity, function length, commit messages, file layout.
    * **Does it match the code around it?**
      * e.g., code style, quality trade-offs.
    * **Does it reinvent anything?**
      * Duplication should be explained.
    * **Is it well-tested?**
      * Err toward asking for full coverage.
    * **Is it idiomatic?**
      * Mastering your team's languages.
    * **Is the diff reasonably sized?**
      * Large diffs are hard to review and hard to get right.
      * You should discourage them and look for smaller atoms that could be delivered separately.
* **Beyond the Code**
  * First is the tone of feedback.
    * Always be supportive; we err toward trusting our colleagues' diligence.
  * Second is latency.
    * Long waits are frustrating.
    * Treating code reviews as top priority tasks and preempting any nonurgent work.
  * Third, avoid the ping-pong review pair.
    * Try to get some diversity of perspective in the code review stream.
  * Finally, when to hold the line on quality and when to bend.
    * When it comes to incorrectness, draw a hard line.
