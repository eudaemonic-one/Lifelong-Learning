# Make Slices, Maps, and Channels

* Slices, maps and channels can be created with the built-in make function
* The memory is initialized with zero values

| Call            | Type    | Description                               |
| --------------- | ------- | ----------------------------------------- |
| `make(T, n)`    | slice   | slice of type T with length n             |
| `make(T, n, c)` |         | capacity c                                |
| `make(T)`       | map     | map of type T                             |
| `make(T, n)`    |         | initial room for approximately n elements |
| `make(T)`       | channel | unbuffered channel of type T              |
| `make(T, n)`    |         | buffered channel with buffer size n       |

```go
s := make([]int, 10, 100)      // slice with len(s) == 10, cap(s) == 100
m := make(map[string]int, 100) // map with initial room for ~100 elements
c := make(chan int, 10)        // channel with a buffer size of 10
```

* Slices, arrays and maps can also be created with composite literals

```go
s := []string{"f", "o", "o"} // slice with len(s) == 3, cap(s) == 3
a := [...]int{1, 2}          // array with len(a) == 2
m := map[string]float64{     // map with two key-value elements
    "e":  2.71828,
    "pi": 3.1416,
}
```
