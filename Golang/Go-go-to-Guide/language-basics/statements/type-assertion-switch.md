# Type Assertions and Type Switches

## Type Assertions

* A **type assertion** doesn’t really convert an interface to another data type, but it provides access to an interface’s concrete value, which is typically what you want
* The type assertion `x.(T)` asserts that the concrete value stored in `x` is of type `T`, and that `x` is not nil
  * If `T` is not an interface, it asserts that the dynamic type of `x` is identical to `T`
  * If `T` is an interface, it asserts that the dynamic type of `x` implements `T`

```go
var x interface{} = "foo"

var s string = x.(string)
fmt.Println(s)     // "foo"

s, ok := x.(string)
fmt.Println(s, ok) // "foo true"

n, ok := x.(int)
fmt.Println(n, ok) // "0 false"

n = x.(int)        // ILLEGAL
// panic: interface conversion: interface {} is string, not int
```

## Type Switches

* A **type switch** performs several type assertions in series and runs the first case with a matching type

```go
var x interface{} = "foo"

switch v := x.(type) {
case nil:
    fmt.Println("x is nil")            // here v has type interface{}
case int: 
    fmt.Println("x is", v)             // here v has type int
case bool, string:
    fmt.Println("x is bool or string") // here v has type interface{}
default:
    fmt.Println("type unknown")        // here v has type interface{}
}
// x is bool or string
```
