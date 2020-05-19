# How to Detect Data Races

## Detect Data Races

* Data races can happen easily and are hard to debug
* **Luckily**, the Go runtime is often able to help
* Use `-race` to enable the built-in data race detector

```text
$ go test -race [packages]
$ go run -race [packages]
```

## Example

```go
package main
import "fmt"

func main() {
    i := 0
    go func() {
        i++ // write
    }()
    fmt.Println(i) // concurrent read
}
```

* Running this program with the `-race` options tells us that there’s a race between the write at line 7 and the read at line 9:

```text
$ go run -race main.go
0
==================
WARNING: DATA RACE
Write by goroutine 6:
  main.main.func1()
      /tmp/main.go:7 +0x44

Previous read by main goroutine:
  main.main()
      /tmp/main.go:9 +0x7e

Goroutine 6 (running) created at:
  main.main()
      /tmp/main.go:8 +0x70
==================
Found 1 data race(s)
exit status 66
```

## Details

* The data race detector does not perform any static analysis
* It checks the memory access in runtime and only for the code paths that are actually executed
* It runs on darwin/amd64, freebsd/amd64, linux/amd64 and windows/amd64
* The overhead varies, but typically there’s a 5-10x increase in memory usage, and 2-20x increase in execution time
