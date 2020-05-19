# Go as a Scripting Language: Lightweight, Safe, and Fast

* This example is a simplified version of the Unix grep command
* The program searches the input file for lines containing the given pattern and prints these lines

```go
func main() {
    log.SetPrefix("grep: ")
    log.SetFlags(0) // no extra info in log messages

    if len(os.Args) != 3 {
        fmt.Printf("Usage: %v PATTERN FILE\n", os.Args[0])
        return
    }

    pattern, err := regexp.Compile(os.Args[1])
    if err != nil {
        log.Fatalln(err)
    }

    file, err := os.Open(os.Args[2])
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if pattern.MatchString(line) {
            fmt.Println(line)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
```
