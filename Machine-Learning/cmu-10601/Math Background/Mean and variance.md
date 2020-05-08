# Mean and variance

## Mean

The *mean* of a numeric random variable (also called its *expected value* or *first moment*) is defined as


<center>E(X)=∑x⋅P(X=x)</center>

## Functions of a random variable

If X is a random variable and f is a function, then f(X) is a random variable too.

## Linearity of expectation

<center>E(aX+b) = aE(X)+b</center>

for any random variable X and any constants a,b.

## Variance

Suppose we want to know *how far* our random variable will typically be from its mean. One way to make this question precise is the *variance*, which is defined as

<center>Var(X) = E[(X-E(X))^2]</center>

A related quantity is the *standard deviation*: the square root of the variance.

The moments of X−E(X) are called the *central moments* of X. So, the variance is the second central moment of X.

## Standardizing

Suppose we have a random variable X with E(X)=μ and Var(X)=σ^2. If we define a new random variable

<center>Z = (X−μ)/σ</center>

then we have E(Z)=0 and Var(Z)=1.

We call this process *standardizing* or *z-scoring* the variable X.

## Covariance and correlation

Suppose we have two random variables, X and Y. The covariance of X and Y is defined as

<center>Cov(X,Y)=E[(X−E(X))(Y−E(Y))]</center>

The covariance is a measure of how much big values of X tend to predict big values of Y.

If we standardize X and Y before computing their covariance, the result is called the *correlation* between X and Y. In equations:

<center>Corr(X,Y)=E[(X−E(X))/σ_X (Y−E(Y))/σ_Y]</center>

where σ^2_X=Var(X) and σ^2_Y=Var(Y).

Correlation is always in the range [−1,1] a correlation of 1 means that Y is a linear function of X with positive slope, while a correlation of –1 means that Y is a linear function of X with negative slope.

if two random variables are independent, then they have zero covariance and zero correlation, but the reverse is not true.

## Conditional mean and variance

If we condition on an event A, the distribution of X can change. In this new distribution, X will have a mean and variance; these are called the *conditional* mean and variance of X given A, written E(X∣A) and Var(X∣A).

## Vector-valued random variables

If X∈ℝ^n is a vector-valued random variable, then we define its expected value *componentwise*: that is,

<center>E(X) = {E(X_1), E(X_2), ..., E(X_n)}^T ∈ R^n</center>

We define its variance to be a matrix:

<center>Var(X)=E[(X−E(X))(X−E(X))^T]∈ℝ^{n×n}</center>

The diagonal elements of Var(X) are the variances of the individual components of X; the off-diagonal elements are the covariances of pairs of elements of X.