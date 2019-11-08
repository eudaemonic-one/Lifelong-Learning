# How to Write Go Code

## Introduction

This document introduces the development process of a Go package and go tool, and the way to fetch, build, and install Go packages and commands.

### Code organization

#### Overview

* Go programmers typically keep all their Go code in a single workspace.
* A workspace contains many version control repositories (managed by Git, for example).
* Each repository contains one or more packages.
* Each package consists of one or more Go source files in a single directory.
* The path to a package's directory determines its import path.

#### Workspaces

A workspace is a directory hierarchy with two directories at its root:

* src contains Go source files, and
* bin contains executable commands.

The go tool builds and installs binaries to the bin directory.

The src subdirectory typically contains multiple version control repositories (such as for Git or Mercurial) that track the development of one or more source packages.

e.g.

```text
bin/
    hello                          # command executable
    outyet                         # command executable
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
    hello/
    hello.go               # command
        source
    outyet/
        main.go                # command source
        main_test.go           # test source
    stringutil/
        reverse.go             # package source
        reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
    bmp/
        reader.go              # package source
        writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

Most Go programmers keep all their Go source code and dependencies in a single workspace.

### The *GOPATH* environment variable

The *GOPATH* environment variable specifies the location of workspace. It defaults to be a directory named go inside your home directory, so $HOME/go on Unix.

For convenience, add the workspace's bin subdirectory to your *PATH*:

```shell
$ export PATH=$PATH:$(go env GOPATH)/bin
```

#### Import paths

An *import path* is a string that uniquely identifies a package. A package's import path corresponds to its location inside a workspace or in a remote repository.

```shell
$ cd $GOPATH/src/github.com/user/hello
$ go install
$ hello
Hello, world.
```

This command builds the *hello* command, producing an executable binary. It then installs that binary to the workspace's bin directory as *hello*. Now, we can run program by typing its full path or just type the binary name if have added *$GOPATH/bin* to *PATH*.

#### Your first library

```shell
$ go build github.com/user/stringutil
```

This won't produce an output file. Instead it saves the compiled package in the local build cache.

#### Package names

The first statement in a Go source file must be

```go
package name
```

where name is the package's default name for imports. (All files in a package must use the same name.)

Go's convention is that the package name is the last element of the import path: the package imported as "crypto/rot13" should be named rot13.

#### Testing

Go has a lightweight test framework composed of the go test command and the testing package.

You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T). The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.

```go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

```shell
$ go test github.com/user/stringutil
ok     github.com/user/stringutil 0.165s
```

#### Remote packages

An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial. The go tool uses this property to automatically fetch packages from remote repositories.

```shell
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```
