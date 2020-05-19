# How to Debug Deadlocks

## Deadlocks

* A **deadlock** happens when a group of goroutines are waiting for each other and none of them is able to proceed

```go
func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
```

* The program will get stuck on the channel send operation waiting forever for someone to read the value
* Go is able to detect situations like this **at runtime**

```text
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	.../deadlock.go:7 +0x6c
```

## Debugging Tips

* A goroutine can get stuck
  - either because it’s waiting for a **channel** or
  - because it is waiting for one of the **locks** in the [sync](https://golang.org/pkg/sync/) package
* **Common reasons** are that
  - no other goroutine has **access** to the channel or the lock,
  - a group of goroutines are waiting for **each other** and none of them is able to proceed
* Currently Go only detects when the program as a whole freezes, not when a subset of goroutines get stuck
* With channels it’s often easy to figure out what caused a deadlock
* Programs that make heavy use of mutexes can, on the other hand, be notoriously difficult to debug
