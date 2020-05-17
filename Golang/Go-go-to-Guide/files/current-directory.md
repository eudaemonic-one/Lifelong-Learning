# Find Current Working Directory

## Current Directory

* Use `os.Getwd` to find the path name for the current directory
  * **Warning:** If the current directory can be reached via multiple paths (due to symbolic links), `Getwd` may return any one of them

```go
path, err := os.Getwd()
if err != nil {
    log.Println(err)
}
fmt.Println(path)  // for example /home/user
```

## Current Executable

* Use `os.Executable` to find the path name for the executable that started the current process
* **Warning:**
  * There is no guarantee that the path is still pointing to the correct executable
  * If a symlink was used to start the process, depending on the operating system, the result might be the symlink or the path it pointed to
  * If a stable result is needed, `path/filepath.EvalSymlinks` might help

```go
path, err := os.Executable()
if err != nil {
    log.Println(err)
}
fmt.Println(path) // for example /tmp/go-build872132473/b001/exe/main
```
