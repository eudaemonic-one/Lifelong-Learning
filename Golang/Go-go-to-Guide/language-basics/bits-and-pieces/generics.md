# Generics (Alternatives and Workarounds)

* Go has some built-in generic data types, such as slices and maps, and some generic functions, such as `append` and `copy`
* However, there is **no** mechanism for writing your own

## Find a Well-fitting Interface

* Describe the generic behaviour of your data with an interface
* The `io.Reader` interface, which represents the read end of a stream of data, is a good example:
  * many functions take an `io.Reader` as input,
  * and many data types, including files, network connections, and ciphers, implement this interface

## Use Multiple Functions

* If you only need to support a few data types, consider offering a separate function for each type
* As an example, the two packages strings and bytes come with pretty much the same set of functions
* If this leads to an unmanageable amount of copy and paste, consider using a code generation tool

## Use the Empty Interface

* If little is known about the data, consider using the empty interface `interface{}` in combination with type assertions, and possibly also reflection

## Write an Experience Report

* If none of these solutions are effective, consider submitting an experience report
