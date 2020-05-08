# Lecture 7 Linear Regression

## Regression

* Goal
  * Given a training dataset of pairs (x, y) where x is a vector y is a scalar
  * Learn a function (curve) y' = h(x) that best fits the training data
* Def:
  * D = {(x^i, y^i)}
  * x^i ∈ R^M (input features, independent variables)
  * y^i ∈ R (output, value, dependent variables)
* Linear functions
  * Linear function ≠ Linear decision boundary
  * General case: y = w^Tx+b
* Def: Residuals
  * vevrtical distance between observed output value y^i and predicted output value y^
  * e_i = |y^i - h(x^i)| = |y^i - (w^Tx^i+b)|
* Key idea: find the linear function that minimizes the squares of residuals for a training dataset