# Representing and Manipulating Information

## Information Storage

bytes := blocks of eight bits

address := identify every byte of memory by a unique number

virtual address space := the set of all possible address

### Hexadecimal Notation

when a value x is a power of two, that is, x=2^n, for n written in the form i+4j, where 0<=i<=3, we can write x with a leading hex digit of 1 (i=0), 2(i=1), 4(i=2), 8(i=3), followed by j hexadecimal 0s.

decimal notation -> hexadecimal notation := x=q·16+r we use the hexadecimal digit representing r as the least significant digit and generate the remaining digits by repeating the process on q.

### Data Sizes

size := indicating the nominal size of integer and pointer data. for a machine with a w-bit word size, the virtual addresses can range from 0 to 2w − 1.

### Addressing and Byte Ordering

big endian & little endian := IBM and Sun Microsystems machines/Intel-compatible machines

### Introduction to Boolean Algebra

Boolean Algebra := Binary values 1 and 0 encode logic values True and False, while operations ~, &, |, and ^ encode logical operations Not, And, Or, and Exclusive-Or, respectively. a ^ a = 0. (a ^ b) ^ a = b.

### Bit-Level Operations in C

```c
void inplace_swap(int *x, int *y) {
    *y=*x^*y;
    *x=*x^*y;
    *y=*x^*y;
}
```

### Logical Operations in C

C provides a set of logical operators ||, &&, and !, which correspond to the Or, And, and Not operations of logic.

### Shift Operations in C

For an operand x having bit representation [xn−1, xn−2, . . . , x0], the C expression x << k yields a value with bit representation [xn−k−1, xn−k−2,...,x0,0,...0]. That is, x is shifted k bits to the left, dropping off the k most significant bits and filling the right end with k zeros.

Generally, machines support two forms of right shift: logical and arithmetic. A logical right shift fills the left end with k zeros, giving a result [0, . . . , 0, xn−1, xn−2, . . . xk]. An arithmetic right shift fills the left end with k repe- titions of the most significant bit, giving a result [xn−1, . . . , xn−1, xn−1, xn−2, . . . xk].

## Integer Representations

### Two's-Complement Encodings

$$B2T_(w(x))=-x_(w-1)2^(w-1)+\sum_i=0^(w-2)x_i2^i$$

|TMin| = |TMax| + 1

## Conversions Between Signed and Unsigned

conversions between signed and unsigned numbers with the same word size—the numeric values might change, but the bit patterns do not.

## Interger Arithmetic

Unsigned arithmetic can be viewed as a form of modular arithmetic. Unsigned addition is equivalent to computing the sum modulo 2^w.

the w-bit two’s-complement sum can be obtained by performing binary addition of the operands and truncating the result to m bits.

-2^(w-1)+-2^(w-1)

One technique for performing two’s-complement negation at the bit level is to complement the bits and then increment the result.

signed multiplication in C generally is performed by truncating the 2w-bit product to w bits.

⌊a⌋ is defined to be the unique integer a′ such that a′ ≤ a < a′ + 1. As examples, ⌊3.14⌋ = 3, ⌊−3.14⌋ = −4, and ⌊3⌋=3.

## Floating Point

### IEEE Floating-Point Representation

The IEEE floating-point standard represents a number in a form V = (−1)s × M × 2E:

* single sign bit s
* k-bit exponent field
* n-bit fraction field

single-precision := s=1, k=8, n=23 bits either
double-precision := s=1, k=11, n=52

1. Normalized := | s | ≠0 & ≠255 | f | E = e − Bias Bias = 2^(k-1) - 1 M = 1 + f
2. Denormalized := | s | 00000000 | f | E = 1 - f M = f
3. a.Infinity := | s | 11111111 | 0...0 |
4. b.NaN := | s | 11111111 | ≠0 |

Denormalized numbers serve to provide a way to represent 0 and numbers that are very close to 0.0.

Infinity can represent results that overflow, as when we
multiply two very large numbers, or when we divide by zero.

NaN (short for "Not a Number") are returned as the result of an operation where the result cannot be given as a real number or as infinity, as when computing −1 or ∞ − ∞. They can also be useful in some applications for representing uninitialized data.

### Rounding

Round-to-even := Rounding toward even numbers avoids statistical bias in most real-life situations.
Round-toward-zero
Round-down
Round-up

* From int to float, the number cannot overflow, but it may be rounded.
* From int or float to double, the exact numeric value can be preserved be- cause double has both greater range (i.e., the range of representable values), as well as greater precision (i.e., the number of significant bits).
* From double to float, the value can overflow to +∞ or −∞, since the range is smaller. Otherwise, it may be rounded, because the precision is smaller.
* From float or double to int the value will be rounded toward zero.

Floating-point arithmetic must be used very carefully, because it has only limited range and precision, and because it does not obey common mathematical properties such as associativity.

## Miscs

~x+1 is equivalent to -x.

(1<<k)-1 to generate masks.
