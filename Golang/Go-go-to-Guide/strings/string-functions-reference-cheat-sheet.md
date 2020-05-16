# Go String Handling Overview

## String Literals

| Expression                | Result     | Note                                       |
| ------------------------- | ---------- | ------------------------------------------ |
| `""`                      |            | Default zero value for type `string`       |
| `"Japan 日本"`            | Japan 日本 | Go code is Unicode text encoded in `UTF‑8` |
| `"\xe6\x97\xa5"`          | 日         | `\xNN` specifies a byte                    |
| `"\u65E5"`                | 日         | `\uNNNN` specifies a Unicode value         |
| `"\\"`                    | \          | Backslash                                  |
| `"\""`                    | "          | Double quote                               |
| `"\n"`                    |            | Newline                                    |
| `"\t"`                    |            | Tab                                        |
| ``\xe6``                  | \xe6       | Raw string literal*                        |
| `html.EscapeString("<>")` | &lt;&gt;   | HTML escape for <, >, &, ' and "           |
| `url.PathEscape("A B")`   | A%20B      | URL percent-encoding                       |

* See [Escapes and multiline strings](./multiline-string.md) for more on raw strings, escape characters and string encodings

## Concatenate

| Expression     | Result | Note          |
| -------------- | ------ | ------------- |
| `"Ja" + "pan"` | Japan  | Concatenation |

* See [3 tips for efficient string concatenation](./build-append-concatenate-strings-efficiently.md) for how to best use a string builder to concatenate strings without redundant copying

## Equal and Compare (Ignore Case)

| Expression                            | Result | Note                 |
| ------------------------------------- | ------ | -------------------- |
| `"Japan" == "Japan"`                  | true   | Equality             |
| `strings.EqualFold("Japan", "JAPAN")` | true   | Unicode case folding |
| `"Japan" < "japan"`                   | true   | Lexicographic order  |

## Length in Bytes or Runes

| Expression                     | Result | Note            |
| ------------------------------ | ------ | --------------- |
| `len("日")`                    | 3      | Length in bytes |
| `utf8.RuneCountInString("日")` | 1      | in runes        |
| `utf8.ValidString("日")`       | true   | UTF-8?          |

## Index, Substring, Iterate

| Expression     | Result | Note               |
| -------------- | ------ | ------------------ |
| `"Japan"[2]`   | 'p'    | Byte at position 2 |
| `"Japan"[1:3]` | ap     | Byte indexing      |

* A Go range loop iterates over UTF-8 encoded characters ([runes](./rune.md)):

```go
for i, ch := range "Japan 日本" {
    fmt.Printf("%d:%q ", i, ch)
}
// Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'
```

* **Iterating over bytes produces nonsense characters for non-ASCII text**

```go
s := "Japan 日本"
for i := 0; i < len(s); i++ {
    fmt.Printf("%q ", s[i])
}
// Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'
```

## Search (Contains, Prefix/suffix, Index)

| Expression                             | Result | Note                             |
| -------------------------------------- | ------ | -------------------------------- |
| `strings.ContainsAny("Japan", "abc")`  | true   | Is a, b or c in Japan?           |
| `strings.Count("Banana", "ana")`       | 1      | Non-overlapping instances of ana |
| `strings.HasPrefix("Japan", "Ja")`     | true   | Does Japan start with Ja?        |
| `strings.HasSuffix("Japan", "pan")`      |        |                                  |
| `strings.Index("Japan", "abc")`        | -1     | Index of first abc               |
| `strings.IndexAny("Japan", "abc")`     | 1      | a, b or c                        |
| `strings.LastIndex("Japan", "abc")`    | -1     | Index of last abc                |
| `strings.LastIndexAny("Japan", "abc")` | 3      | a, b or c                        |

## Replace (Uppercase/lowercase, Trim)

| Expression                                                   | Result | Note                                                 |
| ------------------------------------------------------------ | ------ | ---------------------------------------------------- |
| `strings.Replace("foo", "o", ".", 2)`                        | f..    | Replace first two “o” with “.” Use -1 to replace all |
| `f := func(r rune) rune { return r + 1} strings.Map(f, "ab")` | bc     | Apply function to each character                     |
| `strings.ToUpper("Japan")`                                   | JAPAN  | Uppercase                                            |
| `strings.ToLower("Japan")`                                   | japan  | Lowercase                                            |
| `strings.Title("ja pan")`                                    | Ja Pan | Initial letters to uppercase                         |
| `strings.TrimSpace(" foo\n")`                                | foo    | Strip leading and trailing white space               |
| `strings.Trim("foo", "fo")`                                  |        | Strip *leading and trailing* f:s and o:s             |
| `strings.TrimLeft("foo", "f")`                               | oo     | *only leading*                                       |
| `strings.TrimRight("foo", "o")`                              | f      | *only trailing*                                      |
| `strings.TrimPrefix("foo", "fo")`                            | o      |                                                      |
| `strings.TrimSuffix("foo", "o")`                             | fo     |                                                      |

## Split by Space or Comma

| Expression                       | Result       | Note               |
| -------------------------------- | ------------ | ------------------ |
| `strings.Fields(" a\t b\n")`     | `["a" "b"]`  | Remove white space |
| `strings.Split("a,b", ",")`      | `["a" "b"]`  | Remove separator   |
| `strings.SplitAfter("a,b", ",")` | `["a," "b"]` | Keep separator     |

## Join Strings with Separator

| Expression                              | Result | Note             |
| --------------------------------------- | ------ | ---------------- |
| `strings.Join([]string{"a", "b"}, ":")` | a:b    | Add separator    |
| `strings.Repeat("da", 2)`               | dada   | 2 copies of “da” |

## Format and Convert

| Expression                   | Result  | Note          |
| ---------------------------- | ------- | ------------- |
| `strconv.Itoa(-42)`          | `"-42"` | Int to string |
| `strconv.FormatInt(255, 16)` | `"ff"`  | Base 16       |
