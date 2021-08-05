# Lecture 5 What's in a name?

## Names Matter

* Each API is a little language
* Primary goals
  * Easy to read
  * Hard to misread
  * Easy to write
* Names should:
  * be largely self-explanatory
  * leverage existing knowledge
  * interact harmoniously with language and each other

### How to Choose Names

* Choose **key nouns** carefully
  * They often become types
* Names can be **literal** or **metaphorical**
  * Literal names have literal associations
    * e.g., matrix suggests inverse, determinant, eigenvalue, etc
  * Metaphorical names enable reasoning by analogy
    * e.g., mail suggests send, inbox, outbox, etc

### Names Drive Development

* Good names drive good development
* Bad names inhibit good development
* Bad names result in bad APIs unless you take action
* **The API talks back to you. Listen!**
* Names may remind you of another API
  * Consider **copying** its vocabulary and structure
  * People who know other API will learn yours easily

### Vocabulary Consistency

* Use words consistently throughout your API
  * Never use the same work for multiple meanings
  * Never use multiple words for the same meaning
* The tighter the scope, the more important is consistency
  * **Within APIs, consistency is critical**
  * In related APIs on a platform, it's highly desirable
* It forced to choose between local and platform consistency, always choose local

### Avoid Abbreviations Except Where Customary

* In the bad old days, storage was scarce and people abbreviated everything
  * Unix is exhibit A; Linux APIs still suffer
* **Ideally, use complete words**
* But sometimes, names just get too long
* Of course you should use `gcd`, `url`, `cos`, etc

### Strive for Symmetry

* **If API has 2 verbs & 2 nouns, support all 4 combinations**
* In other words, good APIs are generally **orthogonal**

### Don't Mislead Your User

* Names have implications & learn these implications
* Don't violate the principle of least astonishment
* **Ignore this advice at your own peril**
  * Can cause unending stream of subtle bugs

### Don't Lie to Your User

* Name method for what it does, not what you wish it did
* If you can't bring yourself to do this, fix the method

### Good Naming Takes Time, But It's Worth It

* Don'be afraid to spend hours on it; I do
* **Don't just list names and chooses**
  * Think about goals and anti-goals for names
  * Think of names consistent with these goals
  * Write out realistic client code and compare
* Discuss names with colleagues; it really helps