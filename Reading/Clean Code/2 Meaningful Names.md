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
