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
	}
}

// Feet and meters and pounds and kilo
// Feet to Meters = divide the length value by 3.281
// Meters to Feet = multiply the length value by 3.281

// Pound to Kilo = divide the mass value by 2.205
// Kilo to Pound = multiply the mass value by 2.205
