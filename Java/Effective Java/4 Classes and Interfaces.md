# Chapter 4. Classes and Interfaces

## Item 15: Minimize the accessibility of classes and members

* “A well-designed component hides all its implementation details, cleanly separating its API from its implementation. Components then communicate only through their APIs and are oblivious to each others’ inner workings. This concept, known as *information hiding* or *encapsulation*, is a fundamental tenet of software design [Parnas72].”
* “Information hiding is important for many reasons, most of which stem from the fact that it *decouples* the components that comprise a system, allowing them to be developed, tested, optimized, used, understood, and modified in isolation.”
* “Java has many facilities to aid in information hiding. The *access control* mechanism [JLS, 6.6] specifies the accessibility of classes, interfaces, and members. The *accessibility* of an entity is determined by the location of its declaration and by which, if any, of the access modifiers (`private`, `protected`, and `public`) is present on the declaration.”
* “The rule of thumb is simple: **make each class or member as inaccessible as possible**.”
  * “By making it package-private, you make it part of the implementation rather than the exported API, and you can modify it, replace it, or eliminate it in a subsequent release without fear of harming existing clients.”
  * “If you make it public, you are obligated to support it forever to maintain compatibility.”
  * “If a package-private top-level class or interface is used by only one class, consider making the top-level class a private static nested class of the sole class that uses it (Item 24).”

* “For members (fields, methods, nested classes, and nested interfaces), there are four possible access levels, listed here in order of increasing accessibility:”
  * “**private**—The member is accessible only from the top-level class where it is declared.”
  * “**package-private**—The member is accessible from any class in the package where it is declared. Technically known as default access, this is the access level you get if no access modifier is specified (except for interface members, which are public by default).”
  * “**protected**—The member is accessible from subclasses of the class where it is declared (subject to a few restrictions [JLS, 6.6.2]) and from any class in the package where it is declared.”
  * “**public**—The member is accessible from anywhere.”
* “That said, both private and package-private members are part of a class’s implementation and do not normally impact its exported API. These fields can, however, “leak” into the exported API if the class implements `Serializable` (Items 86 and 87).”
* “For members of public classes, a huge increase in accessibility occurs when the access level goes from package-private to protected. A protected member is part of the class’s exported API and must be supported forever. Also, a protected member of an exported class represents a public commitment to an implementation detail (Item 19). The need for protected members should be relatively rare.”
* “**If a method overrides a superclass method, it cannot have a more restrictive access level in the subclass than in the superclass** [JLS, 8.4.8.3]. This is necessary to ensure that an instance of the subclass is usable anywhere that an instance of the superclass is usable (the Liskov substitution principle, see Item 15).”
* “To facilitate testing your code, you may be tempted to make a class, interface, or member more accessible than otherwise necessary. This is fine up to a point. It is acceptable to make a private member of a public class package-private in order to test it, but it is not acceptable to raise the accessibility any higher.”
* “**Instance fields of public classes should rarely be public** (Item 16).”
  * **“Classes with public mutable fields are not generally thread-safe.”**
* “You can expose constants via public static final fields, assuming the constants form an integral part of the abstraction provided by the class. By convention, such fields have names consisting of capital letters, with words separated by underscores (Item 68). It is critical that these fields contain either primitive values or references to immutable objects (Item 17).”
* **“It is wrong for a class to have a public static final array field, or an accessor that returns such a field.”**
  * “You can make the public array private and add a public immutable list.”
  * “Alternatively, you can make the array private and add a public method that returns a copy of a private array.”
* “As of Java 9, there are two additional, implicit access levels introduced as part of the *module system*.”
  * “A module is a grouping of packages, like a package is a grouping of classes.”
  * “A module may explicitly export some of its packages via *export declarations* in its *module declaration* (which is by convention contained in a source file named module-info.java).”
  * “Using the module system allows you to share classes among packages within a module without making them visible to the entire world.”
