# Lecture 11 Feature Engineering + Regularization

## Feature Engineering

* Hand-crafted Features
* For NLP, refine embedding features with semantic/syntatic information
* For CV, edge detection, corner detection.

## Nonlinear Features

* aka. nonlinear basis functions
* Let input be some function of **x**
  * original input: **x** ∈ R^M
  * new input: **x**' ∈ R^M'
  * define **x**' = b(**x**) = [b1(x), b2(x), ..., bM'(x)]
* Ex:
  * polynomial bj(x) = x^j
  * radial basis function
  * sigmoid
  * log
* **For a linear model**: still a linear function of b(**x**) even though a nonlinear function of **x**
* Nonlinear features are **require no changes to the model** (i.e. just preprocessing)

### Over-fitting

* Root Mean Square (RMS) Error E_RMS = sqrt(2E(**w**\*)/N)
* more data helps prevent overfitting

## Regularization

* Occam's razor: prefer the simplest hypothesis
  * Model Selection: small number of features
  * Shrinkage: small number of important features
* **Given** objective function: J(θ)
* **Goal** is to find: θ^ = argmin J(θ) + λr(θ) = fit the data + combat overfitting
* **Key idea**: Define regularizer r(θ) such that we tradeoff between fitting the data and keeping the  model simple
* **Choose form of r(θ)**:
  * Ex: q-norm (usually p-norm) r(θ) = ||θ||q = [Σ^M_{m=1}||θm||^q]^(1/q)
  * q = 0 (L0 reg.) yields parameters that are zero values that makes no good computational solutions
  * q = 1 (L1 reg.) yields parameters that are zero values that is subdifferentiable (can use SGD)
  * q = 2 (L2 reg.) yields small values that is differentiable
* As λ increases, train error would increase, test error would decrease and then increase
* Don't regularize the Bias (Intercept) Parameter
* Whitening Data
  * whiten each feature by subtracting its mean and dividing by its variance
  * for regularization, this helps all the features be penalized in the same units
* Regularization and MAP estimation are quivalent for appropriately chosen priors