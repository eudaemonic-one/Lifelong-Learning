# 20 Fibonacci Heaps

Fibonacci heaps are predominantly of theoretical interest.

## Structure of fibonacci heaps

A **Fibonacci heap** is a collection of min-heap-ordered trees. Unlike trees within binomial trees, which are oredered, trees within Fibonacci heaps are rooted but unordered.

For a given Fibonacci heap $H$, we indicate by $t(H)$ the number of trees in the root list of $H$ and by $m(H)$ the number of marked nodes (The boolean-valued field $mark[x]$ indicates whether node $x$ has lost a child since the last time $x$ was made the child of another node.) in $H$. The potential of Fibonacci heap $H$ is then defined by:

> $Φ(H) = t(H) + 2m(H)$

## Mergable-heap operations

Lemma 19.1, which gives properties of binomial trees, holds for unordered binomial trees as well, but with the following variation on property 4.

> 4'. For the unordered binomial tree $U_k$, the root has degree $k$, which is greater than that of any other node. The children of the root are roots of subtrees $U_0, U_1, \cdots, U_{k-1}$ in some order.

    FIB-HEAP-INSERT(H, x)
    1   degree[x] ← 0
    2   p[x] ← NIL
    3   child[x] ← NIL
    4   left[x] ← NIL
    5   right[x] ← NIL
    6   mark[x] ← FALSE
    7   concatenate the root list containing  with root list H
    8   if min[H] = NIL or key[x] < key[min[H]]
    9      then min[H] ← x
    10  n[H] ← n[H] + 1

To determine the amortized cost of FIB-HEAP-INSERT, let $H$ be the input Fibonacci heap and $H'$ be the resulting Fibonacci heap. Then, $t(H') = t(H) + 1$ and $m(H') = m(H)$, and the increase in potential is $1$. Since the actual cost is $Ο(1)$, the amortize cost is $Ο(1) + 1 = Ο(1)$.

    FIB-HEAP-UNION(H1, H2)
    1   H ← MAKE-FIB-HEAP()
    2   min[H] ← min[H1]
    3   concatenate the root list of H2 with the root list of H
    4   if (min[H1] = NIL) or (min[H2] != NIL and min[H2] < min[H1])
    5      then min[H] ← min[H2]
    6   n[H] ← n[H1] + n[H2]
    7   free the objects H1 and H2
    8   return H

The change in potential is

> $Φ(H) - (Φ(H1) + Φ(H2)) = (t(H) + 2m(H)) - ((t(H1) + 2m(H1) + (t(H2) + m(H2)))) = 0$

because $t(H) = t(H1) + t(H2)$ and $m(H) = m(H1) + m(H2)$. The amortized cost of FIB-HEAP-UNION is therefore equal to its $Ο(1)$ actual cost.

    FIB-HEAP-EXTRACT-MIN(H)
    1   z ← min[H]
    2   if z != NIL
    3      then for each child x of z
    4           do add x to the root list of H
    5              p[x] ← NIL
    6           remove z from the root list of H
    7           if z = right[z]
    8              then min[H] ← NIL
    9              else min[H] ← right[z]
    10                  CONSOLIDATE(H)
    11          n[H] ← n[H] - 1
    12  return z

Consolidating the root list consists of repeatedly executing the following steps until every root in the root list has a distinct *degree* value.

> 1. Find two roots x and y in the root list with the same degree, where $key[x] <= key[y]$.
> 2. **Link** y to x: remove y from the root list, and make y a child of x. This operation is performed by the FIB-HEAP-LINK procedure. The field $degree[x]$ is incremented, and the mark ib t, if any, is cleared.

    CONSOLIDATE(H)
    1   for i ← 0 to D(n[H])
    2       do A[i] ← NIL
    3   for each node w in the root list of H
    4       do x ← w
    5          d ← degree[x]
    6           while A[d] != NIL
    7               do y ← A[d] ▷ Another node with the same degree as x.
    8                  if key[x] > key[y]
    9                     then exchange x <--> y
    10                 FIB-HEAP-LINK(H, y, x)
    11                 A[d] ← NIL
    12                 d ← d + 1
    13          A[d] ← x
    14  min[H] ← NIL
    15  for i ← 0 to D(n[H])
    16      do if A[i] != NIL
    17            then add A[i] to the root list of H
    18                 if min[H] is NIL or key[A[i]] < key[min[H]]
    19                    then min[H] ← A[i]

    FIB-HEAP-LINK(H, y, x)
    1   remove y from the root list of H
    2   make y a child of x, incrementing degree[x]
    3   mark[y] ← FALSE

The amortized cost of extracting the minimum node is $Ο(\lg n)$.

## Decreasing a key and deleting a node

    FIB-HEAP-DECREASE-KEY(H, x, k)
    1   if k > key[x]
    2      then error "new key is greater than current key"
    3   key[x] ← k
    4   y ← p[x]
    5   if y != NIL and key[x] < key[y]
    6      then CUT(H, x, y)
    7           CASCADING-CUT(H, y)
    8   if key[x] < key[min[H]]
    9      then min[H] ← x

    CUT(H, x, y)
    1   remove x from the child list of y, decrementing degree[y]
    2   add x to the root list of H
    3   p[x] ← NIL
    4   mark[x] ← FALSE

    CASCADING-CUT(H, y)
    1   z ← p[y]
    2   if z != NIL
    3      then if mark[y] = FALSE
    4              then mark[y] ← TRUE
    5              else CUT(H, y, z)
    6                   CASCADING-CUT(H, z)

We use the *mark* fields to obtain the desired time bounds. They record a little piece of the history of each node. Suppose that the following events have happened to node $x$:

1. at some time, x was a root,
2. then x was linked to another node,
3. then two children of x were removed by cuts.

The amortized cost of FIB-HEAP-DECREASE-KEY is at most

> $Ο(c) + 4 - c = Ο(1)$

    FIB-HEAP-DELETE(H, x)
    1   FIB-HEAP-DECREASE-KEY(H, x, -∞)
    2   FIB-HEAP-EXTRACT-MIN(H)

Since $D(n) = Ο(\lg n)$, the amortized time of FIB-HEAP-DELETE is $Ο(\lg n)$.

## Bounding the maximum degree

> Lemma 20.1
>
> Let $x$ be any node in a Fibonacci heap, and suppose that $degree[x] = k$. Let $y_1, y_2, \cdots, y_k$ denote the children of $x$ in the order in which they were linked to $x$, from the earliest to the latest. Then, $degree[y_1] >= 0$ and $degree[y_i] >= i - 2$ for $i = 2, 3, \cdots, k$.

> Lemma 20.2
>
> For all integers $k >= 0$,
> $F_{k+2} = 1 + \sum_{i=0}^k F_i$

> Lemma 20.3
>
> Let $x$ be any node in a Fibonacci heap, and let $k = degree[x]$. Then, $size(x) >= F_{k+2} >= Φ^k$, where $Φ = (1 + \sqrt 5)/2$.

> Corollary 20.4
>
> The maximum degree $D(n)$ of any node in an n-node Fibonacci heap is $Ο(\lg n)$.
