# 2 Patterns for a do-while Loop in Go

* There is no **do-while loop** in Go
* To emulate the C/Java code

```text
do {
    work();
} while (condition);
```

* You may use a for loop in one of these two ways:

```go
for ok := true; ok; ok = condition {
    work()
}
```

## Repeat-until Loop

* To write a **repeat-until loop**

```text
repeat
    work();
until condition;
```

* simply change the condition in the code above to its complement:

```go
for ok := true; ok; ok = !condition {
    work()
}
```
