# Broadcast a Signal on a Channel

* **All readers receive zero values on a closed channel**

## Broadcast a Signal on a Channel

* In this example the `Publish` function returns a channel, which is used to broadcast a signal when a message has been published

```go
// Print text after the given time has expired.
// When done, the wait channel is closed.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
    ch := make(chan struct{})
    go func() {
        time.Sleep(delay)
        fmt.Println("BREAKING NEWS:", text)
        close(ch) // Broadcast to all receivers.
    }()
    return ch
}
```

* Notice that we use a channel of empty structs: `struct{}`
* This clearly indicates that the channel will only be used for signalling, not for **passing data**

```go
func main() {
    wait := Publish("Channels let goroutines communicate.", 5*time.Second)
    fmt.Println("Waiting for news...")
    <-wait
    fmt.Println("Time to leave.")
}

// Waiting for news...
// BREAKING NEWS: Channels let goroutines communicate.
// Time to leave.
```
