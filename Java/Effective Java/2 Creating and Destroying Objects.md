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

## Item 2: Consider a builder when faced with many constructor parameters

* “Static factories and constructors share a limitation: they do not scale well to large numbers of optional parameters.”
* Telescoping Constructor Pattern
  * "You provide a constructor with only the required parameters, another with a single optional parameter, a third with two optional parameters, and so on, culminating in a constructor with all the optional parameters.”
  * “The telescoping constructor pattern works, but it is hard to write client code when there are many parameters, and harder still to read it.”
* JavaBeans Pattern
  * “You call a parameterless constructor to create the object and then call setter methods to set each required parameter and each optional parameter of interest.”
  * “Because construction is split across multiple calls, a JavaBean may be in an inconsistent state partway through its construction.”
  * “A related disadvantage is that the JavaBeans pattern precludes the possibility of making a class immutable (Item 17) and requires added effort on the part of the programmer to ensure thread safety.”
* **Builder Pattern**
  * “Instead of making the desired object directly, the client calls a constructor (or static factory) with all of the required parameters and gets a builder object. Then the client calls setter-like methods on the builder object to set each optional parameter of interest. Finally, the client calls a parameterless build method to generate the object, which is typically immutable."
  * “The builder is typically a static member class (Item 24) of the class it builds. ”
  * “The builder’s setter methods return the builder itself so that invocations can be chained, resulting in a *fluent* API.”
  * **“The Builder pattern is well suited to class hierarchies.”**
  * “A minor advantage of builders over constructors is that builders can have multiple varargs parameters because each parameter is specified in its own method.”
  * “While the cost of creating this builder is unlikely to be noticeable in practice, it could be a problem in performance-critical situations.”
  * “Builder pattern is more verbose than the telescoping constructor pattern, so it should be used only if there are enough parameters to make it worthwhile, say four or more.”
  * “It’s often better to start with a builder in the first place.”
* **“In summary, the Builder pattern is a good choice when designing classes whose constructors or static factories would have more than a handful of parameters.”**
