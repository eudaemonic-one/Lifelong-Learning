# Lecture 17 MLE/MAP + Naive Bayes

## Probabilistic Learning

| Function Approximation                                       | Probabilistic Learning                                       |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Previously, we assumed that our output was generated using a deterministic target function | Today, we assume that our output is sampled from a conditional probability distribution |
| x^(i) ~ p*(·)                                                | x^(i) ~ p*(·)                                                |
| y^(i) ~ c*(x^(i))                                            | y^(i) ~ p*(·\|x^(i))                                         |
| Our goal was to learn a hypothesis h(x) that best approximates c*(x) | Our goal is to learn a probability distribution p(y\|x)      |

### Probabilistics

* Discrete Random Variable X
* Probability Mass Function (pmf) p(x)
* Continuous Random Variable X
* Probability Density Function (pdf)
  * P(a<=X<=b) = ∫f(x)dx
* Cumulative Distribution Function
  * F(x) = P(X<=x) = 
    * Σp(x')
    * ∫f(x')dx'
* P(A|B) = P(A,B) / P(B)
* Beta Distribution
  * pdf: f(Φ|α,β) = (1/B(α,β)) x^(α-1) (1-x)^(β-1)
* Dirichlet Distribution
  * pdf: p(Φ|α) = (1/B(α)) ∏_{k=1}^K  Φ_k^{α_k-1}
* Expectation
  * E[X] = Σxp(x)
  * E[X] = ∫xf(x)dx
* Variance
  * Var(X) = E[(X-E[X])^2] = E[X^2] - E[X]^2
  * μ = E[X]

## Likelihood Function

* Suppose we have N samples $D = \{x^{(1)}, x^{(2)},\cdots, x^{(N)}\}$ from a random variable X
* The **likelihood function**:
  * X is **discrete** with pmf $p(x|\theta)$
    * $L(\theta) = p(x^{(1)}|\theta) p(x^{(2)}|\theta) ... p(x^{(N)}|\theta)$
  * X is **continuous** with pdf $f(x|\theta)$
    * $L(\theta) = f(x^{(1)}|\theta)f(x^{(2)}|\theta) \cdots f(x^{(N)}|\theta)$
* The **log-likelihood function**:
  * X is **discrete** with pmf $p(x|\theta)$
    * $\ell(\theta) = \log p(x^{(1)}|\theta) + \log p(x^{(2)}|\theta) + \cdots + \log p(x^{(N)}|\theta)$
  * X is **continuous** with pdf $f(x|\theta)$
    * $\ell(\theta) = \log f(x^{(1)}|\theta) + \log f(x^{(2)}|\theta) + \cdots + \log f(x^{(N)}|\theta)$
* In both cases, the **likelihood** tells us how likely one sample is relative to another
* The **joint likelihood function**:
  * X and Y are **discrete** with pdf $p(x,y|\theta)$
    * $L(\theta) = p(x^{(1)},y^{(1)}|\theta) p(x^{(2)},y^{(2)}|\theta) \cdots p(x^{(N)},y^{(N)}|\theta)$
  * X and Y are **continuous** with pdf $f(x,y|\theta)$
    * $L(\theta) = f(x^{(1)},y^{(1)}|\theta)f(x^{(2)},y^{(2)}|\theta) \cdots f(x^{(N)},y^{(N)}|\theta)$

## Maximum Likelihood Estimation (MLE)

* Suppose we have data $D = \{x^{(i)}\}_{i=1}^N$
* Choose the parameters that maximize the likelihood of the data.
  * $\theta^{MLE} = argmax_{\theta} \prod_{i=1}^N p(x^{(i)}|\theta)$

### Reciple for Closed-form MLE

* Assume data was generated i.i.d. from some model (i.e. write the generative story)
  * $x^{(i)} \sim p(x|\theta)$
* Write log-likelihood
  * $\ell(\theta) = \log p(x^{(1)}|\theta) + \cdots + \log p(x^{(N)}|\theta)$
* Compute partial derivatives (i.e. gradient)
  * $d\ell(\theta)/d\theta_1$
  * $d\ell(\theta)/d\theta_M$
* Set derivatives to zero and solve for \theta
  * $d\ell(\theta)/d\theta_m = 0$ for all $m \in \{1, \cdots, M\}$
  * $\theta^{MLE}$ = solution to system of M equations and M variables
* Compute the second derivative and check that $\ell(\theta)$ is concave down at $\theta^{MLE}$

### Example: MLE of Exponential Distribution

* Goal
  * pdf of Exponential(λ): $f(x) = \lambda e^{-\lambda x}$
  * Suppose Xi ~ Exponential(λ) for 1 <= i <= N
  * Find MLE for data $D=\{x^{(i)}\}_{i=1}^N$
* Steps
  * First write down log-likelihood of sample
    * $\ell(\lambda) = N \log(\lambda) - \lambda \Sigma x^{(i)}$
  * Compute first derivative, set to zero, solve for λ
    * $d\ell(\lambda)/d\lambda = \frac{N}{\lambda} - \Sigma_{i=1}^N x^{(i)} = 0$
    * $\lambda^{MLE} = \frac{N}{\Sigma x^{(i)}}$
  * Compute second derivative and check that it is concave down at $\lambda^{MLE}$

### Example: MLE of Bernoulli

* Model $x^{(i)} \sim Bernoulli(\Phi)$
* Log-likelihood $D = \{x^{(1)}, \cdots, x^{(N)}\}$
  * $\ell(\Phi) = \log p(D|\Phi) \\ = \log \prod_{i=1}^N p(x^{(i)})|\Phi) \\ = \log \prod_{i=1}^N \Phi^{(x^{(i)})} (1-\Phi)^{(1-x^{(i)})} \\ = N_1 \log \Phi + N_0 \log (1-\Phi)$
  * $N_1 = \# of (x^{(i)}=1) N_0 = \# of (x^{(i)}=0)$
  * $d\ell(\Phi)/d\Phi = \frac{N_1}{\Phi} - \frac{N_0}{1-\Phi}$
  * Set to zero and solve
    * $\frac{N_1}{\Phi} - \frac{N_0}{1-\Phi} = 0$ => $\Phi^{MLE} = \frac{N_1}{N_1+N_0} = \frac{N_1}{N}$

## Maximum a posteriori (MAP) Estimation

* Choose the parameters that maximize the posterior of the parameters given the data
  * $\theta^{MAP} = argmax_ {\theta} \prod_{i=1}^N p(\theta|x^{(i)}) = argmax_{\theta} \prod_{i=1}^N p(x^{(i)}|\theta)p(\theta)$
  * where $p(\theta)$ is seen as prior (usually a pdf)

### General MAP

* MLE: $p(D|\theta)$
* MAP: $p(\theta|D) = \frac{p(D|\theta)p(\theta)}{p(D)}$  <=> posterior = likelihood·prior / not a function of **\theta** (Bayes Rule)
* $$\theta^{MAP} = argmax_{\theta} p(\theta|D) \\ = argmax_{\theta} \log p(\theta|D) \\ = argmax_{\theta} \log (\frac{p(D|\theta)p(\theta)}{p(D)}) \\ = argmax_{\theta} \log (p(D|\theta)p(\theta)) \\ = \ell^{MAP}(\theta)$$

### MAP of Beta-Bernoulli

* Model:
  * $\Phi \sim Beta(\alpha,\beta)$
  * $x^{(1)} \sim Bernoulli(\Phi)$
  * $x^{(2)} \sim Bernoulli(\Phi)$
  * ...
  * $x^{(N)} \sim Bernoulli(\Phi)$
* Log-likelihood
  * $$\ell^{MAP}(\Phi) = \log [p(D|\Phi)f(\Phi|\alpha,\beta)] \\ = \log [(\Phi^{N_1}(1-\Phi)^{N_0})(\frac{1}{B(\alpha,\beta)}\Phi^{(\alpha-1)}(1-\Phi)^{\beta-1})] \\ = \log [\Phi^{(N_1+\alpha-1)}(1-\Phi)^{(N_0+\beta-1)})\frac{1}{B(\alpha,\beta)}] \\ = (N_1+\alpha-1)\log{\Phi}(N_0+\beta-1)\log{1-\Phi}-1 \\ = N_1'\log{\Phi} + N_0'\log{(1-\Phi)} - \log{(B(\alpha,\beta))}$$
  * Φ are parameters α,β are hyperparameters
  * Derivative
    * $\frac{d\ell^{MAP}(\Phi)}{d\Phi} = \frac{N_1'}{\Phi} - \frac{N_0'}{1-\Phi}$
  * Set to zero and solve
    * $\Phi^{MAP} = \frac{N_1'}{N_1'+N_0'} = \frac{N_1+\alpha-1}{N_1+\alpha-1+N_0+\beta-1}$
* Ex #1: Suppose D={8Heads, 2Tails}
  * ΦMLE = 8/10 = 0.8
  * Now if Φ ~ Beta(α=101, β=101)
    * ΦMAP = (8+101-1) / (8+101-1+2+101-1) = 108/(108+102) ≈ 0.5
  * prior are psuedo counts
  * Now if Φ ~ Beta(α=101, β=1)
    * ΦMAP =  108 / (108+2) ≈ 1.0
* Ex #2: D={108Heads, 102Tails}
  * ΦMLE = 108/(108+102)

## Naive Bayes

### MLE for Bernoulli Naive Bayes

* Data:
  * $y \in \{0,1\}$
  * $x \in \{0,1\}^M$
* Model:
  * $y \sim Bernoulli(\Phi) = p(y|\Phi)$
  * $x_1 \sim Bernoulli(\Phi_{y,1}) = p(x_1|y,\Phi)$
  * $x_M \sim Bernoulli(\Phi_{y,M}) = p(x_M|y,\Phi)$
  * $\Phi \in [0,1]$
  * $\theta = [\theta_{H1} \theta_{H2} ... \theta_{HM} \theta_{T1} \theta_{T2} ... \theta_{TM}]$
  * $$p(x_1,x_2, \cdots, x_M, y|\Phi,\theta) \\ = p(y|\Phi) p(x_1|y,\theta) p(x_2|y,\theta) \cdots p(x_M|y,\theta) \\ = p(y|\Phi) \prod_{m=1}^M p(x_m|y, \theta_{H,m}, \theta_{T,m}) \\ =\Phi^y(1-\Phi)^{(1-y)} \prod_{m=1}^M \theta_{y,m}^{X_m} (1-\theta_{y,m})^{(1-x_m)}$$
* Def: two random variable x, y are conditionally independent given random variable Z written  $X ⫫ Z$ if and only if $p(x,y|Z) = p(x|Z) p(y|Z)$
* Naive Bayes Assumption
  * The features might not be independent
  * $p(x|y) = \prod_{m=1}^M p(x_M|y)$
  * that is $x_q$ and $x_r$ are conditionally independent given y
* Log-likelihood
  * $$\ell(\Phi,\theta) \\ = \log \prod_{i=1}^N p(x^{(i)},y^{(i)}|\Phi,\theta) \\ = \Sigma_{i=1}^N [\log p(y^{(i)}|\Phi) + \Sigma_{m=1}^M \log p(x_m^{(i)}|y^{(i)},\theta)] \\ = N_{y=H}\log{\Phi} + N_{y=T}\log{1-\Phi} + \Sigma_{m=1}^M N_{x_m=1,y=H}\log{\theta_{H,M}} + N_{x_m=0,y=H}\log{1-\theta_{H,M}} + \Sigma_{m=1}^M N_{x_m=1,y=T}\log{\theta_{T,M}} + N_{x_m=0,y=T}\log{1-\theta_{T,M}}$$
* Case a) $\Phi$
  * Take partial derivatives wrt $\Phi$
  * $\frac{d\ell(\Phi,\theta)}{d\Phi} = \frac{N_{y=1}}{\Phi} - \frac{N_{y=0}}{1-\Phi}$
  * Set to zero and solve $\Phi$
  * $\Phi^{MLE} = \frac{N_{y=1}}{N_{y=1}+N_{y=0}} = \frac{N_{y=1}}{N}$
* Case b) $\theta$ elements
  * Take partial derivatives wrt $\theta_{H,M}$ (Case for y=H and feature m)
  * $\frac{d\ell(\Phi,\theta)}{d\theta_{H,M}} = \frac{N_{x_m=1,y=H}}{\theta_{H,M}} - \frac{N_{x_m=0,y=H}}{1-\Phi_{H,M}}$
  * Set to zero and solve for $\theta_{H,M}$
  * $\theta_{H,M}^{MLE} = \frac{N_{x_m=1,y=H}}{N_{x_m=1,y=H}+N_{x_m=0,y=H}} = \frac{N_{x_m=1,y=H}}{N_{y=H}}$

