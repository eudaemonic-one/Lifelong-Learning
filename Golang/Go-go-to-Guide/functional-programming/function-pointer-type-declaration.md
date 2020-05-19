# Function Types and Values

* Function types and function values can be used and passed around just like other values:
* The second call to `Map` uses a **function literal** (or **lambda**)

```go
type Operator func(x float64) float64

// Map applies op to each element of a.
func Map(op Operator, a []float64) []float64 {
    res := make([]float64, len(a))
    for i, x := range a {
        res[i] = op(x)
    }
    return res
}

func main() {
    op := math.Abs
    a := []float64{1, -2}
    b := Map(op, a)
    fmt.Println(b) // [1 2]

    c := Map(func(x float64) float64 { return 10 * x }, b)
    fmt.Println(c) // [10, 20]
}
```

## Details

* A function type describes the set of all functions with the same parameter and result types
  * The value of an uninitialized variable of function type is nil
  * The parameter names are optional
* The following two function types are identical

```go
func(x, y int) int
func(int, int) int
```
