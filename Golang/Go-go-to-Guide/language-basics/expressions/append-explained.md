# How to Append Anything (Element, Slice, or String) to a Slice

## Append Function Basics

* With the built-in append function you can use a slice as a dynamic array
* The function appends any number of elements to the end of a slice:
  * if there is enough capacity, the underlying array is reused;
  * if not, a new underlying array is allocated and the data is copied over
* Append **returns the updated slice**
* Therefore you need to store the result of an append, often in the variable holding the slice itself:

```go
a := []int{1, 2}
a = append(a, 3, 4) // a == [1 2 3 4]
```

## Append One Slice to Another

* You can **concatenate two slices** using the three dots notation:

```go
a := []int{1, 2}
b := []int{11, 22}
a = append(a, b...) // a == [1 2 11 22]
```

* The `...` unpacks `b`
* The result does not depend on whether the **arguments overlap**

```go
a := []int{1, 2}
a = append(a, a...) // a == [1 2 1 2]
```

### Append String to Byte Slice

```go
slice := append([]byte("Hello "), "world!"...)
```
