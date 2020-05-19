# Error Handling Best Practice

## Errors

* Go has two different error-handling mechanisms:
  * most functions return errors;
  * only a truly unrecoverable condition, such as an out-of-range index, produces a run-time exception, known as a panic
* Go’s multivalued return makes it easy to return a detailed error message alongside the normal return value
* By convention, such messages have type `error`, a simple built-in interface:

```go
type error interface {
    Error() string
}
```

## Error Handling Example

* The `os.Open` function returns a non-nil `error` value when it fails to open a file

```go
func Open(name string) (file *File, err error)

f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

## Custom Errors

* To create a simple string-only error you can use `errors.New`:

```go
err := errors.New("Houston, we have a problem")
```

* The `error` interface requires only an `Error` method, but specific `error` implementations often have additional methods, allowing callers to inspect the details of the error
