# Anonymous Functions and Closures

* **A function literal (or lambda) is a function without a name**
* n this example a **function literal** is passed as the `less` argument to the `sort.Slice` function

```go
func Slice(slice interface{}, less func(i, j int) bool)

people := []string{"Alice", "Bob", "Dave"}
sort.Slice(people, func(i, j int) bool {
    return len(people[i]) < len(people[j])
})
fmt.Println(people)
// Output: [Bob Dave Alice]
```

* You can also use an intermediate variable
  * Note that the `less` function is a **closure**: it references the `people` variable, which is declared outside the function

```go
people := []string{"Alice", "Bob", "Dave"}
less := func(i, j int) bool {
    return len(people[i]) < len(people[j])
}
sort.Slice(people, less)
```

## Closures

* Function literals in Go are **closures**: they may refer to variables defined in an enclosing function. Such variables
  * are shared between the surrounding function and the function literal,
  * survive as long as they are accessible
* In this example, the function literal uses the local variable `n` from the enclosing scope to count the number of times it has been invoked

```go
// New returns a function Count.
// Count prints the number of times it has been invoked.
func New() (Count func()) {
    n := 0
    return func() {
        n++
        fmt.Println(n)
    }
}

func main() {
    f1, f2 := New(), New()
    f1() // 1
    f2() // 1 (different n)
    f1() // 2
    f2() // 2
}
```
