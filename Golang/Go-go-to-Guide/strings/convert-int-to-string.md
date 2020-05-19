# Convert Between `int`, `int64`, and `string`

## `int`/`int64` to `string`

```go
s := strconv.Itoa(97) // s == "97"
s := string(97) // s == "a"

var n int64 = 97
s := strconv.FormatInt(n, 10) // s == "97" (decimal)

var n int64 = 97
s := strconv.FormatInt(n, 16) // s == "61" (hexadecimal)
```

## `string` to `int`/`int64`

```go
s := "97"
if n, err := strconv.Atoi(s); err == nil {
    fmt.Println(n+1)
} else {
    fmt.Println(s, "is not an integer.")
}
// Output: 98

s := "97"
n, err := strconv.ParseInt(s, 10, 64)
if err == nil {
    fmt.Printf("%d of type %T", n, n)
}
// Output: 97 of type int64
```

## `int` to `int64` (and Back)

```go
var n int = 97
m := int64(n) // safe

var m int64 = 2 << 32
n := int(m)    // truncated on machines with 32-bit ints
fmt.Println(n) // either 0 or 4,294,967,296
```

## General Formatting (Width, Indent, Sign)

```go
s := fmt.Sprintf("%+8d", 97)
// s == "     +97" (width 8, right justify, always show sign)
```
