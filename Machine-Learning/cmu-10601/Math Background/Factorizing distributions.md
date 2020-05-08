# Factorizing distributions

## Factored form

*factored form*: we write the full joint probability distribution as a product of several terms, with each individual term depending only on a subset of the variables.

As with any compression scheme, we trade off accuracy against size: we may not be able to represent our exact distribution in compact form, but if we can get close, the size savings may be well worth it.

## Example: rusty robot

The joint table has 32 entries. But, if we represent it in factored form as

<center>P(O,M,R,W,U) = P(O)P(M)P(R)P(W∣O,R)P(U∣W,M)</center>

then the individual tables have

<center>2+2+2+8+8 = 22</center>

entries total.

## Broadcasting

*broadcasting*: replicating an array along its missing dimensions

<center>P(A)P(B∣A)P(C∣B)=P(A,B)P(C∣B)=P(A∣B)P(B)P(C∣B)=P(A∣B)P(B,C)=P(A∣B)P(B∣C)P(C)</center>

## Bayes rule

<center>P(A|B) = P(B|A)P(A) / P(B)</center>

<center>P(A|B,C) = P(B|A,C)P(A,C) / P(B,C)</center>

## Independence

Two random variables X and Y are said to be *independent* if knowledge of one does not tell us anything about the value of the other — that is, if P(X∣Y)=P(X)P(X∣Y)=P(X) and P(Y∣X)=P(Y)P(Y∣X)=P(Y).

## Conditional independence

if X and Y are independent after conditioning on an event Z=z (that is, in the distribution P(X,Y∣Z=z) then we say X and Y are *conditionally independent* given Z=z.