# 4 Recurrences

## The substitution method

The substitution method for solving recurrences entails two steps:

1. Guess the form of the solution
2. Use mathematical induction to find the constants and show that the solution works.

The substitution method can be used to establish either upper or lower bounds on a recurrence.

> **Making a good guess**
>
> If a recurrence is similar to one you have seen before, then guessing a similar solution is reasonable. Or, recursion trees are good helpers
>
> **Subtleties**
>
> Mathematical induction doesn't work unless we prove the exact form of the inductive hypothesis. We overcome our difficulty by subtracting a lower-order term from our previous guess. For examle, old guess is $T(n) \leq cn$, while new guess could be $T(n) \leq cn - b$, where $b \geq 0$.
>
> **Avoiding pitfalls**
>
> It is easy to err in the use of asymptotic notation.
>
> **Changing variables**
>
> Sometimes, a little algebraic manipulation can make an unknown recurrence similar to one you have seen before.
> As an example, consider the recurrence $T(n) = 2T(\lfloor \sqrt(n) \rfloor) + \lg n$ which looks difficult. We can simplify values and rename $m = \lg n$ which yields $T(n) = 2T(2^{m/2}) + m$. We can now rename $S(m) = T(2^m)$ to produce the new recurrence $S(m) = 2S(m/2) + m$. The new recurrence has the solution: $S(m) = O(m\lg m)$. Changing back from $S(m)$ to $T(n)$, we obtain $T(n) = T(2^m) = S(m) = O(m\lg m) = O(\lg n\lg {\lg n})$.

## The recursion-tree method

In a **recursion tree**, each node represents the cost of a single subproblem somewhere in the set of recursive function invocations. We sum the costs within each level of the tree to obtain a set of per-level costs, and then we sum all the per-level costs to determine the total cost of all levels of the recursion. Recursion trees are particularly useful when the recurrence describes the running time of a divide-and-conquer algorithm.

## The master method

> Theorem 4.1 (Master theorem)
> Let $a \geq 1$ and $b > 1$ be constants, let f(n) be a function, and let T(n) be defined on the nonnegative integers by the recurrence
> $T(n) = aT(n/b) + f(n)$
> where we interpret n/b to mean either $\lfloor n/b \rfloor$ or $\lceil n/b \rceil$. Then T(n) can be bounded asymptotically as follows.
>
> 1. If $f(n) = Ο(n^{\log_b a-ε})$ for some constant $ε > 0$, then $T(n) = Θ(n^{log_b a})$
> 2. If $f(n) = Θ(n^{\log_b a})$, then $T(n) = Θ(n^{\log_b a}lg n)$
> 3. If $f(n) = Ω(n^{log_b{a+ε}})$ for some constant $ε > 0$, and if $af(n/b) \leq cf(n)$ for some constant c < 1 and all sufficiently large n, then $T(n) = Θ(f(n))$
