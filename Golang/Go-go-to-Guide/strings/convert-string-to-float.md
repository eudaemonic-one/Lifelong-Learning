# Convert Between Float and String

## String to Float

* Use the `strconv.ParseFloat` function to parse a string as a floating-point number with the precision specified by bitSize: 32 for `float32`, or 64 for `float64`
  * When `bitSize` is 32, the result still has type `float64`, but it will be convertible to `float32` without changing its value

```go
func ParseFloat(s string, bitSize int) (float64, error)

f := "3.14159265"
if s, err := strconv.ParseFloat(f, 32); err == nil {
    fmt.Println(s) // 3.1415927410125732
}
if s, err := strconv.ParseFloat(f, 64); err == nil {
    fmt.Println(s) // 3.14159265
}
```

## Float to String

* Use the `fmt.Sprintf` method to format a floating-point number as a string

```go
s := fmt.Sprintf("%f", 123.456) // s == "123.456000"
```
