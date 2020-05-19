# Packages Explained: Declare, Import, Download, Document

## Basics

* Every Go program is made up of packages and each package has an **import path**
  * `fmt`
  * `math/rand`
  * `github.com/yourbasic/graph`
* By convention, the **package name** is the same as the last element of the import path:
  * `fmt`
  * `rand`
  * `graph`
* References to other packages’ definitions must always be prefixed with their package names, and only the capitalized names from other packages are accessible

```go
package main

import (
    "fmt"
    "math/rand"

    "github.com/yourbasic/graph"
)

func main() {
    n := rand.Intn(100)
    g := graph.New(n)
    fmt.Println(g)
}
```

## Declare A Package

* Every Go source file starts with a package declaration, which contains only the package name
* For example, the file `src/math/rand/exp.go`, which is part of the implementation of the `math/rand` package, contains the following code

```go
package rand
  
import "math"
  
const re = 7.69711747013104972
...
```

## Package Name Conflicts

* You can customize the name under which you refer to an imported package

```go
package main

import (
    csprng "crypto/rand"
    prng "math/rand"

    "fmt"
)

func main() {
    n := prng.Int() // pseudorandom number
    b := make([]byte, 8)
    csprng.Read(b) // cryptographically secure pseudorandom number
    fmt.Println(n, b)
}
```

## Dot Imports

* If a period . appears instead of a name in an import statement, all the package’s exported identifiers can be accessed without a qualifier
* Dot imports can make programs hard to read and **generally should be avoided**

```go
package main

import (
    "fmt"
    . "math"
)

func main() {
    fmt.Println(Sin(Pi/2)*Sin(Pi/2) + Cos(Pi)/2) // 0.5
}
```

## Package Download

* The [go get](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) command downloads packages named by import paths, along with their dependencies, and then installs the packages

```text
$ go get github.com/yourbasic/graph
```

## Package Documentation

* The [GoDoc](https://godoc.org/) web site hosts documentation for all public Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad:
  * [`https://godoc.org/fmt`](https://godoc.org/fmt)
  * [`https://godoc.org/math/rand`](https://godoc.org/math/rand)
  * [`https://godoc.org/github.com/yourbasic/graph`](https://godoc.org/github.com/yourbasic/graph)
* The [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command extracts and generates documentation for all locally installed Go programs. The following command starts a web server that presents the documentation at `http://localhost:6060/`

```text
$ godoc -http=:6060 &
```
