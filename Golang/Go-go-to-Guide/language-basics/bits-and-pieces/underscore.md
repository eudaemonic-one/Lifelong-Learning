# Blank Identifier (Underscore)

* The blank identifier `_` is an anonymous placeholder
* It may be used like any other identifier in a declaration, but it does not introduce a binding

## Ignore Values

* The blank identifier provides a way to ignore left-hand side values in an assignment

```go
_, present := timeZone["CET"]

sum := 0
for _, n := range a {
	sum += n
}
```

## Import for Side Effects

* It can also be used to import a package solely for its side effects

```go
import _ "image/png" // init png decoder function
```

## Silence the Compiler

* It can be used to during development to avoid compiler errors about unused imports and variables in a half-written program

```go
package main

import (
    "fmt"
    "log"
    "os"
)

var _ = fmt.Printf // DEBUG: delete when done

func main() {
    f, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    _ = f // TODO: read file
}
```

* For an automatic solution, use the goimports tool, which rewrites a Go source file to have the correct imports
