Chapter 7. Lambdas and Streams

## Item 42: Prefer lambdas to anonymous classes

* “Historically, interfaces (or, rarely, abstract classes) with a single abstract method were used as *function types*. Their instances, known as *function objects*, represent functions or actions.”
* “Since JDK 1.1 was released in 1997, the primary means of creating a function object was the *anonymous class* (Item 24).”
* “Anonymous classes were adequate for the classic objected-oriented design patterns requiring function objects, notably the *Strategy* pattern [Gamma95].”


```java
// Anonymous class instance as a function object - obsolete!
Collections.sort(words, new Comparator<String>() {
    public int compare(String s1, String s2) {
        return Integer.compare(s1.length(), s2.length());
    }
});
```

* “The `Comparator` interface represents an *abstract strategy* for sorting; the anonymous class above is a *concrete strategy* for sorting strings.”
* “The verbosity of anonymous classes, however, made functional programming in Java an unappealing prospect.”
* “In Java 8, the language formalized the notion that interfaces with a single abstract method are special and deserve special treatment. These interfaces are now known as *functional interfaces*, and the language allows you to create instances of these interfaces using *lambda expressions*, or *lambdas* for short.”


```java
// Lambda expression as function object (replaces anonymous class)
Collections.sort(words,
        (s1, s2) -> Integer.compare(s1.length(), s2.length()));
```

* “Note that the types of the lambda (`Comparator<String>`), of its parameters (`s1` and `s2`, both `String`), and of its return value (`int`) are not present in the code. The compiler deduces these types from context, using a process known as *type inference*. ”
* **“Omit the types of all lambda parameters unless their presence makes your program clearer.”**
  * “If the compiler generates an error telling you it can’t infer the type of a lambda parameter, then specify it.”
  * “Sometimes you may have to cast the return value or the entire lambda expression, but this is rare.”
* “Incidentally, the comparator in the snippet can be made even more succinct if a *comparator construction method* is used in place of a lambda (Items 14. 43):”


```java
Collections.sort(words, comparingInt(String::length));
```

* “In fact, the snippet can be made still shorter by taking advantage of the `sort` method that was added to the `List` interface in Java 8:”


```java
words.sort(comparingInt(String::length));
```

* “Item 34 says that enum instance fields are preferable to constant-specific class bodies. Lambdas make it easy to implement constant-specific behavior using the former instead of the latter. ”


```java
// Enum with function object fields & constant-specific behavior
public enum Operation {
    PLUS  ("+", (x, y) -> x + y),
    MINUS ("-", (x, y) -> x - y),
    TIMES ("*", (x, y) -> x * y),
    DIVIDE("/", (x, y) -> x / y);

    private final String symbol;
    private final DoubleBinaryOperator op;

    Operation(String symbol, DoubleBinaryOperator op) {
        this.symbol = symbol;
        this.op = op;
    }

    @Override public String toString() { return symbol; }

    public double apply(double x, double y) {
        return op.applyAsDouble(x, y);
    }
}
```

* “Unlike methods and classes, **lambdas lack names and documentation; if a computation isn’t self-explanatory, or exceeds a few lines, don’t put it in a lambda**.”
  * “One line is ideal for a lambda, and three lines is a reasonable maximum.”
* “Lambdas share with anonymous classes the property that you can’t reliably serialize and deserialize them across implementations.”
  * “Therefore, **you should rarely, if ever, serialize a lambda** (or an anonymous class instance).”
  * “If you have a function object that you want to make serializable, such as a `Comparator`, use an instance of a private static nested class (Item 24).”
* **“Don’t use anonymous classes for function objects unless you have to create instances of types that aren’t functional interfaces.”**

## Item 43: Prefer method references to lambdas

* **“Java provides a way to generate function objects even more succinct than lambdas: *method references*.”**

```java
map.merge(key, 1, (count, incr) -> count + incr);
```

* “Note that this code uses the `merge` method, which was added to the `Map` interface in Java 8. If no mapping is present for the given key, the method simply inserts the given value; if a mapping is already present, `merge` applies the given function to the current value and the given value and overwrites the current value with the result. This code represents a typical use case for the `merge` method.”
* “The code reads nicely, but there’s still some boilerplate. The parameters `count` and `incr` don’t add much value, and they take up a fair amount of space. Really, all the lambda tells you is that the function returns the sum of its two arguments. As of Java 8, `Integer` (and all the other boxed numerical primitive types) provides a static method `sum` that does exactly the same thing.”

```java
map.merge(key, 1, Integer::sum);
```

* “The more parameters a method has, the more boilerplate you can eliminate with a method reference. In some lambdas, however, the parameter names you choose provide useful documentation, making the lambda more readable and maintainable than a method reference, even if the lambda is longer.”

| Method Ref Type   | Example                  | Lambda Equivalent                                |
| ----------------- | ------------------------ | ------------------------------------------------ |
| Static            | `Integer::parseInt`      | `str -> Integer.parseInt(str)`                   |
| Bound             | `Instant.now()::isAfter` | `Instant then = Instant.now(); then.isAfter(t);` |
| Unbound           | `String::toLowerCase`    | `str -> str.toLowerCase()`                       |
| Class Constructor | `TreeMap<K,V>::new`      | `() -> new TreeMap<K,V>`                         |
| Array Constructor | `int[]::new`             | `len -> new int[len]`                            |

* “In summary, method references often provide a more succinct alternative to lambdas. **Where method references are shorter and clearer, use them; where they aren’t, stick with lambdas.**”


## Item 44: Favor the use of standard functional interfaces

* “Now that Java has lambdas, best practices for writing APIs have changed considerably. For example, the *Template Method* pattern [Gamma95], wherein a subclass overrides a primitive method to specialize the behavior of its superclass, is far less attractive. The modern alternative is to provide a static factory or constructor that accepts a function object to achieve the same effect. More generally, you’ll be writing more constructors and methods that take function objects as parameters.”

```java
// Unnecessary functional interface; use a standard one instead.
@FunctionalInterface interface EldestEntryRemovalFunction<K,V>{
    boolean remove(Map<K,V> map, Map.Entry<K,V> eldest);
}
```

* “This interface would work fine, but you shouldn’t use it, because you don’t need to declare a new interface for this purpose. The `java.util.function` package provides a large collection of standard functional interfaces for your use.”
* **“If one of the standard functional interfaces does the job, you should generally use it in preference to a purpose-built functional interface.”**
* “There are forty-three interfaces in `java.util.Function`. You can’t be expected to remember them all, but if you remember six basic interfaces, you can derive the rest when you need them.”
  * “The `Operator` interfaces represent functions whose result and argument types are the same.”
  * “The `Predicate` interface represents a function that takes an argument and returns a boolean.”
  * “The `Function` interface represents a function whose argument and return types differ.”
  * “The `Supplier` interface represents a function that takes no arguments and returns (or “supplies”) a value. ”
  * “Finally, `Consumer` represents a function that takes an argument and returns nothing, essentially consuming its argument.”

| Interface           | Function Signature    | Example               |
| ------------------- | --------------------- | --------------------- |
| `UnaryOperator<T>`  | `T apply(T t)`        | `String::toLowerCase` |
| `BinaryOperator<T>` | `T apply(T t1, T t2)` | `BigInteger::add`     |
| `Predicate<T>`      | `boolean test(T t)`   | `Collection::isEmpty` |
| `Function<T,R>`     | `R apply(T t)`        | `Arrays::asList`      |
| `Supplier<T>`       | `T get()`             | `Instant::now`        |
| `Consumer<T>`       | `void accept(T t)`    | `System.out::println` |

* “There are also three variants of each of the six basic interfaces to operate on the primitive types `int`, `long`, and `double`. Their names are derived from the basic interfaces by prefixing them with a primitive type. So, for example, a predicate that takes an `int` is an `IntPredicate`, and a binary operator that takes two `long` values and returns a `long` is a `LongBinaryOperator`. ”
* “There are nine additional variants of the `Function` interface, for use when the result type is primitive. The source and result types always differ, because a function from a type to itself is a `UnaryOperator`. If both the source and result types are primitive, prefix `Function` with `SrcToResult`, for example `LongToIntFunction (six variants)`.”
* “There are two-argument versions of the three basic functional interfaces for which it makes sense to have them: `BiPredicate<T,U>`, `BiFunction<T,U,R>`, and `BiConsumer<T,U>`. ”
* “There are also `BiFunction` variants returning the three relevant primitive types: `ToIntBiFunction<T,U>`, `ToLongBiFunction<T,U>`, and `ToDoubleBiFunction<T,U>`.”
* “There are two-argument variants of `Consumer` that take one object reference and one primitive type: `ObjDoubleConsumer<T>`, `ObjIntConsumer<T>`, and `ObjLongConsumer<T>`.”

* “Most of the standard functional interfaces exist only to provide support for primitive types. **Don’t be tempted to use basic functional interfaces with boxed primitives instead of primitive functional interfaces.**”
* “Finally, there is the `BooleanSupplier` interface, a variant of `Supplier` that returns `boolean` values. This is the only explicit mention of the `boolean` type in any of the standard functional interface names, but `boolean` return values are supported via `Predicate` and its four variant forms.”
* “Consider our old friend `Comparator<T>`, which is structurally identical to the `ToIntBiFunction<T,T>` interface. Even if the latter interface had existed when the former was added to the libraries, it would have been wrong to use it. There are several reasons that `Comparator` deserves its own interface. ”
  * “First, its name provides excellent documentation every time it is used in an API, and it’s used a lot.”
  * “Second, the `Comparator` interface has strong requirements on what constitutes a valid instance, which comprise its *general contract*. By implementing the interface, you are pledging to adhere to its contract.”
  * “Third, the interface is heavily outfitted with useful default methods to transform and combine comparators.”
* **“You should seriously consider writing a purpose-built functional interface in preference to using a standard one if you need a functional interface that shares one or more of the following characteristics with `Comparator`:”**
  * “It will be commonly used and could benefit from a descriptive name.”
  * “It has a strong contract associated with it.”
  * “It would benefit from custom default methods.”
* **“Always annotate your functional interfaces with the `@FunctionalInterface` annotation.”**
  * “It is a statement of programmer intent that serves three purposes: it tells readers of the class and its documentation that the interface was designed to enable lambdas; it keeps you honest because the interface won’t compile unless it has exactly one abstract method; and it prevents maintainers from accidentally adding abstract methods to the interface as it evolves.”
* **“Do not provide a method with multiple overloadings that take different functional interfaces in the same argument position if it could create a possible ambiguity in the client.”**
* **“In summary, now that Java has lambdas, it is imperative that you design your APIs with lambdas in mind. Accept functional interface types on input and return them on output. It is generally best to use the standard interfaces provided in `java.util.function.Function`, but keep your eyes open for the relatively rare cases where you would be better off writing your own functional interface.”**

## Item 45: Use streams judiciously

* “The streams API was added in Java 8 to ease the task of performing bulk operations, sequentially or in parallel.”
* “This API provides two key abstractions: the *stream*, which represents a finite or infinite sequence of data elements, and the *stream pipeline*, which represents a multistage computation on these elements. ”
* “The elements in a stream can come from anywhere. Common sources include collections, arrays, files, regular expression pattern matchers, pseudorandom number generators, and other streams.”
* “The data elements in a stream can be object references or primitive values. Three primitive types are supported: `int`, `long`, and `double`.”
* “A stream pipeline consists of a source stream followed by zero or more *intermediate operations* and one *terminal operation*.”
* “**Stream pipelines are evaluated *lazily***: evaluation doesn’t start until the terminal operation is invoked, and data elements that aren’t required in order to complete the terminal operation are never computed.”
* “**The streams API is *fluent***: it is designed to allow all of the calls that comprise a pipeline to be chained into a single expression. ”
* “By default, stream pipelines run sequentially.”
  * “Making a pipeline execute in parallel is as simple as invoking the `parallel` method on any stream in the pipeline, but it is seldom appropriate to do so (Item 48).”
* “The streams API is sufficiently versatile that practically any computation can be performed using streams, but just because you can doesn’t mean you should.”
* **“Overusing streams makes programs hard to read and maintain.”**

```java
// Tasteful use of streams enhances clarity and conciseness
public class Anagrams {
   public static void main(String[] args) throws IOException {
      Path dictionary = Paths.get(args[0]);
      int minGroupSize = Integer.parseInt(args[1]);

      try (Stream<String> words = Files.lines(dictionary)) {
         words.collect(groupingBy(word -> alphabetize(word)))
           .values().stream()
           .filter(group -> group.size() >= minGroupSize)
           .forEach(g -> System.out.println(g.size() + ": " + g));
      }
   }

   // alphabetize method is the same as in original version
}
```

* **“In the absence of explicit types, careful naming of lambda parameters is essential to the readability of stream pipelines.”**
* “**Using helper methods is even more important for readability in stream pipelines than in iterative code** because pipelines lack explicit type information and named temporary variables.”
* “Ideally, you should **refrain from using streams to process `char` values**.”
* **“Refactor existing code to use streams and use them in new code only where it makes sense to do so.”**
* “Stream pipelines express repeated computation using function objects (typically lambdas or method references), while iterative code expresses repeated computation using code blocks.”
* “There are some things you can do from code blocks that you can’t do from function objects:”
  * “From a code block, you can read or modify any local variable in scope; from a lambda, you can only read final or effectively final variables [JLS 4.12.4], and you can’t modify any local variables.”
  * “From a code block, you can `return` from the enclosing method, `break` or `continue` an enclosing loop, or throw any checked exception that this method is declared to throw”
  * **“If a computation is best expressed using these techniques, then it’s probably not a good match for streams.”**
* “Conversely, streams make it very easy to do some things:”
  * “Uniformly transform sequences of elements”
  * “Filter sequences of elements”
  * “Combine sequences of elements using a single operation (for example to add them, concatenate them, or compute their minimum)”
  * “Accumulate sequences of elements into a collection, perhaps grouping them by some common attribute”
  * “Search a sequence of elements for an element satisfying some criterion”
  * **“If a computation is best expressed using these techniques, then it is a good candidate for streams.”**
* “One thing that is hard to do with streams is to access corresponding elements from multiple stages of a pipeline simultaneously: once you map a value to some other value, the original value is lost.”
  * “One workaround is to map each value to a *pair object* containing the original value and the new value, but this is not a satisfying solution, especially if the pair objects are required for multiple stages of a pipeline.”
  * “When it is applicable, a better workaround is to invert the mapping when you need access to the earlier-stage value.”

```java
// Iterative Cartesian product computation
private static List<Card> newDeck() {
    List<Card> result = new ArrayList<>();
    for (Suit suit : Suit.values())
        for (Rank rank : Rank.values())
            result.add(new Card(suit, rank));
    return result;
}
```

```java
// Stream-based Cartesian product computation
private static List<Card> newDeck() {
    return Stream.of(Suit.values())
        .flatMap(suit ->
            Stream.of(Rank.values())
                .map(rank -> new Card(suit, rank)))
        .collect(toList());
}
```

* “If you’re not sure which version you prefer, the iterative version is probably the safer choice. If you prefer the stream version and you believe that other programmers who will work with the code will share your preference, then you should use it.”
* “In summary, some tasks are best accomplished with streams, and others with iteration. Many tasks are best accomplished by combining the two approaches. There are no hard and fast rules for choosing which approach to use for a task, but there are some useful heuristics. In many cases, it will be clear which approach to use; in some cases, it won’t.”
* **“If you’re not sure whether a task is better served by streams or iteration, try both and see which works better.”**

## Item 46: Prefer side-effect-free functions in streams

* “Streams isn’t just an API, it’s a paradigm based on functional programming.”
* “In order to obtain the expressiveness, speed, and in some cases parallelizability that streams have to offer, you have to adopt the paradigm as well as the API.”
* **“The most important part of the streams paradigm is to structure your computation as a sequence of transformations where the result of each stage is as close as possible to a *pure function* of the result of the previous stage.”**
  * “A pure function is one whose result depends only on its input: it does not depend on any mutable state, nor does it update any state.”
  * “In order to achieve this, any function objects that you pass into stream operations, both intermediate and terminal, should be free of side-effects.”

```java
// Uses the streams API but not the paradigm--Don't do this!
Map<String, Long> freq = new HashMap<>();
try (Stream<String> words = new Scanner(file).tokens()) {
    words.forEach(word -> {
        freq.merge(word.toLowerCase(), 1L, Long::sum);
    });
}
```

* “What’s wrong with this code? After all, it uses streams, lambdas, and method references, and gets the right answer. Simply put, it’s not streams code at all; it’s iterative code masquerading as streams code. It derives no benefits from the streams API, and it’s (a bit) longer, harder to read, and less maintainable than the corresponding iterative code. The problem stems from the fact that this code is doing all its work in a terminal `forEach` operation, using a lambda that mutates external state (the frequency table).”
  * “A `forEach` operation that does anything more than present the result of the computation performed by a stream is a “bad smell in code,” as is a lambda that mutates state.”

```java
// Proper use of streams to initialize a frequency table
Map<String, Long> freq;
try (Stream<String> words = new Scanner(file).tokens()) {
    freq = words
        .collect(groupingBy(String::toLowerCase, counting()));
}
```

* **“The `forEach` operation should be used only to report the result of a stream computation, not to perform the computation.”**
  * “Occasionally, it makes sense to use `forEach` for some other purpose, such as adding the results of a stream computation to a preexisting collection.”
* “The improved code uses a *collector*, which is a new concept that you have to learn in order to use streams.”
  * “For starters, you can ignore the `Collector` interface and think of a collector as an opaque object that encapsulates a *reduction* strategy.”
* “The collectors for gathering the elements of a stream into a true Collection are straightforward. There are three such collectors: `toList()`, `toSet()`, and `toCollection(collectionFactory)`.”


```java
// Pipeline to get a top-ten list of words from a frequency table
List<String> topTen = freq.keySet().stream()
    .sorted(comparing(freq::get).reversed())
    .limit(10)
    .collect(toList());
```

* “The `comparing` method is a comparator construction method (Item 14) that takes a key extraction function.”

* **“It is customary and wise to statically import all members of `Collectors` because it makes stream pipelines more readable.”**
* “So what about the other thirty-six methods in `Collectors`? Most of them exist to let you collect streams into maps, which is far more complicated than collecting them into true collections. Each stream element is associated with a *key* and a *value*, and multiple stream elements can be associated with the same key.”
* “The simplest map collector is `toMap(keyMapper, valueMapper)`, which takes two functions, one of which maps a stream element to a key, the other, to a value.”
  * “If multiple stream elements map to the same key, the pipeline will terminate with an `IllegalStateException`.”


```java
// Using a toMap collector to make a map from string to enum
private static final Map<String, Operation> stringToEnum =
    Stream.of(values()).collect(
        toMap(Object::toString, e -> e));
```

* “The more complicated forms of `toMap`, as well as the `groupingBy` method, give you various ways to provide strategies for dealing with such collisions.”
  * “One way is to provide the `toMap` method with a *merge function* in addition to its key and value mappers.”
* “The three-argument form of `toMap` is also useful to make a map from a key to a chosen element associated with that key. ”

```java
// Collector to generate a map from key to chosen element for key
Map<Artist, Album> topHits = albums.collect(
   toMap(Album::artist, a->a, maxBy(comparing(Album::sales))));
```

* “Another use of the three-argument form of `toMap` is to produce a collector that imposes a last-write-wins policy when there are collisions.”


```java
// Collector to impose last-write-wins policy
toMap(keyMapper, valueMapper, (v1, v2) -> v2)”
```

* “The third and final version of `toMap` takes a fourth argument, which is a map factory, for use when you want to specify a particular map implementation such as an `EnumMap` or a `TreeMap`.”
* “There are also variant forms of the first three versions of `toMap`, named `toConcurrentMap`, that run efficiently in parallel and produce `ConcurrentHashMap` instances.”
* “In addition to the `toMap` method, the `Collectors` API provides the `groupingBy` method, which returns collectors to produce maps that group elements into categories based on a *classifier function*.”
  * “The classifier function takes an element and returns the category into which it falls. This category serves as the element’s map key.”
  * “The simplest version of the `groupingBy` method takes only a classifier and returns a map whose values are lists of all the elements in each category.”


```java
words.collect(groupingBy(word -> alphabetize(word)))
```

* “If you want `groupingBy` to return a collector that produces a map with values other than lists, you can specify a *downstream collector* in addition to a classifier.”
  * “A downstream collector produces a value from a stream containing all the elements in a category.”
  * “The simplest use of this parameter is to pass `toSet()`, which results in a map whose values are sets of elements rather than lists.”
* “Alternatively, you can pass `toCollection(collectionFactory)`, which lets you create the collections into which each category of elements is placed. This gives you the flexibility to choose any collection type you want.”
* “Another simple use of the two-argument form of `groupingBy` is to pass `counting()` as the downstream collector. This results in a map that associates each category with the number of elements in the category, rather than a collection containing the elements.”


```java
Map<String, Long> freq = words
        .collect(groupingBy(String::toLowerCase, counting()));
```

* “The third version of `groupingBy` lets you specify a map factory in addition to a downstream collector. Note that this method violates the standard telescoping argument list pattern: the `mapFactory` parameter precedes, rather than follows, the `downStream` parameter.”
  * “This version of `groupingBy` gives you control over the containing map as well as the contained collections, so, for example, you can specify a collector that returns a `TreeMap` whose values are `TreeSets`.”
* “The `groupingByConcurrent` method provides variants of all three overloadings of `groupingBy`. These variants run efficiently in parallel and produce `ConcurrentHashMap` instances. ”
* “There is also a rarely used relative of `groupingBy` called `partitioningBy`. In lieu of a classifier method, it takes a predicate and returns a map whose key is a `Boolean`.”
* “The collectors returned by the `counting` method are intended *only* for use as downstream collectors. The same functionality is available directly on `Stream`, via the `count` method, so there is never a reason to say `collect(counting())`. ”
* “There are fifteen more `Collectors` methods with this property. They include the nine methods whose names begin with `summing`, `averaging`, and `summarizing` (whose functionality is available on the corresponding primitive stream types).”
* “They also include all overloadings of the `reducing` method, and the `filtering`, `mapping`, `flatMapping`, and `collectingAndThen` methods.”
* “There are three `Collectors` methods we have yet to mention. Though they are in `Collectors`, they don’t involve collections.”
  * “The first two are `minBy` and `maxBy`, which take a comparator and return the minimum or maximum element in the stream as determined by the comparator.”
* “The final `Collectors` method is `joining`, which operates only on streams of `CharSequence` instances such as strings. In its parameterless form, it returns a collector that simply concatenates the elements.”
  * “Its one argument form takes a single `CharSequence` parameter named `delimiter` and returns a collector that joins the stream elements, inserting the delimiter between adjacent elements.”
  * “The three argument form takes a prefix and suffix in addition to the delimiter. The resulting collector generates strings like the ones that you get when you print a collection, for example `[came, saw, conquered]`.”
* **“In summary, the essence of programming stream pipelines is side-effect-free function objects. This applies to all of the many function objects passed to streams and related objects.”**
* “The terminal operation `forEach` should only be used to report the result of a computation performed by a stream, not to perform the computation.”
* “In order to use streams properly, you have to know about collectors. The most important collector factories are `toList`, `toSet`, `toMap`, `groupingBy`, and `joining`.”


## Item 47: Prefer Collection to Stream as a return type

* “Many methods return sequences of elements. Prior to Java 8, the obvious return types for such methods were the collection interfaces `Collection`, `Set`, and `List`; `Iterable`; and the array types.”
  * “The norm was a collection interface.”
  * “If the method existed solely to enable for-each loops or the returned sequence couldn’t be made to implement some `Collection` method (typically, `contains(Object))`, the `Iterable` interface was used.”
  * “If the returned elements were primitive values or there were stringent performance requirements, arrays were used.”
  * “In Java 8, streams were added to the platform, substantially complicating the task of choosing the appropriate return type for a sequence-returning method.”
* “If an API returns only a stream and some users want to iterate over the returned sequence with a for-each loop, those users will be justifiably upset.”
  * “The only thing preventing programmers from using a for-each loop to iterate over a stream is `Stream`’s failure to extend `Iterable`.”


```java
// Hideous workaround to iterate over a stream
for  (ProcessHandle ph : (Iterable<ProcessHandle>)
                        ProcessHandle.allProcesses()::iterator)
```

* “This client code works, but it is too noisy and opaque to use in practice.”
* “A better workaround is to use an adapter method.”
  * “Note that no cast is necessary in the adapter method because Java’s type inference works properly in this context:”


```java
// Adapter from  Stream<E> to Iterable<E>
public static <E> Iterable<E> iterableOf(Stream<E> stream) {
    return stream::iterator;
}
```

* “Conversely, a programmer who wants to process a sequence using a stream pipeline will be justifiably upset by an API that provides only an Iterable.”

```java
// Adapter from Iterable<E> to Stream<E>
public static <E> Stream<E> streamOf(Iterable<E> iterable) {
    return StreamSupport.stream(iterable.spliterator(), false);
}
```

* “The `Collection` interface is a subtype of `Iterable` and has a stream method, so it provides for both iteration and stream access.”
* “Therefore, **`Collection` or an appropriate subtype is generally the best return type for a public, sequence-returning method**.”
* **“Arrays also provide for easy iteration and stream access with the `Arrays.asList` and `Stream.of` methods.”**
* “But **do not store a large sequence in memory just to return it as a collection**.”
  * “If the sequence you’re returning is large but can be represented concisely, consider implementing a special-purpose collection.”

```java
// Returns the power set of an input set as custom collection
public class PowerSet {
   public static final <E> Collection<Set<E>> of(Set<E> s) {
      List<E> src = new ArrayList<>(s);
      if (src.size() > 30)
         throw new IllegalArgumentException("Set too big " + s);
      return new AbstractList<Set<E>>() {
         @Override public int size() {
            return 1 << src.size(); // 2 to the power srcSize
         }

         @Override public boolean contains(Object o) {
            return o instanceof Set && src.containsAll((Set)o);
         }

         @Override public Set<E> get(int index) {
            Set<E> result = new HashSet<>();
            for (int i = 0; index != 0; i++, index >>= 1)
               if ((index & 1) == 1)
                  result.add(src.get(i));
            return result;
         }
      };
   }
}
```

* “This highlights a disadvantage of using `Collection` as a return type rather than `Stream` or `Iterable`: `Collection` has an int-returning size method, which limits the length of the returned sequence to `Integer.MAX_VALUE`, or $2^{31} − 1$. The `Collection` specification does allow the size method to return $2^{31} − 1$ if the collection is larger, even infinite, but this is not a wholly satisfying solution.”
* “In order to write a `Collection` implementation atop `AbstractCollection`, you need implement only two methods beyond the one required for Iterable: `contains` and `size`.”
  * “Often it’s easy to write efficient implementations of these methods. If it isn’t feasible, perhaps because the contents of the sequence aren’t predetermined before iteration takes place, return a stream or iterable, whichever feels more natural.”
* “There are times when you’ll choose the return type based solely on ease of implementation.”

```java
// Returns a stream of all the sublists of its input list
public class SubLists {
   public static <E> Stream<List<E>> of(List<E> list) {
      return Stream.concat(Stream.of(Collections.emptyList()),
         prefixes(list).flatMap(SubLists::suffixes));
   }

   private static <E> Stream<List<E>> prefixes(List<E> list) {
      return IntStream.rangeClosed(1, list.size())
         .mapToObj(end -> list.subList(0, end));
   }

   private static <E> Stream<List<E>> suffixes(List<E> list) {
      return IntStream.range(0, list.size())
         .mapToObj(start -> list.subList(start, list.size()));
   }
}
```

* “Note that we generate the prefixes and suffixes by mapping a stream of consecutive `int` values returned by `IntStream.range` and `IntStream.rangeClosed`. This idiom is, roughly speaking, the stream equivalent of the standard `for`-loop on integer indices.”
* “Thus, our sublist implementation is similar in spirit to the obvious nested `for`-loop:”


```java
for (int start = 0; start < src.size(); start++)
    for (int end = start + 1; end <= src.size(); end++)
        System.out.println(src.subList(start, end));
```

* “It is possible to translate this `for`-loop directly into a stream. The result is more concise than our previous implementation, but perhaps a bit less readable.”

```java
// Returns a stream of all the sublists of its input list
public static <E> Stream<List<E>> of(List<E> list) {
   return IntStream.range(0, list.size())
      .mapToObj(start ->
         IntStream.rangeClosed(start + 1, list.size())
            .mapToObj(end -> list.subList(start, end)))
      .flatMap(x -> x);
}
```

* “In summary, when writing a method that returns a sequence of elements, remember that some of your users may want to process them as a stream while others may want to iterate over them. Try to accommodate both groups.”
  * **“If it’s feasible to return a collection, do so.”**
  * “If you already have the elements in a collection or the number of elements in the sequence is small enough to justify creating a new one, return a standard collection such as `ArrayList`.”
  * “Otherwise, consider implementing a custom collection as we did for the power set.”
  * “If it isn’t feasible to return a collection, return a stream or iterable, whichever seems more natural.”
  * “If, in a future Java release, the `Stream` interface declaration is modified to extend `Iterable`, then you should feel free to return streams because they will allow for both stream processing and iteration.”


## Item 48: Use caution when making streams parallel

* “Writing concurrent programs in Java keeps getting easier, but writing concurrent programs that are correct and fast is as difficult as it ever was. Safety and liveness violations are a fact of life in concurrent programming, and parallel stream pipelines are no exception.”
* “Even under the best of circumstances, **parallelizing a pipeline is unlikely to increase its performance if the source is from `Stream.iterate`, or the intermediate operation limit is used**.”
* **“Do not parallelize stream pipelines indiscriminately.”**
  * “The performance consequences may be disastrous.”
* **“As a rule, performance gains from parallelism are best on streams over `ArrayList`, `HashMap`, `HashSet`, and `ConcurrentHashMap` instances; arrays; `int` ranges; and `long` ranges.”**
  * “What these data structures have in common is that they can all be accurately and cheaply split into subranges of any desired sizes, which makes it easy to divide work among parallel threads.”
  * “The abstraction used by the streams library to perform this task is the *spliterator*, which is returned by the `spliterator` method on `Stream` and `Iterable`.”
  * “Another important factor that all of these data structures have in common is that they provide good-to-excellent *locality of reference* when processed sequentially: sequential element references are stored together in memory.”
  * “If you write your own `Stream`, `Iterable`, or `Collection` implementation and you want decent parallel performance, you must override the `spliterator` method and test the parallel performance of the resulting streams extensively.”

* **“The nature of a stream pipeline’s terminal operation also affects the effectiveness of parallel execution.”**
  * “The best terminal operations for parallelism are reductions, where all of the elements emerging from the pipeline are combined using one of `Stream`’s `reduce` methods, or prepackaged reductions such as `min`, `max`, `count`, and `sum`.”
  * “The *short-circuiting* operations `anyMatch`, `allMatch`, and `noneMatch` are also amenable to parallelism.”
  * “The operations performed by `Stream`’s collect method, which are known as *mutable reductions*, are not good candidates for parallelism because the overhead of combining collections is costly.”
* “**Not only can parallelizing a stream lead to poor performance, including liveness failures; it can lead to incorrect results and unpredictable behavior** (*safety failures*).”
  * “Safety failures may result from parallelizing a pipeline that uses mappers, filters, and other programmer-supplied function objects that fail to adhere to their specifications.”
* “Even assuming that you’re using an efficiently splittable source stream, a parallelizable or cheap terminal operation, and non-interfering function objects, you won’t get a good speedup from parallelization unless the pipeline is doing enough real work to offset the costs associated with parallelism.”
  * “As a very rough estimate, the number of elements in the stream times the number of lines of code executed per element should be at least a hundred thousand [Lea14].”
* “It’s important to remember that parallelizing a stream is strictly a performance optimization. As is the case for any optimization, you must test the performance before and after the change to ensure that it is worth doing (Item 67). Ideally, you should perform the test in a realistic system setting.”
* **“Under the right circumstances, it is possible to achieve near-linear speedup in the number of processor cores simply by adding a `parallel` call to a stream pipeline.”**

```java
// Prime-counting stream pipeline - benefits from parallelization
static long pi(long n) {
    return LongStream.rangeClosed(2, n)
        .mapToObj(BigInteger::valueOf)
        .filter(i -> i.isProbablePrime(50))
        .count();
}

// Prime-counting stream pipeline - parallel version
static long pi(long n) {
    return LongStream.rangeClosed(2, n)
        .parallel()
        .mapToObj(BigInteger::valueOf)
        .filter(i -> i.isProbablePrime(50))
        .count();
}
```

* “If you are going to parallelize a stream of random numbers, start with a `SplittableRandom` instance rather than a `ThreadLocalRandom` (or the essentially obsolete `Random`).”
  * “`SplittableRandom` is designed for precisely this use, and has the potential for linear speedup.”
  * “`ThreadLocalRandom` is designed for use by a single thread, and will adapt itself to function as a parallel stream source, but won’t be as fast as `SplittableRandom`.”
  * “`Random` synchronizes on every operation, so it will result in excessive, parallelism-killing contention.”
* **“In summary, do not even attempt to parallelize a stream pipeline unless you have good reason to believe that it will preserve the correctness of the computation and increase its speed. The cost of inappropriately parallelizing a stream can be a program failure or performance disaster. If you believe that parallelism may be justified, ensure that your code remains correct when run in parallel, and do careful performance measurements under realistic conditions. If your code remains correct and these experiments bear out your suspicion of increased performance, then and only then parallelize the stream in production code.”**
