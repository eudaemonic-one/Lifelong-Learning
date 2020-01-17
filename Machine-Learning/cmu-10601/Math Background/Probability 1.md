# Probability 1

## Atomic events

Many events can’t be predicted with total certainty. We use the idea of probability to say how likely they are to happen.

We think of a **universe** or **sample space** S that is a set of possible **atomic events**.

P(a)≥0 is the probability that an event a∈S occurs, and P(¬a)=1−P(a) is the probability that it does not.

Probability is always between 0 and 1: impossible = 0, certain = 1.

## Compound events

A compound event E⊆S happens exactly when one of the atomic events a∈E happens.

<center>P(E)=∑_{a∈E} P(a)</center>

The sum of all atomic events' probabilities must be 1:

<center>P(S) = 1</center>

## Venn diagrams

We picture the universe as a large rectangle, and events as shapes that are subsets of the rectangle.

## Experiment

We can use probabilities to describe **experiments** that we are thinking about conducting: we make a probability model that describes what we think might happen.  Atomic events represent possible outcomes of the experiment, and the universe contains all possible outcomes.

## Uniform vs. non-uniform distributions

* Fair Die/Uniform Model
* Weighted Die/Non-uniform Model

## Disjoint union

If A and B are disjoint events (i.e., A∩B=∅), then

<center>P(A∪B) = P(A) + P(B).</center>

## Non-disjoint union

<center>P(A∪B) = P(A) + P(B) - P(A∩B)</center>

## Sum rule

<center>P(A) = P(A|B) + P(A|¬B)</center>

## Logic rules

<center>P(¬(A∧B)) = P(¬A∨¬B)</center>

<center>P(¬(A∨B))=P(¬A∧¬B)</center>