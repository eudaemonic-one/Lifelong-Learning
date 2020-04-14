# Lecture 22 Design Case Study: Java Functional APIs and streams

## Lambdas

* Comes from lambda-calculus
* A lambda is simply an anonomous function
  * without a corresponding identifier
* Function objects in Java 8 (2014)
  * `Arrays.sort(words, (s1, s2) -> s1.length() - s2.length())`
  * type inference

### Lambda Syntax

| Syntax                                 | Example                                                      |
| -------------------------------------- | ------------------------------------------------------------ |
| parameter -> expression                | `x -> x * x`                                                 |
| parameter -> block                     | `i -> { int res = 1; return result;}`                        |
| (parameters) -> expression             | `(x, y) -> x * y`                                            |
| (parameters) -> block                  | `(x, y) -> {int z = 1; return x + y + z;}`                   |
| (parameter declarations) -> expression | `(double x, double y) -> Math.sqrt(x\*x + y\*y)`             |
| (parameter declarations) -> block      | `(int i, int j) -> { int res = 1; for (int k = i; k < j; k++) res \*= k; return res;}` |

### Functional Interfaces

* **Java has no functional types, only functional interfaces**
  * **Interfaces with only one explicit abstract method**
    * Singkle Abstract Method (SAM interface)
  * Optionally annotated with `@FunctionalInterface`
    * Do it, for the same reason you use `@Override`
  * **A lambda is essentially a functional interface literal**
  * Some functional interfaces:
    * `Runnable`
    * `Callable`
    * `Comparator`
    * `ActionListener`
  * Many more in package `java.util.function`

### Subtle Difference Between Lambdas & Anonymous Classes

```java
class Enclosing {
  Supplier<Object> lambda() {
    return () -> this;
  }
  Supplier<Object> anon() {
    return new Supplier<Object>() {
      public Object get() {return this;}
    }
  }
}
```

### Method References

* **A more succinct alternative to lambdas**
* Lambdas are succint
  * `map.merge(key, 1, (count, incr) -> count + incr)`
* But method references can be more so
  * `map.merge(key, 1, Integer::sum)`
* The more parameters, the bigger the win
  * But parameter names may provide documentation
  * **If you use a lambda, choose parameter names carefully**

| Type              | Example                  | Lambda Equivalent*                                   |
| ----------------- | ------------------------ | ---------------------------------------------------- |
| Static            | `Integer::parseInt`      | `str -> Integer.parseInt(str)`                       |
| Bound             | `Instant.now()::isAfter` | `Instant then = Instant.now(); t -> then.isAfter(t)` |
| Unbound           | `String::toLowerCase`    | `str -> str.toLowerCase()`                           |
| Class Constructor | `TreeMap<K,V>::new`      | `() -> new TreeMap<K,V>()`                           |
| Array Constructor | `int[]::new`             | `len -> new int[len]`                                |

## Streams

* A bunch of data objects (typically from a collection, array, or input device) for bulk data processing
* Streams are processed lazily
  * Data is pulled by terminal operation, not pushed by source
  * Intermediate operations can be fused
  * Intermediate results typically not stored
    * But there are exceptions (e.g. `sorted`)

### Simple Stream Examples - Mapping, Filtering, Sorting

```java
List <String> longStrings = stringList.stream()
  .filter(s -> s.length() > 3)
  .collect(Collectors.toList());

List<String> firstLetters = stringList.stream()
  .map(s -> s.substring(0,1))
  .collect(Collectors.toList());

List<String> firstLettersOfLongStrings = stringList.stream()
  .filter(s -> s.length() > 3)
  .map(s -> s.substring(0,1))
  .collect(Collectors.toList());

List<String> sortedFirstLettersWithoutDups = stringList.stream()
  .map(s -> s.substring(0,1))
  .distinct()
  .sorted()
  .collect(Collectors.toList());

try (Stream<String> lines = Files.lines(Path.get(filename))) {
  lines.forEach(System.out::println);
}

try (Stream<String> lines = Files.lines(Path.get(filename))) {
  lines.map(String::trim)
    .filter(s -> !s.isEmpty())
    .sorted()
    .forEach(System.out::println);
}

boolean allStringHaveLengthThree = stringList.stream()
  .allMatch(s -> s.length() == 3);

boolean anyStringHaveLengthThree = stringList.stream()
  .anyMatch(s -> s.length() == 3);
```
