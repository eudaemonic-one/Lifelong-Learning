# Type, Value, and Equality of Interfaces

## Interface Type

* **An interface type consists of a set of method signatures. A variable of interface type can hold any value that implements these methods**
* In this example both Temp and `*Point` implement the `MyStringer` interface
* Actually, `*Temp` also implements `MyStringer`, since the method set of a pointer type `*T` is the set of all methods with receiver `*T` or `T`

```go
type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
```

## Structural Typing

* **A type implements an interface by implementing its methods and no explicit declaration is required**
* In fact, the Temp, `*Temp` and `*Point` types also implement the standard library `fmt.Stringer interface`. The `String` method in this interface is used to print values passed as an operand to functions such as `fmt.Println`

```go
var x MyStringer

x = Temp(24)
fmt.Println(x) // 24 °C

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

## The Empty Interface

* The interface type that specifies no methods is known as the empty interface.
* An empty interface can hold values of any type since every type implements at least zero methods
* The `fmt.Println` function is a chief example that takes any number of arguments of any type

```go
func Println(a ...interface{}) (n int, err error)
```

## Interface Values

* An **interface value** consists of a **concrete value** and a **dynamic type**: `[Value, Type]`
* In a call to `fmt.Printf`, you can use `%v` to print the concrete value and `%T` to print the dynamic type

```go
var x MyStringer
fmt.Printf("%v %T\n", x, x) // <nil> <nil>

x = Temp(24)
fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

x = &Point{1, 2}
fmt.Printf("%v %T\n", x, x) // (1,2) *main.Point

x = (*Point)(nil)
fmt.Printf("%v %T\n", x, x) // <nil> *main.Point
```

* The **zero value** of an interface type is nil, which is represented as `[nil, nil]`
* Calling a method on a nil interface is a run-time error
  * However, it’s quite common to write methods that can handle a receiver value `[nil, Type]`, where `Type` isn’t nil
* You can use **type assertions** or **type switches** to access the dynamic type of an interface value

### Type Assertions

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

### Type Switches

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

## Equality

* Two interface values are equal
  * if they have equal concrete values **and** identical dynamic types,
  * or if both are nil

* A value `t` of interface type `T` and a value `x` of non-interface type `X` are equal if
  * `t`’s concrete value is equal to `x`
  * **and** `t`’s dynamic type is identical to `X`

```go
var x MyStringer
fmt.Println(x == nil) // true

x = (*Point)(nil)
fmt.Println(x == nil) // false
```

### Nil is not `nil`

```go
func Foo() error {
    var err *os.PathError = nil
    // …
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)        // <nil>
    fmt.Println(err == nil) // false
}
```

* An interface value is equal to `nil` only if both its value and dynamic type are `nil`. In the example above, `Foo()` returns `[nil, *os.PathError]` and we compare it with `[nil, nil]`
* To avoid this problem use a variable of type `error` instead, for example a named return value
* **Best practice:** Use the built-in `error` interface type, rather than a concrete type, to store and return error values

```go
func Foo() (err error) {
    // …
    return // err is unassigned and has zero value [nil, nil]
}

func main() {
    err := Foo()
    fmt.Println(err)        // <nil>
    fmt.Println(err == nil) // true
}
```
