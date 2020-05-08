# 15 Dynamic Programming

## Dynamic programming

Dynamic programming is typically applied to **optimization problems**.

1. Characterize the structure of an optimal solution.
2. Recursively define the value of an optimal solution.
3. Compute the value of an optimal solution in a bottom-up fashion.
4. Construct an optimal solution from computed information.

## Assembly-line scheduling

    FASTEST-WAY(a, t, e, x, n)
    1   f1[1] ← e1 + a1,1
    2   f2[1] ← e2 + a2,1
    3   for j ← 2 to n
    4       do if f1[j-1] + a1,j ≤ f2[j-1] + t2,j-1 + a1,j
    5             then f1[j] ← f1[j-1] + a1,j
    6                  l1[j] ← 1
    7             else f1[j] ← f2[j-1] + t2,j-1 + a1,j
    8                  l1[j] ← 2
    9          if f2[j-1] + a2,j ≤ f1[j-1] + t1,j-1 + a2,j
    10            then f2[j] ← f2[j-1] + a2,j
    11                 l2[j] ← 1
    12            else f2[j] ← f1[j-1] + t1,j-1 + a2,j
    13                 l2[j] ← 2
    14   if f1[n] + x1 ≤ f2[n] + x2
    15      then f* = f1[n] + x1
    16           l* = 1
    17      else f* = f2[n] + x2
    18           l* = 2

## Matrix-chain multiplication

    MATRIX-MULTIPLY(A, B)
    1   if columns[A] ≠ rows[B]
    2      then error "incompatible dimensions"
    3      else for i ← 1 to rows[A]
    4           do for j ← 1 to columns[B]
    5               do C[i, j] ← 0
    6                  for k ← 1 to columns[A]
    7                      do C[i, j] ← C[i, j] + A[i, k]·B[k, j]

    MATRIX-CHAIN-ORDER(p)
    1   n ← length[p] - 1
    2   for i ← 1 to n
    3       do m[i, i] ← 0
    4   for l ← 2 to n
    5       do for i ← 1 to n - l + 1
    6           do j ← 1 to i + l - 1
    7              m[i, j] ← ∞
    8              for k ← i to j - 1
    9                  do q ← m[i, k] + m[k+1, j] + pi-1pkpj
    10                    if q < m[i, j]
    11                       then m[i, j] ← q
    12                            s[i, j] ← k
    13  return m and s

## Optimal substructure

You will find yourself following a common pattern in discovering optimal substructure:

1. You show that a solution to the problem consists of making a choice, such as choosing a preceding assembly-line station or choosing an index at which to split matrix chain. Making this choice leaves one or more subproblems to be solved.
2. You suppose that for a given problem, you are given the choice that leads to an optimal solution. You do not concern yourself yet with how to determine this choice. You just assume that is has been given to you.
3. Given this choice, you determine which subproblems ensue and how to best characterize the resulting space of subproblems.
4. You show that the solutions to the subproblems used within the optimal solutioon to the problem must themselves be optimal by using a "cut-and-paste" technique. You do so by supposing that each of the subproblem solutions is not optimal and then deriving a contradiction.

Optimal substructure vairs across problemm domains in two ways:

1. how many subprobles are used in an optimal solution to the original problem and
2. how many choices we have in determining which subproblem(s) to use in an optimal solution.

Informally, the running time of a dynamic-programming algorithm depends on the product of two factors: the number of subproblems and how many choices we look at for each subproblem.

Dynamic programming uses optimal substructure in a bottom-up fashion. That is, we first find optimal solutions to subproblems and, having solved the subproblems, we find an optimal solution to the problem. Finding an optimal solution to the problem entails making a choice among subproblems as to which we will use in solving the problem.

## Overlapping subproblems

When a recursive algorithm revisits the same problem over and over again, we say that the optimization problem has overlapping subproblems.

## Memoization

A memoized recursive algorithm maintains an entry in a table for the solution to each subproblem. Each table entry initially contains a special value to indicate that the entry has yet to be filled in. When the subproblem is first encountered during the executin of the recursive algorithm, its solution is computed and then stored in the table. Each subsequent time that the subproblem is encountered, the value stored in the table is simply looked up and returned.

In general practice, if all subproblems must be solved at least once, a bottom-up dynamic-programmin algorithm usually outperforms a top-down memoized algorithm by a constant factor, because there is no overhead for recursion and less overhead for maintaining the table.

## Longest common subsequence

    LCS-LENGTH(X, Y)
    1   m ← length[X]
    2   n ← length[Y]
    3   for i ← 1 to m
    4       do c[i, 0] ← 0
    5   for j ← 1 to n
    6       do c[0, j] ← 0
    7   for i ← 1 to m
    8       do for j ← 1 to n
    9              do if xi = yj
    10                   then c[i, j] ← c[i - 1, j - 1] + 1
    11                        b[i, j] ← "↖"
    12                   else if c[i - 1, j] ≥ c[i, j - 1]
    13                           then c[i, j] ← c[i - 1, j]
    14                                b[i, j] ← "↑"
    15                           else c[i, j] ← c[i, j - 1]
    16                                b[i, j] ← "←"
    17  return c and b

    PRINT-LCS(b, X, i, j)
    1   if i = 0 or j = 0
    2       then return
    3   if b[i, j] = "↖"
    4       then PRINT-LCS(b, X, i - 1, j - 1)
    5            print xi
    6   elseif b[i, j] = "↑"
    7       then PRINT-LCS(b, X, i - 1, j)
    8   else PRINT-LCS(b, X, i, j - 1)

For the b table, this procedure takes time $Ο(m+n)$.

## Optimal binary search trees

    OPTIMAL-BST(p, q, n)
    1   for i ← 1 to n + 1
    2       do e[i, i - 1] ← qi-1
    3          w[i, i - 1] ← qi-1
    4   for l ← 1 to n
    5       do for i ← 1 to n - l + 1
    6           do j ← i + l - 1
    7               e[i, j] ← ∞
    8               w[i, j] ← w[i, j - 1] + pj + qj
    9               for r ← i to j
    10                  do t ← e[i, r - 1] + e[r + 1, j] + w[i, j]
    11                      if t < e[i, j]
    12                          then e[i, j] ← t
    13                               root[i, j] ← t
    14  return e and root
