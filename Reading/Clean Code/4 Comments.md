# Chapter 4: Comments

* "Don't command bad code - rewrite it."
* The proper use of comments is to compensate for our failure to express ourself in code.
* Comments often lie. The older a comment is, and the farther away it is from the code it describes, the more likely it is to be just plain wrong.

## Comments Do Not Make Up for Bad Code

* Clear and expressive code with few comments is far superior to cluttered and cmoplex code with lots of comments.

## Explain Yourself in Code

* In many cases it's simply a matter of creating a function that says the same thing as the comment you want to write.
  * `// Check to see if the employee is eligible for full benefits` vs. `if (employee.isEligibleForFullBenefits())`
