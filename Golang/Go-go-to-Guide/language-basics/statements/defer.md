# Defer a Function Call (with Return Value)

## Defer Statement Basics

* A `defer` statement postpones the execution of a function until the surrounding function returns, either normally or through a panic

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
// Hello
// World
```

* Deferred calls are executed even when the function panics:

```go
func main() {
    defer fmt.Println("World")
    panic("Stop")
    fmt.Println("Hello")
}

// World
// panic: Stop
//
// goroutine 1 [running]:
// main.main()
//     ../main.go:3 +0xa0
```

### Order of Execution

* The deferred call’s **arguments are evaluated immediately**, even though the function call is not executed until the surrounding function returns
* If there are several deferred function calls, they are executed in **last-in-first-out** order

```go
func main() {
    fmt.Println("Hello")
    for i := 1; i <= 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("World")
}
// Hello
// World
// 3
// 2
// 1
```

### Use func to Return a Value

* Deferred anonymous functions may access and modify the surrounding function’s named return parameters
* In this example, the foo function returns “Change World”

```go
func foo() (result string) {
    defer func() {
        result = "Change World" // change value at the very last moment
    }()
    return "Hello World"
}
```

## Common Applications

* Defer is often used to perform clean-up actions, such as closing a file or unlocking a mutex
* Such actions should be performed both when the function returns normally and when it panics

### Close a File

* In this example, defer statements are used to ensure that all files are closed before leaving the `CopyFile` function, whichever way that happens

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```
