# Command-line Arguments and Flags

## Command-line Arguments

* The `os.Args` variable holds the command-line arguments – starting with the program name – which are passed to a Go program

```go
func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage:", os.Args[0], "PATTERN", "FILE")
        return
    }
    pattern := os.Args[1]
    file := os.Args[2]
    // ...
}
```

```text
$ go build grep.go
$ ./grep
Usage: ./grep PATTERN FILE
```

## Flag Parsing

* The [flag](https://golang.org/pkg/flag/) package implements basic command-line flag parsing
