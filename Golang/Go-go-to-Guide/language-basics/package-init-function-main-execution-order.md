# Package Initialization and Program Execution Order

## Basics

* First the `main` package is initialized
  * Imported packages are initialized before the package itself
  * Packages are initialized one at a time:
    * first package-level variables are initialized in declaration order,
    * then the `init` functions are run
* Finally the `main` function is called.

## Program Execution

* Program execution begins by initializing the `main` package and then calling the function `main`
* When `main` returns, the program exits
* It **does not wait** for other goroutines to complete

## Package Initialization

* Package-level variables are initialized in **declaration order**, but after any of the variables they **depend** on
* Initialization of variables declared in multiple files is done in **lexical file name order**
* Variables declared in the first file are declared before any of the variables declared in the second file
* Initialization cycles are **not allowed**
* Dependency analysis is performed **per package**; only references referring to variables, functions, and methods declared in the current package are considered

### Example

* In this example, taken directly from the Go language specification, the initialization order is d, b, c, a

```go
var (
    a = c + b
    b = f()
    c = f()
    d = 3
)

func f() int {
    d++
    return d
}
```

## Init Function

* Variables may also be initialized using `init` functions
* `func init() { … }`
* Multiple such functions may be defined
* They cannot be called from inside a program
  * A package with **no imports** is initialized
    * by assigning initial values to all its package-level variables,
    * followed by calling all `init` functions in the order they appear in the source
  * Imported packages are initialized before the package itself
  * Each package is initialized **once**, regardless if it’s imported by multiple other packages
* It follows that there can be **no cyclic dependencies**
* Package initialization happens in a single goroutine, sequentially, one package at a time

## Warning

* Lexical ordering according to file names is not part of the formal language specification
* To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler
