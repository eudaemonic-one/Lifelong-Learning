# Lecture 18 Naive Bayes + Generative vs. Discriminative

## Naive Bayes

### Model 1: Bernoulli Naive Bayes (MLE)

* Data:
  * Binary feature vectors, Binary labels
  * $x \in \{0,1\}^M$ $y \in \{0,1\}$
* Generative Story:
  * $y \sim Bernoulli(\phi)$
  * $ x_1 \sim Bernoulli(\theta_{y,1}) $
  * $ x_2 \sim Bernoulli(\theta_{y,2}) $
  * ...
  * $ x_M \sim Bernoulli(\theta_{y,M}) $
* Model:
  * $p_{Φ,θ}(x,y)$
    * $= p_{\Phi,\theta}(x_1, x_2, ..., x_M, y)$
    * $= p_\Phi(y) \prod_{m=1}^M p_\theta(x_m|y)$
    * $= [(\Phi)^y (1-\Phi)^{(1-y)} \prod_{m=1}^M (\theta_{y,m})^{x_m}(1-\theta_{y,m})^{(1-x_m)}]$
  * MLE
    * Training: Find the class-conditional MLE parameters
    * Count Variables:
      * $N_{y=1} = \Sigma_{i=1}^{N} I_A(y^{(i)}=1)$
      * $N_{y=0} = \Sigma_{i=1}^{N} I_A(y^{(i)}=0)$
      * $N_{y=0,x_m=1} = \Sigma_{i=1}^{N} I_A(y^{(i)}=0 \wedge x_m^{(i)}=1)$
    * **MLE estimators**:
      * $\Phi = \frac{N_{y=1}}{N}$
      * $\theta_{0,m} = \frac{N_{y=0,x_m=1}}{N_{y=0}}$
      * $\theta_{1,m} = \frac{N_{y=1,x_m=1}}{N_{y=1}}$
      * $\forall m \in \{1,...,M\}$

### A Shortcoming of MLE

* suppose we never observe the word “unicorn” in a real news article. 
* Now suppose we observe the word “unicorn” at test time.
* What is the posterior probability that the article was a real article?

### Model 1: Bernoulli Naive Bayes (MAP)

* Generative Story:
  * The parameters are drawn once for the entire dataset
  * $$\bold{for} \space m \in \{1,...,M\}: \\ \space \space \bold{for} \space y \in \{0,1\}:\\ \space \space \space \space \theta_{m,y} \sim Beta(\alpha, \beta) \\ \bold{for} \space i \in \{1,...,N\}: \\ \space \space y^{(i)} \sim Bernoulli(\Phi) \\ \space \space \bold{for} \space m \in \{1,...,M\}:\\ \space \space \space \space x_m^{(i)} \sim Bernoulli(\theta_{y^{(i)},m})$$
* Likelihood
  * $$\ell_{MAP}(\Phi,\theta) \\ = log[p(\Phi,\theta|\alpha,\beta)p(D|\Phi,\theta)] \\ = log[(p(\Phi|\alpha,\beta)\prod_{m=1}^M p(\theta_{0,m}|\alpha,\beta))(\prod_{i=1}^N p(x^{(i)},y^{(i)}|\Phi,\theta))]$$
* MAP Estimators
  * Take derivatives, set to zero and solve
  * $\Phi = \frac{N_{y=1}}{N}$
  * $\theta_{0,m} = \frac{(\alpha-1)+N_{y=0,x_m=1}}{(\alpha-1)+(\beta-1)+N_{y=0}}$
  * $\theta_{1,m} = \frac{(\alpha-1)+N_{y=1,x_m=1}}{(\alpha-1)+(\beta-1)+N_{y=1}}$

### Model 2: Multinomial Naive Bayes

* Support: Integer vector (word IDs) $x=[x_1,x_2,\cdots,x_M]$ where $x_m \in \{1,\cdots,K\}$ a word id
* Generative Story:
  * $$\bold{for} \space i \in \{1,...,N\}: \\ \space \space y^{(i)} \sim Bernoulli(\Phi) \\ \space \space \bold{for} \space j \in \{1,...,M_i\}:\\ \space \space \space \space x_j^{(i)} \sim Multinomial(\theta_{y^{(i)},1})$$
* Model:
  * $$p_{\Phi,\theta}(x,y) = p_\Phi(y) \prod_{k=1}^K p_{\theta_k}(x_k|y) \\ = (\Phi)^y(1-\Phi)^{(1-y)} \prod_{j=1}^{M_i} \theta_{y,x_j}$$

### Model 3: Gaussian Naive Bayes

* Support: $x \in R^K$
* Model: Product of **prior** and the event model
  * $$p(x,y) = p(x_1,\cdots,x_K,y) = p(y)\prod_{k=1}^K p(x_k|y)$$
  * Gaussian Naive Bayes assumes that $p(x_k|y)$ is given by a Normal Distribution

### Model 4: Multiclass Naive Bayes

* Model:
  * The only change is that we permit y to range over C classes
  * $$p(x,y) = p(x_1,\cdots,x_K,y) = p(y)\prod_{k=1}^K p(x_k|y)$$
  * Now, $y \sim Multinomial(\Phi,1)$ and we have a separate conditional distribution $p(x_k|y)$ for each of the C classes

### Generic Naive Bayes Model

* Support: Depends on the choice of event model, $P(X_k|Y)$
* Model: Product of **prior** and the event model
  * $P(X,Y) = P(Y) \prod_{k=1}^K P(X_k|Y)$
* Training: Find the **class-conditional** MLE parameters
  * For P(Y), we find the MLE using all the data. For each $P(X_k|Y)$ we condition on the data with the corresponding
* Classification: Find the class that maximizes the posterior
  * $\hat{y} = argmax_y p(y|x) (posterior) \\ = argmax_y \frac{p(x|y)p(y)}{p(x)} (by \space Bayes's \space rule) \\ = argmax_y p(x|y)p(y)$

## Discriminative and Generative Classifiers

* **Generative Classifiers**
  * Ex: Naive Bayes
  * Define a joint model of the observations x and the labels y: $p(x,y)$
  * Learning maximizes (joint) likelihood
  * Use Bayes's Rule to classify based on the posterior: $p(y|x) = p(x|y)p(y)/p(x)$
* **Discriminative Classifiers**
  * Ex: Logistic Regression
  * Directly model the conditional: $p(y|x)$
  * Learning maximizes conditional likelihood

|      | Generative                                    | Discriminative                                |
| ---- | --------------------------------------------- | --------------------------------------------- |
| MLE  | $\prod_i p(x^{(i)},y^{(i)}|\theta)$           | $\prod_i p(y^{(i)}|x^{(i)},\theta)$           |
| MAP  | $p(\theta) \prod_i p(x^{(i)},y^{(i)} \theta)$ | $p(\theta) \prod_i p(y^{(i)}|x^{(i)},\theta)$ |

* Finite Sample Analysis
  * Assume that we are learning from a finite training dataset
  * If model assumptions are correct: Naive Bayes is a more efficient learner (requires fewer samples) than Logistic Regression
  * If model assumptions are incorrect: Logistic Regression has lower asymtotic error, and does better than Naive Bayes
* Learning (Parameter Estimation)
  * Naive Bayes
    * Parameters are decoupled -> Closed form solution for MLE
  * Logistic Regression
    * Parameters are coupled -> No closed form solution - must use iterative optimization techniques instead
* Learning (MAP Estimation of Parameters)
  * Bernoulli Naive Bayes
    * Parameters are probabilities -> Beta prior (usually) pushes probabilities away from zero/ one extremes
  * Logistic Regression:
    * Parameters are not probabilities -> Gaussian prior encourages parameters to be close to zero
* Naive Bayes vs. Logistic Regression
  * Naive Bayes: Features x are assumed to be conditionally independent given y. (i.e. Naive Bayes Assumption)
  * Logistic Regression: No assumptions are made about the form of the features x. They can be dependent and correlated in any fashion.

## Structured Prediction

* Most of the models we’ve seen so far were for classification
  * Given observations: $x=(x_1,x_2,\cdots,x_k)$
  * Predict a label: $y$
* Many real-world problems require structured prediction
  * Given observations: $x=(x_1,x_2,\cdots,x_k)$
  * Predict a structure: $y=(y_1,y_2,\cdots,y_J)$
* $\hat{\vec{y}} = argmax_{\vec{y}} p(\vec{y}|\vec{x})$ where $\vec{y} \in Y$ and $|Y|$ is very large
* Some classification problems benefit from **latent structure**

