# 19 Binomial Heaps

|Procedure|Binary heap (worst-case)|Binomial heap (worst-case)|Fibonacci heap (amortized)|
|-|-|-|-|
|MAKE-HEAP|$Θ(1)$|$Θ(1)$|$Θ(1)$|
|INSERT|$Θ(\lg n)$|$Θ(\lg n)$|$Θ(1)$|
|MINIMUM|$Θ(1)$|$Θ(\lg n)$|$Θ(1)$|
|EXRACT-MIN|$Θ(\lg n)$|$Θ(\lg n)$|$Θ(\lg n)$|
|UNION|$Θ(n)$|$Θ(\lg n)$|$Θ(1)$|
|DECREASE-KEY|$Θ(\lg n)$|$Θ(\lg n)$|$Θ(1)$|
|DELETE|$Θ(\lg n)$|$Θ(\lg n)$|$Θ(\lg n)$|

## Binomial trees and binomial heaps

> Lemma 19.1 Properties of binomial trees
> For the binomial tree $B_k$,
>
> 1. there are $2^k$ nodes,
> 2. the height of the tree is $k$,
> 3. there are exactly $C_i^k$ nodes at depth $i$ for $i = 0, 1, \cdots, k$, and
> 4. the root has degree $k$, which is greater than that of any other node; moreover if the children of the root are numbered from left to right by $k-1, k-2, \cdots, 0$, child $i$ is the root of a subtree $B_i$.

> Corollary 19.2
>
> The maximum degree of any node in an n-node binomial tree is $\lg n $.

A **binomial heap** $H$ is a set of binomial trees that satisfies the following **binomial-heap properties**.

1. Each binomial tree in $H$ obeys the **min-heap property**: the key of a node is greater than or equal to the key of its parent. We say that each such tree is **min-heap-ordered**.
2. For any nonnegative integer $k$, there is at most one binomial tree in $H$ whose root has degree $k$.

Each binomial tree within a binomial heap can be stored in the left-child, right siblin representation way. The roots of the binomial trees within a binomial heap are organized in a linked list, which we refer to as the **root list**. The degrees of the roots strictly increase as we traverse the root list.

## Operations on binomial heaps

    BINOMIAL-HEAP-MINIMUM(H)
    1   y ← NIL
    2   x ← head[H]
    3   min ← ∞
    4   while x != NIL
    5       do if key[x] < min
    6             then min ← key[x]
    7                  y ← x
    8           x ← sibling[x]
    9   return y

    BINOMIAL-LINK(y, z)
    1   p[y] ← z
    2   sibling[y] ← child[z]
    3   child[z] ← y
    4   degree[z] ← degree[z] + 1

    BINOMIAL-HEAP-UNION(H1, H2)
    1   H ← MAKE-BINOMIAL-HEAP()
    2   head[H] ← BINOMIAL-HEAP-MERGE(H1, H2)
    3   free the objects H1 and H2 but not the lists they point to
    4   if head[H] = NIL
    5       then return H
    6   prev-x ← NIL
    7   x ← head[H]
    8   next-x ← sibling[x]
    9   while next-x != NIL
    10  do if (degree[x] != degree[next-x]) or (sibling[next-x] != NIL and degree[sibling[next-x]] = degree[x])
    11        then prev-x ← x
    12             x ← next-x
    13        else if key[x] <= key[next-x]
    14                then sibling[x] ← sibling[next-x]
    15                     BINOMIAL-LINK(next-x, x)
    16                else if prev-x = NIL
    17                        then head[H] ← next-x
    18                        else sibling[prev-x] ← next-x
    19             BINOMIAL-LINK(x, next-x)
    20             x ← next-x
    21        next-x ← sibling[x]
    22  return H

    BINOMIAL-HEAP-INSERT(H, x)
    1   H' ← MAKE-BINOMIAL-HEAP()
    2   p[x] ← NIL
    3   child[x] ← NIL
    4   sibling[x] ← NIL
    5   degree[x] ← 0
    6   head[H'] ← x
    7   H ← BINOMIAL-HEAP-UNION(H, H')

    BINOMIAL-HEAP-EXTRACT-MIN(H)
    1   find the root x with the minimum key in the root list of H and remove x from the root list of H
    2   H' ← MAKE-BINOMIAL-HEAP()
    3   reverse the order of the linked list of x's children, and set HEAD[H'] to point to the head of the resulting list
    4   H ← BINOMIAL-HEAP-UNION(H', H)
    5   return x

    BINOMIAL-HEAP-DECREASE-KEY(H, x, k)
    1   if k > key[x]
    2       then error "new key is greater than current key"
    3   key[x] ← k
    4   y ← x
    5   z ← p[y]
    6   while z != NIL and key[y] < key[z]
    7       do exchange key[y] <--> key[z]
    8          ▷ If y and z have satellite fields, exchange them, too.
    9          y ← z
    10         z ← p[y]

    BINOMIAL-HEAP-DELETE(H, x)
    1   BINOMIAL-HEAP-DECREASE-KEY(H, x, -∞)
    2   BINOMIAL-HEAP-EXTRACT-MIN(H)
