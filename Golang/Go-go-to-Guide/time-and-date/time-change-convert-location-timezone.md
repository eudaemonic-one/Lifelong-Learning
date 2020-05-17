# Time Zones

## Convert Location Based Timezone

* Each `Time` has an associated `Location`, which is used for display purposes
* The method `In` returns a time with a specific location. Changing the location in this way changes only the presentation; it does not change the instant in time
* 

```go
// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
    loc, err := time.LoadLocation(name)
    if err == nil {
        t = t.In(loc)
    }
    return t, err
}

for _, name := range []string{
	"",
	"Local",
	"Asia/Shanghai",
	"America/Metropolis",
} {
	t, err := TimeIn(time.Now(), name)
	if err == nil {
		fmt.Println(t.Location(), t.Format("15:04"))
	} else {
		fmt.Println(name, "<time unknown>")
	}
}

// UTC 19:32
// Local 20:32
// Asia/Shanghai 03:32
// America/Metropolis <time unknown>
```
