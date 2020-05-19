# The 3 Ways to Sort in Go

## Sort a Slice of `int`s, `float64`s, or `string`s

* Use one of the functions
  * `sort.Ints`
  * `sort.Float64s`
  * `sort.Strings`
* Package `radix` contains a drop-in replacement for sort.Strings, which can be more than twice as fast in some settings

```go
s := []int{4, 2, 3, 1}
sort.Ints(s)
fmt.Println(s) // [1 2 3 4]
```

## Sort with Custom Comparator

* Use the function `sort.Slice`. It sorts a slice using a provided function `less(i, j int) bool`
  To sort the slice while keeping the original order of equal elements, use `sort.SliceStable` instead

```go
family := []struct {
    Name string
    Age  int
}{
    {"Alice", 23},
    {"David", 2},
    {"Eve", 2},
    {"Bob", 25},
}

// Sort by age, keeping original order or equal elements.
sort.SliceStable(family, func(i, j int) bool {
    return family[i].Age < family[j].Age
})
fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
```

### Sort Custom Data Structures

* Use the generic `sort.Sort` and `sort.Stable` functions
* They sort any collection that implements the `sort.Interface` interface

```go
type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
```

```go
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    family := []Person{
        {"Alice", 23},
        {"Eve", 2},
        {"Bob", 25},
    }
    sort.Sort(ByAge(family))
    fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}
```

## Bonus: Sort a Map by Key and Value

* A map is an **unordered** collection of key-value pairs
* If you need a stable iteration order, you must maintain a separate data structure

```go
m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
// Output:
// Alice 2
// Bob 3
// Cecil 1
```

## Performance and Implementation

* All algorithms in the Go sort package make $O(n log n)$ comparisons in the worst case, where $n$ is the number of elements to be sorted
