# 5 Switch Statement Patterns

## Basic Switch With Default

* A `switch` statement runs the first case equal to the condition expression
* The cases are evaluated from top to bottom, stopping when a case succeeds
* If no case matches and there is a default case, its statements are executed
* **Unlike C and Java, the case expressions do not need to be constants**

```go
switch time.Now().Weekday() {
case time.Saturday:
    fmt.Println("Today is Saturday.")
case time.Sunday:
    fmt.Println("Today is Sunday.")
default:
    fmt.Println("Today is a weekday.")
}
```

## No Condition

* A `switch` without a condition is the same as switch true

```go
switch hour := time.Now().Hour(); { // missing expression means "true"
case hour < 12:
    fmt.Println("Good morning!")
case hour < 17:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}
```

## Case List

```go
func WhiteSpace(c rune) bool {
    switch c {
    case ' ', '\t', '\n', '\f', '\r':
        return true
    }
    return false
}
```

## Fallthrough

* A `fallthrough` statement transfers control to the next case
* It may be used only as the final statement in a clause

```go
switch 2 {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")
    fallthrough
case 3:
    fmt.Println("3")
}

// 2
// 3
```
