# Days in a Month

* To compute the last day of a month, you can use the fact that `time.Date` accepts values outside their usual ranges – the values are normalized during the conversion

```go
func main() {
    t := Date(2000, 3, 0) // the day before 2000-03-01
    fmt.Println(t)        // 2000-02-29 00:00:00 +0000 UTC
    fmt.Println(t.Day())  // 29
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
```

* AddDate normalizes its result in the same way
  * For example, adding one month to October 31 yields December 1, the normalized form of November 31

```go
t = Date(2000, 10, 31).AddDate(0, 1, 0) // a month after October 31
fmt.Println(t)                          // 2000-12-01 00:00:00 +0000 UTC
```
