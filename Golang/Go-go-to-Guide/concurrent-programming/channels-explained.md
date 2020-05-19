# Channels Offer Synchronized Communication

* A **channel** is a mechanism for goroutines to **synchronize execution** and **communicate** by passing values

## Channel

* A new channel value can be made using the built-in function `make`

```go
// unbuffered channel of ints
ic := make(chan int)

// buffered channel with room for 10 strings
sc := make(chan string, 10)
```

* **To send** a value on a channel, use `<-` as a binary operator
* **To receive** a value on a channel, use it as a unary operator

```go
ic <- 3   // Send 3 on the channel.
n := <-sc // Receive a string from the channel.
```

* The `<-` operator specifies the channel direction, **send** or **receive**
* If no direction is given, the channel is **bi-directional**

```go
chan Sushi    // can be used to send and receive values of type Sushi
chan<- string // can only be used to send strings
<-chan int    // can only be used to receive ints
```

## Buffered and Unbuffered Channels

* If the capacity of a channel is zero or absent, the channel is **unbuffered** and the sender blocks until the receiver has received the value
* If the channel **has a buffer**, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value
* Receivers always block until there is data to receive
* Sending or receiving from a `nil` channel blocks forever

## Closing a Channel

* The `close` function records that no more values will be sent on a channel
* Note that it is **only necessary** to close a channel if a receiver is looking for a close
  * After calling `close`, and after any previously sent values have been received, receive operations will return a zero value without blocking
  * A multi-valued receive operation additionally returns an indication of whether the channel is closed
  * Sending to or closing a closed channel causes a run-time panic. Closing a nil channel also causes a run-time panic

```go
ch := make(chan string)
go func() {
    ch <- "Hello!"
    close(ch)
}()

fmt.Println(<-ch) // Print "Hello!".
fmt.Println(<-ch) // Print the zero value "" without blocking.
fmt.Println(<-ch) // Once again print "".
v, ok := <-ch     // v is "", ok is false.

// Receive values from ch until closed.
for v := range ch {
    fmt.Println(v) // Will not be executed.
}
```

## Example

* In the following example we let the `Publish` function return a channel, which is used to broadcast a message when the text has been published

```go
// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}
```

* Note that we use a channel of empty structs to indicate that the channel will only be used for **signalling**, not for passing data
* This is how you might use the function

```go
wait := Publish("important news", 2 * time.Minute)
// Do some more work.
<-wait // Block until the text has been published.
```
