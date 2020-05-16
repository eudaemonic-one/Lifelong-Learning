# Escapes and Multiline Strings

## Raw String Literals

* Raw string literals, delimited by **backticks** (back quotes), are interpreted literally
* They can contain line breaks, and backslashes have no special meaning

```go
const s = `First line
Second line`
fmt.Println(s)
// Output:
// First line
// Second line
```

### Backtick Escape

* It’s not possible to include a backtick in a raw string literal, but you can do

```go
fmt.Println("`" + "foo" + "`") // Output: `foo`
```

## Interpreted String Literals

* To insert escape characters, use interpreted string literals delimited by **double quotes**

```go
const s = "\tFirst line\n" +
"Second line"
fmt.Println(s)
// Output:
//     First line
// Second line
```

### Double Quote Escape

* Use `\"` to insert a double quote in an interpreted string literal:

```go
fmt.Println("\"foo\"") // "foo"
```

## Escape HTML

* Use `html.EscpapeString` to encode a string so that it can be safely placed inside HTML text. The function escapes the five characters <, >, &, ' and "
* `html.UnescapeString` does the inverse transformation

```go
const s = `"Foo's Bar" <foobar@example.com>`
fmt.Println(html.EscapeString(s))
// Output:
// &#34;Foo&#39;s Bar&#34; &lt;foobar@example.com&gt;
```

## Escape URL

* Use `url.PathEscape` in package `net/url` to encode a string so it can be safely placed inside a URL
  * The function uses percent-encoding
* `url.PathUnescape` does the inverse transformation

```go
const s = `Foo's Bar?`
fmt.Println(url.PathEscape(s))
// Output:
// Foo%27s%20Bar%3F
```

## All Escape Characters

* Arbitrary character values can be encoded with backslash escapes and used in string or rune literals
	* `\x` followed by exactly two hexadecimal digits,
	* `\` followed by exactly three octal digits,
	* `\u` followed by exactly four hexadecimal digits,
	* `\U` followed by exactly eight hexadecimal digits,
	* where the escapes `\u` and `\U` represent Unicode code points
