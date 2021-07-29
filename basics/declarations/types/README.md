# Types

The type declaration defines a new *named type* that has the same *underyling type* as an existing type: `type name underlying-type`. And for every type `T`, there is a corresponding *conversion operation* `T(x)`.

To illustrate type declarations, let's turn the different temperature scales into different types:

```golang
type Celcius float64
type Farenheit float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC     Celsius = 0
  BoilingC      Celsius = 100
)

func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32) }

func FToC(f Farenheit) Celsius { return Celsius((f - 32) * 5/9) }
```
