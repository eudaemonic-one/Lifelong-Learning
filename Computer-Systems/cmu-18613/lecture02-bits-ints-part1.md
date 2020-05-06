# Lecture 02 Bits, Bytes, & Integers I

## Representing Information As Bits

* Each bit is 0 or 1
* By encoding/interpreting sets of bits in various ways
* Why bits? Electronic implementation.
* Represent number

### Encoding Byte Values

* Byte = 8 bits

## Bit-level manipulation

### Boolean Algebra

* Encode "True" as `1` and "False" as `0`.
* And `A & B = 1` when both `A = 1` and `B = 1`
* Or `A | B = 1` when either `A = 1` or `B = 1`
* Not `~A = 1` when `A = 0`
* Exclusive-Or (Xor) `A ^ B = 1` when either `A = 1` or `B = 1`, but not both

### Operate on Bit Vectors

* All of the properties of Boolean algebra apply.

### Representing & Manipulating Sets

* Width w bit vector represents subsets of $\{0,\cdots,w-1\}$
* $a_j=1$ if $j \in A$
* `&` Intersection
* `|` Union
* `^` Symmetric difference
* `~` Complement

### Logic Operations

* View `0` as "Fasle"
* Anything nonzero as "True"
* Always return `0` or `1`
* Early termination
* Aside: `p & *p` (avoid null pointer access)

### Shift Operations

* Left Shift: `x << y`

  * Shift bit-vector x right y positions
    * Throw away extra bits on left

  * Fill with 0's on right

* Right Shift: `x >> y`

  * Shift bit-vector x right y positions
    * Throw away extra bits on right
  * Logical shift
    * Fill with 0's on left
  * Arithmetic shift
    * Replicate most significant bit one left

* Undefined Behavior

  * Shift amount < 0 or >= word size

## Integers

### Representation: `unsigned` and `signed`

* C does not mandate using two's complement

* Sign Bit
  * For 2's complement, most significant bit indicates sign
    * 0 for nonnegative
    * 1 for negative

* Unsigned Values
  * $UMin = 0$
  * $UMax = 2^w-1$

* Two's Complement Values
  * $TMin = -2^{w-1}$
  * $TMax = 2^{w-1}-1$
  * $Minus 1 = 111...1$
  * $|TMin| = TMax + 1$
  * $UMax = 2 * TMax + 1$

### Conversion, Casting

* Mappings between unsigned and two's complement numbers: **Keep bits representations and reinterpret**
* Constants
  * By default are considered to be signed integers
  * Unsigned if have "U" as suffix
* Casting

  * Explicit casting
  * Implicit casting
  * Expression Evaluation
    * If there is a mix of unsigned and signed in single expression, **signed values implicitly cast to unsigned**
    * Including comparison operations `<`, `>`, `==`, `<=`, `>=`

### Expanding, Truncating

* Sign Extension
  * Given w-bit signed integer x and convert it into w+k-bit integer with same value
  * Make k copies of sign bit
  * $X' = x_{w-1}, ..., x_{w-1}, x_{w-2}, ..., x_0$
  * C automatically performs sign extension
* Truncation

  * Given k+*w*-bit signed or unsigned integer X and convert it to *w*-bit integer X’ with same value for “small enough” X
  * Drop top k bits
  * $X' = x_{w-1}, x_{w-2}, ..., x_0$
  * Unsigned: mod operation Signed: similar to mod
