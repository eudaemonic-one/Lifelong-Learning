# Last Item in a Slice/Array

## Read Last Element

```go
a := []string{"A", "B", "C"}
s := a[len(a)-1] // C
```

* Go doesn't have negative indexing like Python does
  * This is a deliberate design decision — keeping the language simple can help save you from subtle bugs

## Remove Last Element

```go
a = a[:len(a)-1] // [A B]
```

### Watch Out for Memory Leaks

* If the slice is permanent and the element temporary, you may want to remove the reference to the element before slicing it off

```go
a[len(a)-1] = "" // Erase element (write zero value)
a = a[:len(a)-1] // [A B]
```
