# Lecture 08 Linear Regression + Optimization for ML



## Optimization by Random Guessing

1. Pick a random θ
2. Evaluate J(θ)
3. Repeat steps 1 and 2 many times
4. Return θ that gives smallests J(θ)

For Linear Regression:

* objective function is Mean Squared Error (MSE) • MSE = J(w, b) = J(θ1, θ2) = 1/N Σ(y_i-θ^Tx^i)^2
* contour plot: each line labeled with MSE – lower means a better fit
* minimum corresponds to parameters (w,b) = (θ1, θ2) that best fit some training dataset

For Linear Regression:

* target function h*(x) is **unknown***
* *only have access to h*(x) through **training examples** (x(i),y(i))
* want h(x; **θ**(t)) that **best approximates** h*(x)
* **enable generalization** w/inductive bias that restricts hypothesis class to **linear functions**

## Gradient Descent

Algorithm

1. Choose an initial point θ
2. Repeat:
   1. Compute gradient g = ▽J(θ)
   2. Choose a step size γ
   3. Update θ<-θ-γg
3. Return θ when stopping criterion is met

Remarks:

* Starting point
  * θ = 0
  * θ randomly
* Stopping criterion
  * ||▽J(θ)||_2 < ε ε=10^-8
* Step size
  * fixed value γ = 0.1
  * exact line search
  * backtracking line search

## Gradient for Linear Regression

* MSE
  * J(θ) = 1/N ΣJ^i(θ) where J^i(θ) = 1/2(y^i - θ^Tx^i)^2 (1/2 doesn't affect argmin)
  * d J^i(θ) / dθ_j
    * = (y^i - θ^Tx^i) d/dθ\_j (y^i - θ^Tx^i)
    * = (y^i - θ^Tx^i) d/dθ\_j (y^i - Σθ\_m x\_m^i)
    * = -(y^i - θ^Tx^i) x\_j
  * ▽J^i(θ)
    * = [d J^i(θ) / dθ\_1, d J^i(θ) / dθ\_2, ... , d J^i(θ) / dθ\_m]^T
    * = - (y^i - θ^Tx^i) x^i (which is a scalar multiple with vector)
  * ▽J(θ)
    * = ▽(1/N ΣJ^i(θ))
    * = 1/N Σ▽J^i(θ)
    * = 1/N Σ-(y^i - θ^Tx^i)x^i

## Convexity

* Function f: R^M->R is **convex**
  * if x1 ∈ R^M x2 ∈ R^M
  * f(tx\_1+(1-t)x\_2) <= tf(x\_1) + (1-t)f(x\_2)
* Each **local minimum** is a **global minimum**
* A *nonconvex* function is **not convex**
* Each **local minimum** in nonconvex function is **not** necessarily a **global minimum**
* Each **local minimum** of a **convex** function is also a **global minimum**
* A **strictly convex** function has a **unique global minimum**

## Convexity and Linear Regression

* The **Mean Squared Error** function, which we minimize for learning the parameters of Linear Regression, **is convex**!
  * ...but in the general case it is **not strictly convex**.

## Closed Form Solution

* θ^
  * = argmin{θ}
  * = 1 / N  Σ 1/2 (y^(i)-(θ^Tx^(i)))^2
  * = (**X**^T**X**)^(-1)(**X**^T**Y**)
* **X**^T M × N
* **X** N × M
* **Y** N × 1