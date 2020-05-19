# How to Best Implement an Iterator

* Go has a built-in range loop for iterating over slices, arrays, strings, maps and channels
* To iterate over other types of data, an iterator function with callbacks is a clean and fairly efficient abstraction

### Basic Iterator Pattern

```go
// Iterate calls the f function with n = 1, 2, and 3.
func Iterate(f func(n int)) {
    for i := 1; i <= 3; i++ {
        f(i)
    }
}
```

* In use:

```go
Iterate(func(n int) { fmt.Println(n) })
// 1
// 2
// 3
```

## Iterator with Break

```go
// Iterate calls the f function with n = 1, 2, and 3.
// If f returns true, Iterate returns immediately
// skipping any remaining values.
func Iterate(f func(n int) (skip bool)) {
    for i := 1; i <= 3; i++ {
        if f(i) {
            return
        }
    }
}
```

* In use:

```go
Iterate(func(n int) (skip bool) {
	fmt.Println(n)
	return n == 2
})

// 1
// 2
```
