# 3 Ways to Trim Whitespace (or Other Characters) From a String

## Trim Space

* Use the `strings.TrimSpace` function to remove leading and trailing whitespace as defined by Unicode
  * To remove other leading and trailing characters, use `strings.Trim`
  * To remove only the leading or the trailing characters, use `strings.TrimLeft` or `strings.TrimRight`

```go
s := strings.TrimSpace("\t Goodbye hair!\n ")
fmt.Printf("%q", s) // "Goodbye hair!"
```
