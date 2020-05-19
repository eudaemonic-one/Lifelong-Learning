# Data Races Explained

## Data Races

* A **data race** happens when two goroutines access the same variable concur­rently, and at least one of the accesses is a write
* Data races are quite common and can be very hard to debug
* This function has a data race and it’s behavior is undefined
  * It may, for example, print the number 1

```go
func race() {
    wait := make(chan struct{})
    n := 0
    go func() {
        n++ // read, increment, write
        close(wait)
    }()
    n++ // conflicting access
    <-wait
    fmt.Println(n) // Output: <unspecified>
}
```

* The two goroutines, g1 and g2, participate in a race and there is no way to know in which order the operations will take place
* The following is one out of many possible outcomes

| g1                            | g2                            |
| ----------------------------- | ----------------------------- |
| Read the value 0 from `n`.    |                               |
|                               | Read the value 0 from `n`.    |
| Incre­ment value from 0 to 1. |                               |
| Write 1 to `n`.               |                               |
|                               | Incre­ment value from 0 to 1. |
|                               | Write 1 to `n`.               |
| Print `n`, which is now 1.    |                               |
|                               |                               |

## How to Avoid Data Races

* The only way to avoid data races is to **synchronize access to all mutable data that is shared between threads**
* In Go, you would normally use a **channel** or a **lock**. (Lower-lever mechanisms are available in the [`sync`](https://golang.org/pkg/sync/) and [`sync/atomic`](https://golang.org/pkg/sync/atomic/) packages.)
* The preferred way to handle concurrent data access in Go is to use a channel to pass the actual data from one goroutine to the next
* The motto is: **“Don’t communicate by sharing memory; share memory by communicating.”**

```go
func sharingIsCaring() {
    ch := make(chan int)
    go func() {
        n := 0 // A local variable is only visible to one goroutine.
        n++
        ch <- n // The data leaves one goroutine...
    }()
    n := <-ch // ...and arrives safely in another.
    n++
    fmt.Println(n) // Output: 2
}
```

* In this code the channel does double duty:
  * it passes the data from one goroutine to another,
  * and it acts as a point of synchronization
* The sending goroutine will wait for the other goroutine to receive the data and the receiving goroutine will wait for the other goroutine to send the data
