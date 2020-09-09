# Lecture 3 The Process of API Design - How to build good (and bad) APIs

## How to Design Good APIs

* **Gather requirements** skeptically, including use cases
  * Often you'll get proposed solutions instead
    * Better solutions may exist
  * Your job is to extract true requirements
    * You need **use-cases**; if you don't get them, keep trying
  * You may get requirements that don't make sense
    * Ask questions until you see eye-to-eye
  * You may get requirements that are wrong
    * Push back
  * You may get requirements that are contradictory
    * Broker a compromise
  * Requirements will change as you proceed
  * Key question: **What problems should API solve?**
  * Also important: **What problems shouldn't API solve?**
  * Maintain a **requirements document**
  * **Ask yourself if the API should be designed**
    * Superfluous, impossible, unethical, too vague
    * If the problem can't be fixed, **fail fast**!
* **Choose an abstraction** (model) that appears to address use cases
  * **Embed use cases in an underlying structure**
    * Note their similarities and differences
  * Generally, a model will emerge
* **Compose a short API sketch** for abstraction
  * **Start with short spec - one page is ideal**
  * **At this stage, comprehensibility and agility are more important than completeness**
  * Bounce spec off as many people as possible
    * API Design is **not** a solitary activity
  * **Don't** fall in love with your spec too soon
* **Apply API sketch** to use cases to see if it works
  * Start before you've implemented the API
  * Start before you've even specified it properly
  * Continue writing to API as you flesh it out
  * Early code lives on as **examples**, unit tests
    * **Among the most important code you'll ever write**
* **Show API** to anyone who will look at it
* **Write prototype** implementation of API
  * You will find embarrasing errors in your API
    * Fix them and move on
  * You may also find subtle performance problems
* **Flesh out** the documentation & harden implementation
  * Now you have an artifact you can share more widely
  * Distribute, but ensure people know it's subject to change
  * If you're lucky, you'll get bug reports & feature requests
  * Use the API feedback while you can
* **Keep refining it** as long as you can
  * Try API on at least 3 use cases before release
  * Ideally, get different people to write the use casesHow to design bad APIs

### Issue Tracking

* Throughout process, maintain a list of design issues
  * Write down all the options you know of
  * Say which were ruled out and why
  * When you decide, say which was chosen and why
* Prevent wasting time on solved issues
* Provides rationale for the resulting API

### Key Design Artifacts

* Requirements document
* Issues list
* Use-case code
* Maintain throughout design and retain when done
* They guide the design process

## How to Design Bad APIs

## An API Design Case

* What shape should the API have

  * **Static utility method**
    * Simpler
    * Nicer client code
    * Can fit nicely into a preexitsting class of collection utils
  * Instantiable "chooser" class
    * More flexible
    * Could cause performance problems

* What type should the input be (and why)?

  * `Collection<T>`
    * Can access elements with `iterator()` or `toArray()`
    * Can't save element list across calls
  * `List<T>`
  * `Stream<T>`
    * Streams can be infinitely long
    * Mathematically speaking, result isn't well defined
  * `T[]` (array)

* ```java
  public class RandomChooser {
    // Returns a uniform randomly chosen object from given objects.
    // @throws IllegalArgumentException if list is empty
    public static <T> T randomElement(List<T> objects);
  }
  
  // Design 1b
  public class RandomChooser {
    // Returns a uniform randomly chosen object from given objects.
    // @throws IllegalArgumentException if given objects is empty
    public static <T> T randomElement(List<T> objects);
    
    // Returns a uniform randomly chosen object from given objects using given source of randomness. Use this version only when you need to specify a source of randomness.
    // @throws IllegalArgumentException if list is empty
    public static <T> T randomElement(List<T> objects, Random rnd);
  }
  
  // Design 1c
  public class RandomChooser {
    // Returns a uniform randomly chosen object from given enum type.
    public static <T extends Enum<T>> T randomElement(Class<T> type) {
      return randomElement(Arrays.asList(type.getEnumConstants()));
    }
  }
  ```

* Is it really worth having an API whose implementation is this small?
  * Yes
    * It guides programmers to the functionality
  * Martin Fowlers calls these **highlight methods**
* Random selector API design summary
  * Correct input type was not obvious
    * Wrong type can make API slow or unimplementable
  * It's easy to forget corner cases
  * Convenience methods can make a huge difference
  * Platform influences API design
  * Even tiny APIs with tiny implementations can be very useful