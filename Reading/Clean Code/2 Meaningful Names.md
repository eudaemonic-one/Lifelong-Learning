# Chapter 2: Meaningful Names

## Use Intention-Revealing Names

* Choosing good names takes time but saves more than it takes.
* The name should tell you why it exists, what it does, and how it is used.
  * `int d; // elapsed time in days` => reveals nothing; can be replaced with:
    * `int elapsedTimeInDays;`
    * `int daysSinceCreation;`
* Choosing names that reveal intent can make it much easier to understand and change code.

```java
public List<int[]> getThem() {
  List<int[]> list1 = new ArrayList<int[]>();
  for (int[] x : theList)
    if (x[0] == 4)
      list1.add(x);
  return list1;
}
```

* With these simple name changes, it’s not difficult to understand what’s going on.

```java
public List<int[]> getFlaggedCells() {
  List<int[]> flaggedCells = new ArrayList<int[]>();
  for (int[] cell : gameBoard)
    if (cell[STATUS_VALUE] == FLAGGED)
      flaggedCells.add(cell);
  return flaggedCells;
}
```

## Avoid Disinformation

* We should avoid words whose entrenched meanings vary from our intended meaning.
  * Do not refer to a grouping of accounts as an `accountList` unless it's actually a `List`.
* Beware of using names which vary in small ways.
  * `XYZControllerForEfficientHandlingOfStrings` vs. `XYZControllerForEfficientStorageOfStrings`
* A truly awful example of disinformative names would be the use of lower-case `L` or uppercase `O` as variable names.

## Make Meaningful Distinctions

* Number-series naming (`a1`, `a2`, ..., `aN`) is the opposite of intentional naming.
* Noise words are another meaningless distinction.
  * Imagine that you have a `Product` class. If you have another called `ProductInfo` or `ProductData`, you have made the name different without making them mean anything different.
* Noise words are redundant.
  * The word `variable` should never appear in a variable name. The word `table` should never appear in a table name.
* Distinguish names in such a way that the reader knows what the differences offer.
  * In the absence of specific conventions, the variable `moneyAmount` is indistinguishable from `money`, `customerInfo` is indistinguishable from `customer`, `accountData` is indistinguishable from `account`, `theMessage` is indistinguishable from `message`.

## Use Pronounceable Names

* If you can't pronounce it, you can't discuss it without sounding like an idiot.
  * `private Date genymdhms;` vs. `private Date generationTimestamp;`

## Use Searchable Names

* Single-letter names and numeric constants have a particular problem in that they are not easy to locate across a body of text.
* The length of a name should correspond to the size of its scope.

## Avoid Encodings

* Encoded names are seldom pronounceable and are easy to mis-type.

### Hungarian Notation

* In early days, the programmers need Hungarian Notation to help them remember the types.
  * e.g., `phoneString`, `voidPtr`.
* In modern languages we have much richer type systems, and the compiler remember and enforce the types.
  * Java programmers don't need type encoding.
  * So nowadays, HN makes it harder to change the name and read the code, and easy to mislead the reader.

### Member Prefixes

* You don't need to prefix member variables with `m_` anymore.

### Interfaces and Implementations

* Prefer to leave interfaces unadorned.
  * `IShapeFactory` vs. `ShapeFactory`?

## Avoid Mental Mapping

* Readers shouldn't have to mentally translate your names into other names they already know.
  * Generally arises from a choice to use neither problem domain terms nor solution domain terms.
  * In most other contexts than for-loops a single-letter name is a poor choice; it's just a place holder that the reader must mentally map to the actual concept.
* The professional understands that *clarity is king*.

## Class Names

* Classes and objects should have noun or noun phrase names like `Customer`, `WikiPage`, `Account`, and `AddressParser`.
* Avoid words like `Manager`, `Processor`, `Data`, or `Info` in the name of a class.

## Method Names

* Methods should have verb or verb phrase names like `postPayment`, `deletePage`, or `save`.
* Accessors, mutators, and predicates should be named for their value and prefixed with `get`, `set`, and `is` according to the javabean standard.
* When constructors are overloaded, use static factory methods with names that describe the arguments.
  * Consider enforcing their use by making the corresponding constructors private.

## Don't Be Cute

* Choose clarity over entertainment value.
  * e.g., `whack()` vs. `kill()`, `eatMyShorts` vs. `abort()`.

## Pick One Word per Concept

* Pick one word for one abstract concept and stick with it.
  * e.g., have `fetch`, `retrieve`, `get` as equivalent methods of different classes is confusing.
* The names ahve to be consistent in order for you to pick the correct method without any additional exploration.
  * e.g., have a `controller` and a `manager` and a `driver` in the same code base is confusing.

## Don't Pun

* Avoid using the same word for two purposes.
  * e.g., `add`, `insert`, `append`.
