package main

import (
	"bufio"
	"fmt"
	"golangBook/chapterTwo/tempconv"
	"os"
	"strconv"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	ucl := strings.ToLower(scanner.Text())
	// 	fmt.Println(ucl)
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "error:", err)
	// 	os.Exit(1)
	// }

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("Hits this line")
		for _, arg := range os.Args[1:] {
			fmt.Println("Reaches this point, at the beginning ")
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

			fmt.Println("Reaches this point, at the end")
		}
	}
}
