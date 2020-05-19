# Get Slices of Keys and Values from a Map

## Map to Slice

```go
keys := make([]keyType, 0, len(myMap))
values := make([]valueType, 0, len(myMap))

for k, v := range myMap {
	keys = append(keys, k)
	values = append(values, v)
}
```
