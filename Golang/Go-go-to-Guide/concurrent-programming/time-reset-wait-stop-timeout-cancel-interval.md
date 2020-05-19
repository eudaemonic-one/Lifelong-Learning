# Timer and Ticker: Events in the Future

* **Timers and Tickers let you execute code in the future, once or repeatedly**

## Timeout (Timer)

* `time.After` waits for a specified duration and then sends the current time on the returned channel:

```go
select {
case news := <-AFP:
	fmt.Println(news)
case <-time.After(time.Hour):
	fmt.Println("No news in an hour.")
}
```

* The underlying `time.Timer` will not be recovered by the garbage collector until the timer fires
* If this is a concern, use `time.NewTimer` instead and call its `Stop` method when the timer is no longer needed:

```go
for alive := true; alive; {
	timer := time.NewTimer(time.Hour)
	select {
	case news := <-AFP:
		timer.Stop()
		fmt.Println(news)
	case <-timer.C:
		alive = false
		fmt.Println("No news in an hour. Service aborting.")
	}
}
```

## Repeat (Ticker)

* `time.Tick` returns a channel that delivers clock ticks at even intervals:

```go
go func() {
	for now := range time.Tick(time.Minute) {
		fmt.Println(now, statusUpdate())
	}
}()
```

* The underlying `time.Ticker` will not be recovered by the garbage collector
* If this is a concern, use `time.NewTicker` instead and call its `Stop` method when the ticker is no longer needed

## Wait, Act and Cancel

* `time.AfterFunc` waits for a specified duration and then calls a function in its own goroutine
* It returns a `time.Timer` that can be used to cancel the call:

```go
func Foo() {
    timer = time.AfterFunc(time.Minute, func() {
        log.Println("Foo run for more than a minute.")
    })
    defer timer.Stop()

    // Do heavy work
}
```
