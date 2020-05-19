# Object-oriented Programming without Inheritance

* **Go doesn’t have inheritance – instead composition, embedding and interfaces support code reuse and polymorphism**

## Object-oriented Programming with Inheritance

* Inheritance in traditional object-oriented languages offers three features in one
* When a Dog inherits from an Animal
  * e `Dog` class reuses code from the `Animal` class,
  * a variable `x` of type `Animal` can refer to either a `Dog` or an `Animal`,
  * `x.Eat()` will choose an `Eat` method based on what type of object `x` refers to
* In object-oriented lingo, these features are known as **code reuse**, **polymorphism** and **dynamic dispatch**
* All of these are available in Go, using separate constructs:
  - **composition** and **embedding** provide code reuse,
  - **interfaces** take care of polymorphism and dynamic dispatch

## Code Reuse by Composition

* If a `Dog` needs some or all of the functionality of an `Animal`, simply use **composition**
  * This gives you full freedom to use the `Animal` part of your `Dog` as needed

```go
type Animal struct {
	// …
}

type Dog struct {
	beast Animal
	// …
}
```

## Code Reuse by Embedding

* If the `Dog` class inherits **the exact behavior** of an `Animal`, this approach can result in some tedious coding

```go
type Animal struct {
	// …
}

func (a *Animal) Eat()   { … }
func (a *Animal) Sleep() { … }
func (a *Animal) Breed() { … }

type Dog struct {
	beast Animal
	// …
}

func (a *Dog) Eat()   { a.beast.Eat() }
func (a *Dog) Sleep() { a.beast.Sleep() }
func (a *Dog) Breed() { a.beast.Breed() }
```

* This code pattern is known as **delegation**
* Go uses **embedding** for situations like this
* The declaration of the `Dog` struct and it’s three methods can be reduced to

```go
type Dog struct {
	Animal
	// …
}
```

## Polymorphism and Dynamic Dispatch with Interfaces

* **Keep your interfaces short, and introduce them only when needed**
* Further down the road your project might have grown to include more animals
* At this point you can introduce polymorphism and dynamic dispatch using interfaces
* If you need to put all your pets to sleep, you can define a `Sleeper` interface

```go
type Sleeper interface {
	Sleep()
}

func main() {
	pets := []Sleeper{new(Cat), new(Dog)}
	for _, x := range pets {
		x.Sleep()
	}
}
```

* No explicit declaration is required by the `Cat` and `Dog` types
* Any type that provides the methods named in an interface may be treated as an implementation
