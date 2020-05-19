# Find the Type of an Object

## Use `fmt` for a String Type Description

* You can use the `%T` flag in the `fmt` package to get a Go-syntax representation of the type
  * The empty interface denoted by `interface{}` can hold values of any type

```go
var x interface{} = []int{1, 2, 3}
xType := fmt.Sprintf("%T", x)
fmt.Println(xType) // "[]int"
```

## A Type Switch Lets You Choose Between Types

```go
var x interface{} = 2.3
switch v := x.(type) {
case int:
    fmt.Println("int:", v)
case float64:
    fmt.Println("float64:", v)
default:
    fmt.Println("unknown")
}
// Output: float64: 2.3
```

## Reflection Gives Full Type Information

```go
var x interface{} = []int{1, 2, 3}
xType := reflect.TypeOf(x)
xValue := reflect.ValueOf(x)
fmt.Println(xType, xValue) // "[]int [1 2 3]"
```
