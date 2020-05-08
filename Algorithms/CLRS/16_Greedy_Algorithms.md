# 16 Greedy Algorithm

A **greedy algorithm** always makes the choice that looks best at the moment.

## An activity-selection problem

    RECURSIVE-ACTIVITY-SELECTOR(s, f, i, j)
    1   m ← i + 1
    2   while m < j and sm < fi
    3       do m ← m + 1
    4   if m < j
    5       then return {am} ∪ RECURSIVE-ACTIVITY-SELECTOR(s, f, m, j)
    6       else return φ

    GREEDY-ACTIVITY-SELECTOR(s, f, i, j)
    1   n ← length[s]
    2   A ← {a1}
    3   i ← 1
    4   for i ← 2 to n
    5       do if sm ≤ fi
    6           then A ← A ∪ {am}
    7                i ← m
    8   return A

## Greedy strategy

1. Determine the optimal substructure of the problem.
2. Develop a recursive solution.
3. Prove that at any stage of the recursion, one of the optimal choices is the greedy choice. Thus, it is always safe to make the greedy choice.
4. Show that all but one of the subproblems induced by having make the greedy choice are empty.
5. Develop a recursive algorithm that implements the greedy strategy.
6. Convert the recursive algorithm to an iterative algorithm.

We design greedy algorithms according to the following sequence of steps:

1. Cast the optimization problem as one in which we make a choice and are left with one subproblem to solve.
2. Prove that there is always an optimal solution to the original problem that makes the greedy choice, so that the greedy choice is always safe.
3. Demonstrate that, having made the greedy choice, what remains is a subproblem with the property that if we combine an optimal solution to the subproblem with the greedy choice we have made, we arrive at an optimal solution to the original problem.

## Greedy-choice property

* A globally optimal solution can be arrived at by making a locally optimal choice.
* A problem exhibits **optimal substructure** if an optimal solution to the problem contains within it optimal solutions to subproblems.

In dynamic programming, we make a choice at each step, but the choice usually depends on the solutions to subproblems. In a greedy algorithm, we make whatever choice seems best at the moment and then solve the subproblem arising after the choice is made.

## Greedy versus dynamic programming

* 0-1 knapsack problem
* fractional knapsack problem

## Huffman codes

    HUFFMAN(C)
    1   n ← |C|
    2   Q ← C
    3   for i ← 1 to n - 1
    4       do allocate a new node z
    5           left[z] ← x ← EXTRACT-MIN(Q)
    6           right[z] ← y ← EXTRACT-MIN(Q)
    7           f[z] ← f[x] + f[y]
    8           INSERT(Q, z)
    9   return EXTRACT-MIN(Q)   ▷ Return the root the tree
