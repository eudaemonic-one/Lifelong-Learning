# Chapter 3. Methods Common to All Objects

* “All of `Object`s nonfinal methods (`equals`, `hashCode`, `toString`, `clone`, and `finalize`) have explicit general contracts because they are designed to be overridden.”

## Item 10: Obey the general contract when overriding equals

* “The easiest way to avoid problems is not to override the equals method, in which case each instance of the class is equal only to itself.”

  * “Each instance of the class is inherently unique.”
    * e.g., `Thread`
  * “There is no need for the class to provide a “logical equality” test.”
    * e.g., `java.util.regex.Pattern`
  * “A superclass has already overridden equals, and the superclass behavior is appropriate for this class.”
    * e.g., `Set` inherits `equals` from `AbstractSet`
  * “The class is private or package-private, and you are certain that its equals method will never be invoked.”
* “So when is it appropriate to override `equals`? It is when a class has a notion of *logical equality* that differs from mere object identity and a superclass has not already overridden `equals`. ”
* “When you override the `equals` method, you must adhere to its **general contract**.”

  * “*Reflexive*: For any non-null reference value `x`, `x.equals(x)` must return `true`.”
  * “*Symmetric*: For any non-null reference values `x` and `y`, `x.equals(y)` must return `true` if and only if `y.equals(x)` returns `true`.”
  * “*Transitive*: For any non-null reference values `x`, `y`, `z`, if `x.equals(y)` returns `true` and `y.equals(z)` returns `true`, then `x.equals(z)` must return `true`.”
  * “*Consistent*: For any non-null reference values `x` and `y`, multiple invocations of `x.equals(y)` must consistently return true or consistently return `false`, provided no information used in equals comparisons is modified.”
  * “For any non-null reference value `x`, `x.equals(null)` must return `false`.”
* **“Once you’ve violated the equals contract, you simply don’t know how other objects will behave when confronted with your object.”**
* “Putting it all together, here’s a recipe for a high-quality equals method:”
  * **“Use the `==` operator to check if the argument is a reference to this object.”**
  * **“Use the `instanceof` operator to check if the argument has the correct type.”**
  * **“Cast the argument to the correct type.”**
  * **“For each “significant” field in the class, check if that field of the argument matches the corresponding field of this object.”**
* **“When you are finished writing your equals method, ask yourself three questions: Is it symmetric? Is it transitive? Is it consistent?”**
  * **“Always override hashCode when you override equals (Item 11).”**
  * **“Don’t try to be too clever.”** 
  * **“Don’t substitute another type for Object in the equals declaration.”**
* **“In summary, don’t override the `equals` method unless you have to: in many cases, the implementation inherited from `Object` does exactly what you want. If you do override `equals`, make sure to compare all of the class’s significant fields and to compare them in a manner that preserves all five provisions of the `equals` contract.”**