# Declarations

A *declaration* names a program entity and specifies some or all of its properties.

There are four major kinds of declarations: `var`, `const`, `type`, and `func`.

A function declaration has a name, a list of parameters, an optional list of results, and the function body. The result list is omitted if the function does not return anything.

## Variables

A `var` declaration creates a variable of a particular type, attaches a name to it, and sets its initial value.

```go
var <name> <type> = <expression>
```

Either the `type` or the `= expression` part may be omitted, but not both. If the expression is omitted, the initial value is the *zero value* for the type, which is `0` for numbers, `false` for booleans, `""` for strings, and `nil` for interfaces and reference types.

> In Go there is no such thing as an unitialized variable.

It is possible to declare and optionally initialize a set of variables in a single declaration.

```go
var i, k, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
var f, err = os.Open(name) // file and error
```

### Short Variable Declarations

Within a function, an alternate form called a *short variable declaration* may be used.

```go
// name := expression
anim := gif.GIF{LoopCount: nframes}
freq := rand.Float64()
t := 0.0
i, j := 0, 1
i, j = j, i // swap values
```

> A `var` declaration tends to be reserved for local variables that need an explicit type that differs from that of the initializer expression, or for when the variable will be assigned a value later and its initial value is unimportant.

### Pointers

A *pointer* value is the *address* of a variable. With a pointer, we can read or update the value of a variable indirectly, without using or even knowing the name of a variable, if indeed it has a name.

If a variable is declared `var x int`, the expression `&x` ("address of x") yields a pointer to an integer variable, that is, a value of type `*int`, which is pronunced "pointer to int". If this vallue is called `p`, we say "p points to x". The variable to which `p` points is written `*p`, this expression yields the value of that variable, an `int`.

```go
x := 1
p := &x
fmt.Println(*p) // 1
*p = 2
fmt.Println(x) // 2
```

The zero value for a pointer of any type is `nil`. Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.

```golang
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil); // "true false false"
```

### The `new` Function

Another way to create a variable is to use the built-in function `new`, that creates an *unnamed variable* of type `T`, initializes it to the zero value of `T`, and returns its address, which is a value of type `*T`.

```go
p := new(int) // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // 0
*p = 2
fmt.Println(*p) // 2
```

> Two variables, whose type carries no information and is therefore of size zero, such as `struct{}` or `[0]int`, may, depending on the implementation, have the same address.

### Lifetime

The *lifetime* of a variable is the interval of time during which it exists as the program executes.

* *Package-level variable*: entire execution of the program.

* *Local variable*: dynamic lifetime, a new instance gets created each time the declaration statement is executed, and the variable lives on until it becomes *unreachable*, at which point its storage may be recycled.

## Types
