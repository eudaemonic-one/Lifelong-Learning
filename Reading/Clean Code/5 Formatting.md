# Chapter 5: Formatting

## The Purpose of Formatting

* The coding style and readability set precedents that continue to affect maintainability and extensibility long after the original code has been changed beyond recognition.

## Vertical Formatting

* In Java, file size is closely related to class size.
  * It appears to be possible to build significant systems out of files that are typically 200 lines long, with an upper limit of 500.

### The Newspaper Metaphor

* We would like a source file to be like a newspaper article. The name should be simple but explanatory. The name, by itself, should be sufficient to tell us whether we are in the right module or not. The topmost parts of the source file should provide the high-level concepts and algorithms. Detail should increase as we move downward, until at the end we find the lowest level functions and details in the source file.

### Vertical Openness Between Concepts

* Each expression or a clause, and each group of lines should be separated from each other with blank lines.

### Vertical Density

* Lines of code that are tightly related should appear vertically dense.

### Vertical Distance

* Concepts that are closely related should be kept vertically close to each other.
* Closely related concepts should not be separated into different files unless you have a very good reason.
* **Variable Declarations**
  * Variables should be declared as close to their usage as possible.
  * Control variables for loops should usually be declared within the loop statement.
* **Instance variables** should be declared at the top of the class.
* **Dependent Functions**
  * If one function calls another, they should be vertically close, and the caller should be above the callee. if at all possible. This gives the program a natural flow.
* **Conceptual Affinity**
  * Affinity might be caused because a group of functions perform a similar operation.

### Vertical Ordering

* A function that is called should be below a function that does the calling

## Horizontal Formatting

* We should strive to keep our lines short.
  * Set line length limit to 120 is reasonable.

### Horizontal Openness and Density

* Use horizontal white spaces to associate things that are strongly related and disassociate things that are more weakly related.
  * Surround the assignment operatos with white spaces to accentuate them.

### Horizontal Alignment

* If there are long lists that need to be aligned, the problem is the length of the lists, not the lack of alignment.

### Indentation

* There is information that pertains to the file as a whole, to the individual classes within the file, to the methods within the classes, to the blocks within the methods, and recursively to the blocks within the blocks.
* Programmers visually line up lines on the left to see what scope they appear in.
* **Breaking Indentation**
  * Go back and put the indentation back if the indentation rule was broken.

### Dummy Scopes

* Try to avoid dummy statements like `while` or `for`.
  * Unless you make that semicolon *visible* by indenting it on it's own line, it's just too hard to see.

## Team Rules

* A team of developers should agree upon a single formatting style, and then every member of that team should use that style.

## Uncle Bob's Formatting Rules

```java
public int getWidestLineNumber() {
  return widestLineNumber;
}

public LineWidthHistogram getLineWidthHistogram() {
  return lineWidthHistogram;
}

public double getMeanLineWidth() {
  return (double)totalChars/lineCount;
}

public int getMedianLineWidth() {
  Integer[] sortedWidths = getSortedWidths();
  int cumulativeLineCount = 0;
  for (int width : sortedWidths) {
    cumulativeLineCount += lineCountForWidth(width);
    if (cumulativeLineCount > lineCount/2)
      return width;
  }
  throw new Error(“Cannot get here”);
}

private int lineCountForWidth(int width) {
  return lineWidthHistogram.getLinesforWidth(width).size();
}

private Integer[] getSortedWidths() {
  Set<Integer> widths = lineWidthHistogram.getWidths();
  Integer[] sortedWidths = (widths.toArray(new Integer[0]));
  Arrays.sort(sortedWidths);
  return sortedWidths;
}
```
