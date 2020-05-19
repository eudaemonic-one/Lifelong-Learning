# Pick the Right One: `int` vs. `int64`

## Use `int` for Indexing

* An **index**, **length** or **capacity** should normally be an `int`
* The `int` type is either 32 or 64 bits, and always big enough to hold the maximum possible length of an array

## Use `int64` and Friends for Data

* The types `int8`, `int16`, `int32`, and `int64` (and their unsigned counterparts) are best suited for **data**
* An `int64` is the typical choice when memory isn’t an issue
* In particular, you can use a `byte`, which is an alias for `uint8`, to be extra clear about your intent
* Similarly, you can use a `rune`, which is an alias for `int32`, to emphasize than an integer represents a code point

## Examples

* In this code, the slice elements and the `max` variable have type `int64`, while the index and the length of the slice have type `int`

```go
func Max(a []int64) int64 {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
```

* The implementation of `time.Duration` is a typical example from the standard library where an int64 is used to store data:
  * A Duration represents the time between two instants as a nanosecond count
  * This limits the largest possible duration to about 290 years

```go
type Duration int64
```
