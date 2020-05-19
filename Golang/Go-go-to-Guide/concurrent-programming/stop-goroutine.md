# How to Kill a Goroutine

* **One goroutine can't forcibly stop another**

## Kill Goroutines

* To make a goroutine stoppable, let it listen for a stop signal on a channel

```go
quit := make(chan struct{})
go func() {
    for {
        select {
        case <-quit:
            return
        default:
            // …
        }
    }
}()
// …
close(quit)
```

* Sometimes it’s convenient to use a single channel for both data and signalling

```go
// Generator returns a channel that produces the numbers 1, 2, 3,…
// To stop the underlying goroutine, close the channel.
func Generator() chan int {
    ch := make(chan int)
    go func() {
        n := 1
        for {
            select {
            case ch <- n:
                n++
            case <-ch:
                return
            }
        }
    }()
    return ch
}

func main() {
    number := Generator()
    fmt.Println(<-number)
    fmt.Println(<-number)
    close(number)
    // …
}

// 1
// 2
```
