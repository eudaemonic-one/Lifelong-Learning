# 2 Getting Started

## Insertion Sort

    INSERTION-SORT(A)
    1   for j ← 2 to length[A]
    2       do key ← A[j]
    3           ▷Insert A[j] into sorted sequence A[1..j-1].
    4           i ← j - 1
    5           while i > 0 and A[i] > key
    6               do A[i + 1] ← A[i]
    7                   i ← i - 1
    8           A[i + 1] ← key

## Loop invariants and the correctness of algorithms

* **Initialization:** It is true prior to the first iteration of the loop.
* **Maintenance:** It is true before an iteration of the loop, it remains true before the next iteration.
* **Termination:** When the loop terminates, the invariant gives us a usefull property that helps show that the algorithm is correct.

## Pseudocode

## Random-access Machine (RAM) model

The **running time** of an algorithm on a particular input is the number of primitive operations of "steps" executed.

## Worst-case and average-case analysis

We shall usually concentrate on finding only the **worst-case running time**.

## Divide and Conquer

* **Divide** the problem into a number of subproblems.
* **Conquer** the subproblems by solving them recursively. If the subproblem sizes are small enough, however, just solve the subproblems in a straightforward manner.
* **Combine** the solutions to the subproblems into the solution for the original problem.

## Merge Sort

    MERGE(A, p, q, r)
    1   n1 ← q - p + 1
    2   n2 ← r - q
    3   create arrays L[1..n1+1] and R[1..n2+1]
    4   for i ← 1 to n1
    5       do L[i] ← A[p + i - 1]
    6   for j ← 1 to n2
    7       do R[j] ← A[q + j]
    8   L[n1 + 1] ← ∞
    9   R[n2 + 1] ← ∞
    10  i ← 1
    11  j ← 1
    12  for k ← p to r
    13      do if L[i] <= R[j]
    14            then A[k] ← L[i]
    15                 i ← i + 1
    16            else A[k] ← R[j]
    17                 j ← j + 1

    MERGE-SORT(A, p, r)
    1   if p < r
    2       then q ← ⌊(p + r)/2⌋
    3           MERGE-SORT(A, p, q)
    4           MERGE-SORT(A, q + 1, r)
    5           MERGE(A, p, q, r)

## Bubble Sort

    BUBBLE-SORT(A)
    1   for i ← 1 to length[A]
    2       do for j ← length[A] downto i + 1
    3           do if A[j] < A[j - 1]
    4               then exchange A[j] ←→ A[j - 1]
