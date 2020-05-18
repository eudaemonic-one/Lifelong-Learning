# Type Alias Explained

* An **alias declaration** has the form
* `type T1 = T2`
* as opposed to a standard **type definition**
* `type T1 T2`
* An alias declaration doesn’t create a new distinct type different from the type it’s created from
* It just introduces an alias name `T1`, an alternate spelling, for the type denoted by `T2`
* Type aliases are **not** meant for everyday use
  * They were introduced to support **gradual code repair** while moving a type between packages during large-scale refactoring
