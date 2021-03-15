# Lecture 5: Physical Layer III

## Line Coding

### Synchronization

* Synchronization of clocks in transmitters and receivers
  * clock drift causes a loss of synchronization
* Asynchronization transmission
  * Avoids synchronization loss by
    * specifying a short maximum length for the bit sequences (so that clock doesn't drif much within sequence)
    * and resetting the clock in the beginning of each bit sequence (by using a start bit)
  * ASCII code: 7 bits to represent 128 letters, symbols, and control characters
    * Asynchronous transmission sends sequences of 8 bits = one start bit + 7 ASCII bits
* Synchronous transmission
  * Improves efficiency by transmitting longer sequences of bits, called packets (variable length)
  * Requires extra information to indicate the end of the packet

### Encoding

* Encoding converts a binary information sequence into a digital signal
* Encoding can be done one bit at a time or in blocks of multiple bits called a symbol
* Transmission is synchronous, i.e., a clock is used to sample the signal
* Why do we need encoding?
  * To meet certain electrical constraints (e.g., avoid long sequence of zeros/ones)
  * Create control symbols, besides regular data symbols (e.g., start or end of frame)
  * Can do error detection or error correction

![line_coding_examples](images/lecture05-physical3/line_coding_examples.png)

#### mB/nB Encoding

* m data bits are coded as symbols of n line bits
* Each valid symbol has at least two 1s: get dense transitions
* Example: FDDI uses 4B/5B

![4b5b_encoding](images/lecture05-physical3/4b5b_encoding.png)

## Error Detection and Correction

### Error Control

* Channels introduce errors in digital communications
* Two basic approaches:
  * Error detection & retransmission (ARQ)
  * Forward error correction (FEC)
* Redundancy: only a subset of all possible blocks can be valid codewords
* Undetectable error: when channel transforms a codeword into another valid codeword
* Good codes should maximize sepration between valid codewords

### Single Parity Check

* Check Bit: $b_{k+1} = b_1 + b_2 + b_3 + \cdots + b_k modulo 2$
* Receiver checks to see if # of 1s is even
* Coverage: all error patterns with odd # of errors can be detected

### Two-Dimensional Parity Check

* Arrange information as columns
* Add single parity bit to each column
* Add a final parity column
* Used in early error control systems

### Internet Checksum

* Several Internet protocols (e.g., IP, TCP, UDP) use check bits to detect errors in the header
* A checksum is calculated for header contents and included in a special field
* Checksum is potentially recalculated at every router, so algorithm selected for ease of implementation in software

### Polynomial Code

* Cyclic redundancy check (CRC)
* Implemented using shift-register circuits
* Most data communication standards use polynomial code for error detection
* $b(x) = x^{n-k}i(x) + r(x)$

![crc](images/lecture05-physical3/crc.png)

![find_polynomial_codeword](images/lecture05-physical3/find_polynomial_codeword.png)

* Undetectable error patterns
  * If e(x) is a multiple of g(x), that is, e(x) is a non-zero codeword
  * Choose the generator polynomial so that selected error patterns can be detected

![undetectable_error_patterns](images/lecture05-physical3/undetectable_error_patterns.png)

#### Design Good Polynomial Code

* Select generator polynomial so that likely error patterns are not multiples of g(x)
* Detect Single Errors
  * $e(x) = x^i$ for error in location i + 1
  * If g(x) has more than 1 term, it cannot divide $x^i$
* Detect Double Errors
  * $e(x) = x^i + x^j = x^i (x^{j-i} + 1)$ where j > i
  * If g(x) has more than 1 term, it cannot divide $x^i$
  * If g(x) is a primitive polynomial, it cannot divide $x^m + 1$ for all $m < 2^{n-k} - 1$
  * Primitive polynomials can be found by consulting coding theory books

### Hamming Code

* Class of error-correcting code
* Capable of correcting all single-error patterns
* Probably optimal for 1-bit errors
* Very less redundancy, e.g., 1-bit error proof - add O(logN) bits of redundancy for n bit sequence
* e.g., m=3 Hamming Code
  * Information bits b1, b2, b3, b4
  * Equations for parity checks b5, b6, b7
    * b5 = b1 + b3 + b4
    * b6 = b1 + b2 + b4
    * b7 = b2 + b3 + b4