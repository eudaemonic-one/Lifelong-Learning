# Slices/Arrays Explained: Create, Index, Slice, Iterate

## Basics

* A slice doesn’t store any data, it just describes a section of an underlying array
  * When you change an element of a slice, you modify the corresponding element of its underlying array, and other slices that share the same underlying array will see the change
  * A slice can grow and shrink within the bounds of the underlying array
  * Slices are indexed in the usual way: `s[i]` accesses the `i`th element, starting from zero

## Construction

```go
var s []int                   // a nil slice
s1 := []string{"foo", "bar"}
s2 := make([]int, 2)          // same as []int{0, 0}
s3 := make([]int, 2, 4)       // same as new([4]int)[:2]
fmt.Println(len(s3), cap(s3)) // 2 4
```

* The default **zero value** of a slice is nil. The functions len, cap and append all regard nil as an empty slice with 0 capacity
* You create a slice either by a **slice literal** or a call to the `make` function, which takes the **length** and an optional **capacity** as arguments
* The built-in `len` and `cap` functions retrieve the length and capacity

## Slicing

```go
a := [...]int{0, 1, 2, 3} // an array
s := a[1:3]               // s == []int{1, 2}        cap(s) == 3
s = a[:2]                 // s == []int{0, 1}        cap(s) == 4
s = a[2:]                 // s == []int{2, 3}        cap(s) == 2
s = a[:]                  // s == []int{0, 1, 2, 3}  cap(s) == 4
```

* You can also create a slice by slicing an existing array or slice
  * A slice is formed by specifying a low bound and a high bound: `a[low:high]`
  * This selects a half-open range which includes the first element, but excludes the last
  * When you slice a slice, the indexes are relative to the slice itself, not to the backing array
  * The high bound is not bound by the slice’s length, but by it’s capacity, which means you can extend the length of the slice
  * Trying to extend beyond the capacity causes a panic

## Iteration

```go
s := []string{"Foo", "Bar"}
for i, v := range s {
    fmt.Println(i, v)
}
// 0 Foo
// 1 Bar
```

* The range expression, s, is **evaluated once** before beginning the loop
* The iteration values are assigned to the respective iteration variables, i and v, as in an **assignment statement**
* The second iteration variable is optional
* If the slice is `nil`, the number of iterations is 0

## Append and Copy

### Append

#### Append Function Basics

* With the built-in append function you can use a slice as a dynamic array
* The function appends any number of elements to the end of a slice
	* if there is enough capacity, the underlying array is reused
	* if not, a new underlying array is allocated and the data is copied over
* Append **returns the updated slice**
  * Therefore you need to store the result of an append, often in the variable holding the slice itself

```go
a := []int{}
a = append(a, 3, 4) // a == [3 4]
```

##### Why Doesn’t Append Work Every Time

```go
a := []byte("ba")

a1 := append(a, 'd')
a2 := append(a, 'g')

fmt.Println(string(a1)) // bag
fmt.Println(string(a2)) // bag
```

* If there is room for more elements, append reuses the underlying array

```go
a := []byte("ba")
fmt.Println(len(a), cap(a)) // 2 32
```

* This means that the slices `a`, `a1` and `a2` will refer to the same underlying array in our example
* To avoid this, we need to use two separate byte arrays

```go
const prefix = "ba"

a1 := append([]byte(prefix), 'd')
a2 := append([]byte(prefix), 'g')

fmt.Println(string(a1)) // bad
fmt.Println(string(a2)) // bag
```

#### Append One Slice to Another

* You can concatenate two slices using the three dots notation
* The `...` unpacks `b`

```go
a := []int{1, 2}
b := []int{11, 22}
a = append(a, b...) // a == [1 2 11 22]

a := []int{1, 2}
a = append(a, a...) // a == [1 2 1 2]
```

##### Variadic Functions (...T)

###### Basics

* If the **last parameter** of a function has type `...T` it can be called with any number of trailing arguments of type `T`

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

###### Pass Slice Elements to a Variadic Function

* You can pass the elements of a slice s directly to a variadic function using the `s...` notation
* In this case no new slice is created

```go
primes := []int{2, 3, 5, 7}
fmt.Println(Sum(primes...)) // 17
```

#### Append String to Byte Slice

```go
slice := append([]byte("Hello "), "world!"...)
```

#### Performance

* Appending a single element takes **constant amortized time**

### Copy

* The built-in copy function copies elements into a destination slice `dst` from a source slice `src`
* It returns the number of elements copied, which will be the minimum of `len(dst)` and `len(src)`
* The result does not depend on whether the arguments overlap

```go
func copy(dst, src []Type) int
copy(dst []byte, src string) int
```

#### Examples

```go
var s = make([]int, 3)
n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}]

s := []int{0, 1, 2}
n := copy(s, s[1:]) // n == 2, s == []int{1, 2, 2}

var b = make([]byte, 5)
copy(b, "Hello, world!") // b == []byte("Hello")
```

## Stacks and Queues

* The idiomatic way to implement a stack or queue in Go is to use a slice directly

### A Basic Stack (LIFO) Data Structure

* To push you use the built-in append function, and
* To pop you slice off the top element.

```go
var stack []string

stack = append(stack, "world!") // Push
stack = append(stack, "Hello ")

for len(stack) > 0 {
    n := len(stack) - 1 // Top element
    fmt.Print(stack[n])

    stack = stack[:n] // Pop
}

// Hello world!
```

#### Performance

* Appending a single element to a slice takes **constant amortized time**
* If the stack is permanent and the elements temporary, you may want to remove the top element before popping the stack to avoid memory leaks

```go
// Pop
stack[n] = "" // Erase element (write zero value)
stack = stack[:n]
```

### 2 Basic Queue (FIFO) Implementations

* To enqueue you use the built-in append function, and
* To dequeue you slice off the first element.

```go
var queue []string

queue = append(queue, "Hello ") // Enqueue
queue = append(queue, "world!")

for len(queue) > 0 {
    fmt.Print(queue[0]) // First element
    queue = queue[1:]   // Dequeue
}
```

#### Watch Out for Memory Leaks

* You may want to remove the first element before dequeuing

```go
// Dequeue
queue[0] = "" // Erase element (write zero value)
queue = queue[1:]
```

#### Linked List

* The `container/list` package implements a doubly linked list which can be used as a queue

```go
queue := list.New()

queue.PushBack("Hello ") // Enqueue
queue.PushBack("world!")

for queue.Len() > 0 {
    e := queue.Front() // First element
    fmt.Print(e.Value)

    queue.Remove(e) // Dequeue
}

// Hello world!
```
