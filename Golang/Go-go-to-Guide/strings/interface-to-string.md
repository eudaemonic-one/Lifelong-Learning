# Convert Interface to String

## Interface to String

```go
var x interface{} = "abc"
str := fmt.Sprintf("%v", x)

var x interface{} = []int{1, 2, 3}
str := fmt.Sprintf("%v", x)
fmt.Println(str) // "[1 2 3]"
```
