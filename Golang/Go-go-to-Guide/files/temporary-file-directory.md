# Create a Temporary File or Directory

## File

* Use `ioutil.TempFile` in package `io/ioutil` to create a **globally unique temporary file**
  * It’s your own job to **remove** the file when it’s **no longer needed**
* The call to `ioutil.TempFile`
  * creates a new file with a name starting with `"prefix"` in the directory `"dir"`,
  * opens the file for reading and writing,
  * and returns the new `*os.File`

```go
file, err := ioutil.TempFile("dir", "prefix")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(file.Name())

fmt.Println(file.Name()) // For example "dir/prefix054003078"
```

### Add a Suffix to the Temporary File Name

* Starting with Go 1.11, if the second string given to `TempFile` includes a `"*"`, the random string replaces this `"*"`

```go
file, err := ioutil.TempFile("dir", "myname.*.bat")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(file.Name())

fmt.Println(file.Name()) // For example "dir/myname.054003078.bat"
```

## Directory

* Use `ioutil.TempDir` in package `io/ioutil` to create a **globally unique temporary directory**
* The call to `ioutil.TempDir`
  * creates a new directory with a name starting with `"prefix"` in the directory `"dir"`
  * and returns the path of the new directory

```go
dir, err := ioutil.TempDir("dir", "prefix")
if err != nil {
	log.Fatal(err)
}
defer os.RemoveAll(dir)
```
