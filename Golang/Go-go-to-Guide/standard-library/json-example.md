# How to use JSON with Go

* **The JSON data-interchange format is easy for humans to read and write, and efficient for machines to parse and generate**

## Default Types

* The default Go types for decoding and encoding JSON are
  - `bool` for JSON booleans,
  - `float64` for JSON numbers,
  - `string` for JSON strings, and
  - `nil` for JSON null
* Additionally, `time.Time` and the numeric types in the `math/big` package can be automatically encoded as JSON strings
* Note that JSON **doesn’t** support basic integer types
  * They can often be approximated by floating-point numbers

## Encode (Marshal) Struct to JSON

* The `json.Marshal` function in package `encoding/json` generates JSON data

```go
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

basket := FruitBasket{
    Name:    "Standard",
    Fruit:   []string{"Apple", "Banana", "Orange"},
    Id:      999,
    private: "Second-rate",
    Created: time.Now(),
}

var jsonData []byte
jsonData, err := json.Marshal(basket)
if err != nil {
    log.Println(err)
}
fmt.Println(string(jsonData))

// "Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999,"Created":"2018-04-09T23:00:00Z"}
```

* Only data that can be represented as JSON will be encoded
  * Only the exported (public) fields of a struct will be present in the JSON output
    * **Other fields are ignored**
  * A field with a `json:` **struct tag** is stored with its tag name instead of its variable name
  * Pointers will be encoded as the values they point to, or `null` if the pointer is `nil`

## Pretty Print

* Replace `json.Marshal` with `json.MarshalIndent` in the example above to indent the JSON output

```go
jsonData, err := json.MarshalIndent(basket, "", "    ")

// {
//     "Name": "Standard",
//     "Fruit": [
//         "Apple",
//         "Banana",
//         "Orange"
//     ],
//     "ref": 999,
//     "Created": "2018-04-09T23:00:00Z"
// }
```

## Decode (Unmarshal) JSON to Struct

* The `json.Unmarshal` function in package `encoding/json` parses JSON data

```go
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64 `json:"ref"`
    Created time.Time
}

jsonData := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

var basket FruitBasket
err := json.Unmarshal(jsonData, &basket)
if err != nil {
    log.Println(err)
}
fmt.Println(basket.Name, basket.Fruit, basket.Id)
fmt.Println(basket.Created)

// Standard [Apple Banana Orange] 999
// 2018-04-09 23:00:00 +0000 UTC
```

* Note that `Unmarshal` allocated a new slice all by itself
  * This is how unmarshaling works for slices, maps and pointers
* For a given JSON key `Foo`, `Unmarshal` will attempt to match the struct fields in this order:
  * an exported (public) field with a struct tag `json:"Foo"`,
  * an exported field named `Foo`, or
  * an exported field named `FOO`, `FoO`, or some other case-insensitive match
* **Only** fields thar are found in the destination type will be decoded:
  * This is useful when you wish to pick only a few specific fields
  * In particular, any unexported fields in the destination struct will be unaffected

## Arbitrary Objects and Arrays

* The `encoding/json` package uses
  * `map[string]interface{}` to store arbitrary JSON objects, and
  * `[]interface{}` to store arbitrary JSON arrays
* It will unmarshal any valid JSON data into a plain `interface{}` value
* Consider this JSON data:

```text
{
    "Name": "Eve",
    "Age": 6,
    "Parents": [
        "Alice",
        "Bob"
    ]
}
```

* The `json.Unmarshal` function will parse it into a map whose keys are strings, and whose values are themselves stored as empty interface values:

```go
map[string]interface{}{
    "Name": "Eve",
    "Age":  6.0,
    "Parents": []interface{}{
        "Alice",
        "Bob",
    },
}
```

* We can iterate through the map with a range statement and use a type switch to access its values

```go
jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

var v interface{}
json.Unmarshal(jsonData, &v)
data := v.(map[string]interface{})

for k, v := range data {
    switch v := v.(type) {
    case string:
        fmt.Println(k, v, "(string)")
    case float64:
        fmt.Println(k, v, "(float64)")
    case []interface{}:
        fmt.Println(k, "(array):")
        for i, u := range v {
            fmt.Println("    ", i, u)
        }
    default:
        fmt.Println(k, v, "(unknown)")
    }
}

// Name Eve (string)
// Age 6 (float64)
// Parents (array):
//      0 Alice
//      1 Bob
```

## JSON File Example

* The `json.Decoder` and `json.Encoder` types in package `encoding/json` offer support for reading and writing streams, e.g. files, of JSON data
* The code in this example
  * reads a stream of JSON objects from a `Reader` (`strings.Reader`),
  * removes the `Age` field from each object,
  * and then writes the objects to a `Writer` (`os.Stdout`)

```go
const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
`
reader := strings.NewReader(jsonData)
writer := os.Stdout

dec := json.NewDecoder(reader)
enc := json.NewEncoder(writer)

for {
    // Read one JSON object and store it in a map.
    var m map[string]interface{}
    if err := dec.Decode(&m); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }

    // Remove all key-value pairs with key == "Age" from the map.
    for k := range m {
        if k == "Age" {
            delete(m, k)
        }
    }

    // Write the map as a JSON object.
    if err := enc.Encode(&m); err != nil {
        log.Println(err)
    }
}

// {"Name":"Alice"}
// {"Name":"Bob"}
```
