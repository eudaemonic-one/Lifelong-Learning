# Operators: Complete List

## Arithmetic

| Operator | Name                | Types                                     |
| -------- | ------------------- | ----------------------------------------- |
| `+`      | sum                 | integers, floats, complex values, strings |
| `-`      | difference          | integers, floats, complex values          |
| `*`      | product             |                                           |
| `/`      | quotient            |                                           |
| `%`      | remainder           | integers                                  |
| `&`      | bitwise AND         |                                           |
| `|`      | bitwise OR          |                                           |
| `^`      | bitwise XOR         |                                           |
| `&^`     | bit clear (AND NOT) |                                           |
| `<<`     | left shift          | integer << unsigned integer               |
| `>>`     | right shift         | integer >> unsigned integer               |

## Comparison

| Operator | Name             | Types                     |
| -------- | ---------------- | ------------------------- |
| `==`     | equal            | comparable                |
| `!=`     | not equal        |                           |
| `<`      | less             | integers, floats, strings |
| `<=`     | less or equal    |                           |
| `>`      | greater          |                           |
| `>=`     | greater or equal |                           |

* Boolean, integer, floats, complex values and strings are comparable
* Strings are ordered lexically byte-wise
* Two pointers are equal if they point to the same variable or if both are nil
* Two channel values are equal if they were created by the same call to make or if both are nil
* Two interface values are equal if they have identical dynamic types and equal concrete values or if both are nil
* A value `x` of non-interface type `X` and a value t of interface type `T` are equal if `t`’s dynamic type is identical to `X` and `t`’s concrete value is equal to `x`
* Two struct values are equal if their corresponding **non-blank fields are equal**
* Two array values are equal if their corresponding elements are equal

## Logical

| Operator | Name            | Description                              |
| -------- | --------------- | ---------------------------------------- |
| `&&`     | conditional AND | `p && q`  means "if p then q else false" |
| `||`     | conditional OR  | `p || q`  means "if p then true else q"  |
| `!`      | NOT             | `!p`  means "not p"                      |

## Pointers and Channels

| Operator | Name                | Description                                    |
| -------- | ------------------- | ---------------------------------------------- |
| `&`      | address of          | `&x`  generates a pointer to `x`               |
| `*`      | pointer indirection | `*x`  denotes the variable pointed to by `x`   |
| `<-`     | receive             | `<-ch`  is the value received from channel `ch |

## Operator Precedence

### Unary Operator

* Unary operators have the highest priority and bind the strongest

### Binary Operators (MACAO)

| Operator | Name                           | Description        |
| -------- | ------------------------------ | ------------------ |
| 1        | `*` `/` `%` `<<` `>>` `&` `&^` | **M**ultiplicative |
| 2        | `+` `-` `|` `^`                | **A**dditive       |
| 3        | `==` `!=` `<` `<=` `>` `>=`    | **C**omparison     |
| 4        | `&&`                           | **A**nd            |
| 5        | `||`                           | **O**r             |

* Binary operators of the same priority associate from **left to right**

### Statement Operators

* The `++` and `- -` operators **form statements** and fall outside the operator hierarchy

### Examples

| Expression             | Evaluation Order               |
| ---------------------- | ------------------------------ |
| `x / y * z`            | `(x / y) * z`                  |
| `*p++`                 | `(*p)++`                       |
| `^a >> b`              | `(^a) >> b`                    |
| `1 + 2*a[i]`           | `1 + (2*a[i])`                 |
| `m == n+1 && <-ch > 0` | `(m == (n+1)) && ((<-ch) > 0)` |
