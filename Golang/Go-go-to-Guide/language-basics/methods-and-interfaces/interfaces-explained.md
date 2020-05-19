# Type, Value, and Equality of Interfaces

## Interface Type

* An interface type consists of **a set of method signatures**
* A variable of interface type can hold **any** value that implements these methods

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

* Actually, `*Temp` also implements `MyStringer`, since the method set of a pointer type `*T` is the set of all methods with receiver `*T` or `T`
* When you call a method on an interface value, the method of its underlying type is executed

```go
var x MyStringer

x = Temp(24)
fmt.Println(x.String()) // 24 °C

x = &Point{1, 2}
fmt.Println(x.String()) // (1,2)
```

## Structural Typing

* A type implements an interface by **implementing its methods**
* **No explicit declaration is required**
* In fact, the `Temp`, `*Temp` and `*Point` types also implement the standard library `fmt.Stringer` interface
* The `String` method in this interface is used to print values passed as an operand to functions such as `fmt.Println`

```go
var x MyStringer

x = Temp(24)
fmt.Println(x) // 24 °C

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

## The Empty Interface

* The interface type that specifies no methods is known as the empty interface
* `interface{}`
* An empty interface can hold values of any type since every type implements at least zero methods

```go
var x interface{}

x = 2.4
fmt.Println(x) // 2.4

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

## Interface Values

* An **interface value** consists of a **concrete value** and a **dynamic type**: `[Value, Type]`

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
* You can use type assertions or type switches to access the dynamic type of an interface value

## Equality

* Two interface values are equal
  * if they have equal concrete values and identical dynamic types,
  * **or** if both are nil
* A value `t` of interface type `T` and a value `x` of non-interface type `X` are equal if
  * `t`’s concrete value is equal to `x`
  * **and** `t`’s dynamic type is identical to `X`

```go
var x MyStringer
fmt.Println(x == nil) // true

x = (*Point)(nil)
fmt.Println(x == nil) // false
```
