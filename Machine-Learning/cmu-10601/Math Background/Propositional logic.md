# Propositional logic

## Arguments

* state assumptions clearly
* describe, for each step, which earlier assumptions or conclusions are inputs, and what conclusion we're now drawing for later use
* organize arguments into digestible (and independently checkable) chunks
* write out everything in enough detail that anyone who comes along later can follow the argument.

## Propositional logic

In propositional logic, variables a, b, ... (called *propositions* or *predicates*) represent truth values T or F, and connectives ∨, ∧, ¬, → represent OR, AND, NOT, and IMPLIES.

## Proofs

**From premises ϕ and ϕ→ψ, conclude ψ.**

The premises ϕ and ϕ→ψ are assumptions or previously proven statements; ϕ and ψ can be single variables or more complex statements containing connectives.

## Inference rules

- ∧ introduction: if we separately prove ϕ and ψ, then that constitutes a proof of ϕ∧ψ.
- ∧ elimination: from ϕ∧ψ we can conclude ϕ and ψ.
- ∨ introduction: from ϕ we can conclude ϕ∨ψ for any ψ.
- ∨ elimination (also called proof by cases): if we know ϕ∨ψ (the cases) and we have both ϕ→χ and ψ→χ (the case-specific proofs), then we can conclude χ.
- Associativity: both ∧ and ∨ are associative (it doesn't matter how we parenthesize an expression like a∧b∧c∧d, so in fact we often just leave the parentheses out).
- Distributivity: ∧ and ∨ distribute over one another; for example, a∧(b∨c) is equivalent to (a∧b)∨(a∧c).
- Commutativity: both ∧ and ∨ are commutative (symmetric in the order of their arguments), so we can re-order their arguments however we please. For example, b∨c∨a is equivalent to a∨b∨c.

## Resolution

Suppose we have two statements of the form ϕ∨χ and ¬ϕ∨ψ, resolution let us conclude χ∨ψ. That is, we delete ϕ and ¬ϕ from our formulas, and join together what's left using ∨.