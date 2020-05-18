# Package Documentation

## `godoc.org` Website

* The [GoDoc](https://godoc.org/) website hosts docu­men­tation for all public Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad

## Local `godoc` Server

* The godoc command extracts and generates documentation for all locally installed Go programs, both your own code and the standard libraries
* The following command starts a web server that presents the documentation at `http://localhost:6060/`

```text
$ godoc -http=:6060 &
```

## `go doc` Command-line Tool

* The go doc command prints plain text documentation to standard output:

```text
$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```

## Create Documentation

* To document a function, type, constant, variable, or even a complete package, write a regular comment directly preceding its declaration, with no blank line in between

```text
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
…
```

## Runnable Documentation Examples

* You can add example code snippets to the package documentation; this code is verified by running it as a **test**

### Examples Are Tests

* Examples are compiled (and optionally executed) as part of a package's test suite
* As with typical tests, examples are functions that **reside in** a package's `_test.go` files
* Unlike normal test functions, though, example functions take no arguments and **begin with** the word `Example` instead of `Test`

```go
package stringutil_test

import (
    "fmt"

    "github.com/golang/example/stringutil"
)

func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
```

### Output Comments

* As it executes the example, the testing framework **captures data written to standard output** and then **compares the output against the example's "Output:" comment**
* The test passes if the test's output matches its output comment
* Examples without output comments are useful for demonstrating code that cannot run as unit tests, such as that which accesses the network, while guaranteeing the example at least compiles

```go
func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: golly
}
```

```text
$ go test
--- FAIL: ExampleReverse (0.00s)
got:
olleh
want:
golly
FAIL
```

```go
func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
}
```

```text
$ go test -v
=== RUN TestReverse
--- PASS: TestReverse (0.00s)
PASS
ok  	github.com/golang/example/stringutil	0.009s
```
