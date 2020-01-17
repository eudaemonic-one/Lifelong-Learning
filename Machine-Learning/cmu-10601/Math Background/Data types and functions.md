# Data types and functions

## Data types

In C,

* struct { int a, float b } defines teh type Z × R of pairs of an integer and a real number.
* union { int a, float n } defines the type Z ∪ R of integers together with reals.

## Tagged unions

Tagged unions prevent us from (accidentally or on purpose) treating an object of one type as an object of another type.

## Functions

p(x, y) = 3x^2 + xy => float p(float x, float y) { return 3 * pow(x, 2) + x * y; }

## Anonymous functions

lambda