# Lecture 15 PAC Learning

## PAC

* Can we bound R(h) (Unknown) in terms of R^(h) (known)? Yes
* PAC stands for Probably Approximately Correct
* PAC Learner yiedls hypothesis h ∈ H which is approximately correct (R(h)≈0) with high probability (P(R(h)≈0)≈1)
* PAC Criterion
  * Pr(|R(h)-R^(h)| <= ε) >= 1 - δ
  * R^(h) is based on a random sample of traning drawn from p*(x)
* Sample complexity is min number of training examples N s.t. the PAC Criterion is satisfied for ε and δ
* a hypothesis h ∈ H is consistant with the training data D if R^(h)=0
* For bounds:
  * Two cases for c*
    * Realizable case: c* ∈ H
    * Agnostic case: c* ∉ H or c* ∈ H
  * Two cases for |H|
    * Finite: |H| < +∞
    * Infinite: |H| = +∞
* Theorem 1: Sample Complexity
  * N >= 1/ε [ln(|H|)+ln(1/δ)] labeled examples are sufficient to ensure that with probability (1-δ) all h ∈ H with R^(h) = 0 have R(h) <= ε

