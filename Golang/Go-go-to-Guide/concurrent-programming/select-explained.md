# Select Waits on a Group of Channels

## Select

* The `select` statement waits for multiple send or receive operations simul­taneously
  * The statement blocks as a whole until one of the operations becomes unblocked
  * If several cases can proceed, a single one of them will be chosen at random

```go
// blocks until there's data available on ch1 or ch2
select {
case <-ch1:
    fmt.Println("Received from ch1")
case <-ch2:
    fmt.Println("Received from ch2")
}
```

* Send and receive operations on a `nil` channel block forever
* This can be used to disable a channel in a select statement:

```go
ch1 = nil // disables this channel
select {
case <-ch1:
    fmt.Println("Received from ch1") // will not happen
case <-ch2:
    fmt.Println("Received from ch2")
}
```

## Default Case

* The `default` case is always able to proceed and runs if all other cases are blocked

```go
// never blocks
select {
case x := <-ch:
    fmt.Println("Received", x)
default:
    fmt.Println("Nothing available")
}
```

## Examples

### An Infinite Random Binary Sequence

```go
rand := make(chan int)
for {
    select {
    case rand <- 0: // no statement
    case rand <- 1:
    }
}
```

### A Blocking Operation with a Timeout

* The function `time.After` is part of the standard library; it waits for a specified time to elapse and then sends the current time on the returned channel

```go
select {
case news := <-AFP:
    fmt.Println(news)
case <-time.After(time.Minute):
    fmt.Println("Time out: No news in one minute")
}
```

### A Statement that Blocks Forever

* A `select` statement blocks until **at least one** of it’s cases can proceed
* With zero cases this will never happen

```go
select {}
```

* A typical use would be at the end of the main function in some multithreaded programs
* When main returns, the program exits and it does not wait for other goroutines to complete
