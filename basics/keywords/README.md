# Reserved words

## Keywords

Go has 25 *keywords*:

```go
break
case
chan
const
continue
default
defer
else
fallthrough
for
func
go
goto
if
import
interface
map
package
range
return
select
struct
switch
type
var
```

## Predeclared

In addition, there are about three dozen *predeclared* names for built-in constants, types, and functions. These names are *not reserved*, so you may use them in declarations, but beware of the potential for confusion.

### Constants

```go
true
false
iota
nil
```

### Types

```go
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
```

### Functions

```go
make
len
cap
new
append
copy
close
delete
complex
real
imag
panic
recover
```
