# Variadic Functions (...T)

## Basic

* If the **last parameter** of a function has type `...T` it can be called with **any number** of trailing arguments of type `T`
* The actual type of `...T` inside the function is `[]T`

```go
func Sum(nums ...int) int {
    res := 0
    for _, n := range nums {
        res += n
    }
    return res
}

func main()
    fmt.Println(Sum())        // 0
    fmt.Println(Sum(1, 2, 3)) // 6
}
```

## Pass Slice Elements to a Variadic Function

* You can pass the elements of a slice `s` directly to a variadic function using the `s...` notation
* In this case no new slice is created

```go
primes := []int{2, 3, 5, 7}
fmt.Println(Sum(primes...)) // 17
```

## Append is Variadic

* The built-in append function is variadic and can be used to append any number of elements to a slice
* As a special case, you can append a string to a byte slice

```go
var buf []byte
buf = append(buf, 'a', 'b')
buf = append(buf, "cd"...)
fmt.Println(buf) // [97 98 99 100]
```
