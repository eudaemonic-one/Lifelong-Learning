# How to Best Clear a Slice: Empty vs. nil

## Remove All Elements

* To remove all elements, simply set the slice to `nil`
  * This will release the underlying array to the garbage collector (assuming there are no other references)

```go
a := []string{"A", "B", "C", "D", "E"}
a = nil
fmt.Println(a, len(a), cap(a)) // [] 0 0
```

## Keep Allocated Memory

```go
a := []string{"A", "B", "C", "D", "E"}
a = a[:0]
fmt.Println(a, len(a), cap(a)) // [] 0 5
fmt.Println(a[:2]) // [A B]
```

## Empty Slice vs. nil Slice

* In practice, **nil slices** and **empty slices** can often be treated in the same way
  * they have zero length and capacity,
  * they can be used with the same effect in for loops and append functions,
  * and they even look the same when printed

```go
var a []int = nil
fmt.Println(len(a)) // 0
fmt.Println(cap(a)) // 0
fmt.Println(a)      // []

var a []int = nil
var a0 []int = make([]int, 0)

fmt.Println(a == nil)  // true
fmt.Println(a0 == nil) // false

fmt.Printf("%#v\n", a)  // []int(nil)
fmt.Printf("%#v\n", a0) // []int{}
```

* **the nil slice is the preferred style**
  * Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects
  * When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.
