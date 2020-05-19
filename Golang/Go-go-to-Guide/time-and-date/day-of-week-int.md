# How to Find the Day of Week

* The `Weekday` function returns returns the day of the week of a `time.Time`

```go
func (t Time) Weekday() Weekday

weekday := time.Now().Weekday()
fmt.Println(weekday)      // "Tuesday"
fmt.Println(int(weekday)) // "2"
```

## Type Weekday

* The `time.Weekday` type specifies a day of the week (Sunday = 0, …)

```go
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```
