# 3 Dots in 4 Places

## Variadic Function Parameters

* If the **last parameter** of a function has type `...T`, it can be called with any number of trailing arguments of type `T`
* The actual type of `...T` inside the function is `[]T`

```go
func Sum(nums ...int) int {
    res := 0
    for _, n := range nums {
        res += n
    }
    return res
}
```

## Arguments to Variadic Functions

* You can pass a slice `s` directly to a variadic function if you unpack it with the `s...` notation
* In this case **no new slice is created**

```go
primes := []int{2, 3, 5, 7}
fmt.Println(Sum(primes...)) // 17
```

## Array Literals

* In an array literal, the `...` notation specifies a length equal to the number of elements in the literal

```go
stooges := [...]string{"Moe", "Larry", "Curly"} // len(stooges) == 3
```

## The go Command

* Three dots are used by the go command as a wildcard when describing package lists
* This command tests all packages in the current directory and its subdirectories

```text
$ go test ./...
```
