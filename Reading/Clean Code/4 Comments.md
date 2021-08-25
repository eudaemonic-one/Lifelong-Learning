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
