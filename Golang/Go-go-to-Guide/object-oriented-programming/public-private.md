# Public vs. Private

## Package Encapsulation

* **A package is the smallest unit of private encapsulation in Go**
  * All identifiers defined within a package are visible throughout that package
  * When importing a package you can access only its **exported** identifiers
  * An identifier is exported if it begins with a **capital letter**
* Exported and unexported identifiers are used to describe the public interface of a package and to guard against certain programming errors
* **Warning:** Unexported identifiers is not a security measure and it does not hide or protect any information

## Example

* In this package, the only exported identifiers are `StopWatch` and `Start`

```go
package timer

import "time"

// A StopWatch is a simple clock utility.
// Its zero value is an idle clock with 0 total time.
type StopWatch struct {
    start   time.Time
    total   time.Duration
    running bool
}

// Start turns the clock on.
func (s *StopWatch) Start() {
    if !s.running {
        s.start = time.Now()
        s.running = true
    }
}
```

* The `StopWatch` and its exported methods can be imported and used in a different package

```go
package main

import "timer"

func main() {
    clock := new(timer.StopWatch)
    clock.Start()
    if clock.running { // ILLEGAL
        // …
    }
}

// ../main.go:8:15: clock.running undefined (cannot refer to unexported field or method clock.running)
```
