# Get Year, Month, Day From Time

* The `Date` function returns the year, month and day of a `time.Time`

```go
func (t Time) Date() (year int, month Month, day int)

year, month, day := time.Now().Date()
fmt.Println(year, month, day)      // For example 2009 November 10
fmt.Println(year, int(month), day) // For example 2009 11 10

t := time.Now()
year := t.Year()   // type int
month := t.Month() // type time.Month
day := t.Day()     // type int
```

* The `time.Month` type specifies a month of the year (January = 1, …)

```go
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```
