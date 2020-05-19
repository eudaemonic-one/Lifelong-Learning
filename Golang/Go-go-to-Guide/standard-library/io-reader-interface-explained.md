# How to Use the `io.Reader` Interface

## Basics

* The `io.Reader` interface represents an entity from which you can read a stream of bytes

```go
type Reader interface {
        Read(buf []byte) (n int, err error)
}
```

* Read reads up to `len(buf)` bytes into buf and returns the number of bytes read – it returns an `io.EOF` error when the stream ends
* The standard library provides numerous Reader implementations (including in-memory byte buffers, files and network connections), and Readers are accepted as input by many utilities (including the HTTP client and server implementations)

## Use a Built-in Reader

* As an example, you can create a Reader from a string using the `strings.Reader` function and then pass the Reader directly to the `http.Post` function in package `net/http`
  * The Reader is then used as the source for the data to be posted
* Since `http.Post` uses a `Reader` instead of a `[]byte` it’s trivial to, for instance, use the contents of a file instead

```go
r := strings.NewReader("my request")
resp, err := http.Post("http://foo.bar",
	"application/x-www-form-urlencoded", r)
```

### Read Directly from a Byte Stream

* You can use the `Read` function directly

```go
r := strings.NewReader("abcde")

buf := make([]byte, 4)
for {
	n, err := r.Read(buf)
	fmt.Println(n, err, buf[:n])
	if err == io.EOF {
		break
	}
}

// 4 <nil> [97 98 99 100]
// 1 <nil> [101]
// 0 EOF []
```

* Use `io.ReadFull` to read exactly `len(buf)` bytes into `buf`

```go
r := strings.NewReader("abcde")

buf := make([]byte, 4)
if _, err := io.ReadFull(r, buf); err != nil {
	log.Fatal(err)
}
fmt.Println(buf)

if _, err := io.ReadFull(r, buf); err != nil {
	fmt.Println(err)
}

// [97 98 99 100]
// unexpected EOF
```

* Use `ioutil.ReadAll` to read everything:

```go
r := strings.NewReader("abcde")

buf, err := ioutil.ReadAll(r)
if err != nil {
	log.Fatal(err)
}
fmt.Println(buf)

// [97 98 99 100 101]
```

### Buffered Reading and Scanning

* The `bufio.Reader` and `bufio.Scanner` types wrap a `Reader` creating another `Reader` that also implements the interface but provides buffering and some help for textual input

```go
const input = `Beware of bugs in the above code;
I have only proved it correct, not tried it.`

scanner := bufio.NewScanner(strings.NewReader(input))
scanner.Split(bufio.ScanWords) // Set up the split function.

count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Println(err)
}
fmt.Println(count)

// 16
```
