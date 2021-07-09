# Chapter 9. General Programming

## Item 57: Minimize the scope of local variables

* “By minimizing the scope of local variables, you increase the readability and maintainability of your code and reduce the likelihood of error.”
* **“The most powerful technique for minimizing the scope of a local variable is to declare it where it is first used.”**
  * “Declaring a local variable prematurely can cause its scope not only to begin too early but also to end too late.”
* **“Nearly every local variable declaration should contain an initializer.”**
  * “If you don’t yet have enough information to initialize a variable sensibly, you should postpone the declaration until you do.”
    * “One exception to this rule concerns `try`-`catch` statements. If a variable is initialized to an expression whose evaluation can throw a checked exception, the variable must be initialized inside a `try` block (unless the enclosing method can propagate the exception).”
    * “If the value must be used outside of the `try` block, then it must be declared before the `try` block, where it cannot yet be “sensibly initialized.”
* “Loops present a special opportunity to minimize the scope of variables. The `for` loop, in both its traditional and for-each forms, allows you to declare *loop variables*, limiting their scope to the exact region where they’re needed.”
  * “Therefore, **prefer `for` loops to `while` loops**, assuming the contents of the loop variable aren’t needed after the loop terminates.”
* “If you need access to the iterator, perhaps to call its `remove` method, the preferred idiom uses a traditional `for` loop in place of the for-each loop:”


```java
// Idiom for iterating when you need the iterator
for (Iterator<Element> i = c.iterator(); i.hasNext(); ) {
    Element e = i.next();
    ... // Do something with e and i
}
```

* “Here is another loop idiom that minimizes the scope of local variables:”

```java
for (int i = 0, n = expensiveComputation(); i < n; i++) {
    ... // Do something with i;
}
```

* “The important thing to notice about this idiom is that it has *two* loop variables, `i` and `n`, both of which have exactly the right scope. The second variable, `n`, is used to store the limit of the first, thus avoiding the cost of a redundant computation in every iteration.”
* “A final technique to minimize the scope of local variables is to **keep methods small and focused**.”
  * “If you combine two activities in the same method, local variables relevant to one activity may be in the scope of the code performing the other activity. To prevent this from happening, simply separate the method into two: one for each activity.”


## Item 58: Prefer for-each loops to traditional `for` loops

```java
// Not the best way to iterate over a collection!
for (Iterator<Element> i = c.iterator(); i.hasNext(); ) {
    Element e = i.next();
    ... // Do something with e
}
```

```java
// Not the best way to iterate over an array!
for (int i = 0; i < a.length; i++) {
    ... // Do something with a[i]
}
```

* “These idioms are better than `while` loops (Item 57), but they aren’t perfect. The iterator and the index variables are both just clutter—all you need are the elements.”
* “Furthermore, they represent opportunities for error. The iterator occurs three times in each loop and the index variable four, which gives you many chances to use the wrong variable.”


```java
// The preferred idiom for iterating over collections and arrays
for (Element e : elements) {
    ... // Do something with e
}
```

* “When you see the colon (`:`), read it as “in.” Thus, the loop above reads as “for each element `e` in `elements`.”
* “There is no performance penalty for using for-each loops, even for arrays.”


```java
// Can you spot the bug?
enum Suit { CLUB, DIAMOND, HEART, SPADE }
enum Rank { ACE, DEUCE, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT,
            NINE, TEN, JACK, QUEEN, KING }
...
static Collection<Suit> suits = Arrays.asList(Suit.values());
static Collection<Rank> ranks = Arrays.asList(Rank.values());

List<Card> deck = new ArrayList<>();
for (Iterator<Suit> i = suits.iterator(); i.hasNext(); )
    for (Iterator<Rank> j = ranks.iterator(); j.hasNext(); )
        deck.add(new Card(i.next(), j.next()));
```

* “The problem is that the `next` method is called too many times on the iterator for the outer collection (`suits`). It should be called from the outer loop so that it is called once per suit, but instead it is called from the inner loop, so it is called once per card. After you run out of suits, the loop throws a `NoSuchElementException`.”
* “If instead you use a nested for-each loop, the problem simply disappears.”

```java
// Preferred idiom for nested iteration on collections and arrays
for (Suit suit : suits)
    for (Rank rank : ranks)
        deck.add(new Card(suit, rank));
```

* “Unfortunately, there are three common situations where you *can’t* use for-each:”
  * “**Destructive filtering**—If you need to traverse a collection removing selected elements, then you need to use an explicit iterator so that you can call its remove method. You can often avoid explicit traversal by using `Collection`’s `removeIf` method, added in Java 8.”
  * “**Transforming**—If you need to traverse a list or array and replace some or all of the values of its elements, then you need the list iterator or array index in order to replace the value of an element.”
  * “**Parallel iteration**—If you need to traverse multiple collections in parallel, then you need explicit control over the iterator or index variable so that all iterators or index variables can be advanced in lockstep.”
* “Not only does the for-each loop let you iterate over collections and arrays, it lets you iterate over any object that implements the `Iterable` interface, which consists of a single method. ”

```java
public interface Iterable<E> {
    // Returns an iterator over the elements in this iterable
    Iterator<E> iterator();
}
```

* **“In summary, the for-each loop provides compelling advantages over the traditional `for` loop in clarity, flexibility, and bug prevention, with no performance penalty. Use for-each loops in preference to `for` loops wherever you can.”**

## Item 59: Know and use the libraries

* “Suppose you want to generate random integers between zero and some upper bound. Faced with this common task, many programmers would write a little method that looks something like this:”

```java
// Common but deeply flawed!
static Random rnd = new Random();

static int random(int n) {
    return Math.abs(rnd.nextInt()) % n;
}
```

* “This method may look good, but it has three flaws.”
  * “The first is that if `n` is a small power of two, the sequence of random numbers will repeat itself after a fairly short period.”
  * “The second flaw is that if `n` is not a power of two, some numbers will, on average, be returned more frequently than others. If n is large, this effect can be quite pronounced.”
  * “The third flaw in the `random` method is that it can, on rare occasions, fail catastrophically, returning a number outside the specified range.”
* “To write a version of the `random` method that corrects these flaws, you’d have to know a fair amount about pseudorandom number generators, number theory, and two’s complement arithmetic. Luckily, you don’t have to do this—it’s been done for you. It’s called `Random.nextInt(int)`.”
* **“By using a standard library, you take advantage of the knowledge of the experts who wrote it and the experience of those who used it before you.”**
* “As of Java 7, you should no longer use `Random`.”
  * “For most uses, **the random number generator of choice is now `ThreadLocalRandom`**.”
  * “It produces higher quality random numbers, and it’s very fast.”

  * “For fork join pools and parallel streams, use `SplittableRandom`.”
* “A second advantage of using the libraries is that you don’t have to waste your time writing ad hoc solutions to problems that are only marginally related to your work.”
* “A third advantage of using standard libraries is that their performance tends to improve over time, with no effort on your part.”
* “A fourth advantage of using libraries is that they tend to gain functionality over time.”
  * “If a library is missing something, the developer community will make it known, and the missing functionality may get added in a subsequent release.”

* “A final advantage of using the standard libraries is that you place your code in the mainstream.”
  * “Such code is more easily readable, maintainable, and reusable by the multitude of developers.”
* **“Numerous features are added to the libraries in every major release, and it pays to keep abreast of these additions.”**
* “The libraries are too big to study all the documentation [Java9-api], but **every programmer should be familiar with the basics of `java.lang`, `java.util`, and `java.io`, and their subpackages**.”
* “Several libraries bear special mention. The collections framework and the streams library (Items 45–48) should be part of every programmer’s basic toolkit, as should parts of the concurrency utilities in `java.util.concurrent`. This package contains both high-level utilities to simplify the task of multithreaded programming and low-level primitives to allow experts to write their own higher-level concurrent abstractions.”
* **“To summarize, don’t reinvent the wheel.”**

## Item 60: Avoid `float` and `double` if exact answers are required

* “The `float` and `double` types are designed primarily for scientific and engineering calculations. They perform *binary floating-point arithmetic*, which was carefully designed to furnish accurate approximations quickly over a broad range of magnitudes. They do not, however, provide exact results and should not be used where exact results are required.”
* “**The `float` and `double` types are particularly ill-suited for monetary calculations** because it is impossible to represent 0.1 (or any other negative power of ten) as a float or double exactly.”
  * “The right way to solve this problem is to **use `BigDecimal`, `int`, or `long` for monetary calculations**.”
* “Here’s a straightforward transformation of the previous program to use the `BigDecimal` type in place of `double`. Note that `BigDecimal`’s `String` constructor is used rather than its `double` constructor. This is required in order to avoid introducing inaccurate values into the computation [Bloch05, Puzzle 2]:”


```java
public static void main(String[] args) {
    final BigDecimal TEN_CENTS = new BigDecimal(".10");
    int itemsBought = 0;
    BigDecimal funds = new BigDecimal("1.00");
    for (BigDecimal price = TEN_CENTS;
            funds.compareTo(price) >= 0;
            price = price.add(TEN_CENTS)) {
        funds = funds.subtract(price);
        itemsBought++;
    }
    System.out.println(itemsBought + " items bought.");
    System.out.println("Money left over: $" + funds);
}
```

* “There are, however, two disadvantages to using `BigDecimal`: it’s a lot less convenient than using a primitive arithmetic type, and it’s a lot slower. The latter disadvantage is irrelevant if you’re solving a single short problem, but the former may annoy you.”
* **“In summary, don’t use float or double for any calculations that require an exact answer.”**
  * “Use `BigDecimal` if you want the system to keep track of the decimal point and you don’t mind the inconvenience and cost of not using a primitive type.”
  * “Using `BigDecimal` has the added advantage that it gives you full control over rounding, letting you select from eight rounding modes whenever an operation that entails rounding is performed. This comes in handy if you’re performing business calculations with legally mandated rounding behavior.”
  * “Using `BigDecimal` has the added advantage that it gives you full control over rounding, letting you select from eight rounding modes whenever an operation that entails rounding is performed. This comes in handy if you’re performing business calculations with legally mandated rounding behavior.”
  * “If performance is of the essence, you don’t mind keeping track of the decimal point yourself, and the quantities aren’t too big, use `int` or `long`.”
  * “If the quantities don’t exceed nine decimal digits, you can use `int`; if they don’t exceed eighteen digits, you can use `long`. If the quantities might exceed eighteen digits, use `BigDecimal`.”


## Item 61: Prefer primitive types to boxed primitives

* “Java has a two-part type system, consisting of *primitives*, such as `int`, `double`, and `boolean`, and *reference types*, such as `String` and `List`. Every primitive type has a corresponding reference type, called a boxed primitive. The boxed primitives corresponding to `int`, `double`, and `boolean` are `Integer`, `Double`, and `Boolean`.”
* “There are three major differences between primitives and boxed primitives.”
  * “First, primitives have only their values, whereas boxed primitives have identities distinct from their values.”
  * “Second, primitive types have only fully functional values, whereas each boxed primitive type has one nonfunctional value, which is `null`, in addition to all the functional values of the corresponding primitive type.”
  * “Last, primitives are more time- and space-efficient than boxed primitives.”
* **“Applying the `==` operator to boxed primitives is almost always wrong.”**
* “In practice, if you need a comparator to describe a type’s natural order, you should simply call `Comparator.naturalOrder()`, and if you write a comparator yourself, you should use the comparator construction methods, or the static compare methods on primitive types (Item 14).”


```java
Comparator<Integer> naturalOrder = (iBoxed, jBoxed) -> {
    int i = iBoxed, j = jBoxed; // Auto-unboxing
    return i < j ? -1 : (i == j ? 0 : 1);
};
```

* “In nearly every case **when you mix primitives and boxed primitives in an operation, the boxed primitive is auto-unboxed**.”
  * “If a null object reference is auto-unboxed, you get a `NullPointerException`.”
* “So when should you use boxed primitives? They have several legitimate uses.”
  * “The first is as elements, keys, and values in collections. You can’t put primitives in collections, so you’re forced to use boxed primitives.”
  * “You must use boxed primitives as type parameters in parameterized types and methods (Chapter 5), because the language does not permit you to use primitives.”

  * “Finally, you must use boxed primitives when making reflective method invocations (Item 65).”
* **“In summary, use primitives in preference to boxed primitives whenever you have the choice. Primitive types are simpler and faster. If you must use boxed primitives, be careful! Autoboxing reduces the verbosity, but not the danger, of using boxed primitives. When your program compares two boxed primitives with the `==` operator, it does an identity comparison, which is almost certainly not what you want. When your program does mixed-type computations involving boxed and unboxed primitives, it does unboxing, and when your program does unboxing, it can throw a `NullPointerException`. Finally, when your program boxes primitive values, it can result in costly and unnecessary object creations.”**
