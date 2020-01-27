# Lecture 03 Decision Trees (Part II)

## Decision Tree

Ex:

* A
  * 0 -> B
    * red -> +
    * blue -> -
    * green -> +
  * 1 -> C
    * 0 -> +
    * 1 -> D
      * old -> +
      * new -> -
* h(x) = x -> y
* x' = {A=1, B=blue, C=0, D=new}
* h(x') = +

## Algorithm 3: Decision Tree

```text
def h(x):
		# 3 cases
    1. internal node: test an attribute Xm on that node
    2. branch for node: select branch corresponding to the value of Xm in x
    3. leaf node: predict stored value of y (variant: return p(y|x))

def train(D): # geeneric version of the algorithm. special cases ID3, CART.
		root = new Node(data = D)
		return train_tree(root)

def train_tree(node):
		0. If node's data is perfectly classified (when error rate is zero on node's data using majority vote), then stop and return node with label = majority_vote(node's data)
		1. Xm = best attribute (pick the attribute that maximizes some spliting criterion: 1 - error rate) on which to split the node's data
				a splitting criterion is a function that measures the effectiveness of splitting on a particular attribute
						- error rate or 1 - error rate
						- Gini gain
						- Mutual information
						- random
		2. Let Xm be the decision attribute for node
		3. For each value v of attribute Xm: create a branch labeled with v
		4. Partition the node's data into descendents: D_Xm=v = {(x, y) ∈ node's data | Xm=v} for each v
		5. Recurse on each branch: for each value v, and corresponding branch, add a new node node_v = train_tree(new Node(data = D_Xm=v))
```

## Ex: Decision Tree Learning

| Y    | A    | B    | C    |
| ---- | ---- | ---- | ---- |
| -    | 1    | 0    | 0    |
| -    | 1    | 0    | 1    |
| -    | 1    | 0    | 0    |
| +    | 0    | 0    | 1    |
| +    | 1    | 1    | 0    |
| +    | 1    | 1    | 1    |
| +    | 1    | 1    | 0    |
| +    | 1    | 1    | 1    |

root = new Node(D with [5+, 3-])

* A
  * 0 - > [1+, 0-] -> +
  * 1 -> [4+, 3-] -> +
* B
  * 0 -> [1+, 3-] -> -
  * 1 -> [4+, 0-] -> +
* C
  * 0 -> [2+, 2-] -> + / - tie
  * 1 -> [3+, 1-] -> +

Classifiers on A, B, and C: hA, hB, hC

* error(hA, D) = 3 / 8
* error(hB, D) = 1 / 8
* error(hC, D) = 3 / 8 regardless of choice of whether + or -

Then, we choose B

* B
  * 0 -> [1+, 3-] A
    * 0 -> [1+, 0-] -> +
    * 1 -> [0+, 3-] -> -
  * 1 -> [4+, 0-] **+**
* B
  * 0 -> [1+, 3-] C
    * 0 -> [0+, 2-] -> -
    * 1 -> [1+, 1-] -> + / - tie
  * 1 -> [4+, 0-] **+**

Classifier on A and C

* error_A = 0 / 4

* error_C = 1 / 4

## Splitting Criteria

### Gini Impurity

* Given a random variable Y, over K classes {1,2,...,k}

* G(y) = Σ^K_{k=1} P(Y=k)P(Y≠K)
  * = Σ^K_{k=1} P(Y=k)(1-P(Y=K))
  * = Σ^K_{k=1} P(Y=k) Σ\_{j≠K}P(Y=j)
  * = 1 - Σ^K_{k=1} [P(Y=k)]^2
  * expected error rate for Solution 2

* Consider the case Y is the outcome of a weighted dice roll
  * P(Y=K) = probability of lands on side K
  * P(Y≠K) = probability of lands on any other side
* Goal: to predict the next dice roll given the weight of the die (e.g. P(Y=3) = 90%)
  * Solution 1: predict most possible side every time (e.g. Y = 3)
    * expected error rate = 1 - P(Y=y\*) = P(Y≠y\*) = Σ_{k≠y\*} P(Y=k) (e.g. 10%)
  * Solution 2: roll another die same weighted side and predict whether it lands on
    * expected error rate: G(y)
* G(Y, D) Given a dataset D, then get P(y=k) = N_y=k / N

### Gini Gain

* G(y, Xm; D) = G(y; D) - (P(Xm=0)G(y; D_Xm=0) + P(Xm=1)G(y; D_Xm=1))
* Indistinguishable with Mutual Information

### Mutual Information

Let x be a random variable x ∈ X => attribute

Let y be a random variable y ∈ Y => class

* Entrophy
  * H(Y) = -Σ_y∈Y P(Y=y)log_2 P(Y=y)
* Specific Conditional Entrophy
  * H(Y|X=x) = -Σ_y∈Y P(Y=y|X=x)log_2 P(Y=y|X=x)
* Conditional Entrophy
  * H(Y|X) = H(Y) = Σ_x∈X P(X=x)H(Y|X=x)
* Mutual Information
  * I(Y;X) = H(Y) - H(Y|X) = H(X) - H(X|Y)

Entrophy measures the expected number of bits to code one random draw from X.

For a decision tree, we want to reduce the entrophy of the random variable we are trying to predict.

Mutual Information => If we know X, how much does this reduce our uncertainty about Y.

Gini gain and Mutual information are statisically indistinguishable.

## Decision Tree Learning as Search

* Search space: all possible decision trees
* Node: single decision tree
* Edge: connects a child to a parent from which it could have been created with 1 attribute addition.
* Decision Tree Learning = greedy search, maximizing splitting criterion at each step
* Edge weight: negative of the splitting criterion

Big question: How is it that your ML algorithms can generalize the unseen examples?

**ID3** = Decision Tree Learning with Mutual Information as the splitting criterion

We say that the **inductive bias** of a ML algorithm is the principal by which it generalizes to unseen examples.

Inductive bias of ID3 is the smallest tree that matches the data with high mutual information attributes near the top.