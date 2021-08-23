# Chapter 3: Functions

## Small

* Functions should be very small.

### Blocks and Indenting

* Blocks within `if` statements, `else` statements, `while` statements, and so on should be one line long.
* Functions should not be large enough to hold nested structures.

## Do One Thing

* Functions should do one thing. They should do it well. They should do it only.
  * If a function does only those steps that are one level below the stated name of the function, then the function is doing one thing.
  * Doing more than "one thing" is if you can extract another function from it with a name that is not merely a restatement of its implementation.

### Sections within Functions

* Functions that do one thing cannot be reasonably divided into sections.

## One Level of Abstraction per Function

* Once details are mixed with essential concepts, more and more details tend to accrete within the function.

### Reading Code from Top to Bottom: The Stepdown Rule

* We want every function to be followed by those at the next level of abstraction so that we can read the program.

## Switch Statements

* By their nature, `switch` statements always do N things.
  * However, we can make sure that each `switch` statement is buried in a low-level class and is never repeated.
* There are several problems with this function:
  * It's large.
  * It does more than one thing.
  * It violates the Single Responsibility Principle (SRP) because there is more than one reason for it to change.
  * It violates the Open Closed Principle (OCP) because it must change whenever new types are added.
  * There are an unlimited number of other functions that will have the same structure.

```java
public Money calculatePay(Employee e) throws InvalidEmployeeType {
  switch (e.type) {
    case COMMISSIONED:
      return calculateCommissionedPay(e);
    case HOURLY:
      return calculateHourlyPay(e);
    case SALARIED:
      return calculateSalariedPay(e);
    default:
      throw new InvalidEmployeeType(e.type);
  }
}
```

* The solution to this problem is to bury the `switch` statement in the basement of an $$ABSTRACT FACTORY$$, and never let anyone see it.

```java
public abstract class Employee {
  public abstract boolean isPayday();
  public abstract Money calculatePay();
  public abstract void deliverPay(Money pay);
}
-----------------
public interface EmployeeFactory {
  public Employee makeEmployee(EmployeeRecord r) throws InvalidEmployeeType;
}
-----------------
public class EmployeeFactoryImpl implements EmployeeFactory {
  public Employee makeEmployee(EmployeeRecord r) throws InvalidEmployeeType {
    switch (r.type) {
      case COMMISSIONED:
        return new CommissionedEmployee(r) ;
      case HOURLY:
        return new HourlyEmployee(r);
      case SALARIED:
        return new SalariedEmploye(r);
      default:
        throw new InvalidEmployeeType(r.type);
    }
  }
}
```

## Use Descriptive Names

* A long descriptive name is better than a short enigmatic name.
* You should try several different names and read the code with each in place.
* Be consistent in your names.
  * Use the same phrases, nouns, and verbs in the function names you choose for your modules.

## Function Arguments

* The ideal number of arguments for a function is zero. Next comes one, followed by two. Three arguments should be avoided where possible. More than three requires very special justification.
  * Arguments take a lot of conceptual power. The argument is at a different level of abstraction than the function name and forces you to know a detail.
  * Arguments are even harder from a testing point of view. Imagine the difficulty of writing all the test cases to ensure that all the various combinations of arguments work properly.
  * Output arguments are hard to understand than input arguments because they often cause us to do a double-take.

### Common Monadic Forms

* Two very common reasons to pass a single argument into a function:
  * query (e.g., `boolean fileExists("MyFile")`)
  * transform (e.g., `InputStream fileOpen("MyFile")`)
* A less common one:
  * event (e.g., `void passwordAttemptFailedNTimes(int attempts)`)
* Try to avoid any monadic functions that don't follow these forms.
  * Using an output argument instead of a return value for a transformation is confusing.

### Flag Arguments

* Passing a boolean into a function loudly proclaims that this function does more than one thing and is plain confusing to a poor reader.
  * e.g., `render(boolean isSuite)` vs. `renderForSuite()`

### Dyadic Functions

* A function with two arguments is harder to understand than a monadic function,
  * e.g., `writeField(name)` is easier to understand than `writeField(outputStream, name)`. Whereas `outputStream` and `name` have neither a natural cohesion, nor a natural ordering.
* There are times, of course, where two arguments are appropriate.
  * e.g., `Point p = new Point(0, 0)`
  * Even obvious dyadic functions like `assertEquals(expected actual)` are problematic. The two arguments have no natural ordering.

### Triads

* Functions that take three arguments are significantly harder to understand than dyads.
  * e.g., `assertEquals(message, expected, actual)` always require a double-take to check the ordering of arguments.
  * e.g., `assertEquals(1.0, amount, .001)` is not quite so insidious.

### Argument Objects

* When a function seems to need more than two or three arguments, it is likely that some of those arguments ought to be wrapped into a class of their own.
  * e.g., `Circle makeCircle(double x, double y, double radius);` vs. `Circle makeCircle(Point center, double radius)`

### Argument Lists

* If the variable arguments are all treated identically, they are equivalent to a single argument of type `List`.
  * e.g., `public String format(String format, Object... args)`

### Verbs and Keywords

* The function and argument should form a very nice verb/noun pair.
  * e.g., `write(name)` vs. `writeField(name)`
* Using this form we encode the names of the arguments into the function name.
  * e.g., `assertEquals` vs. `assertExpectedEqualsActual(expect, actual)`

## Have No Side Effects

* Side effects are lies. Your function promises to do one thing, but it also does other *hidden* things.
  * They create strange temporal couplings and order dependencies.

### Output Arguments

* In general, output arguments should be avoided.
  * In OO languages `this` is *intended* to act as an output argument.
  * If your function must change the state of something, have it change the state of its owning object.
  * e.g., `public void appendFooter(StringBuffer report);` vs. `report.appendFooter();`

## Command Query Separation

* Functions should either do something or answer something, but nor both.
  * Either your function should change the stage of an object, or it should return some information about that object.
    * e.g., `public boolean set(String attribute, String value);` vs. `setAndCheckIfExists`.
  * The real solution is to separate the command from the query so that the ambiguity cannot occur.
    * e.g., `if (attributeExists("username")) setAttribute("username", "unclebob");`

## Prefer Exceptions to Returning Error Codes

* Returning error codes from command functions is a subtle violation of command query separation. It promotes commands being used as expressions in the predicates of `if` statements. It also leads to deeply nested structures becaue the caller must deal with error immediately.
  * e.g., `if (deletePage(page) == E_OK)`

```java
“if (deletePage(page) == E_OK) {
   if (registry.deleteReference(page.name) == E_OK) {
     if (configKeys.deleteKey(page.name.makeKey()) == E_OK){
       logger.log("page deleted");
     } else {
       logger.log("configKey not deleted");
     }
  } else {
    logger.log("deleteReference from registry failed");
  }
} else {
  logger.log("delete failed");
  return E_ERROR;
}
```

### Extract Try/Catch Blocks

* It is better to extract the bodies of the `try` and `catch` blocks out into functions of their own to avoid mixing error processing with normal processings.

```java
public void delete(Page page) {
  try {
    deletePageAndAllReferences(page);
  } catch (Exception e) {
    logError(e);
  }
}

private void deletePageAndAllReferences(Page page) throws Exception {
  deletePage(page);
  registry.deleteReference(page.name);
  configKeys.deleteKey(page.name.makeKey());
}

private void logError(Exception e) {
  logger.log(e.getMessage());
}
```

### Error Handling Is One Thing

* A function that handles errors should do nothing else.
  * If the keyword `try` exists in a function, it should be the very first word in the function and that there should be nothing after the `catch`/`finally` blocks.
* Returning error codes usually implies that there is some class or enum in which all the error codes are defined, which enforces other classes to be recomplied and redeployed when the `Error` enum changes.
* When you use exceptions rather than error codes, then new exceptions are *derivatives* of the exception class, which can be added without recompilation or redeployment.

## Don't Repeat Yourself

* Duplication makes it hard to make modification and creates more opportunities for errors of omission.

## Structured Programming

* Dijkstra said that every function, and every block within a function, should have one entry and one exit.
  * So if you keep functions small, then the occasional multiple `return`, `break`, or `continue` statement does no harm and can sometimes even be more expressive than the single-entry, single-exit rule.
  * On the other hand, `goto` only makes sense in large functions, so it should be avoided.

## How Do You Write Functions Like This

* Refine the code, split out functions, change names, eliminate duplication.

## Conclusion

* “If you follow the rules herein, your functions will be short, well named, and nicely organized. But never forget that your real goal is to tell the story of the system, and that the functions you write need to fit cleanly together into a clear and precise language to help you with that telling.”
