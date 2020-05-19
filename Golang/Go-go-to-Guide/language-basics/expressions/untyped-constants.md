# Untyped Numeric Constants with No Limits

* Constants may be **typed** or **untyped**

```go
const a uint = 17
const b = 55
```

* An untyped constant has **no limits**
* The inferred type is determined by the syntax of the value:
  - `123` gets type `int`, and
  - `123.4` becomes a `float64`

```go
const big = 10000000000  // Ok, even though it's too big for an int.
const bigger = big * 100 // Still ok.
var i int = big / 100    // No problem: the new result fits in an int.

// Compile time error: "constant 10000000000 overflows int"
var j int = big
```

## Enumerations

* Go does not have enumerated types
* Instead, you can use the special name `iota` in a single `const` declaration to get a series of increasing values
* When an initialization expression is omitted for a `const`, it reuses the preceding expression

```go
const (
    red = iota // red == 0
    blue       // blue == 1
    green      // green == 2
)
```
