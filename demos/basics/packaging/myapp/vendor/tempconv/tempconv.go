// Performs Celcius and Farenheit conversions
package tempconv

import "fmt"

type Celsius float64
type Farenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%g°C", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%g°F", f) }
