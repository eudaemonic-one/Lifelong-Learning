# How to Reverse a String by Byte or Rune

## Byte by Byte

```go
// Reverse returns a string with the bytes of s in reverse order.
func Reverse(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}
```

## Rune by Rune

```go
// ReverseRune returns a string with the runes of s in reverse order.
// Invalid UTF-8 sequences, if any, will be reversed byte by byte.
func ReverseRune(s string) string {
    res := make([]byte, len(s))
    prevPos, resPos := 0, len(s)
    for pos := range s {
        resPos -= pos - prevPos
        copy(res[resPos:], s[prevPos:pos])
        prevPos = pos
    }
    copy(res[0:], s[prevPos:])
    return string(res)
}
```

### Example Usage

```go
for _, s := range []string{
	"Ångström",
	"Hello, 世界",
	"\xff\xfe\xfd", // invalid UTF-8
} {
	fmt.Printf("%q\n", ReverseRune(s))
}
// "mörtsgnÅ"
// "界世 ,olleH"
// "\xfd\xfe\xff"
```
