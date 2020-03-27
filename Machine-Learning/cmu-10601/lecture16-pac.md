# Lecture 16 PAC Learning

## Sample Complexity Results

|                | Realizable                          | Agonostic                             |
| -------------- | ----------------------------------- | ------------------------------------- |
| Finite \|H\|   | N >= (1/ε)[log(\|H\|)+log(1/δ)]     | N >= (1/(2ε^2))[log(\|H\|)+log(2/δ)]  |
| Infinite \|H\| | N = O((1/ε)[VC(H)log(1/ε)+log(1/δ)] | N = O((1/ε^2)[VC(H)log(1/ε)+log(1/δ)] |

* Bound is inversely linear in epsilon (for realizable cases)/ inversely quadratic in epsilon (for agonostic cases)
* Bound is only logarithmic in |H|

## VC Dimension

### Shattering

* H[S] – the set of splittings of dataset S using concepts (hypothesis) from H
* H shatters S if |H[S]| = 2^|S|
* |H(S)|= # of splitting of S by H < 2^|S|

### VC-dimension

* The VC-dimension of a hypothesis space H is the cardinality of (size of) the largest set S that can be shattered by H
* If arbitrarily large finite sets can be shattered by H, then VCdim(H)=∞
* To show that VC-dimension is d:
  * there exists a set of d points that can be shattered
  * there is no set of d+1 points that can be shattered
* Fact: If H is finite, then VCdim(H) <= log(|H|)
* VCdim vs Shattering
  * Proving VC Dimension requires us to show that there exists (∃) a dataset of size d that can be shattered and that there does not exist (∄) a dataset of size d+1 that can be shattered
  * Proving that a particular dataset can be shattered requires us to show that for all (∀) labelings of the dataset, our hypothesis class contains a hypothesis that can correctly classify it

## SLT-style Corollaries

* Corollary 1 (Realizable, Finite|H|)
  * For some δ > 0, with probability at least (1 - δ), for any h in H consistent with the training data (i.e. R^(h)=0)
  * $R(h) \leq \frac{1}{N}[\ln{(|H|)} + \ln{(\frac{1}{\delta})}]$
* Corollary 2 (Agnostic, Finite|H|)
  * For some δ > 0, with probability at least (1 - δ), for all hypotheses h in H
  * $R(h) \leq \hat{R}(h) + \sqrt{(\frac{1}{2N}[\ln{(|H|)} + \ln{(\frac{2}{\delta})}]}$
* Corollary 3 (Realizable, Infinite|H|)
  * For some δ > 0, with probability at least (1 - δ), for any h in H consistent with the training data (i.e. $\hat{R}(h)=0$)
  * $R(h) \leq O(\frac{1}{N}[VC(H)\ln{(\frac{N}{VC(H)})} + \ln{(\frac{1}{\delta})}])$
* Corollary 4 (Agnostic, Infinite|H|)
  * For some δ > 0, with probability at least (1 - δ), for all hypotheses h in H
  * $R(h) \leq \hat{R}(h) + O(\sqrt{(\frac{1}{N}[VC(H) + \ln{(\frac{1}{\delta})}})$

## Generalization & Overfitting Problems

* key idea: tradeoff between low train error and keeping H simple (low VCdim)
* Ex: Linear Seperable in R^M
  * VC(H) = M+1
  * How to tradeoff?
  * Use a regularizer r(**θ**) = Σ_{m=1}^M|**θ**m|
  * **θ** = argmin J(**θ**) + r(**θ**)

## Big Idea: ML Recipe

* Given data D={**x**\^(i), **y**\^(i)}\_{i=1}\^N
* Choose a decision function h**θ**(**x**) parameterized by **θ**
* Choose an objective function J_D relies on a data
* Learned by choosing parameters that optimizes the objective J(**θ**) **θ**=argmin J_D(**θ**)
* Predict on new test example **x_new** using h y=h**θ**(**x_new**)

### Decision Functions:

* Perceptron: h**θ**(**x**) = sign(**θ**^T**x**)
* Linear Regression: h**θ**(**x**) = **θ**^T**x**
* Discriminative Models: h**θ**(x) = argmax p**θ**(y|**x**)
  * Logistic Regression p(y=1|x) = σ(**θ**^T**x**)
* Generative Models: h**θ**(x) = argmax p**θ**(y|**x**)
  * Naive Bayes p(**x**,y) = p(y) Π_{m=1}^M p(x_m|y)
* Neural Network for classification: p(y|**x**) = σ(W\^(2)σ(W\^(1)^T+b^(1))+b^(2))

### Objective Function

* Maximum Likelihood Estimation (MLE): J(θ) = -Σ_{i=1}^N log(p(x^(i), y^(i)))
* Maximum Conditional Likelihood Estimation (MCLE): J(θ) = -Σ_{i=1}^N log(p(y^(i) | x^(i)))
* L2 Regularizer: J'(θ) = J(θ) + λ||θ||_2^2
* L1 Regularizer: J'(θ) = J(θ) + λ||θ||_1

### Optimization Method

* Gradient Descent
* SGD where J(θ) = (1/N) Σ J^(i)(θ)
  * θ ← θ - γ▽J^(i)(θ) for i ~ Uniform({1, ..., N})
* Mini-batch SGD
* Closed Form
  * Compute partial derivatives
  * set to zero and solve