# Lecture 24 Case Study \- JDK 1.02 Libraries (The First Java Release)

* Only 7 Packages covered - 4 Core, 3 AWT
  * `java.lang` - 21classes, 2 interfaces
  * `java.util` - 10classes, 2 interfaces
  * `java.io` - 23 classes, 3 interfaces
  * `java.net` - 11 classes, 3 interfaces

## `java.lang` Critique

* `Object` - Methods `equals` and `hashCode` enabled global equivalence relation, hash-based collections
* `String` - Immutability very daring for the time, a huge win
  * Parameterless and copy constructors were misguided
  * Hash values specified down to the bit (in JLS!) were a very bad idea
* `StringBuffer` - Mutable Companion Class / Builder
  * Overall structure is great
  * But should never have been synchronized
* `Integer`, `Long`, `Float`, `Double`, `Character`, `Boolean`
  * Immutable is beautiful
  * `equals` semantics differ from primitive `==`
    * Necessary but confusing
  * Hash values specified down to the bit
  * `Boolean(booleanValue)` - should have been a static factory
  * `MIN_VALUE` - Inconsistent meaning in floating point and integral types
  * Hideous hierarchy inversion - `Boolean.getBoolean(String)` returns `true` if named system property exists and is equal to `"true"` (WAT?)
  * What happended to `Byte` and `Short`? (They showed up in JDK1.1)
* `Number`
  * Failed attempt to abstract over wrapped numerical primitives
  * Impossible task; should have been omitted
* `Math`
  * Reusing well-known C names and semantics a wise choice
  * Providing only floating point `random` was a bad idea
  * `abs(Integer.MIN_VALUE)` mayeb should have thrown exception
* `Throwable` - Exception hierarchy is broken
  * Unchecked (and checked) throwables don't have a single root
  * Checked exceptions heavily overused in all APIs
    * `InterruptedException`, `IOException` particularly painful
  * Little consistency in organizing exceptions
  * Many botches, e.g., `NoSuchMethodError`, `NoSuchMethodException`
  * "I don't believe much thought was given to the exception hierarchy." - Frank Yellin
* `Thread`
  * Extending `Runnable` violates LSP
  * `stop`, `suspend`, `resume` - should have been omitted
    * Deprecated in 1.2 (1998); removed in Java 11 (2018)
  * State transitions ill-defined; it took years to get this right (5.0, 2004)
  * Thread priorities are complex and not portable but necessary?
* `ThreadGroup` - Pretty much useless and filled with errors
  * Originally thought to have security uses - it didn't
  * "Putting your cooperating threads into a hierarchy seemed like a good idea at the time." - Frank Yellin
* `Class`
  * `newInstance` limited; broke exception safety, but way better than nothing
* `Process` - Extremely useful but complete nonportable
  * With the rubicon crossed, should have included other such features
  * Serious usability issues concerning output, error streams
* `System`, `Runtime` - should have been combined
  * `in`, `out`, `err` are public fields we soon came to regret
    * "This came straight from C/Unix." - Frank Yellin
    * Final fields, but methods addedd to mutate them in 1.1
    * JLS memory model (1.5) needed special language to deal with this mess
  * `arraycopy` violates naming conventions, type safety
* `Runnable`
  * `void` return a bit sad, but correct for its time
* `Cloneable` - utterly broken, should have been omitted
  * It lacks a `clone` method
  * Worse, creating objects without invoking constructors is evil
  * Sent us down the wrong path for serialization
* `ClassLoader` - Very powerful but ill-understood
  * Left huge bug/complexity tail in its wake (e.g., "shadow type system")
* `SecurityManager` - designed to enable foreign code
  * Clever idea for its time; it seemed that security was designed in
  * Never made good on promise-shared memory security is intractable
    * To this day! (Search web for: site:oracle.com applet security vulnerability)
  * Security is a cross-cutting concern; made Java code uglier
* `Compiler` - initially forward-looking, then obsolete

## `java.util` Critique

* `Vector`, `Hashtable` - Fantastic! Most useful data structures
  * C++ STL takes a book to describe and still no a hash table (2011)
  * Synchronizing everything was probably a missable, but defensive
  * Minor botches, e.g., `Hashtable.contains(Object value)`, Hashtable -> HashTable
  * Implementation dependent tuning parameters hinder evolution
* `Enumeration` - did the job
  * External iteration was correct for its day
  * Names too long for such commonly used functionality
    * Luckily they left the good real estate (`Iterator`) for me
* `Dictionary`
  * Should have been interface not abstract class
  * Blessing in disguise: In 1.2, it was clear that `Map` was the real deal
  * Should have been excluded
* `Stack` - Hideously broken toy
  * Extreme is-a/has-a confusion (extends `Vector`)
  * A great piece of real estate wasted forever
* `Date` - Should have been immutable!
  * "Java's worst botch" - Doug Lea
* `Random` - Decent, but should have been an interface
  * "Interfaces were second-class citizens. The concept of programming to interfaces never occurred to anyone." - Frank Yellin
  * Like hash codes, it was a mistake to specify results bit-for-bit
* `BitSet` - Not clear that it paid for itself, but not bad
  * Many important operations missing; added in 1.2
  * Size abstraction ill-defined; fixed in 1.2
* `Properties` - A huge mess
  * Should not be a subclass of `Hashtable` (is-a/has-a confusion)
  * Hierarchical defaults are complexity for its own sake
  * On-disk format is so poor that it's not amenable to a BNF
  * And it's with us forever because of system properties
  * API says keys, values must be strings; implementation permits arbitrary objects
* `StringTokenizer` - A broken toy
  * Unrelated `StreamTokenizer` class was a bad sign
  * Multiple ways of doing things have multiplied
* `Observable`/`Observer` - Another broken toy
  * Interface is overly complex (`hasChanged`, `clearChanged`)
  * To a first approximation, no one ever used it

## `java.io` Critique

* Looked decent, but didn't stand up to serious use
* Inconvenient to program against
  * Buffering requires wrapping in `Buffered{Input/Output}Stream`
  * All exceptions checked (they inherit from `IOException`)
  * `InputStream.skip` treats its argument as a hint
    * Using it safely requires a large amount of boilerplate
* Lacked orthogonality, e.g., `RandomAccessFile`
* Performance was inherently poor
* Did not scale well
  * `java.nio` was intended to address this in 1.4
  * It didn't
* Use of bytes for printing was an understandable defect
  * That is still with us today

## `java.net` Critique

* A fine abstraction of BSD networking to this day
* Few TCP options supported; fixed by `SocketOptions` in 1.1
* Extensibility turned out to be largely unnecessary
* And the edifice wasn't terribly well designed
  * `URL.hashCode` and `equals` are blocking operations and violated their general contracts - `URI` replaced `URL` in 1.4 (2002)
  * Capitalzation of acronyms throughout is yucky

## `java.awt` Critique

* AWT Threading Critique
  * No one understand the AWT threading model in those days
    * Maybe there wasn't one
  * All published AWT/Swing examples through 2000 were broken
    * In those days it didn't matter much
  * `java.awt.EventQueue.invokeLater` not added till 1.2
* Best viewed as a proof of concept
* Impressive for how quickly it was developed
* Showed the power of Java in the browser
  * Essential for Java's success (at the time)
* Never did much more than demo applets
  * But that's understandable
* No good reason for AWT's zombie to exist in Swing
  * Should have veneered over it

## Overall Libraries Critique

* Small and manageable
  * A programmer could quickly and easily understand the entire platform
* Far simpler than C++; far safer and more powerful than C
* It was clear that many APIs hadn't been seriously used
  * Some toys should have been left out
  * Some APIs should have been simplified
  * There were many small API flaws
* Oversynchronization was common
  * Design philosophy: synchronize every method, don't worry about threads
* On balance, libraries were **good enough**
  * Astonishing considering time constraints