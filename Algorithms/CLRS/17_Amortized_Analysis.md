# 17 Amortized Analysis

## Aggregate analysis

In **aggregate analysis**, we show that for all n, a sequence of n operations takes worst-case time T(n) in total. In the worst case, the average cost, or **amortized cost**, per operation is therefore T(n)/n.

In **stack operations**, for any value of n, any sequence of n PUSH, POP, and MULTIPOP operations takes a total of $Ο(n)$ time. The average cost of an operation is $Ο(n)/n = Ο(1)$.

Consider the problem of **incrementing a k-bit binary bit counter** that counts upward from 0.

    INCREMENT(A)
    1   i ← 0
    2   while i < length[A] and A[i] = 1
    3       do A[i] ← 0
    4          i ← i + 1
    5   if i < length[A]
    6       then A[i] ← 1

An 8-bit binary counter as its value goes from 0 to 5 by a sequence of 5 operations is as follows.

|Counter value|A[7]|A[6]|A[5]|A[4]|A[3]|A[2]|A[1]|A[0]|Total cost|
|-------------|----|----|----|----|----|----|----|----|----------|
|0|0|0|0|0|0|0|0|**0**|0|
|1|0|0|0|0|0|0|**0**|**1**|1|
|2|0|0|0|0|0|0|1|**0**|3|
|3|0|0|0|0|0|**0**|**1**|**1**|4|
|4|0|0|0|0|0|1|0|**0**|7|
|5|0|0|0|0|0|1|**0**|**1**|8|

For $i > ⌊\lg n⌋$, bit A[i] never flips at all. The total number of flips in the sequence is thus $\sum^{⌊\lg n⌋}_{i=0}⌊n/2^i⌋ < n\sum^∞_{i=0}1/2^i = 2n$. The worst time for a sequence of n INCREMENT operations on an initially zero counter is therefore $Ο(n)$. The average cost of each operation and therefore the amortized cost per operation, is $Ο(n)/n=Ο(1)$.

## The accounting method

In the **accounting method** of amortized analysis, we assign differing charges to different operations, with some operations charged more or less than they actually cost. The amount we charge an operation is called its **amortized cost**. When an operation's amortized cost exceeds its actual cost, the difference is assigned to specific objects in the data structure as **credit**. Credit can be used later on to help pay for operations whose amortized cost is less than their actual cost.

In **stack operations**, we assign the following amortized costs: PUSH 2, POP 0, MULTIPOP 0.

## The potential method

The **potential method** of amortized analysis represents the prepaid work as "potential energy", or just "potential", that can be released for future operation.
