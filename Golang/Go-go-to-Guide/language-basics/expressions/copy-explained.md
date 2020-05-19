# How to Use the Copy Function

## Copy

* The built-in copy function copies elements into a destination slice `dst` from a source slice `src`
* `func copy(dst, src []Type) int`
* It returns the number of elements copied, which will be the **minimum** of `len(dst)` and `len(src)`
* As a **special case**, it’s legal to copy bytes from a string to a slice of bytes
* `copy(dst []byte, src string) int`

## Examples

### Copy from One Slice to Another

```go
var s = make([]int, 3)
n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}
```

### Copy from a Slice to Itself

```go
s := []int{0, 1, 2}
n := copy(s, s[1:]) // n == 2, s == []int{1, 2, 2}
```

### Copy from a String to a Byte Slice

```go
var b = make([]byte, 5)
copy(b, "Hello, world!") // b == []byte("Hello")
```
