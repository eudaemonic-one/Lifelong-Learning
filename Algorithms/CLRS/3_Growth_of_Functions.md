# 3 Growth of Functions

## Θ-notation

For a given function $g(n)$, we denote by $Θ(g(n))$ the set of functions

> $Θ(g(n))$ = \{$f(n)$: there exist positive constants $c_1$, $c_2$, and $n_0$ such that $0 \leq c_1g(n) \leq f(n) \leq c_2g(n)$ for all $n \geq n_0$\}

We say that $g(n)$ is an **asymptotically tight bound** for $f(n)$. It bounds a funtion from above and below.

The definition of $Θ(g(n))$ requires that every member of $f(n) ∈ Θ(g(n))$ be **asymptotically nonnegative**, that is, that $f(n)$ be nonnegative whenever n is sufficiently large.

## Ο-notation

For a given function $g(n)$, we denote by $Ο(g(n))$ the set of functions

> $Ο(g(n))$ = \{$f(n)$: there exist positive constants $c$ and $n_0$ such that $0 \leq f(n) \leq cg(n)$ for all $n \geq n_0$\}

## Ω-notation

For a given function $g(n)$, we denote by $Ω(g(n))$ the set of functions

> $Ω(g(n))$ = \{$f(n)$: there exist positive constants $c$ and $n_0$ such that $0 \leq cg(n) \leq f(n)$ for all $n \geq n_0$\}

From the definitions of the asymptotic notations, it is easy to prove the following important theorem.

> Theorem 3.1
>
> For any two functions $f(n)$ and $g(n)$, we have $f(n) = Θ(g(n))$ if and only if $f(n) = Ο(g(n))$ and $f(n) = Ω(g(n))$.

## Asymptotic notation in equations and inequalities

Using asymptotic notations can help eliminate inessential detail and clutter in an equation. For example, the formula $2n^2 + 3n + 1 = 2n^2 + Θ(n)$ $means that $2n^2 + 3n + 1 = 2n^2 + f(n)$, where $f(n)$ is some function in the set $Θ(n)$.

We can interpret equations where asymptotic notations appear on the left-hand side using the rule: *No matter how the anonymous functions are chosen on the left of the equal sign, there is a way to choose the anonymous functions on the right of the equal sign to make the equation valid*.

## ο-notation

For a given function $g(n)$, we denote by $ο(g(n))$ the set of functions

> $ο(g(n))$ = \{$f(n)$: for any positive constant $c > 0$, there exists a constant $n_0 > 0$ such that $0 \leq f(n) < cg(n)$ for all $n \geq n_0$\}

## ω-notation

One way to define it is by

> $f(n) ∈ ω(g(n))$ if and only if $g(n) ∈ ο(f(n))$

For a given function g(n), we denote by ω(g(n)) the set of functions

> $ω(g(n))$ = \{$f(n)$: for any positive constant $c > 0$, there exists a constant $n_0 > 0$ such that $0 \leq cg(n) < f(n)$ for all $n \geq n_0$\}

## Comparison of functions

> **Transitivity:**
>
> * $f(n) = Θ(g(n))$ and $g(n) = Θ(h(n))$ imply $f(n) = Θ(h(n))$
> * $f(n) = Ο(g(n))$ and $g(n) = Ο(h(n))$ imply $f(n) = Ο(h(n))$
> * $f(n) = Ω(g(n))$ and $g(n) = Ω(h(n))$ imply $f(n) = Ω(h(n))$
> * $f(n) = ο(g(n))$ and $g(n) = ο(h(n))$ imply $f(n) = ο(h(n))$
> * $f(n) = ω(g(n))$ and $g(n) = ω(h(n))$ imply $f(n) = ω(h(n))$
>
> **Reflexivity:**
>
> * $f(n) = Θ(f(n))$
> * $f(n) = Ο(f(n))$
> * $f(n) = Ω(f(n))$
>
> **Symmetry:**
>
> * $f(n) = Θ(g(n))$ if and only if $g(n) = Θ(f(n))$
>
> **Transpose symmetry:**
>
> * $f(n) = Ο(g(n))$ if and only of $g(n) = Ω(f(n))$
> * $f(n) = ο(g(n))$ if and only of $g(n) = ω(f(n))$
>
> **Analogy between the asymptotic comparison of two functions f and g and the comparison of two real numbers a and b:**
>
> * $f(n) = Ο(g(n)) ≈ a \leq b$
> * $f(n) = Ω(g(n)) ≈ a \geq b$
> * $f(n) = Θ(g(n)) ≈ a = b$
> * $f(n) = ο(g(n)) ≈ a < b$
> * $f(n) = ω(g(n)) ≈ a > b$
