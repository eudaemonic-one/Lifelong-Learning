# Lecture 13 Neural Network + Backpropagation

## Activation Functions

* Sigmoid / Logistic Function
  * $\frac{1}{1 + \exp{(-\alpha)})}$
* Tanh
  * Like logistic function but shifted to range $[-1, +1]$
* reLU often used in vision tasks
  * rectified linear unit
  * Linear with cutoff at zero
  * $max(0, wx+b)$
  * Soft version: $\log{(\exp{(x)} + 1)}$

## Objective Function

* Quadratic Loss
  * the same objective as Linear Regression
  * i.e. MSE
* Cross Entropy
  * the same objective as Logistic Regression
  * i.e. negative log likelihood
  * this requires probabilities, so we add an additional "softmax" layer at the end of our network
  * steeper

|               | Forward                                 | Backward                                              |
| ------------- | --------------------------------------- | ----------------------------------------------------- |
| Quadratic     | $J = 1/2 (y - y^*)^2$                   | $\frac{dJ}{dy} = y - y^*$                             |
| Cross Entropy | $J = y^*\log{(y)} + (1-y^*)\log{(1-y)}$ | $\frac{dJ}{dy} = \frac{y^*}{y} + \frac{(1-y^*)}{y-1}$ |

## Multi-class Output

* Softmax: $y_k = \frac{\exp{(b_k)}}{\sum_{l=1}^{K} \exp{(b_l)}}$
* Loss: $J = \sum_{k=1}^K y_k^* \log{(y_k)}$

## Chain Rule

* Def #1 Chain Rule
  * $y = f(u)$
  * $u = g(x)$
  * $\frac{dy}{dx} = \frac{dy}{du}·\frac{du}{dx}$
* Def #2 Chain Rule
  * $y = f(u_1,u_2)$
  * $u_2 = g_2(x)$
  * $u_1 = g_1(x)$
  * $\frac{dy}{dx} = \frac{dy}{du_1}·\frac{du_1}{dx} + \frac{dy}{du_2}·\frac{du_2}{dx}$
* Def #3 Chain Rule
  * $y = f(u)$
  * $u = g(x)$
  * $\frac{dy}{dx} = \sum_{j=1}^J \frac{dy_i}{du_j}·\frac{du_j}{dx_k}, \forall i,k$
  * Backpropagation is just repeated application of the chain rule
* Computation Graphs
  * not a Neural Network diagram

## Backpropagation

* Backprop Ex #1
  * $y = f(x,z) = \exp(xz) + \frac{xz}{\log(x)} + \frac{\sin(\log(x))}{xz}$
  * Forward Computation
    * Given $x = 2, z = 3$
    * $a = xz, b = log(x), c = sin(b), d = exp(a), e = a / b, f = c / a$
    * $y = d + e + f$
  * Backgward Computation
    * $gy = dy/dy = 1$
    * $gf = dy/df = 1, de = dy/dc = 1, gd = dy/gd = 1$
    * $gc = dy/dc = dy/df·df/dc = gf(1/a)$
    * $gb = dy/db = dy/de·de/db + dy/dc·dc/db = (ge)(-a/b^2) + (gc)(cos(b))$
    * $ga = dy/da =  dy/dc·de/da + dy/dd·dd/da + dy/df·df/da = (ge)(1/b) + (gd)(exp(a)) + (gf)(-c/a^2)$
    * $gx = (ga)(z) + (gb)(1/x)$
    * $g_z = (ga)(x)$
  * Updates for Backprop
    * $gx = \frac{dy}{dx} = \sum_{k=1}^K \frac{dy}{du_k}·\frac{du_k}{x} = \sum_{k=1}^K (gu_k)(\frac{du_k}{dx})$
    * Reuse forward computation in backward computation
    * Reuse backward computation within itself

## Neural Network Training

* Consider a 2-hidden layer neural nets
* parameters are $\theta = [\alpha^{(1)}, \alpha^{(2)}, \beta]$
* SGD training
  * Iterate until convergence:
    * Sample $i \in {1, \cdots, N}$
    * Compute gradient by backprop
      * $g\alpha^{(1)} = \nabla \alpha^{(1)}J^{(i)}(\theta)$
      * $g\alpha^{(2)} = \nabla \alpha^{(2)}J^{(i)}(\theta)$
      * $g\beta = \nabla \beta J^{(i)}(\theta)$
      * $J^{(i)}(\theta) = \ell(h_\theta(x^{(i)}), y^{(i)})$
    * Step opposite the gradient
      * $\alpha^{(1)} \leftarrow \alpha^{(1)} - \gamma g\alpha^{(1)}$
      * $\alpha^{(2)} \leftarrow \alpha^{(2)} - \gamma g\alpha^{(2)}$
      * $\beta \leftarrow \beta - \gamma g\beta$
* Backprop Ex #2: for neural network
  * Given: decision function $\hat{y} = hθ(x) = \sigma((\alpha^{(3)})^T)·\sigma((\alpha^{(2)})^T·\sigma((\alpha^{(1)})^T·x))$
  * loss function $J = \ell(\hat{y},y^*) = y^*\log(\hat{y}) + (1-y^*)\log(1-\hat{y})$
  * Forward
    * Given $x, \alpha^{(1)}, \alpha^{(2)}, \alpha^{(3)}, y^*$
    * $z^{(0)} = x$
    * for $i = 1, 2, 3$
    * $u^{(i)} = (\alpha^{(1)})^T·z^{(i-1)}$
    * $z^{(i)} = \sigma(u^{(i)})$
  * $\hat{y} = z^{(3)}$
  * $J = \ell(\hat{y}, y^*)$

