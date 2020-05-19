# Measure Execution Time

## Measure a Piece of Code

```go
start := time.Now()
// Code to measure
duration := time.Since(start)

// Formatted string, such as "2h3m0.5s" or "4.503μs"
fmt.Println(duration)

// Nanoseconds as int64
fmt.Println(duration.Nanoseconds())
```

## Measure a Function Call

```go
func foo() {
    defer duration(track("foo"))
    // Code to measure
}

func track(msg string) (string, time.Time) {
    return msg, time.Now()
}

func duration(msg string, start time.Time) {
    log.Printf("%v: %v\n", msg, time.Since(start))
}
```

## Benchmarks

* The testing package has support for benchmarking that can be used to examine the performance of your code
