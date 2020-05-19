# Constructors Deconstructed (Best Practice)

* **Go doesn't have explicit constructors**
* The idiomatic way to set up new data structures is to use proper **zero values** coupled with **factory** functions

## Zero Value

* Try to make the default zero value useful and document its behavior
* Sometimes this is all that’s needed

```go
// A StopWatch is a simple clock utility.
// Its zero value is an idle clock with 0 total time.
type StopWatch struct {
    start   time.Time
    total   time.Duration
    running bool
}

var clock StopWatch // Ready to use, no initialization needed.
```

* `StopWatch` takes advantage of the useful zero values of `time.Time`, `time.Duration` and `bool`
* In turn, users of `StopWatch` can benefit from *its* useful zero value

## Factory

* If the zero value doesn’t suffice, use factory functions named `NewFoo` or just `New`

```go
scanner := bufio.NewScanner(os.Stdin)
err := errors.New("Houston, we have a problem")
```
