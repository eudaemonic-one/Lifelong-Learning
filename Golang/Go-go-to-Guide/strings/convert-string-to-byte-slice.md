# Convert Between Byte Array/Slice and String

## Basics

* When you convert between a string and a byte slice (array), you get a brand new slice that contains the same bytes as the string, and vice versa
  * The conversion **doesn’t change** the data;
  * the only difference is that strings are **immutable**, while byte slices can be modified
  * If you need to manipulate the characters (runes) of a string, you may want to convert the string to a **rune slice** instead

## Convert String to Bytes

* When you convert a string to a byte slice, you get a new slice that contains the same bytes as the string
  * Note that the character € is encoded in UTF-8 using 3 bytes

```go
b := []byte("ABC€")
fmt.Println(b) // [65 66 67 226 130 172]
```

## Convert Bytes to String

* When you convert a slice of bytes to a string, you get a new string that contains the same bytes as the slice

```go
s := string([]byte{65, 66, 67, 226, 130, 172})
fmt.Println(s) // ABC€
```

## Performance

* These conversions create a new slice or string, and therefore have time complexity proportional to the number of bytes that are processed

### More Efficient Alternative

* In some cases, you might be able to use a string builder, which can concatenate strings without redundant copying: [Efficient string concatenation](./build-append-concatenate-strings-efficiently.md)
