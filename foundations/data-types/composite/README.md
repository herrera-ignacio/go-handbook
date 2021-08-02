# Composite Types

* Arrays
* Slices

## Arrays

An array is a **fixed-length sequence** of zero or more elements of a **particular type**. Because of their fixed length, arrays are rarely used directly in Go. *Slices* are more versatile.

```go
var a [3]int             // array of 3 integers
b := [...]int{1,2,3}
fmt.Println(a[0]);       // print the first element
fmt.Println(a[len(a)-1]) // print the last element, a[2]

// Print the indices and elements
for i, v := range a {
  fmt.Printf("%d %d\n", i, v)
}

// Print the elements only
for _, v := range a {
  fmt.Printf("%d\n", v)
}
```

## Slices

*Slices* represent **variable-length** sequences whose elements all have the same type. A slice type is written `[]T` , where the elements have type `T`; it looks like an array type without a size.

*Slices* give access to a subsequence of the elements of an array, which is known as the slice's *underlying array*. A slice has three components: a pointer, a length, and a capacity. Multiple slices can share the same underlying array and may refer to overlapping parts of the array.

```go
// underlying array
months := [...]string{1: "January", /* ... */, 12: "December"}

// overlapping slices for the second quarter and the northern summer
Q2 := months[4:7]
summer := months[6:9]

fmt.Println(Q2) // ["April" "May" "June"]
fmt.Println(summer) // ["June" "July" "August"]

// Check June is included in both as example
for _, s := range summer {
  for _, q := range Q2 {
    if s == q {
      fmt.Printf("%s appears in both\n", s)
    }
  }
}
```
