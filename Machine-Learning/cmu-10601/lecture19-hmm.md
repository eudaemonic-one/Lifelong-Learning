# Lecture 19 Hidden Markov Model

## Finite State Machine

### Ex: Tunnel Closure

* let $y_t$ be state of system at time t
* $p(y_t|y_{t-1},y_{t-2},\cdots,y_1) = p(y_t|y_{t-1})$
  * => $y_t ⫫ y_i | y_{t-1}$
  * => $p(y_1,\cdots,y_t) \\ = \prod_{i=1}^T p(y_t|y_{t-1},\cdots,y_1) \space (by \space Chain \space Rule) \\ = \prod_{t=1}^T p(y_t|y_{t-1}) \space (by \space 1st \space order \space Markov \space assumption)$
* 1st order Markov Model can be viewed as Finite State Machine
* A HMM provides a joint distribution over the tunnel states/travel times with an assumption of dependence between adjacent tunnel states

## From Mixture Model to HMM

* Naive Bayes: $P(X,Y) = \prod_{t=1}^T P(X_t|Y_t)p(Y_t)$

| Y1   |      | Y2   |      | Y3   |      | Y4   |      | Y5   |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| ↓    |      | ↓    |      | ↓    |      | ↓    |      | ↓    |
| X1   |      | X2   |      | X3   |      | X4   |      | X5   |

* HMM: $P(X,Y) = P(Y_1)(\prod_{t=1}^T P(X_t|Y_t))(\prod_{t=2}^T p(Y_t|Y_{t-1}))$
  * $P(X,Y|Y_0) = (\prod_{t=1}^T P(X_t|Y_t))(\prod_{t=1}^T p(Y_t|Y_{t-1}))$ where $p(y_1|y_0) = p(y_1)$

| Y0   | →    | Y1   | →    | Y2   | →    | Y3   | →    | Y4   | →    | Y5   |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
|      |      | ↓    |      | ↓    |      | ↓    |      | ↓    |      | ↓    |
|      |      | X1   |      | X2   |      | X3   |      | X4   |      | X5   |

## Supervised Learning for HMM

### MLE of Categorical Distribution

* Suppose we have a **dataset** obtained by repeatedly rolling a M-sided (weighted) die N times. That is we have data
  * $D = \{x^{(i)}\}_{i=1}^N$
  * where $x^{(i)} \in \{1,\cdots,M\}$ and $x^{(i)} \sim Categorical(\Phi)$
* A random variable is **Categorical** written $X \sim Categorical(\Phi)$ if and only if $P(X=x) = p(x;\Phi) = \Phi_x $where $x^{(i)} \in \{1,\cdots,M\}$ and $\Sigma_{m=1}^M \Phi_m = 1$
* The l**og-likelihood** of the data becomes:
  * $\ell(\Phi) = \Sigma_{i=1}^N \log{\Phi_{x^{(i)}}} s.t. \Sigma_{m=1}^M \Phi_m = 1$
* Solving thing constrained optimization problem yields the maimum likelihood estimator (MLE):
  * $\Phi_m^{MLE} = \frac{N_{x=m}}{N} = \frac{\Sigma_{i=1}^N I_A(x^{(i)}=m)}{N}$

## Hidden Markov Model

* Emission matrix, **A**, where $P(X_t=k|Y_t=j)=A_{j,k},\forall t,k$
* Transition matrix, **B**, where $P(Y_t=k|Y_{t-1}=j) = B_{j,k}, \forall t,k$
* Initial probs, **C**, where $P(Y_1=k)=C_k, \forall k$

### MLE for HMMs

* Data: $D=\{(\vec{x}^{(i)},\vec{y}^{(i)}\}_{i=1}^N$
  * N = # of sentences T = sentence length
* Likelihood:
  * $\ell(A,B,C) = \Sigma_{i=1}^N \log{p(\vec{x}^{(i)},\vec{y}^{(i)}|A,B,C)} \\ = \Sigma_{i=1}^N [\log{p(y_1^{(i)}|C)} + \Sigma_{t=2}^T \log{p(y_t^{(i)}|y_{t-1}^{(i)},B)} + \Sigma_{t=1}^T \log{p(x_t^{(i)}|y_t^{(i)},A)}]$
  * = $\Sigma_{i=1}^N initial + transition + emission$
* MLE:
  * $\hat{A},\hat{B},\hat{C} = argmax_{A,B,C} \ell(A,B,C)$
  * =>
  * $\hat{C} = argmax_C \Sigma_{i=1}^N \log{p(y_1^{(i)}|C)}$
  * $\hat{B} = argmax_B \Sigma_{i=1}^N \Sigma_{t=2}^T \log{p(y_t^{(i)}|y_{t-1}^{(i)},B)}$
  * $\hat{A} = argmax_A \Sigma_{i=1}^N \Sigma_{t=1}^T \log{p(x_t^{(i)}|y_t^{(i)},A)}$
  * can solve in closed form because each is Categorical
  * $\hat{C}_k = \frac{\#(y_1^{(i)}=k)}{N}$
  * $\hat{B}_{jk} = \frac{\#(y_t^{(i)}=k \and y_{t-1}^{(i)}=j)}{\#(y_{t-1}^{(i)}=j)}$
  * $\hat{A}_{jk} = \frac{\#(x_t^{(i)}=k \and y_t^{(i)}=j)}{\#(y_t^{(i)}=j)}$

### Supervised Learning for HMMs

* Learning an HMM decomposes into solving two independent Mixture Models
* Each can be solved in closed form

### Unsupervised Learning for HMMs

* Unlike discriminative models $p(y|x)$, generative models $p(x,y)$ can maximize the likelihood of the data $D = \{x^{(i)},x^{(2)},\cdots,x^{(N)}\}$
* This unsupervised learning setting can be achieved by finding parameters that maximize the marginal likelihood
  * Since we don't observe y, we define the marginal probability:
    * $p_\theta(x) = \sum_{y \in Y} p_\theta(x,y)$
    * The log-likelihood of the data is thus:
    * $\ell(\theta) = \log{\prod_{i=1}^N p_\theta(x^{(i)})} = \sum{i=1}^N \log{\sum{y \in Y} p_\theta(x^{(i)},y)}$
* We optimize using the Expectation-Maximization algorithm