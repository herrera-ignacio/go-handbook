# Composite Types

* Arrays
* Slices
* Maps
* Structs

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

## Maps

In Go, a *map* is a reference to a hash table, and a map type is written `map[K]V`, where `K` and `V` are the types of its keys and values. All of the keys in a given map are of the same type, and all of the values are of the same type. The key type `K` must be comparable using `==`.

> A Hash Map is an unordered collection of key/value pairs in which all the keys are distinct, and the value associated with a given key can be retrieved, updated, or removed using a constant number of key comparisons on the average, no matter how large the hash table.

```go
// Option 1
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34

// Option 2
ages := map[string]int{
  "alice": 31,
  "charlie": 34,
}
```

## Struct

A *struct* is an aggregate data type that groupes together zero or more named values of arbitrary types as a single entity. Each value is called a *field*. Fields are collected into a single entity that can be copied as unit, passed to functions and returned by them, stored in arrays, and so on.

```go
type Employee struct {
  ID int
  Name string
  Address string
  DoB time.Time
  Position string
  Salary int
  ManagerID int
}

// Struct literal
type Point struct{ X, Y int} 

p := Point{1, 2}
```

> See [example](../../../demos/basics/structs/treesort.go)
