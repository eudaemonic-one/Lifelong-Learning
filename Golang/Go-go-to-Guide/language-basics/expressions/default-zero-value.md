# Default Zero Values for All Go Types

* Variables declared without an initial value are set to their zero values:
  - `0` for all **integer** types,
  - `0.0` for **floating point** numbers,
  - `false` for **booleans**,
  - `""` for **strings**,
  - `nil` for **interfaces**, **slices**, **channels**, **maps**, **pointers** and **functions**
* The elements of an **array** or **struct** will have its fields zeroed if no value is specified

```go
type T struct {
    n int
    f float64
    next *T
}
fmt.Println([2]T{}) // [{0 0 <nil>} {0 0 <nil>}]
```
