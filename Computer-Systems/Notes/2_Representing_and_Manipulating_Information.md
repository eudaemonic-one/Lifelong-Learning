# Representing and Manipulating Information

## Information Storage

bytes := blocks of eight bits

address := identify every byte of memory by a unique number

virtual address space := the set of all possible address

### Hexadecimal Notation

when a value x is a power of two, that is, $x=2^n$, for n written in the form $i+4j$, where $0<=i<=3$, we can write x with a leading hex digit of $1 (i=0), 2(i=1), 4(i=2), 8(i=3)$, followed by j hexadecimal 0s.

decimal notation -> hexadecimal notation := $x=q·16+r$ we use the hexadecimal digit representing r as the least significant digit and generate the remaining digits by repeating the process on q.

### Data Sizes

size := indicating the nominal size of integer and pointer data. for a machine with a w-bit word size, the virtual addresses can range from $0$ to $2^w − 1$.

### Addressing and Byte Ordering

big endian := the most significant byte comes first (IBM and Sun Microsystems machines)

little endian := the least significant byte comes first (Intel-compatible machines)

### Introduction to Boolean Algebra

Boolean Algebra := Binary values 1 and 0 encode logic values True and False, while operations ~, &, |, and ^ encode logical operations Not, And, Or, and Exclusive-Or, respectively. a ^ a = 0. (a ^ b) ^ a = b.

One useful application of bit vectors is to represent finite sets. We can encode anysubset $A ⊆ {0,1,...,w−1}$ with a bit vector $[a_{w−1},...,a_1,a_0]$, where $a_i = 1$ if and only if $i ∈ A$.

We can selectively enable or disable different signals by specifying a bit-vector mask, where a 1 in bit position i indicates that signal i is enabled, and a 0 indicates that it is disabled. Thus, the mask represents the set of enabled signals.

### Bit-Level Operations in C

| for Or, & for And, ~ for Not, and ^ for Exclusive-Or.

```c
void inplace_swap(int *x, int *y) {
    *y=*x^*y;
    *x=*x^*y;
    *y=*x^*y;
}
```

### Logical Operations in C

C provides a set of logical operators ||, &&, and !, which correspond to the Or, And, and Not operations of logic.

The logical operations returns either 1 or 0, indicating a result of either True or False, respectively.

Logical operators do not evaluate their second argument if the result of the expression can be determined by evaluat- ing the first argument.

### Shift Operations in C

For an operand x having bit representation $[x_{n−1}, x_{n−2}, . . . , x_0]$, the C expression x << k yields a value with bit representation $[x_{n−k−1}, x_{n−k−2},...,x_0,0,...0]$. That is, x is shifted k bits to the left, dropping off the k most significant bits and filling the right end with k zeros.

Generally, machines support two forms of right shift: logical and arithmetic. A logical right shift fills the left end with k zeros, giving a result $[0, . . . , 0, x_{n−1}, x_{n−2}, . . . x_k]$. An arithmetic right shift fills the left end with k repetitions of the most significant bit, giving a result $[x_{n−1}, . . . , x_{n−1}, x_{n−1}, x_{n−2}, . . . x_k]$.

In practice, almost all compiler/machine combinations use arithmetic right shifts for signed data, and many programmers assume this to be the case.

Java, on the other hand, has a precise definition of how right shifts should be performed. The expression x >> k shifts x arithmetically by k positions, while x >>> k shifts it logically.

When shifting amount $k >= w$, the shift amount is effectively computed as $k mod w$. However, this behavior is not guaranteed for C programs.

## Integer Representations

### Two's-Complement Encodings

$$B2T_{w}(x)=-x_{w-1}2^{w-1}+\sum_{i=0}^{w-2}x_{i}2^{i}$$

The two’s-complement range is asymmetric:

$$|TMin| = |TMax| + 1$$

### Conversions Between Signed and Unsigned

The effect of casting is to keep the bit values identical but change how these bits are interpreted.

Conversions between signed and unsigned numbers with the same word size—the numeric values might change, but the bit patterns do not.

### Signed vs. Unsigned in C

When an operation is performed where one operand is signed and the other is unsigned, C implicitly casts the signed argument to unsigned and performs the operations assuming the numbers are nonnegative.

### Expanding the Bit Representation of a Number

To convert an unsigned number to a larger data type, we can simply add leading zeros to the representation; this operation is known as *zero extension*. For converting a two’s- complement number to a larger data type, the rule is to perform a *sign extension*, adding copies of the most significant bit to the representation.

### Truncating Numbers

When truncating a w-bit number $x = [x_{w−1}, x_{w−2}, . . . , x_0]$ to a k-bit number, we drop the high-order w − k bits, giving a bit vector $x′ = [x_{k−1}, x_{k−2}, . . . , x_0]$. Truncating a number can alter its value — a form of overflow.

## Interger Arithmetic

Unsigned arithmetic can be viewed as a form of modular arithmetic. Unsigned addition is equivalent to computing the sum modulo $2^w$.

The w-bit two’s-complement sum of two numbers has the exact same bit-level representation as the unsigned sum.

$$-2^{w-1}+-2^{w-1}$$

One technique for performing two’s-complement negation at the bit level is to complement the bits and then increment the result. In C, we can state that for any integer value x, computing the expressions $-x$ and $~x + 1$ will give identical results.

Signed multiplication in C generally is performed by truncating the 2w-bit product to w bits.

Consider the task of generating code for the expression x * K, for some constant K, we can compute the effect of these bits on the product using either of two different forms:

* Form A: $(x<<n) + (x<<n−1) + . . . + (x<<m)$
* Form B: $(x<<n+1) - (x<<m)$

The two different shifts — logical and arithmetic — serve the purpose for division of unsigned and two’s complement numbers, respectively.

Integer division always rounds toward zero.

## Floating Point

### IEEE Floating-Point Representation

The IEEE floating-point standard represents a number in a form $V = (−1)^s × M × 2^E$:

* single sign bit s
* k-bit exponent field
* n-bit fraction field

* single-precision := s=1, k=8, n=23 s=bits[31] exp=bits[30:23] frac=bits[22:0]
* double-precision := s=1, k=11, n=52 s=bits[63] exp=bits[62:52] frac=bits[51:0]

1. Normalized := | s | ≠0 & ≠255 | f | $E = e − Bias$ $Bias = 2^{k-1} - 1$ $M = 1 + f$
2. Denormalized := | s | 00000000 | f | $E = 1 - Bias$ $M = f$
3. a.Infinity := | s | 11111111 | 0...0 |
4. b.NaN := | s | 11111111 | ≠0 |

Denormalized numbers serve to provide a way to represent 0 and numbers that are very close to 0.0.

Infinity can represent results that overflow, as when we
multiply two very large numbers, or when we divide by zero.

NaN (short for "Not a Number") are returned as the result of an operation where the result cannot be given as a real number or as infinity, as when computing $−1$ or $∞ − ∞$. They can also be useful in some applications for representing uninitialized data.

### Rounding

Round-to-even := Rounding toward even numbers avoids statistical bias in most real-life situations.
Round-toward-zero
Round-down
Round-up

### Floating-Point Operations

Floating point operations lack associativity and distributivity, but are commutative.

### Floating Point in C

* From int to float, the number cannot overflow, but it may be rounded.
* From int or float to double, the exact numeric value can be preserved because double has both greater range (i.e., the range of representable values), as well as greater precision (i.e., the number of significant bits).
* From double to float, the value can overflow to $+∞$ or $−∞$, since the range is smaller. Otherwise, it may be rounded, because the precision is smaller.
* From float or double to int the value will be rounded toward zero.

Floating-point arithmetic must be used very carefully, because it has only limited range and precision, and because it does not obey common mathematical properties such as associativity.

## Miscs

$~x+1$ is equivalent to $-x$.

$(1<<k)-1$ to generate masks.
