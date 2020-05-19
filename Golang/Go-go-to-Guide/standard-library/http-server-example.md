# Hello World HTTP Server Example

## A Basic Web Server

* If you access the URL `http://localhost:8080/world` on a machine where the program below is running, you will be greeted by the hello world page

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
```

* The call to `http.HandleFunc` tells the `net.http` package to handle all requests to the web root with the `HelloServer` function
* The call to http.ListenAndServe tells the server to listen on the TCP network address `:8080`. This function blocks until the program is terminated
* Writing to an `http.ResponseWriter` sends data to the HTTP client
* An `http.Request` is a data structure that represents a client HTTP request.
* `r.URL.Path` is the path component of the requested URL. In this case, `"/world"` is the path component of `"http://localhost:8080/world"
