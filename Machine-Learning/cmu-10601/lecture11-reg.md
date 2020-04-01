# Lecture 11 Feature Engineering + Regularization

## Feature Engineering

* Hand-crafted Features
* For NLP, refine embedding features with semantic/syntatic information
* For CV, edge detection, corner detection.

## Nonlinear Features

* aka. nonlinear basis functions
* Let input be some function of $x$
  * original input:$x \in R^M$
  * new input: $x' \in R^{M'}$
  * where $M' > M$ usually
  * define $x' = b(x) = [b_1(x), b_2(x), \cdots, b_M'(x)]$
* Ex:
  * polynomial $b_j(x) = x^j$
  * radial basis function $b_j(x) = \exp{(\frac{-(x-\mu_j)^2}{2\sigma_)j^2})}$
  * sigmoid $b_j(x) = \frac{1}{1+\exp{(-w_jx)}}$
  * log $b_j(x) = \log{(x)}$
* **For a linear model**: still a linear function of $b(x)$ even though a nonlinear function of $x$
* Nonlinear features are **require no changes to the model** (i.e. just preprocessing)

### Overfitting

* The problem of **overfitting** is when the model captures the noise in the training data instead of the underlying structure
  * Design Trees (e.g. when tree is too deep)
  * KNN (e.g. when k is small)
  * Perceptron (e.g. when sample isn't representative)
  * Linear Regression (e.g. with nonlinear features)
  * Logistic Regression (e.g. with many rare features)
* Root Mean Square (RMS) Error $E_{RMS} = \sqrt{(2E(w^*)/N)}$
* **more data helps prevent overfitting**

## Regularization

* Occam's razor: prefer the simplest hypothesis
  * **Model Selection**: small number of features
  * **Shrinkage**: small number of important features
* **Given** objective function: $J(\theta)$
* **Goal** is to find: $\hat{\theta} = argmin_\theta J(\theta) + \lambda r(\theta)$ = fit the data + combat overfitting
* **Key idea**: Define regularizer $r(\theta)$ such that we tradeoff between fitting the data and keeping the  model simple
* **Choose form of $r(\theta)$**:
  * Ex: q-norm (usually p-norm) $r(\theta) = ||\theta||_q = [Σ_{m=1}^M||\theta_m||^q]^{\frac{1}{q}}$

| q    | $r(\theta)$                              | yields params that are | name    | optimization notes              |
| ---- | ---------------------------------------- | ---------------------- | ------- | ------------------------------- |
| 0    | $||\theta||_0 = \sum 1(\theta_n \neq 0)$ | zero values            | L0 reg. | no good computational solutions |
| 1    | $||\theta||_1 = \sum |\theta_m|$         | zero values            | L1 reg. | subdifferentiable (can use SGD) |
| 2    | $(||\theta||)^2 = \sum \theta_m^2$       | small values           | L2 reg. | differentiable                  |

* As $\lambda$ increases, train error would increase, test error would decrease and then increase
* Don't regularize the Bias (Intercept) Parameter
  * Otherwise the learning algorithms wouldn’t be invariant to a shift in the y-values
* Whitening Data
  * whiten each feature by subtracting its mean and dividing by its variance
  * for regularization, this helps all the features be penalized in the same units

### Regularization as MAP

* L1 and L2 regularization can be interpreted as maximum a-posteriori (MAP) estimation of the parameters
* Regularization and MAP estimation are quivalent for appropriately chosen priors