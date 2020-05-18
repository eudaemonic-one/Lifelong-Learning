# Shuffle a Slice or Array

* The `rand.Shuffle` function in package `math/rand` shuffles an input sequence using a given swap function

```go
a := []int{1, 2, 3, 4, 5, 6, 7, 8}
rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

// [5 8 6 4 3 7 2 1]
```
