# Pointers Explained

* **A pointer is a vari­able that con­tains the address of an object**

## Basics

* Structs and arrays are **copied** when used in assignments and passed as arguments to functions
* With pointers this can be avoided
* Pointers store **addresses** of objects
* The addresses can be passed around more efficiently than the actual objects
* A pointer has type `*T`

```go
type Student struct {
    Name string
}

var ps *Student = new(Student) // ps holds the address of the new struct

ps := new(Student)
```

## Address Operator

* The `&` operator returns the address of an object

```go
s := Student{"Alice"} // s holds the actual struct
ps := &s              // ps holds the address of the struct
```

* The `&` operator can also be used with **composite literals**

```go
ps := &Student{"Alice"}
```

## Pointer Indirection

* For a pointer `x`, the **pointer indirection** `*x` denotes the value which `x` points to
* Pointer indirection is rarely used, since Go can automatically take the address of a variable

```go
ps := new(Student)
ps.Name = "Alice" // same as (*ps).Name = "Alice"
```

## Pointers as Parameters

* When using a pointer to modify an object, you’re affecting all code that uses the object

```go
// Bob is a function that has no effect.
func Bob(s Student) {
    s.Name = "Bob" // changes only the local copy
}

// Charlie sets pp.Name to "Charlie".
func Charlie(ps *Student) {
    ps.Name = "Charlie"
}

func main() {
    s := Student{"Alice"}

    Bob(s)
    fmt.Println(s) // prints {Alice}

    Charlie(&s)
    fmt.Println(s) // prints {Charlie}
}
```
