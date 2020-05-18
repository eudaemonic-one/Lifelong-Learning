# User-friendly Access to `crypto/rand`

* Go has two packages for random numbers:
  * `math/rand` implements a large selection of pseudo-random number generators
  * `crypto/rand` implements a cryptographically secure pseudo-random number generator with a limited interface.

```go
import (
    crand "crypto/rand"
    rand "math/rand"

    "encoding/binary"
    "fmt"
    "log"
)

func main() {
    var src cryptoSource
    rnd := rand.New(src)
    fmt.Println(rnd.Intn(1000)) // a truly random number 0 to 999
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
    return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
    err := binary.Read(crand.Reader, binary.BigEndian, &v)
    if err != nil {
        log.Fatal(err)
    }
    return v
}
```
