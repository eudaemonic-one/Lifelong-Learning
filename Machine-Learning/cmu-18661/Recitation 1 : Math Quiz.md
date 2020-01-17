# Recitation 1 : Math Quiz

## Q1 Probability Rules

P(A or B) = P(A)+P(B)-P(A and B)

P(A and B) = P(A) P(B|A)

P(B|A) = P(A and B) / P(B)

## Q2 Bayes Rule

P(CMU) = a / 100

P(sleep deprived|CMU) = b / 100

P(sleep deprived|^CMU) = b / (200 * 100)

=>

P(sleep deprived)

​    = P(sleep deprived|CMU) P(CMU) + P(sleep deprived|^CMU)P(^CMU)

​    = P(sleep deprived and CMU) + P(sleep deprived and ^CMU)

=>

P(CMU / sleep deprived)

​    = P(CMU and sleep deprived) / P(sleep deprived)

​    = P(CMU) P(sleep deprived|CMU) / (P(sleep deprived|CMU) P(CMU) + P(sleep deprived|^CMU)P(^CMU))

## Q3 Mean and Variance

**Consider a continuous random variable x that is distributed uniformly.**

between 2 and 4. What are the mean and variance of x

E[x] = ∫_a^b x·f(·) dx

f(x)

​    = 1/2, 2 <= x <= 4

​    = 0, others

E[x] = ∫_2^4 x·(1/2) dx

​    = 3

Var(x)

​    = E[x^2] - E[x]^2

​    = ∫_2^4 x^2·(1/2) dx - e^2

​    = 1/3

## Q4 Matrix Algebra

(AB)^T = B^T A^T

## Q5 Convexity

Convex := draw straight line between any two point on the function curve

f(tx1+(1-t)x2) <= tf(x1)+(1-t)f(x2)

f((x1+x2)/2) <= (f(x1)+f(x2))/2

## Q6 Calculus

f(w) = l(w)+λw^2

f'(w) = l'(w)+2λw

## Q7 Matrix Calculus

## Q8 Eigenvalues

## Q9 Geometry

(i)What does the equation w⊤x+b=0 for x ∈ R3 represent if w ∈ R3  and b ∈ Rare two constants?

Line

(ii) a·b=0 a^T·b=0

(iii) What is the shortest distance between w⊤x + b1 = 0 and w⊤x + b2 = 0?

|b2-b1| / √(wTw)

## Q10 Vector-Vector Multiply

vector x (n*1)

vector y (n*1)

x^T y : inner product

x y^T : outer product

## Q11 Singular Value Decomposition

**Let the singular value decomposition of A be A = UΣV^T . Which of the following is equivalent to A^T A?**

A^T A

​    = (U∑V^T)^T (U∑V)

​    = V∑^T U^T U
