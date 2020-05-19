# Maps Explained: Create, Add, Get, Delete

## Create A New Map

```go
var m map[string]int                // nil map of string-int pairs

m1 := make(map[string]float64)      // Empty map of string-float64 pairs
m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

m3 := map[string]float64{           // Map literal
    "e":  2.71828,
    "pi": 3.1416,
}
fmt.Println(len(m3))                // Size of map: 2
```

* A map (or dictionary) is an **unordered** collection of **key-value** pairs, where each key is **unique**
* You create a new map with a `make` statement or a **map literal**
* The default **zero value** of a map is `nil`
  * A nil map is equivalent to an empty map except that **elements can’t be added**
* The `len` function returns the **size** of a map, which is the number of key-value pairs

### Assignment to Entry in Nil Map

```go
var m map[string]float64
m["pi"] = 3.1416
// panic: assignment to entry in nil map
```

* You have to initialize the map using the `make` function (or a map literal) before you can add any elements

```go
m := make(map[string]float64)
m["pi"] = 3.1416
```

## Add, Update, Get and Delete Keys/Values

```go
m := make(map[string]float64)

m["pi"] = 3.14             // Add a new key-value pair
m["pi"] = 3.1416           // Update value
fmt.Println(m)             // Print map: "map[pi:3.1416]"

v := m["pi"]               // Get value: v == 3.1416
v = m["pie"]               // Not found: v == 0 (zero value)

_, found := m["pi"]        // found == true
_, found = m["pie"]        // found == false

if x, found := m["pi"]; found {
    fmt.Println(x)
}                           // Prints "3.1416"

delete(m, "pi")             // Delete a key-value pair
fmt.Println(m)              // Print map: "map[]"
```

* When you index a map you get two return values; the second one (which is optional) is a boolean that indicates if the key exists
* If the key doesn’t exist, the first value will be the default zero value

## For-each Range Loop

```go
m := map[string]float64{
    "pi": 3.1416,
    "e":  2.71828,
}
fmt.Println(m) // "map[e:2.71828 pi:3.1416]"

for key, value := range m { // Order not specified 
    fmt.Println(key, value)
}
```

* Iteration order is **not specified** and may vary from iteration to iteration
* If an entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced
* If an entry is created during iteration, that entry **may or may not** be produced during the iteration

## Performance and Implementation

* Maps are backed by hash tables
* Add, get and delete operations run in constant expected time. The time complexity for the add operation is amortized
* The comparison operators == and != must be defined for the key type
