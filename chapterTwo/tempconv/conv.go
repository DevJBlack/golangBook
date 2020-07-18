package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius((k - 273.15)) }

// FeToM converts feet to meter
func FeToM(fe Feet) Meter { return Meter(fe / 3.281) }

// MToFe converts meters to feet
func MToFe(m Meter) Feet { return Feet(m * 3.281) }

// PToKilo converts Pounds into Kilo
func PToKilo(p Pound) Kilo { return Kilo(p / 2.205) }

// KiloToP converts Kilo into pounds
func KiloToP(kilo Kilo) Pound { return Pound(kilo * 2.205) }
