# Lecture 21 Introduction to concurrency, part4: In the trenches of parallelism

## Policies for Thread Safety

* Thread-confined state - mutate but don't share
  * Stack-confined
    * Primitive local variables are *never* shared between threads
    * Fast and cheap
  * Unshared object references
    * The thread that creates an object must take action to share (“publish”)
  * Thread-local variables
    * Shared object with a separate value for each thread
    * Rarely needed but invaluable (e.g., for user ID or transaction ID)
* Shared read-only state - share but don't mutate
  * Immutable data is always safe to share
* Shared thread-safe - object synchronizes itself internally
  * Thread-safe objects that perform internal synchronization
  * Better use ones from `java.util.concurrent`
* Shared guarded - client synchronizes objects externally
  * Shared objects that must be locked by user
  * Can be error prone: burden is on user
  * High concurrency can be difficult to achieve
    * Lock granularity is the entire object
  * You’re generally better off avoiding guarded objects

## Parallel Prefix Sums Algorithm

```text
void iterativePrefixSums(long[] a) {
	int gap = 1;
	for (; gap < a.length; gap *= 2) {
	  parfor(int i = gap-1; i+gap < a.length; i += 2*gap) {
 			a[i+gap] = a[i] + a[i+gap];
	  }
	}
  for (; gap > 0; gap /= 2) {
    parfor(int i = gap-1; i < a.length; i += 2*gap) {
      a[i] = a[i] + ((i-gap >= 0) ? a[i-gap] : 0);
    }
  }
}
```

* Work: O(n)
* Span: O(ln n)

## Fork/join in Java

* For a thread or more to do some work
* Join the threads to obtain the result of the work
* `java.util.concurrent.ForkJoinPool` class
  * implements `ExecutorService`
  * Executes `java.util.concurrent.ForkJoinTask<V>` or `java.util.concurrent.RecursiveTask<V>` or `java.util.concurrent.RecursiveAction`

### The RecursiveAction Abstract Class

```java
public class MyActionFoo extends RecursiveAction {
  public MyActionFoo(...) {
    store the data fields we need
  }
  @Override
  public void compute() {
    if (the task is small) {
      do he work here;
      return;
    }
    invokeAll(new MyActionFoo(...), // smaller
              new MyActionFoo(...). // subtasks
              ...)
  }
}
```

