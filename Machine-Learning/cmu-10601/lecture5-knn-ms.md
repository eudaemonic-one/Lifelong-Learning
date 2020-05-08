# Lecture 5 k-Nearest Neighbors + Model Selection

## Overfitting

### Overfitting and Underfitting

* Underfitting
  * too simple
  * unable captures the trends in the data
  * exhibits too much bias
* Overfitting
  * too complex
  * fitting the noise in the data
  * fitting random statistical fluctuations inherent in the “sample” of training data
  * does not have enough bias

### Overfitting

* Consider a hypothesis h and its
  * Error rate over training data: error_train(h)
  * True error rate over all data: error_true(h)
* We say h **overfits** the training data if
  * error_true(h) > error(h, D_train)
* Amount of overfitting
  * error_true(h)-error_train(h)

### Avoid Overfitting

* Do not grow tree beyond some **maximum depth**
* Do not split if splitting criterion is **below some threshold**
* Stop growing when the split is **not statistically signifcant**
* Grow the entire tree, then **prune**
  * Use a third validation set

### DTs in the Wild

* DTs are one of the most popular classification methods for practical applications
* DTs can be applied to a wide variety of problems including classification, regression, density estimation, etc

## K-Nearest Neighbors

* Def: **Classification**
  * D = {x^i, y^i}_{i=1}^N
  * Every i x ∈ R^M (real valued vectors of length M)
  * Every y^i ∈ {1,2,...,L}
* M = number of features
* N = number of examples = |D|
* Def: **Binary Classification**
  * Above where y^i ∈ {0,1}
  * |y| = 2
* Def:
  * Hypothesis (aka. Decision Rule)
    * for Binary Class
    * h : R^M -> {+, -}
    * Train time: learn h
    * Test time: Given x, predict y = h(x)
* Ex: 2D Binary Class (M = 2, |y| = 2)
* Linear Decision Boundary
* Nonlinear Decision Boundary

### K-Nearest Neighbor Classifier

```
# K-Nearest Neighbor Classifier

def train(D):
    store D

def predict(x):
    Assign the most common label of the nearest k points in D
```

### Distance Functions

* KNN requires a distance function
  * g: R^M × R^M -> R
* Euclidean distance
  * g(u, v) = sqrt(Σ(u_m-v_m)^2)
* Manhattan distance
  * g(u, v) = Σ|u_m-v_m|

* What is the inductive bias of KNN?
  * Similar points should have similar labels
  * Feature scale could influence classification results

## Model Selection

* model
* model parameters
* learning algorithm
* hyperparameters