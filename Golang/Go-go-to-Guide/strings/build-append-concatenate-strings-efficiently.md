# Efficient String Concatenation

## Clean and Simple String Building (fmt)

* For simple cases where performance is a non-issue, [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend. It’s clean, simple and fairly efficient

```go
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

## High-performance String Concatenation

* A [`strings.Builder`](https://golang.org/pkg/strings/#Builder) is used to efficiently append strings using write methods
  * It offers a subset of the [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) methods that allows it to safely avoid extra copying when converting a builder to a string
  * You can use the [`fmt`](https://golang.org/pkg/fmt/) package for formatting since the builder implements the [`io.Writer`](https://yourbasic.org/golang/io-writer-interface-explained/) interface

```go
var b strings.Builder
b.Grow(32)
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&b, "%d:%d, ", i+1, p)
}
s := b.String()   // no copying
s = s[:b.Len()-2] // no copying (removes trailing ", ")
fmt.Println(s)
// Output: 1:2, 2:3, 3:5, 4:7, 5:11, 6:13
```
