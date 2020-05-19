# Optional Parameters, Default Parameter Values and Method Overloading

* By design, Go does **not** support
  * **optional parameters**,
  * **default parameter values**,
  * or **method overloading**
* However, there are
  * [variadic functions](./variadic-function.md) (functions that accept a variable number of arguments),
  * and **dynamic method dispatch** is supported through interfaces
* The idiomatic way to emulate optional parameters and method overloading in Go is to write several methods with different names
* For example, the `sort` package has five different functions for sorting a slice:
  * the generic `sort.Slice` and `sort.SliceStable`,
  * and the three more specific `sort.Float64s`, `sort.Ints`, and `sort.Strings`
