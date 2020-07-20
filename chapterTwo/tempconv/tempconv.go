// Package tempconv performs converstions for multiple of examples
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Pound float64
type Kilo float64
type Feet float64
type Meter float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	MaxLoadP      Pound   = 1000
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.3f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func (fe Feet) String() string {
	return fmt.Sprintf("%.2f", fe)
}

func (m Meter) String() string {
	return fmt.Sprintf("%.4f", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%.2f", p)
}

func (kilo Kilo) String() string {
	return fmt.Sprintf("%.3f", kilo)
}
