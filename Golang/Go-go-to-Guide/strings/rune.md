# Runes and Character Encoding

## Characters, ASCII and Unicode

* The `rune` type is an alias for `int32`, and is used to emphasize than an integer represents a code point
* **ASCII** defines 128 characters, identified by the **code points** 0–127
  * It covers English letters, Latin numbers, and a few other characters
* **Unicode**, which is a superset of ASCII, defines a codespace of 1,114,112 code points
  * Unicode version 10.0 covers 139 modern and historic scripts (including the runic alphabet, but not Klingon) as well as multiple symbol sets

## Strings and UTF-8 Encoding

* A `string` is a sequence of bytes, not runes
* However, strings often contain Unicode text encoded in UTF-8, which encodes all Unicode code points using one to four bytes
* ASCII characters are encoded with one byte, while other code points use more

```go
fmt.Println([]byte("café")) // [99 97 102 195 169]
fmt.Println([]rune("café")) // [99 97 102 233]
```
