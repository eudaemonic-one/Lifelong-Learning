# Waiting for Goroutines

## Wait Goroutines

* A `sync.WaitGroup` waits for a group of goroutines to finish

```go
var wg sync.WaitGroup
wg.Add(2)
go func() {
    // Do work.
    wg.Done()
}()
go func() {
    // Do work.
    wg.Done()
}()
wg.Wait()
```

* First the main goroutine calls `Add` to set the number of goroutines to wait for
* Then two new goroutines run and call `Done` when finished
* At the same time, `Wait` is used to block until these two goroutines have finished
* **Note:** A `WaitGroup` must not be copied after first use
