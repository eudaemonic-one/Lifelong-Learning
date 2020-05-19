# Conversions [Complete List]

## Basics

* The expression `T(x)` converts the value `x` to the type `T`

```go
x := 5.1
n := int(x) // convert float to int
```

* The conversion rules are extensive but predictable:
  * all conversions between typed expressions must be explicitly stated,
  * illegal conversions are caught by the compiler
* Conversions to and from numbers and strings may **change the representation** and have a **run-time cost**
* All other conversions only change the type but not the representation of `x`

## Interfaces

* To “convert” an interface to a string, struct or map you should use a **type assertion** or a **type switch**
* A type assertion doesn’t really convert an interface to another data type, but it provides access to an interface’s concrete value, which is typically what you want

## Integers

* When converting to a shorter integer type, the value is **truncated** to fit in the result type’s size
* When converting to a longer integer type,
  - if the value is a signed integer, it is **signed extended**;
  - otherwise it is **zero extended**

```go
a := uint16(0x10fe) // 0001 0000 1111 1110
b := int8(a)        //           1111 1110 (truncated to -2)
c := uint16(b)      // 1111 1111 1111 1110 (sign extended to 0xfffe)
```

## Floats

* When converting a floating-point number to an integer, the **fraction is discarded** (truncation towards zero)
* When converting an integer or floating-point number to a floating-point type, the result value is **rounded to the precision** specified by the destination type

```go
var x float64 = 1.9
n := int64(x) // 1
n = int64(-x) // -1

n = 1234567890
```

* **Warning:** In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent

## Integer to String

* When converting an integer to a string, the value is interpreted as a Unicode code point, and the resulting string will contain the character represented by that code point, encoded in UTF-8
* If the value does not represent a valid code point (for instance if it’s negative), the result will be `"\ufffd"`, the Unicode replacement character �

```go
string(97) // "a"
string(-1) // "\ufffd" == "\xef\xbf\xbd"
```

* Use `strconv.Itoa` to get the decimal string representation of an integer

```go
strconv.Itoa(97) // "97"
```

## Strings and Byte Slices

* Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice
* Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string

```go
string([]byte{97, 230, 151, 165}) // "a日"
[]byte("a日")                     // []byte{97, 230, 151, 165}
```

## Strings and Rune Slices

* Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings
* Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string

```go
string([]rune{97, 26085}) // "a日"
[]rune("a日")             // []rune{97, 26085}
```

## Underlying Type

* A non-constant value can be converted to type `T` if it has the same underlying type as `T`

```go
type (
	T1 int64
	T2 T1
)

var n int64 = 12345
fmt.Println(n)                // 12345
fmt.Println(time.Duration(n)) // 12.345µs
```

## Implicit Conversions

* The only implicit conversion in Go is when an untyped constant is used in a situation where a type is required

```go
var x float64
x = 1 // Same as x = float64(1)

t := 2 * time.Second // Same as t := time.Duration(2) * time.Second
```

* The implicit conversions are necessary since there is no mixing of numeric types in Go
  * You can only multiply a `time.Duration` with another `time.Duration`
* When the type can’t be inferred from the context, an untyped constant is converted to a `bool`, `int`, `float64`, `complex128`, `string` or `rune` depending on the syntactical format of the constant

```go
n := 1   // Same as n := int(1)
x := 1.0 // Same as x := float64(1.0)
s := "A" // Same as s := string("A")
c := 'A' // Same as c := rune('A')
```

* Illegal implicit conversions are caught by the compiler

```go
var b byte = 256 // Same as var b byte = byte(256)
// ../main.go:2:6: constant 256 overflows byte
```

## Pointers

* The Go compiler does not allow conversions between pointers and integers
* Package `unsafe` implements this functionality under restricted circumstances
