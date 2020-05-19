# Generate Random Numbers, Characters and Slice Elements

## Go Pseudo-random Number Basics

* Use the `rand.Seed` and `rand.Int63` functions in package `math/rand` to generate a non-negative pseudo-random number of type `int64`
* Similarly, `rand.Float64` generates a pseudo-random float x, where $0 \leq x < 1$
* **Warning:** Without an initial call to rand.Seed, you will get the same sequence of numbers each time you run the program

```go
rand.Seed(time.Now().UnixNano())
n := rand.Int63() // for example 4601851300195147788

x := rand.Float64() // for example 0.49893371771268225
```

### Several Random Sources

* If needed, you can create a new random generator of type Rand with its own source, and then use its methods to generate random numbers

```go
generator := rand.New(rand.NewSource(time.Now().UnixNano()))
n := generator.Int63()
x := generator.Float64()
```

## Integers and Characters in a Given Range

### Number Between a and b

* Use `rand.Intn(m)`, which returns a pseudo-random number n, where $0 \leq n < m$

```go
n := a + rand.Intn(b-a+1) // a ≤ n ≤ b
```

### Character Between 'a' and 'z'

```go
c := 'a' + rune(rand.Intn('z'-'a'+1)) // 'a' ≤ c ≤ 'z'
```

## Random Element From Slice

* To generate a character from an arbitrary set, choose a random index from a slice of characters:

```go
chars := []rune("AB⌘")
c := chars[rand.Intn(len(chars))] // for example '⌘'
```
