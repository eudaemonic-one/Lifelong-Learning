# List All Files (Recursively) in a Directory

## Directory Listing

* Use the `ioutil.ReadDir` function in package `io/ioutil`
* It returns a sorted slice containing elements of type `os.FileInfo`

```go
files, err := ioutil.ReadDir(".")
if err != nil {
    log.Fatal(err)
}
for _, f := range files {
    fmt.Println(f.Name())
}

// dev
// etc
// tmp
// usr
```

## Visit All Files and Folders in a Directory Tree

* Use the `filepath.Walk` function in package `path/filepath`
  * It walks a file tree calling a function of type `filepath.WalkFunc` for each file or directory in the tree, including the root
  * The files are walked in lexical order
  * Symbolic links are not followed

```go
err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.Size())
    return nil
})
if err != nil {
    log.Println(err)
}

// . 1644
// dev 1644
// dev/null 0
// dev/random 0
```
