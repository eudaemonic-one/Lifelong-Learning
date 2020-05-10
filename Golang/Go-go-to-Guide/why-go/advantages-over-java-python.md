# Why Go? – Key advantages you may have overlooked

## Minimalism

* The **built-in** frameworks for testing and profiling are small and easy to learn, but still fully functional
* It’s possible to **debug** and **profile** an optimized binary running in production through an HTTP server
* Go has automatically generated documentation with testable examples. Once again, the **interface is minimal**, and there is very little to learn
* Go is **strongly** and **statically** typed with no implicit conversions, but the syntactic overhead is still surprisingly small
* Programs are constructed from packages that offer clear **code separation** and allow efficient management of dependencies
* Structurally typed interfaces provide runtime **polymorphism** through dynamic dispatch
* **Concurrency** is an integral part of Go, supported by goroutines, channels and the select statement

## Features For The Future

* New features are considered only if there is a pressing need demonstrated by **experience reports** from real-world projects

## Code Transparency

* You **always** need to know **exactly what** your coding is doing
* and **sometimes** need to **estimate the resources** (time and memory) it uses

## Compability

* Go 1 has succinct and strict compatibility guarantees for the core language and standard packages – Go programs that work today should continue to work with future releases of Go 1
* Go is an **open source** project and comes with a BSD-style license that permits commercial use, modification, distribution, and private use

## Performance

* The exterior of Go is far from flashy, but there is a **fine-tuned engine** underneath
* First, Go is a **compiled** language
	* An executable Go program typically consists of a **single standalone binary**, with no separate dynamic libraries or virtual machines, which can be **directly deployed**
* **Size and speed of generated code** will vary depending on target architecture
	* Go code generation is fairly mature and the major OSes (Linux, macOS, Windows) and architectures (Intel x86/x86-64, ARM64, WebAssembly, ARM), as well as many others, are supported
* Go is **garbage collected**, protecting against memory leaks
* The **standard libraries** are typically of high quality, with optimized code using efficient algorithms
* **Build speeds**, in absolute terms, are currently fairly good
	* More importantly, Go is **designed** to make compilation and dependency analysis easy, making it possible to create programming tools that scales well with growing projects
