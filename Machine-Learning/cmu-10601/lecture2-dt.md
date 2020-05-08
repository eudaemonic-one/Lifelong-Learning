# Lecture 2 Decision Tree

## Function Approximation

* Implement a simple function which returns sin(x)
* Plots of y=sin(x) are good data for c*.
* Medical Diagnosis
* N = 5, M = 4

i | sick (y^(i)) | sneezing | caughing | doctor | fox | **x**
- | - | - | - | - | - | -
1 | - | Y | N | N | N | **x**^(1)
2 | - | N | Y | N | N | **x**^(2)
3 | + | Y | Y | N | N | **x**^(3)
4 | - | Y | Y | Y | N | **x**^(4)
5 | + | N | Y | N | Y | **x**^(5)

* Problem setting
  * Set of possible inputs X
  * Set of possible outputs Y
  * Unknown target function c^*: X->Y
  * Set of candidate hypotheses H={h|h:X->Y}

### Training

Learner is given:

* Training examples D={(**x**^(1),y^(1)),...,(**x**^(n),y^(n))} of unknown target function c*, where y^(i) =c*(**x**^(i))

Learner Produce:

* Hypothesis h∈H that best approximates c^*

### Testing

To evaluate:

* Loss function l=y×y->R mesures how "bad" y^=h(**x**) are compared to c*(**x**)
* Practioner chooses functions
  * Ex: Regression y∈R l(y,y^)=(y-y^)^2 "squared loss"
  * Ex: Classification y is discrete l(y,y^)=0 if y=y^ 1 otherwise "zero-one loss"
  * Definition: Error Rate
    * Let D be a dataset and h∈H be a hypothesis
    * error(D,h) = \sigma_{i=1}^{|D|} Π(y^(i)≠h(**x**))
    * error(D,h)∈[0,1] and can use this to measure some notion of a "best" approximation.

Learner is given:

* Another dataset D^{test}={(**x**^(1),y^(1)),...,(**x**^(n),y^(n))}
  * What is our average liss on D^{test}

## Algorithm 0: Memorizer

```text
def train(D):
    store dataset D

def h(x):
    if Εx^(i)∈D s.t. x^(i)=x:
        return y^(i)
    else:
        return y∈Y randomly
```

Q: Does memorization = learning?
A: Yes, but not a form of generalization.

## Algorithm 1: Majority Vote Classifier

```text
def train(D):
    store v=majority_vote(D)
    return the class y∈Y that appears most often in D

def h(x):
    return v
```

## Algorithm 2: Decision Stump

```text
def train(D):
    pick an attribute, m
    divide dataset D on m
        D^(0)={(x,y)∈D|x_m=0}
        D^(1)={(x,y)∈D|x_m=1}
    two votes：
        v^(0)=majority_vote(D^(0))
        v^(1)=majority_vote(D^(1))

def h(x):
    if x_m==0:
        return v^(0)
    else x_m==1:
        return v^(1)
```

## ML as Function Approximation

* Problem setting
  * Set of possible inputs x
  * Set of possible outputs y
  * Unknown target function c*: X->Y
  * Set of candidate hypotheses H={h|h:x->y}
* Input space
* Output space
* Unknown target function
* Hypothesis space
* Training examples
