# Features Overview

As a recent high-level language, Go has the benefit of hindsight, and **the basics are done well**:

* Garbage collection
* Package system
* Static Type system
* First-class functions
* Lexical scope
* System-call interface
* Immutable strings (text generally encoded in UTF-8)
* Guarantees backwards compatibility

But it has comparatively few features and is unlikely to add more. For instance is **missing**:

* No implicit number conversions
* No constructors/destructors
* No operator overloading
* No default parameter values
* No inheritance
* No generics
* No exceptions
* No macros
* No function annotations
* No thread-local storage

## Type system

Go has enough of a type system to avoid most of the careless mistakes from dynamic languages, but it has a simpler type system than comparable typed languages.

> In practice, Go gives programmers much of the safety and run-time performance benefits of a relatively strong type system without the burden of a complex one.

## Contemporary Computer System Design

Its built-in data types and most library data structures are crafted to work naturally without explicit initialization or implicit constructors, so relatively few memory allocations and memory writes are hidden in the code.

Go's aggregate types (structs and arrays) hold their elements directly, requiring less storage and fewer allocations and pointer indirections that languages that use indirect fields.

And since the modern computer is a parallel machine, Go has concurrency features based on CSP (*Communicating Sequential Processes*).

The variable size stacks of Go's lightweight threads or *goroutines* are initially small enough that creating one goroutine is cheap and creating a million is practical.

## Standard Library

Go's standard library provides clean building blocks and APIs for I/O, text processing, graphics, cryptography, networking and distributed applications, with support for many standard file formats and protocols.

The libraries and tools make extensive use of contention to reduce the need for configuration and explanation, thus simplifying program logic and making diverse Go programs more similar to each other and thus easier to learn.

Projects built using the go tool use only file and identifier names and an occasional special comment to determine all the libraries, executables, tests, becnhmarks, examples, platform-specific variants, and documentation for a project; the Go source itself contains the build-specification.
