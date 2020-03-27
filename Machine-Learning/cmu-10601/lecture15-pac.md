# Lecture 15 RNN + PAC Learning

## Recurrent Neural Networks (RNN)

### Time Series Data

* a prediction task with variable length input/output

### RNN

* Inputs: **x** = (x\_1, x\_2, ..., x_T), x\_i ∈ R^I
* Hidden units: **h** = (h\_1, h\_2, ..., h\_T), h\_i ∈ R^J
* Outputs: **y** = (y\_1, y\_2, ..., y\_T), y\_i ∈ R^K
* Nonlinearity: H
* Definition of Elman Network
  * ht = H(W\_{xh}x\_t + W\_{hh}h\_{t-1} + b\_h)
  * yt = W\_{hy}h\_t+ b\_y
* If T=1, then we have a standard feed-forward neural net with one hidden layer

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

* Can we bound R(h) (Unknown) in terms of R^(h) (known)? Yes
* PAC stands for Probably Approximately Correct
* PAC Learner yiedls hypothesis h ∈ H which is approximately correct (R(h)≈0) with high probability (P(R(h)≈0)≈1)
* PAC Criterion
  * Pr(|R(h)-R^(h)| <= ε) >= 1 - δ
  * R^(h) is based on a random sample of traning drawn from p*(x)
* Sample complexity is min number of training examples N s.t. the PAC Criterion is satisfied for ε and δ
* a hypothesis h ∈ H is consistant with the training data D if R^(h)=0
* For bounds:
  * Two cases for c*
    * Realizable case: c* ∈ H
    * Agnostic case: c* ∉ H or c* ∈ H
  * Two cases for |H|
    * Finite: |H| < +∞
    * Infinite: |H| = +∞
* Theorem 1: Sample Complexity
  * N >= 1/ε [ln(|H|)+ln(1/δ)] labeled examples are sufficient to ensure that with probability (1-δ) all h ∈ H with R^(h) = 0 have R(h) <= ε

