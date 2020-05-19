# Access Environment Variables

* Use the `os.Setenv`, `os.Getenv` and `os.Unsetenv` functions to access environment variables

```go
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/bash"

os.Unsetenv("SHELL")
fmt.Printf("%q\n", os.Getenv("SHELL")) // ""

os.Setenv("SHELL", "/bin/dash")
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/dash"
```

* The `os.Environ` function returns a slice of "key=value" strings listing all environment variables

```go
for _, s := range os.Environ() {
    kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
    fmt.Printf("key:%q value:%q\n", kv[0], kv[1])
}

// key:"SHELL" value:"/bin/bash"
// key:"SESSION" value:"ubuntu"
// key:"TERM" value:"xterm-256color"
// key:"LANG" value:"en_US.UTF-8"
// key:"XMODIFIERS" value:"@im=ibus"
// ...
```
