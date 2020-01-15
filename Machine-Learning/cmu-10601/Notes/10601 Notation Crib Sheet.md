# 10601 Notation Crib Sheet

## Scalars, Vectors, Matrices

**Scalars** are either lowercase letters *x, y, z, α, β, γ* or uppercase Latin letters *N, M, T*. The latter are typically used to indicate a count (e.g. number of examples, features, timesteps) and are often accompanied by a corresponding index n, m, t (e.g. current example, feature, timestep).

**Vectors** are bold lowercase letters $x = [x1, x2, . . . , xM ]T$ and are typically assumed to be column vectors—hence the transposed row vector in this example. When handwritten, a vector is indicated by an over-arrow ⃗x = [x1, x2, . . . , xM ]T .

**Matrices** are bold uppercase letters, in which subscripts are used as indices into structured objects such as vectors or matrices.

## Sets

**Sets** are represented by caligraphic uppercase letters X , Y, D. We often index a set by labels in parenthesized superscripts $S = {s(1), s(2), . . . , s(S)}$, where $S = |S|$.

## Random Variables

**Random variables** are also uppercase Latin letters X, Y, Z, but their use is typically apparent from context. When a random variable Xi and a scalar xi are upper/lower-case versions of each other, we typically mean that the scalar is a value taken by the random variable.

When possible, we try to reserve Greek letters for **parameters θ, ϕ** or **hyperparameters α, β, γ**.

For a random variable X, we write $X ∼ Gaussian(µ, σ2)$ to indicate that X follows a 1D Gaussian distribution with mean µ and variance σ2. We write $x ∼ Gaussian(µ, σ2)$ to say that x is a value **sampled** from the same distribution.

A **conditional probability distribution** over random variable X given Y and Z is written $P(X|Y, Z)$ and its **probability mass function** (pmf) or **probability density function** (pdf) is $p(x|y, z)$.

The **expectation** of a random variable X is $E[X]$.

## Functions and Derivatives

Suppose we have a function $f(x)$. We write its partial derivative with respect to x as $∂f(x)/∂x$ or $df(x)/dx$. We also denote its first derivative as $f′(x)$, its second derivative as $f′′(x)$, and so on. For a multivariate function $f(x) = f(x_1, . . . , x_M)$, we write its gradient with respect to x as $∇_xf(x)$ and frequently omit the subscript, i.e. $∇f(x)$, when it is clear from context—it might not be for a gradient such as $∇_{y}g(x, y)$.

## Common Conventions

$N$ number of training examples
$M$ number of feature types
$K$ number of classes
$n$ or i current training example
$m$ current feature type
$k$ current class
$Z$ set of integers
$R$ set of reals
$R^M$ set of real-valued vectors of length M
$\{0,1\}^M$ set of binary vectors of length M
