# Convert Between Rune Array/Slice and String

## Convert String to Runes

* When you convert a string to a rune slice, you get a new slice that contains the Unicode code points (runes) of the string
* For an invalid UTF-8 sequence, the rune value will be **0xFFFD** for each **invalid byte**
* You can also use a range loop to access the code points of a string

```go
r := []rune("ABC€")
fmt.Println(r)        // [65 66 67 8364]
fmt.Printf("%U\n", r) // [U+0041 U+0042 U+0043 U+20AC]
```

## Convert Runes to String

* When you convert a slice of runes to a string, you get a new string that is the concatenation of the runes converted to UTF-8 encoded strings
* Values outside the range of valid Unicode code points are converted to \uFFFD, the Unicode replacement character �

```go
s := string([]rune{'\u0041', '\u0042', '\u0043', '\u20AC', -1})
fmt.Println(s) // ABC€�
```
