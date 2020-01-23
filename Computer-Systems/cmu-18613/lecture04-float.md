# Lecture 04 Floating point

## Fractional Binary Numbers

* Bits to right of “binary point” represent fractional powers of 2
* Represents rational number: $Σ_{k=-j}^i b_k \times 2^k$

* Representable numbers
  * Limitation #1 Can only exactly represent numbers of the form x/2k
  * Limitation #2 Limited range of numbers

## IEEE floating point standard

* Nemerical Form: $(-1)^s M 2^E$

  * Sign bit **s** determines whether number is negative or positive
  * Significand **M **normally a fractional value in range $[1.0,2.0)$
  * Exponent **E** weights value by power of two

* Encoding: |s|exp|frac|

  * MSB s is sign bit **s**
  * **exp** field encodes **E** (but not equal)
  * **frac** field encodes **M** (but not equal)

* Precision options

  * Single precision: 32 bits

    * ≈ 7 decimal digits, $10^{\pm 38}$

    * | s    | exp    | frac    |
      | ---- | ------ | ------- |
      | 1    | 8-bits | 23-bits |

  * Double precision: 64 bits

    * ≈ 16 decimal digits, $10^{\pm 308}$

    * | s    | exp     | frac    |
      | ---- | ------- | ------- |
      | 1    | 11-bits | 52-bits |

  * Other formats: half precision, quad precision

* Three "kinds" of floating point numbers

  * Normalized
    * When: **exp ≠ 0 and exp ≠ 11...11**
    * Exponent coded as a biased value: **E = exp - Bias**
      * **Bias = $2^{(k-1)} - 1$**, where k is number of exponent bits
        * Single precision: 127
        * Double precision: 1023
    * Significand coded with implied leading 1: **M = 1.xxx...x**
      * Minimum when frac = 000...0 (M = 1.0)
      * Maximum when frac = 111...1 (M = 2.0 - ε)
      * Get extra leading bit for free
  * Denormalized
    * exp = **00...00**
    * Exponent value: **E = 1 - bias**
    * Significand coded with implied leading 0: **M = 0.xxx...x**
    * Cases
      * exp = 000...0 frac = 000...0
        * zero value
        * distinct values: +0 or -0
      * exp = 000...0 frac ≠ 000...0
        * numbers closest to 0.0
        * equi-spaced
  * Special
    * exp = **11...11**
    * Cases
      * exp = 111...1 frac = 000...0
        * Infinity
        * Operation that overflows
        * Both negative and positive
        * E.g. 1.0/0.0=-1.0/-0.0 = +∞, 1.0/-0.0 = -∞
      * exp = 111...1 frac 000...0
        * Not-a-Number (NaN)
        * Represents case when no numeric value can be determined
        * E.g. sqrt(-1), ∞-∞, ∞×0

## Rounding, addition, multiplication

### Rouding

* Towards zero
* Round down
* Round up
* Nearest Even (Default)
  * Round to nearest, but if half-way in-between then round to nearest even
  * 1.BBGRXXX
    * **G**: Guard bit: LSB of result
    * **R**: Round bit: 1st bit removed
    * **X**: Sticky bit: OR of remaining bits

### Floating Point Multiplication

* $(-1)^{s_1} M_1 2^{E_1} \times (-1)^{s_2} M_2 2^{E_2}$
* Exact result: $(-1)^S M 2^E$
  * Sign s: s1 ^ s2
  * Significand M: M1 × M2
  * Exponent E: E1 + E2
* Fixing
  * If M >= 2, shift M right, increment E
  * If E out of range, overflow
  * Round M to fit frac precision

### Floating Point Addition

* $(-1)^{s_1} M_1 2^{E_1} + (-1)^{s_2} M_2 2^{E_2}$
  * Assume E1 > E2
* Exact result: $(-1)^S M 2^E$
  * Sign s, significand M:
    * Result of signed align & add
  * Exponent E: E1
* Fixing
  * If M >= 2, shift M right, increment
  * If M < 1, shift M left k positions, decrement E by k
  * Overflow if E out of range
  * Round M to fit frac precision