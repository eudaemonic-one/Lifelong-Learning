# Lecture 14 Deep Learning

## Matrix Calculus

| Types of Derivative | scalar    | vector        | matrix        |
| ------------------- | --------- | ------------- | ------------- |
| Scalar              | dy/dx     | d**y**/dx     | d**Y**/dx     |
| vector              | dy/d**x** | d**y**/d**x** | d**Y**/d**x** |
| Matrix              | dy/d**X** | d**y**/d**X** | d**Y**/d**X** |

## Backpropagation

### Automatic Differentiation – Reverse Mode  (aka. Backpropagation)

* Forward Computation
  * Write an algorithm for evaluating the function y = f(x). The algorithm defines a directed acyclic graph, where each variable is a node (i.e. the “computation graph”)
  * Visit each node in topological order.
    * For variable ui with inputs v1,…, vN
      * a. Compute ui = gi (v1,…, vN)
      * b. Store the result at the node

### Backward Computation (Version A)

* Initialize dy/dy = 1.
* Visit each node vj in reverse topological order.
  * Let u1,…, uM denote all the nodes with vj as an input
  * Assuming that y = h(u) = h(u_1,…, u_M)
  * and u = g(v) or equivalently ui = gi (v_1,…, v_j ,…, v_N) for all i
    * a. We already know dy/dui for all i
    * b. Compute dy/dvj as below dy/dvj = Σ_{i=1}^{M}dy/du_i du_i/dv_j

### Backward Computation (Version B)

* Initialize all partial derivatives dy/duj to 0 and dy/dy = 1.
* Visit each node in reverse topological order.
  * For variable ui = gi (v1,…, vN)
    * a. We already know dy/dui
    * b. Increment dy/dvj by (dy/dui )(dui /dvj )

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

* dJ(θ)/dθi ≈ (J(θ+ε·**d**i) - J(θ-ε·**d**i)) / 2ε
* **d**i is a 1-hot vector consisting of all zeros except for the ith entry of **d**i, which has value 1
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
  * Max-pooling layer
  * Fully-connected (linear) layer
    * Suppose input is a 3D tensor x
    * Stretch out into a long vector x_hat = [x_hat_1, ..., x_hat_(C×H×W)]
    * Then standard linear layer: y = α^Tx_hat + α0
  * ReLU layer
  * Softmax

### CNN Visualizations

* Color images consist of 3 floats per pixel for RGB (red, green blue) color values
* Convolution must also be 3-dimensional

