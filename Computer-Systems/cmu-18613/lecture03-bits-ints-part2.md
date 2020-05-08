# Lecture 03 Bits, Bytes & Integers II

## Addition, Negation, Multiplication

### Unsigned Addition

* Ignores carry output
* $s = UAdd_w(u, v) = u + v \space mod 2^w$

### Two's Complement Addition

* TAdd and UAdd have identical bit-level behavior
* Discard carry
* $TAdd_w(u, v) =$
  * $u + v + 2^w, u + v<TMin_w$ (Negative overflow)
  * $u+v, TMin_w <= u + v <= TMax_w$
  * $u+v-2^w, TMax_w < u + v$ (Positive overflow)

### Multiplication

* exact results can be bigger than w bits

#### Unsigned Multiplication in C

* Ignores high order w bits
* $UMult_w(u, v) = u · v \space mod 2^w$

#### Signed Multiplication in C

* Ignores high order w bits
* Some of which are different for signed vs. unsigned multiplication
* Lower bits are the same

#### Power-of-2 Multiply with Shift

* `u << k` gives $u * 2^k$
* Both signed and unsigned

#### Unsigned Power-of-2 Divide with Shift

* `u >> k` gives lower $(u / 2^k)$
* Uses logical shift

#### Signed Power-of-2 Divide with Shift

* u >> k gives lower $(u / 2^k)$
* Uses arithmetic shift
* **Rounds wrong** direction when x < 0

#### Correct Power-of-2 Divide

* Compute as lower $((x+2^k-1) / 2^k)$
* In C: `(x + (1<<k) - 1) >> k`

### Negation

* Negate through complement and increase
  * `~x + 1 = -x`
* $x = 0$ => $-x = 0$
* $x = TMin$ => $-x = TMin$

## Why Should I Use Unsigned

* Do use when performing modular arithmetic
* Do use when using bits to represent sets
* Do use in system programming
  * Bit masks, device commands

## Representation in memory, pointers, strings

### Byte-Oriented Memory Organization

* Programs refer to data by address
* **Note: system provides private address spaces to each “process”**
  * Think of a process as a program being executed
  * So, a program can clobber its own data, but not that of others

### Machine Words

* Any given computer has a "Word Size"
  * Nominal size of integer-valued data and of address

### Word-Oriented Memory Organization

* Addresses specify byte locations
  * Address of first byte in word

### Byte Ordering

* Big Endian: Sun, PPC Mac, Internet
  * Least significant byte has highest address
* Little Endian: x86, ARM, iOS, and linux
  * Least significant byte has lowest address