# 3 Rules for Efficient Parallel Computation

* **Dividing a large computation into work units for parallel processing is more of an art than a science**

## Rules of Thumb

* Divide the work into units that take about 100μs to 1ms to compute.
  - If the work units are too small, the administrative overhead of dividing the problem and scheduling sub-problems might be too large
  - If the units are too big, the whole computation may have to wait for a single slow work item to finish. This slowdown can happen for many reasons, such as scheduling, interrupts from other processes, and unfortunate memory layout
  - Note that the number of work units is **independent** of the number of CPUs
* Try to minimize the amount of data sharing
  - Concurrent writes can be very costly, particularly so if goroutines execute on separate CPUs
  - Sharing data for reading is often much **less of a problem**
* Strive for good locality when accessing data
  - If data can be kept in **cache memory**, data loading and storing will be dramatically faster
  - Once again, this is particularly important for **writing**
* Whatever strategies you are using, don’t forget to **benchmark** and **profile** your code

## Example

* The following example shows how to divide a costly computation and distribute it on all available CPUs
* This is the code we want to optimize

```go
type Vector []float64

// Convolve computes w = u * v, where w[k] = Σ u[i]*v[j], i + j = k.
// Precondition: len(u) > 0, len(v) > 0.
func Convolve(u, v Vector) Vector {
    n := len(u) + len(v) - 1
    w := make(Vector, n)
    for k := 0; k < n; k++ {
        w[k] = mul(u, v, k)
    }
    return w
}

// mul returns Σ u[i]*v[j], i + j = k.
func mul(u, v Vector, k int) float64 {
    var res float64
    n := min(k+1, len(u))
    j := min(k, len(v)-1)
    for i := k - j; i < n; i, j = i+1, j-1 {
        res += u[i] * v[j]
    }
    return res
}
```

* The idea is simple: identify work units of suitable size and then run each work unit in a separate goroutine
* Here is a parallel version of `Convolve`

```go
func Convolve(u, v Vector) Vector {
    n := len(u) + len(v) - 1
    w := make(Vector, n)

    // Divide w into work units that take ~100μs-1ms to compute.
    size := max(1, 1000000/n)

    var wg sync.WaitGroup
    for i, j := 0, size; i < n; i, j = j, j+size {
        if j > n {
            j = n
        }
        // These goroutines share memory, but only for reading.
        wg.Add(1)
        go func(i, j int) {
            for k := i; k < j; k++ {
                w[k] = mul(u, v, k)
            }
            wg.Done()
        }(i, j)
    }
    wg.Wait()
    return w
}
```

* When the work units have been defined, it’s often **best** to leave the scheduling to the runtime and the operating system
* However, if needed, you can tell the runtime how many goroutines you want executing code simultaneously

```go
func init() {
    numcpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}
```
