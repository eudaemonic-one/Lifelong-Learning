# Generate a Random String

## Random String

```go
rand.Seed(time.Now().UnixNano())
chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö" +
    "0123456789")
length := 8
var b strings.Builder
for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
}
str := b.String() // E.g. "ExcbsVQs"
```

* **Warning:** To generate a password, you should use cryptographically secure pseudorandom numbers

## Random String with Restrictions

```go
rand.Seed(time.Now().UnixNano())
digits := "0123456789"
specials := "~=+%^*/()[]{}/!@#$?|"
all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
    "abcdefghijklmnopqrstuvwxyz" +
    digits + specials
length := 8
buf := make([]byte, length)
buf[0] = digits[rand.Intn(len(digits))]
buf[1] = specials[rand.Intn(len(specials))]
for i := 2; i < length; i++ {
    buf[i] = all[rand.Intn(len(all))]
}
rand.Shuffle(len(buf), func(i, j int) {
    buf[i], buf[j] = buf[j], buf[i]
})
str := string(buf) // E.g. "3i[g0|)z"
```
