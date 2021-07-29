# Scope

A declaration associates a name with a program entity, such as a function or a variable. The *scope* of a declaration is the **part of the source code where a use of the declared name refers to that declaration**.

A declaration's **lexical scope determines its scope**.

* Declarations outside any function, that is, at *package level*, can be referred to from any file in the same package.
* Imported packages are declared at the *file level*, so they can be referred to from the same file, but not form another file in the same package.

At the package level, **the order in which declarations appeaar has no effect on their scope**, so a declaration may refer to itself or to another that follows it, letting us declare recursive or mutually recursive types and functions.

## Lifetime

**Don't confuse scope with lifetime**. The scope of a declaration is a region of the program text; **it is a compile-time property**. The lifetime of a variable is the range of time during execution when the variable can be referred to by other parts of the program; it is a run-time property.

## Syntactic & Lexical Blocks

A *synatic block* is a sequence of statements enclosed in braces like those that surround the body of a function or a loop.

A name declared inside a syntactic block is not visible outside that block. The block **encloses its declarations and determines their scope**.

We can generalize this notion of blocks to include other groupings of declarations that are not explicitly surrounded entire by braces in the source code; we'll call them all *lexical blocks*.
