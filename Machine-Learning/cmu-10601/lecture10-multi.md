# Lecture 10 Binary Logistic Regression + Multinomial Logistic Regression

* Define a linear classifier (logistic regression)
* Define an objective function (likelihood)
* Optimize it with gradient descent to learn parameters
* Predict the class with highest probability under the model

## Logistic Regression Background

* h(x) = sign(θ^Tx) is indifferentiable
* Use a differentiable function instead:
  * pθ(y=1|x) = 1 / (1 + exp(-θ^Tx)) = sigmoid(θ^Tx)
  * **sigmoid = logistic(u) = 1 / (1 + e^(-u))**
* **Simple linear classification. Despite the name logistic regression.**

## Logistic Regression

* Data: Inputs are continuous vectors of length M. Outputs are discrete D = {x^(i), y^(i)}^N_i=1 where x ∈ R^M and y ∈ {0, 1}
* Model: Logistic function applied to dot product of parameters with input vector.
  * p**θ**(y=1|x)  = 1 / (1 + exp(-**θ**^Tx)) = sigmoid(**θ**^Tx)
* Learning: Finds the parameters that minimize some objective function.
  * **θ*** = argmin J(**θ**)
* Prediction: Output is the most probable class
  * y^ = argmax{y ∈ {0, 1}} p**θ**(y|x) 

## Binary Logistic Regression

* Assume: Use **x**
* Model:
  * φ = σ(**θ**^T**x**) where σ(u) = 1 / (1 + exp(-u))
  * y ~ Bernoulli(φ)
  * p(y|**x**) =
    * σ(**θ**^T**x**) if y = 1
    * 1 - σ(**θ**^T**x**) if y = 0
* Objective function:
  * l(**θ**) = Σ^N_i=1 log p(y^(i)|x(i), **θ**)
  * J(**θ**) = - 1 / N l(**θ**) = 1 / N  Σ^N_i=1 - log p(y^(i)|x(i)) negative average conditional log-likelihood for logistic regression is convex
  * **θ**MLE = argmax l(**θ**) = argmin -l(**θ**) = argmin -1/Nl(**θ**)
  * dJ^(i)(**θ**)/d(**θ**m)
    * = d/d**θ**m (- log p(y^(i)|x(i), **θ**)) =
      * d/d**θ**m -log[σ(**θ**^T**x**^(i))] if y^(i)=1
      * d/d**θ**m -log[1 - σ(**θ**^T**x**^(i))] if y^(i)= 0
    * = - (y^(i) - σ(**θ**^T**x**^(i)))x^(i)_m
    *  = - (truth - probability of y = 1) mth feature
  * ▽J^(i)(**θ**) = - (y^(i) - σ(**θ**^T**x**^(i)))x^(i)
  * ▽J(**θ**) = 1 / N Σ▽J^(i)(**θ**)
* Find θ^ by gradient descent of SGD
* Predict
  * y^
    * = argmax{y∈{0,1}} p(y|x)
    * 1 if p(y=1|x) >= 0.5
    * 0 otherwise
    * = sign(**θ**^T**x**) ∈ {0, 1}

### Maximum Conditional Likelihood Estimation

* Gradient Descent
  * take larger - more certain - steps opposite the gradient
* Stochastic Gradient Descent
  * take many small steps opposite the gradient
* Newton's Method
  * use second derivatives to better follow curvature

### Mini-Batch SGD

* Gradient Descent
  * Compute true gradient exactly from all N examples
* Stochastic Gradient Descent
  * Approximate true gradient by the gradient of one randomly chosen example
* Mini-Batch SGD
  * Approximate true gradient by the average gradient of S randomly chosen examples

## Multinomial Logistic Regression

* Ex: Polar Bears, Sea Lions, Whales
* Three hyperplanes:
  * **θ**p·**x** = 0
  * **θ**s·**x** = 0
  * **θ**w·**x** = 0
* P(y=p|x) ∝ exp(**θ**p^T**x**) / Z(**x**)
* P(y=s|x) ∝ exp(**θ**s^T**x**) / Z(**x**)
* P(y=w|x) ∝ exp(**θ**w^T**x**) / Z(**x**)
*  Z(**x**,θ) = exp(**θ**p·**x**) + exp(**θ**s·**x**) + exp(**θ**w·**x**)
* Def: Multiclass Classification
  * x ∈ R^M, y ∈ {1,2,...,k}
  * Model: p(y|x)  = exp(**θ**y·**x**) / Σ_{K ∈ Y} exp(**θ**K·**x**)
    * where **θ** = K × M matrix = [**θ**1, **θ**1, **θ**K]
    * **φ** = [φ1, φ2, ..., φK]^T
    * φK = P(y=K|**x**)
    * y ~ Categorical(**φ**)
  * Learning by MLE
    * Negative Conditional Log-likelihood
    * Compute derivatives
  * Optimization with SGD
  * Predict most probable class