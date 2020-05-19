# Create, Initialize and Compare Structs

## Struct Types

* A struct is a typed collection of fields, useful for grouping data into records
  * To define a new **struct type**, you list the names and types of each field
  * The default **zero value** of a struct has all its fields zeroed
  * You can access individual fields with **dot notation**

```go
type Student struct {
    Name string
    Age  int
}

var a Student    // a == Student{"", 0}
a.Name = "Alice" // a == Student{"Alice", 0}
```

## 2 Ways to Create and Initialize a New Struct

* The `new` keyword can be used to create a new struct
* It returns a pointer to the newly created struct

```go
var pa *Student   // pa == nil
pa = new(Student) // pa == &Student{"", 0}
pa.Name = "Alice" // pa == &Student{"Alice", 0}
```

* You can also create and initialize a struct with a **struct literal**
  * An element list that contains keys does not need to have an element for each struct field
    * Omitted fields get the zero value for that field
  * An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared
  * A literal may omit the element list; such a literal evaluates to the zero value for its type

```go
b := Student{ // b == Student{"Bob", 0}
    Name: "Bob",
}
    
pb := &Student{ // pb == &Student{"Bob", 8}
    Name: "Bob",
    Age:  8,
}

c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
d := Student{}             // d == Student{"", 0}
```

## Compare Structs

* You can compare struct values with the comparison operators == and !=
* Two values are equal if their corresponding fields are equal

```go
d1 := Student{"David", 1}
d2 := Student{"David", 2}
fmt.Println(d1 == d2) // false
```
