## Lecture 12 Neural Network

## Background

* Neural Network Model
  * Independent variables
  * weights
  * Hidden Layer
  * Weights
  * Dependent variable (Prediction)
* Artificial Model
  * Neuron: node in a directed acyclic graph (DAG)
  * Weight: multiplier on each edge
  * Activation Function: nonlinear thresholding function, which allows a neuron to "fire" when the input value is sufficiently high
  * Artificial Neural Network: collection of neurons into a DAG, which define some differentiable function

## Example #1: Neural Network with 1 Hidden Layer and 2 Hidden Units

* Let σ be the activation function
* If σ is sigmoid: σ(α) = 1 / (1 + exp(-α))
* xi ∈ R
* zi ∈ (0, 1) if σ is sigmoid
* zi ∈ R more generally
* z1 = σ(α11x1 + α12x2 + α10)
* z2 = σ(α21x1 + α22x2 + α20)
* y = σ(β1z1 + β2z2 + β0) = σ(β1 σ(α11x1 + α12x2 + α10) + β2 σ(α21x1 + α22x2 + α20) + β0)
* (Each is a logistic regression model function)
* (Don't forget the intercept terms)
* y => Pr[Y=1|x1α1β1] => predict using Bayes Optimal Classifier y^ = h_αβ(x) = 1 if y > 0.5; 0 otherwise

## Example #2: 1D Face Recognition

* D = {(1+μ, 0), (3+μ, 1)}
* Is D for classification or regression? Both!
* Which line is learned by linear regression on data set? Z_B(x)
  * Z_A(x) = wAx+bA
  * Z_B(x) = wBx+bB
  * Z_C(x) = wCx+bC
* Which sigmoid is learned by logistic regression?
  * h_A(x) = σ(Z_A(x))
  * h_B(x) = σ(Z_B(x))
  * h_C(x) = σ(Z_C(x))
* What happens if increasing intercept b?
  * to z(x)? Shift up OR shift left
  * to h(x)? Shift left
  * Shift left
* Which changes in h_A(x) if increasing wA? steeper sigmoid
* What is the decision boundary for h_C(x)? the point x = 2
* What is h_E(x) = σ((h_C(x) + h_D(x))/2)
  * not σ((Z_C(x) + Z_D(x))/2)
  * h_E is the first neural network
  * decision boundary is a nonlinear function of x

## Neural Network Parameters

* nonconvex
* no unique set of parameters

## Architectures

* Number of hidden layers (depth)
* Number of units per hidden layer (width)
* Type of activation function (nonlinearity)
* Form of objective function
* How to initialize parameters

## Example #3: Arbitrart Feedward Neural Network (Matrix Form)

* Parameters
  * x1 ... xm
  * d1 ... d2
  * α ∈ R^(M×D1)
  * β ∈ R^(D1)
* Computation
  * z^(1) = σ((α^(1))^T+b^(1))
  * σ applied elementwise to the vector ((α^(1)^T)x+b^(1))
  * z^(2) = σ(()(α^(2))^T)z^(1)+b^(2))
  * y = σ(β^T z^(2) + β0)
* Fold in the intercept terms?
  * Assume x1 = 1 z1^(1) = 1 z1^(2) = 1
  * drop β0, b^(1), b^(2)
  * Caution: tricky to implement

## Building a Neural Net

* D = M
* D < M
* D > M => Feature Engineering
* Theoretical answer:
  * A neural network with 1 hidden layer is a universal function approximator
  * For any continuous function g(x), there exists a 1-hidden-layer neural net hθ(x) such that |hθ(x) - g(x| < ε for all x, assuming sigmoid activation
* Empirical answer:
  * After 2006, deep networks are easier to train than shallow networks for many problems