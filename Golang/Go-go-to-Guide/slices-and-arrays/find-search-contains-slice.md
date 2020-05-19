# Find Element in Slice/Array with Linear or Binary Search

## Linear Search

```go
// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}
```

## Binary Search

* If the array is **sorted**, you can use a binary search instead
* This will be much more efficient, since binary search runs in worst-case logarithmic time, making $O(log n)$ comparisons, where n is the size of the slice
* There are the three custom binary search functions: `sort.SearchInts`, `sort.SearchStrings` or `sort.SearchFloat64s`
  * They all have the signature `func SearchType(a []Type, x Type) int`
  * and return
    * the smallest index i at which `x <= a[i]`
    * or `len(a)` if there is no such index
  * The slice must be sorted in **ascending order**

```go
a := []string{"A", "C", "C"}

fmt.Println(sort.SearchStrings(a, "A")) // 0
fmt.Println(sort.SearchStrings(a, "B")) // 1
fmt.Println(sort.SearchStrings(a, "C")) // 1
fmt.Println(sort.SearchStrings(a, "D")) // 3
```

### Generic Binary Search

* There is also a **generic binary search** function `sort.Search`
  * `func Search(n int, f func(int) bool) int`
  * It returns
    * the smallest index `i` at which `f(i)` is true,
    * or `n` if there is no such index
  * It requires that `f` is false for some (possibly empty) prefix of the input range and then true for the remainder

```go
a := []string{"A", "C", "C"}
x := "C"

i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
if i < len(a) && a[i] == x {
    fmt.Printf("Found %s at index %d in %v.\n", x, i, a)
} else {
    fmt.Printf("Did not find %s in %v.\n", x, a)
}
// Output: Found C at index 1 in [A C C].
```

## The Map Option

* If you are doing repeated searches and updates, you may want to use a map instead of a slice
* A map provides lookup, insert, and delete operations in $O(1)$ expected amortized time
