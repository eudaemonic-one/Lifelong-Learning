# 4 Basic Range Loop (for-each) Patterns

## Basic for-each Loop (Slice or Array)

```go
a := []string{"Foo", "Bar"}
for i, s := range a {
    fmt.Println(i, s)
}
// 0 Foo
// 1 Bar
```

* The range expression, `a`, is evaluated once before beginning the loop
* The iteration values are assigned to the respective iteration variables, `i` and `s`, as in an assignment statement
* The second iteration variable is optional
* For a nil slice, the number of iterations is 0

## String Iteration: Runes or Bytes

```go
for i, ch := range "日本語" {
    fmt.Printf("%#U starts at byte position %d\n", ch, i)
}
// U+65E5 '日' starts at byte position 0
// U+672C '本' starts at byte position 3
// U+8A9E '語' starts at byte position 6
```

* The index is the first byte of a UTF-8-encoded code point; the second value, of type `rune`, is the value of the code point
* For an invalid UTF-8 sequence, the second value will be 0xFFFD, and the iteration will advance a single byte
* To loop over individual bytes, simply use a normal for loop and string indexing:

```go
const s = "日本語"
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}
// e6 97 a5 e6 9c ac e8 aa 9e
```

## Map Iteration: Keys and Values

```go
m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
for k, v := range m {
    fmt.Println(k, v)
}
// two 2
// three 3
// one 1
```

* If a map entry that has not yet been reached is removed during iteration, this value will not be produced
* If a map entry is created during iteration, that entry may or may not be produced
* For a nil map, the number of iterations is 0

## Channel Iteration

```go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()
for n := range ch {
    fmt.Println(n)
}
// 1
// 2
// 3
```

* For channels, the iteration values are the successive values sent on the channel until closed
* For a nil channel, the range loop blocks forever
