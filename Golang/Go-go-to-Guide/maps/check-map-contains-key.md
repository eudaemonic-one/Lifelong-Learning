# 3 Ways to Find a Key in a Map

## Basics

* When you index a map in Go you get two return values; the second one (which is optional) is a boolean that indicates if the key exists
* If the key doesn’t exist, the first value will be the default zero value.

## Check Second Return Value

```go
m := map[string]float64{"pi": 3.14}
v, found := m["pi"] // v == 3.14  found == true
v, found = m["pie"] // v == 0.0   found == false
_, found = m["pi"]  // found == true
```

## Use Second Return Value Directly in an if Statement

```go
m := map[string]float64{"pi": 3.14}
if v, found := m["pi"]; found {
    fmt.Println(v)
}
// Output: 3.14
```

## Check for Zero Value

```go
m := map[string]float64{"pi": 3.14}

v := m["pi"] // v == 3.14
v = m["pie"] // v == 0.0 (zero value)
```

* **Warning:** This approach doesn't work if the zero value is a possible key
