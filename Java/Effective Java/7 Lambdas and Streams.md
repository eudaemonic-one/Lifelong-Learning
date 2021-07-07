# Chapter 7. Lambdas and Streams

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