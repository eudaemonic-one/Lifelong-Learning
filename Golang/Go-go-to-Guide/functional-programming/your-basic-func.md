# Functional Programming in Go (Case Study)

* **A graph implementation based entirely on functions**

## Introduction

* This text is about the implementation of a Go tool based entirely on functions – the API contains only immutable data types, and the code is built on top of a `struct` with five `func` fields
* It’s a tool for building **virtual graphs**
* In a virtual graph no vertices or edges are stored in memory, they are instead computed as needed
* The tool is part of a larger library of generic graph algorithms:
  * the package `graph` contains the basic graph library and
  * the subpackage `graph/build` is the tool for building virtual graphs

### The Peterson Graph

* To describe this graph in a conventional graph library, you would typically need to enumerate the edges…

```text
{0, 1}, {0, 4}, {0, 5}, {1, 2}, {1, 6}, {2, 3}, {2, 7}, {3, 4},
{3, 8}, {4, 9}, {5, 7}, {5, 8}, {6, 8}, {6, 9}, and {7, 9}
```

* You get a Petersen graph if you draw a pentagon with a pentagram inside, with five spokes
* This example from the graph/build documentation corresponds to the mathematical description

```go
// Build a Petersen graph.
pentagon := build.Cycle(5)
pentagram := pentagon.Complement()
petersen := pentagon.Match(pentagram, build.AllEdges())
```

* As you can see, the `Cycle`, `Complement` and `Match` functions implement basic concepts in graph theory: we start with a **cycle graph** of order 5, compute its **complement**, and then combine these two graphs by **matching** their vertices

## A Generic Graph

* It’s also possible to define a new graph by writing a **function** that describes the **edge set** of the graph
* This code example shows how to build a directed graph containing all edges (*v*, *w*) for which *v* is odd and *w* even
* The `build.Generic` function returns a virtual graph with 10 vertices; its edge set consists of all edges (v, w), v ≠ w, for which the anonymous function returns true

```go
// Define a graph by a function.
g := build.Generic(10, func(v, w int) bool {
    // Include all edges with v odd and w even.
    return v%2 == 1 && w%2 == 0
})
```

## Implementation

```go
type Virtual struct {
    // The `order` field is, in fact, a constant function.
    // It returns the number of vertices in the graph.
    order int

    // The `edge` and `cost` functions define a weighted graph
    // without self-loops.
    //
    //  • edge(v, w) returns true whenever (v, w) belongs to
    //    the graph; the value is disregarded when v == w.
    //
    //  • cost(v, w) returns the cost of (v, w);
    //    the value is disregarded when edge(v, w) is false.
    //
    edge func(v, w int) bool
    cost func(v, w int) int64

    // The `degree` and `visit` functions can be used to improve
    // performance. They MUST BE CONSISTENT with edge and cost.
    // If not implemented, the `generic` or `generic0` implementation
    // is used instead. The `Consistent` test function should be used
    // to check compliance.
    //
    //  • degree(v) returns the outdegree of vertex v.
    //
    //  • visit(v) visits all neighbors w of v for which w ≥ a in
    //    NUMERICAL ORDER calling do(w, c) for edge (v, w) of cost c.
    //    If a call to do returns true, visit MUST ABORT the iteration
    //    and return true; if successful it should return false.
    //    Precondition: a ≥ 0.
    //
    degree func(v int) int
    visit  func(v int, a int,
        do func(w int, c int64) (skip bool)) (aborted bool)
}
```

### Cycle Graph

```go
g := generic0(n, func(v, w int) (edge bool) {
	switch v - w {
	case 1 - n, -1, 1, n - 1:
		edge = true
	}
	return
})
```

```go
// Precondition : n ≥ 3.
g.degree = func(v int) int { return 2 }
```

```go
// Precondition : n ≥ 3.
g.visit = func(v int, a int,
    do func(w int, c int64) bool) (aborted bool) {
    var w [2]int
    switch v {
    case 0:
        w = [2]int{1, n - 1}
    case n - 1:
        w = [2]int{0, n - 2}
    default:
        w = [2]int{v - 1, v + 1}
    }
    for _, w := range w {
        if w >= a && do(w, 0) {
            return true
        }
    }
    return
}
```

### Tensor Product

```go
func (g1 *Virtual) Tensor(g2 *Virtual) *Virtual {
	m, n := g1.Order(), g2.Order()

	g := generic0(m*n, func(v, w int) (edge bool) {
		v1, v2 := v/n, v%n
		w1, w2 := w/n, w%n
		return g1.Edge(v1, w1) && g2.Edge(v2, w2)
	})

	g.degree = func(v int) (deg int) {
		v1, v2 := v/n, v%n
		return g1.degree(v1) * g2.degree(v2)
	}

	g.visit = func(v int, a int,
		do func(w int, c int64) bool) (aborted bool) {
		v1, v2 := v/n, v%n
		a1, a2 := a/n, a%n
		return g1.visit(v1, a1,
			func(w1 int, c int64) (skip bool) {
			if w1 == a1 {
				return g2.visit(v2, a2,
					func(w2 int, c int64) (skip bool) {
					return do(n*w1+w2, 0)
				})
			}
			return g2.visit(v2, 0,
				func(w2 int, c int64) (skip bool) {
				return do(n*w1+w2, 0)
			})
		})
	}
	return g
}
```
