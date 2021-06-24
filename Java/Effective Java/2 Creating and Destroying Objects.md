# Chapter 2. Creating and Destroying Objects

## Item 1: Consider static factory methods instead of constructors

* “A class can provide a public static factory method, which is simply a static method that returns an instance of the class.”
* “A class can provide its clients with static factory methods instead of, or in addition to, public constructors.”
* **“One advantage of static factory methods is that, unlike constructors, they have names.”**
  * “A static factory with a well-chosen name is easier to use and the resulting client code easier to read.”
* **“A second advantage of static factory methods is that, unlike constructors, they are not required to create a new object each time they’re invoked.”**
  * “This allows immutable classes (Item 17) to use preconstructed instances, or to cache instances as they’re constructed, and dispense them repeatedly to avoid creating unnecessary duplicate objects.”
* **“A third advantage of static factory methods is that, unlike constructors, they can return an object of any subtype of their return type.”**
  * “One application of this flexibility is that an API can return objects without making their classes public. Hiding implementation classes in this fashion leads to a very compact API.”
  * “Prior to Java 8, interfaces couldn’t have static methods. By convention, static factory methods for an interface named Type were put in a noninstantiable companion class (Item 4) named Types.”

* **“A fourth advantage of static factories is that the class of the returned object can vary from call to call as a function of the input parameters.”**
  * “The existence of implementation classes can be invisible to clients.”
  * “Clients neither know nor care about the class of the object they get back from the factory.”
* **“A fifth advantage of static factories is that the class of the returned object need not exist when the class containing the method is written.”**
* **“The main limitation of providing only static factory methods is that classes without public or protected constructors cannot be subclassed.”**
* **“A second shortcoming of static factory methods is that they are hard for programmers to find.”**
* **Common names for static factory methods:**
  * **from** - “A *type-conversion* method that takes a single parameter and returns a corresponding instance of this type.”
  * **of** - “An *aggregation* method that takes multiple parameters and returns an instance of this type that incorporates them.”
  * **valueOf** - “A more verbose alternative to `from` and `of`.”
  * **instance** or **getInstance** - “Returns an instance that is described by its parameters (if any) but cannot be said to have the same value.”
  * **create** or **newInstance** - “Like `instance` or `getInstance`, except that the method guarantees that each call returns a new instance.”
  * **get*Type*** - “Like `getInstance`, but used if the factory method is in a different class.”
  * **new*Type*** - “Like `newInstance`, but used if the factory method is in a different class.”
  * ***type*** - “A concise alternative to `getType` and `newType`.”
