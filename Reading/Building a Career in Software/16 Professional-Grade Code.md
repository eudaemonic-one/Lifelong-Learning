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
