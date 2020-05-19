# Read a File (stdin) Line By Line

## Read From File

* Use a `bufio.Scanner` to read a file line by line

```go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

## Read From Stdin

* Use `os.Stdin` to read from the standard input stream

```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	log.Println(err)
}
```

### Read From Any Stream

* A `bufio.Scanner` can read from any stream of bytes, as long as it implements the `io.Reader` interface
