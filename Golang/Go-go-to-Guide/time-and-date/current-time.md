# How to Get Current Timestamp

* Use `time.Now` and one of `time.Unix` or `time.UnixNano` to get a timestamp

```go
now := time.Now()      // current local time
sec := now.Unix()      // number of seconds since January 1, 1970 UTC
nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

fmt.Println(now)  // time.Time
fmt.Println(sec)  // int64
fmt.Println(nsec) // int64

// 2009-11-10 23:00:00 +0000 UTC m=+0.000000000
// 1257894000
// 1257894000000000000
```
