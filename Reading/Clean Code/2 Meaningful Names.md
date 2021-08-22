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
