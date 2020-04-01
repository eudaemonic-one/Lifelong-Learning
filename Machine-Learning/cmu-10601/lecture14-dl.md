# Lecture 14 Deep Learning

## Matrix Calculus

| Types of Derivative | scalar                 | vector                        | matrix                        |
| ------------------- | ---------------------- | ----------------------------- | ----------------------------- |
| Scalar              | $\frac{dy}{dx}$        | $\frac{d\bold{y}}{dx}$        | $\frac{d\bold{Y}}{dx}$        |
| vector              | $\frac{dy}{d\bold{x}}$ | $\frac{d\bold{y}}{d\bold{x}}$ | $\frac{d\bold{Y}}{d\bold{x}}$ |
| Matrix              | $\frac{dy}{d\bold{X}}$ | $\frac{d\bold{y}}{d\bold{X}}$ | $\frac{d\bold{Y}}{d\bold{X}}$ |

## Backpropagation

### Automatic Differentiation – Reverse Mode  (aka. Backpropagation)

* Forward Computation
  * Write an algorithm for evaluating the function $y = f(x)$. The algorithm defines a directed acyclic graph, where each variable is a node (i.e. the “computation graph”)
  * Visit each node in topological order.
    * For variable ui with inputs $v_1,…, v_N$
      * a. Compute $u_i = g_i(v_1,…, v_N)$
      * b. Store the result at the node

### Backward Computation (Version A)

* Initialize $\frac{dy}{dy} = 1$.
* Visit each node vj in reverse topological order.
  * Let $u_1, \cdots, u_M$ denote all the nodes with $v_j$ as an input
  * Assuming that $y = h(u) = h(u_1,\cdots,u_M)$
  * and $u = g(v)$ or equivalently $u_i = g_i (v_1,\cdots,v_j,\cdots,v_N)$ for all i
    * a. We already know $\frac{dy}{du_i}$ for all i
    * b. Compute $\frac{dy}{dv_j}$ as below $\frac{dy}{dv_j} = \sum_{i=1}^{M} \frac{dy}{du_i} \frac{du_i}{dv_j}$

### Backward Computation (Version B)

* Initialize all partial derivatives $\frac{dy}{du_j}=0$ and $\frac{dy}{dy} = 1$.
* Visit each node in reverse topological order.
  * For variable $u_i = g_i(v_1,\cdots,v_N)$
    * a. We already know $\frac{dy}{du_i}$
    * b. Increment $\frac{dy}{dv_j}$ by $\frac{dy}{du_i} \frac{du_i}{dv_j}$

### Why is the backpropagation algorithm efficient

* Reuses computation from the forward pass in the backward pass
* Reuses partial derivatives throughout the backward pass (but only if the algorithm reuses shared computation in the forward pass)
* Key idea: partial derivatives in the backward pass should be thought of as variables stored for reuse

### SGD with Backprop

```text
procedure SGD(Training data D, test data Dt)
	Initialize parameters α, β
	for e ∈ {1, 2, ..., E} do
	  for (x,y) ∈ D do
	  	Compute neural network layers:
	  	o = object(x, a, b, z, y_hat, J) = NNForward(x, y, α, β)
	  	Compute gradients via backprop
	  	gα = ▽αJ = NNBackward(x, y, α, β, o)
	  	gβ = ▽βJ = NNBackward(x, y, α, β, o)
	  	Update parameters:
	  	α ← α - γgα
	  	β ← α - γgβ
    Evaluate training mean cross-entropy JD(α, β)
    Evaluate test mean cross-entropy JDt(α, β)
  return parameters α, β
```

## Other Approaches to Differentiation

### Finite Difference Method

* $\frac{dJ(\theta)}{d\theta_i} \approx \frac{(J(\theta+\epsilon \cdot d_i) - J(\theta-\epsilon \cdot d_i))}{2\epsilon}$
* $\bold{d}_i$ is a 1-hot vector consisting of all zeros except for the ith entry of $\bold{d}_i$, which has value 1
* Suffers from issues of floating point precision, in practice
* Typically only appropriate to use on small examples with an appropriately chosen epsilon

## Deep Learning

### Convolution

* Basic idea:
  * Pick a 3x3 matrix F of weights
  * Slide this over an image and compute the “inner product” (similarity) of F and the corresponding field of the image, and replace the pixel in the center of the field with the output of the inner product operation
* Key point:
  * Different convolutions extract different types of low-level “features” from an image
  * All that we need to vary to generate these different features is the weights of
* A **convolution matrix** is used in image processing for tasks such as edge detection, blurring, sharpening, etc.

### Downsampling

* Downsampling can use elaborately designed stride
* Downsampling by averaging used to be a common approach
  * This is a special case of convolution where the weights are fixed to a uniform distribution
* Max-pooling is another (common) form of downsampling
  * take the max value within the same range as the equivalently-sized convolution

### Convolutional Neural Network (CNN)

* Typical layers include:
  * Convolutional layer
    * Treat convolution matrix as parameters and learn them
  * Max-pooling layer
    * Another form of downsampling
    * take the max value within the same range as the equivalently-sized convolution
    * Max() is not differentiable, but subdifferentiable
  * Fully-connected (linear) layer
    * Suppose input is a 3D tensor $\vec{x} = C \times W \times D$
    * Stretch out into a long vector $\hat{x} = [\hat{x}_1, ..., \hat{x}_{(C×D×W)}]$
    * Then standard linear layer: $y = \alpha^T\hat{x} + \alpha_0$ where $\alpha \in R^{A \times B}, |\hat{x}|=A,|\hat{y}|=B$
  * ReLU layer
    * Input: $\vec{x} \in R^K$
    * Output: $\vec{y} \in R^K$
    * Forward:
      * $\vec{y} = \sigma(\vec{(x)})$
      * $\sigma(a) = max(0,a)$
    * Backward:
      * $\frac{dJ}{dx_i} = \frac{dJ}{dy_i} \frac{dy_i}{dx_i}$
      * where $\frac{dy_i}{dx_i} = 1 \space if \space x_i > 0 \space otherwise \space 0$ is sub-derivative
  * Softmax
    * Input: $\vec{x} \in R^K$
    * Output: $\vec{y} \in R^K$
    * Forward:
      * $y_i = \frac{\exp{(x_i)}}{\sum_{k=1}^K \exp{(x_k)}}$
    * Backward:
      * $\frac{dJ}{dx_j} = \sum_{i=1}^K \frac{dJ}{dy_i} \frac{dy_i}{dx_j}$
      * where $\frac{dy_i}{dx_j} = y_i(1-y_i) \space if \space i=j \space otherwise \space -y_iy_j$

### CNN Visualizations

* Color images consist of 3 floats per pixel for RGB (red, green blue) color values
* Convolution must also be 3-dimensional

