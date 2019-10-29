# 18 B-Trees

B-trees are balanced search trees designed to work well on magnetic disks or other direct-access secondary storage devices. B-trees are similar to red-black trees, but they are better at minimizing disk I/O operations.

## Definition of B-trees

A **B-tree** T is a rooted tree having the following properties:

> 1. Every node x has the following fields:
>   a. $n[x]$, the number of keys currently stored in node $x$,
>   b. the $n[x]$ keys themselves, stored in nondecreasing order,
>   c. $leaf[x]$, a boolean value that is TRUE if $x$ is a leaf and FALSE if $x$ is an internal node.
> 2. Each internal nodes $x$ also contains $n[x]+1$ pointers $c_1[x],c_2[x],\cdots,c_{n[x]+1}[x]$ to its children. Leaf nodes have no children, so their $c_i$ fields are undefined.
> 3. The keys $key_i[x]$ separate the ranges of keys stored in each subtree: if $k_i$ is any key stored in the subtree with root $c_i[x]$, then $k_1 \leq key_1[x] \leq k_2 \leq key_2[x] \leq \cdots \leq key_{n[x]}[x] \leq k_{n[x]+1}$.
> 4. All leaves have the same depth, which is the tree's height h.
> 5. There are lower and upper bounds on the number of keys a node can contain. These bounds can be expressed in terms of a fixed integer $t \geq 2$ called the **minimum degree** of B-tree:
>   a. Every node other than the root must have at least $t - 1$ keys. Every internal node other than the root thus has at least $t$ children. If the tree is nonempty, the root must have at least one key.
>   b. Every node can contain at most $2t - 1$ keys. Therefore, an internal node can have at most $2t$ children. We say that a node is **full** if it it contains exactly $2t - 1$ keys.

A common variant, known as a **B+-tree**, stores all the satellite information in the leaves and stores only keys and child pointers in the internal nodes, thus maximizing the branching factor of the internal nodes.

> Theorem 18.1
>
> If $n \geq 1$, then for any n-key B-tree T of height h and minimum degree $t \geq 2$.
> 
> $h \leq log_i{(n+1)/2}$

## Basic operations on B-trees

    B-TREE-SEARCH(x, k)
    1   i ← 1
    2   while i <= n[x] and k > key_i[x]
    3       do i ← i + 1
    4   if i <= n[x] and k = key_i[x]
    5       then return (x, i)
    6   if leaf[x]
    7       then return NIL
    8       else DISK-READ(c_i[x])
    9            return B-TREE-SEARCH(c_i[x], k)

    B-TREE-CREATE(T)
    1   x ← ALLOCATE-NODE()
    2   leaf[x] ← TRUE
    3   n[x] ← 0
    4   DISK-WRITE(x)
    5   root[T] ← x

    B-TREE-SPLIT-CHILD(x, i, y)
    1   z ← ALLOCATE-NODE()
    2   leaf[z] ← leaf[y]
    3   n[z] ← t - 1
    4   for j ← 1 to t - 1
    5       do key_j[z] ← key_j+t[y]
    6   if not leaf[y]
    7       then for j ← 1 to t
    8           do c_j[z] ← c_j+t[y]
    9   n[y] ← t - 1
    10  for j ← n[x] + 1 downto i + 1
    11      do c_j+t[x] ← key_j[x]
    12  c_i+1[x] ← z
    13  for j ← n[x] downto i
    14      do key_j+1[x] ← key_j[x]
    15  key_i[x] ← key_t[y]
    16  n[x] ← n[x] + 1
    17  DISK-WRITE(y)
    18  DISK-WRITE(z)
    19  DISK-WRITE(x)

    B-TREE-INSERT(T, k)
    1   r ← root[T]
    2   if n[r] = 2t - 1
    3       then s ← ALLOCATE-NODE()
    4            root[T] ← s
    5            leaf[s] ← False
    6            n[s] ← 0
    7            c_1[s] ← r
    8            B-TREE-SPLIT-CHILD(s, 1, r)
    9            B-TREE-INSERT-NONFULL(s, k)
    10      else B-TREE-INSERT-NONFULL(r, k)

    B-TREE-INSERT-NONFULL(x, k)
    1   i ← n[x]
    2   if leaf[x]
    3       then while i >= 1 and k < key_i[x]
    4               do key_i+1[x] ← key_i[x]
    5            key_i+1[x] ← k
    6            n[x] ← n[x] + 1
    7            DISK-WRITE(x)
    8       else while i >= 1 and k < key_i[x]
    9               do i ← i - 1
    10           i ← i + 1
    11           DISK-READ(c_i[x])
    12           if n[c_i[x]] = 2t - 1
    13               then B-TREE-SPLIT-CHILD(x, i, c_i[x])
    14                    if k > key_i[x]
    15                        then i ← i + 1
    16           B-TREE-INSERT-NONFULL(c_i[x], k)

