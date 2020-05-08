# Lecture 6 Perceptron

## Model Selection

- *Def*: **model selection** is the process by which we choose the “best” model from among a set of candidates
- *Def*: **hyperparameter optimization** is the process by which we choose the “best” hyperparameters from among a set of candidates **(could be called a special case of model selection)**

## Cross Validation

### KNN Train Test Error

* Assume dataset D is 40% y^i = 0 and 60% y^i = 1
* if k = 1 train error = 0
* if k = N train error = 40% (majority vote)
* Which k to pick
  * 1200 train samples 300 test samples
    * 900 train samples and 300 validation samples
    * pick the k with the lower validation error rate
* **Cross validation**
  * is a method of estimating loss on held out data
  * Input:** training data, learning algorithm, loss function
  * **Output:** an estimate of loss function on held-out data
  * Divide data into folds
  * Concatenate all the predictions and evaluate loss (almost equivalent to averaging to loss over the folds)
* N-fold cross validation = cross validation with N folds

## Perceptron

### Linear Models for Classification

* Key idea: Try to learn hyperplane directly
* Decision function: h(x) = sign(θ^Tx) for y ∈ {-1, +1}
* y = h(x)
  * = sign(w_1x_1 + w_2x_2 + ... + w_Nx_N)
  * = sign(w^Tx+b)
* Def: a vector a is orthogonal to vector b if a·b = 0
* Def: a · b = a^T b
* Def: the l2 norm of vector u is ||u||_2 = √(Σ__{m=1}^|u| (u_m)^2)
* Def: vector project of a onto b where ||b||_2 = 1
  * c = (a·b)b
* Def: vector project of a onto b
  * c = (a · b \/ ||b||\_2) (b \/ ||b||_2) = (a · b / (||b||\_2^2)) b

### Hyperplane

* 2D line
* 3D plane
* 4D hyperplane
* Def: Hyperplane S = {x: x^Tx + b = 0}
* Def: Half space
  * S+ = {x: w^Tx + b > 0}
  * S- = {x: w^Tx + b < 0}

## Online vs. Batch Learning

* Online Learning
  * Gradually learn as each example is received
  * For i = 1, 2, 3, ...
    * **Receive** an unlabeled instance x^i
    * **Predict** y' = h_θ(x^i)
    * **Receive** trye label y^i
    * **Suffer** loss if a mistake was made, y' ≠ y^i
    * **Update** parameters θ
  * **Goal**:
    * **Minimize** the number of **mistakes**
* Batch Learning
  * Learn from all the examples at once

## Perceptron Algorithm

* Initialize parameters
  * w = [w_1, w_2, ..., w_M]^T = 0 = [0, 0, ..., 0]^T (weights)
  * b = 0 (intercept term/bias term)
* For i = 1, 2, 3, ...
  * **Receive** an unlabeled instance x^i
  * **Predict** y^ = h(x) = sign(w^T + b)
    * where sign(a)
      * = +1 if a >= 0
      * = -1 others
* **Receive** trye label y^i
* **Suffer** loss if a mistake was made, y^ ≠ y^i
  * If positive mistake (y^ ≠ y^i and y^i = +1)
    * w <- w + x^i
    * b <- b + 1
  * If negative mistake (y^ ≠ y^i and y^i = -1)
    * w <- w - x^i
    * b <- b - 1

### Hypothesis Class

* set of all hyperplanes (aka. Linear decision boundaries) in M-dimensional space where M = number of features
* h : R^M -> {+1, -1}
* H = { h(·) : exists w ∈ R^M, b ∈ R such that h(x) = sign(w^Tx+b) }