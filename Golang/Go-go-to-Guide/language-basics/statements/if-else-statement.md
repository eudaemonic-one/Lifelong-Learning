# 4 Basic if-else Statement Patterns

## Basic Syntax

```go
if x > max {
    x = max
}

if x <= y {
    min = x
} else {
    min = y
}
```

## With Init Statement

```go
if x := f(); x <= y {
    return x
}
```

* The expression may be preceded by a **simple statement**, which executes before the expression is evaluated
* The **scope** of `x` is limited to the if statement

## Nested if Statements

```go
if x := f(); x < y {
    return x
} else if x > z {
    return z
} else {
    return y
}
```

* Complicated conditionals are often best expressed in Go with a **switch statement**

## Ternary ? Operator Alternatives

* You **can’t** write a short one-line conditional in Go; there is **no** ternary conditional operator
* Instead of `res = expr ? x : y`
* You write

```go
if expr {
    res = x
} else {
    res = y
}

func Min(x, y int) int {
    if x <= y {
        return x
    }
    return y
}
```
