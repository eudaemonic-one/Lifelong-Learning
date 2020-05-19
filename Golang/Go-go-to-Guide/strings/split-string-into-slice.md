# 3 Ways to Split a String into a Slice

## Split on Comma or Other Substring

* Use the `strings.Split` function to split a string into its comma separated values

```go
s := strings.Split("a,b,c", ",")
fmt.Println(s)
// Output: [a b c]
```

* To include the separators, use `strings.SplitAfter`
* To split only the first n values, use `strings.SplitN` and `strings.SplitAfterN`
* You can use `strings.TrimSpace` to strip leading and trailing whitespace from the resulting strings

## Split by Whitespace and Newline

* Use the `strings.Fields` function to split a string into substrings removing any space characters, including newlines

```go
s := strings.Fields(" a \t b \n")
fmt.Println(s)
// Output: [a b]
```

## Split on Regular Expression

```go
a := regexp.MustCompile(`a`)              // a single `a`
fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

zp := regexp.MustCompile(` *, *`)             // spaces and one comma
fmt.Printf("%q\n", zp.Split("a,b ,  c ", -1)) // ["a" "b" "c "]
```
