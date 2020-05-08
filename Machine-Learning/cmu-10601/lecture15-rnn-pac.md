# Lecture 15 RNN + PAC Learning

## Recurrent Neural Networks (RNN)

### Time Series Data

* a prediction task with variable length input/output

### RNN

* Inputs: $\bold{x} = (x_1, x_2, \cdots, x_T), x_i \in R^I$
* Hidden units: $\bold{h} = (h_1, h_2, \cdots, h_T), h_i \in R^J$
* Outputs: $\bold{y} = (y_1, y_2, \cdots, y_T), y_i \in R^K$
* Nonlinearity: $H$
* Definition of Elman Network
  * $h_t = H(W_{xh}x_t + W_{hh}h_{t-1} + b_h)$
  * $y_t = W_{hy}h_t+ b_y$
* If T=1, then we have a standard feed-forward neural net with one hidden layer
* By unrolling the RNN through time, we can share parameters and accommodate arbitrary length input/output pairs

| y1     |      | y2     |      | y3     |      | y4     |      | y5     |
| ------ | ---- | ------ | ---- | ------ | ---- | ------ | ---- | ------ |
| ↑      |      | ↑      |      | ↑      |      | ↑      |      | ↑      |
| **h1** | →    | **h2** | →    | **h3** | →    | **h4** | →    | **h5** |
| ↑      |      | ↑      |      | ↑      |      | ↑      |      | ↑      |
| **x1** |      | **x2** |      | **x3** |      | **x4** |      | **x5** |

#### Bidirectional RNN

#### Deep RNNs

#### Deep Bidirectional RNNs

* the upper level hidden units have input from two previous layers (i.e. wider input)

### Long Short-Term Memory (LSTM)

* Standard RNNs have trouble learning long distance dependencies and LSTM combat this issue
* Motivation
  * LSTM units have a rich internal structure
  * The various “gates” determine the propagation of information and can choose to “remember” or “forget” information
  * Input gate: masks out the standard RNN inputs
  * Forget gate: masks out the previous cell
  * Cell: stores the input/forget mixture
  * Output gate: masks out the values of the next hidden

## PAC

### Two Types of Error

* True Error (aka. **expected risk**)
  * $R(h) = P_{x \sim p^*(x)}(c^*(x) \neq h(x))$
  * Always unknown
* Train Error (aka. **empirical risk**)
  * $\hat{R}(h) = P_{x \sim S}(c^*(x) \neq h(x))$
  * We can measure this on the training data

### PAC / SLR Model

* Generate instances from unknown distribution $p^*$
  * $x^{(i)} \sim p^*(x), \forall i$
* Oracle labels each instance with unknown function $c^*$
  * $y^{(i)} = c^*(x^{(i)}), \forall i$
* Learning algorithm chooses hypothesis $h \in H$ with low(est) training error $\hat{R}(h)$
  * $\hat{h} = argmin_h \hat{R}(h)$
* Goal: Choose an $h$ with low generalization error $R(h)$

### Three Hypotheses of Interest

* The **true function** $c^*$ is the one we are trying to learn and that labeled the training data:
  * $y^{(i)} = c^*(x^{(i)}), \forall i$
* The **expected risk minimizer** has lowest true error:
  * $h^* = argmin_{h \in H} R(h)$
* The **empirical risk minimizer** has lowest training error:
  * $\hat{h} = argmin_{h \in H} \hat{R}(h)$

### PAC Learning

* Can we bound $R(h)$ (Unknown) in terms of $\hat{R}h$ (known)? Yes
* PAC stands for **Probably Approximately Correct**
* PAC Learner yiedls hypothesis $h \in H$ which is approximately correct ($R(h) \approx 0$) with high probability ($P(R(h) \approx 0)\approx 1$)
* PAC Criterion
  * $Pr(|R(h)-\hat{R}(h)| \leq \epsilon) \geq 1 - \delta$
  * $\hat{R}(h)$ is based on a random sample of traning drawn from p*(x)
* Sample complexity is min number of training examples N s.t. the PAC Criterion is satisfied for $epsilon$ and $delta$
* a hypothesis $h \in H$ is consistant with the training data D if $\hat{R}(h)=0$
* For bounds:
  * Two cases for $c^*$
    * Realizable case: $c^* \in H$
    * Agnostic case: $c^* \notin H or c^* \in H$
  * Two cases for $|H|$
    * Finite: $|H| < +\infin$
    * Infinite: $|H| = +\infin$
