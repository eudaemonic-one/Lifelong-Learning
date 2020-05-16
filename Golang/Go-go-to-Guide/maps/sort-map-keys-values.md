# Sort a Map by Key or Value

## Sort a Map by Key

* A map is an **unordered** collection of key-value pairs
* If you need a stable iteration order, you must maintain a separate data structure
  * Also, starting with Go 1.12, the fmt package prints maps in key-sorted order to ease testing

```go
m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
	fmt.Println(k, m[k])
}

// Alice 23
// Bob 25
// Eve 2
```

## Sort a Map by Value

```go
type pair struct {
	Key string
	Value int
}

m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

pairs := make([]pair, 0)
for k, v := range m {
	pairs = append(pairs, pair{k, v})
}
sort.Strings(pairs, func(i, j int) bool {
	return pairs[i].Value < pairs[j].Value
})

for _, pair := range pairs {
	fmt.Println(pair.Key, pair.Value)
}

// Eve 2
// Alice 23
// Bob 25
```
