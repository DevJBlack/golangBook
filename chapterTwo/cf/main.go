// Cf converts its numeric arugment to Celsius and Fahrenheit.package main

package main

import (
	"fmt"
	"golangBook/chapterTwo/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("A difference of 1K has the same magantude as 1C: %s = %s\n", k, tempconv.KToC(k))

		fe := tempconv.Feet(t)
		m := tempconv.Meter(t)
		p := tempconv.Pound(t)
		kilo := tempconv.Kilo(t)

		fmt.Printf("Feet to Meters: %s = %s, Meters to Feet: %s = %s\n", fe, tempconv.FeToM(fe), m, tempconv.MToFe(m))

		fmt.Printf("Pounds to Kilo: %s = %s, Kilo to Pounds: %s = %s\n", p, tempconv.PToKilo(p), kilo, tempconv.KiloToP(kilo))

	}
}
