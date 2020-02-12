# Lecture 09 Stochastic Gradient Descent + Probabilistic Learning

## Stochastic Gradient Descent

Objective function J(θ) = ΣJ^(i)(θ)

```text
procedure SGD(D, θ^(0))
    θ ← θ^(0)
    while not converged do
        i ~ Uniform({1,2,...,N})
            θ ← θ - γ▽J^(i)(θ)
    return θ
```



```text
procedure SGD(D, θ^(0))
    θ ← θ^(0)
    while not converged do
        for i ∈ shuffle({1,2,...,N}) do
            θ ← θ - γ▽J^(i)(θ)
    return θ
```

* It is common to implemement SGD using sampling **without** replacement
* epoch - single pass through the training data
* For GD, only **one update** per epoch
* For SGD, **N updates** per epoch (N = # of train examples)
* SGD reduces MSE much more rapidly than GD

### SGD for Linear Regression

SGD applied to Linear Regression is called the “Least Mean Squares” algorithm

```text
procedure LMS(D, θ^(0))
    θ ← θ^(0)
    while not converged do
        for i ∈ shuffle({1,2,...,N}) do
            g ← (θ^Tx^(i) - y^(i))x^(i)
            θ ← θ - γg
    return θ
```

### GD for Linear Regression

Gradient Descent for Linear Regression repeatedly takes steps opposite the gradient of the objective function

```text
procedure GDLR(D, θ^(0))
    θ ← θ^(0)
    while not converged do
        g ← Σ(θ^Tx^(i) - y^(i))x^(i)
        θ ← θ - γg
    return θ
```

### ## Probabilistic Learning

### Bayes Classifier

* An **oracle** knows everything (e.g. usually unknown p*(y|x))
* Optimal classifier for 0/1 loss function
  * y ∈ {0,1}
  * y^ = h(x) =
    * 1 if p(y=1|x) >= p(y=0|x)
    * 0 otherwise
    * = argmax{y ∈ {0,1}} p(y|x)
* Reducible error
* Irreducible error

### Maximum Likelihood Estimation

* Choose parameters that make the data most likely
* Bad Idea #1: Bernoulli Classifier
  * Assumption:
    * Ignore **x**
  * Model: y ~ Bernoulli(φ)
  * p(y|**x**) =
    * φ if y = 1
    * 1 - φ if y = 0
  * Conditional log-likelihood
    * l(φ) = log p(D|φ) = Σlog p(y^(i)|x^(i)) = logφ + log(1-φ) + logφ + logφ = 3logφ + log(1-φ)
  * φMLE = argmax{φ ∈ {0,1}} l(φ)
  * Bayes Classifier
    * y^ = h\_φMLE(**x**) = argmax{y ∈ {0,1}} p(y|**x**,φMLE) = 1
    * Majority Vote