# Remove All Duplicate Whitespace

## Remove Duplicate Whitespace

```go
space := regexp.MustCompile(`\s+`)
s := space.ReplaceAllString("Hello  \t \n world!", " ")
fmt.Printf("%q", s) // "Hello world!"
```

* `\s+` is a regular expression:
  * the character class `\s` matches a space, tab, new line, carriage return or form feed,
  * and `+` says “one or more of those”

### Trim Leading and Trailing Space

* To trim leading and trailing whitespace, use the `strings.TrimSpace` function
