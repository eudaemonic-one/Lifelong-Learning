# Useful Expressions

Notes and excerpts on useful or tricky expressions.

## Array

### Avoid overcounting/duplicates in array

```python
# Very important here! We don't use `i > 0` because we always want
# to count the first element in this recursive step even if it is the same
# as one before. To avoid overcounting, we just ignore the duplicates
# after the first element.
if i > start and nums[i] == nums[i-1]:
    continue
```

```python
if nums[i] > target:
    break
```

### Rotate a 2-dimension matrix

```python
for row in board:
    do_something()
for col in zip(*board):
    do_something_else()
```

### Initiation of lists

```python
a = [0 * n] # All elements will reference the identical zero in memory
b = [0] for _ in range(n)
a != b
```

## Hashmap

### defaultdict

When key is missing, the factory function will return the default value of the corresponding type. factory_function can be list, set, or str.

```python
dict = defaultdict(factory_function)
```

## Bit Manipulation

### Reduce the least significant one

```python
i & (i-1) # will reduce the least significant one
```

### Add two numbers

```python
x, y = int(a, 2), int(b, 2)
while y:
    x, y = x ^ y, (x & y) << 1
return bin(x)[2:]
```
