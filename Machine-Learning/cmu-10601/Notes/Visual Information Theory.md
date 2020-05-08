# Visual Information Theory

[Reference Link](http://colah.github.io/posts/2015-09-Visual-Information/)

## Visualizing Probability Distributions

## Aside: Simpson's Paradox

Unintuitive statistical situation

## Codes

Binary Code

## Variable-Length Codes

To minimize the message length, we’d ideally like all codewords to be short, but we especially want the commonly used ones to be.

Entropy of the distribution

## The Space of Codewords

No codeword should be the prefix of another codeword. This is called the prefix property, and codes that obey it are called prefix codes.

## Optimal Encodings

Distribute our budget in proportion to how common an event is.

## Calculating Entropy

Recall that the cost of a message of length L is 12L. We can invert this to get the length of a message that costs a given amount: log2(1cost). Since we spend p(x) on a codeword for x, the length is log2(1p(x)). Those are the best choices of lengths.
