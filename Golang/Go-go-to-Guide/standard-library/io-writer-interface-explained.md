# How to Use the `io.Writer` Interface

## Basics

* The `io.Writer` interface represents an entity to which you can write a stream of bytes

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

* `Write` writes up to `len(p)` bytes from `p` to the underlying data stream – it returns the number of bytes written and any error encountered that caused the write to stop early

## How to Use a Built-in Writer

* You can write directly into a `bytes.Buffer` using the `fmt.Fprintf` function
  * `bytes.Buffer` has a `Write` method, and
  * `fmt.Fprintf` takes a `Writer` as its first argument

```go
var buf bytes.Buffer
fmt.Fprintf(&buf, "Size: %d MB.", 85)
s := buf.String()) // s == "Size: 85 MB."
```

* Similarly, you can write directly into files or other streams, such as http connections
* You can compute the hash value of a file by copying the file into the io.Writer function of a suitable `hash.Hash` object

## Optimize String Writes

* Some Writers in the standard library have an additional `WriteString` method
* This method can be more efficient than the standard `Write` method since it writes a string directly without allocating a byte slice
* You can take direct advantage of this optimization by using the `io.WriteString()` function
  * If w implements a WriteString method, it is invoked directly
  * Otherwise, w.Write is called exactly once

```go
func WriteString(w Writer, s string) (n int, err error)
```
