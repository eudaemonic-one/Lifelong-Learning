# Redeclaring Variables

* You can’t redeclare a variable which has already been declared in the same block

```go
func main() {
	m := 0
	m := 1
	fmt.Println(m)
}

// ../main.go:3:4: no new variables on left side of :=
```

* However, variables can be redeclared in short multi-variable declarations where at least one new variable is introduced

```go
func main() {
	m := 0
	m, n := 1, 2
	fmt.Println(m, n)
}
```

* Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-blank variables is new
* Redeclaration does not introduce a new variable; it just assigns a new value to the original
