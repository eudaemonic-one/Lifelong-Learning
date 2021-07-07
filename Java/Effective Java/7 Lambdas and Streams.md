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
