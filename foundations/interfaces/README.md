# Interfaces

An *interface type* is an abstract type. It doesn't expose the representation or internal structure of its values; it reveals only some of their methods.

```go
package io

type Reader interface {
  Read(p []byte) (n int, err error)
}

type Closer interface {
  Close() error
}

type ReadWriter interface {
  Reader
  Writer
}

type ReadWriteCloser interface {
  Reader
  Writer
  Closer
}
```

## Interface Satisfaction

```golang
var w io.Writer
w = os.Stdout // OK
w = new(bytes.Buffer) // OK
w = time.Second // compile error

var wrc io.ReadWriteCloser
rwc = os.Stdout // OK
rwc = new(bytes.Buffer) // compile error

var any interface{}
```

## Type Assertions

```golang
var w io.Writer
w = os.Stdout
f := w.(*os.File) // success f == os.Stdout
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```

### Type switch

```golang
switch x.(type) {
  case nil: // ...
  case int, uint: // ...
  case bool: // ...
  case string: // ...
  default: // ...
}
```
