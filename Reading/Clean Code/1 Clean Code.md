# Chapter 1: Clean Code

## There Will Be Code

* Code => express the requirements ultimately.

## Bad Code

* Good code matters.

## The Total Cost of Owning a Mess

* Add more staff to the project in hopes of increasing productivity => make more and more messes.

## The Grand Redesign in the Sky

* Eventually management bend to the demands of the developers to redesign the code base and authorize the grand redesign in the sky.
* The redesign team must build a new system that does everything that the old systems does and also keep up with the continuously changes to the old system.

## Attitude

* Managers defend the schedule and requirements; it's your job to defend the code with passion.

## The Primal Conundrum

* The only way to make the deadline is to keep code as clean as possible at all times.

## The Art of Clean Code?

* This "code-sense" is the key.
* A programmer who writes clean code is an artist who can take a blank screen through a series of transformations until it is an elegantly coded system.

## What Is Clean Code?

* Bjarne Stroustrup - “I like my code to be elegant and efficient. The logic should be straightforward to make it hard for bugs to hide, the dependencies minimal to ease maintenance, error handling complete according to an articulated strategy, and performance close to optimal so as not to tempt people to make the code messy with unprincipled optimizations. Clean code does one thing well.”
* Grady Booch - “Clean code is simple and direct. Clean code reads like well-written prose. Clean code never obscures the designer’s intent but rather is full of crisp abstractions and straightforward lines of control.”
* "Big" Dave Thomas - “Clean code can be read, and enhanced by a developer other than its original author. It has unit and acceptance tests. It has meaningful names. It provides one way rather than many ways for doing one thing. It has minimal dependencies, which are explicitly defined, and provides a clear and minimal API. Code should be literate since depending on the language, not all necessary information can be expressed clearly in code alone.”
* Michael Feathers - “I could list all of the qualities that I notice in clean code, but there is one overarching quality that leads to all of them. Clean code always looks like it was written by someone who cares. There is nothing obvious that you can do to make it better. All of those things were thought about by the code’s author, and if you try to imagine improvements, you’re led back to where you are, sitting in appreciation of the code someone left for you—code left by someone who cares deeply about the craft.”
* Ron Jeffries -
  * “In recent years I begin, and nearly end, with Beck’s rules of simple code. In priority order, simple code: Runs all the tests; Contains no duplication; Expresses all the design ideas that are in the system; Minimizes the number of entities such as classes, methods, functions, and the like. Of these, I focus mostly on duplication. When the same thing is done over and over, it’s a sign that there is an idea in our mind that is not well represented in the code. I try to figure out what it is. Then I try to express that idea more clearly.”
  * “Expressiveness to me includes meaningful names, and I am likely to change the names of things several times before I settle in. With modern coding tools such as Eclipse, renaming is quite inexpensive, so it doesn’t trouble me to change. Expressiveness goes beyond names, however. I also look at whether an object or method is doing more than one thing. If it’s an object, it probably needs to be broken into two or more objects. If it’s a method, I will always use the Extract Method refactoring on it, resulting in one method that says more clearly what it does, and some submethods saying how it is done.”
  * “Duplication and expressiveness take me a very long way into what I consider clean code, and improving dirty code with just these two things in mind can make a huge difference. There is, however, one other thing that I’m aware of doing, which is a bit harder to explain.”
  * “Reduced duplication, high expressiveness, and early building of simple abstractions. That’s what makes clean code for me.”
* Ward Cunningham - “You know you are working on clean code when each routine you read turns out to be pretty much what you expected. You can call it beautiful code when the code also makes it look like the language was made for the problem.”

## Schools of Thought

* The rightness within a school does not invalidate the teachings of a different school.

## We Are Authors

* Making it easy to read actually makes it easier to write.

## The Boy Scout Rule

* The cleanup of code doesn't have to be something big.

## Prequel and Principles

* This book is a prequel to a book Agile Software Development: Principles, Patterns, and Practices (PPP).
* This book has references to the Single Responsibility Principles (SRP), the Open Closed Principle (OCP), and the Dependency Inversion Principle (DIP).
