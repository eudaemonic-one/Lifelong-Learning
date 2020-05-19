# Mutual Exclusive Lock (Mutex)

* **Mutexes let you synchronize data access by explicit locking, without channels**

## Use with Caution

* For this type of locking to be safe, it’s crucial that all accesses to the shared data, both reads and writes, are performed **only** when a goroutine holds the lock
* One mistake by a single goroutine is enough to introduce a data race and break the program
* Because of this you should consider designing a custom data structure with a clean API and make sure that all the synchronization is done **internally**
* In this example we build a safe and easy-to-use concurrent data structure, `AtomicInt`, that stores a single integer
* Any number of goroutines can safely access this number through the `Add` and `Value` methods

```go
// AtomicInt is a concurrent data structure that holds an int.
// Its zero value is 0.
type AtomicInt struct {
    mu sync.Mutex // A lock than can be held by one goroutine at a time.
    n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add(n int) {
    a.mu.Lock() // Wait for the lock to be free and then take it.
    a.n += n
    a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
    a.mu.Lock()
    n := a.n
    a.mu.Unlock()
    return n
}

func main() {
    wait := make(chan struct{})
    var n AtomicInt
    go func() {
        n.Add(1) // one access
        close(wait)
    }()
    n.Add(1) // another concurrent access
    <-wait
    fmt.Println(n.Value()) // 2
}
```
