# Basic Data Types

- [Basic Data Types](#basic-data-types)
  - [Numbers](#numbers)
    - [Integers](#integers)
    - [Floating Point](#floating-point)
    - [Complex](#complex)
  - [Booleans](#booleans)
  - [Strings](#strings)
  - [Constants](#constants)


## Numbers

Each numeric type determines the syze and signedness of its values.

### Integers

Go provides both **signed and unsigned integer arithmetic**. There are four distinct sizes of signed integers: `int8`, `int16`, `int32`, and `int64`, and corresponding unsigned versions `uint8`, `uint16`, `uint32`, `uint64`. There are also two types calles just `int` and `uint` that are the **natural and most efficient size** for the signed and unsigned integers on a particular platform.

The type `rune` is a synonym for `int32` and conventionally indicates that a value is a Unicode code point. Similarly the type `byte` is a synonym for `uint8`, and emphasizes that the value is a piece of raw data rather than a small numeric quantity.

Finally, there is an unsigned integer type `uintptr`, whose width is not specified but is sufficient to hold all the bits of a pointer value. The `uintptr` type is used for low-level programming, such as the boundary of a Go program with a C library or an operating system.

### Floating Point

Go provides two sizes of floating-point numbers, `float32` and `float64`. Their arithmetic properties are governed by the IEEE 754 standard implemented by all modern CPUs.

A `float32` provides approximately six decimal digits of precision, whereas a `float64` provides about 15 digits; `float64` should be preferred for most purposes because `float32` computations accumulate error rapidly unless one is quite careful, and the [smallest possible integer that cannot be exactly represented as a `float32`](https://stackoverflow.com/a/49758066/9248356) is not large:

```go
var f float32 = 16777216 // 1 << 24
fmt.Println(f == f+1) // "true"!
```

Floating-point numbers can be written literally using decimals:

```go
const e = 2.71828 // aprox
```

Very small or large numbers are better written in scientific notation:

```go
const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34
```

### Complex

Go provides two sizes of complex numers, `complex64` and `complex128`, whose components are `float32` and `float64` respectvely. The built-in function `complex` creates a complex number from its real and imaginary components, and the built-in `real` and `imag` functions extract those components:

```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y) // "(-5+10i)"
fmt.Println(real(x*y)) // -5
fmtPrintln(imag(x*y)) // 10
```

The declarations can be simplified:

```go
x := 1 + 2i
y := 3 + 4i
```

## Booleans

A value of type `bool`, or `boolean`, has only two possible values, `true` and `false`. The condition in `if` and `for` statementsa are booleans, and comparison operators produce a boolean result. The unary operator `!` is a logical negation.

## Strings

A string is an **immutable sequence of bytes**. Text strings are conventionally interpreted as UTF-8 encoded sequences of Unicode code points (runes).

The built-in `len` function returns the number of bytes (not runes) in a string, and the operation `s[i]` returns the *i*-th byte of string *s*.

```go
s := "hello, world"
fmt.Println(len(s)) // 12
fmt.Println(s[0], s[7]) // 104 119 ('h' and 'w')
```

So the *i*-th byte of a string is not necessarily the *i*-th *character* of a string, because the encoding of a non-ASCII code point requires two or more bytes.

Four standard packages are particularly important for manipuilating strings: `bytes`, `strings`, `strconv`, and `unicode`.

* `strings`: searching, replacing, comparing, trimming, splitting, and joining strings.

* `bytes`: manipulating slices of bytes, of type `[]byte`, which share some properties with strings. Building up strings incrementally can be done efficiently using the `bytes.Buffer` type.

* `strconv`: provides functions for converting boolean, integer, and floating-point values to and from their string representations, and functions for quoting and unquoting strings.

* `unicode`: provides functions like `IsDigit`, `IsLetter`, `IsUpper` and `IsLower` for classifiying runes.

## Constants

Constants refer to *fixed values* that the program may not alter during its execution. These fixed values are also called *literals*. They are treated just like regular variables except that their values cannot be modified after their initial definition.

Constants can be of any of the *basic data types* like an integrer constant, a floating constant, a character constant, or a string literal. There are also enumartion constants as well.

You can use `const` prefix to declara constants with a specific type as follows:

```go
package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int

   area = LENGTH * WIDTH
   fmt.Printf("value of area : %d", area)   
}
```
