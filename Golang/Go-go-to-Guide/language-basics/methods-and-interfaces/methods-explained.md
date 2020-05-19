# Methods Explained

* **Go doesn't have classes, but you can define methods on types**
* You can define methods on any type declared in a type definition
  - A method is a function with an extra **receiver** argument
  - The receiver sits between the `func` keyword and the method name

```go
type House struct {
    garage bool
}

func (p *House) HasGarage() bool { return p.garage }

func main() {
    house := new(House)
    fmt.Println(house.HasGarage()) // Prints "false" (zero value)
}
```

## Conversions and Methods

* If you convert a value to a different type, the new value will have the methods of the new type, but not the old

```go
type MyInt int

func (m MyInt) Positive() bool { return m > 0 }

func main() {
    var m MyInt = 2
    m = m * m // The operators of the underlying type still apply.

    fmt.Println(m.Positive())        // Prints "true"
    fmt.Println(MyInt(3).Positive()) // Prints "true"

    var n int
    n = int(m) // The conversion is required.
    n = m      // ILLEGAL
}
// ../main.go:14:4: cannot use m (type MyInt) as type int in assignment
```

* It’s idiomatic in Go to convert the type of an expression to access a specific method

```go
var n int64 = 12345
fmt.Println(n)                // 12345
fmt.Println(time.Duration(n)) // 12.345µs
```
