# Learn Go Programming

## Overview

### Features of Go Programming

* Support for environment adopting patterns similar to dynamic languages. For example, type inference (x := 0 is valid declaration of a variable x of type int)
* Compilation time is fast.
* Inbuilt concurrency support: lightweight processes (via go routines), channels, select statement.
* Go programs are simple, concise, and safe.
* Support for Interfaces and Type embedding.
* Production of statically linked native binaries without external dependencies.

### Features Excluded Intentionally

* Support for type inheritance
* Support for method or operator overloading
* Support for circular dependencies among packages
* Support for pointer arithmetic
* Support for assertions
* Support for generic programming

## Program Structure

* Package Declaration
* Import Packages
* Functions
* Variables
* Statements and Expressions
* Comments

## Basic Syntax

### Tokens in Go

A token is either a keyword, an identifier, a constant, a string literal, or a symbol.

### Line Separator

The Go compiler internally places “;” as the statement terminator to indicate the end of one logical entity.

### Comments

Comments start with /* and terminates with the characters */

### Identifiers

identifier = letter { letter | unicode_digit }
