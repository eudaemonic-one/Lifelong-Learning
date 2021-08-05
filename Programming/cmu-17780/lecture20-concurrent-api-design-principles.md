# Lecture 20 Concurrent API Design Principles

# Eight Concurrent API Design Principles

* Design for scalability, i.e., reduced contention

  * `LongAdder` could be emulated with `AtomicLong`
    * But adders would contend with adders
  * `ForkJoinPool` design driven by work-stealing queues
    * Which use deques to eliminate worker-stealer contention
  * Scalability also drove design of `ConcurrentHashMap`, `ConcurrentSkipListMap`, `CopyOnWriteArrayList`
  * Tradeoff: When you aim for scalability you usually weaken a guaranteed total order

* Provide alternatives to blocking

  * Blocking is usually bad
  * If possible, eliminate by magic
    * e.g., reading from `ConcurrentHashMap`
  * If not, different use cases call for different policies
    * Timeout, helping out
  * `java.util.concurrent.locks` API offers: try, try with timeout, try interruptibly, try uninteruptibly
  * Alternatives: completion-based, reactive streams
  * Tensions: green threads (loom)

* Provide easy-to-use check-then-act methods

  * Provide reduced contention and hopefully ease of use

  * Most are of the form of `tryOp` but others include `opIfCondition`, `queue.poll`, `remove`

  * Useful for nonconcurrent APIs too

  * But it doesn't really say what makes them error-prone

    * Bad return values and bad namings are suspects

  * Concurrent collections synchronize internally; you can't prevent concurrent access

    * This works for synchronized collections

    * ```java
      Map<String,String> syncMap = Collections.synchronizedMap(new HashMap<>());
      synchronized(syncMap) {
        if (!syncMap.containsKey("foo"))
          syncMap.put("foo", "bar");
      }
      ```

    * But **not** for concurrent collections

    * **Never synchronize on a concurrent collection**

  * Instead, use atomic read-modify-write methods

    * `V pubIfAbsent(K key, V value);`
    * `boolean remove(Object key, Object value); // if equal`
    * `V replace(K key, V value); // if present`
    * `boolean replace(K key, V oldValue, V newValue); // if equal`
    * `V computeIfAbsent(K key, Function<...> mappingFn);`
    * `V computeIfPresent(K key, BiFunction<...> remapFn)`;
    * `V merge(K key, V value, BiFunction<...> remapFn);`

  * ```java
    private final ConcurrentMap<T,T> map = new ConcurrentHashMap<>();
    public T intern(T t) {
      T oldVal = map.putIfAbsent(t, t);
      return oldVal == null ? t : oldVal;
    }
    ```

  * Check-then-act errors are among the most common concurrency bugs

    * Unless check-act is atomic, it's a race condition
    * Some programmers erroneously think that just by using thread-safe concurrent collections, their code is thread-safe

* Support live data

  * Some collections are continually updated & never stable
    * People still want to iterate over them
    * Requires defining weak consistency for iterators
    * This is plenty good enough
  * Somethings are surprisingly impossible
    * `ConcurrentLinkedQueue` doesn't know its own size
    * `size()` method can take arbitrary amount of time
    * Violates principle of least astonishment
    * But justifiable in this case
    * Time & space overhead to fix wound have hurt primary use case
    * And accomplished little; size not well-defined

* Support cancellation

  * Make it possible to cancel tasks, remove queued items, shut down services, etc
  * Vastly complicates design and specs
    * Do it anyway - your users need it
    * Tensioin between power and complexity
    * In this case, power wins
  * Consider support for gentle and abrupt versions
    * e.g., `ExecutorService`'s `shutdown` and `shutdownNow`

* Decouple tasks and execution

  * Greatly increases flexibility and performance
  * Special case of "don't let implementation influence interface"
  * Threads violate this principle
  * Drove design of `Executor`, `Future`, `CompletableFuture`, `Flow`
  * May seem obvious now, but a revelation at the time
  * Some newer languages, e.g., Kotlin, don't need executors because they have an `async` keyword
    * Today's APIs are tomorrow's language features

* Provide skeletal implementations

  * Some people won't be able to use abstraction
  * Make the hart work behind it available to them
  * May be **impossible** for them to do it themselves
    * Native code, VM-coupled code, etc
  * In Doug's words "Open up a dogfood store!"
  * Examples
    * `AbstractQueuedSynchronizer`
    * `AbstractOwnableSynchronizer`
    * `AbstractQueuedLongSynchronizer`

* Provide the low-level abstractions beneath your high-level abstractions

  * Some people won't be able to use abstraction family
  * Make your hard work behind it available to them too
  * Again, may be **impossible** for users to DIY
    * Native code, VM-coupled code, etc
  * Examples
    * `java.lang.reflect.VarHandle`
    * `java.util.concurrent.locks.LockSupport`

* Bonus - `java.util.concurrent` war story

  * In retrospect `CompletableFuture` should have been the same interface as `Future` - Doug Lea
  * Doug wanted a callback on `finish` in `Future`
    * But he let a collaborator talk him out of it
    * Because he is too nice
  * Now we have a separate class and interface
    * And a stopgap: Guava folks wrote `ListenableFuture`
  * Takeaways:
    * If you don't get it right the first time, API will bear scars
    * Don't be too nice; listen, but have courage of your convictions