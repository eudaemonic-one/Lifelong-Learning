# Pointer vs. Value Receiver

## Basic Guidelines

* For a given type, don’t mix value and pointer receivers
* If in doubt, use pointer receivers (they are safe and extendable)

## Pointer Receivers

* You **must** use pointer receivers
  - if any method needs to mutate the receiver,
  - for structs that contain a `sync.Mutex` or similar synchronizing field (they musn’t be copied)
* You **probably want** to use pointer receivers
  - for large structs or arrays (it can be more efficient),
  - in all other cases

## Value Receivers

* You **probably want** to use value receivers
  - for `map`, `func` and `chan` types,
  - for simple basic types such as `int` or `string`,
  - for small arrays or structs that are value types, with no mutable fields and no pointers
* You **may want** to use value receivers
  - for slices with methods that do not reslice or reallocate the slice
