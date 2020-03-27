# Lecture 13 Neural Network + Backpropagation

## Activation Functions

* Sigmoid / Logistic Function
  * 1 / (1 + exp(-α))
* Tanh
  * Like logistic function but shifted to range [-1, +1]
* reLU often used in vision tasks
  * rectified linear unit
  * Linear with cutoff at zero
  * max(0, wx+b)
  * Soft version: log(exp(x) + 1)

## Objective Function

* Quadratic Loss
  * the same objective as Linear Regression
  * i.e. MSE
* Cross Entropy
  * the same objective as Logistic Regression
  * i.e. negative log likelihood
  * this requires probabilities, so we add an additional "softmax" layer at the end of our network
  * steeper

|               | Forward                         | Backward                        |
| ------------- | ------------------------------- | ------------------------------- |
| Quadratic     | J = 1/2 (y - y\*)^2             | dJ/dy = y - y\*                 |
| Cross Entropy | J = y\*log(y) + (1-y\*)log(1-y) | dJ/dy = y\*1/y + (1-y\*)1/(y-1) |

## Multi-class Output

* Softmax: y_k = exp(b_k) / Σ_{l=1}^{K} exp(b_l)

## Chain Rule

* Def #1
  * y = f(u)
  * u = g(x)
  * dy/dx = dy/du·du/dx
* Def #2
  * y = f(u_1,u_2)
  * u2 = g2(x)
  * u1 = g1(x)
  * dy/dx = dy/du1·du1/dx + dy/du2·du2/dx
* **Def #3 Chain Rule**
  * y = f(**u**)
  * **u** = g(x)
  * dy/dx = Σ^K_{k=1}dy/duk·duk/dx
  * Holds for any intermediate quantities
* Computation Graphs
  * not a Neural Network diagram

## Backpropagation

* Backprop Ex #1
  * y = f(x,z) = exp(xz) + xz/log(x) + sin(log(x))/xz
  * Forward Computation
    * Given x = 2, z = 3
    * a = xz, b = log(x), c = sin(b), d = exp(a), e = a / b, f = c / a
    * y = d + e + f
  * Backgward Computation
    * gy = dy/dy = 1
    * gf = dy/df = 1, de = dy/dc = 1, gd = dy/gd = 1
    * gc = dy/dc = dy/df·df/dc = gf(1/a)
    * gb = dy/db = dy/de·de/db + dy/dc·dc/db = (ge)(-a/b^2) + (gc)(cos(b))
    * ga = dy/da =  dy/dc·de/da + dy/dd·dd/da + dy/df·df/da = (ge)(1/b) + (gd)(exp(a)) + (gf)(-c/a^2)
    * gx = (ga)(z) + (gb)(1/x)
    * Gz = (ga)(x)
  * Updates for Backprop
    * gx = dy/dx = Σ^K\_{k=1}dy/duk·duk/x = Σ^K_{k=1}(guk)(duk/dx)
    * Reuse forward computation in backward computation
    * Reuse backward computation within itself

## Neural Network Training

* Consider a 2-hidden layer neural nets
* parameters are θ = [α^(1), α^(2), β]
* SGD training
  * Iterate until convergence:
    * Sample i ∈ {1, ..., N}
    * Compute gradient by backprop
      * gα^(1) = ▽ α^(1)J^(i)(θ)
      * gα^(2) = ▽ α^(2)J^(i)(θ)
      * gβ = ▽β J^(i)(θ)
      * J^(i)(θ) = l(hθ(x^(i)), y^(i))
    * Step opposite the gradient
      * α^(1) <- α^(1) - γgα^(1)
      * α^(2) <- α^(2) - γgα^(2)
      * β <- β - γgβ
* Backprop Ex #2: for neural network
  * Given: decision function y^ = hθ(x) = σ((α^(3))^T)·σ((α^(2))^T·σ((α^(1))^T·x))
  * loss function J = l(y^,y\*) = y\*log(y^) + (1-y*)log(1-y^)
  * Forward
    * Given x, α^(1), α^(2), α^(3), y*
    * z^(0) = x
    * for i = 1, 2, 3
    * u^(i) = (α^(1))^T·z^(i-1)
    * z^(i) = σ(u^(i))
  * y^ = z^(3)
  * J = l(y^, y*)

