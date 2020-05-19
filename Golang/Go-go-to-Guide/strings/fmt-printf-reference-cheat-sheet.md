# fmt.Printf Formatting Tutorial

## Basics

* With the Go `fmt` package you can format numbers and strings padded with spaces or zeroes, in different bases, and with optional quotes
* You submit a **template string** that contains the text you want to format plus some **annotation verbs** that tell the `fmt` functions how to format the trailing arguments

### `Printf`

* In this example, `fmt.Printf` formats and writes to standard output:
  * the **template string** is `"Binary: %b\\%b"`,
  * the **annotation verb** `%b` formats a number in binary, and
  * the **special value** `\\` is a backslash

```go
fmt.Printf("Binary: %b\\%b", 4, 5) // Prints `Binary: 100\101`
```

### `Sprintf` (Format Without Printing)

* Use fmt.Sprintf to format a string without printing it:

```go
s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
```

### Find `fmt` Errors with `vet`

```go
fmt.Printf("Binary: %b\\%b", 4) // An argument to Printf is missing.
// Output: Binary: 100\%!b(MISSING)
```

* To catch this type of errors early, you can use the `vet` command - it can find calls whose arguments do not align with the format string

```text
$ go vet example.go
example.go:8: missing argument for Printf("%b"): format reads arg 2, have only 1 args
```

## Cheat Sheet

### Default Formats and Type

* **Value:** `[]int64{0, 1}`

| Format        | Verb  | Description           |
| ------------- | ----- | --------------------- |
| [0 1]         | `%v`  | Default format        |
| []int64{0, 1} | `%#v` | Go-syntax format      |
| []int64       | `%T`  | The type of the value |

### Integer (Indent, Based, Sign)

* **Value:** `15`

| Format | Verb   | Description                                |
| ------ | ------ | ------------------------------------------ |
| 15     | `%d`   | Base 10                                    |
| +15    | `%+d`  | Always show sign                           |
| ␣␣15   | `%4d`  | Pad with spaces (width 4, right justified) |
| 15␣␣   | `%-4d` | Pad with spaces (width 4, left justified)  |
| 0015   | `%04d` | Pad with zeroes (width 4)                  |
| 1111   | `%b`   | Base 2                                     |
| 17     | `%o`   | Base 8                                     |
| f      | `%x`   | Base 16, lowercase                         |
| F      | `%X`   | Base 16, uppercase                         |
| 0xf    | `%#x`  | Base 16, with leading 0x                   |

### Character (Quoted, Unicode)

* **Value:** `65` (Unicode letter A)

| Format     | Verb  | Description            |
| ---------- | ----- | ---------------------- |
| A          | `%c`  | Character              |
| 'A'        | `%q`  | Quoted character       |
| U+0041     | `%U`  | Unicode                |
| U+0041 'A' | `%#U` | Unicode with character |

### Boolean (true/false)

* Use `%t` to format a boolean as `true` or `false`

### Pointer (hex)

* Use `%p` to format a pointer in base 16 notation with leading `0x`

### Float (Indent, Precision, Scientific Notation)

* **Value:** `123.456`

| Format       | Verb    | Description                               |
| ------------ | ------- | ----------------------------------------- |
| 1.234560e+02 | `%e`    | Scientific notation                       |
| 123.456000   | `%f`    | Decimal point, no exponent                |
| ␣␣123.46     | `%8.2f` | Width 8, precision 2                      |
| 123.456      | `%g`    | Exponent as needed, necessary digits only |

### String or Byte Slice (Quote, Indent, hex)

* **Value:** "café"

| Format         | Verb   | Description             |
| -------------- | ------ | ----------------------- |
| café           | `%s`   | Plain string            |
| ␣␣café         | `%6s`  | Width 6, right justify  |
| café␣␣         | `%-6s` | Width 6, left justify   |
| "café"         | `%q`   | Quoted string           |
| 636166c3a9     | `%x`   | Hex dump of byte values |
| 63 61 66 c3 a9 | `% x`  | Hex dump with spaces    |
