# Probability 2

## Random variables

If F∈S→ℝ is a function that maps atomic events to real numbers, we will say that F is a **random variable**.

## Joint distribution

To construct our sample space, we take our atomic events to be all of the possible joint settings of the variables.

## Probability tables

Tables are often called **multidimensional arrays** or **tensors**, and their dimensions are called **modes**.

Importantly, a probability table is *not just* a matrix or tensor: its dimensions are labeled with random variables. When we work with a probability table, the labels are what matter, not the order that we happen to write them down.

## Probability tables and compound events

To get the probability of a compound event, we can use the sum rule: we just add together the appropriate table entries.

## Marginal distribution

Given a joint distribution over several random variables (say height, weight, and eye color) we can ask for the distribution over some subset of the variables (say just height and eye color).  We say that the second distribution is a *marginal* of the first one; the process of forming a marginal is called *marginalizing*.

## Conditional distribution

Given a joint distribution, we can also ask what happens if we fix the value of one or more of the variables. The result is called a *conditional* distribution; we say that we are *conditioning on* or *observing* the random variables that we are fixing.

## Notations for conditional distributions

If we write P(X,Y∣Z,W), we mean the conditional distribution of the random variables X and Y, given the values of the random variables Z and W.
