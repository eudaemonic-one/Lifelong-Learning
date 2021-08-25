# Chapter 4: Comments

* "Don't command bad code - rewrite it."
* The proper use of comments is to compensate for our failure to express ourself in code.
* Comments often lie. The older a comment is, and the farther away it is from the code it describes, the more likely it is to be just plain wrong.

## Comments Do Not Make Up for Bad Code

* Clear and expressive code with few comments is far superior to cluttered and cmoplex code with lots of comments.

## Explain Yourself in Code

* In many cases it's simply a matter of creating a function that says the same thing as the comment you want to write.
  * `// Check to see if the employee is eligible for full benefits` vs. `if (employee.isEligibleForFullBenefits())`

## Good Comments

### Legal Comments

* Write certain comments for legal reasons.
  * e.g., refer to a standard license or other external document.

### Informative Comments

* To explain the return value of an abstract method:
  * e.g., `// Returns an instance of the Responder being tested.` or to use the name of the function to convey the information like `responderBeingTested`.
* To explain the string format:
  * e.g., `// format matched kk:mm:ss EEE, MMM dd, yyyy`

### Explanation of Intent

* Provide the intent behind a decision.

```java
public void testConcurrentAddWidgets() throws Exception {
  WidgetBuilder widgetBuilder = new WidgetBuilder(new Class[]{BoldWidget.class});
  String text = ”’’’bold text’’’”;
  ParentWidget parent = new BoldWidget(new MockWidgetRoot(), ”’’’bold text’’’”);
  AtomicBoolean failFlag = new AtomicBoolean();
  failFlag.set(false);
   
  // This is our best attempt to get a race condition by creating large number of threads.
  for (int i = 0; i < 25000; i++) {
     WidgetBuilderThread widgetBuilderThread = new WidgetBuilderThread(widgetBuilder, text, parent, failFlag);
     Thread thread = new Thread(widgetBuilderThread);
     thread.start();
  }
  assertEquals(false, failFlag.get());
}
```

### Clarification

* Sometimes it is just helpful to translate the meaning of some obscure argument or return value into something that's readable.
* Clarifying comment can be a substantial risk if it's incorrect.

```java
assertTrue(a.compareTo(a) == 0);    // a == a
assertTrue(a.compareTo(b) != 0);    // a != b
```

### Warning of Consequences

* Sometimes it is useful to warn other programmers about certain consequences.
  * e.g., `Don't run unless you have some time to kill.` vs. `@Ignore("Takes too long to run")`
  * e.g., `“//SimpleDateFormat is not thread safe, so we need to create each instance independently.`

### TODO Comments

* It is sometime reasonable to leave "Todo" notes in the form of `//TODO` comments.
  * It might be a reminder to delete a deprecated feature or a plea for someone else to look at a problem.
  * It might be a request for someone else to think of a betetr name or a reminder to make a change that is dependent on a planned event.

```java
//TODO-MdM these are not needed
// We expect this to go away when we do the checkout model
protected VersionInfo makeVersion() throws Exception {
  return null;
}
```

### Amplification

* A comment may be used to amplify the importance of something that may otherwise seem inconsequential.

```java
String listItemContent = match.group(3).trim();
// the trim is real important.  It removes the starting spaces that could cause the item to be recognized as another list.
new ListItemWidget(this, listItemContent, this.level + 1);
return buildList(text.substring(match.end()));
```

### Javadocs in Public APIs

* Javadocs are quite helpful for public APIs.

## Bad Comments

### Mumbling

* If you decide to write a comment, then spend the time necessary to make sure it is the best comment you can write.
  * Any comment that forces you to look in another module for the meaning of that comment has failed to communicate to you and is not worth the bits it consumes.

### Redundant Comments

* Redundant comments probably take longer to read than the code itself.

```java
// Utility method that returns when this.closed is true. Throws an exception if the timeout is reached.
public synchronized void waitForClose(final long timeoutMillis) throws Exception {
  if(!closed) {
    wait(timeoutMillis);
    if(!closed)
      throw new Exception("MockResponseSender could not be closed");
  }
}
```

### Misleading Comments

* Sometimes, with all the best intentions, a programmer makes a statement in his comments that isn't precise enough to be accurate.

### Mandated Comments

* It is just plain silly to have a rule that says that every function must have a javadoc, or every variable must have a comment.

```java
/**
 * 
 * @param title The title of the CD
 * @param author The author of the CD
 * @param tracks The number of tracks on the CD
 * @param durationInMinutes The duration of the CD in minutes
 */
public void addCD(String title, String author, int tracks, int durationInMinutes) {
  CD cd = new CD();
  cd.title = title;
  cd.author = author;
  cd.tracks = tracks;
  cd.duration = duration;
  cdList.add(cd);
}
```

### Journal Comments

* Sometimes people add a comment to the start of a module every time they edit it. These comments accumulate as a kind of journal, or log, of every change that has ever been made.

### Noise Comments

* Sometimes you see comments that are nothing but noise. They restate the obvious and provide no new information.

```java
/**
  * Default constructor.
  */
protected AnnualDateRule() {}
```

* Replace the temptation to create noise with the determination to clean your code.

```java
private void startSending() {
  try {
    doSending();
  } catch(SocketException e) {
    // normal. someone stopped the request.
  } catch(Exception e) {
    try {
      response.add(ErrorResponder.makeExceptionString(e));
      response.closeAll();
    } catch(Exception e1) {
      //Give me a break!
    }
  }
}
```

```java
private void startSending() {
  try {
    doSending();
  } catch(SocketException e) {
    // normal. someone stopped the request.
  } catch(Exception e) {
    addExceptionAndCloseResponse(e);
  }
}
   
private void addExceptionAndCloseResponse(Exception e) {
  try {
    response.add(ErrorResponder.makeExceptionString(e));
    response.closeAll();
  } catch(Exception e1) {}
}
```

### Scary Noise

* cut-paste error.

### Don't Use a Comment When You Can Use a Function or a Variable

```java
// does the module from the global list <mod> depend on the subsystem we are part of?
if (smodule.getDependSubsystems().contains(subSysMod.getSubSystem()))
```

* This could be rephrased without the comment as

```java
ArrayList moduleDependees = smodule.getDependSubsystems();
String ourSubSystem = subSysMod.getSubSystem();
if (moduleDependees.contains(ourSubSystem))
```

### Position Markers

* A banner is startling and obvious if you don't see banners very often. So use them very sparingly, and only when the benefit is significant.

### Closing Brace Comments

* If you find yourself wanting to mark your closing braces, try to shorten your functions instead.

### Attributions and Bylines

* Source code control systems are very good at remembering who added what, when.

### Commented-Out Code

* Don't leave commented-out code because they may gathers and never be deleted.

### HTML Comments

* If comments are going to be extracted by some tool (like Javadoc) to appear in a Web page, then it should be the responsibility of that tool, and not the programmer, to adorn the comments with appropriate HTML.

### Nonlocal Information

* Don't offer systemwide information in the context of a local comment.

```java
/**
 * Port on which fitnesse would run. Defaults to 8082.
 *
 * @param fitnessePort
 */
public void setFitnessePort(int fitnessePort) {
  this.fitnessePort = fitnessePort;
}
```

### Too Much Information

* Don't put interesting historical discussions or irrelevant descriptions of details into your comments.

### Inobvious Connection

* The connection between a comment and the code it describes should be obvious.

```java
/*
 * start with an array that is big enough to hold all the pixels
 * (plus filter bytes), and an extra 200 bytes for header info
 */
this.pngBytes = new byte[((this.width + 1) * this.height * 3) + 200];
```

### Function Headers

* A well-chosen name for a small function that does one thing is usually better than a comment header.

### Javadocs in Nonpublic Code

* Javadocs are not intended for private code consumption.

### Example

```java
/**
 * This class Generates prime numbers up to a user specified
 * maximum.  The algorithm used is the Sieve of Eratosthenes.
 * <p>
 * Eratosthenes of Cyrene, b. c. 276 BC, Cyrene, Libya --
 * d. c. 194, Alexandria.  The first man to calculate the
 * circumference of the Earth.  Also known for working on
 * calendars with leap years and ran the library at Alexandria.
 * <p>
 * The algorithm is quite simple.  Given an array of integers
 * starting at 2.  Cross out all multiples of 2.  Find the next
 * uncrossed integer, and cross out all of its multiples.
 * Repeat untilyou have passed the square root of the maximum
 * value.
 *
 * @author Alphonse
 * @version 13 Feb 2002 atp
 */
import java.util.*;

public class GeneratePrimes {
  /**
   * @param maxValue is the generation limit.
   */
public static int[] generatePrimes(int maxValue) {
  if (maxValue >= 2) { // the only valid case
    // declarations
    int s = maxValue + 1; // size of array
    boolean[] f = new boolean[s];
    int i;
    // initialize array to true.
    for (i = 0; i < s; i++)
      f[i] = true;
    
    // get rid of known non-primes
    f[0] = f[1] = false;
    
    // sieve
    int j;
    for (i = 2; i < Math.sqrt(s) + 1; i++) {
      if (f[i]) { // if i is uncrossed, cross its multiples.
        for (j = 2 * i; j < s; j += i)
          f[j] = false; // multiple is not prime
      }
    }
    
    // how many primes are there?
    int count = 0;
    for (i = 0; i < s; i++) {
      if (f[i])
        count++; // bump count.
    }
    
    int[] primes = new int[count];
    
    // move the primes into the result
    for (i = 0, j = 0; i < s; i++) {
      if (f[i])  // if prime
        primes[j++] = i;
    }
    
    return primes;  // return the primes
  } else // maxValue < 2
    return new int[0]; // return null array if bad input.
  }
}
```

```java
/**
 * This class Generates prime numbers up to a user specified
 * maximum.  The algorithm used is the Sieve of Eratosthenes.
 * Given an array of integers starting at 2:
 * Find the first uncrossed integer, and cross out all its
 * multiples.  Repeat until there are no more multiples
 * in the array.
 */

public class PrimeGenerator {
  private static boolean[] crossedOut;
  private static int[] result;

  public static int[] generatePrimes(int maxValue) {
    if (maxValue < 2)
      return new int[0];
    else {
      uncrossIntegersUpTo(maxValue);
      crossOutMultiples();
      putUncrossedIntegersIntoResult();
      return result;
    }
  }

  private static void uncrossIntegersUpTo(int maxValue) {
    crossedOut = new boolean[maxValue + 1];
    for (int i = 2; i < crossedOut.length; i++)
      crossedOut[i] = false;
  }

  private static void crossOutMultiples() {
    int limit = determineIterationLimit();
    for (int i = 2; i <= limit; i++)
      if (notCrossed(i))
        crossOutMultiplesOf(i);
  }

  private static int determineIterationLimit() {
    // Every multiple in the array has a prime factor that
    // is less than or equal to the root of the array size,
    // so we don’t have to cross out multiples of numbers
    // larger than that root.
    double iterationLimit = Math.sqrt(crossedOut.length);
    return (int) iterationLimit;
  }

  private static void crossOutMultiplesOf(int i) {
    for (int multiple = 2*i; multiple < crossedOut.length; multiple += i)
      crossedOut[multiple] = true;
  }

  private static boolean notCrossed(int i) {
    return crossedOut[i] == false;
  }

  private static void putUncrossedIntegersIntoResult() {
    result = new int[numberOfUncrossedIntegers()];
    for (int j = 0, i = 2; i < crossedOut.length; i++)
      if (notCrossed(i))
        result[j++] = i;
  }

  private static int numberOfUncrossedIntegers() {
    int count = 0;
    for (int i = 2; i < crossedOut.length; i++)
      if (notCrossed(i))
        count++;
    return count;
  }
}
```
